const MockERC20 = artifacts.require("MockERC20");
const BlockhashStore = artifacts.require("BlockhashStore");
const VORCoordinator = artifacts.require("VORCoordinator");
const VORD20 = artifacts.require("VORD20");
const VORRandomnessRequestMock = artifacts.require("VORRandomnessRequestMock")

module.exports = function (deployer, network) {
    switch (network) {
        default:
        case "development":
        case "develop":
            deployer.then(async () => {
                const erc20 = await deployer.deploy(MockERC20, 'xFUND', 'xFUND', web3.utils.toWei('1000000000', 'ether'));
                const xfund = erc20.address

                const block = await deployer.deploy(BlockhashStore);
                const vor = await deployer.deploy(VORCoordinator, xfund, block.address);

                // For tests, you must substitute the correct values
                await deployer.deploy(VORD20, vor.address, xfund, web3.utils.fromAscii('keyHash'), web3.utils.toWei('1', 'ether'));
                await deployer.deploy(VORRandomnessRequestMock);
            });
            break
        case "rinkeby":
        case "rinkeby-fork":
            deployer.then(async () => {
                // const block = await deployer.deploy(BlockhashStore);
                await deployer.deploy(VORCoordinator, "0x245330351344f9301690d5d8de2a07f5f32e1149", "0x95AE62E3E2261615970375CC8af8c7E6923627Fa");
            });
            break
        case "mainnet":
        case "mainnet-fork":
            deployer.then(async () => {
                // const block = await deployer.deploy(BlockhashStore);
                await deployer.deploy(VORCoordinator, "0x892A6f9dF0147e5f079b0993F486F9acA3c87881", "0x540FCdd99F4EC8cDac1345D152857B1B20e4d5f6");
            });
            break
    }

};