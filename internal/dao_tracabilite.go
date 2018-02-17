// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internal

import (
	"math/big"
	"strings"
)

// DAOTracabiliteABI is the input ABI used to generate the binding from.
const DAOTracabiliteABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Objects\",\"outputs\":[{\"name\":\"Date\",\"type\":\"uint256\"},{\"name\":\"Kind\",\"type\":\"uint8\"},{\"name\":\"Metadata\",\"type\":\"string\"},{\"name\":\"TraceCount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"AddQRCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"}],\"name\":\"toInternalId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Traces\",\"outputs\":[{\"name\":\"From\",\"type\":\"bytes32\"},{\"name\":\"To\",\"type\":\"bytes32\"},{\"name\":\"Date\",\"type\":\"uint256\"},{\"name\":\"Kind\",\"type\":\"uint8\"},{\"name\":\"Metadata\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Indexes\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ContainerList\",\"outputs\":[{\"name\":\"From\",\"type\":\"bytes32\"},{\"name\":\"To\",\"type\":\"bytes32\"},{\"name\":\"Date\",\"type\":\"uint256\"},{\"name\":\"Kind\",\"type\":\"uint8\"},{\"name\":\"Metadata\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"GetHistoryHash\",\"outputs\":[{\"name\":\"traceHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guidFrom\",\"type\":\"string\"},{\"name\":\"guidTo\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"AddObjectToContainer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"AddContainer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guid\",\"type\":\"string\"},{\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"AddUnit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"ObjectNotification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"ActionNotification\",\"type\":\"event\"}]"

// DAOTracabiliteBin is the compiled bytecode used for deploying new contracts.
const DAOTracabiliteBin = `0x6060604052341561000f57600080fd5b610ee38061001e6000396000f3006060604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630b181f1c81146100a857806336ecf9081461016f5780636d325d8c1461021657806382ef93ea14610279578063bc86314714610348578063bf5af742146103d5578063c96f48aa146103eb578063cc6d7b2c1461043e578063d09cf21614610513578063dedde18f146105a6575b600080fd5b34156100b357600080fd5b6100be600435610639565b604051808581526020018460028111156100d457fe5b60ff16815260408101839052606082820381016020830190815285546002600182161561010002600019019091160491830182905291608001908590801561015d5780601f106101325761010080835404028352916020019161015d565b820191906000526020600020905b81548152906001019060200180831161014057829003601f168201915b50509550505050505060405180910390f35b341561017a57600080fd5b61020260046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061066495505050505050565b604051901515815260200160405180910390f35b341561022157600080fd5b61026760046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061067995505050505050565b60405190815260200160405180910390f35b341561028457600080fd5b61028f6004356106de565b6040518581526020810185905260408101849052606081018360038111156102b357fe5b60ff16815260408282038101602083019081528454600260018216156101000260001901909116049183018290529160600190849080156103355780601f1061030a57610100808354040283529160200191610335565b820191906000526020600020905b81548152906001019060200180831161031857829003601f168201915b5050965050505050505060405180910390f35b341561035357600080fd5b61035e60043561070f565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561039a578082015183820152602001610382565b50505050905090810190601f1680156103c75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103e057600080fd5b61028f6004356107bf565b34156103f657600080fd5b61026760046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965050933593506107ef92505050565b341561044957600080fd5b61020260046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f0160208091040260200160405190810160405281815292919060208401838380828437509496506109d495505050505050565b341561051e57600080fd5b61020260046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650610a7995505050505050565b34156105b157600080fd5b61020260046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f016020809104026020016040519081016040528181529291906020840183838082843750949650610a8795505050505050565b6000602081905290815260409020805460018201546003830154919260ff9091169160029091019084565b600061067283836000610a91565b9392505050565b6000816040518082805190602001908083835b602083106106ab5780518252601f19909201916020918201910161068c565b6001836020036101000a038019825116818451161790925250505091909101925060409150505180910390209050919050565b60036020819052600091825260409091208054600182015460028301549383015491939092909160ff169060040185565b60016020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107b75780601f1061078c576101008083540402835291602001916107b7565b820191906000526020600020905b81548152906001019060200180831161079a57829003601f168201915b505050505081565b600260208190526000918252604090912080546001820154928201546003830154919392909160ff169060040185565b6000806107fa610dd5565b600080610805610dd5565b61080d610dd5565b600060649650866040518059106108215750595b818152601f19601f830116810160200160405290509550600094505b881561088857600a89069350600a890498508360300160f860020a0286868060010197508151811061086b57fe5b906020010190600160f860020a031916908160001a90535061083d565b8992508483510160405180591061089c5750595b818152601f19601f830116810160200160405290509150600090505b825181101561090b578281815181106108cd57fe5b016020015160f860020a900460f860020a028282815181106108eb57fe5b906020010190600160f860020a031916908160001a9053506001016108b8565b5060005b8481101561096957858160018703038151811061092857fe5b016020015160f860020a900460f860020a0282845183018151811061094957fe5b906020010190600160f860020a031916908160001a90535060010161090f565b816040518082805190602001908083835b602083106109995780518252601f19909201916020918201910161097a565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051809103902097505050505050505092915050565b60008060006109e286610679565b91506109ed85610679565b60008381526002602081905260409091208481556001808201849055429282019290925560038101805460ff1916909217909155909150600401848051610a38929160200190610de7565b50610a47868383600288610be3565b610a55858383600388610be3565b50600090815260208190526040902060030180546001908101909155949350505050565b600061067283836002610a91565b6000610672838360015b600080610a9d85610679565b6000818152600160205260409020909150858051610abf929160200190610de7565b50600081815260208190526040902042815560019081018054859260ff1990911690836002811115610aed57fe5b02179055506000818152602081905260409020600201848051610b14929160200190610de7565b50600081815260208190526040808220600301919091557f147a59d0ce1e1ef8e354b9fcbf2a4f901fd6e7f2da431ec98e482e1b11c41cb490849086905180836002811115610b5f57fe5b60ff16815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610ba0578082015183820152602001610b88565b50505050905090810190601f168015610bcd5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a1509392505050565b610beb610e65565b84815260208101849052426040820152600060608201846003811115610c0d57fe5b90816003811115610c1a57fe5b90525060808201839052600086815260208190526040902060030154610c419088906107ef565b6000818152600360205260409020909150829081518155602082015160018201556040820151816002015560608201518160030160006101000a81548160ff02191690836003811115610c9057fe5b0217905550608082015181600401908051610caf929160200190610de7565b509050507f5b15b5f04dd4d8e09768a9c816b465b956cd4be67e9a6ef5e827b4343887819e84888560405180846003811115610ce757fe5b60ff1681526020018060200180602001838103835285818151815260200191508051906020019080838360005b83811015610d2c578082015183820152602001610d14565b50505050905090810190601f168015610d595780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b83811015610d8f578082015183820152602001610d77565b50505050905090810190601f168015610dbc5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a150505050505050565b60206040519081016040526000815290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e2857805160ff1916838001178555610e55565b82800160010185558215610e55579182015b82811115610e55578251825591602001919060010190610e3a565b50610e61929150610e9a565b5090565b60a06040519081016040908152600080835260208301819052908201819052606082015260808101610e95610dd5565b905290565b610eb491905b80821115610e615760008155600101610ea0565b905600a165627a7a72305820f937539bd26ffd5366727cc762c1466f11431d7997353012538f7754f18301210029`

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

// GetHistoryHash is a free data retrieval call binding the contract method 0xc96f48aa.
//
// Solidity: function GetHistoryHash(guid string, index uint256) constant returns(traceHash bytes32)
func (_DAOTracabilite *DAOTracabiliteCaller) GetHistoryHash(opts *bind.CallOpts, guid string, index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DAOTracabilite.contract.Call(opts, out, "GetHistoryHash", guid, index)
	return *ret0, err
}

// GetHistoryHash is a free data retrieval call binding the contract method 0xc96f48aa.
//
// Solidity: function GetHistoryHash(guid string, index uint256) constant returns(traceHash bytes32)
func (_DAOTracabilite *DAOTracabiliteSession) GetHistoryHash(guid string, index *big.Int) ([32]byte, error) {
	return _DAOTracabilite.Contract.GetHistoryHash(&_DAOTracabilite.CallOpts, guid, index)
}

// GetHistoryHash is a free data retrieval call binding the contract method 0xc96f48aa.
//
// Solidity: function GetHistoryHash(guid string, index uint256) constant returns(traceHash bytes32)
func (_DAOTracabilite *DAOTracabiliteCallerSession) GetHistoryHash(guid string, index *big.Int) ([32]byte, error) {
	return _DAOTracabilite.Contract.GetHistoryHash(&_DAOTracabilite.CallOpts, guid, index)
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
// Solidity: function Objects( bytes32) constant returns(Date uint256, Kind uint8, Metadata string, TraceCount uint256)
func (_DAOTracabilite *DAOTracabiliteCaller) Objects(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Date       *big.Int
	Kind       uint8
	Metadata   string
	TraceCount *big.Int
}, error) {
	ret := new(struct {
		Date       *big.Int
		Kind       uint8
		Metadata   string
		TraceCount *big.Int
	})
	out := ret
	err := _DAOTracabilite.contract.Call(opts, out, "Objects", arg0)
	return *ret, err
}

// Objects is a free data retrieval call binding the contract method 0x0b181f1c.
//
// Solidity: function Objects( bytes32) constant returns(Date uint256, Kind uint8, Metadata string, TraceCount uint256)
func (_DAOTracabilite *DAOTracabiliteSession) Objects(arg0 [32]byte) (struct {
	Date       *big.Int
	Kind       uint8
	Metadata   string
	TraceCount *big.Int
}, error) {
	return _DAOTracabilite.Contract.Objects(&_DAOTracabilite.CallOpts, arg0)
}

// Objects is a free data retrieval call binding the contract method 0x0b181f1c.
//
// Solidity: function Objects( bytes32) constant returns(Date uint256, Kind uint8, Metadata string, TraceCount uint256)
func (_DAOTracabilite *DAOTracabiliteCallerSession) Objects(arg0 [32]byte) (struct {
	Date       *big.Int
	Kind       uint8
	Metadata   string
	TraceCount *big.Int
}, error) {
	return _DAOTracabilite.Contract.Objects(&_DAOTracabilite.CallOpts, arg0)
}

// Traces is a free data retrieval call binding the contract method 0x82ef93ea.
//
// Solidity: function Traces( bytes32) constant returns(From bytes32, To bytes32, Date uint256, Kind uint8, Metadata string)
func (_DAOTracabilite *DAOTracabiliteCaller) Traces(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _DAOTracabilite.contract.Call(opts, out, "Traces", arg0)
	return *ret, err
}

// Traces is a free data retrieval call binding the contract method 0x82ef93ea.
//
// Solidity: function Traces( bytes32) constant returns(From bytes32, To bytes32, Date uint256, Kind uint8, Metadata string)
func (_DAOTracabilite *DAOTracabiliteSession) Traces(arg0 [32]byte) (struct {
	From     [32]byte
	To       [32]byte
	Date     *big.Int
	Kind     uint8
	Metadata string
}, error) {
	return _DAOTracabilite.Contract.Traces(&_DAOTracabilite.CallOpts, arg0)
}

// Traces is a free data retrieval call binding the contract method 0x82ef93ea.
//
// Solidity: function Traces( bytes32) constant returns(From bytes32, To bytes32, Date uint256, Kind uint8, Metadata string)
func (_DAOTracabilite *DAOTracabiliteCallerSession) Traces(arg0 [32]byte) (struct {
	From     [32]byte
	To       [32]byte
	Date     *big.Int
	Kind     uint8
	Metadata string
}, error) {
	return _DAOTracabilite.Contract.Traces(&_DAOTracabilite.CallOpts, arg0)
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

// DAOTracabiliteActionNotificationIterator is returned from FilterActionNotification and is used to iterate over the raw logs and unpacked data for ActionNotification events raised by the DAOTracabilite contract.
type DAOTracabiliteActionNotificationIterator struct {
	Event *DAOTracabiliteActionNotification // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAOTracabiliteActionNotificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTracabiliteActionNotification)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAOTracabiliteActionNotification)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAOTracabiliteActionNotificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTracabiliteActionNotificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTracabiliteActionNotification represents a ActionNotification event raised by the DAOTracabilite contract.
type DAOTracabiliteActionNotification struct {
	Action   uint8
	From     string
	Metadata string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterActionNotification is a free log retrieval operation binding the contract event 0x5b15b5f04dd4d8e09768a9c816b465b956cd4be67e9a6ef5e827b4343887819e.
//
// Solidity: event ActionNotification(action uint8, from string, metadata string)
func (_DAOTracabilite *DAOTracabiliteFilterer) FilterActionNotification(opts *bind.FilterOpts) (*DAOTracabiliteActionNotificationIterator, error) {

	logs, sub, err := _DAOTracabilite.contract.FilterLogs(opts, "ActionNotification")
	if err != nil {
		return nil, err
	}
	return &DAOTracabiliteActionNotificationIterator{contract: _DAOTracabilite.contract, event: "ActionNotification", logs: logs, sub: sub}, nil
}

// WatchActionNotification is a free log subscription operation binding the contract event 0x5b15b5f04dd4d8e09768a9c816b465b956cd4be67e9a6ef5e827b4343887819e.
//
// Solidity: event ActionNotification(action uint8, from string, metadata string)
func (_DAOTracabilite *DAOTracabiliteFilterer) WatchActionNotification(opts *bind.WatchOpts, sink chan<- *DAOTracabiliteActionNotification) (event.Subscription, error) {

	logs, sub, err := _DAOTracabilite.contract.WatchLogs(opts, "ActionNotification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTracabiliteActionNotification)
				if err := _DAOTracabilite.contract.UnpackLog(event, "ActionNotification", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DAOTracabiliteObjectNotificationIterator is returned from FilterObjectNotification and is used to iterate over the raw logs and unpacked data for ObjectNotification events raised by the DAOTracabilite contract.
type DAOTracabiliteObjectNotificationIterator struct {
	Event *DAOTracabiliteObjectNotification // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DAOTracabiliteObjectNotificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOTracabiliteObjectNotification)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DAOTracabiliteObjectNotification)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DAOTracabiliteObjectNotificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOTracabiliteObjectNotificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOTracabiliteObjectNotification represents a ObjectNotification event raised by the DAOTracabilite contract.
type DAOTracabiliteObjectNotification struct {
	Kind     uint8
	Metadata string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterObjectNotification is a free log retrieval operation binding the contract event 0x147a59d0ce1e1ef8e354b9fcbf2a4f901fd6e7f2da431ec98e482e1b11c41cb4.
//
// Solidity: event ObjectNotification(kind uint8, metadata string)
func (_DAOTracabilite *DAOTracabiliteFilterer) FilterObjectNotification(opts *bind.FilterOpts) (*DAOTracabiliteObjectNotificationIterator, error) {

	logs, sub, err := _DAOTracabilite.contract.FilterLogs(opts, "ObjectNotification")
	if err != nil {
		return nil, err
	}
	return &DAOTracabiliteObjectNotificationIterator{contract: _DAOTracabilite.contract, event: "ObjectNotification", logs: logs, sub: sub}, nil
}

// WatchObjectNotification is a free log subscription operation binding the contract event 0x147a59d0ce1e1ef8e354b9fcbf2a4f901fd6e7f2da431ec98e482e1b11c41cb4.
//
// Solidity: event ObjectNotification(kind uint8, metadata string)
func (_DAOTracabilite *DAOTracabiliteFilterer) WatchObjectNotification(opts *bind.WatchOpts, sink chan<- *DAOTracabiliteObjectNotification) (event.Subscription, error) {

	logs, sub, err := _DAOTracabilite.contract.WatchLogs(opts, "ObjectNotification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOTracabiliteObjectNotification)
				if err := _DAOTracabilite.contract.UnpackLog(event, "ObjectNotification", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
