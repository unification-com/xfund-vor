# Oracle

API backend that Oracle client contracts on Ethereum make requests to. The backend utulizes Solidity contract ABIs to generate types for interacting with Ethereum contracts.

## Store

Oracle key store: `keystorage.file` in the configuration.

### keystore.json format

Each account's private key is a standard go-ethereum v3 keystore (scrypt), so `.keys[n].keystore` can be
extracted and used with any Ethereum tool. Every key is encrypted under one key, which the oracle reads from
`-k` (a file, or the key itself) and which also authenticates the HTTP API. Only the account selected by
`keystorage.account` (else the first) is ever decrypted.

```json
{
  "version": 2,
  "keys": [
    {
      "account": "oracle",
      "address": "0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1",
      "keystore": { "address": "90f8bf6a...", "crypto": { "kdf": "scrypt", "...": "..." }, "id": "...", "version": 3 },
      "registered": true,
      "block_number": 13452142
    }
  ]
}
```

### Migrating a v0.1.x keystore

v0.1.x keystores (`cipherprivate` keys and a `hash`) used weak encryption under a guessable key, and the oracle
refuses to start on one. Back up the directory, then convert the keystore under a new key of at least 16
characters:

```bash
openssl rand -hex 32 > password.new && chmod 600 password.new
oracle migrate-keystore -c config.json -k password --new-key password.new && mv password.new password
```

Without `--new-key` the new key is asked for, twice, on the terminal. Every key is converted and verified from
disk before the old file is overwritten and replaced. Afterwards, update the key in `oracle-cli settings`, and
delete the backup once the oracle runs: the old keystore and its key together expose the private keys.

## Configuration

Oracle settings

### Expected config.json format

```json
{
  "contract_address": "0x22F043993312CB050E7F7A5C1207f68a05D3ef66",
  "eth_http_host": "http://localhost:7545",
  "eth_ws_host": "http://localhost:7545",
  "network_id": 5777,
  "serve": {
    "host": "0.0.0.0",
    "port": 8888
  },
  "log_file": "./log.out",
  "log_level": "DEBUG",
  "keystorage": {
    "file": "./keystore.json",
    "account": "username"
  }
}
```

## Run tests

Before running tests, it is necessary to deploy contracts

```sh
go test --run [function_name] -v -args [config.json path]
```