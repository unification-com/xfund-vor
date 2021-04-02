# Oracle

API backend that Oracle client contracts on Ethereum make requests to. The backend utulizes Solidity contract ABIs to generate types for interacting with Ethereum contracts.

## Store

Oracle key store

### Expected keystore.json format

```json
{
  "keys": [
    {
      "account": "bestuser",
      "cipherprivate": "0xf54ca099a480e75a417a676855aed602f559d27f6f461f3754667b0b8af11ba6"
    },
    {
      "account": "username",
      "cipherprivate": "..."
    }
  ]
}

```

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