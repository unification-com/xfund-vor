// SPDX-License-Identifier: MIT
pragma solidity 0.6.6;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./interfaces/XFundTokenInterface.sol";
import "./interfaces/BlockHashStoreInterface.sol";
import "./VOR.sol";
import "./VORRequestIDBase.sol";
import "./VORConsumerBase.sol";

/**
 * @title VORCoordinator coordinates on-chain verifiable-randomness requests
 * @title with off-chain responses
 */
contract VORCoordinator is Ownable, ReentrancyGuard, VOR, VORRequestIDBase {
    using SafeMath for uint256;
    using Address for address;

    XFundTokenInterface internal xFUND;
    BlockHashStoreInterface internal blockHashStore;

    uint256 public constant EXPECTED_GAS_FIRST_FULFILMENT = 90550;
    uint256 public constant EXPECTED_GAS = 56290;

    uint256 private totalGasDeposits;
    uint256 private gasTopUpLimit;

    constructor(address _xfund, address _blockHashStore) public {
        xFUND = XFundTokenInterface(_xfund);
        blockHashStore = BlockHashStoreInterface(_blockHashStore);
        gasTopUpLimit = 1 ether;
    }

    struct Callback {
        // Tracks an ongoing request
        address callbackContract; // Requesting contract, which will receive response
        // Amount of xFUND paid at request time. Total xFUND = 1e9 * 1e18 < 2^96, so
        // this representation is adequate, and saves a word of storage when this
        // field follows the 160-bit callbackContract address.
        uint96 randomnessFee;
        // Commitment to seed passed to oracle by this contract, and the number of
        // the block in which the request appeared. This is the keccak256 of the
        // concatenation of those values. Storing this commitment saves a word of
        // storage.
        bytes32 seedAndBlockNum;
    }

    struct ServiceAgreement {
        // Tracks oracle commitments to VOR service
        address payable vOROracle; // Oracle committing to respond with VOR service
        uint96 fee; // Minimum payment for oracle response. Total xFUND=1e9*1e18<2^96
        bool providerPaysGas; // True if provider will pay gas
    }

    struct Consumer {
        uint256 amount;
        mapping(address => uint256) providers;
    }

    /* (provingKey, seed) */
    mapping(bytes32 => Callback) public callbacks;
    /* provingKey */
    mapping(bytes32 => ServiceAgreement) public serviceAgreements;
    /* oracle */
    /* xFUND balance */
    mapping(address => uint256) public withdrawableTokens;
    /* provingKey */
    /* consumer */
    mapping(bytes32 => mapping(address => uint256)) private nonces;
    /* (consumer, struct Consumer) */
    mapping(address => Consumer) private gasDeposits;

    mapping(address => bool) public consumerPreviousFulfillment;

    event RandomnessRequest(
        bytes32 keyHash,
        uint256 seed,
        address sender,
        uint256 fee,
        bytes32 requestID
    );

    event NewServiceAgreement(bytes32 keyHash, uint256 fee);

    event ChangeFee(bytes32 keyHash, uint256 fee);

    event RandomnessRequestFulfilled(bytes32 requestId, uint256 output);

    event GasToppedUp(address indexed consumer, address indexed provider, uint256 amount);

    event SetGasTopUpLimit(address indexed sender, uint256 oldLimit, uint256 newLimit);

    event GasRefundedToProvider(address indexed consumer, address indexed provider, uint256 amount);

    event SetProviderPaysGas(address indexed provider, bool providerPays);

    /**
     * @dev getTotalGasDeposits - get total gas deposited in Router
     * @return uint256
     */
    function getTotalGasDeposits() external view returns (uint256) {
        return totalGasDeposits;
    }

    /**
     * @dev getGasTopUpLimit - get gas top up limit
     * @return uint256
     */
    function getGasTopUpLimit() external view returns (uint256) {
        return gasTopUpLimit;
    }

    /**
     * @dev getProviderAddress - get provider address
     * @return address
     */
    function getProviderAddress(bytes32 _keyHash) external view returns (address) {
        return serviceAgreements[_keyHash].vOROracle;
    }

    /**
     * @dev setGasTopUpLimit set the max amount of ETH that can be sent
     * in a topUpGas Tx. Router admin calls this to set the maximum amount
     * a Consumer can send in a single Tx, to prevent large amounts of ETH
     * being sent.
     *
     * @param _gasTopUpLimit amount in wei
     * @return success
     */
    function setGasTopUpLimit(uint256 _gasTopUpLimit) external onlyOwner returns (bool success) {
        require(_gasTopUpLimit > 0, "_gasTopUpLimit must be > 0");
        uint256 oldGasTopUpLimit = gasTopUpLimit;
        gasTopUpLimit = _gasTopUpLimit;
        emit SetGasTopUpLimit(msg.sender, oldGasTopUpLimit, _gasTopUpLimit);
        return true;
    }

    /**
     * @notice Commits calling address to serve randomness
     * @param _fee minimum xFUND payment required to serve randomness
     * @param _oracle the address of the node with the proving key
     * @param _publicProvingKey public key used to prove randomness
     * @param _providerPaysGas true if provider will pay gas
     */
    function registerProvingKey(
        uint256 _fee,
        address payable _oracle,
        uint256[2] calldata _publicProvingKey,
        bool _providerPaysGas
    ) external {
        bytes32 keyHash = hashOfKey(_publicProvingKey);
        address oldVOROracle = serviceAgreements[keyHash].vOROracle;
        require(oldVOROracle == address(0), "please register a new key");
        require(_oracle != address(0), "_oracle must not be 0x0");
        serviceAgreements[keyHash].vOROracle = _oracle;
        // Yes, this revert message doesn't fit in a word
        require(_fee <= 1e9 ether, "you can't charge more than all the xFUND in the world, greedy");
        serviceAgreements[keyHash].fee = uint96(_fee);
        serviceAgreements[keyHash].providerPaysGas = _providerPaysGas;
        emit NewServiceAgreement(keyHash, _fee);
    }

    /**
     * @notice Changes the provider's commission
     * @param _publicProvingKey public key used to prove randomness
     * @param _fee minimum xFUND payment required to serve randomness
     */
    function changeFee(uint256[2] calldata _publicProvingKey, uint256 _fee) external {
        bytes32 keyHash = hashOfKey(_publicProvingKey);
        require(serviceAgreements[keyHash].vOROracle == _msgSender(), "only oracle can change the commission");
        require(_fee <= 1e9 ether, "you can't charge more than all the xFUND in the world, greedy");
        serviceAgreements[keyHash].fee = uint96(_fee);
        emit ChangeFee(keyHash, _fee);
    }

        /**
     * @dev setProviderPaysGas - provider calls for setting who pays gas
     * for sending the fulfillRequest Tx
     * @param _providerPays bool - true if provider will pay gas
     * @return success
     */
    function setProviderPaysGas(uint256[2] calldata _publicProvingKey, bool _providerPays) external returns (bool success) {
        bytes32 keyHash = hashOfKey(_publicProvingKey);
        require(serviceAgreements[keyHash].vOROracle == _msgSender(), "only oracle can change who will pay gas");
        serviceAgreements[keyHash].providerPaysGas = _providerPays;
        emit SetProviderPaysGas(msg.sender, _providerPays);
        return true;
    }

    /**
     * @dev Allows the oracle operator to withdraw their xFUND
     * @param _recipient is the address the funds will be sent to
     * @param _amount is the amount of xFUND transferred from the Coordinator contract
     */
    function withdraw(address _recipient, uint256 _amount) external hasAvailableFunds(_amount) {
        withdrawableTokens[_msgSender()] = withdrawableTokens[_msgSender()].sub(_amount);
        assert(xFUND.transfer(_recipient, _amount));
    }

    /**
     * @notice creates the request for randomness
     *
     * @param _keyHash ID of the VOR public key against which to generate output
     * @param _consumerSeed Input to the VOR, from which randomness is generated
     * @param _feePaid Amount of xFUND sent with request. Must exceed fee for key
     *
     * @dev _consumerSeed is mixed with key hash, sender address and nonce to
     * @dev obtain preSeed, which is passed to VOR oracle, which mixes it with the
     * @dev hash of the block containing this request, to compute the final seed.
     *
     * @dev The requestId used to store the request data is constructed from the
     * @dev preSeed and keyHash.
     */
    function randomnessRequest(
        bytes32 _keyHash,
        uint256 _consumerSeed,
        uint256 _feePaid
    ) external sufficientXFUND(_feePaid, _keyHash) {
        require(address(_msgSender()).isContract(), "request can only be made by a contract");

        xFUND.transferFrom(_msgSender(), address(this), _feePaid);

        uint256 nonce = nonces[_keyHash][_msgSender()];
        uint256 preSeed = makeVORInputSeed(_keyHash, _consumerSeed, _msgSender(), nonce);
        bytes32 requestId = makeRequestId(_keyHash, preSeed);

        // Cryptographically guaranteed by preSeed including an increasing nonce
        assert(callbacks[requestId].callbackContract == address(0));
        callbacks[requestId].callbackContract = _msgSender();

        assert(_feePaid < 1e27); // Total xFUND fits in uint96
        callbacks[requestId].randomnessFee = uint96(_feePaid);

        callbacks[requestId].seedAndBlockNum = keccak256(abi.encodePacked(preSeed, block.number));
        emit RandomnessRequest(_keyHash, preSeed, _msgSender(), _feePaid, requestId);
        nonces[_keyHash][_msgSender()] = nonces[_keyHash][_msgSender()].add(1);
    }

    /**
     * @dev topUpGas consumer contract calls this function to top up gas
     * Gas is the ETH held by this contract which is used to refund Tx costs
     * to the VOR provider for fulfilling a request.
     *
     * To prevent silly amounts of ETH being sent, a sensible limit is imposed.
     *
     * Can only top up for authorised providers
     *
     * @param _provider address of VOR provider
     * @return success
     */
    function topUpGas(address _provider) external payable nonReentrant returns (bool success) {
        uint256 amount = msg.value;
        // msg.sender is the address of the Consumer's smart contract
        address consumer = msg.sender;
        require(address(consumer).isContract(), "only a contract can top up gas");
        require(amount > 0, "cannot top up zero");
        require(amount <= gasTopUpLimit, "cannot top up more than gasTopUpLimit");

        // total held by Router contract
        totalGasDeposits = totalGasDeposits.add(amount);

        // Total held for consumer contract
        gasDeposits[consumer].amount = gasDeposits[consumer].amount.add(amount);

        // Total held for consumer contract/provider pair
        gasDeposits[consumer].providers[_provider] = gasDeposits[consumer].providers[_provider].add(amount);

        emit GasToppedUp(consumer, _provider, amount);
        return true;
    }

    /**
     * @notice Returns the serviceAgreements key associated with this public key
     * @param _publicKey the key to return the address for
     */
    function hashOfKey(uint256[2] memory _publicKey) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(_publicKey));
    }

    /**
     * @notice Called by the node to fulfill requests
     *
     * @param _proof the proof of randomness. Actual random output built from this
     */
    function fulfillRandomnessRequest(bytes memory _proof) public {
        (bytes32 currentKeyHash, Callback memory callback, bytes32 requestId, uint256 randomness) =
            getRandomnessFromProof(_proof);

        // Pay oracle
        address payable oracle = serviceAgreements[currentKeyHash].vOROracle;
        withdrawableTokens[oracle] = withdrawableTokens[oracle].add(callback.randomnessFee);

        address gasPayer = (serviceAgreements[currentKeyHash].providerPaysGas) ? oracle : callback.callbackContract;

        // Forget request. Must precede callback (prevents reentrancy)
        delete callbacks[requestId];
        uint256 gasLeftStart = gasleft();
        callBackWithRandomness(requestId, randomness, callback.callbackContract);
        uint256 gasUsedToCall = gasLeftStart - gasleft();

        if (gasPayer != oracle) {
            require(refundGas(callback.callbackContract, oracle, gasUsedToCall));
        }

        emit RandomnessRequestFulfilled(requestId, randomness);
    }

    // Offsets into fulfillRandomnessRequest's _proof of various values
    //
    // Public key. Skips byte array's length prefix.
    uint256 public constant PUBLIC_KEY_OFFSET = 0x20;
    // Seed is 7th word in proof, plus word for length, (6+1)*0x20=0xe0
    uint256 public constant PRESEED_OFFSET = 0xe0;

    function callBackWithRandomness(bytes32 requestId, uint256 randomness, address consumerContract) internal {
        // Dummy variable; allows access to method selector in next line. See
        // https://github.com/ethereum/solidity/issues/3506#issuecomment-553727797
        VORConsumerBase v;
        bytes memory resp = abi.encodeWithSelector(v.rawFulfillRandomness.selector, requestId, randomness);
        // The bound b here comes from https://eips.ethereum.org/EIPS/eip-150. The
        // actual gas available to the consuming contract will be b-floor(b/64).
        // This is chosen to leave the consuming contract ~200k gas, after the cost
        // of the call itself.
        uint256 b = 206000;
        require(gasleft() >= b, "not enough gas for consumer");
        // A low-level call is necessary, here, because we don't want the consuming
        // contract to be able to revert this execution, and thus deny the oracle
        // payment for a valid randomness response. This also necessitates the above
        // check on the gasleft, as otherwise there would be no indication if the
        // callback method ran out of gas.
        //
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = consumerContract.call(resp);
        // Avoid unused-local-variable warning. (success is only present to prevent
        // a warning that the return value of consumerContract.call is unused.)
        (success);
    }

    function getRandomnessFromProof(bytes memory _proof)
        internal
        view
        returns (
            bytes32 currentKeyHash,
            Callback memory callback,
            bytes32 requestId,
            uint256 randomness
        )
    {
        // blockNum follows proof, which follows length word (only direct-number
        // constants are allowed in assembly, so have to compute this in code)
        uint256 blocknumOffset = 0x20 + PROOF_LENGTH;
        // _proof.length skips the initial length word, so not including the
        // blocknum in this length check balances out.
        require(_proof.length == blocknumOffset, "wrong proof length");
        uint256[2] memory publicKey;
        uint256 preSeed;
        uint256 blockNum;
        assembly {
            // solhint-disable-line no-inline-assembly
            publicKey := add(_proof, PUBLIC_KEY_OFFSET)
            preSeed := mload(add(_proof, PRESEED_OFFSET))
            blockNum := mload(add(_proof, blocknumOffset))
        }
        currentKeyHash = hashOfKey(publicKey);
        requestId = makeRequestId(currentKeyHash, preSeed);
        callback = callbacks[requestId];
        require(callback.callbackContract != address(0), "no corresponding request");
        require(callback.seedAndBlockNum == keccak256(abi.encodePacked(preSeed, blockNum)), "wrong preSeed or block num");

        bytes32 blockHash = blockhash(blockNum);
        if (blockHash == bytes32(0)) {
            blockHash = blockHashStore.getBlockhash(blockNum);
            require(blockHash != bytes32(0), "please prove blockhash");
        }
        // The seed actually used by the VOR machinery, mixing in the blockhash
        uint256 actualSeed = uint256(keccak256(abi.encodePacked(preSeed, blockHash)));
        // solhint-disable-next-line no-inline-assembly
        assembly {
            // Construct the actual proof from the remains of _proof
            mstore(add(_proof, PRESEED_OFFSET), actualSeed)
            mstore(_proof, PROOF_LENGTH)
        }
        randomness = VOR.randomValueFromVORProof(_proof); // Reverts on failure
    }

    /**
     * @dev refundGas - private function called by fulfillRequest, when Consumer is expected to pay
     * the gas for fulfilling a request.
     *
     * @param _consumer address of the consumer contract
     * @param _provider address of the VOR provider
     * @param _gasUsedToCall amount of gas consumed calling the Consumer's
     * @return success if the execution was successful.
     */
    function refundGas(address _consumer, address payable _provider, uint256 _gasUsedToCall) private returns (bool){
        // calculate how much should be refunded to the provider
        uint256 baseGas = EXPECTED_GAS_FIRST_FULFILMENT;
        if(consumerPreviousFulfillment[_consumer]) {
            baseGas = EXPECTED_GAS;
        }

        consumerPreviousFulfillment[_consumer] = true;

        uint256 totalGasUsed = baseGas + _gasUsedToCall;
        uint256 ethRefund = totalGasUsed.mul(tx.gasprice);

        // check there's enough
        require(
            gasDeposits[_consumer].providers[_provider] >= ethRefund
            && totalGasDeposits >= ethRefund,
            "Router: not enough ETH to refund"
        );
        // update total held by Router contract
        totalGasDeposits = totalGasDeposits.sub(ethRefund);

        // update total held for consumer contract
        gasDeposits[_consumer].amount = gasDeposits[_consumer].amount.sub(ethRefund);

        // update total held for consumer contract/provider pair
        gasDeposits[_consumer].providers[_provider] = gasDeposits[_consumer].providers[_provider].sub(ethRefund);

        emit GasRefundedToProvider(_consumer, _provider, ethRefund);
        // refund the provider
        Address.sendValue(_provider, ethRefund);
        return true;
    }

    /**
     * @dev Reverts if amount is not at least what was agreed upon in the service agreement
     * @param _feePaid The payment for the request
     * @param _keyHash The key which the request is for
     */
    modifier sufficientXFUND(uint256 _feePaid, bytes32 _keyHash) {
        require(_feePaid >= serviceAgreements[_keyHash].fee, "Below agreed payment");
        _;
    }

    /**
     * @dev Reverts if amount requested is greater than withdrawable balance
     * @param _amount The given amount to compare to `withdrawableTokens`
     */
    modifier hasAvailableFunds(uint256 _amount) {
        require(withdrawableTokens[_msgSender()] >= _amount, "can't withdraw more than balance");
        _;
    }
}