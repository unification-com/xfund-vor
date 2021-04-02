package keystorage

type KeyStorageKeyModel struct {
	Account string `json:"account"`
	// encrypted private key
	CipherPrivate string `json:"cipherprivate"`
	Private       string `json:"-"`
	// true if public key generated from this private is already
	// registered in VORCoordinator
	Registered bool `json:"registered"`
	// last checked block number (may be set manually)
	BlockNumber int64 `json:"block_number"`
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

func (d KeyStorageKeyModel) GetRegistered() bool {
	return d.Registered
}

func (d KeyStorageKeyModel) GetBlockNumber() int64 {
	return d.BlockNumber
}

func (d KeyStorageKeyModel) SetAccount(account string) {
	d.Account = account
}

func (d KeyStorageKeyModel) SetCipherPrivate(cipherPrivate string) {
	d.CipherPrivate = cipherPrivate
}

func (d KeyStorageKeyModel) SetPrivate(private string) {
	d.Private = private
}

type KeyStorageModel struct {
	Key  []*KeyStorageKeyModel `json:"keys"`
	Hash string                `json:"hash"`
	// used to store decrypted api token(key) which is being used
	Token string `json:"-"`
	// used to store decrypted private key which is being used
	PrivateKey string `json:"-"`
}

func (d KeyStorageModel) GetKey() []*KeyStorageKeyModel {
	return d.Key
}

func (d KeyStorageModel) GetHash() string {
	return d.Hash
}

func (d KeyStorageModel) GetToken() string {
	return d.Token
}

func (d *KeyStorageModel) GetPrivateKey() string {
	return d.PrivateKey
}

func (d *KeyStorageModel) SetKey(key []*KeyStorageKeyModel) {
	d.Key = key
}

func (d *KeyStorageModel) SetHash(hash string) {
	d.Hash = hash
}

func (d *KeyStorageModel) SetToken(token string) {
	d.Token = token
}
