const MockERC20 = artifacts.require("MockERC20");
const BlockhashStore = artifacts.require("BlockhashStore");
const VORCoordinator = artifacts.require("VORCoordinator");
const VORD20 = artifacts.require("VORD20");

module.exports = function(deployer) {
  deployer.then(async () => {
    // const xfund = "0x0000000000000000000000000000000000000000"
    // When deploying to mainnet, you need to change to the real address of the contract
    const erc20 = await deployer.deploy(MockERC20, 'xFUND', 'xFUND', web3.utils.toWei('1000000000', 'ether'));
    const xfund = erc20.address

    const block = await deployer.deploy(BlockhashStore);
    const vor = await deployer.deploy(VORCoordinator, xfund, block.address);

    // For tests, you must substitute the correct values
    await deployer.deploy(VORD20, vor.address, xfund, web3.utils.fromAscii('keyHash'), web3.utils.toWei('1', 'ether'));
  });
};