# Interacting with the DnD example

If you haven't already, run through the [full implementation example](./implementation.md)

This guide assumes you are familiar with Solidity development frameworks
such as Truffle/Hardhat.

## Prerequisites

The contract will be deployed on Rinkeby testnet. You will need:

1. wallets with test Eth - for example from the [faucet](https://faucet.rinkeby.io/)
2. xFUNDMOCK tokens for each wallet interacting, including the owner. See
   [contracts](../contracts.md) for the Rinkeby testnet address, and call the
   `gimme()` function to get 10 xFUNDMOCK tokens.

## Deployment

Write your deployment script for your contract (this will be dependent on your framework
of choice). You will need to pass the two contract values to the `constructor`
for both `xFUNDMOCK` and the `VORCoordinator`. These can be obtained from
[contracts](../contracts.md).

Once you're ready, deploy your contract on Rinkeby. You can also optionally
verify the contract source code on Etherscan.

## Initialise

Once deployed, we need to call a couple of functions to set things up.

### DnD.increaseVorAllowance

The first thing we need to do is allow the `VORCoordinator` smart contract
to spend xFUNDMOCK on behalf of our DnD smart contract. This allows us to
request random numbers, and for the `VORCoordinator` to transfer fees to
the VOR Oracle Provider.

`increaseVorAllowance` can be called from your framework's development console,
or if you have verified the source code on Etherscan, from their "Write Contract"
feature.

The amount should be suitably high to cover any future costs. Many services
(for example Uniswap), do this once with the maximum allowable `uint256`
value "115792089237316195423570985008687907853269984665640564039457584007913129639935"

::: tip
`increaseVorAllowance` must be called by the wallet that deployed the DnD
example smart contract
:::

### xFUNDMOCK.increaseAllowance

_Each_ wallet that interacts with the DnD smart contract and intends to call
the `rollForHit` function will need to allow the DnD smart contract to spend
xFUNDMOCK tokens on their behalf.

The `increaseAllowance` function can be executed on Etherscan (see [contracts](../contracts.md)
for the Rinkeby address), using Metamask etc. Similarly, the amount
should be sufficiently high, and the `spender` value should be the address 
of the deployed DnD smart contract.

## Interact

Now we have the boring stuff out of the way, we can actually run the game.
As before, these functions can be called via your Solidity framework of choice
or via Etherscan.

### DnD.addMonster

This must be called by the wallet that deployed the contract. It's possible
to add up to 20 monsters, with each one requiring a short name (for example
"Orc"), and an AC value between 1 - 20.

### DnD.monsters

`monsters` has a built-in getter, which will accept the Id of a monster
and retrieve its stats. The result returned will be the name and AC.

### DnD.changeStrModifier

Each wallet calling the `rollForHit` can optionally change their STR modifier,
up to 5. This modifier gets added to the d20 roll when the randomness request
is fulfilled.

### DnD.rollForHit

A wallet can call the `rollForHit` to generate a random d20 roll to determin
if they hit the monster or not. There are 4 required parameters:

1. The `monsetrId` of the monster being hit. See `DnD.monsters` above
2. A `seed` value - this can be any number, and gets mixed in with the
   random number generation
3. The `keyhash` for the provider you wish to fulfil the request. This can be
   obtained from the [providers](../providers.md) page.
4. the `fee` amount required to pay for the request. This will be deducted
   from the wallet calling this function.
   
::: warn
A VOR Provider's base fee is available on the [providers](../providers.md) page
or by querying the `VORCoordinator` smart contract. This is the minimum
amount they will accept for fulfilling a request.

However, providers may also set granular fees on a per-consumer contract
basis, because not every contract is the same and more complex consumer
contracts will incur higher gas costs for the provider.

The current fee for your contract can be acquired by calling the
`VORCoordinator.getProviderGranularFee()` function and passing the
provider's `keyHash` along with the address of your contract.

The `VORCoordinator` address is available on the [contracts](../contracts.md)
page.
:::

Once a request has been initialised, the VOR Provider Oracle will see
a `RandomnessRequest` event emitted in the `VORCoordinator` smart 
contract, generate a random number, and call your contract's 
`fulfillRandomness` function to complete the process.

### DnD.getLastResult

It's possible to monitor the DnD contract's event logs and watch for the
`HitResult` event. This indicates that the VOR Provider Oracle has fulfilled
the randomness request.

Calling the `DnD.getLastResult` function requires 2 parameters:

1. `player` - the wallet address of the player that rolled
2. `monsterId` - the ID of the monster they attempted to hit

The result returned will be something along the lines of:

```
[
  roll: '17',
  modified: '19',
  result: 'hit',
  isRolling: false
]
```
