package models

type KeyStorageKeyModel struct {
	Account       string `json:"account"`
	CipherPrivate string `json:"cipherprivate"`
}

func (d KeyStorageKeyModel) GetAccount() string {
	return d.Account
}
func (d KeyStorageKeyModel) GetCipherPrivate() string {
	return d.CipherPrivate
}

type KeyStorageModel struct {
	Key *[]KeyStorageKeyModel
}

func (d KeyStorageModel) GetKey() *[]KeyStorageKeyModel {
	return d.Key
}
