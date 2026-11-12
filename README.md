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

A complete Dockerised development environment is available, which is useful for both
contributing to the development of VOR and when developing your own VOR-enabled smart
contracts.

The development environment can be run using:

```bash
make dev-env
```

The image (`docker/Dockerfile`) holds Go, Foundry's `anvil` and `cast`, and the oracle built from this tree.
The environment (`docker/scripts/dev-env.sh`) will:

1. Start an `anvil` chain mining a block a second, with 20 accounts from the mnemonic below
2. Deploy the VOR contracts from the compiled bytecode in `abigenBindings/bin`, at the addresses below
3. Send test xFUND to accounts 2-19
4. Run the `oracle`, which registers its proving key

The container exposes port `8545` allowing the chain to be accessible
via http://127.0.0.1:8545

Additionally, port `8445` is exposed, so the oracle's admin commands (`oracle about`, `oracle withdraw`, ...)
can reach it.

#### Dev environment configuration

- Wallet mnemonic: `myth like bonus scare over problem client lizard pioneer submit female collect` (ganache's
  deterministic one, so the accounts and contract addresses are the ones ganache gave)
- Chain URL: `http://127.0.0.1:8545`
- Network/Chain ID: `696969`
- VORCoordinator Contract address: `0xCfEB869F69431e42cdB54A4F4f105C19C080A601`
- BlockHashStore contract address: `0x5b1869D9A4C187F2EAa108f3062412ecf0526b24`
- VOR Oracle Wallet Address: `0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0`
- VOR Oracle KeyHash: `0x1a7a24165e904cb38eb8344affcf8fdee72ac11b5c542428b35eef5769c409f0`
- VOR Oracle API key: `0pear3uoznba36fwzoaspwrvc164bkjd`

### Admin commands in the Dev environment

Build the oracle (`make build-oracle`); with the dev environment running, its admin commands reach it using the
dev config and API key, for example:

```bash
./oracle/build/oracle about -c ./docker/assets/config.json --api-key ./docker/assets/pass
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

This runs, in the same image, the `go` unit and integration tests against a local `anvil` chain with the test
contracts, then the end-to-end flow (`docker/scripts/e2e.sh`): on two local chains, the oracle registers its
proving key, fulfils dice rolls on `VORD20` with EIP-1559 and legacy transactions, and has a second chain added
to it.

### Polygon Mumbai

TestXFUND:      0xb07C72acF3D7A5E9dA28C56af6F93862f8cc8196  
BlockHashStore: 0xf6b5d6eafE402d22609e685DE3394c8b359CaD31  
VORCoordinator: 0xc12678b997ce94e9f3921B65AD144565dC20Aefc   

Provider: `0x06b5959C3D8212a5718c4b5d6827aa7b2f29E2D5`  
**Public Key**: `0x0499d4b80d710ad8ca1ae12ac0682a3e3e004637cd10ccf47a6badc1f45b56cfcbc7504ed77f062e2b9922c13daeb001004a19ca632a8354f6c65d4b00991c85c8`  
**Key Hash**: `0x91c07383e58d6a27a1e8d89806b5959c3d8212a5718c4b5d6827aa7b2f29e2d5`  
**Global Fee**: 0.000001 xFUNDMOCK  
**Wait Time**: 10 block confirmations
