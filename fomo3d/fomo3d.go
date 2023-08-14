// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fomo3d

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

// Fomo3dMetaData contains all meta data concerning the Fomo3d contract.
var Fomo3dMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"playerBookAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasuryAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"buyerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onAffiliatePayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onBuyAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keysBought\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"genAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"potAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"airDropPot\",\"type\":\"uint256\"}],\"name\":\"onEndTx\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNewPlayer\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"affiliateID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliateAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"affiliateName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onNewName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onReLoadAndDistribute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"playerID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"playerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"playerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compressedData\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compressedIDs\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winnerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"winnerName\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWon\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"P3DAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"genAmount\",\"type\":\"uint256\"}],\"name\":\"onWithdrawAndDistribute\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activated_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"airDropPot_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"airDropTracker_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_affCode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_team\",\"type\":\"uint256\"}],\"name\":\"buyXid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eth\",\"type\":\"uint256\"}],\"name\":\"calcKeysReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gap\",\"type\":\"uint256\"}],\"name\":\"changeRndGap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inc\",\"type\":\"uint256\"}],\"name\":\"changeRndInc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endCurrentRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fees_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gen\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"p3d\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBuyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRoundInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPlayerInfoByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"}],\"name\":\"getPlayerVaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTimeLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_keys\",\"type\":\"uint256\"}],\"name\":\"iWantXKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pIDxAddr_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pIDxName_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"playerBook\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"plyrNames_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plyrRnds_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"eth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keys\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ico\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plyr_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"win\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gen\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lrnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"laff\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"potSplit_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gen\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"p3d\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rID_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_affCode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_team\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eth\",\"type\":\"uint256\"}],\"name\":\"reLoadXid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_laff\",\"type\":\"uint256\"}],\"name\":\"receivePlayerInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"receivePlayerNameList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rndTmEth_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"round_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"plyr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"team\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"strt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keys\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mask\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ico\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"icoGen\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"icoAvg\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawDust\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Fomo3dABI is the input ABI used to generate the binding from.
// Deprecated: Use Fomo3dMetaData.ABI instead.
var Fomo3dABI = Fomo3dMetaData.ABI

// Fomo3d is an auto generated Go binding around an Ethereum contract.
type Fomo3d struct {
	Fomo3dCaller     // Read-only binding to the contract
	Fomo3dTransactor // Write-only binding to the contract
	Fomo3dFilterer   // Log filterer for contract events
}

// Fomo3dCaller is an auto generated read-only Go binding around an Ethereum contract.
type Fomo3dCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fomo3dTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Fomo3dTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fomo3dFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Fomo3dFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fomo3dSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Fomo3dSession struct {
	Contract     *Fomo3d           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Fomo3dCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Fomo3dCallerSession struct {
	Contract *Fomo3dCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Fomo3dTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Fomo3dTransactorSession struct {
	Contract     *Fomo3dTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Fomo3dRaw is an auto generated low-level Go binding around an Ethereum contract.
type Fomo3dRaw struct {
	Contract *Fomo3d // Generic contract binding to access the raw methods on
}

// Fomo3dCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Fomo3dCallerRaw struct {
	Contract *Fomo3dCaller // Generic read-only contract binding to access the raw methods on
}

// Fomo3dTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Fomo3dTransactorRaw struct {
	Contract *Fomo3dTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFomo3d creates a new instance of Fomo3d, bound to a specific deployed contract.
func NewFomo3d(address common.Address, backend bind.ContractBackend) (*Fomo3d, error) {
	contract, err := bindFomo3d(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fomo3d{Fomo3dCaller: Fomo3dCaller{contract: contract}, Fomo3dTransactor: Fomo3dTransactor{contract: contract}, Fomo3dFilterer: Fomo3dFilterer{contract: contract}}, nil
}

// NewFomo3dCaller creates a new read-only instance of Fomo3d, bound to a specific deployed contract.
func NewFomo3dCaller(address common.Address, caller bind.ContractCaller) (*Fomo3dCaller, error) {
	contract, err := bindFomo3d(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Fomo3dCaller{contract: contract}, nil
}

// NewFomo3dTransactor creates a new write-only instance of Fomo3d, bound to a specific deployed contract.
func NewFomo3dTransactor(address common.Address, transactor bind.ContractTransactor) (*Fomo3dTransactor, error) {
	contract, err := bindFomo3d(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Fomo3dTransactor{contract: contract}, nil
}

// NewFomo3dFilterer creates a new log filterer instance of Fomo3d, bound to a specific deployed contract.
func NewFomo3dFilterer(address common.Address, filterer bind.ContractFilterer) (*Fomo3dFilterer, error) {
	contract, err := bindFomo3d(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Fomo3dFilterer{contract: contract}, nil
}

// bindFomo3d binds a generic wrapper to an already deployed contract.
func bindFomo3d(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Fomo3dABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fomo3d *Fomo3dRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fomo3d.Contract.Fomo3dCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fomo3d *Fomo3dRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fomo3d.Contract.Fomo3dTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fomo3d *Fomo3dRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fomo3d.Contract.Fomo3dTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fomo3d *Fomo3dCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fomo3d.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fomo3d *Fomo3dTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fomo3d.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fomo3d *Fomo3dTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fomo3d.Contract.contract.Transact(opts, method, params...)
}

// Activated is a free data retrieval call binding the contract method 0xd53b2679.
//
// Solidity: function activated_() view returns(bool)
func (_Fomo3d *Fomo3dCaller) Activated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "activated_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Activated is a free data retrieval call binding the contract method 0xd53b2679.
//
// Solidity: function activated_() view returns(bool)
func (_Fomo3d *Fomo3dSession) Activated() (bool, error) {
	return _Fomo3d.Contract.Activated(&_Fomo3d.CallOpts)
}

// Activated is a free data retrieval call binding the contract method 0xd53b2679.
//
// Solidity: function activated_() view returns(bool)
func (_Fomo3d *Fomo3dCallerSession) Activated() (bool, error) {
	return _Fomo3d.Contract.Activated(&_Fomo3d.CallOpts)
}

// AirDropPot is a free data retrieval call binding the contract method 0xd87574e0.
//
// Solidity: function airDropPot_() view returns(uint256)
func (_Fomo3d *Fomo3dCaller) AirDropPot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "airDropPot_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AirDropPot is a free data retrieval call binding the contract method 0xd87574e0.
//
// Solidity: function airDropPot_() view returns(uint256)
func (_Fomo3d *Fomo3dSession) AirDropPot() (*big.Int, error) {
	return _Fomo3d.Contract.AirDropPot(&_Fomo3d.CallOpts)
}

// AirDropPot is a free data retrieval call binding the contract method 0xd87574e0.
//
// Solidity: function airDropPot_() view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) AirDropPot() (*big.Int, error) {
	return _Fomo3d.Contract.AirDropPot(&_Fomo3d.CallOpts)
}

// AirDropTracker is a free data retrieval call binding the contract method 0x11a09ae7.
//
// Solidity: function airDropTracker_() view returns(uint256)
func (_Fomo3d *Fomo3dCaller) AirDropTracker(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "airDropTracker_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AirDropTracker is a free data retrieval call binding the contract method 0x11a09ae7.
//
// Solidity: function airDropTracker_() view returns(uint256)
func (_Fomo3d *Fomo3dSession) AirDropTracker() (*big.Int, error) {
	return _Fomo3d.Contract.AirDropTracker(&_Fomo3d.CallOpts)
}

// AirDropTracker is a free data retrieval call binding the contract method 0x11a09ae7.
//
// Solidity: function airDropTracker_() view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) AirDropTracker() (*big.Int, error) {
	return _Fomo3d.Contract.AirDropTracker(&_Fomo3d.CallOpts)
}

// CalcKeysReceived is a free data retrieval call binding the contract method 0xce89c80c.
//
// Solidity: function calcKeysReceived(uint256 _rID, uint256 _eth) view returns(uint256)
func (_Fomo3d *Fomo3dCaller) CalcKeysReceived(opts *bind.CallOpts, _rID *big.Int, _eth *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "calcKeysReceived", _rID, _eth)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcKeysReceived is a free data retrieval call binding the contract method 0xce89c80c.
//
// Solidity: function calcKeysReceived(uint256 _rID, uint256 _eth) view returns(uint256)
func (_Fomo3d *Fomo3dSession) CalcKeysReceived(_rID *big.Int, _eth *big.Int) (*big.Int, error) {
	return _Fomo3d.Contract.CalcKeysReceived(&_Fomo3d.CallOpts, _rID, _eth)
}

// CalcKeysReceived is a free data retrieval call binding the contract method 0xce89c80c.
//
// Solidity: function calcKeysReceived(uint256 _rID, uint256 _eth) view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) CalcKeysReceived(_rID *big.Int, _eth *big.Int) (*big.Int, error) {
	return _Fomo3d.Contract.CalcKeysReceived(&_Fomo3d.CallOpts, _rID, _eth)
}

// Fees is a free data retrieval call binding the contract method 0x2ce21999.
//
// Solidity: function fees_(uint256 ) view returns(uint256 gen, uint256 p3d)
func (_Fomo3d *Fomo3dCaller) Fees(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "fees_", arg0)

	outstruct := new(struct {
		Gen *big.Int
		P3d *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gen = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.P3d = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Fees is a free data retrieval call binding the contract method 0x2ce21999.
//
// Solidity: function fees_(uint256 ) view returns(uint256 gen, uint256 p3d)
func (_Fomo3d *Fomo3dSession) Fees(arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	return _Fomo3d.Contract.Fees(&_Fomo3d.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x2ce21999.
//
// Solidity: function fees_(uint256 ) view returns(uint256 gen, uint256 p3d)
func (_Fomo3d *Fomo3dCallerSession) Fees(arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	return _Fomo3d.Contract.Fees(&_Fomo3d.CallOpts, arg0)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() view returns(uint256)
func (_Fomo3d *Fomo3dCaller) GetBuyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "getBuyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() view returns(uint256)
func (_Fomo3d *Fomo3dSession) GetBuyPrice() (*big.Int, error) {
	return _Fomo3d.Contract.GetBuyPrice(&_Fomo3d.CallOpts)
}

// GetBuyPrice is a free data retrieval call binding the contract method 0x018a25e8.
//
// Solidity: function getBuyPrice() view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) GetBuyPrice() (*big.Int, error) {
	return _Fomo3d.Contract.GetBuyPrice(&_Fomo3d.CallOpts)
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, address, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_Fomo3d *Fomo3dCaller) GetCurrentRoundInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "getCurrentRoundInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(common.Address), *new([32]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	out8 := *abi.ConvertType(out[8], new([32]byte)).(*[32]byte)
	out9 := *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	out10 := *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	out11 := *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	out12 := *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)
	out13 := *abi.ConvertType(out[13], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, out12, out13, err

}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, address, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_Fomo3d *Fomo3dSession) GetCurrentRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Fomo3d.Contract.GetCurrentRoundInfo(&_Fomo3d.CallOpts)
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() view returns(uint256, uint256, uint256, uint256, uint256, uint256, uint256, address, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_Fomo3d *Fomo3dCallerSession) GetCurrentRoundInfo() (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Fomo3d.Contract.GetCurrentRoundInfo(&_Fomo3d.CallOpts)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(address _addr) view returns(uint256, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_Fomo3d *Fomo3dCaller) GetPlayerInfoByAddress(opts *bind.CallOpts, _addr common.Address) (*big.Int, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "getPlayerInfoByAddress", _addr)

	if err != nil {
		return *new(*big.Int), *new([32]byte), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(address _addr) view returns(uint256, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_Fomo3d *Fomo3dSession) GetPlayerInfoByAddress(_addr common.Address) (*big.Int, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Fomo3d.Contract.GetPlayerInfoByAddress(&_Fomo3d.CallOpts, _addr)
}

// GetPlayerInfoByAddress is a free data retrieval call binding the contract method 0xee0b5d8b.
//
// Solidity: function getPlayerInfoByAddress(address _addr) view returns(uint256, bytes32, uint256, uint256, uint256, uint256, uint256)
func (_Fomo3d *Fomo3dCallerSession) GetPlayerInfoByAddress(_addr common.Address) (*big.Int, [32]byte, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Fomo3d.Contract.GetPlayerInfoByAddress(&_Fomo3d.CallOpts, _addr)
}

// GetPlayerVaults is a free data retrieval call binding the contract method 0x63066434.
//
// Solidity: function getPlayerVaults(uint256 _pID) view returns(uint256, uint256, uint256)
func (_Fomo3d *Fomo3dCaller) GetPlayerVaults(opts *bind.CallOpts, _pID *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "getPlayerVaults", _pID)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPlayerVaults is a free data retrieval call binding the contract method 0x63066434.
//
// Solidity: function getPlayerVaults(uint256 _pID) view returns(uint256, uint256, uint256)
func (_Fomo3d *Fomo3dSession) GetPlayerVaults(_pID *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Fomo3d.Contract.GetPlayerVaults(&_Fomo3d.CallOpts, _pID)
}

// GetPlayerVaults is a free data retrieval call binding the contract method 0x63066434.
//
// Solidity: function getPlayerVaults(uint256 _pID) view returns(uint256, uint256, uint256)
func (_Fomo3d *Fomo3dCallerSession) GetPlayerVaults(_pID *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Fomo3d.Contract.GetPlayerVaults(&_Fomo3d.CallOpts, _pID)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() view returns(uint256)
func (_Fomo3d *Fomo3dCaller) GetTimeLeft(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "getTimeLeft")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() view returns(uint256)
func (_Fomo3d *Fomo3dSession) GetTimeLeft() (*big.Int, error) {
	return _Fomo3d.Contract.GetTimeLeft(&_Fomo3d.CallOpts)
}

// GetTimeLeft is a free data retrieval call binding the contract method 0xc7e284b8.
//
// Solidity: function getTimeLeft() view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) GetTimeLeft() (*big.Int, error) {
	return _Fomo3d.Contract.GetTimeLeft(&_Fomo3d.CallOpts)
}

// IWantXKeys is a free data retrieval call binding the contract method 0xcf808000.
//
// Solidity: function iWantXKeys(uint256 _keys) view returns(uint256)
func (_Fomo3d *Fomo3dCaller) IWantXKeys(opts *bind.CallOpts, _keys *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "iWantXKeys", _keys)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IWantXKeys is a free data retrieval call binding the contract method 0xcf808000.
//
// Solidity: function iWantXKeys(uint256 _keys) view returns(uint256)
func (_Fomo3d *Fomo3dSession) IWantXKeys(_keys *big.Int) (*big.Int, error) {
	return _Fomo3d.Contract.IWantXKeys(&_Fomo3d.CallOpts, _keys)
}

// IWantXKeys is a free data retrieval call binding the contract method 0xcf808000.
//
// Solidity: function iWantXKeys(uint256 _keys) view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) IWantXKeys(_keys *big.Int) (*big.Int, error) {
	return _Fomo3d.Contract.IWantXKeys(&_Fomo3d.CallOpts, _keys)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fomo3d *Fomo3dCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fomo3d *Fomo3dSession) Name() (string, error) {
	return _Fomo3d.Contract.Name(&_Fomo3d.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fomo3d *Fomo3dCallerSession) Name() (string, error) {
	return _Fomo3d.Contract.Name(&_Fomo3d.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fomo3d *Fomo3dCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fomo3d *Fomo3dSession) Owner() (common.Address, error) {
	return _Fomo3d.Contract.Owner(&_Fomo3d.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fomo3d *Fomo3dCallerSession) Owner() (common.Address, error) {
	return _Fomo3d.Contract.Owner(&_Fomo3d.CallOpts)
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_(address ) view returns(uint256)
func (_Fomo3d *Fomo3dCaller) PIDxAddr(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "pIDxAddr_", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_(address ) view returns(uint256)
func (_Fomo3d *Fomo3dSession) PIDxAddr(arg0 common.Address) (*big.Int, error) {
	return _Fomo3d.Contract.PIDxAddr(&_Fomo3d.CallOpts, arg0)
}

// PIDxAddr is a free data retrieval call binding the contract method 0x10f01eba.
//
// Solidity: function pIDxAddr_(address ) view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) PIDxAddr(arg0 common.Address) (*big.Int, error) {
	return _Fomo3d.Contract.PIDxAddr(&_Fomo3d.CallOpts, arg0)
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_(bytes32 ) view returns(uint256)
func (_Fomo3d *Fomo3dCaller) PIDxName(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "pIDxName_", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_(bytes32 ) view returns(uint256)
func (_Fomo3d *Fomo3dSession) PIDxName(arg0 [32]byte) (*big.Int, error) {
	return _Fomo3d.Contract.PIDxName(&_Fomo3d.CallOpts, arg0)
}

// PIDxName is a free data retrieval call binding the contract method 0x2e19ebdc.
//
// Solidity: function pIDxName_(bytes32 ) view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) PIDxName(arg0 [32]byte) (*big.Int, error) {
	return _Fomo3d.Contract.PIDxName(&_Fomo3d.CallOpts, arg0)
}

// PlayerBook is a free data retrieval call binding the contract method 0x2fa241fc.
//
// Solidity: function playerBook() view returns(address)
func (_Fomo3d *Fomo3dCaller) PlayerBook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "playerBook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlayerBook is a free data retrieval call binding the contract method 0x2fa241fc.
//
// Solidity: function playerBook() view returns(address)
func (_Fomo3d *Fomo3dSession) PlayerBook() (common.Address, error) {
	return _Fomo3d.Contract.PlayerBook(&_Fomo3d.CallOpts)
}

// PlayerBook is a free data retrieval call binding the contract method 0x2fa241fc.
//
// Solidity: function playerBook() view returns(address)
func (_Fomo3d *Fomo3dCallerSession) PlayerBook() (common.Address, error) {
	return _Fomo3d.Contract.PlayerBook(&_Fomo3d.CallOpts)
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_(uint256 , bytes32 ) view returns(bool)
func (_Fomo3d *Fomo3dCaller) PlyrNames(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "plyrNames_", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_(uint256 , bytes32 ) view returns(bool)
func (_Fomo3d *Fomo3dSession) PlyrNames(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _Fomo3d.Contract.PlyrNames(&_Fomo3d.CallOpts, arg0, arg1)
}

// PlyrNames is a free data retrieval call binding the contract method 0x2660316e.
//
// Solidity: function plyrNames_(uint256 , bytes32 ) view returns(bool)
func (_Fomo3d *Fomo3dCallerSession) PlyrNames(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _Fomo3d.Contract.PlyrNames(&_Fomo3d.CallOpts, arg0, arg1)
}

// PlyrRnds is a free data retrieval call binding the contract method 0xa2bccae9.
//
// Solidity: function plyrRnds_(uint256 , uint256 ) view returns(uint256 eth, uint256 keys, uint256 mask, uint256 ico)
func (_Fomo3d *Fomo3dCaller) PlyrRnds(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Eth  *big.Int
	Keys *big.Int
	Mask *big.Int
	Ico  *big.Int
}, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "plyrRnds_", arg0, arg1)

	outstruct := new(struct {
		Eth  *big.Int
		Keys *big.Int
		Mask *big.Int
		Ico  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Eth = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Keys = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Mask = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Ico = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PlyrRnds is a free data retrieval call binding the contract method 0xa2bccae9.
//
// Solidity: function plyrRnds_(uint256 , uint256 ) view returns(uint256 eth, uint256 keys, uint256 mask, uint256 ico)
func (_Fomo3d *Fomo3dSession) PlyrRnds(arg0 *big.Int, arg1 *big.Int) (struct {
	Eth  *big.Int
	Keys *big.Int
	Mask *big.Int
	Ico  *big.Int
}, error) {
	return _Fomo3d.Contract.PlyrRnds(&_Fomo3d.CallOpts, arg0, arg1)
}

// PlyrRnds is a free data retrieval call binding the contract method 0xa2bccae9.
//
// Solidity: function plyrRnds_(uint256 , uint256 ) view returns(uint256 eth, uint256 keys, uint256 mask, uint256 ico)
func (_Fomo3d *Fomo3dCallerSession) PlyrRnds(arg0 *big.Int, arg1 *big.Int) (struct {
	Eth  *big.Int
	Keys *big.Int
	Mask *big.Int
	Ico  *big.Int
}, error) {
	return _Fomo3d.Contract.PlyrRnds(&_Fomo3d.CallOpts, arg0, arg1)
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_(uint256 ) view returns(address addr, bytes32 name, uint256 win, uint256 gen, uint256 aff, uint256 lrnd, uint256 laff)
func (_Fomo3d *Fomo3dCaller) Plyr(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr common.Address
	Name [32]byte
	Win  *big.Int
	Gen  *big.Int
	Aff  *big.Int
	Lrnd *big.Int
	Laff *big.Int
}, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "plyr_", arg0)

	outstruct := new(struct {
		Addr common.Address
		Name [32]byte
		Win  *big.Int
		Gen  *big.Int
		Aff  *big.Int
		Lrnd *big.Int
		Laff *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Win = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Gen = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Aff = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Lrnd = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Laff = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_(uint256 ) view returns(address addr, bytes32 name, uint256 win, uint256 gen, uint256 aff, uint256 lrnd, uint256 laff)
func (_Fomo3d *Fomo3dSession) Plyr(arg0 *big.Int) (struct {
	Addr common.Address
	Name [32]byte
	Win  *big.Int
	Gen  *big.Int
	Aff  *big.Int
	Lrnd *big.Int
	Laff *big.Int
}, error) {
	return _Fomo3d.Contract.Plyr(&_Fomo3d.CallOpts, arg0)
}

// Plyr is a free data retrieval call binding the contract method 0xde7874f3.
//
// Solidity: function plyr_(uint256 ) view returns(address addr, bytes32 name, uint256 win, uint256 gen, uint256 aff, uint256 lrnd, uint256 laff)
func (_Fomo3d *Fomo3dCallerSession) Plyr(arg0 *big.Int) (struct {
	Addr common.Address
	Name [32]byte
	Win  *big.Int
	Gen  *big.Int
	Aff  *big.Int
	Lrnd *big.Int
	Laff *big.Int
}, error) {
	return _Fomo3d.Contract.Plyr(&_Fomo3d.CallOpts, arg0)
}

// PotSplit is a free data retrieval call binding the contract method 0xc519500e.
//
// Solidity: function potSplit_(uint256 ) view returns(uint256 gen, uint256 p3d)
func (_Fomo3d *Fomo3dCaller) PotSplit(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "potSplit_", arg0)

	outstruct := new(struct {
		Gen *big.Int
		P3d *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gen = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.P3d = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PotSplit is a free data retrieval call binding the contract method 0xc519500e.
//
// Solidity: function potSplit_(uint256 ) view returns(uint256 gen, uint256 p3d)
func (_Fomo3d *Fomo3dSession) PotSplit(arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	return _Fomo3d.Contract.PotSplit(&_Fomo3d.CallOpts, arg0)
}

// PotSplit is a free data retrieval call binding the contract method 0xc519500e.
//
// Solidity: function potSplit_(uint256 ) view returns(uint256 gen, uint256 p3d)
func (_Fomo3d *Fomo3dCallerSession) PotSplit(arg0 *big.Int) (struct {
	Gen *big.Int
	P3d *big.Int
}, error) {
	return _Fomo3d.Contract.PotSplit(&_Fomo3d.CallOpts, arg0)
}

// RID is a free data retrieval call binding the contract method 0x624ae5c0.
//
// Solidity: function rID_() view returns(uint256)
func (_Fomo3d *Fomo3dCaller) RID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "rID_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RID is a free data retrieval call binding the contract method 0x624ae5c0.
//
// Solidity: function rID_() view returns(uint256)
func (_Fomo3d *Fomo3dSession) RID() (*big.Int, error) {
	return _Fomo3d.Contract.RID(&_Fomo3d.CallOpts)
}

// RID is a free data retrieval call binding the contract method 0x624ae5c0.
//
// Solidity: function rID_() view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) RID() (*big.Int, error) {
	return _Fomo3d.Contract.RID(&_Fomo3d.CallOpts)
}

// RndTmEth is a free data retrieval call binding the contract method 0x5893d481.
//
// Solidity: function rndTmEth_(uint256 , uint256 ) view returns(uint256)
func (_Fomo3d *Fomo3dCaller) RndTmEth(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "rndTmEth_", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RndTmEth is a free data retrieval call binding the contract method 0x5893d481.
//
// Solidity: function rndTmEth_(uint256 , uint256 ) view returns(uint256)
func (_Fomo3d *Fomo3dSession) RndTmEth(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Fomo3d.Contract.RndTmEth(&_Fomo3d.CallOpts, arg0, arg1)
}

// RndTmEth is a free data retrieval call binding the contract method 0x5893d481.
//
// Solidity: function rndTmEth_(uint256 , uint256 ) view returns(uint256)
func (_Fomo3d *Fomo3dCallerSession) RndTmEth(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Fomo3d.Contract.RndTmEth(&_Fomo3d.CallOpts, arg0, arg1)
}

// Round is a free data retrieval call binding the contract method 0x24c33d33.
//
// Solidity: function round_(uint256 ) view returns(uint256 plyr, uint256 team, uint256 end, bool ended, uint256 strt, uint256 keys, uint256 eth, uint256 pot, uint256 mask, uint256 ico, uint256 icoGen, uint256 icoAvg)
func (_Fomo3d *Fomo3dCaller) Round(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Plyr   *big.Int
	Team   *big.Int
	End    *big.Int
	Ended  bool
	Strt   *big.Int
	Keys   *big.Int
	Eth    *big.Int
	Pot    *big.Int
	Mask   *big.Int
	Ico    *big.Int
	IcoGen *big.Int
	IcoAvg *big.Int
}, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "round_", arg0)

	outstruct := new(struct {
		Plyr   *big.Int
		Team   *big.Int
		End    *big.Int
		Ended  bool
		Strt   *big.Int
		Keys   *big.Int
		Eth    *big.Int
		Pot    *big.Int
		Mask   *big.Int
		Ico    *big.Int
		IcoGen *big.Int
		IcoAvg *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Plyr = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Team = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Ended = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Strt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Keys = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Eth = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Pot = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Mask = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Ico = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.IcoGen = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.IcoAvg = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Round is a free data retrieval call binding the contract method 0x24c33d33.
//
// Solidity: function round_(uint256 ) view returns(uint256 plyr, uint256 team, uint256 end, bool ended, uint256 strt, uint256 keys, uint256 eth, uint256 pot, uint256 mask, uint256 ico, uint256 icoGen, uint256 icoAvg)
func (_Fomo3d *Fomo3dSession) Round(arg0 *big.Int) (struct {
	Plyr   *big.Int
	Team   *big.Int
	End    *big.Int
	Ended  bool
	Strt   *big.Int
	Keys   *big.Int
	Eth    *big.Int
	Pot    *big.Int
	Mask   *big.Int
	Ico    *big.Int
	IcoGen *big.Int
	IcoAvg *big.Int
}, error) {
	return _Fomo3d.Contract.Round(&_Fomo3d.CallOpts, arg0)
}

// Round is a free data retrieval call binding the contract method 0x24c33d33.
//
// Solidity: function round_(uint256 ) view returns(uint256 plyr, uint256 team, uint256 end, bool ended, uint256 strt, uint256 keys, uint256 eth, uint256 pot, uint256 mask, uint256 ico, uint256 icoGen, uint256 icoAvg)
func (_Fomo3d *Fomo3dCallerSession) Round(arg0 *big.Int) (struct {
	Plyr   *big.Int
	Team   *big.Int
	End    *big.Int
	Ended  bool
	Strt   *big.Int
	Keys   *big.Int
	Eth    *big.Int
	Pot    *big.Int
	Mask   *big.Int
	Ico    *big.Int
	IcoGen *big.Int
	IcoAvg *big.Int
}, error) {
	return _Fomo3d.Contract.Round(&_Fomo3d.CallOpts, arg0)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Fomo3d *Fomo3dCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Fomo3d *Fomo3dSession) Staking() (common.Address, error) {
	return _Fomo3d.Contract.Staking(&_Fomo3d.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Fomo3d *Fomo3dCallerSession) Staking() (common.Address, error) {
	return _Fomo3d.Contract.Staking(&_Fomo3d.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fomo3d *Fomo3dCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fomo3d *Fomo3dSession) Symbol() (string, error) {
	return _Fomo3d.Contract.Symbol(&_Fomo3d.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fomo3d *Fomo3dCallerSession) Symbol() (string, error) {
	return _Fomo3d.Contract.Symbol(&_Fomo3d.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Fomo3d *Fomo3dCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fomo3d.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Fomo3d *Fomo3dSession) Treasury() (common.Address, error) {
	return _Fomo3d.Contract.Treasury(&_Fomo3d.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Fomo3d *Fomo3dCallerSession) Treasury() (common.Address, error) {
	return _Fomo3d.Contract.Treasury(&_Fomo3d.CallOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Fomo3d *Fomo3dTransactor) Activate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "activate")
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Fomo3d *Fomo3dSession) Activate() (*types.Transaction, error) {
	return _Fomo3d.Contract.Activate(&_Fomo3d.TransactOpts)
}

// Activate is a paid mutator transaction binding the contract method 0x0f15f4c0.
//
// Solidity: function activate() returns()
func (_Fomo3d *Fomo3dTransactorSession) Activate() (*types.Transaction, error) {
	return _Fomo3d.Contract.Activate(&_Fomo3d.TransactOpts)
}

// BuyXid is a paid mutator transaction binding the contract method 0x8f38f309.
//
// Solidity: function buyXid(uint256 _affCode, uint256 _team) payable returns()
func (_Fomo3d *Fomo3dTransactor) BuyXid(opts *bind.TransactOpts, _affCode *big.Int, _team *big.Int) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "buyXid", _affCode, _team)
}

// BuyXid is a paid mutator transaction binding the contract method 0x8f38f309.
//
// Solidity: function buyXid(uint256 _affCode, uint256 _team) payable returns()
func (_Fomo3d *Fomo3dSession) BuyXid(_affCode *big.Int, _team *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.BuyXid(&_Fomo3d.TransactOpts, _affCode, _team)
}

// BuyXid is a paid mutator transaction binding the contract method 0x8f38f309.
//
// Solidity: function buyXid(uint256 _affCode, uint256 _team) payable returns()
func (_Fomo3d *Fomo3dTransactorSession) BuyXid(_affCode *big.Int, _team *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.BuyXid(&_Fomo3d.TransactOpts, _affCode, _team)
}

// ChangeRndGap is a paid mutator transaction binding the contract method 0x816f172a.
//
// Solidity: function changeRndGap(uint256 gap) returns()
func (_Fomo3d *Fomo3dTransactor) ChangeRndGap(opts *bind.TransactOpts, gap *big.Int) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "changeRndGap", gap)
}

// ChangeRndGap is a paid mutator transaction binding the contract method 0x816f172a.
//
// Solidity: function changeRndGap(uint256 gap) returns()
func (_Fomo3d *Fomo3dSession) ChangeRndGap(gap *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.ChangeRndGap(&_Fomo3d.TransactOpts, gap)
}

// ChangeRndGap is a paid mutator transaction binding the contract method 0x816f172a.
//
// Solidity: function changeRndGap(uint256 gap) returns()
func (_Fomo3d *Fomo3dTransactorSession) ChangeRndGap(gap *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.ChangeRndGap(&_Fomo3d.TransactOpts, gap)
}

// ChangeRndInc is a paid mutator transaction binding the contract method 0x993aaefa.
//
// Solidity: function changeRndInc(uint256 inc) returns()
func (_Fomo3d *Fomo3dTransactor) ChangeRndInc(opts *bind.TransactOpts, inc *big.Int) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "changeRndInc", inc)
}

// ChangeRndInc is a paid mutator transaction binding the contract method 0x993aaefa.
//
// Solidity: function changeRndInc(uint256 inc) returns()
func (_Fomo3d *Fomo3dSession) ChangeRndInc(inc *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.ChangeRndInc(&_Fomo3d.TransactOpts, inc)
}

// ChangeRndInc is a paid mutator transaction binding the contract method 0x993aaefa.
//
// Solidity: function changeRndInc(uint256 inc) returns()
func (_Fomo3d *Fomo3dTransactorSession) ChangeRndInc(inc *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.ChangeRndInc(&_Fomo3d.TransactOpts, inc)
}

// EndCurrentRound is a paid mutator transaction binding the contract method 0xf4e3bdc8.
//
// Solidity: function endCurrentRound() returns()
func (_Fomo3d *Fomo3dTransactor) EndCurrentRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "endCurrentRound")
}

// EndCurrentRound is a paid mutator transaction binding the contract method 0xf4e3bdc8.
//
// Solidity: function endCurrentRound() returns()
func (_Fomo3d *Fomo3dSession) EndCurrentRound() (*types.Transaction, error) {
	return _Fomo3d.Contract.EndCurrentRound(&_Fomo3d.TransactOpts)
}

// EndCurrentRound is a paid mutator transaction binding the contract method 0xf4e3bdc8.
//
// Solidity: function endCurrentRound() returns()
func (_Fomo3d *Fomo3dTransactorSession) EndCurrentRound() (*types.Transaction, error) {
	return _Fomo3d.Contract.EndCurrentRound(&_Fomo3d.TransactOpts)
}

// ReLoadXid is a paid mutator transaction binding the contract method 0x349cdcac.
//
// Solidity: function reLoadXid(uint256 _affCode, uint256 _team, uint256 _eth) returns()
func (_Fomo3d *Fomo3dTransactor) ReLoadXid(opts *bind.TransactOpts, _affCode *big.Int, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "reLoadXid", _affCode, _team, _eth)
}

// ReLoadXid is a paid mutator transaction binding the contract method 0x349cdcac.
//
// Solidity: function reLoadXid(uint256 _affCode, uint256 _team, uint256 _eth) returns()
func (_Fomo3d *Fomo3dSession) ReLoadXid(_affCode *big.Int, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.ReLoadXid(&_Fomo3d.TransactOpts, _affCode, _team, _eth)
}

// ReLoadXid is a paid mutator transaction binding the contract method 0x349cdcac.
//
// Solidity: function reLoadXid(uint256 _affCode, uint256 _team, uint256 _eth) returns()
func (_Fomo3d *Fomo3dTransactorSession) ReLoadXid(_affCode *big.Int, _team *big.Int, _eth *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.ReLoadXid(&_Fomo3d.TransactOpts, _affCode, _team, _eth)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(uint256 _pID, address _addr, bytes32 _name, uint256 _laff) returns()
func (_Fomo3d *Fomo3dTransactor) ReceivePlayerInfo(opts *bind.TransactOpts, _pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "receivePlayerInfo", _pID, _addr, _name, _laff)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(uint256 _pID, address _addr, bytes32 _name, uint256 _laff) returns()
func (_Fomo3d *Fomo3dSession) ReceivePlayerInfo(_pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.ReceivePlayerInfo(&_Fomo3d.TransactOpts, _pID, _addr, _name, _laff)
}

// ReceivePlayerInfo is a paid mutator transaction binding the contract method 0x49cc635d.
//
// Solidity: function receivePlayerInfo(uint256 _pID, address _addr, bytes32 _name, uint256 _laff) returns()
func (_Fomo3d *Fomo3dTransactorSession) ReceivePlayerInfo(_pID *big.Int, _addr common.Address, _name [32]byte, _laff *big.Int) (*types.Transaction, error) {
	return _Fomo3d.Contract.ReceivePlayerInfo(&_Fomo3d.TransactOpts, _pID, _addr, _name, _laff)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(uint256 _pID, bytes32 _name) returns()
func (_Fomo3d *Fomo3dTransactor) ReceivePlayerNameList(opts *bind.TransactOpts, _pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "receivePlayerNameList", _pID, _name)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(uint256 _pID, bytes32 _name) returns()
func (_Fomo3d *Fomo3dSession) ReceivePlayerNameList(_pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _Fomo3d.Contract.ReceivePlayerNameList(&_Fomo3d.TransactOpts, _pID, _name)
}

// ReceivePlayerNameList is a paid mutator transaction binding the contract method 0x8f7140ea.
//
// Solidity: function receivePlayerNameList(uint256 _pID, bytes32 _name) returns()
func (_Fomo3d *Fomo3dTransactorSession) ReceivePlayerNameList(_pID *big.Int, _name [32]byte) (*types.Transaction, error) {
	return _Fomo3d.Contract.ReceivePlayerNameList(&_Fomo3d.TransactOpts, _pID, _name)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Fomo3d *Fomo3dTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Fomo3d *Fomo3dSession) RenounceOwnership() (*types.Transaction, error) {
	return _Fomo3d.Contract.RenounceOwnership(&_Fomo3d.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Fomo3d *Fomo3dTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Fomo3d.Contract.RenounceOwnership(&_Fomo3d.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fomo3d *Fomo3dTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fomo3d *Fomo3dSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fomo3d.Contract.TransferOwnership(&_Fomo3d.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fomo3d *Fomo3dTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fomo3d.Contract.TransferOwnership(&_Fomo3d.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Fomo3d *Fomo3dTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Fomo3d *Fomo3dSession) Withdraw() (*types.Transaction, error) {
	return _Fomo3d.Contract.Withdraw(&_Fomo3d.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Fomo3d *Fomo3dTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Fomo3d.Contract.Withdraw(&_Fomo3d.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_Fomo3d *Fomo3dTransactor) WithdrawDust(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fomo3d.contract.Transact(opts, "withdrawDust")
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_Fomo3d *Fomo3dSession) WithdrawDust() (*types.Transaction, error) {
	return _Fomo3d.Contract.WithdrawDust(&_Fomo3d.TransactOpts)
}

// WithdrawDust is a paid mutator transaction binding the contract method 0xcfb550f1.
//
// Solidity: function withdrawDust() returns()
func (_Fomo3d *Fomo3dTransactorSession) WithdrawDust() (*types.Transaction, error) {
	return _Fomo3d.Contract.WithdrawDust(&_Fomo3d.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fomo3d *Fomo3dTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fomo3d.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fomo3d *Fomo3dSession) Receive() (*types.Transaction, error) {
	return _Fomo3d.Contract.Receive(&_Fomo3d.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Fomo3d *Fomo3dTransactorSession) Receive() (*types.Transaction, error) {
	return _Fomo3d.Contract.Receive(&_Fomo3d.TransactOpts)
}

// Fomo3dOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Fomo3d contract.
type Fomo3dOwnershipTransferredIterator struct {
	Event *Fomo3dOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Fomo3dOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fomo3dOwnershipTransferred)
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
		it.Event = new(Fomo3dOwnershipTransferred)
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
func (it *Fomo3dOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fomo3dOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fomo3dOwnershipTransferred represents a OwnershipTransferred event raised by the Fomo3d contract.
type Fomo3dOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Fomo3d *Fomo3dFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Fomo3dOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fomo3d.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Fomo3dOwnershipTransferredIterator{contract: _Fomo3d.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Fomo3d *Fomo3dFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Fomo3dOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fomo3d.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fomo3dOwnershipTransferred)
				if err := _Fomo3d.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Fomo3d *Fomo3dFilterer) ParseOwnershipTransferred(log types.Log) (*Fomo3dOwnershipTransferred, error) {
	event := new(Fomo3dOwnershipTransferred)
	if err := _Fomo3d.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fomo3dOnAffiliatePayoutIterator is returned from FilterOnAffiliatePayout and is used to iterate over the raw logs and unpacked data for OnAffiliatePayout events raised by the Fomo3d contract.
type Fomo3dOnAffiliatePayoutIterator struct {
	Event *Fomo3dOnAffiliatePayout // Event containing the contract specifics and raw log

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
func (it *Fomo3dOnAffiliatePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fomo3dOnAffiliatePayout)
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
		it.Event = new(Fomo3dOnAffiliatePayout)
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
func (it *Fomo3dOnAffiliatePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fomo3dOnAffiliatePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fomo3dOnAffiliatePayout represents a OnAffiliatePayout event raised by the Fomo3d contract.
type Fomo3dOnAffiliatePayout struct {
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	RoundID          *big.Int
	BuyerID          *big.Int
	Amount           *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnAffiliatePayout is a free log retrieval operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: event onAffiliatePayout(uint256 indexed affiliateID, address affiliateAddress, bytes32 affiliateName, uint256 indexed roundID, uint256 indexed buyerID, uint256 amount, uint256 timeStamp)
func (_Fomo3d *Fomo3dFilterer) FilterOnAffiliatePayout(opts *bind.FilterOpts, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (*Fomo3dOnAffiliatePayoutIterator, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _Fomo3d.contract.FilterLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return &Fomo3dOnAffiliatePayoutIterator{contract: _Fomo3d.contract, event: "onAffiliatePayout", logs: logs, sub: sub}, nil
}

// WatchOnAffiliatePayout is a free log subscription operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: event onAffiliatePayout(uint256 indexed affiliateID, address affiliateAddress, bytes32 affiliateName, uint256 indexed roundID, uint256 indexed buyerID, uint256 amount, uint256 timeStamp)
func (_Fomo3d *Fomo3dFilterer) WatchOnAffiliatePayout(opts *bind.WatchOpts, sink chan<- *Fomo3dOnAffiliatePayout, affiliateID []*big.Int, roundID []*big.Int, buyerID []*big.Int) (event.Subscription, error) {

	var affiliateIDRule []interface{}
	for _, affiliateIDItem := range affiliateID {
		affiliateIDRule = append(affiliateIDRule, affiliateIDItem)
	}

	var roundIDRule []interface{}
	for _, roundIDItem := range roundID {
		roundIDRule = append(roundIDRule, roundIDItem)
	}
	var buyerIDRule []interface{}
	for _, buyerIDItem := range buyerID {
		buyerIDRule = append(buyerIDRule, buyerIDItem)
	}

	logs, sub, err := _Fomo3d.contract.WatchLogs(opts, "onAffiliatePayout", affiliateIDRule, roundIDRule, buyerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fomo3dOnAffiliatePayout)
				if err := _Fomo3d.contract.UnpackLog(event, "onAffiliatePayout", log); err != nil {
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

// ParseOnAffiliatePayout is a log parse operation binding the contract event 0x590bbc0fc16915a85269a48f74783c39842b7ae9eceb7c295c95dbe8b3ec7331.
//
// Solidity: event onAffiliatePayout(uint256 indexed affiliateID, address affiliateAddress, bytes32 affiliateName, uint256 indexed roundID, uint256 indexed buyerID, uint256 amount, uint256 timeStamp)
func (_Fomo3d *Fomo3dFilterer) ParseOnAffiliatePayout(log types.Log) (*Fomo3dOnAffiliatePayout, error) {
	event := new(Fomo3dOnAffiliatePayout)
	if err := _Fomo3d.contract.UnpackLog(event, "onAffiliatePayout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fomo3dOnBuyAndDistributeIterator is returned from FilterOnBuyAndDistribute and is used to iterate over the raw logs and unpacked data for OnBuyAndDistribute events raised by the Fomo3d contract.
type Fomo3dOnBuyAndDistributeIterator struct {
	Event *Fomo3dOnBuyAndDistribute // Event containing the contract specifics and raw log

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
func (it *Fomo3dOnBuyAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fomo3dOnBuyAndDistribute)
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
		it.Event = new(Fomo3dOnBuyAndDistribute)
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
func (it *Fomo3dOnBuyAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fomo3dOnBuyAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fomo3dOnBuyAndDistribute represents a OnBuyAndDistribute event raised by the Fomo3d contract.
type Fomo3dOnBuyAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EthIn          *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnBuyAndDistribute is a free log retrieval operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: event onBuyAndDistribute(address playerAddress, bytes32 playerName, uint256 ethIn, uint256 compressedData, uint256 compressedIDs, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount)
func (_Fomo3d *Fomo3dFilterer) FilterOnBuyAndDistribute(opts *bind.FilterOpts) (*Fomo3dOnBuyAndDistributeIterator, error) {

	logs, sub, err := _Fomo3d.contract.FilterLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return &Fomo3dOnBuyAndDistributeIterator{contract: _Fomo3d.contract, event: "onBuyAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnBuyAndDistribute is a free log subscription operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: event onBuyAndDistribute(address playerAddress, bytes32 playerName, uint256 ethIn, uint256 compressedData, uint256 compressedIDs, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount)
func (_Fomo3d *Fomo3dFilterer) WatchOnBuyAndDistribute(opts *bind.WatchOpts, sink chan<- *Fomo3dOnBuyAndDistribute) (event.Subscription, error) {

	logs, sub, err := _Fomo3d.contract.WatchLogs(opts, "onBuyAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fomo3dOnBuyAndDistribute)
				if err := _Fomo3d.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
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

// ParseOnBuyAndDistribute is a log parse operation binding the contract event 0xa7801a70b37e729a11492aad44fd3dba89b4149f0609dc0f6837bf9e57e2671a.
//
// Solidity: event onBuyAndDistribute(address playerAddress, bytes32 playerName, uint256 ethIn, uint256 compressedData, uint256 compressedIDs, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount)
func (_Fomo3d *Fomo3dFilterer) ParseOnBuyAndDistribute(log types.Log) (*Fomo3dOnBuyAndDistribute, error) {
	event := new(Fomo3dOnBuyAndDistribute)
	if err := _Fomo3d.contract.UnpackLog(event, "onBuyAndDistribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fomo3dOnEndTxIterator is returned from FilterOnEndTx and is used to iterate over the raw logs and unpacked data for OnEndTx events raised by the Fomo3d contract.
type Fomo3dOnEndTxIterator struct {
	Event *Fomo3dOnEndTx // Event containing the contract specifics and raw log

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
func (it *Fomo3dOnEndTxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fomo3dOnEndTx)
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
		it.Event = new(Fomo3dOnEndTx)
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
func (it *Fomo3dOnEndTxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fomo3dOnEndTxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fomo3dOnEndTx represents a OnEndTx event raised by the Fomo3d contract.
type Fomo3dOnEndTx struct {
	CompressedData *big.Int
	CompressedIDs  *big.Int
	PlayerName     [32]byte
	PlayerAddress  common.Address
	EthIn          *big.Int
	KeysBought     *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	PotAmount      *big.Int
	AirDropPot     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnEndTx is a free log retrieval operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: event onEndTx(uint256 compressedData, uint256 compressedIDs, bytes32 playerName, address playerAddress, uint256 ethIn, uint256 keysBought, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount, uint256 potAmount, uint256 airDropPot)
func (_Fomo3d *Fomo3dFilterer) FilterOnEndTx(opts *bind.FilterOpts) (*Fomo3dOnEndTxIterator, error) {

	logs, sub, err := _Fomo3d.contract.FilterLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return &Fomo3dOnEndTxIterator{contract: _Fomo3d.contract, event: "onEndTx", logs: logs, sub: sub}, nil
}

// WatchOnEndTx is a free log subscription operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: event onEndTx(uint256 compressedData, uint256 compressedIDs, bytes32 playerName, address playerAddress, uint256 ethIn, uint256 keysBought, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount, uint256 potAmount, uint256 airDropPot)
func (_Fomo3d *Fomo3dFilterer) WatchOnEndTx(opts *bind.WatchOpts, sink chan<- *Fomo3dOnEndTx) (event.Subscription, error) {

	logs, sub, err := _Fomo3d.contract.WatchLogs(opts, "onEndTx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fomo3dOnEndTx)
				if err := _Fomo3d.contract.UnpackLog(event, "onEndTx", log); err != nil {
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

// ParseOnEndTx is a log parse operation binding the contract event 0x500e72a0e114930aebdbcb371ccdbf43922c49f979794b5de4257ff7e310c746.
//
// Solidity: event onEndTx(uint256 compressedData, uint256 compressedIDs, bytes32 playerName, address playerAddress, uint256 ethIn, uint256 keysBought, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount, uint256 potAmount, uint256 airDropPot)
func (_Fomo3d *Fomo3dFilterer) ParseOnEndTx(log types.Log) (*Fomo3dOnEndTx, error) {
	event := new(Fomo3dOnEndTx)
	if err := _Fomo3d.contract.UnpackLog(event, "onEndTx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fomo3dOnNewNameIterator is returned from FilterOnNewName and is used to iterate over the raw logs and unpacked data for OnNewName events raised by the Fomo3d contract.
type Fomo3dOnNewNameIterator struct {
	Event *Fomo3dOnNewName // Event containing the contract specifics and raw log

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
func (it *Fomo3dOnNewNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fomo3dOnNewName)
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
		it.Event = new(Fomo3dOnNewName)
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
func (it *Fomo3dOnNewNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fomo3dOnNewNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fomo3dOnNewName represents a OnNewName event raised by the Fomo3d contract.
type Fomo3dOnNewName struct {
	PlayerID         *big.Int
	PlayerAddress    common.Address
	PlayerName       [32]byte
	IsNewPlayer      bool
	AffiliateID      *big.Int
	AffiliateAddress common.Address
	AffiliateName    [32]byte
	AmountPaid       *big.Int
	TimeStamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnNewName is a free log retrieval operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(uint256 indexed playerID, address indexed playerAddress, bytes32 indexed playerName, bool isNewPlayer, uint256 affiliateID, address affiliateAddress, bytes32 affiliateName, uint256 amountPaid, uint256 timeStamp)
func (_Fomo3d *Fomo3dFilterer) FilterOnNewName(opts *bind.FilterOpts, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (*Fomo3dOnNewNameIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _Fomo3d.contract.FilterLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return &Fomo3dOnNewNameIterator{contract: _Fomo3d.contract, event: "onNewName", logs: logs, sub: sub}, nil
}

// WatchOnNewName is a free log subscription operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(uint256 indexed playerID, address indexed playerAddress, bytes32 indexed playerName, bool isNewPlayer, uint256 affiliateID, address affiliateAddress, bytes32 affiliateName, uint256 amountPaid, uint256 timeStamp)
func (_Fomo3d *Fomo3dFilterer) WatchOnNewName(opts *bind.WatchOpts, sink chan<- *Fomo3dOnNewName, playerID []*big.Int, playerAddress []common.Address, playerName [][32]byte) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}
	var playerAddressRule []interface{}
	for _, playerAddressItem := range playerAddress {
		playerAddressRule = append(playerAddressRule, playerAddressItem)
	}
	var playerNameRule []interface{}
	for _, playerNameItem := range playerName {
		playerNameRule = append(playerNameRule, playerNameItem)
	}

	logs, sub, err := _Fomo3d.contract.WatchLogs(opts, "onNewName", playerIDRule, playerAddressRule, playerNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fomo3dOnNewName)
				if err := _Fomo3d.contract.UnpackLog(event, "onNewName", log); err != nil {
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

// ParseOnNewName is a log parse operation binding the contract event 0xdd6176433ff5026bbce96b068584b7bbe3514227e72df9c630b749ae87e64442.
//
// Solidity: event onNewName(uint256 indexed playerID, address indexed playerAddress, bytes32 indexed playerName, bool isNewPlayer, uint256 affiliateID, address affiliateAddress, bytes32 affiliateName, uint256 amountPaid, uint256 timeStamp)
func (_Fomo3d *Fomo3dFilterer) ParseOnNewName(log types.Log) (*Fomo3dOnNewName, error) {
	event := new(Fomo3dOnNewName)
	if err := _Fomo3d.contract.UnpackLog(event, "onNewName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fomo3dOnReLoadAndDistributeIterator is returned from FilterOnReLoadAndDistribute and is used to iterate over the raw logs and unpacked data for OnReLoadAndDistribute events raised by the Fomo3d contract.
type Fomo3dOnReLoadAndDistributeIterator struct {
	Event *Fomo3dOnReLoadAndDistribute // Event containing the contract specifics and raw log

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
func (it *Fomo3dOnReLoadAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fomo3dOnReLoadAndDistribute)
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
		it.Event = new(Fomo3dOnReLoadAndDistribute)
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
func (it *Fomo3dOnReLoadAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fomo3dOnReLoadAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fomo3dOnReLoadAndDistribute represents a OnReLoadAndDistribute event raised by the Fomo3d contract.
type Fomo3dOnReLoadAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnReLoadAndDistribute is a free log retrieval operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: event onReLoadAndDistribute(address playerAddress, bytes32 playerName, uint256 compressedData, uint256 compressedIDs, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount)
func (_Fomo3d *Fomo3dFilterer) FilterOnReLoadAndDistribute(opts *bind.FilterOpts) (*Fomo3dOnReLoadAndDistributeIterator, error) {

	logs, sub, err := _Fomo3d.contract.FilterLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return &Fomo3dOnReLoadAndDistributeIterator{contract: _Fomo3d.contract, event: "onReLoadAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnReLoadAndDistribute is a free log subscription operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: event onReLoadAndDistribute(address playerAddress, bytes32 playerName, uint256 compressedData, uint256 compressedIDs, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount)
func (_Fomo3d *Fomo3dFilterer) WatchOnReLoadAndDistribute(opts *bind.WatchOpts, sink chan<- *Fomo3dOnReLoadAndDistribute) (event.Subscription, error) {

	logs, sub, err := _Fomo3d.contract.WatchLogs(opts, "onReLoadAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fomo3dOnReLoadAndDistribute)
				if err := _Fomo3d.contract.UnpackLog(event, "onReLoadAndDistribute", log); err != nil {
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

// ParseOnReLoadAndDistribute is a log parse operation binding the contract event 0x88261ac70d02d5ea73e54fa6da17043c974de1021109573ec1f6f57111c823dd.
//
// Solidity: event onReLoadAndDistribute(address playerAddress, bytes32 playerName, uint256 compressedData, uint256 compressedIDs, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount)
func (_Fomo3d *Fomo3dFilterer) ParseOnReLoadAndDistribute(log types.Log) (*Fomo3dOnReLoadAndDistribute, error) {
	event := new(Fomo3dOnReLoadAndDistribute)
	if err := _Fomo3d.contract.UnpackLog(event, "onReLoadAndDistribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fomo3dOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the Fomo3d contract.
type Fomo3dOnWithdrawIterator struct {
	Event *Fomo3dOnWithdraw // Event containing the contract specifics and raw log

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
func (it *Fomo3dOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fomo3dOnWithdraw)
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
		it.Event = new(Fomo3dOnWithdraw)
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
func (it *Fomo3dOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fomo3dOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fomo3dOnWithdraw represents a OnWithdraw event raised by the Fomo3d contract.
type Fomo3dOnWithdraw struct {
	PlayerID      *big.Int
	PlayerAddress common.Address
	PlayerName    [32]byte
	EthOut        *big.Int
	TimeStamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: event onWithdraw(uint256 indexed playerID, address playerAddress, bytes32 playerName, uint256 ethOut, uint256 timeStamp)
func (_Fomo3d *Fomo3dFilterer) FilterOnWithdraw(opts *bind.FilterOpts, playerID []*big.Int) (*Fomo3dOnWithdrawIterator, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _Fomo3d.contract.FilterLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return &Fomo3dOnWithdrawIterator{contract: _Fomo3d.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: event onWithdraw(uint256 indexed playerID, address playerAddress, bytes32 playerName, uint256 ethOut, uint256 timeStamp)
func (_Fomo3d *Fomo3dFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *Fomo3dOnWithdraw, playerID []*big.Int) (event.Subscription, error) {

	var playerIDRule []interface{}
	for _, playerIDItem := range playerID {
		playerIDRule = append(playerIDRule, playerIDItem)
	}

	logs, sub, err := _Fomo3d.contract.WatchLogs(opts, "onWithdraw", playerIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fomo3dOnWithdraw)
				if err := _Fomo3d.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// ParseOnWithdraw is a log parse operation binding the contract event 0x8f36579a548bc439baa172a6521207464154da77f411e2da3db2f53affe6cc3a.
//
// Solidity: event onWithdraw(uint256 indexed playerID, address playerAddress, bytes32 playerName, uint256 ethOut, uint256 timeStamp)
func (_Fomo3d *Fomo3dFilterer) ParseOnWithdraw(log types.Log) (*Fomo3dOnWithdraw, error) {
	event := new(Fomo3dOnWithdraw)
	if err := _Fomo3d.contract.UnpackLog(event, "onWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fomo3dOnWithdrawAndDistributeIterator is returned from FilterOnWithdrawAndDistribute and is used to iterate over the raw logs and unpacked data for OnWithdrawAndDistribute events raised by the Fomo3d contract.
type Fomo3dOnWithdrawAndDistributeIterator struct {
	Event *Fomo3dOnWithdrawAndDistribute // Event containing the contract specifics and raw log

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
func (it *Fomo3dOnWithdrawAndDistributeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fomo3dOnWithdrawAndDistribute)
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
		it.Event = new(Fomo3dOnWithdrawAndDistribute)
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
func (it *Fomo3dOnWithdrawAndDistributeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fomo3dOnWithdrawAndDistributeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fomo3dOnWithdrawAndDistribute represents a OnWithdrawAndDistribute event raised by the Fomo3d contract.
type Fomo3dOnWithdrawAndDistribute struct {
	PlayerAddress  common.Address
	PlayerName     [32]byte
	EthOut         *big.Int
	CompressedData *big.Int
	CompressedIDs  *big.Int
	WinnerAddr     common.Address
	WinnerName     [32]byte
	AmountWon      *big.Int
	NewPot         *big.Int
	P3DAmount      *big.Int
	GenAmount      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOnWithdrawAndDistribute is a free log retrieval operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: event onWithdrawAndDistribute(address playerAddress, bytes32 playerName, uint256 ethOut, uint256 compressedData, uint256 compressedIDs, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount)
func (_Fomo3d *Fomo3dFilterer) FilterOnWithdrawAndDistribute(opts *bind.FilterOpts) (*Fomo3dOnWithdrawAndDistributeIterator, error) {

	logs, sub, err := _Fomo3d.contract.FilterLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return &Fomo3dOnWithdrawAndDistributeIterator{contract: _Fomo3d.contract, event: "onWithdrawAndDistribute", logs: logs, sub: sub}, nil
}

// WatchOnWithdrawAndDistribute is a free log subscription operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: event onWithdrawAndDistribute(address playerAddress, bytes32 playerName, uint256 ethOut, uint256 compressedData, uint256 compressedIDs, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount)
func (_Fomo3d *Fomo3dFilterer) WatchOnWithdrawAndDistribute(opts *bind.WatchOpts, sink chan<- *Fomo3dOnWithdrawAndDistribute) (event.Subscription, error) {

	logs, sub, err := _Fomo3d.contract.WatchLogs(opts, "onWithdrawAndDistribute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fomo3dOnWithdrawAndDistribute)
				if err := _Fomo3d.contract.UnpackLog(event, "onWithdrawAndDistribute", log); err != nil {
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

// ParseOnWithdrawAndDistribute is a log parse operation binding the contract event 0x0bd0dba8ab932212fa78150cdb7b0275da72e255875967b5cad11464cf71bedc.
//
// Solidity: event onWithdrawAndDistribute(address playerAddress, bytes32 playerName, uint256 ethOut, uint256 compressedData, uint256 compressedIDs, address winnerAddr, bytes32 winnerName, uint256 amountWon, uint256 newPot, uint256 P3DAmount, uint256 genAmount)
func (_Fomo3d *Fomo3dFilterer) ParseOnWithdrawAndDistribute(log types.Log) (*Fomo3dOnWithdrawAndDistribute, error) {
	event := new(Fomo3dOnWithdrawAndDistribute)
	if err := _Fomo3d.contract.UnpackLog(event, "onWithdrawAndDistribute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
