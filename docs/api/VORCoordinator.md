# VORCoordinator

Coordinates on-chain verifiable-randomness requests

## Functions:
- [`constructor(address _xfund, address _blockHashStore) public`](#VORCoordinator-constructor-address-address-)
- [`getProviderAddress(bytes32 _keyHash) external`](#VORCoordinator-getProviderAddress-bytes32-)
- [`getProviderFee(bytes32 _keyHash) external`](#VORCoordinator-getProviderFee-bytes32-)
- [`getProviderGranularFee(bytes32 _keyHash, address _consumer) external`](#VORCoordinator-getProviderGranularFee-bytes32-address-)
- [`registerProvingKey(uint256 _fee, address payable _oracle, uint256[2] _publicProvingKey) external`](#VORCoordinator-registerProvingKey-uint256-address-payable-uint256-2--)
- [`changeFee(uint256[2] _publicProvingKey, uint256 _fee) external`](#VORCoordinator-changeFee-uint256-2--uint256-)
- [`changeGranularFee(uint256[2] _publicProvingKey, uint256 _fee, address _consumer) external`](#VORCoordinator-changeGranularFee-uint256-2--uint256-address-)
- [`withdraw(address _recipient, uint256 _amount) external`](#VORCoordinator-withdraw-address-uint256-)
- [`randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) external`](#VORCoordinator-randomnessRequest-bytes32-uint256-uint256-)
- [`hashOfKey(uint256[2] _publicKey) public`](#VORCoordinator-hashOfKey-uint256-2--)
- [`fulfillRandomnessRequest(bytes _proof) public`](#VORCoordinator-fulfillRandomnessRequest-bytes-)
- [`callBackWithRandomness(bytes32 requestId, uint256 randomness, address consumerContract) internal`](#VORCoordinator-callBackWithRandomness-bytes32-uint256-address-)
- [`getRandomnessFromProof(bytes _proof) internal`](#VORCoordinator-getRandomnessFromProof-bytes-)

## Events:
- [`RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)`](#VORCoordinator-RandomnessRequest-bytes32-uint256-address-uint256-bytes32-)
- [`NewServiceAgreement(bytes32 keyHash, uint256 fee)`](#VORCoordinator-NewServiceAgreement-bytes32-uint256-)
- [`ChangeFee(bytes32 keyHash, uint256 fee)`](#VORCoordinator-ChangeFee-bytes32-uint256-)
- [`ChangeGranularFee(bytes32 keyHash, address consumer, uint256 fee)`](#VORCoordinator-ChangeGranularFee-bytes32-address-uint256-)
- [`RandomnessRequestFulfilled(bytes32 requestId, uint256 output)`](#VORCoordinator-RandomnessRequestFulfilled-bytes32-uint256-)

## Modifiers:
- [`sufficientXFUND(uint256 _feePaid, bytes32 _keyHash)`](#VORCoordinator-sufficientXFUND-uint256-bytes32-)
- [`hasAvailableFunds(uint256 _amount)`](#VORCoordinator-hasAvailableFunds-uint256-)

<a name="VORCoordinator-constructor-address-address-"></a>
### Function `constructor(address _xfund, address _blockHashStore) public `
No description
<a name="VORCoordinator-getProviderAddress-bytes32-"></a>
### Function `getProviderAddress(bytes32 _keyHash) external  -> address`
getProviderAddress - get provider address

<a name="VORCoordinator-getProviderFee-bytes32-"></a>
### Function `getProviderFee(bytes32 _keyHash) external  -> uint96`
getProviderFee - get provider's base fee

<a name="VORCoordinator-getProviderGranularFee-bytes32-address-"></a>
### Function `getProviderGranularFee(bytes32 _keyHash, address _consumer) external  -> uint96`
getProviderGranularFee - get provider's granular fee for selected consumer

<a name="VORCoordinator-registerProvingKey-uint256-address-payable-uint256-2--"></a>
### Function `registerProvingKey(uint256 _fee, address payable _oracle, uint256[2] _publicProvingKey) external `
No description
#### Parameters:
- `_fee`: minimum xFUND payment required to serve randomness

- `_oracle`: the address of the node with the proving key

- `_publicProvingKey`: public key used to prove randomness
<a name="VORCoordinator-changeFee-uint256-2--uint256-"></a>
### Function `changeFee(uint256[2] _publicProvingKey, uint256 _fee) external `
No description
#### Parameters:
- `_publicProvingKey`: public key used to prove randomness

- `_fee`: minimum xFUND payment required to serve randomness
<a name="VORCoordinator-changeGranularFee-uint256-2--uint256-address-"></a>
### Function `changeGranularFee(uint256[2] _publicProvingKey, uint256 _fee, address _consumer) external `
No description
#### Parameters:
- `_publicProvingKey`: public key used to prove randomness

- `_fee`: minimum xFUND payment required to serve randomness
<a name="VORCoordinator-withdraw-address-uint256-"></a>
### Function `withdraw(address _recipient, uint256 _amount) external `
Allows the oracle operator to withdraw their xFUND

#### Parameters:
- `_recipient`: is the address the funds will be sent to

- `_amount`: is the amount of xFUND transferred from the Coordinator contract
<a name="VORCoordinator-randomnessRequest-bytes32-uint256-uint256-"></a>
### Function `randomnessRequest(bytes32 _keyHash, uint256 _consumerSeed, uint256 _feePaid) external `
_consumerSeed is mixed with key hash, sender address and nonce to
obtain preSeed, which is passed to VOR oracle, which mixes it with the
hash of the block containing this request, to compute the final seed.

The requestId used to store the request data is constructed from the
preSeed and keyHash.
#### Parameters:
- `_keyHash`: ID of the VOR public key against which to generate output

- `_consumerSeed`: Input to the VOR, from which randomness is generated

- `_feePaid`: Amount of xFUND sent with request. Must exceed fee for key


<a name="VORCoordinator-hashOfKey-uint256-2--"></a>
### Function `hashOfKey(uint256[2] _publicKey) public  -> bytes32`
No description
#### Parameters:
- `_publicKey`: the key to return the address for
<a name="VORCoordinator-fulfillRandomnessRequest-bytes-"></a>
### Function `fulfillRandomnessRequest(bytes _proof) public `
No description
#### Parameters:
- `_proof`: the proof of randomness. Actual random output built from this
<a name="VORCoordinator-callBackWithRandomness-bytes32-uint256-address-"></a>
### Function `callBackWithRandomness(bytes32 requestId, uint256 randomness, address consumerContract) internal `
No description
<a name="VORCoordinator-getRandomnessFromProof-bytes-"></a>
### Function `getRandomnessFromProof(bytes _proof) internal  -> bytes32 currentKeyHash, struct VORCoordinator.Callback callback, bytes32 requestId, uint256 randomness`
No description

<a name="VORCoordinator-RandomnessRequest-bytes32-uint256-address-uint256-bytes32-"></a>
### Event `RandomnessRequest(bytes32 keyHash, uint256 seed, address sender, uint256 fee, bytes32 requestID)`
No description
<a name="VORCoordinator-NewServiceAgreement-bytes32-uint256-"></a>
### Event `NewServiceAgreement(bytes32 keyHash, uint256 fee)`
No description
<a name="VORCoordinator-ChangeFee-bytes32-uint256-"></a>
### Event `ChangeFee(bytes32 keyHash, uint256 fee)`
No description
<a name="VORCoordinator-ChangeGranularFee-bytes32-address-uint256-"></a>
### Event `ChangeGranularFee(bytes32 keyHash, address consumer, uint256 fee)`
No description
<a name="VORCoordinator-RandomnessRequestFulfilled-bytes32-uint256-"></a>
### Event `RandomnessRequestFulfilled(bytes32 requestId, uint256 output)`
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
