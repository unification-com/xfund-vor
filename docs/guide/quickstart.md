# Quickstart

An example implementation can be found in https://github.com/unification-com/xfund-vor/blob/master/contracts/examples/VORD20.sol

## Integration

In order to request randomness,you will need to import the `VORCoordinator.sol` smart contract
and set up some simple functions within your smart contract. It is advisable to also implement
OpenZeppelin's `Ownable` functionality to secure some of the required functions:

1. Add the packages to your project:

```
yarn add @unification-com/xfund-vor @openzeppelin/contracts
```

2. In your smart contract, import `VORCoordinator.sol` and `Ownable.sol`:

```solidity
import "@openzeppelin/contracts/access/Ownable.sol";
import "@unification-com/xfund-vor/contracts/VORCoordinator.sol";
```

3. Extend your contract, adding `is VORConsumerBase, Ownable`:

```solidity
contract MockConsumer is VORConsumerBase, Ownable {
```

4. Ensure your `constructor` function has at least two parameters to accept the `VORCoordinator` 
   and `xFUND` smart contract addresses, and pass them to the `ConsumerBase`:

```solidity
constructor(address _vorCoordinator, address _xfund)
    public VORConsumerBase(_vorCoordinator, _xfund) {
        // other stuff...
    }
```

5. Implement a `topUpGas` function to enable gas refunds to providers who require refunds for data fulfilment:

```solidity
function topUpGas(bytes32 _sKeyHash, uint256 _amount) public onlyOwner {
    topUpGas(_sKeyHash, _amount);
}
```

6. Optionally implement a `withdrawXFUND` function, to allow you to remove any xFUND held by your contract:

```solidity
function withdrawXFUND(address to, uint256 value) public onlyOwner {
    require(xFUND.transfer(to, value), "Not enough xFUND");
}
```

7. Implement a `requestRandomness` function, for example:

```solidity
function requestRandomness(uint256 _userProvidedSeed, bytes32 _keyHash, unit256 _fee) 
public 
onlyOwner 
returns (bytes32 requestId) {
    require(xFUND.balanceOf(address(this)) >= _fee, "Not enough xFUND to pay fee");
    requestId = requestRandomness(_keyHash, _fee, _userProvidedSeed);
    emit RandomnessRequested(requestId);
}
```

8. Implement the `fulfillRandomness` function for data Providers to send data, for example:

```solidity
function fulfillRandomness(bytes32 requestId, uint256 randomness) internal override {
    uint256 randVal = randomness.mod(999).add(1);
    emit RandomnessReceived(requestId, randVal);
}
```

You should now be ready to compile and deploy your smart contract with your preferred toolchain
(Truffle, Hardhat etc.)

## Initialisation

Once integrated, compiled and deployed, you will need to send some transactions to the
Ethereum blockchain in order to initialise the fee payment and data acquisition environment
This involves:

1) Increasing the `xFUND` token allowance on the `VORCoordinator` smart contract, in order for the `VORCoordinator`
   to accept and pay xFUND fees to VOR providers. This need only be run once, if the initial
   allowance is set high enough.
2) Transfer some `xFUND` tokens to your smart contract, that is integrating the Consumer Library.
   This allows you to submit randomness requests, and your contract to pay fees. The required amount
   of `xFUND` to pay for a request fee is sent to the `VORCoordinator` with each request.

   **Note**: The `xFUNDMOCK` Token on Rinkeby testnet has a faucet function, `gimme()` which can be used
   to grab some test tokens.
3) "Topping up" gas payments on the `VORCoordinator` - a small amount of ETH will be held by the `VORCoordinator`
   on your behalf in order to reimburse VOR providers for the cost of sending a Tx to your contract
   and submitting the data to you. This can be periodically topped up in small amounts, and can
   also be withdrawn by you in its entirety at any time.

Once these steps have been run through, you will be able to initialise data requests via your
smart contract.

## Requesting Randomness

Once the environment has been initialised, you will be able to request randomness


