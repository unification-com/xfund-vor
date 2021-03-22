package keystorage

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"oracle/models"
	"os"
)

type Keystorage struct {
	log      *logrus.Logger
	File     *os.File
	KeyStore *models.KeyStorageModel
}

func NewKeyStorage(log *logrus.Logger, filePath string) (*Keystorage, error) {
	var err error
	keystoreFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "keystorage",
			"function": "NewKeyStorage",
			"action":   "reading file",
			"result":   err.Error(),
		}).Error()
		return nil, err
	}

	data, err := ioutil.ReadAll(keystoreFile)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "keystorage",
			"function": "NewKeyStorage",
			"action":   "init KeyStore object",
			"result":   err.Error(),
		}).Error()
		return nil, err
	}

	var keyStore models.KeyStorageModel

	err = json.Unmarshal(data, &keyStore)
	if err != nil {
		log.WithFields(logrus.Fields{
			"package":  "keystorage",
			"function": "NewKeyStorage",
			"action":   "unmarshal json from file",
			"result":   err.Error(),
		}).Error()
	}

	return &Keystorage{
		log:      log,
		File:     keystoreFile,
		KeyStore: &keyStore,
	}, err
}

func (d *Keystorage) GetFirst() models.KeyStorageKeyModel {
	return (*d.KeyStore.GetKey())[0]
}
