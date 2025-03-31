// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ifiledid

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IfiledidMetaData contains all meta data concerning the Ifiledid contract.
var IfiledidMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"BuyRead\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"}],\"name\":\"ChangeController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"}],\"name\":\"ChangeFtype\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"ChangeKeywords\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ChangePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"}],\"name\":\"DeactivateMfileDid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"DeactivateRead\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"GrantRead\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"RegisterMfileDid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"buyRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"}],\"name\":\"changeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"}],\"name\":\"changeFtype\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"changeKeywords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"}],\"name\":\"deactivateMfileDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"deactivateRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"deactivated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getEncode\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getFtype\",\"outputs\":[{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getKeywords\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"grantRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"encode\",\"type\":\"string\"},{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"registerMfileDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IfiledidABI is the input ABI used to generate the binding from.
// Deprecated: Use IfiledidMetaData.ABI instead.
var IfiledidABI = IfiledidMetaData.ABI

// Ifiledid is an auto generated Go binding around an Ethereum contract.
type Ifiledid struct {
	IfiledidCaller     // Read-only binding to the contract
	IfiledidTransactor // Write-only binding to the contract
	IfiledidFilterer   // Log filterer for contract events
}

// IfiledidCaller is an auto generated read-only Go binding around an Ethereum contract.
type IfiledidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IfiledidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IfiledidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IfiledidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IfiledidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IfiledidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IfiledidSession struct {
	Contract     *Ifiledid         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IfiledidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IfiledidCallerSession struct {
	Contract *IfiledidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IfiledidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IfiledidTransactorSession struct {
	Contract     *IfiledidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IfiledidRaw is an auto generated low-level Go binding around an Ethereum contract.
type IfiledidRaw struct {
	Contract *Ifiledid // Generic contract binding to access the raw methods on
}

// IfiledidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IfiledidCallerRaw struct {
	Contract *IfiledidCaller // Generic read-only contract binding to access the raw methods on
}

// IfiledidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IfiledidTransactorRaw struct {
	Contract *IfiledidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIfiledid creates a new instance of Ifiledid, bound to a specific deployed contract.
func NewIfiledid(address common.Address, backend bind.ContractBackend) (*Ifiledid, error) {
	contract, err := bindIfiledid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ifiledid{IfiledidCaller: IfiledidCaller{contract: contract}, IfiledidTransactor: IfiledidTransactor{contract: contract}, IfiledidFilterer: IfiledidFilterer{contract: contract}}, nil
}

// NewIfiledidCaller creates a new read-only instance of Ifiledid, bound to a specific deployed contract.
func NewIfiledidCaller(address common.Address, caller bind.ContractCaller) (*IfiledidCaller, error) {
	contract, err := bindIfiledid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IfiledidCaller{contract: contract}, nil
}

// NewIfiledidTransactor creates a new write-only instance of Ifiledid, bound to a specific deployed contract.
func NewIfiledidTransactor(address common.Address, transactor bind.ContractTransactor) (*IfiledidTransactor, error) {
	contract, err := bindIfiledid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IfiledidTransactor{contract: contract}, nil
}

// NewIfiledidFilterer creates a new log filterer instance of Ifiledid, bound to a specific deployed contract.
func NewIfiledidFilterer(address common.Address, filterer bind.ContractFilterer) (*IfiledidFilterer, error) {
	contract, err := bindIfiledid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IfiledidFilterer{contract: contract}, nil
}

// bindIfiledid binds a generic wrapper to an already deployed contract.
func bindIfiledid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IfiledidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ifiledid *IfiledidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ifiledid.Contract.IfiledidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ifiledid *IfiledidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ifiledid.Contract.IfiledidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ifiledid *IfiledidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ifiledid.Contract.IfiledidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ifiledid *IfiledidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ifiledid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ifiledid *IfiledidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ifiledid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ifiledid *IfiledidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ifiledid.Contract.contract.Transact(opts, method, params...)
}

// Deactivated is a free data retrieval call binding the contract method 0x84435ce2.
//
// Solidity: function deactivated(string mfileDid) view returns(bool)
func (_Ifiledid *IfiledidCaller) Deactivated(opts *bind.CallOpts, mfileDid string) (bool, error) {
	var out []interface{}
	err := _Ifiledid.contract.Call(opts, &out, "deactivated", mfileDid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Deactivated is a free data retrieval call binding the contract method 0x84435ce2.
//
// Solidity: function deactivated(string mfileDid) view returns(bool)
func (_Ifiledid *IfiledidSession) Deactivated(mfileDid string) (bool, error) {
	return _Ifiledid.Contract.Deactivated(&_Ifiledid.CallOpts, mfileDid)
}

// Deactivated is a free data retrieval call binding the contract method 0x84435ce2.
//
// Solidity: function deactivated(string mfileDid) view returns(bool)
func (_Ifiledid *IfiledidCallerSession) Deactivated(mfileDid string) (bool, error) {
	return _Ifiledid.Contract.Deactivated(&_Ifiledid.CallOpts, mfileDid)
}

// GetController is a free data retrieval call binding the contract method 0x63a27111.
//
// Solidity: function getController(string mfileDid) view returns(string)
func (_Ifiledid *IfiledidCaller) GetController(opts *bind.CallOpts, mfileDid string) (string, error) {
	var out []interface{}
	err := _Ifiledid.contract.Call(opts, &out, "getController", mfileDid)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0x63a27111.
//
// Solidity: function getController(string mfileDid) view returns(string)
func (_Ifiledid *IfiledidSession) GetController(mfileDid string) (string, error) {
	return _Ifiledid.Contract.GetController(&_Ifiledid.CallOpts, mfileDid)
}

// GetController is a free data retrieval call binding the contract method 0x63a27111.
//
// Solidity: function getController(string mfileDid) view returns(string)
func (_Ifiledid *IfiledidCallerSession) GetController(mfileDid string) (string, error) {
	return _Ifiledid.Contract.GetController(&_Ifiledid.CallOpts, mfileDid)
}

// GetEncode is a free data retrieval call binding the contract method 0xf100cfd3.
//
// Solidity: function getEncode(string mfileDid) view returns(string)
func (_Ifiledid *IfiledidCaller) GetEncode(opts *bind.CallOpts, mfileDid string) (string, error) {
	var out []interface{}
	err := _Ifiledid.contract.Call(opts, &out, "getEncode", mfileDid)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEncode is a free data retrieval call binding the contract method 0xf100cfd3.
//
// Solidity: function getEncode(string mfileDid) view returns(string)
func (_Ifiledid *IfiledidSession) GetEncode(mfileDid string) (string, error) {
	return _Ifiledid.Contract.GetEncode(&_Ifiledid.CallOpts, mfileDid)
}

// GetEncode is a free data retrieval call binding the contract method 0xf100cfd3.
//
// Solidity: function getEncode(string mfileDid) view returns(string)
func (_Ifiledid *IfiledidCallerSession) GetEncode(mfileDid string) (string, error) {
	return _Ifiledid.Contract.GetEncode(&_Ifiledid.CallOpts, mfileDid)
}

// GetFtype is a free data retrieval call binding the contract method 0xdff5b013.
//
// Solidity: function getFtype(string mfileDid) view returns(uint8)
func (_Ifiledid *IfiledidCaller) GetFtype(opts *bind.CallOpts, mfileDid string) (uint8, error) {
	var out []interface{}
	err := _Ifiledid.contract.Call(opts, &out, "getFtype", mfileDid)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetFtype is a free data retrieval call binding the contract method 0xdff5b013.
//
// Solidity: function getFtype(string mfileDid) view returns(uint8)
func (_Ifiledid *IfiledidSession) GetFtype(mfileDid string) (uint8, error) {
	return _Ifiledid.Contract.GetFtype(&_Ifiledid.CallOpts, mfileDid)
}

// GetFtype is a free data retrieval call binding the contract method 0xdff5b013.
//
// Solidity: function getFtype(string mfileDid) view returns(uint8)
func (_Ifiledid *IfiledidCallerSession) GetFtype(mfileDid string) (uint8, error) {
	return _Ifiledid.Contract.GetFtype(&_Ifiledid.CallOpts, mfileDid)
}

// GetKeywords is a free data retrieval call binding the contract method 0x380cc539.
//
// Solidity: function getKeywords(string mfileDid) view returns(string[])
func (_Ifiledid *IfiledidCaller) GetKeywords(opts *bind.CallOpts, mfileDid string) ([]string, error) {
	var out []interface{}
	err := _Ifiledid.contract.Call(opts, &out, "getKeywords", mfileDid)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetKeywords is a free data retrieval call binding the contract method 0x380cc539.
//
// Solidity: function getKeywords(string mfileDid) view returns(string[])
func (_Ifiledid *IfiledidSession) GetKeywords(mfileDid string) ([]string, error) {
	return _Ifiledid.Contract.GetKeywords(&_Ifiledid.CallOpts, mfileDid)
}

// GetKeywords is a free data retrieval call binding the contract method 0x380cc539.
//
// Solidity: function getKeywords(string mfileDid) view returns(string[])
func (_Ifiledid *IfiledidCallerSession) GetKeywords(mfileDid string) ([]string, error) {
	return _Ifiledid.Contract.GetKeywords(&_Ifiledid.CallOpts, mfileDid)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string mfileDid) view returns(uint256)
func (_Ifiledid *IfiledidCaller) GetPrice(opts *bind.CallOpts, mfileDid string) (*big.Int, error) {
	var out []interface{}
	err := _Ifiledid.contract.Call(opts, &out, "getPrice", mfileDid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string mfileDid) view returns(uint256)
func (_Ifiledid *IfiledidSession) GetPrice(mfileDid string) (*big.Int, error) {
	return _Ifiledid.Contract.GetPrice(&_Ifiledid.CallOpts, mfileDid)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string mfileDid) view returns(uint256)
func (_Ifiledid *IfiledidCallerSession) GetPrice(mfileDid string) (*big.Int, error) {
	return _Ifiledid.Contract.GetPrice(&_Ifiledid.CallOpts, mfileDid)
}

// Read is a free data retrieval call binding the contract method 0x8c97f99e.
//
// Solidity: function read(string mfileDid, string memoDid) view returns(uint8)
func (_Ifiledid *IfiledidCaller) Read(opts *bind.CallOpts, mfileDid string, memoDid string) (uint8, error) {
	var out []interface{}
	err := _Ifiledid.contract.Call(opts, &out, "read", mfileDid, memoDid)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Read is a free data retrieval call binding the contract method 0x8c97f99e.
//
// Solidity: function read(string mfileDid, string memoDid) view returns(uint8)
func (_Ifiledid *IfiledidSession) Read(mfileDid string, memoDid string) (uint8, error) {
	return _Ifiledid.Contract.Read(&_Ifiledid.CallOpts, mfileDid, memoDid)
}

// Read is a free data retrieval call binding the contract method 0x8c97f99e.
//
// Solidity: function read(string mfileDid, string memoDid) view returns(uint8)
func (_Ifiledid *IfiledidCallerSession) Read(mfileDid string, memoDid string) (uint8, error) {
	return _Ifiledid.Contract.Read(&_Ifiledid.CallOpts, mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_Ifiledid *IfiledidTransactor) BuyRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Ifiledid.contract.Transact(opts, "buyRead", mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_Ifiledid *IfiledidSession) BuyRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Ifiledid.Contract.BuyRead(&_Ifiledid.TransactOpts, mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_Ifiledid *IfiledidTransactorSession) BuyRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Ifiledid.Contract.BuyRead(&_Ifiledid.TransactOpts, mfileDid, memoDid)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_Ifiledid *IfiledidTransactor) ChangeController(opts *bind.TransactOpts, mfileDid string, controller string) (*types.Transaction, error) {
	return _Ifiledid.contract.Transact(opts, "changeController", mfileDid, controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_Ifiledid *IfiledidSession) ChangeController(mfileDid string, controller string) (*types.Transaction, error) {
	return _Ifiledid.Contract.ChangeController(&_Ifiledid.TransactOpts, mfileDid, controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_Ifiledid *IfiledidTransactorSession) ChangeController(mfileDid string, controller string) (*types.Transaction, error) {
	return _Ifiledid.Contract.ChangeController(&_Ifiledid.TransactOpts, mfileDid, controller)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_Ifiledid *IfiledidTransactor) ChangeFtype(opts *bind.TransactOpts, mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _Ifiledid.contract.Transact(opts, "changeFtype", mfileDid, ftype)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_Ifiledid *IfiledidSession) ChangeFtype(mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _Ifiledid.Contract.ChangeFtype(&_Ifiledid.TransactOpts, mfileDid, ftype)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_Ifiledid *IfiledidTransactorSession) ChangeFtype(mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _Ifiledid.Contract.ChangeFtype(&_Ifiledid.TransactOpts, mfileDid, ftype)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_Ifiledid *IfiledidTransactor) ChangeKeywords(opts *bind.TransactOpts, mfileDid string, keywords []string) (*types.Transaction, error) {
	return _Ifiledid.contract.Transact(opts, "changeKeywords", mfileDid, keywords)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_Ifiledid *IfiledidSession) ChangeKeywords(mfileDid string, keywords []string) (*types.Transaction, error) {
	return _Ifiledid.Contract.ChangeKeywords(&_Ifiledid.TransactOpts, mfileDid, keywords)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_Ifiledid *IfiledidTransactorSession) ChangeKeywords(mfileDid string, keywords []string) (*types.Transaction, error) {
	return _Ifiledid.Contract.ChangeKeywords(&_Ifiledid.TransactOpts, mfileDid, keywords)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_Ifiledid *IfiledidTransactor) ChangePrice(opts *bind.TransactOpts, mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _Ifiledid.contract.Transact(opts, "changePrice", mfileDid, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_Ifiledid *IfiledidSession) ChangePrice(mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _Ifiledid.Contract.ChangePrice(&_Ifiledid.TransactOpts, mfileDid, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_Ifiledid *IfiledidTransactorSession) ChangePrice(mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _Ifiledid.Contract.ChangePrice(&_Ifiledid.TransactOpts, mfileDid, price)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_Ifiledid *IfiledidTransactor) DeactivateMfileDid(opts *bind.TransactOpts, mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _Ifiledid.contract.Transact(opts, "deactivateMfileDid", mfileDid, deactivate)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_Ifiledid *IfiledidSession) DeactivateMfileDid(mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _Ifiledid.Contract.DeactivateMfileDid(&_Ifiledid.TransactOpts, mfileDid, deactivate)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_Ifiledid *IfiledidTransactorSession) DeactivateMfileDid(mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _Ifiledid.Contract.DeactivateMfileDid(&_Ifiledid.TransactOpts, mfileDid, deactivate)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_Ifiledid *IfiledidTransactor) DeactivateRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Ifiledid.contract.Transact(opts, "deactivateRead", mfileDid, memoDid)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_Ifiledid *IfiledidSession) DeactivateRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Ifiledid.Contract.DeactivateRead(&_Ifiledid.TransactOpts, mfileDid, memoDid)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_Ifiledid *IfiledidTransactorSession) DeactivateRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Ifiledid.Contract.DeactivateRead(&_Ifiledid.TransactOpts, mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_Ifiledid *IfiledidTransactor) GrantRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Ifiledid.contract.Transact(opts, "grantRead", mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_Ifiledid *IfiledidSession) GrantRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Ifiledid.Contract.GrantRead(&_Ifiledid.TransactOpts, mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_Ifiledid *IfiledidTransactorSession) GrantRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Ifiledid.Contract.GrantRead(&_Ifiledid.TransactOpts, mfileDid, memoDid)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_Ifiledid *IfiledidTransactor) RegisterMfileDid(opts *bind.TransactOpts, mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _Ifiledid.contract.Transact(opts, "registerMfileDid", mfileDid, encode, ftype, controller, price, keywords)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_Ifiledid *IfiledidSession) RegisterMfileDid(mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _Ifiledid.Contract.RegisterMfileDid(&_Ifiledid.TransactOpts, mfileDid, encode, ftype, controller, price, keywords)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_Ifiledid *IfiledidTransactorSession) RegisterMfileDid(mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _Ifiledid.Contract.RegisterMfileDid(&_Ifiledid.TransactOpts, mfileDid, encode, ftype, controller, price, keywords)
}

// IfiledidBuyReadIterator is returned from FilterBuyRead and is used to iterate over the raw logs and unpacked data for BuyRead events raised by the Ifiledid contract.
type IfiledidBuyReadIterator struct {
	Event *IfiledidBuyRead // Event containing the contract specifics and raw log

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
func (it *IfiledidBuyReadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IfiledidBuyRead)
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
		it.Event = new(IfiledidBuyRead)
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
func (it *IfiledidBuyReadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IfiledidBuyReadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IfiledidBuyRead represents a BuyRead event raised by the Ifiledid contract.
type IfiledidBuyRead struct {
	MfileDid common.Hash
	MemoDid  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBuyRead is a free log retrieval operation binding the contract event 0x9a23ee694031b7714b282fb4f89b349db581b6c5fb20fb299897b2dbb5b6510a.
//
// Solidity: event BuyRead(string indexed mfileDid, string memoDid)
func (_Ifiledid *IfiledidFilterer) FilterBuyRead(opts *bind.FilterOpts, mfileDid []string) (*IfiledidBuyReadIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.FilterLogs(opts, "BuyRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IfiledidBuyReadIterator{contract: _Ifiledid.contract, event: "BuyRead", logs: logs, sub: sub}, nil
}

// WatchBuyRead is a free log subscription operation binding the contract event 0x9a23ee694031b7714b282fb4f89b349db581b6c5fb20fb299897b2dbb5b6510a.
//
// Solidity: event BuyRead(string indexed mfileDid, string memoDid)
func (_Ifiledid *IfiledidFilterer) WatchBuyRead(opts *bind.WatchOpts, sink chan<- *IfiledidBuyRead, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.WatchLogs(opts, "BuyRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IfiledidBuyRead)
				if err := _Ifiledid.contract.UnpackLog(event, "BuyRead", log); err != nil {
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

// ParseBuyRead is a log parse operation binding the contract event 0x9a23ee694031b7714b282fb4f89b349db581b6c5fb20fb299897b2dbb5b6510a.
//
// Solidity: event BuyRead(string indexed mfileDid, string memoDid)
func (_Ifiledid *IfiledidFilterer) ParseBuyRead(log types.Log) (*IfiledidBuyRead, error) {
	event := new(IfiledidBuyRead)
	if err := _Ifiledid.contract.UnpackLog(event, "BuyRead", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IfiledidChangeControllerIterator is returned from FilterChangeController and is used to iterate over the raw logs and unpacked data for ChangeController events raised by the Ifiledid contract.
type IfiledidChangeControllerIterator struct {
	Event *IfiledidChangeController // Event containing the contract specifics and raw log

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
func (it *IfiledidChangeControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IfiledidChangeController)
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
		it.Event = new(IfiledidChangeController)
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
func (it *IfiledidChangeControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IfiledidChangeControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IfiledidChangeController represents a ChangeController event raised by the Ifiledid contract.
type IfiledidChangeController struct {
	MfileDid   common.Hash
	Controller string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChangeController is a free log retrieval operation binding the contract event 0xad2c5e83d2fbe2975f5ba70ab5c8c351ef2da897e41ce843a2eb6de9ffbf59ee.
//
// Solidity: event ChangeController(string indexed mfileDid, string controller)
func (_Ifiledid *IfiledidFilterer) FilterChangeController(opts *bind.FilterOpts, mfileDid []string) (*IfiledidChangeControllerIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.FilterLogs(opts, "ChangeController", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IfiledidChangeControllerIterator{contract: _Ifiledid.contract, event: "ChangeController", logs: logs, sub: sub}, nil
}

// WatchChangeController is a free log subscription operation binding the contract event 0xad2c5e83d2fbe2975f5ba70ab5c8c351ef2da897e41ce843a2eb6de9ffbf59ee.
//
// Solidity: event ChangeController(string indexed mfileDid, string controller)
func (_Ifiledid *IfiledidFilterer) WatchChangeController(opts *bind.WatchOpts, sink chan<- *IfiledidChangeController, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.WatchLogs(opts, "ChangeController", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IfiledidChangeController)
				if err := _Ifiledid.contract.UnpackLog(event, "ChangeController", log); err != nil {
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

// ParseChangeController is a log parse operation binding the contract event 0xad2c5e83d2fbe2975f5ba70ab5c8c351ef2da897e41ce843a2eb6de9ffbf59ee.
//
// Solidity: event ChangeController(string indexed mfileDid, string controller)
func (_Ifiledid *IfiledidFilterer) ParseChangeController(log types.Log) (*IfiledidChangeController, error) {
	event := new(IfiledidChangeController)
	if err := _Ifiledid.contract.UnpackLog(event, "ChangeController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IfiledidChangeFtypeIterator is returned from FilterChangeFtype and is used to iterate over the raw logs and unpacked data for ChangeFtype events raised by the Ifiledid contract.
type IfiledidChangeFtypeIterator struct {
	Event *IfiledidChangeFtype // Event containing the contract specifics and raw log

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
func (it *IfiledidChangeFtypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IfiledidChangeFtype)
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
		it.Event = new(IfiledidChangeFtype)
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
func (it *IfiledidChangeFtypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IfiledidChangeFtypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IfiledidChangeFtype represents a ChangeFtype event raised by the Ifiledid contract.
type IfiledidChangeFtype struct {
	MfileDid common.Hash
	Ftype    uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeFtype is a free log retrieval operation binding the contract event 0x935721402a5ae34f86d9500f87d6acfbef247c62cc5c845f49e7ec94a0dd4896.
//
// Solidity: event ChangeFtype(string indexed mfileDid, uint8 ftype)
func (_Ifiledid *IfiledidFilterer) FilterChangeFtype(opts *bind.FilterOpts, mfileDid []string) (*IfiledidChangeFtypeIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.FilterLogs(opts, "ChangeFtype", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IfiledidChangeFtypeIterator{contract: _Ifiledid.contract, event: "ChangeFtype", logs: logs, sub: sub}, nil
}

// WatchChangeFtype is a free log subscription operation binding the contract event 0x935721402a5ae34f86d9500f87d6acfbef247c62cc5c845f49e7ec94a0dd4896.
//
// Solidity: event ChangeFtype(string indexed mfileDid, uint8 ftype)
func (_Ifiledid *IfiledidFilterer) WatchChangeFtype(opts *bind.WatchOpts, sink chan<- *IfiledidChangeFtype, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.WatchLogs(opts, "ChangeFtype", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IfiledidChangeFtype)
				if err := _Ifiledid.contract.UnpackLog(event, "ChangeFtype", log); err != nil {
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

// ParseChangeFtype is a log parse operation binding the contract event 0x935721402a5ae34f86d9500f87d6acfbef247c62cc5c845f49e7ec94a0dd4896.
//
// Solidity: event ChangeFtype(string indexed mfileDid, uint8 ftype)
func (_Ifiledid *IfiledidFilterer) ParseChangeFtype(log types.Log) (*IfiledidChangeFtype, error) {
	event := new(IfiledidChangeFtype)
	if err := _Ifiledid.contract.UnpackLog(event, "ChangeFtype", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IfiledidChangeKeywordsIterator is returned from FilterChangeKeywords and is used to iterate over the raw logs and unpacked data for ChangeKeywords events raised by the Ifiledid contract.
type IfiledidChangeKeywordsIterator struct {
	Event *IfiledidChangeKeywords // Event containing the contract specifics and raw log

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
func (it *IfiledidChangeKeywordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IfiledidChangeKeywords)
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
		it.Event = new(IfiledidChangeKeywords)
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
func (it *IfiledidChangeKeywordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IfiledidChangeKeywordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IfiledidChangeKeywords represents a ChangeKeywords event raised by the Ifiledid contract.
type IfiledidChangeKeywords struct {
	MfileDid common.Hash
	Keywords []string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeKeywords is a free log retrieval operation binding the contract event 0x0d0ecbde8ff7f0cbfa8a5aba6209011f194a2ab60ceb48e047f704427ebbed2e.
//
// Solidity: event ChangeKeywords(string indexed mfileDid, string[] keywords)
func (_Ifiledid *IfiledidFilterer) FilterChangeKeywords(opts *bind.FilterOpts, mfileDid []string) (*IfiledidChangeKeywordsIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.FilterLogs(opts, "ChangeKeywords", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IfiledidChangeKeywordsIterator{contract: _Ifiledid.contract, event: "ChangeKeywords", logs: logs, sub: sub}, nil
}

// WatchChangeKeywords is a free log subscription operation binding the contract event 0x0d0ecbde8ff7f0cbfa8a5aba6209011f194a2ab60ceb48e047f704427ebbed2e.
//
// Solidity: event ChangeKeywords(string indexed mfileDid, string[] keywords)
func (_Ifiledid *IfiledidFilterer) WatchChangeKeywords(opts *bind.WatchOpts, sink chan<- *IfiledidChangeKeywords, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.WatchLogs(opts, "ChangeKeywords", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IfiledidChangeKeywords)
				if err := _Ifiledid.contract.UnpackLog(event, "ChangeKeywords", log); err != nil {
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

// ParseChangeKeywords is a log parse operation binding the contract event 0x0d0ecbde8ff7f0cbfa8a5aba6209011f194a2ab60ceb48e047f704427ebbed2e.
//
// Solidity: event ChangeKeywords(string indexed mfileDid, string[] keywords)
func (_Ifiledid *IfiledidFilterer) ParseChangeKeywords(log types.Log) (*IfiledidChangeKeywords, error) {
	event := new(IfiledidChangeKeywords)
	if err := _Ifiledid.contract.UnpackLog(event, "ChangeKeywords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IfiledidChangePriceIterator is returned from FilterChangePrice and is used to iterate over the raw logs and unpacked data for ChangePrice events raised by the Ifiledid contract.
type IfiledidChangePriceIterator struct {
	Event *IfiledidChangePrice // Event containing the contract specifics and raw log

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
func (it *IfiledidChangePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IfiledidChangePrice)
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
		it.Event = new(IfiledidChangePrice)
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
func (it *IfiledidChangePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IfiledidChangePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IfiledidChangePrice represents a ChangePrice event raised by the Ifiledid contract.
type IfiledidChangePrice struct {
	MfileDid common.Hash
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangePrice is a free log retrieval operation binding the contract event 0x0868ddc93d8bff1d2b8931574f1250319ab20dc6cdf2a7f86d5eee192ea1c986.
//
// Solidity: event ChangePrice(string indexed mfileDid, uint256 price)
func (_Ifiledid *IfiledidFilterer) FilterChangePrice(opts *bind.FilterOpts, mfileDid []string) (*IfiledidChangePriceIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.FilterLogs(opts, "ChangePrice", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IfiledidChangePriceIterator{contract: _Ifiledid.contract, event: "ChangePrice", logs: logs, sub: sub}, nil
}

// WatchChangePrice is a free log subscription operation binding the contract event 0x0868ddc93d8bff1d2b8931574f1250319ab20dc6cdf2a7f86d5eee192ea1c986.
//
// Solidity: event ChangePrice(string indexed mfileDid, uint256 price)
func (_Ifiledid *IfiledidFilterer) WatchChangePrice(opts *bind.WatchOpts, sink chan<- *IfiledidChangePrice, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.WatchLogs(opts, "ChangePrice", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IfiledidChangePrice)
				if err := _Ifiledid.contract.UnpackLog(event, "ChangePrice", log); err != nil {
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

// ParseChangePrice is a log parse operation binding the contract event 0x0868ddc93d8bff1d2b8931574f1250319ab20dc6cdf2a7f86d5eee192ea1c986.
//
// Solidity: event ChangePrice(string indexed mfileDid, uint256 price)
func (_Ifiledid *IfiledidFilterer) ParseChangePrice(log types.Log) (*IfiledidChangePrice, error) {
	event := new(IfiledidChangePrice)
	if err := _Ifiledid.contract.UnpackLog(event, "ChangePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IfiledidDeactivateMfileDidIterator is returned from FilterDeactivateMfileDid and is used to iterate over the raw logs and unpacked data for DeactivateMfileDid events raised by the Ifiledid contract.
type IfiledidDeactivateMfileDidIterator struct {
	Event *IfiledidDeactivateMfileDid // Event containing the contract specifics and raw log

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
func (it *IfiledidDeactivateMfileDidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IfiledidDeactivateMfileDid)
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
		it.Event = new(IfiledidDeactivateMfileDid)
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
func (it *IfiledidDeactivateMfileDidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IfiledidDeactivateMfileDidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IfiledidDeactivateMfileDid represents a DeactivateMfileDid event raised by the Ifiledid contract.
type IfiledidDeactivateMfileDid struct {
	MfileDid   string
	Deactivate bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeactivateMfileDid is a free log retrieval operation binding the contract event 0x5e17ae7a2d292f907502879bb1abf12db289923734c91a05b5d0651531aebf2a.
//
// Solidity: event DeactivateMfileDid(string mfileDid, bool deactivate)
func (_Ifiledid *IfiledidFilterer) FilterDeactivateMfileDid(opts *bind.FilterOpts) (*IfiledidDeactivateMfileDidIterator, error) {

	logs, sub, err := _Ifiledid.contract.FilterLogs(opts, "DeactivateMfileDid")
	if err != nil {
		return nil, err
	}
	return &IfiledidDeactivateMfileDidIterator{contract: _Ifiledid.contract, event: "DeactivateMfileDid", logs: logs, sub: sub}, nil
}

// WatchDeactivateMfileDid is a free log subscription operation binding the contract event 0x5e17ae7a2d292f907502879bb1abf12db289923734c91a05b5d0651531aebf2a.
//
// Solidity: event DeactivateMfileDid(string mfileDid, bool deactivate)
func (_Ifiledid *IfiledidFilterer) WatchDeactivateMfileDid(opts *bind.WatchOpts, sink chan<- *IfiledidDeactivateMfileDid) (event.Subscription, error) {

	logs, sub, err := _Ifiledid.contract.WatchLogs(opts, "DeactivateMfileDid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IfiledidDeactivateMfileDid)
				if err := _Ifiledid.contract.UnpackLog(event, "DeactivateMfileDid", log); err != nil {
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

// ParseDeactivateMfileDid is a log parse operation binding the contract event 0x5e17ae7a2d292f907502879bb1abf12db289923734c91a05b5d0651531aebf2a.
//
// Solidity: event DeactivateMfileDid(string mfileDid, bool deactivate)
func (_Ifiledid *IfiledidFilterer) ParseDeactivateMfileDid(log types.Log) (*IfiledidDeactivateMfileDid, error) {
	event := new(IfiledidDeactivateMfileDid)
	if err := _Ifiledid.contract.UnpackLog(event, "DeactivateMfileDid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IfiledidDeactivateReadIterator is returned from FilterDeactivateRead and is used to iterate over the raw logs and unpacked data for DeactivateRead events raised by the Ifiledid contract.
type IfiledidDeactivateReadIterator struct {
	Event *IfiledidDeactivateRead // Event containing the contract specifics and raw log

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
func (it *IfiledidDeactivateReadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IfiledidDeactivateRead)
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
		it.Event = new(IfiledidDeactivateRead)
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
func (it *IfiledidDeactivateReadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IfiledidDeactivateReadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IfiledidDeactivateRead represents a DeactivateRead event raised by the Ifiledid contract.
type IfiledidDeactivateRead struct {
	MfileDid common.Hash
	MemoDid  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeactivateRead is a free log retrieval operation binding the contract event 0x9d679cd38c09c06f29e83ac51c57df2602453a9336575359346cfab8a8f32606.
//
// Solidity: event DeactivateRead(string indexed mfileDid, string memoDid)
func (_Ifiledid *IfiledidFilterer) FilterDeactivateRead(opts *bind.FilterOpts, mfileDid []string) (*IfiledidDeactivateReadIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.FilterLogs(opts, "DeactivateRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IfiledidDeactivateReadIterator{contract: _Ifiledid.contract, event: "DeactivateRead", logs: logs, sub: sub}, nil
}

// WatchDeactivateRead is a free log subscription operation binding the contract event 0x9d679cd38c09c06f29e83ac51c57df2602453a9336575359346cfab8a8f32606.
//
// Solidity: event DeactivateRead(string indexed mfileDid, string memoDid)
func (_Ifiledid *IfiledidFilterer) WatchDeactivateRead(opts *bind.WatchOpts, sink chan<- *IfiledidDeactivateRead, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.WatchLogs(opts, "DeactivateRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IfiledidDeactivateRead)
				if err := _Ifiledid.contract.UnpackLog(event, "DeactivateRead", log); err != nil {
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

// ParseDeactivateRead is a log parse operation binding the contract event 0x9d679cd38c09c06f29e83ac51c57df2602453a9336575359346cfab8a8f32606.
//
// Solidity: event DeactivateRead(string indexed mfileDid, string memoDid)
func (_Ifiledid *IfiledidFilterer) ParseDeactivateRead(log types.Log) (*IfiledidDeactivateRead, error) {
	event := new(IfiledidDeactivateRead)
	if err := _Ifiledid.contract.UnpackLog(event, "DeactivateRead", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IfiledidGrantReadIterator is returned from FilterGrantRead and is used to iterate over the raw logs and unpacked data for GrantRead events raised by the Ifiledid contract.
type IfiledidGrantReadIterator struct {
	Event *IfiledidGrantRead // Event containing the contract specifics and raw log

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
func (it *IfiledidGrantReadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IfiledidGrantRead)
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
		it.Event = new(IfiledidGrantRead)
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
func (it *IfiledidGrantReadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IfiledidGrantReadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IfiledidGrantRead represents a GrantRead event raised by the Ifiledid contract.
type IfiledidGrantRead struct {
	MfileDid common.Hash
	MemoDid  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGrantRead is a free log retrieval operation binding the contract event 0x94b2915a0e10afcab72fdc47f046e2d9f8c45380746ef36ed2f10afaff5390be.
//
// Solidity: event GrantRead(string indexed mfileDid, string memoDid)
func (_Ifiledid *IfiledidFilterer) FilterGrantRead(opts *bind.FilterOpts, mfileDid []string) (*IfiledidGrantReadIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.FilterLogs(opts, "GrantRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IfiledidGrantReadIterator{contract: _Ifiledid.contract, event: "GrantRead", logs: logs, sub: sub}, nil
}

// WatchGrantRead is a free log subscription operation binding the contract event 0x94b2915a0e10afcab72fdc47f046e2d9f8c45380746ef36ed2f10afaff5390be.
//
// Solidity: event GrantRead(string indexed mfileDid, string memoDid)
func (_Ifiledid *IfiledidFilterer) WatchGrantRead(opts *bind.WatchOpts, sink chan<- *IfiledidGrantRead, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _Ifiledid.contract.WatchLogs(opts, "GrantRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IfiledidGrantRead)
				if err := _Ifiledid.contract.UnpackLog(event, "GrantRead", log); err != nil {
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

// ParseGrantRead is a log parse operation binding the contract event 0x94b2915a0e10afcab72fdc47f046e2d9f8c45380746ef36ed2f10afaff5390be.
//
// Solidity: event GrantRead(string indexed mfileDid, string memoDid)
func (_Ifiledid *IfiledidFilterer) ParseGrantRead(log types.Log) (*IfiledidGrantRead, error) {
	event := new(IfiledidGrantRead)
	if err := _Ifiledid.contract.UnpackLog(event, "GrantRead", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IfiledidRegisterMfileDidIterator is returned from FilterRegisterMfileDid and is used to iterate over the raw logs and unpacked data for RegisterMfileDid events raised by the Ifiledid contract.
type IfiledidRegisterMfileDidIterator struct {
	Event *IfiledidRegisterMfileDid // Event containing the contract specifics and raw log

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
func (it *IfiledidRegisterMfileDidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IfiledidRegisterMfileDid)
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
		it.Event = new(IfiledidRegisterMfileDid)
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
func (it *IfiledidRegisterMfileDidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IfiledidRegisterMfileDidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IfiledidRegisterMfileDid represents a RegisterMfileDid event raised by the Ifiledid contract.
type IfiledidRegisterMfileDid struct {
	MfileDid string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegisterMfileDid is a free log retrieval operation binding the contract event 0x09f50a6181226540ba96339e43e56ab9004a370db287d9dd7db64891edc494fe.
//
// Solidity: event RegisterMfileDid(string mfileDid)
func (_Ifiledid *IfiledidFilterer) FilterRegisterMfileDid(opts *bind.FilterOpts) (*IfiledidRegisterMfileDidIterator, error) {

	logs, sub, err := _Ifiledid.contract.FilterLogs(opts, "RegisterMfileDid")
	if err != nil {
		return nil, err
	}
	return &IfiledidRegisterMfileDidIterator{contract: _Ifiledid.contract, event: "RegisterMfileDid", logs: logs, sub: sub}, nil
}

// WatchRegisterMfileDid is a free log subscription operation binding the contract event 0x09f50a6181226540ba96339e43e56ab9004a370db287d9dd7db64891edc494fe.
//
// Solidity: event RegisterMfileDid(string mfileDid)
func (_Ifiledid *IfiledidFilterer) WatchRegisterMfileDid(opts *bind.WatchOpts, sink chan<- *IfiledidRegisterMfileDid) (event.Subscription, error) {

	logs, sub, err := _Ifiledid.contract.WatchLogs(opts, "RegisterMfileDid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IfiledidRegisterMfileDid)
				if err := _Ifiledid.contract.UnpackLog(event, "RegisterMfileDid", log); err != nil {
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

// ParseRegisterMfileDid is a log parse operation binding the contract event 0x09f50a6181226540ba96339e43e56ab9004a370db287d9dd7db64891edc494fe.
//
// Solidity: event RegisterMfileDid(string mfileDid)
func (_Ifiledid *IfiledidFilterer) ParseRegisterMfileDid(log types.Log) (*IfiledidRegisterMfileDid, error) {
	event := new(IfiledidRegisterMfileDid)
	if err := _Ifiledid.contract.UnpackLog(event, "RegisterMfileDid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
