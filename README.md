[![GitHub release (latest by date)](https://img.shields.io/github/v/release/unification-com/xfund-vor?label=oracle%20version)](https://github.com/unification-com/xfund-vor/releases/latest)
[![npm](https://img.shields.io/npm/v/@unification-com/xfund-vor?label=smart%20contract%20version%20%28npm%29)](https://www.npmjs.com/package/@unification-com/xfund-vor)
![sc unit tests](https://github.com/unification-com/xfund-vor/actions/workflows/test-contracts.yml/badge.svg)

# Verified Open Randomness

A suite of Ethereum smart contracts, and accompanying Provider Oracle software
for running VOR, and integrating VOR into your own smart contracts.

## VOR Integration Quickstart

1. Install the package

```bash
yarn add @unification-com/xfund-vor
```

or

```bash
npm i @unification-com/xfund-vor
```

2. Import `VORConsumerBase.sol` into your contract and pass
   parameters to your `constructor`

```solidity
import "@unification-com/xfund-vor/contracts/VORConsumerBase.sol";

contract MyRandomNumberContract is VORConsumerBase {
    constructor(address _vorCoordinator, address _xfund)
    public VORConsumerBase(_vorCoordinator, _xfund) {
        // other stuff...
    }
}
```

3. Implement a `requestRandomness` function

```solidity
    function requestRandomness(uint256 _userProvidedSeed, bytes32 _keyHash, unit256 _fee) 
    external
    returns (bytes32 requestId) {
        requestId = requestRandomness(_keyHash, _fee, _userProvidedSeed);
        // other stuff...
    }
```

4. Implement the `fulfillRandomness` function for data Providers to send data

```solidity
    function fulfillRandomness(bytes32 requestId, uint256 randomness) internal override {
        // do something with the received number
        uint256 randVal = randomness.mod(999).add(1);
        // other stuff...
    }
```

## Development and Testing

Clone of fork

```bash
git clone https://github.com/unification-com/xfund-vor
npx truffle compile
```

### Dev Environment

A complete Dockerised development environment is available, and can be run using:

```bash
make dev-env
```

Alternatively, run the docker commands as follows:

```bash
docker build -t vor_dev_env -f docker/dev.Dockerfile .
docker run -it -p 8545:8545 -p 8445:8445 vor_dev_env
```

The environment will:

1. Spawn a deterministic `ganach-cli` development chain
2. Compile and deploy the necessary VOR smart contracts
3. Initialise the test accounts, send test tokens and register the Oracle's proving key
4. Run the `oracle` application

The container exposes port `8545` allowing the Ganache chain to be accessible
via http://127.0.0.1:8545

Additionally, port `8445` is exposed, allowing the oracle to be accessed using the `oracle-cli`
tool.

#### Dev environment configuration

- VORCoordinator Contract address: `0xCfEB869F69431e42cdB54A4F4f105C19C080A601`
- BlockHashStore contract address: `0x5b1869D9A4C187F2EAa108f3062412ecf0526b24`
- VOR Oracle Wallet Address: `0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0`
- VOR Oracle KeyHash: `0x1a7a24165e904cb38eb8344affcf8fdee72ac11b5c542428b35eef5769c409f0`
- VOR Oracle API key: `0pear3uoznba36fwzoaspwrvc164bkjd`

### `oraclecli` in the Dev environment

You will need to build the tool:

```bash
make build-oracle-cli
```

and save the following to `$HOME/.oracle-cli_settings_dev_env.json`:

```json
{
  "oracle_host": "127.0.0.1",
  "oracle_port": "8445",
  "oracle_key": "0pear3uoznba36fwzoaspwrvc164bkjd"
}
```

The `oracle-cli` commands should now be available when the dev environment is running, for example:

```bash
./oracle-cli/build/oraclecli about -c $HOME/.oracle-cli_settings_dev_env.json
```

### Unit Testing

Run smart contract tests:

```bash
yarn test
```

Run `oracle` tests:

```bash
make test-oracle
```

This will run the `go` tests in a self-contained, dockerised environment. A Ganache
network will be run within the container for the tests to run against.
