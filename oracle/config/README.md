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