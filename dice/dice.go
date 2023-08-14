// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dice

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

// DiceDiceGame is an auto generated low-level Go binding around an user-defined struct.
type DiceDiceGame struct {
	Wager        *big.Int
	StopGain     *big.Int
	StopLoss     *big.Int
	RequestID    *big.Int
	TokenAddress common.Address
	BlockNumber  uint64
	NumBets      uint32
	Multiplier   uint32
	IsOver       bool
}

// DiceMetaData contains all meta data concerning the Dice contract.
var DiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bankroll\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vrf\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"AwaitingVRF\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"want\",\"type\":\"uint256\"}],\"name\":\"BlockNumberTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"}],\"name\":\"InvalidMultiplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxNumBets\",\"type\":\"uint256\"}],\"name\":\"InvalidNumBets\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotApprovedBankroll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAwaitingVRF\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyRandomizerCanFulfill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"suspensionTime\",\"type\":\"uint256\"}],\"name\":\"PlayerSuspended\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefundFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wager\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxWager\",\"type\":\"uint256\"}],\"name\":\"WagerAboveLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroWager\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wager\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"diceOutcomes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"payouts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numGames\",\"type\":\"uint32\"}],\"name\":\"Dice_Outcome_Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"Dice_Play_Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wager\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"Dice_Refund_Event\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Bankroll\",\"outputs\":[{\"internalType\":\"contractIBankRoll\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"}],\"name\":\"Dice_GetState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"wager\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stopGain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stopLoss\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"numBets\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"multiplier\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isOver\",\"type\":\"bool\"}],\"internalType\":\"structDice.DiceGame\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wager\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"multiplier\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isOver\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"numBets\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"stopGain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stopLoss\",\"type\":\"uint256\"}],\"name\":\"Dice_Play\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Dice_Refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVRFFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomizer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_value\",\"type\":\"bytes32\"}],\"name\":\"randomizerCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"setVRFGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"whithdrawVRF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DiceABI is the input ABI used to generate the binding from.
// Deprecated: Use DiceMetaData.ABI instead.
var DiceABI = DiceMetaData.ABI

// Dice is an auto generated Go binding around an Ethereum contract.
type Dice struct {
	DiceCaller     // Read-only binding to the contract
	DiceTransactor // Write-only binding to the contract
	DiceFilterer   // Log filterer for contract events
}

// DiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiceSession struct {
	Contract     *Dice             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiceCallerSession struct {
	Contract *DiceCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiceTransactorSession struct {
	Contract     *DiceTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiceRaw struct {
	Contract *Dice // Generic contract binding to access the raw methods on
}

// DiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiceCallerRaw struct {
	Contract *DiceCaller // Generic read-only contract binding to access the raw methods on
}

// DiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiceTransactorRaw struct {
	Contract *DiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDice creates a new instance of Dice, bound to a specific deployed contract.
func NewDice(address common.Address, backend bind.ContractBackend) (*Dice, error) {
	contract, err := bindDice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dice{DiceCaller: DiceCaller{contract: contract}, DiceTransactor: DiceTransactor{contract: contract}, DiceFilterer: DiceFilterer{contract: contract}}, nil
}

// NewDiceCaller creates a new read-only instance of Dice, bound to a specific deployed contract.
func NewDiceCaller(address common.Address, caller bind.ContractCaller) (*DiceCaller, error) {
	contract, err := bindDice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiceCaller{contract: contract}, nil
}

// NewDiceTransactor creates a new write-only instance of Dice, bound to a specific deployed contract.
func NewDiceTransactor(address common.Address, transactor bind.ContractTransactor) (*DiceTransactor, error) {
	contract, err := bindDice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiceTransactor{contract: contract}, nil
}

// NewDiceFilterer creates a new log filterer instance of Dice, bound to a specific deployed contract.
func NewDiceFilterer(address common.Address, filterer bind.ContractFilterer) (*DiceFilterer, error) {
	contract, err := bindDice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiceFilterer{contract: contract}, nil
}

// bindDice binds a generic wrapper to an already deployed contract.
func bindDice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dice *DiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dice.Contract.DiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dice *DiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dice.Contract.DiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dice *DiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dice.Contract.DiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dice *DiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dice *DiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dice *DiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dice.Contract.contract.Transact(opts, method, params...)
}

// Bankroll is a free data retrieval call binding the contract method 0xd0bcfc0c.
//
// Solidity: function Bankroll() view returns(address)
func (_Dice *DiceCaller) Bankroll(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dice.contract.Call(opts, &out, "Bankroll")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bankroll is a free data retrieval call binding the contract method 0xd0bcfc0c.
//
// Solidity: function Bankroll() view returns(address)
func (_Dice *DiceSession) Bankroll() (common.Address, error) {
	return _Dice.Contract.Bankroll(&_Dice.CallOpts)
}

// Bankroll is a free data retrieval call binding the contract method 0xd0bcfc0c.
//
// Solidity: function Bankroll() view returns(address)
func (_Dice *DiceCallerSession) Bankroll() (common.Address, error) {
	return _Dice.Contract.Bankroll(&_Dice.CallOpts)
}

// DiceGetState is a free data retrieval call binding the contract method 0x08170595.
//
// Solidity: function Dice_GetState(address player) view returns((uint256,uint256,uint256,uint256,address,uint64,uint32,uint32,bool))
func (_Dice *DiceCaller) DiceGetState(opts *bind.CallOpts, player common.Address) (DiceDiceGame, error) {
	var out []interface{}
	err := _Dice.contract.Call(opts, &out, "Dice_GetState", player)

	if err != nil {
		return *new(DiceDiceGame), err
	}

	out0 := *abi.ConvertType(out[0], new(DiceDiceGame)).(*DiceDiceGame)

	return out0, err

}

// DiceGetState is a free data retrieval call binding the contract method 0x08170595.
//
// Solidity: function Dice_GetState(address player) view returns((uint256,uint256,uint256,uint256,address,uint64,uint32,uint32,bool))
func (_Dice *DiceSession) DiceGetState(player common.Address) (DiceDiceGame, error) {
	return _Dice.Contract.DiceGetState(&_Dice.CallOpts, player)
}

// DiceGetState is a free data retrieval call binding the contract method 0x08170595.
//
// Solidity: function Dice_GetState(address player) view returns((uint256,uint256,uint256,uint256,address,uint64,uint32,uint32,bool))
func (_Dice *DiceCallerSession) DiceGetState(player common.Address) (DiceDiceGame, error) {
	return _Dice.Contract.DiceGetState(&_Dice.CallOpts, player)
}

// GetVRFFee is a free data retrieval call binding the contract method 0xf0034a7e.
//
// Solidity: function getVRFFee() view returns(uint256 fee)
func (_Dice *DiceCaller) GetVRFFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dice.contract.Call(opts, &out, "getVRFFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVRFFee is a free data retrieval call binding the contract method 0xf0034a7e.
//
// Solidity: function getVRFFee() view returns(uint256 fee)
func (_Dice *DiceSession) GetVRFFee() (*big.Int, error) {
	return _Dice.Contract.GetVRFFee(&_Dice.CallOpts)
}

// GetVRFFee is a free data retrieval call binding the contract method 0xf0034a7e.
//
// Solidity: function getVRFFee() view returns(uint256 fee)
func (_Dice *DiceCallerSession) GetVRFFee() (*big.Int, error) {
	return _Dice.Contract.GetVRFFee(&_Dice.CallOpts)
}

// Randomizer is a free data retrieval call binding the contract method 0xf10fb584.
//
// Solidity: function randomizer() view returns(address)
func (_Dice *DiceCaller) Randomizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dice.contract.Call(opts, &out, "randomizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Randomizer is a free data retrieval call binding the contract method 0xf10fb584.
//
// Solidity: function randomizer() view returns(address)
func (_Dice *DiceSession) Randomizer() (common.Address, error) {
	return _Dice.Contract.Randomizer(&_Dice.CallOpts)
}

// Randomizer is a free data retrieval call binding the contract method 0xf10fb584.
//
// Solidity: function randomizer() view returns(address)
func (_Dice *DiceCallerSession) Randomizer() (common.Address, error) {
	return _Dice.Contract.Randomizer(&_Dice.CallOpts)
}

// DicePlay is a paid mutator transaction binding the contract method 0x74af2e59.
//
// Solidity: function Dice_Play(uint256 wager, uint32 multiplier, address tokenAddress, bool isOver, uint32 numBets, uint256 stopGain, uint256 stopLoss) payable returns()
func (_Dice *DiceTransactor) DicePlay(opts *bind.TransactOpts, wager *big.Int, multiplier uint32, tokenAddress common.Address, isOver bool, numBets uint32, stopGain *big.Int, stopLoss *big.Int) (*types.Transaction, error) {
	return _Dice.contract.Transact(opts, "Dice_Play", wager, multiplier, tokenAddress, isOver, numBets, stopGain, stopLoss)
}

// DicePlay is a paid mutator transaction binding the contract method 0x74af2e59.
//
// Solidity: function Dice_Play(uint256 wager, uint32 multiplier, address tokenAddress, bool isOver, uint32 numBets, uint256 stopGain, uint256 stopLoss) payable returns()
func (_Dice *DiceSession) DicePlay(wager *big.Int, multiplier uint32, tokenAddress common.Address, isOver bool, numBets uint32, stopGain *big.Int, stopLoss *big.Int) (*types.Transaction, error) {
	return _Dice.Contract.DicePlay(&_Dice.TransactOpts, wager, multiplier, tokenAddress, isOver, numBets, stopGain, stopLoss)
}

// DicePlay is a paid mutator transaction binding the contract method 0x74af2e59.
//
// Solidity: function Dice_Play(uint256 wager, uint32 multiplier, address tokenAddress, bool isOver, uint32 numBets, uint256 stopGain, uint256 stopLoss) payable returns()
func (_Dice *DiceTransactorSession) DicePlay(wager *big.Int, multiplier uint32, tokenAddress common.Address, isOver bool, numBets uint32, stopGain *big.Int, stopLoss *big.Int) (*types.Transaction, error) {
	return _Dice.Contract.DicePlay(&_Dice.TransactOpts, wager, multiplier, tokenAddress, isOver, numBets, stopGain, stopLoss)
}

// DiceRefund is a paid mutator transaction binding the contract method 0x24f50d66.
//
// Solidity: function Dice_Refund() returns()
func (_Dice *DiceTransactor) DiceRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dice.contract.Transact(opts, "Dice_Refund")
}

// DiceRefund is a paid mutator transaction binding the contract method 0x24f50d66.
//
// Solidity: function Dice_Refund() returns()
func (_Dice *DiceSession) DiceRefund() (*types.Transaction, error) {
	return _Dice.Contract.DiceRefund(&_Dice.TransactOpts)
}

// DiceRefund is a paid mutator transaction binding the contract method 0x24f50d66.
//
// Solidity: function Dice_Refund() returns()
func (_Dice *DiceTransactorSession) DiceRefund() (*types.Transaction, error) {
	return _Dice.Contract.DiceRefund(&_Dice.TransactOpts)
}

// RandomizerCallback is a paid mutator transaction binding the contract method 0xebe93caf.
//
// Solidity: function randomizerCallback(uint256 _id, bytes32 _value) returns()
func (_Dice *DiceTransactor) RandomizerCallback(opts *bind.TransactOpts, _id *big.Int, _value [32]byte) (*types.Transaction, error) {
	return _Dice.contract.Transact(opts, "randomizerCallback", _id, _value)
}

// RandomizerCallback is a paid mutator transaction binding the contract method 0xebe93caf.
//
// Solidity: function randomizerCallback(uint256 _id, bytes32 _value) returns()
func (_Dice *DiceSession) RandomizerCallback(_id *big.Int, _value [32]byte) (*types.Transaction, error) {
	return _Dice.Contract.RandomizerCallback(&_Dice.TransactOpts, _id, _value)
}

// RandomizerCallback is a paid mutator transaction binding the contract method 0xebe93caf.
//
// Solidity: function randomizerCallback(uint256 _id, bytes32 _value) returns()
func (_Dice *DiceTransactorSession) RandomizerCallback(_id *big.Int, _value [32]byte) (*types.Transaction, error) {
	return _Dice.Contract.RandomizerCallback(&_Dice.TransactOpts, _id, _value)
}

// SetVRFGasLimit is a paid mutator transaction binding the contract method 0xd6e23e6c.
//
// Solidity: function setVRFGasLimit(uint256 limit) returns()
func (_Dice *DiceTransactor) SetVRFGasLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _Dice.contract.Transact(opts, "setVRFGasLimit", limit)
}

// SetVRFGasLimit is a paid mutator transaction binding the contract method 0xd6e23e6c.
//
// Solidity: function setVRFGasLimit(uint256 limit) returns()
func (_Dice *DiceSession) SetVRFGasLimit(limit *big.Int) (*types.Transaction, error) {
	return _Dice.Contract.SetVRFGasLimit(&_Dice.TransactOpts, limit)
}

// SetVRFGasLimit is a paid mutator transaction binding the contract method 0xd6e23e6c.
//
// Solidity: function setVRFGasLimit(uint256 limit) returns()
func (_Dice *DiceTransactorSession) SetVRFGasLimit(limit *big.Int) (*types.Transaction, error) {
	return _Dice.Contract.SetVRFGasLimit(&_Dice.TransactOpts, limit)
}

// WhithdrawVRF is a paid mutator transaction binding the contract method 0x745b303e.
//
// Solidity: function whithdrawVRF(address to, uint256 amount) returns()
func (_Dice *DiceTransactor) WhithdrawVRF(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dice.contract.Transact(opts, "whithdrawVRF", to, amount)
}

// WhithdrawVRF is a paid mutator transaction binding the contract method 0x745b303e.
//
// Solidity: function whithdrawVRF(address to, uint256 amount) returns()
func (_Dice *DiceSession) WhithdrawVRF(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dice.Contract.WhithdrawVRF(&_Dice.TransactOpts, to, amount)
}

// WhithdrawVRF is a paid mutator transaction binding the contract method 0x745b303e.
//
// Solidity: function whithdrawVRF(address to, uint256 amount) returns()
func (_Dice *DiceTransactorSession) WhithdrawVRF(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Dice.Contract.WhithdrawVRF(&_Dice.TransactOpts, to, amount)
}

// DiceDiceOutcomeEventIterator is returned from FilterDiceOutcomeEvent and is used to iterate over the raw logs and unpacked data for DiceOutcomeEvent events raised by the Dice contract.
type DiceDiceOutcomeEventIterator struct {
	Event *DiceDiceOutcomeEvent // Event containing the contract specifics and raw log

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
func (it *DiceDiceOutcomeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiceDiceOutcomeEvent)
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
		it.Event = new(DiceDiceOutcomeEvent)
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
func (it *DiceDiceOutcomeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiceDiceOutcomeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiceDiceOutcomeEvent represents a DiceOutcomeEvent event raised by the Dice contract.
type DiceDiceOutcomeEvent struct {
	PlayerAddress common.Address
	Wager         *big.Int
	Payout        *big.Int
	TokenAddress  common.Address
	DiceOutcomes  []*big.Int
	Payouts       []*big.Int
	NumGames      uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDiceOutcomeEvent is a free log retrieval operation binding the contract event 0x090dbd65630d04a5178ecb9346e0cdcc299215a135c6eb7ecd530ce00dfa44d2.
//
// Solidity: event Dice_Outcome_Event(address indexed playerAddress, uint256 wager, uint256 payout, address tokenAddress, uint256[] diceOutcomes, uint256[] payouts, uint32 numGames)
func (_Dice *DiceFilterer) FilterDiceOutcomeEvent(opts *bind.FilterOpts, playerAddress []common.Address) (*DiceDiceOutcomeEventIterator, error) {

	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}

	logs, sub, err := _Dice.contract.FilterLogs(opts, "Dice_Outcome_Event", playerAddressRule)
	if err != nil {
		return nil, err
	}
	return &DiceDiceOutcomeEventIterator{contract: _Dice.contract, event: "Dice_Outcome_Event", logs: logs, sub: sub}, nil
}

// WatchDiceOutcomeEvent is a free log subscription operation binding the contract event 0x090dbd65630d04a5178ecb9346e0cdcc299215a135c6eb7ecd530ce00dfa44d2.
//
// Solidity: event Dice_Outcome_Event(address indexed playerAddress, uint256 wager, uint256 payout, address tokenAddress, uint256[] diceOutcomes, uint256[] payouts, uint32 numGames)
func (_Dice *DiceFilterer) WatchDiceOutcomeEvent(opts *bind.WatchOpts, sink chan<- *DiceDiceOutcomeEvent, playerAddress []common.Address) (event.Subscription, error) {

	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}

	logs, sub, err := _Dice.contract.WatchLogs(opts, "Dice_Outcome_Event", playerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiceDiceOutcomeEvent)
				if err := _Dice.contract.UnpackLog(event, "Dice_Outcome_Event", log); err != nil {
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

// ParseDiceOutcomeEvent is a log parse operation binding the contract event 0x090dbd65630d04a5178ecb9346e0cdcc299215a135c6eb7ecd530ce00dfa44d2.
//
// Solidity: event Dice_Outcome_Event(address indexed playerAddress, uint256 wager, uint256 payout, address tokenAddress, uint256[] diceOutcomes, uint256[] payouts, uint32 numGames)
func (_Dice *DiceFilterer) ParseDiceOutcomeEvent(log types.Log) (*DiceDiceOutcomeEvent, error) {
	event := new(DiceDiceOutcomeEvent)
	if err := _Dice.contract.UnpackLog(event, "Dice_Outcome_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiceDicePlayEventIterator is returned from FilterDicePlayEvent and is used to iterate over the raw logs and unpacked data for DicePlayEvent events raised by the Dice contract.
type DiceDicePlayEventIterator struct {
	Event *DiceDicePlayEvent // Event containing the contract specifics and raw log

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
func (it *DiceDicePlayEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiceDicePlayEvent)
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
		it.Event = new(DiceDicePlayEvent)
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
func (it *DiceDicePlayEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiceDicePlayEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiceDicePlayEvent represents a DicePlayEvent event raised by the Dice contract.
type DiceDicePlayEvent struct {
	PlayerAddress common.Address
	RequestId     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDicePlayEvent is a free log retrieval operation binding the contract event 0x46c1ffae04f53c49213e4ec73b7c175b1c4bad68b56ca6366dcac8f8491068be.
//
// Solidity: event Dice_Play_Event(address indexed playerAddress, uint256 requestId)
func (_Dice *DiceFilterer) FilterDicePlayEvent(opts *bind.FilterOpts, playerAddress []common.Address) (*DiceDicePlayEventIterator, error) {

	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}

	logs, sub, err := _Dice.contract.FilterLogs(opts, "Dice_Play_Event", playerAddressRule)
	if err != nil {
		return nil, err
	}
	return &DiceDicePlayEventIterator{contract: _Dice.contract, event: "Dice_Play_Event", logs: logs, sub: sub}, nil
}

// WatchDicePlayEvent is a free log subscription operation binding the contract event 0x46c1ffae04f53c49213e4ec73b7c175b1c4bad68b56ca6366dcac8f8491068be.
//
// Solidity: event Dice_Play_Event(address indexed playerAddress, uint256 requestId)
func (_Dice *DiceFilterer) WatchDicePlayEvent(opts *bind.WatchOpts, sink chan<- *DiceDicePlayEvent, playerAddress []common.Address) (event.Subscription, error) {

	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}

	logs, sub, err := _Dice.contract.WatchLogs(opts, "Dice_Play_Event", playerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiceDicePlayEvent)
				if err := _Dice.contract.UnpackLog(event, "Dice_Play_Event", log); err != nil {
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

// ParseDicePlayEvent is a log parse operation binding the contract event 0x46c1ffae04f53c49213e4ec73b7c175b1c4bad68b56ca6366dcac8f8491068be.
//
// Solidity: event Dice_Play_Event(address indexed playerAddress, uint256 requestId)
func (_Dice *DiceFilterer) ParseDicePlayEvent(log types.Log) (*DiceDicePlayEvent, error) {
	event := new(DiceDicePlayEvent)
	if err := _Dice.contract.UnpackLog(event, "Dice_Play_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiceDiceRefundEventIterator is returned from FilterDiceRefundEvent and is used to iterate over the raw logs and unpacked data for DiceRefundEvent events raised by the Dice contract.
type DiceDiceRefundEventIterator struct {
	Event *DiceDiceRefundEvent // Event containing the contract specifics and raw log

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
func (it *DiceDiceRefundEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiceDiceRefundEvent)
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
		it.Event = new(DiceDiceRefundEvent)
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
func (it *DiceDiceRefundEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiceDiceRefundEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiceDiceRefundEvent represents a DiceRefundEvent event raised by the Dice contract.
type DiceDiceRefundEvent struct {
	Player       common.Address
	Wager        *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDiceRefundEvent is a free log retrieval operation binding the contract event 0xac278bd0c9dff1c7083bac39d01b59001e74bdff34e5c7e75b5362162090cf2a.
//
// Solidity: event Dice_Refund_Event(address indexed player, uint256 wager, address tokenAddress)
func (_Dice *DiceFilterer) FilterDiceRefundEvent(opts *bind.FilterOpts, player []common.Address) (*DiceDiceRefundEventIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _Dice.contract.FilterLogs(opts, "Dice_Refund_Event", playerRule)
	if err != nil {
		return nil, err
	}
	return &DiceDiceRefundEventIterator{contract: _Dice.contract, event: "Dice_Refund_Event", logs: logs, sub: sub}, nil
}

// WatchDiceRefundEvent is a free log subscription operation binding the contract event 0xac278bd0c9dff1c7083bac39d01b59001e74bdff34e5c7e75b5362162090cf2a.
//
// Solidity: event Dice_Refund_Event(address indexed player, uint256 wager, address tokenAddress)
func (_Dice *DiceFilterer) WatchDiceRefundEvent(opts *bind.WatchOpts, sink chan<- *DiceDiceRefundEvent, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _Dice.contract.WatchLogs(opts, "Dice_Refund_Event", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiceDiceRefundEvent)
				if err := _Dice.contract.UnpackLog(event, "Dice_Refund_Event", log); err != nil {
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

// ParseDiceRefundEvent is a log parse operation binding the contract event 0xac278bd0c9dff1c7083bac39d01b59001e74bdff34e5c7e75b5362162090cf2a.
//
// Solidity: event Dice_Refund_Event(address indexed player, uint256 wager, address tokenAddress)
func (_Dice *DiceFilterer) ParseDiceRefundEvent(log types.Log) (*DiceDiceRefundEvent, error) {
	event := new(DiceDiceRefundEvent)
	if err := _Dice.contract.UnpackLog(event, "Dice_Refund_Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
