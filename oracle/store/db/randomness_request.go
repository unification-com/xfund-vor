package db

import (
	"oracle/models"
	"oracle/models/requests"
)

type RandomnessRequestStore struct {
	db *DB
}

func NewRandomnessRequestStore(db *DB) *RandomnessRequestStore {
	return &RandomnessRequestStore{db: db}
}

func (d *RandomnessRequestStore) InsertNewRequest(keyHash string, seed string,
	sender string, requestId string, blockHash string,
	blockNumber uint64, requestTxHash string, status string) (err error) {
	err = d.db.Create(&requests.RandomnessRequestStoreModel{
		KeyHash:        keyHash,
		Seed:           seed,
		Sender:         sender,
		RequestId:      requestId,
		ReqBlockHash:   blockHash,
		ReqBlockNumber: blockNumber,
		RequestTxHash:  requestTxHash,
		Status:         status,
	}).Error
	return
}

func (d RandomnessRequestStore) Migrate() (err error) {
	if !d.db.Migrator().HasTable("randomness_request") {
		err = d.db.Migrator().CreateTable(&requests.RandomnessRequestStoreModel{})
	}
	return
}

func (d *RandomnessRequestStore) FindByRequestId(requestId string) (models.IRandomnessRequestStoreModel, error) {
	result := requests.RandomnessRequestStoreModel{}
	err := d.db.Where("request_id = ?", requestId).First(&result).Error
	return result, err
}

func (d *RandomnessRequestStore) UpdateFulfillment(requestId string, fulfillTxHash string,
	status string, gasUsed uint64, blockNumber uint64) error {

	req := requests.RandomnessRequestStoreModel{}
	err := d.db.Where("request_id = ?", requestId).First(&req).Error
	if err != nil {
		return err
	}
	req.Status = status
	req.FulfillTxHash = fulfillTxHash
	req.FulfillGasUsed = gasUsed
	req.FulfillBlockNum = blockNumber

	err = d.db.Save(&req).Error

	return err
}

func (d RandomnessRequestStore) Last() (models.IRandomnessRequestStoreModel, error) {
	request := requests.RandomnessRequestStoreModel{}
	err := d.db.Last(&request).Error
	return request, err
}

