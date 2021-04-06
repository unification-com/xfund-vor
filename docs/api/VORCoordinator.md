# with off-chain responses


## Functions:
- [`constructor(address _xfund, address _blockHashStore, uint256 _expectedGasFirst, uint256 _expectedGas)`](#VORCoordinator-constructor-address-address-uint256-uint256-)
- [`getTotalGasDeposits()`](#VORCoordinator-getTotalGasDeposits--)
- [`getGasDepositsForConsumer(address _consumer)`](#VORCoordinator-getGasDepositsForConsumer-address-)
- [`getGasDepositsForConsumerProvider(address _consumer, bytes32 _keyHash)`](#VORCoordinator-getGasDepositsForConsumerProvider-address-bytes32-)
- [`getGasTopUpLimit()`](#VORCoordinator-getGasTopUpLimit--)
- [`getProviderAddress(bytes32 _keyHash)`](#VORCoordinator-getProviderAddress-bytes32-)
- [`setGasTopUpLimit(uint256 _gasTopUpLimit)`](#VORCoordinator-setGasTopUpLimit-uint256-)
- [`setBaseGasRates(uint256 _newExpectedGasFirst, uint256 _newExpectedGas)`](#VORCoordinator-setBaseGasRates-uint256-uint256-)
- [`registerProvingKey(uint256 _fee, address payable _oracle, uint256[2] _publicProvingKey, bool _providerPaysGas)`](#VORCoordinator-registerProvingKey-uint256-address-payable-uint256-2--bool-)
- [`changeFee(uint256[2] _publicProvingKey, uint256 _fee)`](#VORCoordinator-changeFee-uint256-2--uint256-)
- [`setProviderPaysGas(uint256[2] _publicProvingKey, bool _providerPays)`](#VORCoordinator-setProviderPaysGas-uint256-2--bool-)
- [`withdraw(address _recipient, uint256 _amount)`](#VORCoordinator-withdraw-address-uint256-)
- [`randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid)`](#VORCoordinator-randomnessRequest-bytes32-uint256-uint256-)
- [`topUpGas(bytes32 _keyHash)`](#VORCoordinator-topUpGas-bytes32-)
- [`withDrawGasTopUpForProvider(bytes32 _keyHash)`](#VORCoordinator-withDrawGasTopUpForProvider-bytes32-)
- [`hashOfKey(uint256[2] _publicKey)`](#VORCoordinator-hashOfKey-uint256-2--)
- [`fulfillRandomnessRequest(bytes _proof)`](#VORCoordinator-fulfillRandomnessRequest-bytes-)

## Events:
- [`RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)`](#VORCoordinator-RandomnessRequest-bytes32-uint256-address-uint256-bytes32-)
- [`NewServiceAgreement(bytes32 keyHash, uint256 fee)`](#VORCoordinator-NewServiceAgreement-bytes32-uint256-)
- [`ChangeFee(bytes32 keyHash, uint256 fee)`](#VORCoordinator-ChangeFee-bytes32-uint256-)
- [`RandomnessRequestFulfilled(bytes32 requestId, uint256 output)`](#VORCoordinator-RandomnessRequestFulfilled-bytes32-uint256-)
- [`GasToppedUp(address consumer, address provider, uint256 amount)`](#VORCoordinator-GasToppedUp-address-address-uint256-)
- [`GasWithdrawnByConsumer(address consumer, address provider, uint256 amount)`](#VORCoordinator-GasWithdrawnByConsumer-address-address-uint256-)
- [`SetGasTopUpLimit(address sender, uint256 oldLimit, uint256 newLimit)`](#VORCoordinator-SetGasTopUpLimit-address-uint256-uint256-)
- [`SetBaseGasRates(address sender, uint256 oldFirstExpected, uint256 newFirstExpected, uint256 oldExpected, uint256 newExpected)`](#VORCoordinator-SetBaseGasRates-address-uint256-uint256-uint256-uint256-)
- [`GasRefundedToProvider(address consumer, address provider, uint256 amount)`](#VORCoordinator-GasRefundedToProvider-address-address-uint256-)
- [`SetProviderPaysGas(address provider, bool providerPays)`](#VORCoordinator-SetProviderPaysGas-address-bool-)

## Modifiers:
- [`sufficientXFUND(uint256 _feePaid, bytes32 _keyHash)`](#VORCoordinator-sufficientXFUND-uint256-bytes32-)
- [`hasAvailableFunds(uint256 _amount)`](#VORCoordinator-hasAvailableFunds-uint256-)

<a name="VORCoordinator-constructor-address-address-uint256-uint256-"></a>
### Function `constructor(address _xfund, address _blockHashStore, uint256 _expectedGasFirst, uint256 _expectedGas)`
No description
<a name="VORCoordinator-getTotalGasDeposits--"></a>
### Function `getTotalGasDeposits() -> uint256`
getTotalGasDeposits - get total gas deposited in VORCoordinator

<a name="VORCoordinator-getGasDepositsForConsumer-address-"></a>
### Function `getGasDepositsForConsumer(address _consumer) -> uint256`
getGasDepositsForConsumer - get gas deposited in VORCoordinator by a consumer

<a name="VORCoordinator-getGasDepositsForConsumerProvider-address-bytes32-"></a>
### Function `getGasDepositsForConsumerProvider(address _consumer, bytes32 _keyHash) -> uint256`
getGasDepositsForConsumerProvider - get gas deposited in VORCoordinator by a consumer
for a given provider

<a name="VORCoordinator-getGasTopUpLimit--"></a>
### Function `getGasTopUpLimit() -> uint256`
getGasTopUpLimit - get gas top up limit

<a name="VORCoordinator-getProviderAddress-bytes32-"></a>
### Function `getProviderAddress(bytes32 _keyHash) -> address`
getProviderAddress - get provider address

<a name="VORCoordinator-setGasTopUpLimit-uint256-"></a>
### Function `setGasTopUpLimit(uint256 _gasTopUpLimit) -> bool success`
setGasTopUpLimit set the max amount of ETH that can be sent
in a topUpGas Tx. Router admin calls this to set the maximum amount
a Consumer can send in a single Tx, to prevent large amounts of ETH
being sent.


#### Parameters:
- `_gasTopUpLimit`: amount in wei

<a name="VORCoordinator-setBaseGasRates-uint256-uint256-"></a>
### Function `setBaseGasRates(uint256 _newExpectedGasFirst, uint256 _newExpectedGas) -> bool success`
setBaseGasRates set the base expected gas values used for calculating gas
refunds


#### Parameters:
- `_newExpectedGasFirst`: expected gas units consumed for first fulfilment

- `_newExpectedGas`: expected gas units consumed for subsequent fulfilments

<a name="VORCoordinator-registerProvingKey-uint256-address-payable-uint256-2--bool-"></a>
### Function `registerProvingKey(uint256 _fee, address payable _oracle, uint256[2] _publicProvingKey, bool _providerPaysGas)`
No description
#### Parameters:
- `_fee`: minimum xFUND payment required to serve randomness

- `_oracle`: the address of the node with the proving key

- `_publicProvingKey`: public key used to prove randomness

- `_providerPaysGas`: true if provider will pay gas
<a name="VORCoordinator-changeFee-uint256-2--uint256-"></a>
### Function `changeFee(uint256[2] _publicProvingKey, uint256 _fee)`
No description
#### Parameters:
- `_publicProvingKey`: public key used to prove randomness

- `_fee`: minimum xFUND payment required to serve randomness
<a name="VORCoordinator-setProviderPaysGas-uint256-2--bool-"></a>
### Function `setProviderPaysGas(uint256[2] _publicProvingKey, bool _providerPays) -> bool success`
setProviderPaysGas - provider calls for setting who pays gas
for sending the fulfillRequest Tx

#### Parameters:
- `_providerPays`: bool - true if provider will pay gas

<a name="VORCoordinator-withdraw-address-uint256-"></a>
### Function `withdraw(address _recipient, uint256 _amount)`
Allows the oracle operator to withdraw their xFUND

#### Parameters:
- `_recipient`: is the address the funds will be sent to

- `_amount`: is the amount of xFUND transferred from the Coordinator contract
<a name="VORCoordinator-randomnessRequest-bytes32-uint256-uint256-"></a>
### Function `randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid)`
_consumerSeed is mixed with key hash, sender address and nonce to
obtain preSeed, which is passed to VOR oracle, which mixes it with the
hash of the block containing this request, to compute the final seed.

The requestId used to store the request data is constructed from the
preSeed and keyHash.
#### Parameters:
- `_keyHash`: ID of the VOR public key against which to generate output

- `_consumerSeed`: Input to the VOR, from which randomness is generated

- `_feePaid`: Amount of xFUND sent with request. Must exceed fee for key


<a name="VORCoordinator-topUpGas-bytes32-"></a>
### Function `topUpGas(bytes32 _keyHash) -> bool success`
topUpGas consumer contract calls this function to top up gas
Gas is the ETH held by this contract which is used to refund Tx costs
to the VOR provider for fulfilling a request.

To prevent silly amounts of ETH being sent, a sensible limit is imposed.

Can only top up for authorised providers


#### Parameters:
- `_keyHash`: ID of the VOR public key against which to generate output

<a name="VORCoordinator-withDrawGasTopUpForProvider-bytes32-"></a>
### Function `withDrawGasTopUpForProvider(bytes32 _keyHash) -> uint256 amountWithdrawn`
withDrawGasTopUpForProvider data consumer contract calls this function to
withdraw any remaining ETH stored in the Router for gas refunds for a specified
data provider.

Consumer contract will then transfer through to the consumer contract's
owner.

NOTE - data provider authorisation is not checked, since a consumer needs to
be able to withdraw for a data provide that has been revoked.


#### Parameters:
- `_keyHash`: ID of the VOR public key against which to generate output

<a name="VORCoordinator-hashOfKey-uint256-2--"></a>
### Function `hashOfKey(uint256[2] _publicKey) -> bytes32`
No description
#### Parameters:
- `_publicKey`: the key to return the address for
<a name="VORCoordinator-fulfillRandomnessRequest-bytes-"></a>
### Function `fulfillRandomnessRequest(bytes _proof)`
No description
#### Parameters:
- `_proof`: the proof of randomness. Actual random output built from this

<a name="VORCoordinator-RandomnessRequest-bytes32-uint256-address-uint256-bytes32-"></a>
### Event `RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)`
No description
<a name="VORCoordinator-NewServiceAgreement-bytes32-uint256-"></a>
### Event `NewServiceAgreement(bytes32 keyHash, uint256 fee)`
No description
<a name="VORCoordinator-ChangeFee-bytes32-uint256-"></a>
### Event `ChangeFee(bytes32 keyHash, uint256 fee)`
No description
<a name="VORCoordinator-RandomnessRequestFulfilled-bytes32-uint256-"></a>
### Event `RandomnessRequestFulfilled(bytes32 requestId, uint256 output)`
No description
<a name="VORCoordinator-GasToppedUp-address-address-uint256-"></a>
### Event `GasToppedUp(address consumer, address provider, uint256 amount)`
No description
<a name="VORCoordinator-GasWithdrawnByConsumer-address-address-uint256-"></a>
### Event `GasWithdrawnByConsumer(address consumer, address provider, uint256 amount)`
No description
<a name="VORCoordinator-SetGasTopUpLimit-address-uint256-uint256-"></a>
### Event `SetGasTopUpLimit(address sender, uint256 oldLimit, uint256 newLimit)`
No description
<a name="VORCoordinator-SetBaseGasRates-address-uint256-uint256-uint256-uint256-"></a>
### Event `SetBaseGasRates(address sender, uint256 oldFirstExpected, uint256 newFirstExpected, uint256 oldExpected, uint256 newExpected)`
No description
<a name="VORCoordinator-GasRefundedToProvider-address-address-uint256-"></a>
### Event `GasRefundedToProvider(address consumer, address provider, uint256 amount)`
No description
<a name="VORCoordinator-SetProviderPaysGas-address-bool-"></a>
### Event `SetProviderPaysGas(address provider, bool providerPays)`
No description

<a name="VORCoordinator-sufficientXFUND-uint256-bytes32-"></a>
### Modifier `sufficientXFUND(uint256 _feePaid, bytes32 _keyHash)`
Reverts if amount is not at least what was agreed upon in the service agreement

#### Parameters:
- `_feePaid`: The payment for the request

- `_keyHash`: The key which the request is for
<a name="VORCoordinator-hasAvailableFunds-uint256-"></a>
### Modifier `hasAvailableFunds(uint256 _amount)`
Reverts if amount requested is greater than withdrawable balance

#### Parameters:
- `_amount`: The given amount to compare to `withdrawableTokens`
