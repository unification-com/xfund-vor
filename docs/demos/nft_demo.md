# NFT Competition Demo

Our first demo is running on Rinkeby testnet, and showcases the potential of VOR by implementing
an NFT competition. Competitions can be created, with an NFT as the prize. Entrants pay a
fee in xFUND to enter, and the winner is selected via VOR. The NFT is then minted for the
winner.

**Contract**: [0xE1426CE899537340E5551cF37Db813B75Ec6C579](https://rinkeby.etherscan.io/address/0xE1426CE899537340E5551cF37Db813B75Ec6C579#code)  
**Github**: [https://github.com/unification-com/vor-demos](https://github.com/unification-com/vor-demos)

## Interaction

For the moment, interaction can be done via Etherscan (UI is forthcoming).

### View a Competition

There is currently only one competition running at the moment, for this wonderful 
[unique piece of art](https://gateway.pinata.cloud/ipfs/QmNuNkxxhdbWAGJkAESvcTPVZZevCKnTfnWtquRmAUzuBj).

The competition has a maximum of 100 entries, and costs 0.1 xFUNDMOCK to enter.

To see the competition details and current entrants:

1. Go to [NFT Competition Contract](https://rinkeby.etherscan.io/address/0xE1426CE899537340E5551cF37Db813B75Ec6C579#readContract) on
Etherscan and click the "Read Contract" button.
2. Select the `getCompetition` function
3. Enter `1` as the `_competitionId`
4. Click the "Query" button

### Entering a competition

To enter a competition, you'll need some xFUNDMOCK tokens, and to approve the NFT competition contract to spend some for you
to pay for entry fees.

#### Getting xFUNDMOCK tokens

1. Go to the [xFUNDMOCK](https://rinkeby.etherscan.io/address/0x245330351344F9301690D5D8De2A07f5F32e1149#writeContract) contract
2. Click the "Connect to Web3" button and connect your wallet.
3. Click "gimme" function, followed by "Write"

This should give you 10 `xFUNDMOCK` tokens.

#### Approving the NFT Contract to pay entry fees

1. Go to the [xFUNDMOCK](https://rinkeby.etherscan.io/address/0x245330351344F9301690D5D8De2A07f5F32e1149#writeContract) contract
2. Click the "Connect to Web3" button and connect your wallet.
3. Click the "increaseAllowance" function.
4. Enter `0xE1426CE899537340E5551cF37Db813B75Ec6C579` as the `spender` (the NFT Competition contract address)
5. Enter `5000000000` (5 xFUNDMOCK) as the `addedValue`
6. Click "Write"

You have now approved the NFT contract to transfer a small amount of xFUNDMOCK on your behalf.

#### Enter the Competition

Now we're ready to enter the competition.

1. Go to [NFT Competition Contract](https://rinkeby.etherscan.io/address/0xE1426CE899537340E5551cF37Db813B75Ec6C579#writeContract) on
   Etherscan and click the "Write Contract" button.
2. Click the "Connect to Web3" button and connect your wallet.
3. Click the `enterCompetition` function.
4. Enter `1` as the `_competitionId`
5. Click the "Write" button.

Once your transaction has been confirmed, you can repeat the steps in the "**View a Competition**" section and you
should see your wallet address in the list of entrants.

## Wen Winner?

The winner will be selected at some point in the future...

In the meantime, check out our [guides](../guide/index.md) on deploying your own VOR-enable smart contract, or the 
[Github repo](https://github.com/unification-com/vor-demos) for this demo.
