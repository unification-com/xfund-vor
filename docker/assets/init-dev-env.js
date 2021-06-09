const { BN } = require('@openzeppelin/test-helpers')

const VORCoordinator = artifacts.require("VORCoordinator")
const MockERC20 = artifacts.require("MockERC20")

module.exports = async function(callback) {

    const vor = await VORCoordinator.deployed()
    const token = await MockERC20.deployed()

    const publicProvingKey = [
        new BN("67315240764067688871012716141531292970097069926279872171462765639266213797159"),
        new BN("44851436639087445619644214552761311037004527461759145665315091825514767327704")
    ]
    const accounts = await web3.eth.getAccounts()
    const owner = accounts[0]
    const oracleAcc = accounts[1]

    console.log("owner:", owner)
    console.log("oracleAcc:", oracleAcc)

    for(let i = 2; i < accounts.length; i += 1) {
        console.log("transfer tokens to", accounts[i])
        try {
            await token.transfer(accounts[i], 100000000000, {from: owner})
        } catch (e) {
            console.error(e)
        }
    }

    console.log("register new proving key for", oracleAcc)
    try {
        const newServiceAgreement = await vor.registerProvingKey(100000000, oracleAcc, publicProvingKey)
        console.log("registered keyHash:", newServiceAgreement.logs[0].args.keyHash)
    } catch (e) {
        console.error(e)
    }

    console.log("done")

    callback()
}
