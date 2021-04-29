package db

import (
	"fmt"
	"gorm.io/gorm"
	"oracle/models/database"
)

func (d *DB) InsertNewRequest(keyHash string, seed string,
	sender string, requestId string, status int, blockHash string,
	blockNumber uint64, txHash string, gasUsed uint64, gasPrice uint64,
	fee uint64) (err error) {
	err = d.Omit("FulfilTx").Create(&database.RandomnessRequest{
		KeyHash:            keyHash,
		Seed:               seed,
		Sender:             sender,
		RequestId:          requestId,
		RequestBlockHash:   blockHash,
		RequestBlockNumber: blockNumber,
		RequestTxHash:      txHash,
		RequestGasUsed:     gasUsed,
		RequestGasPrice:    gasPrice,
		Fee:                fee,
		Status:             status,
	}).Error
	return
}

func (d *DB) UpdateRequestStatus(requestId string, status int, statusReason string) (error) {
	req := database.RandomnessRequest{}
	err := d.Where("request_id = ?", requestId).First(&req).Error
	if err != nil {
		return err
	}
	req.Status = status
	req.StatusReason = statusReason
	err = d.Save(&req).Error

	return err
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (d *DB) FindByRequestId(requestId string) (database.RandomnessRequest, error) {
	result := database.RandomnessRequest{}
	err := d.Where("request_id = ?", requestId).First(&result).Error
	return result, err
}

func (d *DB) GetPaginatedRequests(page, limit, status int, order string) ([]database.RandomnessRequest, int64, error) {
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

func (d *DB) UpdateFulfillment(requestId string,
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

func (d DB) GetLast() (database.RandomnessRequest, error) {
	request := database.RandomnessRequest{}
	err := d.Where("status != ?", database.REQUEST_STATUS_SUCCESS).Last(&request).Error
	if err != nil {
		err = d.Last(&request).Error
	}
	return request, err
}

func (d DB) GetLastXRequests(limit int, consumer string) ([]database.RandomnessRequest, error) {
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

func (d DB) GetMostGasUsed() (database.RandomnessRequest, error) {
	request := database.RandomnessRequest{}
	err := d.Order(fmt.Sprintf("fulfill_gas_used %s", "desc")).Limit(1).First(&request).Error
	return request, err
}

func (d DB) GetLeastGasUsed() (database.RandomnessRequest, error) {
	request := database.RandomnessRequest{}
	err := d.Order(fmt.Sprintf("fulfill_gas_used %s", "asc")).Limit(1).First(&request).Error
	return request, err
}

func (d DB) GetByStatus(status int) ([]database.RandomnessRequest, error) {
	var requests = []database.RandomnessRequest{}
	err := d.Where("status = ?", status).Order(fmt.Sprintf("id %s", "asc")).Find(&requests).Error
	return requests, err
}

func (d DB) GetDistinctConsumers(consumer string) ([]database.RandomnessRequest, error) {
	var requests = []database.RandomnessRequest{}
	var err error

	if len(consumer) > 0 {
		err = d.Distinct("sender").Where("sender = ?", consumer).Order("sender desc").Find(&requests).Error
	} else {
		err = d.Distinct("sender").Order("sender desc").Find(&requests).Error
	}
	return requests, err
}
