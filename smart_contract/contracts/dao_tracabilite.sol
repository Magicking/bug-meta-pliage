pragma solidity ^0.4.19;

contract DAOTracabilite {

	// Mapping 1-to-1 hash <=> Objects
	mapping (bytes32 => Object) public Objects;
	// Mapping 1-to-1 hash(id) <=> id in string format
	mapping (bytes32 => string) public Indexes;
	// Mapping 1-to-1 hash(id) <=> id in string format
	mapping (bytes32 => Point) public ContainerList;

	enum Actions { Pack, Unpack, Transfer }
	enum ObjectKind { QRCode, Unit, Container }

	// Basic object
	struct Object {
		uint       Date;
		ObjectKind Kind;
		string     Metadata;
	}

	// Historical traces of units or container
	struct Point {
		bytes32 From;
		bytes32 To;
		uint    Date;
		Actions Kind;
		string  Metadata;
	}

	function DAOTracabilite() public {
	}

	function toInternalId(string guid) pure public returns (bytes32) {
		return keccak256(guid);
	}

	function addInternal(string guid, string metadata, ObjectKind kind) internal returns (bool) {
		bytes32 id = toInternalId(guid);
		Indexes[id] = guid;
		Objects[id].Date = now;
		Objects[id].Kind = kind;
		Objects[id].Metadata = metadata;
	}
	function AddQRCode(string guid, string metadata) public returns (bool){
		return addInternal(guid, metadata, ObjectKind.QRCode);
	}

	function AddUnit(string guid, string metadata) public returns (bool) {
		return addInternal(guid, metadata, ObjectKind.Unit);
	}

	function AddContainer(string guid, string metadata) public returns (bool){
		return addInternal(guid, metadata, ObjectKind.Container);
	}

	function AddObjectToContainer(string guidFrom, string guidTo, string metadata) public returns (bool) {
		bytes32 fromID = toInternalId(guidFrom);
		bytes32 toID = toInternalId(guidTo);
		ContainerList[fromID].From = fromID;
		ContainerList[fromID].To = toID;
		ContainerList[fromID].Date = now;
		ContainerList[fromID].Kind = Actions.Pack;
		ContainerList[fromID].Metadata = metadata;
		return true;
	}

	function GetHistory(string guid) public view returns (bool) {
		Point[] memory pts = new Point[](2);

		bytes32 id = toInternalId(guid);
		pts[0].From = id;
		pts[0].To = id;
		pts[0].Date = now;
		pts[0].Kind = Actions.Transfer;
		pts[0].Metadata = "";
		pts[1].From = id;
		pts[1].To = id;
		pts[1].Date = now;
		pts[1].Kind = Actions.Transfer;
		pts[1].Metadata = "";
		return true;
	}
	/*
	function TransferContainer(string guid, string fromID, string toID, string metadata) public returns (bool) {
		return true;
	}*/
}
