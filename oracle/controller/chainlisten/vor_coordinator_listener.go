package chainlisten

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	"oracle/chaincall"
	"oracle/config"
	"oracle/contracts/vor_coordinator"
	"oracle/models/database"
	"oracle/service"
	"oracle/store/db"
	"oracle/tools/vor"
)

var (
	randomnessRequestTopic = crypto.Keccak256Hash([]byte("RandomnessRequest(bytes32,uint256,address,uint256,bytes32)"))
	requestFulfilledTopic  = crypto.Keccak256Hash([]byte("RandomnessRequestFulfilled(bytes32,uint256)"))
)

const (
	// maxScanBatchesPerPoll bounds one poll's catch-up, so a long backlog doesn't hold up the job queue.
	maxScanBatchesPerPoll = 100

	maxFulfilmentAttempts = 3
	// a request older than this needs its block hash in the block store to be fulfilled (the EVM only
	// serves the last 256 block hashes)
	staleRequestBlocks = 250
	// a retried request older than this has its block hash stored, in case it becomes stale
	storeBlockHashAfterBlocks = 50
	// a sent fulfilment the node has never seen after this many blocks was dropped, and is sent again
	droppedTxBlocks = 50
)

type VORCoordinatorListener struct {
	chain           config.ChainConfig
	contractAddress common.Address
	client          *ethclient.Client
	contractAbi     abi.ABI
	fromBlock       uint64 // the next block to scan for events
	service         *service.Service
	keyHash         [32]byte
	context         context.Context
	logger          *logrus.Logger
}

// NewVORCoordinatorListener listens for the service's chain's requests.
func NewVORCoordinatorListener(service *service.Service, logger *logrus.Logger, ctx context.Context) (*VORCoordinatorListener, error) {
	client, err := chaincall.Dial(service.Chain.EthHTTPHost)
	if err != nil {
		return nil, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(vor_coordinator.VorCoordinatorABI))
	if err != nil {
		return nil, err
	}
	fromBlock, err := startBlock(service)
	if err != nil {
		return nil, err
	}
	// without the key hash every request would look addressed to someone else, and be ignored
	keyHash, err := service.VORCoordinatorCaller.HashOfKey()
	if err != nil {
		return nil, fmt.Errorf("getting the proving key's hash from the coordinator: %w", err)
	}
	if keyHash == ([32]byte{}) {
		return nil, errors.New("the coordinator returned an empty hash for the proving key")
	}
	return &VORCoordinatorListener{
		chain:           service.Chain,
		client:          client,
		contractAddress: common.HexToAddress(service.Chain.VORCoordinatorContractAddress),
		contractAbi:     contractAbi,
		fromBlock:       fromBlock,
		service:         service,
		context:         ctx,
		keyHash:         keyHash,
		logger:          logger,
	}, nil
}

// log starts an entry for function, naming the chain: one process serves several.
func (d *VORCoordinatorListener) log(function string) *logrus.Entry {
	return d.logger.WithFields(logrus.Fields{"package": "chainlisten", "function": function, "chain": d.chain.Name})
}

// StartPoll polls until its context is cancelled.
func (d *VORCoordinatorListener) StartPoll() (err error) {
	d.log("StartPoll").WithFields(logrus.Fields{"action": "begin polling", "from_block": d.fromBlock}).Info()

	// Poll cadence: a base interval plus bounded jitter, so the floor is ALWAYS non-zero. The old code slept
	// time.Sleep(rand.Int31n(base)) which returns [0,base) - frequently 0 - so a stuck job (e.g. a permanently
	// reverting under-fee request) re-ran the loop with no delay, pegging a core and churning the heap (the
	// 2026-06 CPU/RAM runaway). It also discarded every iteration error and never honoured ctx cancellation.
	base := d.chain.CheckDuration
	for {
		if d.context.Err() != nil {
			d.log("StartPoll").WithField("action", "context cancelled - stopping poller").Info()
			return d.context.Err()
		}

		errEvents := d.ProcessIncommingEvents()
		errJobs := d.CheckJobs()

		// Always wait at least `base` seconds; back off further on an RPC/processing error so a struggling
		// or dead endpoint is not hammered in a tight loop.
		wait := base + rand.Int31n(base)
		if errEvents != nil || errJobs != nil {
			entry := d.log("StartPoll").WithField("action", "poll iteration error - backing off")
			if errEvents != nil {
				entry = entry.WithField("err_events", errEvents.Error())
			}
			if errJobs != nil {
				entry = entry.WithField("err_jobs", errJobs.Error())
			}
			entry.Warn()
			wait = base * 2
		}

		select {
		case <-d.context.Done():
			return d.context.Err()
		case <-time.After(time.Duration(wait) * time.Second):
		}
	}
}

// startBlock is where polling resumes: the chain's saved cursor, else its last request's block, else its
// configured first block, else 1.
func startBlock(service *service.Service) (uint64, error) {
	next, err := service.Store.Db.NextBlock()
	if err != nil {
		return 0, fmt.Errorf("reading the scan cursor: %w", err)
	}
	if next > 0 {
		return next, nil
	}
	lastRequest, err := service.Store.Db.GetLast()
	if err != nil && !db.IsNotFound(err) {
		return 0, fmt.Errorf("getting the last request: %w", err)
	}
	if block := lastRequest.GetRequestBlockNumber(); block != 0 {
		return block, nil
	}
	if first := service.Chain.FirstBlockNumber; first != 0 {
		return first, nil
	}
	return 1, nil
}

// blockRange is an inclusive range of blocks.
type blockRange struct{ from, to uint64 }

// scanRanges splits [from, head] into ranges of at most batch blocks, at most maxBatches of them.
func scanRanges(from, head, batch uint64, maxBatches int) []blockRange {
	var ranges []blockRange
	for from <= head && len(ranges) < maxBatches {
		to := min(from+batch-1, head)
		ranges = append(ranges, blockRange{from, to})
		from = to + 1
	}
	return ranges
}

// blocksSince is how many blocks after past current is; 0 when current isn't after it, as a lagging RPC node
// can report a head behind a block it has already served.
func blocksSince(current, past uint64) uint64 {
	if current <= past {
		return 0
	}
	return current - past
}

// ProcessIncommingEvents scans the blocks since the last scan for the coordinator's events, in ranges RPCs
// accept, saving progress after each range. A range is only passed once all its events are recorded, so an
// error means it is scanned again next time rather than skipped.
func (d *VORCoordinatorListener) ProcessIncommingEvents() error {
	head, err := d.client.BlockNumber(d.context)
	if err != nil {
		return fmt.Errorf("getting the head block: %w", err)
	}
	ranges := scanRanges(d.fromBlock, head, d.chain.EventScanBatchBlocks, maxScanBatchesPerPoll)
	for _, r := range ranges {
		logs, err := d.client.FilterLogs(d.context, ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(r.from),
			ToBlock:   new(big.Int).SetUint64(r.to),
			Addresses: []common.Address{d.contractAddress},
		})
		if err != nil {
			return fmt.Errorf("getting the logs of blocks %d-%d: %w", r.from, r.to, err)
		}
		for _, vLog := range logs {
			if err = d.processLog(vLog); err != nil {
				return fmt.Errorf("block %d, log %d: %w", vLog.BlockNumber, vLog.Index, err)
			}
		}
		if err = d.service.Store.Db.SetNextBlock(r.to + 1); err != nil {
			return fmt.Errorf("saving the scan position: %w", err)
		}
		d.fromBlock = r.to + 1
	}
	if len(ranges) == 0 {
		d.log("ProcessIncommingEvents").WithFields(logrus.Fields{"from_block": d.fromBlock, "head": head}).Info("no new blocks")
	} else {
		d.log("ProcessIncommingEvents").WithFields(logrus.Fields{"from_block": ranges[0].from, "to_block": d.fromBlock - 1, "head": head}).Info("scanned")
	}
	return nil
}

func (d *VORCoordinatorListener) processLog(vLog types.Log) error {
	if len(vLog.Topics) == 0 {
		return nil
	}
	switch vLog.Topics[0] {
	case randomnessRequestTopic:
		return d.recordRequest(vLog)
	case requestFulfilledTopic:
		return d.recordFulfilment(vLog)
	}
	return nil
}

// recordRequest adds a randomness request for this oracle's key to the job queue.
func (d *VORCoordinatorListener) recordRequest(vLog types.Log) error {
	event := vor_coordinator.VorCoordinatorRandomnessRequest{}
	if err := d.contractAbi.UnpackIntoInterface(&event, "RandomnessRequest", vLog.Data); err != nil {
		return fmt.Errorf("unpacking RandomnessRequest: %w", err)
	}
	if event.KeyHash != d.keyHash {
		return nil
	}
	requestId := common.Bytes2Hex(event.RequestID[:])
	entry := d.log("recordRequest").WithField("request_id", requestId)

	existing, err := d.service.Store.Db.FindByRequestId(requestId)
	if err == nil {
		entry.WithField("status", existing.GetStatusString()).Info("request already in db")
		return nil
	}
	if !db.IsNotFound(err) {
		return err
	}

	gasUsed, gasPrice := d.txGas(vLog.TxHash)
	entry.Info("new request")
	return d.service.Store.Db.InsertNewRequest(
		common.Bytes2Hex(event.KeyHash[:]),
		event.Sender.Hex(),
		requestId,
		database.REQUEST_STATUS_INITIALISED,
		vLog.TxHash.Hex(),
		gasUsed,
		gasPrice,
		event.Fee.Uint64(),
	)
}

// recordFulfilment marks one of this oracle's requests fulfilled.
func (d *VORCoordinatorListener) recordFulfilment(vLog types.Log) error {
	event := vor_coordinator.VorCoordinatorRandomnessRequestFulfilled{}
	if err := d.contractAbi.UnpackIntoInterface(&event, "RandomnessRequestFulfilled", vLog.Data); err != nil {
		return fmt.Errorf("unpacking RandomnessRequestFulfilled: %w", err)
	}
	requestId := common.Bytes2Hex(event.RequestId[:])
	if _, err := d.service.Store.Db.FindByRequestId(requestId); db.IsNotFound(err) {
		return nil // another oracle's request
	} else if err != nil {
		return err
	}

	gasUsed, gasPrice := d.txGas(vLog.TxHash)
	d.log("recordFulfilment").WithField("request_id", requestId).Info("confirmed request fulfilment")
	return d.service.Store.Db.UpdateFulfillment(
		requestId,
		database.REQUEST_STATUS_SUCCESS,
		event.Output.String(),
		vLog.BlockHash.Hex(),
		vLog.BlockNumber,
		vLog.TxHash.Hex(),
		gasUsed,
		gasPrice,
	)
}

// txGas returns a mined transaction's gas used and the gas price it paid, for the records; 0s if its receipt
// can't be fetched.
func (d *VORCoordinatorListener) txGas(txHash common.Hash) (gasUsed, gasPrice uint64) {
	receipt, err := d.client.TransactionReceipt(d.context, txHash)
	if err != nil {
		d.log("txGas").WithField("tx_hash", txHash.Hex()).Warn(err.Error())
		return 0, 0
	}
	return receipt.GasUsed, d.paidGasPrice(receipt)
}

// paidGasPrice is the gas price a mined transaction paid: its receipt's effective price, or, from nodes that
// don't report one, the price the transaction offered (which is what a legacy one pays). A dynamic-fee
// transaction's GasPrice is only its cap.
func (d *VORCoordinatorListener) paidGasPrice(receipt *types.Receipt) uint64 {
	if receipt.EffectiveGasPrice != nil {
		return receipt.EffectiveGasPrice.Uint64()
	}
	if tx, _, err := d.client.TransactionByHash(d.context, receipt.TxHash); err == nil {
		return tx.GasPrice().Uint64()
	}
	return 0
}

// isPermanentFulfilmentError reports whether a fulfilment send/estimate error is PERMANENT - one that will
// never succeed by retrying, so the request must be failed terminally rather than re-sent forever. The classic
// case is a request whose paid fee is below the oracle's registered fee: the coordinator reverts ("Below agreed
// payment"), which - with gas estimation enabled (GasLimit 0) - surfaces as an "execution reverted" error BEFORE
// broadcast. Genuinely transient errors (nonce too low, RPC blips) are NOT matched and fall through to the normal
// capped retry. This guard is what stops the fee-revert CPU/RAM runaway.
func isPermanentFulfilmentError(err error) bool {
	if err == nil {
		return false
	}
	msg := strings.ToLower(err.Error())
	for _, marker := range []string{"execution reverted", "below agreed payment", "always failing transaction"} {
		if strings.Contains(msg, marker) {
			return true
		}
	}
	return false
}

// requestEvent finds requestId's RandomnessRequest event in the receipt of the transaction that made it.
func (d *VORCoordinatorListener) requestEvent(requestId string, receipt *types.Receipt) (*vor_coordinator.VorCoordinatorRandomnessRequest, error) {
	for _, vLog := range receipt.Logs {
		if len(vLog.Topics) == 0 || vLog.Topics[0] != randomnessRequestTopic {
			continue
		}
		event := vor_coordinator.VorCoordinatorRandomnessRequest{}
		if err := d.contractAbi.UnpackIntoInterface(&event, "RandomnessRequest", vLog.Data); err != nil {
			return nil, fmt.Errorf("unpacking RandomnessRequest: %w", err)
		}
		if common.Bytes2Hex(event.RequestID[:]) == requestId {
			return &event, nil
		}
	}
	return nil, errors.New("the request's transaction has no RandomnessRequest event for it")
}

// fulfil sends the fulfilment of requestId and records the outcome. Every path changes the request's status,
// so no request is processed the same way twice.
func (d *VORCoordinatorListener) fulfil(requestId string, requestTxReceipt *types.Receipt, currentBlockNum uint64) error {
	entry := d.log("fulfil").WithField("request_id", requestId)
	entry.Info("begin request fulfilment")
	store := d.service.Store.Db

	event, err := d.requestEvent(requestId, requestTxReceipt)
	if err != nil {
		entry.Error(err.Error())
		return store.UpdateRequestStatus(requestId, database.REQUEST_STATUS_FULFILMENT_FAILED, err.Error())
	}
	seed, err := vor.BigToSeed(event.Seed)
	if err != nil {
		entry.Error(err.Error())
		return store.UpdateRequestStatus(requestId, database.REQUEST_STATUS_FULFILMENT_FAILED, err.Error())
	}
	requestBlockHash := requestTxReceipt.BlockHash
	requestBlockNum := requestTxReceipt.BlockNumber.Uint64()
	if err = store.UpdateRequestBlockAndSeed(requestId, requestBlockHash.Hex(), hexutil.EncodeBig(event.Seed), requestBlockNum); err != nil {
		return err
	}

	fTx, err := d.service.FulfillRandomness(seed, requestBlockHash, requestBlockNum)
	if err != nil {
		// A permanent revert (e.g. the request's paid fee is below the registered fee) is caught here by gas
		// estimation BEFORE broadcast - fail it terminally so it drops out of the job queue, instead of
		// re-sending it forever (the fee-revert CPU/RAM runaway). Transient errors (nonce/RPC) fall through to
		// TX_FAILED and the normal capped retry.
		permanent := isPermanentFulfilmentError(err)
		status := database.REQUEST_STATUS_TX_FAILED
		if permanent {
			status = database.REQUEST_STATUS_FULFILMENT_FAILED
		}
		entry.WithField("permanent", permanent).Error(err.Error())
		return store.UpdateRequestStatus(requestId, status, err.Error())
	}
	entry.WithFields(logrus.Fields{
		"request_tx_hash":    requestTxReceipt.TxHash.Hex(),
		"fulfill_tx_hash":    fTx.Hash().Hex(),
		"request_block_hash": requestBlockHash.Hex(),
		"request_block_num":  requestBlockNum,
	}).Info("fulfill tx sent")
	if err = store.UpdateFulfilmentSent(requestId, database.REQUEST_STATUS_SENT, fTx.Hash().Hex(), currentBlockNum); err != nil {
		return fmt.Errorf("fulfilment %s was sent but couldn't be recorded, so may be sent again: %w", fTx.Hash().Hex(), err)
	}
	return nil
}

// failedAttempt describes a fulfilment attempt that failed, for the failed_fulfilments history.
type failedAttempt struct {
	txHash            string
	gasUsed, gasPrice uint64
	reason            string
}

// retry records a failed fulfilment attempt and sends the fulfilment again, unless the request has had too
// many attempts or has become too old to fulfil.
func (d *VORCoordinatorListener) retry(request database.RandomnessRequest, requestTxReceipt *types.Receipt, currentBlockNum uint64, failed failedAttempt) error {
	requestId := request.GetRequestId()
	store := d.service.Store.Db
	if err := store.InsertNewFailedFulfilment(requestId, failed.txHash, failed.gasUsed, failed.gasPrice, failed.reason); err != nil {
		return err
	}
	if request.GetFulfillmentAttempts() >= maxFulfilmentAttempts {
		d.log("retry").WithFields(logrus.Fields{"request_id": requestId, "num_attempts": request.GetFulfillmentAttempts()}).Info("too many failed attempts")
		return store.UpdateRequestStatus(requestId, database.REQUEST_STATUS_FULFILMENT_FAILED, "too many failed attempts")
	}

	requestBlockNum := requestTxReceipt.BlockNumber.Uint64()
	age := blocksSince(currentBlockNum, requestBlockNum)
	foundHash, _, bhErr := d.service.VORCoordinatorCaller.GetBlockHashFromBlockStore(requestBlockNum)
	if age > staleRequestBlocks && (!foundHash || bhErr != nil) {
		d.log("retry").WithField("request_id", requestId).Info("request too old and block hash not in block store")
		return store.UpdateRequestStatus(requestId, database.REQUEST_STATUS_FULFILMENT_FAILED, "request too old and block hash not in block store")
	}
	if age > storeBlockHashAfterBlocks && !foundHash {
		if err := d.recordBlockHash(requestBlockNum, requestTxReceipt.BlockHash.Hex()); err != nil {
			return err
		}
	}
	return d.fulfil(requestId, requestTxReceipt, currentBlockNum)
}

// recordBlockHash stores the request block's hash in the block store contract, so the request can still be
// fulfilled once the block is too old for the EVM to serve its hash.
func (d *VORCoordinatorListener) recordBlockHash(blockNumber uint64, blockHash string) error {
	entry := d.log("recordBlockHash").WithFields(logrus.Fields{"block_num": blockNumber, "block_hash": blockHash})
	bhTx, err := d.service.VORCoordinatorCaller.StoreBlockHash(blockNumber)
	if err != nil {
		// not fatal: the retry can still succeed while the block is young enough
		entry.Error(err.Error())
		return nil
	}
	entry.Info("store tx sent")
	return d.service.Store.Db.InsertNewStoredBlock(blockHash, blockNumber, bhTx.Hash().Hex())
}

// checkSent follows up a sent fulfilment: retried if it reverted or was dropped, left alone while pending or
// once mined successfully (the RandomnessRequestFulfilled event then marks the request done).
func (d *VORCoordinatorListener) checkSent(request database.RandomnessRequest, requestTxReceipt *types.Receipt, currentBlockNum uint64) error {
	requestId := request.GetRequestId()
	entry := d.log("checkSent").WithFields(logrus.Fields{"request_id": requestId, "tx_hash": request.GetFulfillTxHash()})

	sinceSent := blocksSince(currentBlockNum, request.GetLastFulfillSentBlockNumber())
	if sinceSent < 2 {
		entry.WithField("block_diff", sinceSent).Info("not enough blocks since last sent. Wait.")
		return nil
	}

	fulfilTxHash := common.HexToHash(request.GetFulfillTxHash())
	_, isPending, err := d.client.TransactionByHash(d.context, fulfilTxHash)
	if errors.Is(err, ethereum.NotFound) {
		if sinceSent < droppedTxBlocks {
			entry.Info("tx not seen yet")
			return nil
		}
		entry.Warn("tx dropped - sending again")
		return d.service.Store.Db.UpdateRequestStatus(requestId, database.REQUEST_STATUS_TX_FAILED, "fulfilment transaction dropped")
	}
	if err != nil {
		return err
	}
	if isPending {
		entry.Info("tx still pending - ignore")
		return nil
	}

	fulfillReceipt, err := d.client.TransactionReceipt(d.context, fulfilTxHash)
	if err != nil {
		return err
	}
	if fulfillReceipt.Status == types.ReceiptStatusSuccessful {
		entry.Info("tx success. wait for RandomnessRequestFulfilled event")
		return nil
	}
	// todo - try to get actual revert reason
	return d.retry(request, requestTxReceipt, currentBlockNum, failedAttempt{
		txHash:   request.GetFulfillTxHash(),
		gasUsed:  fulfillReceipt.GasUsed,
		gasPrice: d.paidGasPrice(fulfillReceipt),
		reason:   "transaction reverted",
	})
}

func (d *VORCoordinatorListener) preProcessJob(request database.RandomnessRequest, currentBlockNum uint64) error {
	requestId := request.GetRequestId()
	entry := d.log("preProcessJob").WithFields(logrus.Fields{"request_id": requestId, "status": request.GetStatusString()})
	entry.Info("preprocess job")

	requestTxReceipt, err := d.client.TransactionReceipt(d.context, common.HexToHash(request.GetRequestTxHash()))
	if errors.Is(err, ethereum.NotFound) {
		entry.WithField("request_tx", request.GetRequestTxHash()).Warn("request tx not found (yet)")
		return nil
	}
	if err != nil {
		return err
	}

	switch request.GetStatus() {
	case database.REQUEST_STATUS_INITIALISED:
		requestBlockDiff := blocksSince(currentBlockNum, requestTxReceipt.BlockNumber.Uint64())
		if requestBlockDiff < d.chain.WaitConfirmations {
			entry.WithFields(logrus.Fields{"block_diff": requestBlockDiff, "wait_config": d.chain.WaitConfirmations}).
				Info("not enough block confirmations to fulfill request")
			return nil
		}
		return d.fulfil(requestId, requestTxReceipt, currentBlockNum)
	case database.REQUEST_STATUS_TX_FAILED:
		return d.retry(request, requestTxReceipt, currentBlockNum, failedAttempt{
			txHash:   request.GetFulfillTxHash(),
			gasUsed:  request.GetFulfillGasUsed(),
			gasPrice: request.GetFulfillGasPrice(),
			reason:   request.GetStatusReason(),
		})
	case database.REQUEST_STATUS_SENT:
		return d.checkSent(request, requestTxReceipt, currentBlockNum)
	}
	return nil
}

// CheckJobs moves every queued request on a step, returning the errors of any that couldn't be.
func (d *VORCoordinatorListener) CheckJobs() error {
	d.log("CheckJobs").Info("check job queue")
	currentBlockNum, err := d.client.BlockNumber(d.context)
	if err != nil {
		return fmt.Errorf("getting the head block: %w", err)
	}
	requests, err := d.service.Store.Db.GetJobs()
	if err != nil {
		return fmt.Errorf("getting the job queue: %w", err)
	}
	var errs []error
	for _, request := range requests {
		if err = d.preProcessJob(request, currentBlockNum); err != nil {
			d.log("CheckJobs").WithField("request_id", request.GetRequestId()).Error(err.Error())
			errs = append(errs, fmt.Errorf("request %s: %w", request.GetRequestId(), err))
		}
	}
	return errors.Join(errs...)
}
