# BlockhashStore



## Functions:
- [`storeEarliest() external`](#BlockhashStore-storeEarliest--)
- [`getBlockhash(uint256 n) external`](#BlockhashStore-getBlockhash-uint256-)
- [`store(uint256 n) public`](#BlockhashStore-store-uint256-)
- [`storeVerifyHeader(uint256 n, bytes header) public`](#BlockhashStore-storeVerifyHeader-uint256-bytes-)



<a name="BlockhashStore-storeEarliest--"></a>
### Function `storeEarliest() external `
No description
<a name="BlockhashStore-getBlockhash-uint256-"></a>
### Function `getBlockhash(uint256 n) external  -> bytes32`
No description
#### Parameters:
- `n`: the number of the block whose blockhash should be returned
<a name="BlockhashStore-store-uint256-"></a>
### Function `store(uint256 n) public `
No description
#### Parameters:
- `n`: the number of the block whose blockhash should be stored
<a name="BlockhashStore-storeVerifyHeader-uint256-bytes-"></a>
### Function `storeVerifyHeader(uint256 n, bytes header) public `
No description
#### Parameters:
- `n`: the number of the block whose blockhash should be stored

- `header`: the rlp-encoded blockheader of block n+1. We verify its correctness by checking
  that it hashes to a stored blockhash, and then extract parentHash to get the n-th blockhash.


