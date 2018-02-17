// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internal

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DAOTracabiliteABI is the input ABI used to generate the binding from.
const DAOTracabiliteABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Objects\",\"outputs\":[{\"name\":\"Date\",\"type\":\"uint256\"},{\"name\":\"Kind\",\"type\":\"uint8\"},{\"name\":\"Metadata\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"AddQRCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"}],\"name\":\"toInternalId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"}],\"name\":\"GetHistory\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Indexes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ContainerList\",\"outputs\":[{\"name\":\"From\",\"type\":\"bytes32\"},{\"name\":\"To\",\"type\":\"bytes32\"},{\"name\":\"Date\",\"type\":\"uint256\"},{\"name\":\"Kind\",\"type\":\"uint8\"},{\"name\":\"Metadata\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guidFrom\",\"type\":\"string\"},{\"name\":\"guidTo\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"AddObjectToContainer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"AddContainer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"AddUnit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DAOTracabiliteBin is the compiled bytecode used for deploying new contracts.
const DAOTracabiliteBin = `0x6060604052341561000f57600080fd5b610b718061001e6000396000f3006060604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630b181f1c811461009d57806336ecf9081461015c5780636d325d8c1461020357806382d2934714610266578063bc863147146102b7578063bf5af74214610344578063cc6d7b2c14610413578063d09cf216146104e8578063dedde18f1461057b575b600080fd5b34156100a857600080fd5b6100b360043561060e565b604051808481526020018360028111156100c957fe5b60ff168152604082820381016020830190815284546002600182161561010002600019019091160491830182905291606001908490801561014b5780601f106101205761010080835404028352916020019161014b565b820191906000526020600020905b81548152906001019060200180831161012e57829003601f168201915b505094505050505060405180910390f35b341561016757600080fd5b6101ef60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061063195505050505050565b604051901515815260200160405180910390f35b341561020e57600080fd5b61025460046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061064695505050505050565b60405190815260200160405180910390f35b341561027157600080fd5b6101ef60046024813581810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506106ab95505050505050565b34156102c257600080fd5b6102cd60043561087a565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156103095780820151838201526020016102f1565b50505050905090810190601f1680156103365780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561034f57600080fd5b61035a60043561092a565b60405185815260208101859052604081018490526060810183600281111561037e57fe5b60ff16815260408282038101602083019081528454600260018216156101000260001901909116049183018290529160600190849080156104005780601f106103d557610100808354040283529160200191610400565b820191906000526020600020905b8154815290600101906020018083116103e357829003601f168201915b5050965050505050505060405180910390f35b341561041e57600080fd5b6101ef60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061095a95505050505050565b34156104f357600080fd5b6101ef60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506109c895505050505050565b341561058657600080fd5b6101ef60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506109d695505050505050565b600060208190529081526040902080546001820154909160ff9091169060020183565b600061063f838360006109e0565b9392505050565b6000816040518082805190602001908083835b602083106106785780518252601f199092019160209182019101610659565b6001836020036101000a038019825116818451161790925250505091909101925060409150505180910390209050919050565b60006106b5610a63565b600060026040518059106106c65750595b9080825280602002602001820160405280156106fc57816020015b6106e9610a75565b8152602001906001900390816106e15790505b50915061070884610646565b9050808260008151811061071857fe5b9060200190602002015152808260008151811061073157fe5b9060200190602002015160200152428260008151811061074d57fe5b906020019060200201516040015260028260008151811061076a57fe5b9060200190602002015160600190600281111561078357fe5b9081600281111561079057fe5b905250602060405190810160405260008082528390815181106107af57fe5b906020019060200201516080015280826001815181106107cb57fe5b906020019060200201515280826001815181106107e457fe5b9060200190602002015160200152428260018151811061080057fe5b906020019060200201516040015260028260018151811061081d57fe5b9060200190602002015160600190600281111561083657fe5b9081600281111561084357fe5b9052506020604051908101604052600081528260018151811061086257fe5b90602001906020020151608001525060019392505050565b60016020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109225780601f106108f757610100808354040283529160200191610922565b820191906000526020600020905b81548152906001019060200180831161090557829003601f168201915b505050505081565b600260208190526000918252604090912080546001820154928201546003830154919392909160ff169060040185565b600080600061096886610646565b915061097385610646565b600083815260026020819052604090912084815560018101839055429181019190915560038101805460ff191690559091506004018480516109b9929160200190610aaa565b50600192505b50509392505050565b600061063f838360026109e0565b600061063f838360015b6000806109ec85610646565b6000818152600160205260409020909150858051610a0e929160200190610aaa565b50600081815260208190526040902042815560019081018054859260ff1990911690836002811115610a3c57fe5b021790555060008181526020819052604090206002018480516109bf929160200190610aaa565b60206040519081016040526000815290565b60a06040519081016040908152600080835260208301819052908201819052606082015260808101610aa5610a63565b905290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610aeb57805160ff1916838001178555610b18565b82800160010185558215610b18579182015b82811115610b18578251825591602001919060010190610afd565b50610b24929150610b28565b5090565b610b4291905b80821115610b245760008155600101610b2e565b905600a165627a7a7230582046926a9b79a42bf9951f1aa7912d4ac8e5658f79c7069a15520d9388d5acc0d50029`

// DeployDAOTracabilite deploys a new Ethereum contract, binding an instance of DAOTracabilite to it.
func DeployDAOTracabilite(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DAOTracabilite, error) {
	parsed, err := abi.JSON(strings.NewReader(DAOTracabiliteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DAOTracabiliteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DAOTracabilite{DAOTracabiliteCaller: DAOTracabiliteCaller{contract: contract}, DAOTracabiliteTransactor: DAOTracabiliteTransactor{contract: contract}, DAOTracabiliteFilterer: DAOTracabiliteFilterer{contract: contract}}, nil
}

// DAOTracabilite is an auto generated Go binding around an Ethereum contract.
type DAOTracabilite struct {
	DAOTracabiliteCaller     // Read-only binding to the contract
	DAOTracabiliteTransactor // Write-only binding to the contract
	DAOTracabiliteFilterer   // Log filterer for contract events
}

// DAOTracabiliteCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAOTracabiliteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTracabiliteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAOTracabiliteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTracabiliteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAOTracabiliteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTracabiliteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAOTracabiliteSession struct {
	Contract     *DAOTracabilite   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAOTracabiliteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAOTracabiliteCallerSession struct {
	Contract *DAOTracabiliteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DAOTracabiliteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAOTracabiliteTransactorSession struct {
	Contract     *DAOTracabiliteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DAOTracabiliteRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAOTracabiliteRaw struct {
	Contract *DAOTracabilite // Generic contract binding to access the raw methods on
}

// DAOTracabiliteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAOTracabiliteCallerRaw struct {
	Contract *DAOTracabiliteCaller // Generic read-only contract binding to access the raw methods on
}

// DAOTracabiliteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAOTracabiliteTransactorRaw struct {
	Contract *DAOTracabiliteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAOTracabilite creates a new instance of DAOTracabilite, bound to a specific deployed contract.
func NewDAOTracabilite(address common.Address, backend bind.ContractBackend) (*DAOTracabilite, error) {
	contract, err := bindDAOTracabilite(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAOTracabilite{DAOTracabiliteCaller: DAOTracabiliteCaller{contract: contract}, DAOTracabiliteTransactor: DAOTracabiliteTransactor{contract: contract}, DAOTracabiliteFilterer: DAOTracabiliteFilterer{contract: contract}}, nil
}

// NewDAOTracabiliteCaller creates a new read-only instance of DAOTracabilite, bound to a specific deployed contract.
func NewDAOTracabiliteCaller(address common.Address, caller bind.ContractCaller) (*DAOTracabiliteCaller, error) {
	contract, err := bindDAOTracabilite(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAOTracabiliteCaller{contract: contract}, nil
}

// NewDAOTracabiliteTransactor creates a new write-only instance of DAOTracabilite, bound to a specific deployed contract.
func NewDAOTracabiliteTransactor(address common.Address, transactor bind.ContractTransactor) (*DAOTracabiliteTransactor, error) {
	contract, err := bindDAOTracabilite(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAOTracabiliteTransactor{contract: contract}, nil
}

// NewDAOTracabiliteFilterer creates a new log filterer instance of DAOTracabilite, bound to a specific deployed contract.
func NewDAOTracabiliteFilterer(address common.Address, filterer bind.ContractFilterer) (*DAOTracabiliteFilterer, error) {
	contract, err := bindDAOTracabilite(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAOTracabiliteFilterer{contract: contract}, nil
}

// bindDAOTracabilite binds a generic wrapper to an already deployed contract.
func bindDAOTracabilite(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DAOTracabiliteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOTracabilite *DAOTracabiliteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DAOTracabilite.Contract.DAOTracabiliteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOTracabilite *DAOTracabiliteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.DAOTracabiliteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOTracabilite *DAOTracabiliteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.DAOTracabiliteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAOTracabilite *DAOTracabiliteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DAOTracabilite.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAOTracabilite *DAOTracabiliteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAOTracabilite *DAOTracabiliteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.contract.Transact(opts, method, params...)
}

// ContainerList is a free data retrieval call binding the contract method 0xbf5af742.
//
// Solidity: function ContainerList( bytes32) constant returns(From bytes32, To bytes32, Date uint256, Kind uint8, Metadata string)
func (_DAOTracabilite *DAOTracabiliteCaller) ContainerList(opts *bind.CallOpts, arg0 [32]byte) (struct {
	From     [32]byte
	To       [32]byte
	Date     *big.Int
	Kind     uint8
	Metadata string
}, error) {
	ret := new(struct {
		From     [32]byte
		To       [32]byte
		Date     *big.Int
		Kind     uint8
		Metadata string
	})
	out := ret
	err := _DAOTracabilite.contract.Call(opts, out, "ContainerList", arg0)
	return *ret, err
}

// ContainerList is a free data retrieval call binding the contract method 0xbf5af742.
//
// Solidity: function ContainerList( bytes32) constant returns(From bytes32, To bytes32, Date uint256, Kind uint8, Metadata string)
func (_DAOTracabilite *DAOTracabiliteSession) ContainerList(arg0 [32]byte) (struct {
	From     [32]byte
	To       [32]byte
	Date     *big.Int
	Kind     uint8
	Metadata string
}, error) {
	return _DAOTracabilite.Contract.ContainerList(&_DAOTracabilite.CallOpts, arg0)
}

// ContainerList is a free data retrieval call binding the contract method 0xbf5af742.
//
// Solidity: function ContainerList( bytes32) constant returns(From bytes32, To bytes32, Date uint256, Kind uint8, Metadata string)
func (_DAOTracabilite *DAOTracabiliteCallerSession) ContainerList(arg0 [32]byte) (struct {
	From     [32]byte
	To       [32]byte
	Date     *big.Int
	Kind     uint8
	Metadata string
}, error) {
	return _DAOTracabilite.Contract.ContainerList(&_DAOTracabilite.CallOpts, arg0)
}

// GetHistory is a free data retrieval call binding the contract method 0x82d29347.
//
// Solidity: function GetHistory(guid string) constant returns(bool)
func (_DAOTracabilite *DAOTracabiliteCaller) GetHistory(opts *bind.CallOpts, guid string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DAOTracabilite.contract.Call(opts, out, "GetHistory", guid)
	return *ret0, err
}

// GetHistory is a free data retrieval call binding the contract method 0x82d29347.
//
// Solidity: function GetHistory(guid string) constant returns(bool)
func (_DAOTracabilite *DAOTracabiliteSession) GetHistory(guid string) (bool, error) {
	return _DAOTracabilite.Contract.GetHistory(&_DAOTracabilite.CallOpts, guid)
}

// GetHistory is a free data retrieval call binding the contract method 0x82d29347.
//
// Solidity: function GetHistory(guid string) constant returns(bool)
func (_DAOTracabilite *DAOTracabiliteCallerSession) GetHistory(guid string) (bool, error) {
	return _DAOTracabilite.Contract.GetHistory(&_DAOTracabilite.CallOpts, guid)
}

// Indexes is a free data retrieval call binding the contract method 0xbc863147.
//
// Solidity: function Indexes( bytes32) constant returns(string)
func (_DAOTracabilite *DAOTracabiliteCaller) Indexes(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DAOTracabilite.contract.Call(opts, out, "Indexes", arg0)
	return *ret0, err
}

// Indexes is a free data retrieval call binding the contract method 0xbc863147.
//
// Solidity: function Indexes( bytes32) constant returns(string)
func (_DAOTracabilite *DAOTracabiliteSession) Indexes(arg0 [32]byte) (string, error) {
	return _DAOTracabilite.Contract.Indexes(&_DAOTracabilite.CallOpts, arg0)
}

// Indexes is a free data retrieval call binding the contract method 0xbc863147.
//
// Solidity: function Indexes( bytes32) constant returns(string)
func (_DAOTracabilite *DAOTracabiliteCallerSession) Indexes(arg0 [32]byte) (string, error) {
	return _DAOTracabilite.Contract.Indexes(&_DAOTracabilite.CallOpts, arg0)
}

// Objects is a free data retrieval call binding the contract method 0x0b181f1c.
//
// Solidity: function Objects( bytes32) constant returns(Date uint256, Kind uint8, Metadata string)
func (_DAOTracabilite *DAOTracabiliteCaller) Objects(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Date     *big.Int
	Kind     uint8
	Metadata string
}, error) {
	ret := new(struct {
		Date     *big.Int
		Kind     uint8
		Metadata string
	})
	out := ret
	err := _DAOTracabilite.contract.Call(opts, out, "Objects", arg0)
	return *ret, err
}

// Objects is a free data retrieval call binding the contract method 0x0b181f1c.
//
// Solidity: function Objects( bytes32) constant returns(Date uint256, Kind uint8, Metadata string)
func (_DAOTracabilite *DAOTracabiliteSession) Objects(arg0 [32]byte) (struct {
	Date     *big.Int
	Kind     uint8
	Metadata string
}, error) {
	return _DAOTracabilite.Contract.Objects(&_DAOTracabilite.CallOpts, arg0)
}

// Objects is a free data retrieval call binding the contract method 0x0b181f1c.
//
// Solidity: function Objects( bytes32) constant returns(Date uint256, Kind uint8, Metadata string)
func (_DAOTracabilite *DAOTracabiliteCallerSession) Objects(arg0 [32]byte) (struct {
	Date     *big.Int
	Kind     uint8
	Metadata string
}, error) {
	return _DAOTracabilite.Contract.Objects(&_DAOTracabilite.CallOpts, arg0)
}

// ToInternalId is a free data retrieval call binding the contract method 0x6d325d8c.
//
// Solidity: function toInternalId(guid string) constant returns(bytes32)
func (_DAOTracabilite *DAOTracabiliteCaller) ToInternalId(opts *bind.CallOpts, guid string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DAOTracabilite.contract.Call(opts, out, "toInternalId", guid)
	return *ret0, err
}

// ToInternalId is a free data retrieval call binding the contract method 0x6d325d8c.
//
// Solidity: function toInternalId(guid string) constant returns(bytes32)
func (_DAOTracabilite *DAOTracabiliteSession) ToInternalId(guid string) ([32]byte, error) {
	return _DAOTracabilite.Contract.ToInternalId(&_DAOTracabilite.CallOpts, guid)
}

// ToInternalId is a free data retrieval call binding the contract method 0x6d325d8c.
//
// Solidity: function toInternalId(guid string) constant returns(bytes32)
func (_DAOTracabilite *DAOTracabiliteCallerSession) ToInternalId(guid string) ([32]byte, error) {
	return _DAOTracabilite.Contract.ToInternalId(&_DAOTracabilite.CallOpts, guid)
}

// AddContainer is a paid mutator transaction binding the contract method 0xd09cf216.
//
// Solidity: function AddContainer(guid string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteTransactor) AddContainer(opts *bind.TransactOpts, guid string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.contract.Transact(opts, "AddContainer", guid, metadata)
}

// AddContainer is a paid mutator transaction binding the contract method 0xd09cf216.
//
// Solidity: function AddContainer(guid string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteSession) AddContainer(guid string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.AddContainer(&_DAOTracabilite.TransactOpts, guid, metadata)
}

// AddContainer is a paid mutator transaction binding the contract method 0xd09cf216.
//
// Solidity: function AddContainer(guid string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteTransactorSession) AddContainer(guid string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.AddContainer(&_DAOTracabilite.TransactOpts, guid, metadata)
}

// AddObjectToContainer is a paid mutator transaction binding the contract method 0xcc6d7b2c.
//
// Solidity: function AddObjectToContainer(guidFrom string, guidTo string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteTransactor) AddObjectToContainer(opts *bind.TransactOpts, guidFrom string, guidTo string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.contract.Transact(opts, "AddObjectToContainer", guidFrom, guidTo, metadata)
}

// AddObjectToContainer is a paid mutator transaction binding the contract method 0xcc6d7b2c.
//
// Solidity: function AddObjectToContainer(guidFrom string, guidTo string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteSession) AddObjectToContainer(guidFrom string, guidTo string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.AddObjectToContainer(&_DAOTracabilite.TransactOpts, guidFrom, guidTo, metadata)
}

// AddObjectToContainer is a paid mutator transaction binding the contract method 0xcc6d7b2c.
//
// Solidity: function AddObjectToContainer(guidFrom string, guidTo string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteTransactorSession) AddObjectToContainer(guidFrom string, guidTo string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.AddObjectToContainer(&_DAOTracabilite.TransactOpts, guidFrom, guidTo, metadata)
}

// AddQRCode is a paid mutator transaction binding the contract method 0x36ecf908.
//
// Solidity: function AddQRCode(guid string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteTransactor) AddQRCode(opts *bind.TransactOpts, guid string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.contract.Transact(opts, "AddQRCode", guid, metadata)
}

// AddQRCode is a paid mutator transaction binding the contract method 0x36ecf908.
//
// Solidity: function AddQRCode(guid string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteSession) AddQRCode(guid string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.AddQRCode(&_DAOTracabilite.TransactOpts, guid, metadata)
}

// AddQRCode is a paid mutator transaction binding the contract method 0x36ecf908.
//
// Solidity: function AddQRCode(guid string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteTransactorSession) AddQRCode(guid string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.AddQRCode(&_DAOTracabilite.TransactOpts, guid, metadata)
}

// AddUnit is a paid mutator transaction binding the contract method 0xdedde18f.
//
// Solidity: function AddUnit(guid string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteTransactor) AddUnit(opts *bind.TransactOpts, guid string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.contract.Transact(opts, "AddUnit", guid, metadata)
}

// AddUnit is a paid mutator transaction binding the contract method 0xdedde18f.
//
// Solidity: function AddUnit(guid string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteSession) AddUnit(guid string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.AddUnit(&_DAOTracabilite.TransactOpts, guid, metadata)
}

// AddUnit is a paid mutator transaction binding the contract method 0xdedde18f.
//
// Solidity: function AddUnit(guid string, metadata string) returns(bool)
func (_DAOTracabilite *DAOTracabiliteTransactorSession) AddUnit(guid string, metadata string) (*types.Transaction, error) {
	return _DAOTracabilite.Contract.AddUnit(&_DAOTracabilite.TransactOpts, guid, metadata)
}
