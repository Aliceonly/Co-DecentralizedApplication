// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Smartgo

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

// TaskDeployerContractStructABI is the input ABI used to generate the binding from.
const TaskDeployerContractStructABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Cancel_Event\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"cancelEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_studentid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"}],\"name\":\"changename\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_studentid\",\"type\":\"uint256\"}],\"name\":\"changepassword\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sigs\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_PHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_taskname\",\"type\":\"string\"}],\"name\":\"claimTrust\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Confirm_Event\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"confirmtask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskname\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"taskcategory\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Creat_Event\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_launchTime\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_category\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_taskname\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createNewEvent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phonenumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"studentid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"name\":\"createuser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"phonenumber\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"studentid\",\"type\":\"uint256\"}],\"name\":\"Creatuser_Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"claim_Event\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"svalue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transderToContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_tasklist\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"Taskname\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"category\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"state\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"launchTime\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_userlist\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phonenumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"studentid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalanceOfContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_taskname\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_timestamp\",\"type\":\"string\"}],\"name\":\"getTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TaskDeployerContractStruct is an auto generated Go binding around an Ethereum contract.
type TaskDeployerContractStruct struct {
	TaskDeployerContractStructCaller     // Read-only binding to the contract
	TaskDeployerContractStructTransactor // Write-only binding to the contract
	TaskDeployerContractStructFilterer   // Log filterer for contract events
}

// TaskDeployerContractStructCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskDeployerContractStructCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskDeployerContractStructTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskDeployerContractStructTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskDeployerContractStructFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskDeployerContractStructFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskDeployerContractStructSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskDeployerContractStructSession struct {
	Contract     *TaskDeployerContractStruct // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TaskDeployerContractStructCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskDeployerContractStructCallerSession struct {
	Contract *TaskDeployerContractStructCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// TaskDeployerContractStructTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskDeployerContractStructTransactorSession struct {
	Contract     *TaskDeployerContractStructTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// TaskDeployerContractStructRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskDeployerContractStructRaw struct {
	Contract *TaskDeployerContractStruct // Generic contract binding to access the raw methods on
}

// TaskDeployerContractStructCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskDeployerContractStructCallerRaw struct {
	Contract *TaskDeployerContractStructCaller // Generic read-only contract binding to access the raw methods on
}

// TaskDeployerContractStructTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskDeployerContractStructTransactorRaw struct {
	Contract *TaskDeployerContractStructTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskDeployerContractStruct creates a new instance of TaskDeployerContractStruct, bound to a specific deployed contract.
func NewTaskDeployerContractStruct(address common.Address, backend bind.ContractBackend) (*TaskDeployerContractStruct, error) {
	contract, err := bindTaskDeployerContractStruct(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractStruct{TaskDeployerContractStructCaller: TaskDeployerContractStructCaller{contract: contract}, TaskDeployerContractStructTransactor: TaskDeployerContractStructTransactor{contract: contract}, TaskDeployerContractStructFilterer: TaskDeployerContractStructFilterer{contract: contract}}, nil
}

// NewTaskDeployerContractStructCaller creates a new read-only instance of TaskDeployerContractStruct, bound to a specific deployed contract.
func NewTaskDeployerContractStructCaller(address common.Address, caller bind.ContractCaller) (*TaskDeployerContractStructCaller, error) {
	contract, err := bindTaskDeployerContractStruct(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractStructCaller{contract: contract}, nil
}

// NewTaskDeployerContractStructTransactor creates a new write-only instance of TaskDeployerContractStruct, bound to a specific deployed contract.
func NewTaskDeployerContractStructTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskDeployerContractStructTransactor, error) {
	contract, err := bindTaskDeployerContractStruct(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractStructTransactor{contract: contract}, nil
}

// NewTaskDeployerContractStructFilterer creates a new log filterer instance of TaskDeployerContractStruct, bound to a specific deployed contract.
func NewTaskDeployerContractStructFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskDeployerContractStructFilterer, error) {
	contract, err := bindTaskDeployerContractStruct(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractStructFilterer{contract: contract}, nil
}

// bindTaskDeployerContractStruct binds a generic wrapper to an already deployed contract.
func bindTaskDeployerContractStruct(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskDeployerContractStructABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskDeployerContractStruct *TaskDeployerContractStructRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskDeployerContractStruct.Contract.TaskDeployerContractStructCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskDeployerContractStruct *TaskDeployerContractStructRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.TaskDeployerContractStructTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskDeployerContractStruct *TaskDeployerContractStructRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.TaskDeployerContractStructTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskDeployerContractStruct *TaskDeployerContractStructCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskDeployerContractStruct.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.contract.Transact(opts, method, params...)
}

// Tasklist is a free data retrieval call binding the contract method 0x5ac43ca2.
//
// Solidity: function _tasklist(uint256 ) view returns(string Taskname, address sponsor, address beneficiary, string category, uint256 amount, uint256 timestamp, string state, string launchTime)
func (_TaskDeployerContractStruct *TaskDeployerContractStructCaller) Tasklist(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _TaskDeployerContractStruct.contract.Call(opts, &out, "_tasklist", arg0)

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
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) Tasklist(arg0 *big.Int) (struct {
	Taskname    string
	Sponsor     common.Address
	Beneficiary common.Address
	Category    string
	Amount      *big.Int
	Timestamp   *big.Int
	State       string
	LaunchTime  string
}, error) {
	return _TaskDeployerContractStruct.Contract.Tasklist(&_TaskDeployerContractStruct.CallOpts, arg0)
}

// Tasklist is a free data retrieval call binding the contract method 0x5ac43ca2.
//
// Solidity: function _tasklist(uint256 ) view returns(string Taskname, address sponsor, address beneficiary, string category, uint256 amount, uint256 timestamp, string state, string launchTime)
func (_TaskDeployerContractStruct *TaskDeployerContractStructCallerSession) Tasklist(arg0 *big.Int) (struct {
	Taskname    string
	Sponsor     common.Address
	Beneficiary common.Address
	Category    string
	Amount      *big.Int
	Timestamp   *big.Int
	State       string
	LaunchTime  string
}, error) {
	return _TaskDeployerContractStruct.Contract.Tasklist(&_TaskDeployerContractStruct.CallOpts, arg0)
}

// Userlist is a free data retrieval call binding the contract method 0x885b8290.
//
// Solidity: function _userlist(uint256 ) view returns(string name, string phonenumber, uint256 studentid, string password)
func (_TaskDeployerContractStruct *TaskDeployerContractStructCaller) Userlist(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Password    string
}, error) {
	var out []interface{}
	err := _TaskDeployerContractStruct.contract.Call(opts, &out, "_userlist", arg0)

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
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) Userlist(arg0 *big.Int) (struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Password    string
}, error) {
	return _TaskDeployerContractStruct.Contract.Userlist(&_TaskDeployerContractStruct.CallOpts, arg0)
}

// Userlist is a free data retrieval call binding the contract method 0x885b8290.
//
// Solidity: function _userlist(uint256 ) view returns(string name, string phonenumber, uint256 studentid, string password)
func (_TaskDeployerContractStruct *TaskDeployerContractStructCallerSession) Userlist(arg0 *big.Int) (struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Password    string
}, error) {
	return _TaskDeployerContractStruct.Contract.Userlist(&_TaskDeployerContractStruct.CallOpts, arg0)
}

// GetBalanceOfContract is a free data retrieval call binding the contract method 0x22968885.
//
// Solidity: function getBalanceOfContract() view returns(uint256)
func (_TaskDeployerContractStruct *TaskDeployerContractStructCaller) GetBalanceOfContract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TaskDeployerContractStruct.contract.Call(opts, &out, "getBalanceOfContract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOfContract is a free data retrieval call binding the contract method 0x22968885.
//
// Solidity: function getBalanceOfContract() view returns(uint256)
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) GetBalanceOfContract() (*big.Int, error) {
	return _TaskDeployerContractStruct.Contract.GetBalanceOfContract(&_TaskDeployerContractStruct.CallOpts)
}

// GetBalanceOfContract is a free data retrieval call binding the contract method 0x22968885.
//
// Solidity: function getBalanceOfContract() view returns(uint256)
func (_TaskDeployerContractStruct *TaskDeployerContractStructCallerSession) GetBalanceOfContract() (*big.Int, error) {
	return _TaskDeployerContractStruct.Contract.GetBalanceOfContract(&_TaskDeployerContractStruct.CallOpts)
}

// GetTxHash is a free data retrieval call binding the contract method 0x2c40bf6c.
//
// Solidity: function getTxHash(string _taskname, string _timestamp) view returns(bytes32)
func (_TaskDeployerContractStruct *TaskDeployerContractStructCaller) GetTxHash(opts *bind.CallOpts, _taskname string, _timestamp string) ([32]byte, error) {
	var out []interface{}
	err := _TaskDeployerContractStruct.contract.Call(opts, &out, "getTxHash", _taskname, _timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTxHash is a free data retrieval call binding the contract method 0x2c40bf6c.
//
// Solidity: function getTxHash(string _taskname, string _timestamp) view returns(bytes32)
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) GetTxHash(_taskname string, _timestamp string) ([32]byte, error) {
	return _TaskDeployerContractStruct.Contract.GetTxHash(&_TaskDeployerContractStruct.CallOpts, _taskname, _timestamp)
}

// GetTxHash is a free data retrieval call binding the contract method 0x2c40bf6c.
//
// Solidity: function getTxHash(string _taskname, string _timestamp) view returns(bytes32)
func (_TaskDeployerContractStruct *TaskDeployerContractStructCallerSession) GetTxHash(_taskname string, _timestamp string) ([32]byte, error) {
	return _TaskDeployerContractStruct.Contract.GetTxHash(&_TaskDeployerContractStruct.CallOpts, _taskname, _timestamp)
}

// CancelEvent is a paid mutator transaction binding the contract method 0x3f69babd.
//
// Solidity: function cancelEvent(uint256 _timestamp) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) CancelEvent(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.Transact(opts, "cancelEvent", _timestamp)
}

// CancelEvent is a paid mutator transaction binding the contract method 0x3f69babd.
//
// Solidity: function cancelEvent(uint256 _timestamp) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) CancelEvent(_timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.CancelEvent(&_TaskDeployerContractStruct.TransactOpts, _timestamp)
}

// CancelEvent is a paid mutator transaction binding the contract method 0x3f69babd.
//
// Solidity: function cancelEvent(uint256 _timestamp) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) CancelEvent(_timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.CancelEvent(&_TaskDeployerContractStruct.TransactOpts, _timestamp)
}

// Changename is a paid mutator transaction binding the contract method 0xab3c8449.
//
// Solidity: function changename(string _name, uint256 _studentid, string _password) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) Changename(opts *bind.TransactOpts, _name string, _studentid *big.Int, _password string) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.Transact(opts, "changename", _name, _studentid, _password)
}

// Changename is a paid mutator transaction binding the contract method 0xab3c8449.
//
// Solidity: function changename(string _name, uint256 _studentid, string _password) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) Changename(_name string, _studentid *big.Int, _password string) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Changename(&_TaskDeployerContractStruct.TransactOpts, _name, _studentid, _password)
}

// Changename is a paid mutator transaction binding the contract method 0xab3c8449.
//
// Solidity: function changename(string _name, uint256 _studentid, string _password) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) Changename(_name string, _studentid *big.Int, _password string) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Changename(&_TaskDeployerContractStruct.TransactOpts, _name, _studentid, _password)
}

// Changepassword is a paid mutator transaction binding the contract method 0xf6f18e5d.
//
// Solidity: function changepassword(string _password, uint256 _studentid) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) Changepassword(opts *bind.TransactOpts, _password string, _studentid *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.Transact(opts, "changepassword", _password, _studentid)
}

// Changepassword is a paid mutator transaction binding the contract method 0xf6f18e5d.
//
// Solidity: function changepassword(string _password, uint256 _studentid) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) Changepassword(_password string, _studentid *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Changepassword(&_TaskDeployerContractStruct.TransactOpts, _password, _studentid)
}

// Changepassword is a paid mutator transaction binding the contract method 0xf6f18e5d.
//
// Solidity: function changepassword(string _password, uint256 _studentid) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) Changepassword(_password string, _studentid *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Changepassword(&_TaskDeployerContractStruct.TransactOpts, _password, _studentid)
}

// ClaimTrust is a paid mutator transaction binding the contract method 0xf6bf490d.
//
// Solidity: function claimTrust(uint256 _timestamp, bytes _sigs, bytes32 _PHash, string _taskname) payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) ClaimTrust(opts *bind.TransactOpts, _timestamp *big.Int, _sigs []byte, _PHash [32]byte, _taskname string) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.Transact(opts, "claimTrust", _timestamp, _sigs, _PHash, _taskname)
}

// ClaimTrust is a paid mutator transaction binding the contract method 0xf6bf490d.
//
// Solidity: function claimTrust(uint256 _timestamp, bytes _sigs, bytes32 _PHash, string _taskname) payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) ClaimTrust(_timestamp *big.Int, _sigs []byte, _PHash [32]byte, _taskname string) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.ClaimTrust(&_TaskDeployerContractStruct.TransactOpts, _timestamp, _sigs, _PHash, _taskname)
}

// ClaimTrust is a paid mutator transaction binding the contract method 0xf6bf490d.
//
// Solidity: function claimTrust(uint256 _timestamp, bytes _sigs, bytes32 _PHash, string _taskname) payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) ClaimTrust(_timestamp *big.Int, _sigs []byte, _PHash [32]byte, _taskname string) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.ClaimTrust(&_TaskDeployerContractStruct.TransactOpts, _timestamp, _sigs, _PHash, _taskname)
}

// Confirmtask is a paid mutator transaction binding the contract method 0x7b4623cc.
//
// Solidity: function confirmtask(uint256 _timestamp) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) Confirmtask(opts *bind.TransactOpts, _timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.Transact(opts, "confirmtask", _timestamp)
}

// Confirmtask is a paid mutator transaction binding the contract method 0x7b4623cc.
//
// Solidity: function confirmtask(uint256 _timestamp) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) Confirmtask(_timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Confirmtask(&_TaskDeployerContractStruct.TransactOpts, _timestamp)
}

// Confirmtask is a paid mutator transaction binding the contract method 0x7b4623cc.
//
// Solidity: function confirmtask(uint256 _timestamp) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) Confirmtask(_timestamp *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Confirmtask(&_TaskDeployerContractStruct.TransactOpts, _timestamp)
}

// CreateNewEvent is a paid mutator transaction binding the contract method 0x63e0a525.
//
// Solidity: function createNewEvent(string _launchTime, string _category, string _taskname, uint256 _amount) payable returns(uint256)
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) CreateNewEvent(opts *bind.TransactOpts, _launchTime string, _category string, _taskname string, _amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.Transact(opts, "createNewEvent", _launchTime, _category, _taskname, _amount)
}

// CreateNewEvent is a paid mutator transaction binding the contract method 0x63e0a525.
//
// Solidity: function createNewEvent(string _launchTime, string _category, string _taskname, uint256 _amount) payable returns(uint256)
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) CreateNewEvent(_launchTime string, _category string, _taskname string, _amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.CreateNewEvent(&_TaskDeployerContractStruct.TransactOpts, _launchTime, _category, _taskname, _amount)
}

// CreateNewEvent is a paid mutator transaction binding the contract method 0x63e0a525.
//
// Solidity: function createNewEvent(string _launchTime, string _category, string _taskname, uint256 _amount) payable returns(uint256)
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) CreateNewEvent(_launchTime string, _category string, _taskname string, _amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.CreateNewEvent(&_TaskDeployerContractStruct.TransactOpts, _launchTime, _category, _taskname, _amount)
}

// Createuser is a paid mutator transaction binding the contract method 0x3fb93885.
//
// Solidity: function createuser(string name, string phonenumber, uint256 studentid, string password) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) Createuser(opts *bind.TransactOpts, name string, phonenumber string, studentid *big.Int, password string) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.Transact(opts, "createuser", name, phonenumber, studentid, password)
}

// Createuser is a paid mutator transaction binding the contract method 0x3fb93885.
//
// Solidity: function createuser(string name, string phonenumber, uint256 studentid, string password) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) Createuser(name string, phonenumber string, studentid *big.Int, password string) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Createuser(&_TaskDeployerContractStruct.TransactOpts, name, phonenumber, studentid, password)
}

// Createuser is a paid mutator transaction binding the contract method 0x3fb93885.
//
// Solidity: function createuser(string name, string phonenumber, uint256 studentid, string password) returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) Createuser(name string, phonenumber string, studentid *big.Int, password string) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Createuser(&_TaskDeployerContractStruct.TransactOpts, name, phonenumber, studentid, password)
}

// Svalue is a paid mutator transaction binding the contract method 0xfc6378a4.
//
// Solidity: function svalue(address addr, uint256 amount) payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) Svalue(opts *bind.TransactOpts, addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.Transact(opts, "svalue", addr, amount)
}

// Svalue is a paid mutator transaction binding the contract method 0xfc6378a4.
//
// Solidity: function svalue(address addr, uint256 amount) payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) Svalue(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Svalue(&_TaskDeployerContractStruct.TransactOpts, addr, amount)
}

// Svalue is a paid mutator transaction binding the contract method 0xfc6378a4.
//
// Solidity: function svalue(address addr, uint256 amount) payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) Svalue(addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Svalue(&_TaskDeployerContractStruct.TransactOpts, addr, amount)
}

// TransderToContract is a paid mutator transaction binding the contract method 0x460968dd.
//
// Solidity: function transderToContract() payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) TransderToContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.Transact(opts, "transderToContract")
}

// TransderToContract is a paid mutator transaction binding the contract method 0x460968dd.
//
// Solidity: function transderToContract() payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) TransderToContract() (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.TransderToContract(&_TaskDeployerContractStruct.TransactOpts)
}

// TransderToContract is a paid mutator transaction binding the contract method 0x460968dd.
//
// Solidity: function transderToContract() payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) TransderToContract() (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.TransderToContract(&_TaskDeployerContractStruct.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Fallback(&_TaskDeployerContractStruct.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Fallback(&_TaskDeployerContractStruct.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskDeployerContractStruct.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructSession) Receive() (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Receive(&_TaskDeployerContractStruct.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaskDeployerContractStruct *TaskDeployerContractStructTransactorSession) Receive() (*types.Transaction, error) {
	return _TaskDeployerContractStruct.Contract.Receive(&_TaskDeployerContractStruct.TransactOpts)
}

// TaskDeployerContractStructCancelEventIterator is returned from FilterCancelEvent and is used to iterate over the raw logs and unpacked data for CancelEvent events raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructCancelEventIterator struct {
	Event *TaskDeployerContractStructCancelEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractStructCancelEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractStructCancelEvent)
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
		it.Event = new(TaskDeployerContractStructCancelEvent)
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
func (it *TaskDeployerContractStructCancelEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractStructCancelEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractStructCancelEvent represents a CancelEvent event raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructCancelEvent struct {
	Add       common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCancelEvent is a free log retrieval operation binding the contract event 0xebb62664a922b45fff0b6d8448a6f26157da006cd1887a3311f5a678b598dda4.
//
// Solidity: event Cancel_Event(address add, uint256 timestamp)
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) FilterCancelEvent(opts *bind.FilterOpts) (*TaskDeployerContractStructCancelEventIterator, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.FilterLogs(opts, "Cancel_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractStructCancelEventIterator{contract: _TaskDeployerContractStruct.contract, event: "Cancel_Event", logs: logs, sub: sub}, nil
}

// WatchCancelEvent is a free log subscription operation binding the contract event 0xebb62664a922b45fff0b6d8448a6f26157da006cd1887a3311f5a678b598dda4.
//
// Solidity: event Cancel_Event(address add, uint256 timestamp)
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) WatchCancelEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractStructCancelEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.WatchLogs(opts, "Cancel_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractStructCancelEvent)
				if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "Cancel_Event", log); err != nil {
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
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) ParseCancelEvent(log types.Log) (*TaskDeployerContractStructCancelEvent, error) {
	event := new(TaskDeployerContractStructCancelEvent)
	if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "Cancel_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskDeployerContractStructConfirmEventIterator is returned from FilterConfirmEvent and is used to iterate over the raw logs and unpacked data for ConfirmEvent events raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructConfirmEventIterator struct {
	Event *TaskDeployerContractStructConfirmEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractStructConfirmEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractStructConfirmEvent)
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
		it.Event = new(TaskDeployerContractStructConfirmEvent)
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
func (it *TaskDeployerContractStructConfirmEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractStructConfirmEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractStructConfirmEvent represents a ConfirmEvent event raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructConfirmEvent struct {
	Add       common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirmEvent is a free log retrieval operation binding the contract event 0x6b9ce9a8215bda6518e4fb80be40bc0ce2da320e58757a4aef77aca3a1d2fb4e.
//
// Solidity: event Confirm_Event(address add, uint256 timestamp)
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) FilterConfirmEvent(opts *bind.FilterOpts) (*TaskDeployerContractStructConfirmEventIterator, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.FilterLogs(opts, "Confirm_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractStructConfirmEventIterator{contract: _TaskDeployerContractStruct.contract, event: "Confirm_Event", logs: logs, sub: sub}, nil
}

// WatchConfirmEvent is a free log subscription operation binding the contract event 0x6b9ce9a8215bda6518e4fb80be40bc0ce2da320e58757a4aef77aca3a1d2fb4e.
//
// Solidity: event Confirm_Event(address add, uint256 timestamp)
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) WatchConfirmEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractStructConfirmEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.WatchLogs(opts, "Confirm_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractStructConfirmEvent)
				if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "Confirm_Event", log); err != nil {
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
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) ParseConfirmEvent(log types.Log) (*TaskDeployerContractStructConfirmEvent, error) {
	event := new(TaskDeployerContractStructConfirmEvent)
	if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "Confirm_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskDeployerContractStructCreatEventIterator is returned from FilterCreatEvent and is used to iterate over the raw logs and unpacked data for CreatEvent events raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructCreatEventIterator struct {
	Event *TaskDeployerContractStructCreatEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractStructCreatEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractStructCreatEvent)
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
		it.Event = new(TaskDeployerContractStructCreatEvent)
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
func (it *TaskDeployerContractStructCreatEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractStructCreatEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractStructCreatEvent represents a CreatEvent event raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructCreatEvent struct {
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
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) FilterCreatEvent(opts *bind.FilterOpts) (*TaskDeployerContractStructCreatEventIterator, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.FilterLogs(opts, "Creat_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractStructCreatEventIterator{contract: _TaskDeployerContractStruct.contract, event: "Creat_Event", logs: logs, sub: sub}, nil
}

// WatchCreatEvent is a free log subscription operation binding the contract event 0x7de35515d174e93110b2a3f097c27ebf4cb042509117d9a38b69d5b47b1a2ba3.
//
// Solidity: event Creat_Event(address add, string taskname, string taskcategory, uint256 amount, uint256 timestamp)
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) WatchCreatEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractStructCreatEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.WatchLogs(opts, "Creat_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractStructCreatEvent)
				if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "Creat_Event", log); err != nil {
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
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) ParseCreatEvent(log types.Log) (*TaskDeployerContractStructCreatEvent, error) {
	event := new(TaskDeployerContractStructCreatEvent)
	if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "Creat_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskDeployerContractStructCreatuserEventIterator is returned from FilterCreatuserEvent and is used to iterate over the raw logs and unpacked data for CreatuserEvent events raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructCreatuserEventIterator struct {
	Event *TaskDeployerContractStructCreatuserEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractStructCreatuserEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractStructCreatuserEvent)
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
		it.Event = new(TaskDeployerContractStructCreatuserEvent)
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
func (it *TaskDeployerContractStructCreatuserEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractStructCreatuserEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractStructCreatuserEvent represents a CreatuserEvent event raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructCreatuserEvent struct {
	Name        string
	Phonenumber string
	Studentid   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreatuserEvent is a free log retrieval operation binding the contract event 0xc951898903b5802619d88d98319fd44320261ae4adac4bd24849854b79b3477d.
//
// Solidity: event Creatuser_Event(string name, string phonenumber, uint256 studentid)
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) FilterCreatuserEvent(opts *bind.FilterOpts) (*TaskDeployerContractStructCreatuserEventIterator, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.FilterLogs(opts, "Creatuser_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractStructCreatuserEventIterator{contract: _TaskDeployerContractStruct.contract, event: "Creatuser_Event", logs: logs, sub: sub}, nil
}

// WatchCreatuserEvent is a free log subscription operation binding the contract event 0xc951898903b5802619d88d98319fd44320261ae4adac4bd24849854b79b3477d.
//
// Solidity: event Creatuser_Event(string name, string phonenumber, uint256 studentid)
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) WatchCreatuserEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractStructCreatuserEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.WatchLogs(opts, "Creatuser_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractStructCreatuserEvent)
				if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "Creatuser_Event", log); err != nil {
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
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) ParseCreatuserEvent(log types.Log) (*TaskDeployerContractStructCreatuserEvent, error) {
	event := new(TaskDeployerContractStructCreatuserEvent)
	if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "Creatuser_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskDeployerContractStructClaimEventIterator is returned from FilterClaimEvent and is used to iterate over the raw logs and unpacked data for ClaimEvent events raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructClaimEventIterator struct {
	Event *TaskDeployerContractStructClaimEvent // Event containing the contract specifics and raw log

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
func (it *TaskDeployerContractStructClaimEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDeployerContractStructClaimEvent)
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
		it.Event = new(TaskDeployerContractStructClaimEvent)
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
func (it *TaskDeployerContractStructClaimEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDeployerContractStructClaimEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDeployerContractStructClaimEvent represents a ClaimEvent event raised by the TaskDeployerContractStruct contract.
type TaskDeployerContractStructClaimEvent struct {
	Add       common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimEvent is a free log retrieval operation binding the contract event 0x94b84a30e611e2b4293b2a3902a9b9a5030237946d944d181beee0006cb7ab4a.
//
// Solidity: event claim_Event(address add, uint256 timestamp)
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) FilterClaimEvent(opts *bind.FilterOpts) (*TaskDeployerContractStructClaimEventIterator, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.FilterLogs(opts, "claim_Event")
	if err != nil {
		return nil, err
	}
	return &TaskDeployerContractStructClaimEventIterator{contract: _TaskDeployerContractStruct.contract, event: "claim_Event", logs: logs, sub: sub}, nil
}

// WatchClaimEvent is a free log subscription operation binding the contract event 0x94b84a30e611e2b4293b2a3902a9b9a5030237946d944d181beee0006cb7ab4a.
//
// Solidity: event claim_Event(address add, uint256 timestamp)
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) WatchClaimEvent(opts *bind.WatchOpts, sink chan<- *TaskDeployerContractStructClaimEvent) (event.Subscription, error) {

	logs, sub, err := _TaskDeployerContractStruct.contract.WatchLogs(opts, "claim_Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDeployerContractStructClaimEvent)
				if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "claim_Event", log); err != nil {
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
func (_TaskDeployerContractStruct *TaskDeployerContractStructFilterer) ParseClaimEvent(log types.Log) (*TaskDeployerContractStructClaimEvent, error) {
	event := new(TaskDeployerContractStructClaimEvent)
	if err := _TaskDeployerContractStruct.contract.UnpackLog(event, "claim_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
