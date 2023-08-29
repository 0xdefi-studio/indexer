// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bankroll

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

// BankrollMetaData contains all meta data concerning the Bankroll contract.
var BankrollMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidGameAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gameAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"BankRoll_Game_State_Changed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"Bankroll_Token_State_Changed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game\",\"type\":\"address\"}],\"name\":\"getIsGame\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getIsValidWager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"isPlayerSuspended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"game\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrapped\",\"type\":\"address\"}],\"name\":\"setWrappedAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"transferPayout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNativeFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BankrollABI is the input ABI used to generate the binding from.
// Deprecated: Use BankrollMetaData.ABI instead.
var BankrollABI = BankrollMetaData.ABI

// Bankroll is an auto generated Go binding around an Ethereum contract.
type Bankroll struct {
	BankrollCaller     // Read-only binding to the contract
	BankrollTransactor // Write-only binding to the contract
	BankrollFilterer   // Log filterer for contract events
}

// BankrollCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankrollCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankrollTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankrollTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankrollFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankrollFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankrollSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankrollSession struct {
	Contract     *Bankroll         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankrollCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankrollCallerSession struct {
	Contract *BankrollCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BankrollTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankrollTransactorSession struct {
	Contract     *BankrollTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BankrollRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankrollRaw struct {
	Contract *Bankroll // Generic contract binding to access the raw methods on
}

// BankrollCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankrollCallerRaw struct {
	Contract *BankrollCaller // Generic read-only contract binding to access the raw methods on
}

// BankrollTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankrollTransactorRaw struct {
	Contract *BankrollTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBankroll creates a new instance of Bankroll, bound to a specific deployed contract.
func NewBankroll(address common.Address, backend bind.ContractBackend) (*Bankroll, error) {
	contract, err := bindBankroll(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bankroll{BankrollCaller: BankrollCaller{contract: contract}, BankrollTransactor: BankrollTransactor{contract: contract}, BankrollFilterer: BankrollFilterer{contract: contract}}, nil
}

// NewBankrollCaller creates a new read-only instance of Bankroll, bound to a specific deployed contract.
func NewBankrollCaller(address common.Address, caller bind.ContractCaller) (*BankrollCaller, error) {
	contract, err := bindBankroll(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankrollCaller{contract: contract}, nil
}

// NewBankrollTransactor creates a new write-only instance of Bankroll, bound to a specific deployed contract.
func NewBankrollTransactor(address common.Address, transactor bind.ContractTransactor) (*BankrollTransactor, error) {
	contract, err := bindBankroll(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankrollTransactor{contract: contract}, nil
}

// NewBankrollFilterer creates a new log filterer instance of Bankroll, bound to a specific deployed contract.
func NewBankrollFilterer(address common.Address, filterer bind.ContractFilterer) (*BankrollFilterer, error) {
	contract, err := bindBankroll(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankrollFilterer{contract: contract}, nil
}

// bindBankroll binds a generic wrapper to an already deployed contract.
func bindBankroll(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankrollABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankroll *BankrollRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankroll.Contract.BankrollCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankroll *BankrollRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankroll.Contract.BankrollTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankroll *BankrollRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankroll.Contract.BankrollTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bankroll *BankrollCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bankroll.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bankroll *BankrollTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankroll.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bankroll *BankrollTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bankroll.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bankroll *BankrollCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bankroll *BankrollSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bankroll.Contract.DOMAINSEPARATOR(&_Bankroll.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bankroll *BankrollCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bankroll.Contract.DOMAINSEPARATOR(&_Bankroll.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bankroll *BankrollCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bankroll *BankrollSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bankroll.Contract.Allowance(&_Bankroll.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bankroll *BankrollCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bankroll.Contract.Allowance(&_Bankroll.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Bankroll *BankrollCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Bankroll *BankrollSession) Asset() (common.Address, error) {
	return _Bankroll.Contract.Asset(&_Bankroll.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_Bankroll *BankrollCallerSession) Asset() (common.Address, error) {
	return _Bankroll.Contract.Asset(&_Bankroll.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bankroll *BankrollCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bankroll *BankrollSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bankroll.Contract.BalanceOf(&_Bankroll.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bankroll *BankrollCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bankroll.Contract.BalanceOf(&_Bankroll.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Bankroll *BankrollCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Bankroll *BankrollSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.ConvertToAssets(&_Bankroll.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_Bankroll *BankrollCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.ConvertToAssets(&_Bankroll.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Bankroll *BankrollCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Bankroll *BankrollSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.ConvertToShares(&_Bankroll.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_Bankroll *BankrollCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.ConvertToShares(&_Bankroll.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bankroll *BankrollCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bankroll *BankrollSession) Decimals() (uint8, error) {
	return _Bankroll.Contract.Decimals(&_Bankroll.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bankroll *BankrollCallerSession) Decimals() (uint8, error) {
	return _Bankroll.Contract.Decimals(&_Bankroll.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bankroll *BankrollCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bankroll *BankrollSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Bankroll.Contract.Eip712Domain(&_Bankroll.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bankroll *BankrollCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Bankroll.Contract.Eip712Domain(&_Bankroll.CallOpts)
}

// GetIsGame is a free data retrieval call binding the contract method 0x298729d8.
//
// Solidity: function getIsGame(address game) view returns(bool)
func (_Bankroll *BankrollCaller) GetIsGame(opts *bind.CallOpts, game common.Address) (bool, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "getIsGame", game)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsGame is a free data retrieval call binding the contract method 0x298729d8.
//
// Solidity: function getIsGame(address game) view returns(bool)
func (_Bankroll *BankrollSession) GetIsGame(game common.Address) (bool, error) {
	return _Bankroll.Contract.GetIsGame(&_Bankroll.CallOpts, game)
}

// GetIsGame is a free data retrieval call binding the contract method 0x298729d8.
//
// Solidity: function getIsGame(address game) view returns(bool)
func (_Bankroll *BankrollCallerSession) GetIsGame(game common.Address) (bool, error) {
	return _Bankroll.Contract.GetIsGame(&_Bankroll.CallOpts, game)
}

// GetIsValidWager is a free data retrieval call binding the contract method 0xbfd7f9cc.
//
// Solidity: function getIsValidWager(address game, address tokenAddress) view returns(bool)
func (_Bankroll *BankrollCaller) GetIsValidWager(opts *bind.CallOpts, game common.Address, tokenAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "getIsValidWager", game, tokenAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsValidWager is a free data retrieval call binding the contract method 0xbfd7f9cc.
//
// Solidity: function getIsValidWager(address game, address tokenAddress) view returns(bool)
func (_Bankroll *BankrollSession) GetIsValidWager(game common.Address, tokenAddress common.Address) (bool, error) {
	return _Bankroll.Contract.GetIsValidWager(&_Bankroll.CallOpts, game, tokenAddress)
}

// GetIsValidWager is a free data retrieval call binding the contract method 0xbfd7f9cc.
//
// Solidity: function getIsValidWager(address game, address tokenAddress) view returns(bool)
func (_Bankroll *BankrollCallerSession) GetIsValidWager(game common.Address, tokenAddress common.Address) (bool, error) {
	return _Bankroll.Contract.GetIsValidWager(&_Bankroll.CallOpts, game, tokenAddress)
}

// GetWrappedAddress is a free data retrieval call binding the contract method 0xb097d80e.
//
// Solidity: function getWrappedAddress() view returns(address)
func (_Bankroll *BankrollCaller) GetWrappedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "getWrappedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedAddress is a free data retrieval call binding the contract method 0xb097d80e.
//
// Solidity: function getWrappedAddress() view returns(address)
func (_Bankroll *BankrollSession) GetWrappedAddress() (common.Address, error) {
	return _Bankroll.Contract.GetWrappedAddress(&_Bankroll.CallOpts)
}

// GetWrappedAddress is a free data retrieval call binding the contract method 0xb097d80e.
//
// Solidity: function getWrappedAddress() view returns(address)
func (_Bankroll *BankrollCallerSession) GetWrappedAddress() (common.Address, error) {
	return _Bankroll.Contract.GetWrappedAddress(&_Bankroll.CallOpts)
}

// IsPlayerSuspended is a free data retrieval call binding the contract method 0x0a5748a8.
//
// Solidity: function isPlayerSuspended(address player) view returns(bool, uint256)
func (_Bankroll *BankrollCaller) IsPlayerSuspended(opts *bind.CallOpts, player common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "isPlayerSuspended", player)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// IsPlayerSuspended is a free data retrieval call binding the contract method 0x0a5748a8.
//
// Solidity: function isPlayerSuspended(address player) view returns(bool, uint256)
func (_Bankroll *BankrollSession) IsPlayerSuspended(player common.Address) (bool, *big.Int, error) {
	return _Bankroll.Contract.IsPlayerSuspended(&_Bankroll.CallOpts, player)
}

// IsPlayerSuspended is a free data retrieval call binding the contract method 0x0a5748a8.
//
// Solidity: function isPlayerSuspended(address player) view returns(bool, uint256)
func (_Bankroll *BankrollCallerSession) IsPlayerSuspended(player common.Address) (bool, *big.Int, error) {
	return _Bankroll.Contract.IsPlayerSuspended(&_Bankroll.CallOpts, player)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Bankroll *BankrollCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Bankroll *BankrollSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Bankroll.Contract.MaxDeposit(&_Bankroll.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_Bankroll *BankrollCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _Bankroll.Contract.MaxDeposit(&_Bankroll.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Bankroll *BankrollCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Bankroll *BankrollSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Bankroll.Contract.MaxMint(&_Bankroll.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_Bankroll *BankrollCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _Bankroll.Contract.MaxMint(&_Bankroll.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Bankroll *BankrollCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Bankroll *BankrollSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Bankroll.Contract.MaxRedeem(&_Bankroll.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_Bankroll *BankrollCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _Bankroll.Contract.MaxRedeem(&_Bankroll.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Bankroll *BankrollCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Bankroll *BankrollSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Bankroll.Contract.MaxWithdraw(&_Bankroll.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_Bankroll *BankrollCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _Bankroll.Contract.MaxWithdraw(&_Bankroll.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bankroll *BankrollCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bankroll *BankrollSession) Name() (string, error) {
	return _Bankroll.Contract.Name(&_Bankroll.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bankroll *BankrollCallerSession) Name() (string, error) {
	return _Bankroll.Contract.Name(&_Bankroll.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bankroll *BankrollCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bankroll *BankrollSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Bankroll.Contract.Nonces(&_Bankroll.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bankroll *BankrollCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Bankroll.Contract.Nonces(&_Bankroll.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bankroll *BankrollCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bankroll *BankrollSession) Owner() (common.Address, error) {
	return _Bankroll.Contract.Owner(&_Bankroll.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bankroll *BankrollCallerSession) Owner() (common.Address, error) {
	return _Bankroll.Contract.Owner(&_Bankroll.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Bankroll *BankrollCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Bankroll *BankrollSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.PreviewDeposit(&_Bankroll.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_Bankroll *BankrollCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.PreviewDeposit(&_Bankroll.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Bankroll *BankrollCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Bankroll *BankrollSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.PreviewMint(&_Bankroll.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_Bankroll *BankrollCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.PreviewMint(&_Bankroll.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Bankroll *BankrollCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Bankroll *BankrollSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.PreviewRedeem(&_Bankroll.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_Bankroll *BankrollCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.PreviewRedeem(&_Bankroll.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Bankroll *BankrollCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Bankroll *BankrollSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.PreviewWithdraw(&_Bankroll.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_Bankroll *BankrollCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _Bankroll.Contract.PreviewWithdraw(&_Bankroll.CallOpts, assets)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bankroll *BankrollCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bankroll *BankrollSession) Symbol() (string, error) {
	return _Bankroll.Contract.Symbol(&_Bankroll.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bankroll *BankrollCallerSession) Symbol() (string, error) {
	return _Bankroll.Contract.Symbol(&_Bankroll.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Bankroll *BankrollCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Bankroll *BankrollSession) TotalAssets() (*big.Int, error) {
	return _Bankroll.Contract.TotalAssets(&_Bankroll.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Bankroll *BankrollCallerSession) TotalAssets() (*big.Int, error) {
	return _Bankroll.Contract.TotalAssets(&_Bankroll.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bankroll *BankrollCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bankroll.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bankroll *BankrollSession) TotalSupply() (*big.Int, error) {
	return _Bankroll.Contract.TotalSupply(&_Bankroll.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bankroll *BankrollCallerSession) TotalSupply() (*big.Int, error) {
	return _Bankroll.Contract.TotalSupply(&_Bankroll.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bankroll *BankrollTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bankroll *BankrollSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.Approve(&_Bankroll.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bankroll *BankrollTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.Approve(&_Bankroll.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bankroll *BankrollTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bankroll *BankrollSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.DecreaseAllowance(&_Bankroll.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bankroll *BankrollTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.DecreaseAllowance(&_Bankroll.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Bankroll *BankrollTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Bankroll *BankrollSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Deposit(&_Bankroll.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_Bankroll *BankrollTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Deposit(&_Bankroll.TransactOpts, assets, receiver)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bankroll *BankrollTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bankroll *BankrollSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.IncreaseAllowance(&_Bankroll.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bankroll *BankrollTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.IncreaseAllowance(&_Bankroll.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _asset) returns()
func (_Bankroll *BankrollTransactor) Initialize(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "initialize", _asset)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _asset) returns()
func (_Bankroll *BankrollSession) Initialize(_asset common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Initialize(&_Bankroll.TransactOpts, _asset)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _asset) returns()
func (_Bankroll *BankrollTransactorSession) Initialize(_asset common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Initialize(&_Bankroll.TransactOpts, _asset)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Bankroll *BankrollTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Bankroll *BankrollSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Mint(&_Bankroll.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_Bankroll *BankrollTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Mint(&_Bankroll.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bankroll *BankrollTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bankroll *BankrollSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bankroll.Contract.Permit(&_Bankroll.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bankroll *BankrollTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bankroll.Contract.Permit(&_Bankroll.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Bankroll *BankrollTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Bankroll *BankrollSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Redeem(&_Bankroll.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_Bankroll *BankrollTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Redeem(&_Bankroll.TransactOpts, shares, receiver, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bankroll *BankrollTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bankroll *BankrollSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bankroll.Contract.RenounceOwnership(&_Bankroll.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bankroll *BankrollTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bankroll.Contract.RenounceOwnership(&_Bankroll.TransactOpts)
}

// SetGame is a paid mutator transaction binding the contract method 0xb0836e2e.
//
// Solidity: function setGame(address game, bool isValid) returns()
func (_Bankroll *BankrollTransactor) SetGame(opts *bind.TransactOpts, game common.Address, isValid bool) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "setGame", game, isValid)
}

// SetGame is a paid mutator transaction binding the contract method 0xb0836e2e.
//
// Solidity: function setGame(address game, bool isValid) returns()
func (_Bankroll *BankrollSession) SetGame(game common.Address, isValid bool) (*types.Transaction, error) {
	return _Bankroll.Contract.SetGame(&_Bankroll.TransactOpts, game, isValid)
}

// SetGame is a paid mutator transaction binding the contract method 0xb0836e2e.
//
// Solidity: function setGame(address game, bool isValid) returns()
func (_Bankroll *BankrollTransactorSession) SetGame(game common.Address, isValid bool) (*types.Transaction, error) {
	return _Bankroll.Contract.SetGame(&_Bankroll.TransactOpts, game, isValid)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x2a0b5e9e.
//
// Solidity: function setTokenAddress(address tokenAddress, bool isValid) returns()
func (_Bankroll *BankrollTransactor) SetTokenAddress(opts *bind.TransactOpts, tokenAddress common.Address, isValid bool) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "setTokenAddress", tokenAddress, isValid)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x2a0b5e9e.
//
// Solidity: function setTokenAddress(address tokenAddress, bool isValid) returns()
func (_Bankroll *BankrollSession) SetTokenAddress(tokenAddress common.Address, isValid bool) (*types.Transaction, error) {
	return _Bankroll.Contract.SetTokenAddress(&_Bankroll.TransactOpts, tokenAddress, isValid)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x2a0b5e9e.
//
// Solidity: function setTokenAddress(address tokenAddress, bool isValid) returns()
func (_Bankroll *BankrollTransactorSession) SetTokenAddress(tokenAddress common.Address, isValid bool) (*types.Transaction, error) {
	return _Bankroll.Contract.SetTokenAddress(&_Bankroll.TransactOpts, tokenAddress, isValid)
}

// SetWrappedAddress is a paid mutator transaction binding the contract method 0x509cac54.
//
// Solidity: function setWrappedAddress(address wrapped) returns()
func (_Bankroll *BankrollTransactor) SetWrappedAddress(opts *bind.TransactOpts, wrapped common.Address) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "setWrappedAddress", wrapped)
}

// SetWrappedAddress is a paid mutator transaction binding the contract method 0x509cac54.
//
// Solidity: function setWrappedAddress(address wrapped) returns()
func (_Bankroll *BankrollSession) SetWrappedAddress(wrapped common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.SetWrappedAddress(&_Bankroll.TransactOpts, wrapped)
}

// SetWrappedAddress is a paid mutator transaction binding the contract method 0x509cac54.
//
// Solidity: function setWrappedAddress(address wrapped) returns()
func (_Bankroll *BankrollTransactorSession) SetWrappedAddress(wrapped common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.SetWrappedAddress(&_Bankroll.TransactOpts, wrapped)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bankroll *BankrollTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bankroll *BankrollSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.Transfer(&_Bankroll.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bankroll *BankrollTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.Transfer(&_Bankroll.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bankroll *BankrollTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bankroll *BankrollSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.TransferFrom(&_Bankroll.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bankroll *BankrollTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.TransferFrom(&_Bankroll.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bankroll *BankrollTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bankroll *BankrollSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.TransferOwnership(&_Bankroll.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bankroll *BankrollTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.TransferOwnership(&_Bankroll.TransactOpts, newOwner)
}

// TransferPayout is a paid mutator transaction binding the contract method 0x6c025ec2.
//
// Solidity: function transferPayout(address player, uint256 payout, address tokenAddress) returns()
func (_Bankroll *BankrollTransactor) TransferPayout(opts *bind.TransactOpts, player common.Address, payout *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "transferPayout", player, payout, tokenAddress)
}

// TransferPayout is a paid mutator transaction binding the contract method 0x6c025ec2.
//
// Solidity: function transferPayout(address player, uint256 payout, address tokenAddress) returns()
func (_Bankroll *BankrollSession) TransferPayout(player common.Address, payout *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.TransferPayout(&_Bankroll.TransactOpts, player, payout, tokenAddress)
}

// TransferPayout is a paid mutator transaction binding the contract method 0x6c025ec2.
//
// Solidity: function transferPayout(address player, uint256 payout, address tokenAddress) returns()
func (_Bankroll *BankrollTransactorSession) TransferPayout(player common.Address, payout *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.TransferPayout(&_Bankroll.TransactOpts, player, payout, tokenAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Bankroll *BankrollTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Bankroll *BankrollSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Withdraw(&_Bankroll.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_Bankroll *BankrollTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _Bankroll.Contract.Withdraw(&_Bankroll.TransactOpts, assets, receiver, owner)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x1c20fadd.
//
// Solidity: function withdrawFunds(address tokenAddress, address to, uint256 amount) returns()
func (_Bankroll *BankrollTransactor) WithdrawFunds(opts *bind.TransactOpts, tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "withdrawFunds", tokenAddress, to, amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x1c20fadd.
//
// Solidity: function withdrawFunds(address tokenAddress, address to, uint256 amount) returns()
func (_Bankroll *BankrollSession) WithdrawFunds(tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.WithdrawFunds(&_Bankroll.TransactOpts, tokenAddress, to, amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x1c20fadd.
//
// Solidity: function withdrawFunds(address tokenAddress, address to, uint256 amount) returns()
func (_Bankroll *BankrollTransactorSession) WithdrawFunds(tokenAddress common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.WithdrawFunds(&_Bankroll.TransactOpts, tokenAddress, to, amount)
}

// WithdrawNativeFunds is a paid mutator transaction binding the contract method 0x26792aca.
//
// Solidity: function withdrawNativeFunds(address to, uint256 amount) returns()
func (_Bankroll *BankrollTransactor) WithdrawNativeFunds(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.contract.Transact(opts, "withdrawNativeFunds", to, amount)
}

// WithdrawNativeFunds is a paid mutator transaction binding the contract method 0x26792aca.
//
// Solidity: function withdrawNativeFunds(address to, uint256 amount) returns()
func (_Bankroll *BankrollSession) WithdrawNativeFunds(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.WithdrawNativeFunds(&_Bankroll.TransactOpts, to, amount)
}

// WithdrawNativeFunds is a paid mutator transaction binding the contract method 0x26792aca.
//
// Solidity: function withdrawNativeFunds(address to, uint256 amount) returns()
func (_Bankroll *BankrollTransactorSession) WithdrawNativeFunds(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bankroll.Contract.WithdrawNativeFunds(&_Bankroll.TransactOpts, to, amount)
}

// BankrollApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bankroll contract.
type BankrollApprovalIterator struct {
	Event *BankrollApproval // Event containing the contract specifics and raw log

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
func (it *BankrollApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankrollApproval)
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
		it.Event = new(BankrollApproval)
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
func (it *BankrollApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankrollApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankrollApproval represents a Approval event raised by the Bankroll contract.
type BankrollApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bankroll *BankrollFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BankrollApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bankroll.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BankrollApprovalIterator{contract: _Bankroll.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bankroll *BankrollFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BankrollApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bankroll.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankrollApproval)
				if err := _Bankroll.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bankroll *BankrollFilterer) ParseApproval(log types.Log) (*BankrollApproval, error) {
	event := new(BankrollApproval)
	if err := _Bankroll.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankrollBankRollGameStateChangedIterator is returned from FilterBankRollGameStateChanged and is used to iterate over the raw logs and unpacked data for BankRollGameStateChanged events raised by the Bankroll contract.
type BankrollBankRollGameStateChangedIterator struct {
	Event *BankrollBankRollGameStateChanged // Event containing the contract specifics and raw log

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
func (it *BankrollBankRollGameStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankrollBankRollGameStateChanged)
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
		it.Event = new(BankrollBankRollGameStateChanged)
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
func (it *BankrollBankRollGameStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankrollBankRollGameStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankrollBankRollGameStateChanged represents a BankRollGameStateChanged event raised by the Bankroll contract.
type BankrollBankRollGameStateChanged struct {
	GameAddress common.Address
	IsValid     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBankRollGameStateChanged is a free log retrieval operation binding the contract event 0xdbd1aaf28b8f98e5d4e0c894cc0a5d00595709f1f072f52fa3fbab1a0802aa21.
//
// Solidity: event BankRoll_Game_State_Changed(address gameAddress, bool isValid)
func (_Bankroll *BankrollFilterer) FilterBankRollGameStateChanged(opts *bind.FilterOpts) (*BankrollBankRollGameStateChangedIterator, error) {

	logs, sub, err := _Bankroll.contract.FilterLogs(opts, "BankRoll_Game_State_Changed")
	if err != nil {
		return nil, err
	}
	return &BankrollBankRollGameStateChangedIterator{contract: _Bankroll.contract, event: "BankRoll_Game_State_Changed", logs: logs, sub: sub}, nil
}

// WatchBankRollGameStateChanged is a free log subscription operation binding the contract event 0xdbd1aaf28b8f98e5d4e0c894cc0a5d00595709f1f072f52fa3fbab1a0802aa21.
//
// Solidity: event BankRoll_Game_State_Changed(address gameAddress, bool isValid)
func (_Bankroll *BankrollFilterer) WatchBankRollGameStateChanged(opts *bind.WatchOpts, sink chan<- *BankrollBankRollGameStateChanged) (event.Subscription, error) {

	logs, sub, err := _Bankroll.contract.WatchLogs(opts, "BankRoll_Game_State_Changed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankrollBankRollGameStateChanged)
				if err := _Bankroll.contract.UnpackLog(event, "BankRoll_Game_State_Changed", log); err != nil {
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

// ParseBankRollGameStateChanged is a log parse operation binding the contract event 0xdbd1aaf28b8f98e5d4e0c894cc0a5d00595709f1f072f52fa3fbab1a0802aa21.
//
// Solidity: event BankRoll_Game_State_Changed(address gameAddress, bool isValid)
func (_Bankroll *BankrollFilterer) ParseBankRollGameStateChanged(log types.Log) (*BankrollBankRollGameStateChanged, error) {
	event := new(BankrollBankRollGameStateChanged)
	if err := _Bankroll.contract.UnpackLog(event, "BankRoll_Game_State_Changed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankrollBankrollTokenStateChangedIterator is returned from FilterBankrollTokenStateChanged and is used to iterate over the raw logs and unpacked data for BankrollTokenStateChanged events raised by the Bankroll contract.
type BankrollBankrollTokenStateChangedIterator struct {
	Event *BankrollBankrollTokenStateChanged // Event containing the contract specifics and raw log

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
func (it *BankrollBankrollTokenStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankrollBankrollTokenStateChanged)
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
		it.Event = new(BankrollBankrollTokenStateChanged)
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
func (it *BankrollBankrollTokenStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankrollBankrollTokenStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankrollBankrollTokenStateChanged represents a BankrollTokenStateChanged event raised by the Bankroll contract.
type BankrollBankrollTokenStateChanged struct {
	TokenAddress common.Address
	IsValid      bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBankrollTokenStateChanged is a free log retrieval operation binding the contract event 0xd371cc87ec694236589d1bdc930c8276a4335e3b49a288dc5ee58da7f094d7a8.
//
// Solidity: event Bankroll_Token_State_Changed(address tokenAddress, bool isValid)
func (_Bankroll *BankrollFilterer) FilterBankrollTokenStateChanged(opts *bind.FilterOpts) (*BankrollBankrollTokenStateChangedIterator, error) {

	logs, sub, err := _Bankroll.contract.FilterLogs(opts, "Bankroll_Token_State_Changed")
	if err != nil {
		return nil, err
	}
	return &BankrollBankrollTokenStateChangedIterator{contract: _Bankroll.contract, event: "Bankroll_Token_State_Changed", logs: logs, sub: sub}, nil
}

// WatchBankrollTokenStateChanged is a free log subscription operation binding the contract event 0xd371cc87ec694236589d1bdc930c8276a4335e3b49a288dc5ee58da7f094d7a8.
//
// Solidity: event Bankroll_Token_State_Changed(address tokenAddress, bool isValid)
func (_Bankroll *BankrollFilterer) WatchBankrollTokenStateChanged(opts *bind.WatchOpts, sink chan<- *BankrollBankrollTokenStateChanged) (event.Subscription, error) {

	logs, sub, err := _Bankroll.contract.WatchLogs(opts, "Bankroll_Token_State_Changed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankrollBankrollTokenStateChanged)
				if err := _Bankroll.contract.UnpackLog(event, "Bankroll_Token_State_Changed", log); err != nil {
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

// ParseBankrollTokenStateChanged is a log parse operation binding the contract event 0xd371cc87ec694236589d1bdc930c8276a4335e3b49a288dc5ee58da7f094d7a8.
//
// Solidity: event Bankroll_Token_State_Changed(address tokenAddress, bool isValid)
func (_Bankroll *BankrollFilterer) ParseBankrollTokenStateChanged(log types.Log) (*BankrollBankrollTokenStateChanged, error) {
	event := new(BankrollBankrollTokenStateChanged)
	if err := _Bankroll.contract.UnpackLog(event, "Bankroll_Token_State_Changed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankrollDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bankroll contract.
type BankrollDepositIterator struct {
	Event *BankrollDeposit // Event containing the contract specifics and raw log

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
func (it *BankrollDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankrollDeposit)
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
		it.Event = new(BankrollDeposit)
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
func (it *BankrollDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankrollDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankrollDeposit represents a Deposit event raised by the Bankroll contract.
type BankrollDeposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_Bankroll *BankrollFilterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*BankrollDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bankroll.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BankrollDepositIterator{contract: _Bankroll.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_Bankroll *BankrollFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BankrollDeposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bankroll.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankrollDeposit)
				if err := _Bankroll.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_Bankroll *BankrollFilterer) ParseDeposit(log types.Log) (*BankrollDeposit, error) {
	event := new(BankrollDeposit)
	if err := _Bankroll.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankrollEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Bankroll contract.
type BankrollEIP712DomainChangedIterator struct {
	Event *BankrollEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BankrollEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankrollEIP712DomainChanged)
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
		it.Event = new(BankrollEIP712DomainChanged)
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
func (it *BankrollEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankrollEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankrollEIP712DomainChanged represents a EIP712DomainChanged event raised by the Bankroll contract.
type BankrollEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bankroll *BankrollFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BankrollEIP712DomainChangedIterator, error) {

	logs, sub, err := _Bankroll.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BankrollEIP712DomainChangedIterator{contract: _Bankroll.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bankroll *BankrollFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BankrollEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Bankroll.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankrollEIP712DomainChanged)
				if err := _Bankroll.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bankroll *BankrollFilterer) ParseEIP712DomainChanged(log types.Log) (*BankrollEIP712DomainChanged, error) {
	event := new(BankrollEIP712DomainChanged)
	if err := _Bankroll.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankrollInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bankroll contract.
type BankrollInitializedIterator struct {
	Event *BankrollInitialized // Event containing the contract specifics and raw log

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
func (it *BankrollInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankrollInitialized)
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
		it.Event = new(BankrollInitialized)
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
func (it *BankrollInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankrollInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankrollInitialized represents a Initialized event raised by the Bankroll contract.
type BankrollInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bankroll *BankrollFilterer) FilterInitialized(opts *bind.FilterOpts) (*BankrollInitializedIterator, error) {

	logs, sub, err := _Bankroll.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BankrollInitializedIterator{contract: _Bankroll.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bankroll *BankrollFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BankrollInitialized) (event.Subscription, error) {

	logs, sub, err := _Bankroll.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankrollInitialized)
				if err := _Bankroll.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Bankroll *BankrollFilterer) ParseInitialized(log types.Log) (*BankrollInitialized, error) {
	event := new(BankrollInitialized)
	if err := _Bankroll.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankrollOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bankroll contract.
type BankrollOwnershipTransferredIterator struct {
	Event *BankrollOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BankrollOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankrollOwnershipTransferred)
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
		it.Event = new(BankrollOwnershipTransferred)
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
func (it *BankrollOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankrollOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankrollOwnershipTransferred represents a OwnershipTransferred event raised by the Bankroll contract.
type BankrollOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bankroll *BankrollFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BankrollOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bankroll.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BankrollOwnershipTransferredIterator{contract: _Bankroll.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bankroll *BankrollFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BankrollOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bankroll.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankrollOwnershipTransferred)
				if err := _Bankroll.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bankroll *BankrollFilterer) ParseOwnershipTransferred(log types.Log) (*BankrollOwnershipTransferred, error) {
	event := new(BankrollOwnershipTransferred)
	if err := _Bankroll.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankrollTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bankroll contract.
type BankrollTransferIterator struct {
	Event *BankrollTransfer // Event containing the contract specifics and raw log

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
func (it *BankrollTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankrollTransfer)
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
		it.Event = new(BankrollTransfer)
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
func (it *BankrollTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankrollTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankrollTransfer represents a Transfer event raised by the Bankroll contract.
type BankrollTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bankroll *BankrollFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BankrollTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bankroll.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BankrollTransferIterator{contract: _Bankroll.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bankroll *BankrollFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BankrollTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bankroll.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankrollTransfer)
				if err := _Bankroll.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bankroll *BankrollFilterer) ParseTransfer(log types.Log) (*BankrollTransfer, error) {
	event := new(BankrollTransfer)
	if err := _Bankroll.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankrollWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Bankroll contract.
type BankrollWithdrawIterator struct {
	Event *BankrollWithdraw // Event containing the contract specifics and raw log

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
func (it *BankrollWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankrollWithdraw)
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
		it.Event = new(BankrollWithdraw)
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
func (it *BankrollWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankrollWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankrollWithdraw represents a Withdraw event raised by the Bankroll contract.
type BankrollWithdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Bankroll *BankrollFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*BankrollWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bankroll.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BankrollWithdrawIterator{contract: _Bankroll.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Bankroll *BankrollFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BankrollWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Bankroll.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankrollWithdraw)
				if err := _Bankroll.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_Bankroll *BankrollFilterer) ParseWithdraw(log types.Log) (*BankrollWithdraw, error) {
	event := new(BankrollWithdraw)
	if err := _Bankroll.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
