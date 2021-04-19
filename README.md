[![npm version](http://img.shields.io/npm/v/@unification-com/xfund-vor.svg?style=flat)](https://npmjs.org/package/@unification-com/xfund-vor "View this project on npm")
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
