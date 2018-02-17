pragma solidity ^0.4.4;

import "./ConvertLib.sol";

contract DAOTracabilite {
	address owner

	struct Unit {
		string name
	}

	struct Container {
		string name
	}

	function DAOTracabilite() {
		owner = msg.sender;
	}

	function AddUnit(string name, string metadata) returns(uint hash) {
	}

	function AddContainer(string name, string metadata) returns(uint hash){
	}
	
	//PACK
	//UNPACK
	function AddUnitToContainer(uint unitHash, uint containerHash) returns(bool ok) {
	}

	function GetMetadata(uint hash) returns(string metadata) {
	}

	function GetInformation(uint hash) returns(string information) {
	}
}
