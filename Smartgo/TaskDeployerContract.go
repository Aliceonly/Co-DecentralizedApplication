// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TaskDeployerContract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TaskDeployerContractABI is the input ABI used to generate the binding from.
const TaskDeployerContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Cancel_Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Confirm_Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskcategory\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Creat_Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"phonenumber\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"studentid\",\"type\":\"uint256\"}],\"name\":\"Creatuser_Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"claim_Event\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_tasklist\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"Taskname\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"state\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"launchTime\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_userlist\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phonenumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"studentid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"cancelEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_studentid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"}],\"name\":\"changename\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_studentid\",\"type\":\"uint256\"}],\"name\":\"changepassword\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sigs\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_PHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_taskname\",\"type\":\"string\"}],\"name\":\"claimTrust\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"confirmtask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_launchTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_category\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_taskname\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createNewEvent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phonenumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"studentid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"name\":\"createuser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalanceOfContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_taskname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_timestamp\",\"type\":\"string\"}],\"name\":\"getTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"query\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"svalue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transderToContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// TaskDeployerContract is an auto generated Go binding around an Ethereum contract.
type TaskDeployerContract struct {
	TaskDeployerContractCaller     // Read-only binding to the contract
	TaskDeployerContractTransactor // Write-only binding to the contract
	TaskDeployerContractFilterer   // Log filterer for contract events
}

// TaskDeployerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskDeployerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskDeployerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskDeployerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskDeployerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskDeployerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskDeployerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskDeployerContractSession struct {
	Contract     *TaskDeployerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TaskDeployerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskDeployerContractCallerSession struct {
	Contract *TaskDeployerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TaskDeployerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskDeployerContractTransactorSession struct {
	Contract     *TaskDeployerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TaskDeployerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskDeployerContractRaw struct {
	Contract *TaskDeployerContract // Generic contract binding to access the raw methods on
}

// TaskDeployerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskDeployerContractCallerRaw struct {
	Contract *TaskDeployerContractCaller // Generic read-only contract binding to access the raw methods on
}

// TaskDeployerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskDeployerContractTransactorRaw struct {
	Contract *TaskDeployerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskDeployerContract creates a new instance of TaskDeployerContract, bound to a specific deployed contract.
func NewTaskDeployerContract(address common.Address, backend bind.ContractBackend) (*TaskDeployerContract, error) {
	contract, err := bindTaskDeployerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContract{TaskDeployerContractCaller: TaskDeployerContractCaller{contract: contract}, TaskDeployerContractTransactor: TaskDeployerContractTransactor{contract: contract}, TaskDeployerContractFilterer: TaskDeployerContractFilterer{contract: contract}}, nil
}

// NewTaskDeployerContractCaller creates a new read-only instance of TaskDeployerContract, bound to a specific deployed contract.
func NewTaskDeployerContractCaller(address common.Address, caller bind.ContractCaller) (*TaskDeployerContractCaller, error) {
	contract, err := bindTaskDeployerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractCaller{contract: contract}, nil
}

// NewTaskDeployerContractTransactor creates a new write-only instance of TaskDeployerContract, bound to a specific deployed contract.
func NewTaskDeployerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskDeployerContractTransactor, error) {
	contract, err := bindTaskDeployerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractTransactor{contract: contract}, nil
}

// NewTaskDeployerContractFilterer creates a new log filterer instance of TaskDeployerContract, bound to a specific deployed contract.
func NewTaskDeployerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskDeployerContractFilterer, error) {
	contract, err := bindTaskDeployerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractFilterer{contract: contract}, nil
}

// bindTaskDeployerContract binds a generic wrapper to an already deployed contract.
func bindTaskDeployerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskDeployerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskDeployerContract *TaskDeployerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskDeployerContract.Contract.TaskDeployerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskDeployerContract *TaskDeployerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.TaskDeployerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskDeployerContract *TaskDeployerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.TaskDeployerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskDeployerContract *TaskDeployerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskDeployerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskDeployerContract *TaskDeployerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskDeployerContract *TaskDeployerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.contract.Transact(opts, method, params...)
}

// Tasklist is a free data retrieval call binding the contract method 0x5ac43ca2.
//
// Solidity: function _tasklist(uint256 ) view returns(string Taskname, address sponsor, address beneficiary, string category, uint256 amount, uint256 timestamp, string state, string launchTime)
func (_TaskDeployerContract *TaskDeployerContractCaller) Tasklist(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Taskname    string
	Sponsor     common.Address
	Beneficiary common.Address
	Category    string
	Amount      *big.Int
	Timestamp   *big.Int
	State       string
	LaunchTime  string
}, error) {
	var out []interface{}
	err := _TaskDeployerContract.contract.Call(opts, &out, "_tasklist", arg0)

	outstruct := new(struct {
		Taskname    string
		Sponsor     common.Address
		Beneficiary common.Address
		Category    string
		Amount      *big.Int
		Timestamp   *big.Int
		State       string
		LaunchTime  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Taskname = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Sponsor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Beneficiary = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Category = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.LaunchTime = *abi.ConvertType(out[7], new(string)).(*string)

	return *outstruct, err

}

// Tasklist is a free data retrieval call binding the contract method 0x5ac43ca2.
//
// Solidity: function _tasklist(uint256 ) view returns(string Taskname, address sponsor, address beneficiary, string category, uint256 amount, uint256 timestamp, string state, string launchTime)
func (_TaskDeployerContract *TaskDeployerContractSession) Tasklist(arg0 *big.Int) (struct {
	Taskname    string
	Sponsor     common.Address
	Beneficiary common.Address
	Category    string
	Amount      *big.Int
	Timestamp   *big.Int
	State       string
	LaunchTime  string
}, error) {
	return _TaskDeployerContract.Contract.Tasklist(&_TaskDeployerContract.CallOpts, arg0)
}

// Tasklist is a free data retrieval call binding the contract method 0x5ac43ca2.
//
// Solidity: function _tasklist(uint256 ) view returns(string Taskname, address sponsor, address beneficiary, string category, uint256 amount, uint256 timestamp, string state, string launchTime)
func (_TaskDeployerContract *TaskDeployerContractCallerSession) Tasklist(arg0 *big.Int) (struct {
	Taskname    string
	Sponsor     common.Address
	Beneficiary common.Address
	Category    string
	Amount      *big.Int
	Timestamp   *big.Int
	State       string
	LaunchTime  string
}, error) {
	return _TaskDeployerContract.Contract.Tasklist(&_TaskDeployerContract.CallOpts, arg0)
}

// Userlist is a free data retrieval call binding the contract method 0x885b8290.
//
// Solidity: function _userlist(uint256 ) view returns(string name, string phonenumber, uint256 studentid, string password)
func (_TaskDeployerContract *TaskDeployerContractCaller) Userlist(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Password    string
}, error) {
	var out []interface{}
	err := _TaskDeployerContract.contract.Call(opts, &out, "_userlist", arg0)

	outstruct := new(struct {
		Name        string
		Phonenumber string
		Studentid   *big.Int
		Password    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Phonenumber = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Studentid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Password = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Userlist is a free data retrieval call binding the contract method 0x885b8290.
//
// Solidity: function _userlist(uint256 ) view returns(string name, string phonenumber, uint256 studentid, string password)
func (_TaskDeployerContract *TaskDeployerContractSession) Userlist(arg0 *big.Int) (struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Password    string
}, error) {
	return _TaskDeployerContract.Contract.Userlist(&_TaskDeployerContract.CallOpts, arg0)
}

// Userlist is a free data retrieval call binding the contract method 0x885b8290.
//
// Solidity: function _userlist(uint256 ) view returns(string name, string phonenumber, uint256 studentid, string password)
func (_TaskDeployerContract *TaskDeployerContractCallerSession) Userlist(arg0 *big.Int) (struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Password    string
}, error) {
	return _TaskDeployerContract.Contract.Userlist(&_TaskDeployerContract.CallOpts, arg0)
}

// GetBalanceOfContract is a free data retrieval call binding the contract method 0x22968885.
//
// Solidity: function getBalanceOfContract() view returns(uint256)
func (_TaskDeployerContract *TaskDeployerContractCaller) GetBalanceOfContract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaskDeployerContract.contract.Call(opts, &out, "getBalanceOfContract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOfContract is a free data retrieval call binding the contract method 0x22968885.
//
// Solidity: function getBalanceOfContract() view returns(uint256)
func (_TaskDeployerContract *TaskDeployerContractSession) GetBalanceOfContract() (*big.Int, error) {
	return _TaskDeployerContract.Contract.GetBalanceOfContract(&_TaskDeployerContract.CallOpts)
}

// GetBalanceOfContract is a free data retrieval call binding the contract method 0x22968885.
//
// Solidity: function getBalanceOfContract() view returns(uint256)
func (_TaskDeployerContract *TaskDeployerContractCallerSession) GetBalanceOfContract() (*big.Int, error) {
	return _TaskDeployerContract.Contract.GetBalanceOfContract(&_TaskDeployerContract.CallOpts)
}

// GetTxHash is a free data retrieval call binding the contract method 0x2c40bf6c.
//
// Solidity: function getTxHash(string _taskname, string _timestamp) view returns(bytes32)
func (_TaskDeployerContract *TaskDeployerContractCaller) GetTxHash(opts *bind.CallOpts, _taskname string, _timestamp string) ([32]byte, error) {
	var out []interface{}
	err := _TaskDeployerContract.contract.Call(opts, &out, "getTxHash", _taskname, _timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTxHash is a free data retrieval call binding the contract method 0x2c40bf6c.
//
// Solidity: function getTxHash(string _taskname, string _timestamp) view returns(bytes32)
func (_TaskDeployerContract *TaskDeployerContractSession) GetTxHash(_taskname string, _timestamp string) ([32]byte, error) {
	return _TaskDeployerContract.Contract.GetTxHash(&_TaskDeployerContract.CallOpts, _taskname, _timestamp)
}

// GetTxHash is a free data retrieval call binding the contract method 0x2c40bf6c.
//
// Solidity: function getTxHash(string _taskname, string _timestamp) view returns(bytes32)
func (_TaskDeployerContract *TaskDeployerContractCallerSession) GetTxHash(_taskname string, _timestamp string) ([32]byte, error) {
	return _TaskDeployerContract.Contract.GetTxHash(&_TaskDeployerContract.CallOpts, _taskname, _timestamp)
}

// Query is a free data retrieval call binding the contract method 0x2c46b205.
//
// Solidity: function query() view returns(uint256)
func (_TaskDeployerContract *TaskDeployerContractCaller) Query(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaskDeployerContract.contract.Call(opts, &out, "query")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Query is a free data retrieval call binding the contract method 0x2c46b205.
//
// Solidity: function query() view returns(uint256)
func (_TaskDeployerContract *TaskDeployerContractSession) Query() (*big.Int, error) {
	return _TaskDeployerContract.Contract.Query(&_TaskDeployerContract.CallOpts)
}

// Query is a free data retrieval call binding the contract method 0x2c46b205.
//
// Solidity: function query() view returns(uint256)
func (_TaskDeployerContract *TaskDeployerContractCallerSession) Query() (*big.Int, error) {
	return _TaskDeployerContract.Contract.Query(&_TaskDeployerContract.CallOpts)
}

// CancelEvent is a paid mutator transaction binding the contract method 0x3f69babd.
//
// Solidity: function cancelEvent(uint256 _timestamp) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) CancelEvent(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.Transact(opts, "cancelEvent", _timestamp)
}

// CancelEvent is a paid mutator transaction binding the contract method 0x3f69babd.
//
// Solidity: function cancelEvent(uint256 _timestamp) returns()
func (_TaskDeployerContract *TaskDeployerContractSession) CancelEvent(_timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.CancelEvent(&_TaskDeployerContract.TransactOpts, _timestamp)
}

// CancelEvent is a paid mutator transaction binding the contract method 0x3f69babd.
//
// Solidity: function cancelEvent(uint256 _timestamp) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) CancelEvent(_timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.CancelEvent(&_TaskDeployerContract.TransactOpts, _timestamp)
}

// Changename is a paid mutator transaction binding the contract method 0xab3c8449.
//
// Solidity: function changename(string _name, uint256 _studentid, string _password) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) Changename(opts *bind.TransactOpts, _name string, _studentid *big.Int, _password string) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.Transact(opts, "changename", _name, _studentid, _password)
}

// Changename is a paid mutator transaction binding the contract method 0xab3c8449.
//
// Solidity: function changename(string _name, uint256 _studentid, string _password) returns()
func (_TaskDeployerContract *TaskDeployerContractSession) Changename(_name string, _studentid *big.Int, _password string) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Changename(&_TaskDeployerContract.TransactOpts, _name, _studentid, _password)
}

// Changename is a paid mutator transaction binding the contract method 0xab3c8449.
//
// Solidity: function changename(string _name, uint256 _studentid, string _password) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) Changename(_name string, _studentid *big.Int, _password string) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Changename(&_TaskDeployerContract.TransactOpts, _name, _studentid, _password)
}

// Changepassword is a paid mutator transaction binding the contract method 0xf6f18e5d.
//
// Solidity: function changepassword(string _password, uint256 _studentid) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) Changepassword(opts *bind.TransactOpts, _password string, _studentid *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.Transact(opts, "changepassword", _password, _studentid)
}

// Changepassword is a paid mutator transaction binding the contract method 0xf6f18e5d.
//
// Solidity: function changepassword(string _password, uint256 _studentid) returns()
func (_TaskDeployerContract *TaskDeployerContractSession) Changepassword(_password string, _studentid *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Changepassword(&_TaskDeployerContract.TransactOpts, _password, _studentid)
}

// Changepassword is a paid mutator transaction binding the contract method 0xf6f18e5d.
//
// Solidity: function changepassword(string _password, uint256 _studentid) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) Changepassword(_password string, _studentid *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Changepassword(&_TaskDeployerContract.TransactOpts, _password, _studentid)
}

// ClaimTrust is a paid mutator transaction binding the contract method 0xf6bf490d.
//
// Solidity: function claimTrust(uint256 _timestamp, bytes _sigs, bytes32 _PHash, string _taskname) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) ClaimTrust(opts *bind.TransactOpts, _timestamp *big.Int, _sigs []byte, _PHash [32]byte, _taskname string) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.Transact(opts, "claimTrust", _timestamp, _sigs, _PHash, _taskname)
}

// ClaimTrust is a paid mutator transaction binding the contract method 0xf6bf490d.
//
// Solidity: function claimTrust(uint256 _timestamp, bytes _sigs, bytes32 _PHash, string _taskname) returns()
func (_TaskDeployerContract *TaskDeployerContractSession) ClaimTrust(_timestamp *big.Int, _sigs []byte, _PHash [32]byte, _taskname string) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.ClaimTrust(&_TaskDeployerContract.TransactOpts, _timestamp, _sigs, _PHash, _taskname)
}

// ClaimTrust is a paid mutator transaction binding the contract method 0xf6bf490d.
//
// Solidity: function claimTrust(uint256 _timestamp, bytes _sigs, bytes32 _PHash, string _taskname) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) ClaimTrust(_timestamp *big.Int, _sigs []byte, _PHash [32]byte, _taskname string) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.ClaimTrust(&_TaskDeployerContract.TransactOpts, _timestamp, _sigs, _PHash, _taskname)
}

// Confirmtask is a paid mutator transaction binding the contract method 0x7b4623cc.
//
// Solidity: function confirmtask(uint256 _timestamp) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) Confirmtask(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.Transact(opts, "confirmtask", _timestamp)
}

// Confirmtask is a paid mutator transaction binding the contract method 0x7b4623cc.
//
// Solidity: function confirmtask(uint256 _timestamp) returns()
func (_TaskDeployerContract *TaskDeployerContractSession) Confirmtask(_timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Confirmtask(&_TaskDeployerContract.TransactOpts, _timestamp)
}

// Confirmtask is a paid mutator transaction binding the contract method 0x7b4623cc.
//
// Solidity: function confirmtask(uint256 _timestamp) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) Confirmtask(_timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Confirmtask(&_TaskDeployerContract.TransactOpts, _timestamp)
}

// CreateNewEvent is a paid mutator transaction binding the contract method 0x63e0a525.
//
// Solidity: function createNewEvent(string _launchTime, string _category, string _taskname, uint256 _amount) payable returns(uint256)
func (_TaskDeployerContract *TaskDeployerContractTransactor) CreateNewEvent(opts *bind.TransactOpts, _launchTime string, _category string, _taskname string, _amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.Transact(opts, "createNewEvent", _launchTime, _category, _taskname, _amount)
}

// CreateNewEvent is a paid mutator transaction binding the contract method 0x63e0a525.
//
// Solidity: function createNewEvent(string _launchTime, string _category, string _taskname, uint256 _amount) payable returns(uint256)
func (_TaskDeployerContract *TaskDeployerContractSession) CreateNewEvent(_launchTime string, _category string, _taskname string, _amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.CreateNewEvent(&_TaskDeployerContract.TransactOpts, _launchTime, _category, _taskname, _amount)
}

// CreateNewEvent is a paid mutator transaction binding the contract method 0x63e0a525.
//
// Solidity: function createNewEvent(string _launchTime, string _category, string _taskname, uint256 _amount) payable returns(uint256)
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) CreateNewEvent(_launchTime string, _category string, _taskname string, _amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.CreateNewEvent(&_TaskDeployerContract.TransactOpts, _launchTime, _category, _taskname, _amount)
}

// Createuser is a paid mutator transaction binding the contract method 0x3fb93885.
//
// Solidity: function createuser(string name, string phonenumber, uint256 studentid, string password) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) Createuser(opts *bind.TransactOpts, name string, phonenumber string, studentid *big.Int, password string) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.Transact(opts, "createuser", name, phonenumber, studentid, password)
}

// Createuser is a paid mutator transaction binding the contract method 0x3fb93885.
//
// Solidity: function createuser(string name, string phonenumber, uint256 studentid, string password) returns()
func (_TaskDeployerContract *TaskDeployerContractSession) Createuser(name string, phonenumber string, studentid *big.Int, password string) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Createuser(&_TaskDeployerContract.TransactOpts, name, phonenumber, studentid, password)
}

// Createuser is a paid mutator transaction binding the contract method 0x3fb93885.
//
// Solidity: function createuser(string name, string phonenumber, uint256 studentid, string password) returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) Createuser(name string, phonenumber string, studentid *big.Int, password string) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Createuser(&_TaskDeployerContract.TransactOpts, name, phonenumber, studentid, password)
}

// Svalue is a paid mutator transaction binding the contract method 0xfc6378a4.
//
// Solidity: function svalue(address addr, uint256 amount) payable returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) Svalue(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.Transact(opts, "svalue", addr, amount)
}

// Svalue is a paid mutator transaction binding the contract method 0xfc6378a4.
//
// Solidity: function svalue(address addr, uint256 amount) payable returns()
func (_TaskDeployerContract *TaskDeployerContractSession) Svalue(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Svalue(&_TaskDeployerContract.TransactOpts, addr, amount)
}

// Svalue is a paid mutator transaction binding the contract method 0xfc6378a4.
//
// Solidity: function svalue(address addr, uint256 amount) payable returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) Svalue(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Svalue(&_TaskDeployerContract.TransactOpts, addr, amount)
}

// TransderToContract is a paid mutator transaction binding the contract method 0x460968dd.
//
// Solidity: function transderToContract() payable returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) TransderToContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.Transact(opts, "transderToContract")
}

// TransderToContract is a paid mutator transaction binding the contract method 0x460968dd.
//
// Solidity: function transderToContract() payable returns()
func (_TaskDeployerContract *TaskDeployerContractSession) TransderToContract() (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.TransderToContract(&_TaskDeployerContract.TransactOpts)
}

// TransderToContract is a paid mutator transaction binding the contract method 0x460968dd.
//
// Solidity: function transderToContract() payable returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) TransderToContract() (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.TransderToContract(&_TaskDeployerContract.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TaskDeployerContract *TaskDeployerContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Fallback(&_TaskDeployerContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Fallback(&_TaskDeployerContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaskDeployerContract *TaskDeployerContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDeployerContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaskDeployerContract *TaskDeployerContractSession) Receive() (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Receive(&_TaskDeployerContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaskDeployerContract *TaskDeployerContractTransactorSession) Receive() (*types.Transaction, error) {
	return _TaskDeployerContract.Contract.Receive(&_TaskDeployerContract.TransactOpts)
}

// TaskDeployerContractCancelEventIterator is returned from FilterCancelEvent and is used to iterate over the raw logs and unpacked data for CancelEvent events raised by the TaskDeployerContract contract.
type TaskDeployerContractCancelEventIterator struct {
	Event *TaskDeployerContractCancelEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractCancelEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractCancelEvent)
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
		it.Event = new(TaskDeployerContractCancelEvent)
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
func (it *TaskDeployerContractCancelEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractCancelEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractCancelEvent represents a CancelEvent event raised by the TaskDeployerContract contract.
type TaskDeployerContractCancelEvent struct {
	Add       common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCancelEvent is a free log retrieval operation binding the contract event 0xebb62664a922b45fff0b6d8448a6f26157da006cd1887a3311f5a678b598dda4.
//
// Solidity: event Cancel_Event(address add, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) FilterCancelEvent(opts *bind.FilterOpts) (*TaskDeployerContractCancelEventIterator, error) {

	logs, sub, err := _TaskDeployerContract.contract.FilterLogs(opts, "Cancel_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractCancelEventIterator{contract: _TaskDeployerContract.contract, event: "Cancel_Event", logs: logs, sub: sub}, nil
}

// WatchCancelEvent is a free log subscription operation binding the contract event 0xebb62664a922b45fff0b6d8448a6f26157da006cd1887a3311f5a678b598dda4.
//
// Solidity: event Cancel_Event(address add, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) WatchCancelEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractCancelEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContract.contract.WatchLogs(opts, "Cancel_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractCancelEvent)
				if err := _TaskDeployerContract.contract.UnpackLog(event, "Cancel_Event", log); err != nil {
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

// ParseCancelEvent is a log parse operation binding the contract event 0xebb62664a922b45fff0b6d8448a6f26157da006cd1887a3311f5a678b598dda4.
//
// Solidity: event Cancel_Event(address add, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) ParseCancelEvent(log types.Log) (*TaskDeployerContractCancelEvent, error) {
	event := new(TaskDeployerContractCancelEvent)
	if err := _TaskDeployerContract.contract.UnpackLog(event, "Cancel_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskDeployerContractConfirmEventIterator is returned from FilterConfirmEvent and is used to iterate over the raw logs and unpacked data for ConfirmEvent events raised by the TaskDeployerContract contract.
type TaskDeployerContractConfirmEventIterator struct {
	Event *TaskDeployerContractConfirmEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractConfirmEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractConfirmEvent)
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
		it.Event = new(TaskDeployerContractConfirmEvent)
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
func (it *TaskDeployerContractConfirmEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractConfirmEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractConfirmEvent represents a ConfirmEvent event raised by the TaskDeployerContract contract.
type TaskDeployerContractConfirmEvent struct {
	Add       common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmEvent is a free log retrieval operation binding the contract event 0x6b9ce9a8215bda6518e4fb80be40bc0ce2da320e58757a4aef77aca3a1d2fb4e.
//
// Solidity: event Confirm_Event(address add, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) FilterConfirmEvent(opts *bind.FilterOpts) (*TaskDeployerContractConfirmEventIterator, error) {

	logs, sub, err := _TaskDeployerContract.contract.FilterLogs(opts, "Confirm_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractConfirmEventIterator{contract: _TaskDeployerContract.contract, event: "Confirm_Event", logs: logs, sub: sub}, nil
}

// WatchConfirmEvent is a free log subscription operation binding the contract event 0x6b9ce9a8215bda6518e4fb80be40bc0ce2da320e58757a4aef77aca3a1d2fb4e.
//
// Solidity: event Confirm_Event(address add, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) WatchConfirmEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractConfirmEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContract.contract.WatchLogs(opts, "Confirm_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractConfirmEvent)
				if err := _TaskDeployerContract.contract.UnpackLog(event, "Confirm_Event", log); err != nil {
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

// ParseConfirmEvent is a log parse operation binding the contract event 0x6b9ce9a8215bda6518e4fb80be40bc0ce2da320e58757a4aef77aca3a1d2fb4e.
//
// Solidity: event Confirm_Event(address add, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) ParseConfirmEvent(log types.Log) (*TaskDeployerContractConfirmEvent, error) {
	event := new(TaskDeployerContractConfirmEvent)
	if err := _TaskDeployerContract.contract.UnpackLog(event, "Confirm_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskDeployerContractCreatEventIterator is returned from FilterCreatEvent and is used to iterate over the raw logs and unpacked data for CreatEvent events raised by the TaskDeployerContract contract.
type TaskDeployerContractCreatEventIterator struct {
	Event *TaskDeployerContractCreatEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractCreatEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractCreatEvent)
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
		it.Event = new(TaskDeployerContractCreatEvent)
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
func (it *TaskDeployerContractCreatEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractCreatEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractCreatEvent represents a CreatEvent event raised by the TaskDeployerContract contract.
type TaskDeployerContractCreatEvent struct {
	Add          common.Address
	Taskname     string
	Taskcategory string
	Amount       *big.Int
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreatEvent is a free log retrieval operation binding the contract event 0x7de35515d174e93110b2a3f097c27ebf4cb042509117d9a38b69d5b47b1a2ba3.
//
// Solidity: event Creat_Event(address add, string taskname, string taskcategory, uint256 amount, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) FilterCreatEvent(opts *bind.FilterOpts) (*TaskDeployerContractCreatEventIterator, error) {

	logs, sub, err := _TaskDeployerContract.contract.FilterLogs(opts, "Creat_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractCreatEventIterator{contract: _TaskDeployerContract.contract, event: "Creat_Event", logs: logs, sub: sub}, nil
}

// WatchCreatEvent is a free log subscription operation binding the contract event 0x7de35515d174e93110b2a3f097c27ebf4cb042509117d9a38b69d5b47b1a2ba3.
//
// Solidity: event Creat_Event(address add, string taskname, string taskcategory, uint256 amount, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) WatchCreatEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractCreatEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContract.contract.WatchLogs(opts, "Creat_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractCreatEvent)
				if err := _TaskDeployerContract.contract.UnpackLog(event, "Creat_Event", log); err != nil {
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

// ParseCreatEvent is a log parse operation binding the contract event 0x7de35515d174e93110b2a3f097c27ebf4cb042509117d9a38b69d5b47b1a2ba3.
//
// Solidity: event Creat_Event(address add, string taskname, string taskcategory, uint256 amount, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) ParseCreatEvent(log types.Log) (*TaskDeployerContractCreatEvent, error) {
	event := new(TaskDeployerContractCreatEvent)
	if err := _TaskDeployerContract.contract.UnpackLog(event, "Creat_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskDeployerContractCreatuserEventIterator is returned from FilterCreatuserEvent and is used to iterate over the raw logs and unpacked data for CreatuserEvent events raised by the TaskDeployerContract contract.
type TaskDeployerContractCreatuserEventIterator struct {
	Event *TaskDeployerContractCreatuserEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractCreatuserEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractCreatuserEvent)
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
		it.Event = new(TaskDeployerContractCreatuserEvent)
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
func (it *TaskDeployerContractCreatuserEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractCreatuserEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractCreatuserEvent represents a CreatuserEvent event raised by the TaskDeployerContract contract.
type TaskDeployerContractCreatuserEvent struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreatuserEvent is a free log retrieval operation binding the contract event 0xc951898903b5802619d88d98319fd44320261ae4adac4bd24849854b79b3477d.
//
// Solidity: event Creatuser_Event(string name, string phonenumber, uint256 studentid)
func (_TaskDeployerContract *TaskDeployerContractFilterer) FilterCreatuserEvent(opts *bind.FilterOpts) (*TaskDeployerContractCreatuserEventIterator, error) {

	logs, sub, err := _TaskDeployerContract.contract.FilterLogs(opts, "Creatuser_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractCreatuserEventIterator{contract: _TaskDeployerContract.contract, event: "Creatuser_Event", logs: logs, sub: sub}, nil
}

// WatchCreatuserEvent is a free log subscription operation binding the contract event 0xc951898903b5802619d88d98319fd44320261ae4adac4bd24849854b79b3477d.
//
// Solidity: event Creatuser_Event(string name, string phonenumber, uint256 studentid)
func (_TaskDeployerContract *TaskDeployerContractFilterer) WatchCreatuserEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractCreatuserEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContract.contract.WatchLogs(opts, "Creatuser_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractCreatuserEvent)
				if err := _TaskDeployerContract.contract.UnpackLog(event, "Creatuser_Event", log); err != nil {
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

// ParseCreatuserEvent is a log parse operation binding the contract event 0xc951898903b5802619d88d98319fd44320261ae4adac4bd24849854b79b3477d.
//
// Solidity: event Creatuser_Event(string name, string phonenumber, uint256 studentid)
func (_TaskDeployerContract *TaskDeployerContractFilterer) ParseCreatuserEvent(log types.Log) (*TaskDeployerContractCreatuserEvent, error) {
	event := new(TaskDeployerContractCreatuserEvent)
	if err := _TaskDeployerContract.contract.UnpackLog(event, "Creatuser_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskDeployerContractClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the TaskDeployerContract contract.
type TaskDeployerContractClaimEventIterator struct {
	Event *TaskDeployerContractClaimEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractClaimEvent)
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
		it.Event = new(TaskDeployerContractClaimEvent)
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
func (it *TaskDeployerContractClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractClaimEvent represents a ClaimEvent event raised by the TaskDeployerContract contract.
type TaskDeployerContractClaimEvent struct {
	Add       common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimEvent is a free log retrieval operation binding the contract event 0x94b84a30e611e2b4293b2a3902a9b9a5030237946d944d181beee0006cb7ab4a.
//
// Solidity: event claim_Event(address add, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) FilterClaimEvent(opts *bind.FilterOpts) (*TaskDeployerContractClaimEventIterator, error) {

	logs, sub, err := _TaskDeployerContract.contract.FilterLogs(opts, "claim_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractClaimEventIterator{contract: _TaskDeployerContract.contract, event: "claim_Event", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x94b84a30e611e2b4293b2a3902a9b9a5030237946d944d181beee0006cb7ab4a.
//
// Solidity: event claim_Event(address add, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractClaimEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContract.contract.WatchLogs(opts, "claim_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractClaimEvent)
				if err := _TaskDeployerContract.contract.UnpackLog(event, "claim_Event", log); err != nil {
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

// ParseClaimEvent is a log parse operation binding the contract event 0x94b84a30e611e2b4293b2a3902a9b9a5030237946d944d181beee0006cb7ab4a.
//
// Solidity: event claim_Event(address add, uint256 timestamp)
func (_TaskDeployerContract *TaskDeployerContractFilterer) ParseClaimEvent(log types.Log) (*TaskDeployerContractClaimEvent, error) {
	event := new(TaskDeployerContractClaimEvent)
	if err := _TaskDeployerContract.contract.UnpackLog(event, "claim_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
