var ConvertLib = artifacts.require("./ConvertLib.sol");
var DAOT = artifacts.require("./dao_tracabilite.sol");

module.exports = function(deployer) {
  deployer.deploy(ConvertLib);
  deployer.link(ConvertLib, DAOT);
  deployer.deploy(DAOT);
};
