# 
This is only an example implementation and not necessarily suitable for mainnet.

## Functions:
- [`constructor(address vorCoordinator, address xfund, bytes32 keyHash, uint256 fee)`](#VORD20-constructor-address-address-bytes32-uint256-)
- [`rollDice(uint256 userProvidedSeed, address roller)`](#VORD20-rollDice-uint256-address-)
- [`increaseVorAllowance(uint256 _amount)`](#VORD20-increaseVorAllowance-uint256-)
- [`withdrawToken(address to, uint256 value)`](#VORD20-withdrawToken-address-uint256-)
- [`house(address player)`](#VORD20-house-address-)
- [`setKeyHash(bytes32 keyHash)`](#VORD20-setKeyHash-bytes32-)
- [`keyHash()`](#VORD20-keyHash--)
- [`setFee(uint256 fee)`](#VORD20-setFee-uint256-)
- [`fee()`](#VORD20-fee--)

## Events:
- [`DiceRolled(bytes32 requestId, address roller)`](#VORD20-DiceRolled-bytes32-address-)
- [`DiceLanded(bytes32 requestId, uint256 result)`](#VORD20-DiceLanded-bytes32-uint256-)


<a name="VORD20-constructor-address-address-bytes32-uint256-"></a>
### Function `constructor(address vorCoordinator, address xfund, bytes32 keyHash, uint256 fee)`
No description
#### Parameters:
- `vorCoordinator`: address of the VOR Coordinator

- `xfund`: address of the xFUND token

- `keyHash`: bytes32 representing the hash of the VOR provider

- `fee`: uint256 fee to pay the VOR oracle
<a name="VORD20-rollDice-uint256-address-"></a>
### Function `rollDice(uint256 userProvidedSeed, address roller) -> bytes32 requestId`
Warning: if the VOR response is delayed, avoid calling requestRandomness repeatedly
as that would give miners/VOR operators latitude about which VOR response arrives first.
You must review your implementation details with extreme care.


#### Parameters:
- `userProvidedSeed`: uint256 unpredictable seed

- `roller`: address of the roller
<a name="VORD20-increaseVorAllowance-uint256-"></a>
### Function `increaseVorAllowance(uint256 _amount)`
Wrapped around an Ownable modifier to ensure only the contract owner can call it.
Allows contract owner to increase the xFUND allowance for the VORCoordinator contract
enabling it to pay request fees on behalf of this contract's owner.
NOTE: This contract must have an xFUND balance in order to request randomness


#### Parameters:
- `_amount`: uint256 amount to increase allowance by
<a name="VORD20-withdrawToken-address-uint256-"></a>
### Function `withdrawToken(address to, uint256 value)`
Wrapped around an Ownable modifier to ensure only the contract owner can call it.
Allows contract owner to withdraw any xFUND currently held by this contract
<a name="VORD20-house-address-"></a>
### Function `house(address player) -> string`
No description
#### Parameters:
- `player`: address

#### Return Values:
- house as a string
<a name="VORD20-setKeyHash-bytes32-"></a>
### Function `setKeyHash(bytes32 keyHash)`
No description
#### Parameters:
- `keyHash`: bytes32
<a name="VORD20-keyHash--"></a>
### Function `keyHash() -> bytes32`
No description
<a name="VORD20-setFee-uint256-"></a>
### Function `setFee(uint256 fee)`
No description
#### Parameters:
- `fee`: uint256
<a name="VORD20-fee--"></a>
### Function `fee() -> uint256`
No description

<a name="VORD20-DiceRolled-bytes32-address-"></a>
### Event `DiceRolled(bytes32 requestId, address roller)`
No description
<a name="VORD20-DiceLanded-bytes32-uint256-"></a>
### Event `DiceLanded(bytes32 requestId, uint256 result)`
No description

