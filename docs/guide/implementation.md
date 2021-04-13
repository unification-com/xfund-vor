# Full implementation example

In this guide, we'll run through a more detailed example and show you how to
implement a simplified DnD D20 for fighting monsters and calculating hits.

The resulting contract will:

- allow any player to "roll for hit" against a selected monster, using a VOR
  request for the D20 roll
- allow the owner to add up to 20 monsters
- allow any player to register their STR modifier
- allow the contract owner to implement and run some optional helper functions
  included in the `VORConsumerBase` contract.
  
Players calling the `rollToHit` function will pay their own xFUNDMOCK fees.

::: tip
This guide assumes you have read through the [quickstart](./quickstart.md)
and :

1. set up your development and NodeJS environment
2. initialised a new project (using npm, and Truffle/Hardhat etc.)
3. installed the required `@unification-com/xfund-vor` and
   supplementary `@openzeppelin/contracts` packages using 
   `npm`/`yarn`
:::

The full, final contract can be seen in our `vor-demos` Github repo: 
[https://github.com/unification-com/vor-demos/blob/main/contracts/DnD.sol](https://github.com/unification-com/vor-demos/blob/main/contracts/DnD.sol).

## Initial contract skeleton

We'll start with a simple contract outline - some function stubs and the constructor:

Create `contracts/DnD.sol`. All future edits will be in this file

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.6.12;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@unification-com/xfund-vor/contracts/VORConsumerBase.sol";

contract DnD is Ownable, VORConsumerBase {
    using SafeMath for uint256;

    constructor(address _vorCoordinator, address _xfund)
    public
    VORConsumerBase(_vorCoordinator, _xfund) {}

    function rollForHit(
        uint256 _monsterId, 
        uint256 _seed, 
        bytes32 _keyHash, 
        uint256 _fee
    ) 
    external returns (bytes32 requestId) {
        requestId = requestRandomness(_keyHash, _fee, _seed);
    }

    function fulfillRandomness(bytes32 _requestId, uint256 _randomness) 
    internal override {}
}
```

We're importing the `VORConsumerBase`, and some helper contracts from the 
awesome OpenZeppelin suite, and we have the three most basic functions 
required to implement VOR:

1. a `constructor` which passes the addresses of the `VORCoordinatior` and
   `xFUND` to the `VORConsumerBase`.
2. a `rollForHit` function which will ultimately call the `requestRandomness`
   function in `VORConsumerBase`.
3. a `fulfillRandomness` which overrides the same function `VORConsumerBase`, and
   which will be called by the VOR Oracle via the `VORCoordinator` smart contract.
   
::: tip
the `_keyHash` and `_fee` args passed to `rollForHit` function can be
set/retrieved elsewhere in the contract, and passed to the `requestRandomness`
function. Your request function does not necessarily require these arguments,
but the values do ultimately need passing to `requestRandomness`

`_seed` however, should be generated for each individual request.
:::

## Adding some functionality

The above contract isn't very useful. While it requests a random number, it
doesn't actually do anything with it. Let's add a few functions 
that can make use of the random number, and some support variables, structures
and events.

::: tip
_None of the following variables or events are required by VOR_. 
They are required only for the DnD contract's own functionality 
and for the purpose of demoing VOR.
:::

### Contract variables

We'll need a couple of simple contract variables to keep track of our monsters:

```solidity
    // keep track of the monsters added
    uint256 public currentMonsterId;

    // super simple monster stats
    struct Monster {
        string name;
        uint256 ac;
    }

    struct Result {
        uint256 roll;
        uint256 modified;
        string result;
        bool isRolling;
    }
```

Nothing complex here - each monster added will be assgined an incremental
ID `currentMonsterId`, and have some simple stats associated with it.
`Result` will be used to store that last result for a player/monster.

### Mappings

```solidity
    // monsters held in the contract
    mapping (uint256 => Monster) public monsters;
    // player STR modifiers
    mapping (address => uint256) public strModifiers;
    // map request IDs to monster IDs
    mapping(bytes32 => uint256) public requestIdToMonsterId;
    // map request IDs to player addresses, to retrieve STR modifiers
    mapping(bytes32 => address) public requestIdToAddress;
    // store last [player][monster] results
    mapping(address => mapping(uint256 => Result)) lastResult;
```

- `monsters` will be used to store our contracts monsters and their stats
- `strModifiers` will allow players to set their STR modifier, which gets
  added to each dice roll
- `requestIdToMonsterId` will temporarily store data on which request relates
  to which monster while the roll is in progress
- `requestIdToAddress` will similarly track which user rolled.
- `lastResult` will store the last roll and result for a player/monster

### Events

```solidity
    event AddMonster(uint256 monsterId, string name, uint256 ac);
    event ChangeStrModifier(address player, uint256 strMod);
    event HittingMonster(uint256 monsterId, bytes32 requestId);
    event HitResult(uint256 monsterId, bytes32 requestId, address player, string result, uint256 roll, uint256 modified);
```

Again, nothing complex here - just some events that will be emitted during
the functions we'll implement next.

## Monster & Player related functions

We'll add few non-VOR functions first, to support the contract 
owner adding monsters, and players editing their stats. 

::: tip
_Neither of the two following functions are required by VOR_. 
They are required only for the DnD contract's own functionality 
and for the purpose of demoing VOR.
:::

### addMonster

This will just allow the contract owner to add up to 20 monsters. The
`onlyOwner` modifier is part of OpenZeppelin's `Ownable` contract.

```solidity
    function addMonster(string memory _name, uint256 _ac) external onlyOwner {
        require(nextMonsterId <= 20, "too many monsters");
        require(_ac > 0, "monster too weak");
        monsters[nextMonsterId].name = _name;
        monsters[nextMonsterId].ac = _ac;
        emit AddMonster(nextMonsterId, _name, _ac);
        nextMonsterId = nextMonsterId.add(1);
    }
```

Here, we're just taking the name and AC value of a new monster and adding it
to the `monsters` list, then emitting an event. 
Players will select the ID of the monster they are fighting when 
calling the `rollForHit` function.

### changeStrModifier

This function will allow players to set their STR modifier. Anyone can
call, and the result will be stored with the `msg.sender`.

```solidity
    function changeStrModifier(uint256 _strMod) external {
        require(_strMod <= 5, "player too strong");
        strModifiers[msg.sender] = _strMod;
        emit ChangeStrModifier(msg.sender, _strMod);
    }
```

### getLastResult

This function allows queries to the contract to retrieve the last roll
result for a player/monster combo

```solidity
    function getLastResult(address _player, uint256 _monsterId) external view returns (Result memory) {
        return lastResult[_player][_monsterId];
    }
```

### unstickRoll

Because sometimes the dice rolls under the table...

```solidity
    function unstickRoll(address _player, uint256 _monsterId) external onlyOwner {
        lastResult[_player][_monsterId].isRolling = false;
    }
```

## Implementing VOR for Randomness

Now we have some support functions, we can implement the actual randomness
functionality and make the contract do something. We need to extend and
flesh out the `rollForHit` function, which will call the underlying `requestRandomness`
function in the `VORConsumerBase` contract, and the `fulfillRandomness`
function, which will receive and process the random number.

::: tip
The following two functions **are required** to interact with VOR and request
random numbers.
:::

### rollForHit

The `rollForHit` function is a wrapper around the required 
`VORConsumerBase.requestRandomness` function. The call to `requestRandomness`
is required in order to initialise a request to a VOR Provider Oracle.
Technically, this is the only requirement of the function, but we'll need
to do some pre-processing in order to track results and make use of the
returned value.

```solidity
    function rollForHit(uint256 _monsterId, uint256 _seed, bytes32 _keyHash, uint256 _fee) external returns (bytes32 requestId) {
        require(monsters[_monsterId].ac > 0, "monster does not exist");
        require(!lastResult[msg.sender][_monsterId].isRolling, "roll currently in progress");
        xFUND.transferFrom(msg.sender, address(this), _fee);
        requestId = requestRandomness(_keyHash, _fee, _seed);
        emit HittingMonster(_monsterId, requestId);
        requestIdToAddress[requestId] = msg.sender;
        requestIdToMonsterId[requestId] = _monsterId;
        lastResult[msg.sender][_monsterId].isRolling = true;
    }
```

Prior to calling `requestRandomness`, we're just ensuring the monster
is in the contract's database, and that the player/monster combo does 
not currently have a roll in progress. 

The next line transfers the required fee from the player calling the 
function to this DnD contract. The VORCoordinator then transfers
that fee from the DnD contract to itself for later forwarding to the 
VOR Provider Oracle.

::: tip
This is just one method for transferring fees to the VORCoordinator. 
Another is to omit `xFUND.transferFrom` from your request implementation,
and simply keep the contract topped up with xFUND (for example, if the
contract owner is the only address that will ever call the request function). 
In this case, it would also be advisable to implement a function to withdraw
xFUND from your contract.

However, since we want each player to pay their own fees, we are 
transferring as a part of the request process.
:::

Next, the function is making the actual `requestRandomness` call, which 
forwards the request to the `VORCoordinator`. It returns the generated
`requestId` which we then use to map some data about the request:

1. the address of the player who made the request
2. the monster th player is fighting.

When the request is fulfilled, the `requestId` is included in the
fulfilment so that we can retrieve any data associated with the request.

Finally, we set `isRolling` for the current player/monster combo to
prevent any further requests on this combo until it's fulfilled

### fulfillRandomness

The `fulfillRandomness` overrides the `virtual` function in 
`VORConsumerBase`. It is the function that will receive the random number
and ultimately process it.

::: tip
Whilst the `requestRandomness` function can be wrapped around an
arbitrary function, `fulfillRandomness` must be implemented as

`function fulfillRandomness(bytes32 requestId, uint256 randomness)`

This is because the `VORCoordinator`, and `VORConsumerBase` return 
data with specific parameters, and expect this function to be defined
as this in order to be able to correctly fulfil the request.

The internals of the function contain anything, however.
:::

```solidity
    function fulfillRandomness(bytes32 _requestId, uint256 _randomness) internal override {
        uint256 monsterId = requestIdToMonsterId[_requestId];
        address player = requestIdToAddress[_requestId];
        uint256 strModifier = strModifiers[player];
        uint256 roll = _randomness.mod(20).add(1);
        uint256 modified = roll.add(strModifier);
        string memory res = "miss";

        // Critical hit!
        if(roll == 20) {
            res = "nat20";
        } else if (roll == 1) {
            res = "nat1";
        } else {
            // Check roll + STR modifier against monster's AC
            if(modified >= monsters[monsterId].ac) {
                res = "hit";
            } else {
                res = "miss";
            }
        }
        emit HitResult(monsterId, _requestId, player, res, roll, modified);
  
        // store the results
        lastResult[player][monsterId].result = res;
        lastResult[player][monsterId].roll = roll;
        lastResult[player][monsterId].modified = modified;
        lastResult[player][monsterId].isRolling = false;

        // clean up
        delete requestIdToMonsterId[_requestId];
        delete requestIdToAddress[_requestId];
    }
```

The first three lines of our function are just retrieving some data related
to the request. This data was stored when the request was initialised.
It's simply getting the ID of the monster being hit, the player who is
hitting it, and then from that, deriving that player's STR modifier.

Next, the `_randomness` value returned from the VOR Provider Oracle is 
converted into a value between 1 and 20, simulating the roll of a d20.

The player's STR modifier is added to this value, and then checking
whether the player rolled high enough to hit the selected monster.

A roll of 20 is a Natural 20, and always hits (with crit!). 
A roll of 1 is a Natural 1, and always misses (boo!).

Otherwise, we check to see if the `roll + STR` modifier is enough to 
hit the monster.

The result is emitted in the `HitResult` event, and stored in `lastResult`.

## Helper function(s)

The final function we'll write uses the `_increaseVorCoordinatorAllowance` 
helper function included in `VORConsumerBase`. Its role is to permit 
the `VORCoordinator` smart contract to spend xFUNDMOCK on behalf of 
our smart contract, since the `VORCoordinator` needs to be able to 
transfer fees when each request for randomness is made.

For this function, we'll also use the OpenZeppelin `Ownable` contract's
`onlyOwner` modifier to ensure that only the contract owner can run this
function

```solidity
    function increaseVorAllowance(uint256 _amount) external onlyOwner {
        _increaseVorCoordinatorAllowance(_amount);
    }
```

This just calls `VORConsumerBase._increaseVorCoordinatorAllowance` function,
which in turn informs the `xFUNDMOCK` smart contract that we're allowing
`VORCoordinator` to spend `DnD`'s `xFUNDMOCK` tokens to pay for fees.

## Final contract

::: tip
Check out [https://github.com/unification-com/vor-demos/blob/main/contracts/DnD.sol](https://github.com/unification-com/vor-demos/blob/main/contracts/DnD.sol)
for the latest version of this demo contract.
:::

The final contract should look something like this:

```solidity
// SPDX-License-Identifier: MIT
pragma solidity 0.6.12;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
//import "@unification-com/xfund-vor/contracts/VORConsumerBase.sol";
import "../../xfund-vor/contracts/VORConsumerBase.sol";

/** ****************************************************************************
 * @notice Extremely simple DnD roll D20 to Hit using VOR
 * *****************************************************************************
 * @dev The contract owner can add up to 20 monsters. Players can modify their STR
 * modifier, which is pinned to their address. Players call the rollForHit function
 * and pay the associated xFUND fee to roll the D20. The result is returned in
 * fulfillRandomness, which calculates if the player crits, hits or misses.
 */
contract DnD is Ownable, VORConsumerBase {
    using SafeMath for uint256;

    // keep track of the monsters
    uint256 public nextMonsterId;

    // super simple monster stats
    struct Monster {
        string name;
        uint256 ac;
    }

    struct Result {
        uint256 roll;
        uint256 modified;
        string result;
        bool isRolling;
    }

    // monsters held in the contract
    mapping (uint256 => Monster) public monsters;
    // player STR modifiers
    mapping (address => uint256) public strModifiers;
    // map request IDs to monster IDs
    mapping(bytes32 => uint256) public requestIdToMonsterId;
    // map request IDs to player addresses, to retrieve STR modifiers
    mapping(bytes32 => address) public requestIdToAddress;
    // store last [player][monster] results
    mapping(address => mapping(uint256 => Result)) lastResult;

    // Some useful events to track
    event AddMonster(uint256 monsterId, string name, uint256 ac);
    event ChangeStrModifier(address player, uint256 strMod);
    event HittingMonster(uint256 monsterId, bytes32 requestId);
    event HitResult(uint256 monsterId, bytes32 requestId, address player, string result, uint256 roll, uint256 modified);

    /**
    * @notice Constructor inherits VORConsumerBase
    *
    * @param _vorCoordinator address of the VOR Coordinator
    * @param _xfund address of the xFUND token
    */
    constructor(address _vorCoordinator, address _xfund)
    public
    VORConsumerBase(_vorCoordinator, _xfund) {
        nextMonsterId = 1;
    }

    /**
    * @notice addMonster can be called by the owner to add a new monster
    *
    * @param _name string name of the monster
    * @param _ac uint256 AC of the monster
    */
    function addMonster(string memory _name, uint256 _ac) external onlyOwner {
        require(nextMonsterId <= 20, "too many monsters");
        require(_ac > 0, "monster too weak");
        monsters[nextMonsterId].name = _name;
        monsters[nextMonsterId].ac = _ac;
        emit AddMonster(nextMonsterId, _name, _ac);
        nextMonsterId = nextMonsterId.add(1);
    }

    /**
    * @notice changeStrModifier can be called by anyone to change their STR modifier
    *
    * @param _strMod uint256 STR modifier of player
    */
    function changeStrModifier(uint256 _strMod) external {
        require(_strMod <= 5, "player too strong");
        strModifiers[msg.sender] = _strMod;
        emit ChangeStrModifier(msg.sender, _strMod);
    }

    /**
    * @notice rollForHit anyone can call to roll the D20 for hit. Caller (msg.sender)
    * pays the xFUND fees for the request.
    *
    * @param _monsterId uint256 Id of the monster the caller is fighting
    * @param _seed uint256 seed for the randomness request. Gets mixed in with the blockhash of the block this Tx is in
    * @param _keyHash bytes32 key hash of the provider caller wants to fulfil the request
    * @param _fee uint256 required fee amount for the request
    */
    function rollForHit(uint256 _monsterId, uint256 _seed, bytes32 _keyHash, uint256 _fee) external returns (bytes32 requestId) {
        require(monsters[_monsterId].ac > 0, "monster does not exist");
        require(!lastResult[msg.sender][_monsterId].isRolling, "roll currently in progress");
        // Note - caller must have increased xFUND allowance for this contract first.
        // Fee is transferred from msg.sender to this contract. The VORCoordinator.requestRandomness
        // function will then transfer from this contract to itself.
        // This contract's owner must have increased the VORCoordnator's allowance for this contract.
        xFUND.transferFrom(msg.sender, address(this), _fee);
        requestId = requestRandomness(_keyHash, _fee, _seed);
        emit HittingMonster(_monsterId, requestId);
        requestIdToAddress[requestId] = msg.sender;
        requestIdToMonsterId[requestId] = _monsterId;
        lastResult[msg.sender][_monsterId].isRolling = true;
    }

    /**
     * @notice Callback function used by VOR Coordinator to return the random number
     * to this contract.
     * @dev The random number is used to simulate a D20 roll. Result is emitted as follows:
     * 1: Natural 1...
     * 20: Natural 20!
     * roll + strModifier >= monster AC: hit
     * roll + strModifier < monster AC: miss
     *
     * @param _requestId bytes32
     * @param _randomness The random result returned by the oracle
     */
    function fulfillRandomness(bytes32 _requestId, uint256 _randomness) internal override {
        uint256 monsterId = requestIdToMonsterId[_requestId];
        address player = requestIdToAddress[_requestId];
        uint256 strModifier = strModifiers[player];
        uint256 roll = _randomness.mod(20).add(1);
        uint256 modified = roll.add(strModifier);
        string memory res = "miss";

        // Critical hit!
        if(roll == 20) {
            res = "nat20";
        } else if (roll == 1) {
            res = "nat1";
        } else {
            // Check roll + STR modifier against monster's AC
            if(modified >= monsters[monsterId].ac) {
                res = "hit";
            } else {
                res = "miss";
            }
        }
        emit HitResult(monsterId, _requestId, player, res, roll, modified);

        // store the results
        lastResult[player][monsterId].result = res;
        lastResult[player][monsterId].roll = roll;
        lastResult[player][monsterId].modified = modified;
        lastResult[player][monsterId].isRolling = false;

        // clean up
        delete requestIdToMonsterId[_requestId];
        delete requestIdToAddress[_requestId];
    }

    /**
     * @notice getLastResult returns the last result for a specified player/monsterId.
     *
     * @param _player address address of player
     * @param _monsterId uint256 id of monster
     */
    function getLastResult(address _player, uint256 _monsterId) external view returns (Result memory) {
        return lastResult[_player][_monsterId];
    }

    /**
    * @notice unstickRoll allows contract owner to unstick a roll when a request is not fulfilled
    *
    * @param _player address address of player
    * @param _monsterId uint256 id of monster
    */
    function unstickRoll(address _player, uint256 _monsterId) external onlyOwner {
        lastResult[_player][_monsterId].isRolling = false;
    }

    /**
     * @notice Example wrapper function for the VORConsumerBase increaseVorCoordinatorAllowance function.
     * @dev Wrapped around an Ownable modifier to ensure only the contract owner can call it.
     * @dev Allows contract owner to increase the xFUND allowance for the VORCoordinator contract
     * @dev enabling it to pay request fees on behalf of this contract.
     *
     * @param _amount uint256 amount to increase allowance by
     */
    function increaseVorAllowance(uint256 _amount) external onlyOwner {
        _increaseVorCoordinatorAllowance(_amount);
    }
}
```

Next, we'll look at how to [interact](./interaction.md) with the contract,
request randomness and view the results.
