package models

type IKeyStorageKeyModel interface {
	GetAccount() string
	GetCipherPrivate() string
	GetPrivate() string
	SetAccount(account string)
	SetCipherPrivate(cipherPrivate string)
	SetPrivate(private string)
}

type IKeyStorageModel interface {
	GetKey() *[]IKeyStorageKeyModel
	GetHash() string
	GetToken() string
	SetKey(key *[]IKeyStorageKeyModel)
	SetHash(hash string)
	SetToken(token string)
}
