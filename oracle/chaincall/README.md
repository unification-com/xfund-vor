### How to run chaincall unit tests?

1. Deploy smart contracts
2. Write [config.json](../config/README.md)
3. Write [keystore.json](../store/keystorage/README.md)
4. go to chainlisten dir
5. run go test command like this:
` go test --run [function_name] -v -args [config.json path]`