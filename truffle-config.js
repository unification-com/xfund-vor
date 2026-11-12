require("dotenv").config()
const HDWalletProvider = require("@truffle/hdwallet-provider")
const TestRPC = require("ganache-cli")

const {
    ETH_PKEY_TESTNET,
    ETH_PKEY_MAINNET,
    INFURA_PROJECT_ID,
    ETHERSCAN_API,
} = process.env

module.exports = {
    networks: {
        development: {
            provider: TestRPC.provider({
                mnemonic: "myth like bonus scare over problem client lizard pioneer submit female collect",
                gasPrice: "0x4A817C800",
            }),
            network_id: "*", // Any network (default: none)
        },
        // truffle develop console
        develop: {
            host: "127.0.0.1",
            port: 8545,
            network_id: "*",
            defaultEtherBalance: 500,
        },
        goerli: {
            provider: () =>
                new HDWalletProvider({
                    privateKeys: [ETH_PKEY_TESTNET],
                    providerOrUrl: `https://goerli.infura.io/v3/${INFURA_PROJECT_ID}`,
                }),
            network_id: "5",
            gas: 10000000,
            gasPrice: 2000000000,
            skipDryRun: true,
        },
        sepolia: {
            provider: () =>
                new HDWalletProvider({
                    privateKeys: [ETH_PKEY_TESTNET],
                    providerOrUrl: `https://sepolia.infura.io/v3/${INFURA_PROJECT_ID}`,
                }),
            network_id: "11155111",
            gas: 10000000,
            gasPrice: 5000000000,
            skipDryRun: true,
        },
        mainnet: {
            provider: () =>
                new HDWalletProvider({
                    privateKeys: [ETH_PKEY_MAINNET],
                    providerOrUrl: `https://mainnet.infura.io/v3/${INFURA_PROJECT_ID}`,
                }),
            network_id: "1",
            gasPrice: 40000000000,
        },
        polygon: {
            provider: () =>
                new HDWalletProvider({
                    privateKeys: [ETH_PKEY_MAINNET],
                    providerOrUrl: `https://polygon-mainnet.infura.io/v3/${INFURA_PROJECT_ID}`,
                }),
            network_id: "137",
            gasPrice: 40000000000,
        },
        polygon_mumbai: {
            provider: () =>
                new HDWalletProvider({
                    privateKeys: [ETH_PKEY_TESTNET],
                    providerOrUrl: `https://polygon-mumbai.infura.io/v3/${INFURA_PROJECT_ID}`,
                }),
            network_id: "80001",
            gasPrice: 3000000000,
        },
        shibarium: {
            provider: () =>
                new HDWalletProvider({
                    privateKeys: [ETH_PKEY_MAINNET],
                    providerOrUrl: 'https://rpc.shibarium.shib.io',
                }),
            network_id: "109",
            gasPrice: 10000000000, // 10 gwei
            skipDryRun: true,
        },
        qom: {
            provider: () =>
                new HDWalletProvider({
                    privateKeys: [ETH_PKEY_MAINNET],
                    providerOrUrl: 'https://evm-rpc-ql1.foxxone.one',
                }),
            network_id: "766",
            gasPrice: 1000000000, // 1 gwei
            skipDryRun: true,
        },
    },
    plugins: ["truffle-plugin-verify", "solidity-coverage", "@chainsafe/truffle-plugin-abigen"],
    api_keys: {
        etherscan: ETHERSCAN_API,
    },
    compilers: {
        solc: {
            version: '0.6.12',
            settings: {
                optimizer: {
                    enabled: true,
                    runs: 200,
                }
            }
        }
    }
}
