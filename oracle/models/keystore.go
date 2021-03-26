package models

type KeyStorageKeyModel struct {
	Account       string `json:"account"`
	CipherPrivate string `json:"cipherprivate"`
	Private       string `json:"-"`
}

func (d KeyStorageKeyModel) GetAccount() string {
	return d.Account
}
func (d KeyStorageKeyModel) GetCipherPrivate() string {
	return d.CipherPrivate
}

func (d KeyStorageKeyModel) GetPrivate() string {
	return d.Private
}

type KeyStorageModel struct {
	Key   *[]KeyStorageKeyModel `json:"keys"`
	Hash  string                `json:"hash"`
	Token string                `json:"-"`
}

func (d KeyStorageModel) GetKey() *[]KeyStorageKeyModel {
	return d.Key
}

func (d KeyStorageModel) GetHash() string {
	return d.Hash
}

func (d KeyStorageModel) GetToken() string {
	return d.Token
}
