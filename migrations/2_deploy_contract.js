const VORCoordinator = artifacts.require("VORCoordinator");

module.exports = function(deployer) {
  deployer.deploy(VORCoordinator);
};