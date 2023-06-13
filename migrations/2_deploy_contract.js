const MockERC20 = artifacts.require("MockERC20");
const TestXFUND = artifacts.require("TestXFUND");
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
        case "goerli":
        case "goerli-fork":
            deployer.then(async () => {
                // const block = await deployer.deploy(BlockhashStore);
                await deployer.deploy(VORCoordinator, "0xb07C72acF3D7A5E9dA28C56af6F93862f8cc8196", "0xC1e1a7f39fB6E3E1FBAa5d33407F7844e5C843Ff");
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
        case "polygon":
        case "polygon-fork":
            deployer.then(async () => {
                // const block = await deployer.deploy(BlockhashStore);
                await deployer.deploy(VORCoordinator, "0x77a3840f78e4685afaf9c416b36e6eae6122567b", "0x2E9ade949900e19735689686E61BF6338a65B881");
            });
            break
        case "polygon_mumbai":
        case "polygon_mumbai-fork":
            deployer.then(async () => {
                const testXfund = await deployer.deploy(TestXFUND, 'xFUND', 'xFUND', web3.utils.toWei('100', 'ether'), 9);
                const block = await deployer.deploy(BlockhashStore);
                await deployer.deploy(VORCoordinator, testXfund.address, block.address);
            });
            break
        case "sepolia":
        case "sepolia-fork":
            deployer.then(async () => {
                const block = await deployer.deploy(BlockhashStore);
                await deployer.deploy(VORCoordinator, "0xb07C72acF3D7A5E9dA28C56af6F93862f8cc8196", block.address);
            });
            break
    }

};
