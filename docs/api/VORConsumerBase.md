# VORConsumerBase

PURPOSE

Reggie the Random Oracle (not his real job) wants to provide randomness
to Vera the verifier in such a way that Vera can be sure he's not
making his output up to suit himself. Reggie provides Vera a public key
to which he knows the secret key. Each time Vera provides a seed to
Reggie, he gives back a value which is computed completely
deterministically from the seed and the secret key.

Reggie provides a proof by which Vera can verify that the output was
correctly computed once Reggie tells it to her, but without that proof,
the output is indistinguishable to her from a uniform random sample
from the output space.

The purpose of this contract is to make it easy for unrelated contracts
to talk to Vera the verifier about the work Reggie is doing, to provide
simple access to a verifiable source of randomness.

USAGE

Calling contracts must inherit from VORConsumerBase, and can
initialize VORConsumerBase's attributes in their constructor as
shown:

```
  contract VORConsumer {
    constuctor(<other arguments>, address _vorCoordinator, address _xfund)
      VORConsumerBase(_vorCoordinator, _xfund) public {
        <initialization with other arguments goes here>
      }
  }
```
The oracle will have given you an ID for the VOR keypair they have
committed to (let's call it keyHash), and have told you the minimum xFUND
price for VOR service. Make sure your contract has sufficient xFUND, and
call requestRandomness(keyHash, fee, seed), where seed is the input you
want to generate randomness from.

Once the VORCoordinator has received and validated the oracle's response
to your request, it will call your contract's fulfillRandomness method.

The randomness argument to fulfillRandomness is the actual random value
generated from your seed.

The requestId argument is generated from the keyHash and the seed by
makeRequestId(keyHash, seed). If your contract could have concurrent
requests open, you can use the requestId to track which seed is
associated with which randomness. See VORRequestIDBase.sol for more
details. (See "SECURITY CONSIDERATIONS" for principles to keep in mind,
if your contract could have multiple requests in flight simultaneously.)

Colliding `requestId`s are cryptographically impossible as long as seeds
differ. (Which is critical to making unpredictable randomness! See the
next section.)

SECURITY CONSIDERATIONS

A method with the ability to call your fulfillRandomness method directly
could spoof a VOR response with any random value, so it's critical that
it cannot be directly called by anything other than this base contract
(specifically, by the VORConsumerBase.rawFulfillRandomness method).

For your users to trust that your contract's random behavior is free
from malicious interference, it's best if you can write it so that all
behaviors implied by a VOR response are executed *during* your
fulfillRandomness method. If your contract must store the response (or
anything derived from it) and use it later, you must ensure that any
user-significant behavior which depends on that stored value cannot be
manipulated by a subsequent VOR request.

Similarly, both miners and the VOR oracle itself have some influence
over the order in which VOR responses appear on the blockchain, so if
your contract could have multiple VOR requests in flight simultaneously,
you must ensure that the order in which the VOR responses arrive cannot
be used to manipulate your contract's user-significant behavior.

Since the ultimate input to the VOR is mixed with the block hash of the
block in which the request is made, user-provided seeds have no impact
on its economic security properties. They are only included for API
compatability with previous versions of this contract.

Since the block hash of the block which contains the requestRandomness
call is mixed into the input to the VOR *last*, a sufficiently powerful
miner could, in principle, fork the blockchain to evict the block
containing the request, forcing the request to be included in a
different block with a different hash, and therefore a different input
to the VOR. However, such an attack would incur a substantial economic
cost. This cost scales with the number of blocks the VOR oracle waits
until it calls responds to a request.

## Functions:
- [`fulfillRandomness(bytes32 requestId, uint256 randomness) internal`](#VORConsumerBase-fulfillRandomness-bytes32-uint256-)
- [`requestRandomness(bytes32 _keyHash, uint256 _fee, uint256 _seed) internal`](#VORConsumerBase-requestRandomness-bytes32-uint256-uint256-)
- [`_increaseVorCoordinatorAllowance(uint256 _amount) internal`](#VORConsumerBase-_increaseVorCoordinatorAllowance-uint256-)
- [`_withdrawEth(address _to, uint256 _amount) internal`](#VORConsumerBase-_withdrawEth-address-uint256-)
- [`_withdrawXFUND(address _to, uint256 _amount) internal`](#VORConsumerBase-_withdrawXFUND-address-uint256-)
- [`constructor(address _vorCoordinator, address _xfund) public`](#VORConsumerBase-constructor-address-address-)
- [`rawFulfillRandomness(bytes32 requestId, uint256 randomness) external`](#VORConsumerBase-rawFulfillRandomness-bytes32-uint256-)
- [`receive() external`](#VORConsumerBase-receive--)



<a name="VORConsumerBase-fulfillRandomness-bytes32-uint256-"></a>
### Function `fulfillRandomness(bytes32 requestId, uint256 randomness) internal `
VORConsumerBase expects its subcontracts to have a method with this
signature, and will call it once it has verified the proof
associated with the randomness. (It is triggered via a call to
rawFulfillRandomness, below.)


#### Parameters:
- `requestId`: The Id initially returned by requestRandomness

- `randomness`: the VOR output
<a name="VORConsumerBase-requestRandomness-bytes32-uint256-uint256-"></a>
### Function `requestRandomness(bytes32 _keyHash, uint256 _fee, uint256 _seed) internal  -> bytes32 requestId`
The fulfillRandomness method receives the output, once it's provided
by the Oracle, and verified by the vorCoordinator.

The _keyHash must already be registered with the VORCoordinator, and
the _fee must exceed the fee specified during registration of the
_keyHash.

The _seed parameter is vestigial, and is kept only for API
compatibility with older versions. It can't *hurt* to mix in some of
your own randomness, here, but it's not necessary because the VOR
oracle will mix the hash of the block containing your request into the
VOR seed it ultimately uses.


#### Parameters:
- `_keyHash`: ID of public key against which randomness is generated

- `_fee`: The amount of xFUND to send with the request

- `_seed`: seed mixed into the input of the VOR.


#### Return Values:
- requestId unique ID for this request

The returned requestId can be used to distinguish responses to
concurrent requests. It is passed as the first argument to
fulfillRandomness.
<a name="VORConsumerBase-_increaseVorCoordinatorAllowance-uint256-"></a>
### Function `_increaseVorCoordinatorAllowance(uint256 _amount) internal  -> bool`
No description
<a name="VORConsumerBase-_withdrawEth-address-uint256-"></a>
### Function `_withdrawEth(address _to, uint256 _amount) internal  -> bool success`
NOTE: this functions should be wrapped around a, for example,
Ownable function such that only the contract's owner can call it.


#### Parameters:
- `_to`: address to send the eth to

- `_amount`: uint256 amount to withdraw
<a name="VORConsumerBase-_withdrawXFUND-address-uint256-"></a>
### Function `_withdrawXFUND(address _to, uint256 _amount) internal `
NOTE: this functions should be wrapped around a, for example,
Ownable function such that only the contract's owner can call it.


#### Parameters:
- `_to`: the address to withdraw xFUND to

- `_amount`: the amount of xFUND to withdraw
<a name="VORConsumerBase-constructor-address-address-"></a>
### Function `constructor(address _vorCoordinator, address _xfund) public `
No description
#### Parameters:
- `_vorCoordinator`: address of VORCoordinator contract

- `_xfund`: address of xFUND token contract
<a name="VORConsumerBase-rawFulfillRandomness-bytes32-uint256-"></a>
### Function `rawFulfillRandomness(bytes32 requestId, uint256 randomness) external `
No description
<a name="VORConsumerBase-receive--"></a>
### Function `receive() external `
No description


