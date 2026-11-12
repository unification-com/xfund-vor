# Oracle

API backend that Oracle client contracts on Ethereum make requests to. The backend utulizes Solidity contract ABIs to generate types for interacting with Ethereum contracts.

## Store

Oracle key store: `keystorage.file` in the configuration.

### keystore.json format

Each account's private key is a standard go-ethereum v3 keystore (scrypt), so `.keys[n].keystore` can be
extracted and used with any Ethereum tool. Every key is encrypted under one key, which the oracle reads from
`-k` (a file, or the key itself). Only the account selected by `keystorage.account` (else the first) is ever
decrypted. `api_key_sha256` is the hash of the HTTP API key (see below).

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
  ],
  "api_key_sha256": "..."
}
```

### HTTP API key

The HTTP API (used by the admin commands: `oracle about`, `oracle withdraw`, ...) has its own random key,
separate from the keystore key; the keystore keeps only its hash. With the oracle stopped, issue one with:

```bash
oracle api-key -c config.json [--out FILE]
```

It is written to `api_key` beside the config (or `--out`), at 0600, replacing any older key; the admin commands
read it from there (or `--api-key FILE`). First run issues one. Until a key is set, the API refuses every
request.

### Migrating a v0.1.x keystore

v0.1.x keystores (`cipherprivate` keys and a `hash`) used weak encryption under a guessable key, and the oracle
refuses to start on one. Back up the directory, then convert the keystore under a new key of at least 16
characters:

```bash
openssl rand -hex 32 > password.new && chmod 600 password.new
oracle migrate-keystore -c config.json -k password --new-key password.new && mv password.new password
```

Without `--new-key` the new key is asked for, twice, on the terminal. Every key is converted and verified from
disk before the old file is overwritten and replaced. Afterwards, delete the backup once the oracle runs: the old keystore and its key together expose the private keys. The old
key also opened the HTTP API, which now has its own key: issue one (above). Both commands refuse to run while
the oracle is running, as it would overwrite their changes.

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
  "max_gas_price": 150,
  "event_scan_batch_blocks": 2000,
  "keystorage": {
    "file": "./keystore.json",
    "account": "username"
  }
}
```

A config may list several chains under `"chains"`, each with the chain settings shown above (`network_id`,
`eth_http_host`, `contract_address`, `blockhash_store_address`, `first_block`, `check_duration`, `max_gas_price`,
`legacy_gas`, `wait_confirmations`, `event_scan_batch_blocks`, and `fee`, which registers the proving key on a
chain where it isn't registered yet) and an optional `name` (default: the network id);
a config for one chain may instead keep them at the top level, as above. A setting a chain leaves out takes its
default (`first_block` 1, `check_duration` 15 s, `max_gas_price` 150 gwei, `wait_confirmations` 10,
`event_scan_batch_blocks` 2,000); one it sets to 0 stays 0. `serve` defaults to `127.0.0.1:8445`.

Transactions are EIP-1559 where the chain's blocks have a base fee (tip as the node suggests, fee cap twice the
base fee plus the tip), else legacy-priced; `"legacy_gas": true` forces legacy pricing. `max_gas_price` caps the
gas price or fee cap in gwei (0: uncapped), so it must stay above the chain's base fee. Gas limits are always
estimated. Events are read
with `eth_getLogs` in ranges of at most `event_scan_batch_blocks` blocks (default 2,000), as range-limited
RPCs reject wider ones.

## Run tests

Before running tests, it is necessary to deploy contracts

```sh
go test --run [function_name] -v -args [config.json path]
```