pragma solidity ^0.4.20;

contract DAOTracabilite {

	// Mapping 1-to-1 hash <=> Objects
	mapping (bytes32 => Object) public Objects;
	// Mapping 1-to-1 hash(id) <=> id in string format
	mapping (bytes32 => string) public Indexes;
	// Mapping 1-to-1 hash(id) <=> id in string format
	mapping (bytes32 => Point) public ContainerList;
	// Mapping 1-to-1 hash(id+idx) => Point
	mapping (bytes32 => Point) public Traces;

	enum Actions { Create, Pack, TransferFrom, TransferTo }
	enum ObjectKind { QRCode, Unit, Container }

	event ObjectNotification(ObjectKind kind, string metadata);
	event ActionNotification(Actions action, string from, string metadata);

	// Basic object
	struct Object {
		uint       Date;
		ObjectKind Kind;
		string     Metadata;
		uint       TraceCount;
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
		Objects[id].TraceCount = 0;
		ObjectNotification(kind, metadata);
	}

	function AddQRCode(string guid, string metadata) public returns (bool) {
		return addInternal(guid, metadata, ObjectKind.QRCode);
	}

	function AddUnit(string guid, string metadata) public returns (bool) {
		return addInternal(guid, metadata, ObjectKind.Unit);
	}

	function AddContainer(string guid, string metadata) public returns (bool) {
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
		AddTrace(guidFrom, fromID, toID, Actions.TransferFrom, metadata);
		AddTrace(guidTo, fromID, toID, Actions.TransferTo, metadata);
		Objects[fromID].TraceCount++;
		return true;
	}

	function AddTrace(string fromStr, bytes32 from, bytes32 to, Actions kind, string metadata) internal {
		Point memory pts;
		pts.From = from;
		pts.To = to;
		pts.Date = now;
		pts.Kind = kind;
		pts.Metadata = metadata;

		bytes32 idx = GetHistoryHash(fromStr, Objects[from].TraceCount);
		Traces[idx] = pts;
		ActionNotification(kind, fromStr, metadata);
	}

	// Call GetHistoryHash
	function GetHistoryHash(string guid, uint index) public pure returns (bytes32 traceHash) {
		uint maxlength = 100;
		bytes memory reversed = new bytes(maxlength);
		uint i = 0;
		index++;
		while (index != 0) {
			uint remainder = index % 10;
			index = index / 10;
			reversed[i++] = byte(48 + remainder);
		}
		bytes memory guidb = bytes(guid);
		bytes memory s = new bytes(guidb.length + i);
		uint j;
		for (j = 0; j < guidb.length; j++) {
			s[j] = guidb[j];
		}
		for (j = 0; j < i; j++) {
			s[j + guidb.length] = reversed[i - 1 - j];
		}
		return keccak256(string(s));
	}

	/*
	function TransferContainer(string guid, string fromID, string toID, string metadata) public returns (bool) {
		return true;
	}*/
}
