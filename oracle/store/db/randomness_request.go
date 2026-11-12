package db

import (
	"fmt"
	"gorm.io/gorm"
	"oracle/models/database"
)

func (d *ChainDB) InsertNewRequest(keyHash string,
	sender string, requestId string, status int,
	txHash string, gasUsed uint64, gasPrice uint64,
	fee uint64) (err error) {
	err = d.Omit("FulfilTx").Create(&database.RandomnessRequest{
		ChainId:             d.ChainID,
		KeyHash:             keyHash,
		Sender:              sender,
		RequestId:           requestId,
		RequestTxHash:       txHash,
		RequestGasUsed:      gasUsed,
		RequestGasPrice:     gasPrice,
		Fee:                 fee,
		Status:              status,
		FulfillmentAttempts: 0,
	}).Error
	return
}

func (d *ChainDB) UpdateRequestBlockAndSeed(requestId string, blockHash string, seed string, blockNumber uint64) error {
	req := database.RandomnessRequest{}
	err := d.Where("request_id = ?", requestId).First(&req).Error
	if err != nil {
		return err
	}
	req.RequestBlockHash = blockHash
	req.RequestBlockNumber = blockNumber
	req.Seed = seed

	err = d.Save(&req).Error

	return err
}

func (d *ChainDB) UpdateRequestStatus(requestId string, status int, statusReason string) error {
	req := database.RandomnessRequest{}
	err := d.Where("request_id = ?", requestId).First(&req).Error
	if err != nil {
		return err
	}
	req.Status = status
	req.StatusReason = statusReason

	// if the Tx failed as a result of "nonce too low" or "out of gas" increment the num attempts made
	if status == database.REQUEST_STATUS_TX_FAILED {
		req.FulfillmentAttempts = req.FulfillmentAttempts + 1
	}
	err = d.Save(&req).Error

	return err
}

func (d *ChainDB) UpdateFulfilmentSent(requestId string, status int, txHash string, blockNum uint64) error {
	req := database.RandomnessRequest{}
	err := d.Where("request_id = ?", requestId).First(&req).Error
	if err != nil {
		return err
	}
	req.LastFulfillSentBlockNumber = blockNum
	req.Status = status
	req.FulfillTxHash = txHash
	req.FulfillmentAttempts = req.FulfillmentAttempts + 1
	err = d.Save(&req).Error

	return err
}

// PageAndLimit normalises pagination: pages count from 1, and a page holds 1-100 records (default 10).
func PageAndLimit(page, pageSize int) (int, int) {
	if page < 1 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return page, pageSize
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, pageSize := PageAndLimit(page, pageSize)
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

func (d *ChainDB) FindByRequestId(requestId string) (database.RandomnessRequest, error) {
	result := database.RandomnessRequest{}
	err := d.Where("request_id = ?", requestId).First(&result).Error
	return result, err
}

func (d *ChainDB) GetPaginatedRequests(page, limit, status int, order string) ([]database.RandomnessRequest, int64, error) {
	var count int64
	var err error

	var requests = []database.RandomnessRequest{}

	if status >= 0 {
		d.Table("randomness_requests").Where("status = ?", status).Count(&count)
		err = d.Scopes(Paginate(page, limit)).Where("status = ?", status).Order(fmt.Sprintf("id %s", order)).Find(&requests).Error
	} else {
		d.Table("randomness_requests").Count(&count)
		err = d.Scopes(Paginate(page, limit)).Order(fmt.Sprintf("id %s", order)).Find(&requests).Error
	}

	return requests, count, err
}

func (d *ChainDB) UpdateFulfillment(requestId string,
	status int, randomness string, blockHash string,
	blockNumber uint64, txHash string, gasUsed uint64, gasPrice uint64) error {

	req := database.RandomnessRequest{}
	err := d.Where("request_id = ?", requestId).First(&req).Error
	if err != nil {
		return err
	}

	req.Status = status
	req.Randomness = randomness
	req.FulfillBlockHash = blockHash
	req.FulfillBlockNumber = blockNumber
	req.FulfillTxHash = txHash
	req.FulfillGasUsed = gasUsed
	req.FulfillGasPrice = gasPrice

	err = d.Save(&req).Error

	return err
}

func (d *ChainDB) GetLast() (database.RandomnessRequest, error) {
	request := database.RandomnessRequest{}
	err := d.Where("status != ?", database.REQUEST_STATUS_SUCCESS).Last(&request).Error
	if err != nil {
		err = d.Last(&request).Error
	}
	return request, err
}

func (d *ChainDB) GetLastXRequests(limit int, consumer string) ([]database.RandomnessRequest, error) {
	var requests = []database.RandomnessRequest{}
	var err error

	where := map[string]interface{}{"status": database.REQUEST_STATUS_SUCCESS}
	if len(consumer) > 0 {
		where = map[string]interface{}{"status": database.REQUEST_STATUS_SUCCESS, "sender": consumer}
	}

	if limit > 0 {
		err = d.Where(where).Order(fmt.Sprintf("id %s", "desc")).Limit(limit).Find(&requests).Error
	} else {
		err = d.Where(where).Order(fmt.Sprintf("id %s", "desc")).Find(&requests).Error
	}
	return requests, err
}

func (d *ChainDB) GetMostGasUsed() (database.RandomnessRequest, error) {
	request := database.RandomnessRequest{}
	err := d.Where("status = ?", database.REQUEST_STATUS_SUCCESS).Order(fmt.Sprintf("fulfill_gas_used %s", "desc")).Limit(1).First(&request).Error
	return request, err
}

func (d *ChainDB) GetLeastGasUsed() (database.RandomnessRequest, error) {
	request := database.RandomnessRequest{}
	err := d.Where("status = ?", database.REQUEST_STATUS_SUCCESS).Order(fmt.Sprintf("fulfill_gas_used %s", "asc")).Limit(1).First(&request).Error
	return request, err
}

func (d *ChainDB) GetByStatus(status int) ([]database.RandomnessRequest, error) {
	var requests = []database.RandomnessRequest{}
	err := d.Where("status = ?", status).Order(fmt.Sprintf("id %s", "asc")).Find(&requests).Error
	return requests, err
}

func (d *ChainDB) GetInitialisedRequests() ([]database.RandomnessRequest, error) {
	var requests = []database.RandomnessRequest{}
	err := d.Where("status = ?", database.REQUEST_STATUS_INITIALISED).Order(fmt.Sprintf("id %s", "asc")).Find(&requests).Error
	return requests, err
}

func (d *ChainDB) GetStuckOrFailedTx() ([]database.RandomnessRequest, error) {
	var requests = []database.RandomnessRequest{}
	err := d.Where("status = ? OR status = ?", database.REQUEST_STATUS_SENT, database.REQUEST_STATUS_TX_FAILED).Order(fmt.Sprintf("id %s", "asc")).Find(&requests).Error
	return requests, err
}

func (d *ChainDB) GetJobs() ([]database.RandomnessRequest, error) {
	var requests = []database.RandomnessRequest{}
	err := d.Where("status = ? OR status = ? OR status = ?", database.REQUEST_STATUS_INITIALISED, database.REQUEST_STATUS_SENT, database.REQUEST_STATUS_TX_FAILED).Order(fmt.Sprintf("id %s", "asc")).Find(&requests).Error
	return requests, err
}

func (d *ChainDB) GetDistinctConsumers(consumer string) ([]database.RandomnessRequest, error) {
	var requests = []database.RandomnessRequest{}
	var err error

	if len(consumer) > 0 {
		err = d.Distinct("sender").Where("sender = ?", consumer).Order("sender desc").Find(&requests).Error
	} else {
		err = d.Distinct("sender").Order("sender desc").Find(&requests).Error
	}
	return requests, err
}
