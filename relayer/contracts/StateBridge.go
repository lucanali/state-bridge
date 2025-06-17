// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// BeaconBlockHeader is an auto generated low-level Go binding around an user-defined struct.
type BeaconBlockHeader struct {
	Slot          uint64
	ProposerIndex uint64
	ParentRoot    [32]byte
	StateRoot     [32]byte
	BodyRoot      [32]byte
}

// StateBridgeBlockData is an auto generated low-level Go binding around an user-defined struct.
type StateBridgeBlockData struct {
	ParentHash       [32]byte
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	ReceiptsRoot     [32]byte
	Timestamp        *big.Int
	Number           *big.Int
	ProposerV        uint8
	ProposerR        [32]byte
	ProposerS        [32]byte
}

// StateBridgeBlockUpdate is an auto generated low-level Go binding around an user-defined struct.
type StateBridgeBlockUpdate struct {
	BlockHash          [32]byte
	ChallengeTimestamp *big.Int
	Proposer           common.Address
	Challenged         bool
	IsCritical         bool
	ExecutionStateRoot [32]byte
	BlockData          StateBridgeBlockData
}

// StateBridgeFraudProof is an auto generated low-level Go binding around an user-defined struct.
type StateBridgeFraudProof struct {
	CorrectBlockHash   [32]byte
	ExecutionStateRoot [32]byte
	Proof              [][]byte
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"genesisValidatorsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"forkVersion\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isCritical\",\"type\":\"bool\"}],\"name\":\"BlockSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"Challenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TreasuryWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ValidatorSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SLASH_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"correctBlockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executionStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"proof\",\"type\":\"bytes[]\"}],\"internalType\":\"structStateBridge.FraudProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"challengeBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getUpdate\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCritical\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"executionStateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"proposerV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"proposerR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"proposerS\",\"type\":\"bytes32\"}],\"internalType\":\"structStateBridge.BlockData\",\"name\":\"blockData\",\"type\":\"tuple\"}],\"internalType\":\"structStateBridge.BlockUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"headers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"proposerIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bodyRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structBeaconBlockHeader\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lightClientState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"genesisValidatorsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"defaultForkVersion\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"head\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractValidatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"period\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"syncCommitteeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"proof\",\"type\":\"bytes[]\"}],\"name\":\"setSyncCommittee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"executionStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isCritical\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"proposerV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"proposerR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"proposerS\",\"type\":\"bytes32\"}],\"internalType\":\"structStateBridge.BlockData\",\"name\":\"blockData\",\"type\":\"tuple\"}],\"name\":\"submitBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"syncCommitteeRootByPeriod\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"updates\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"challengeTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"challenged\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCritical\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"executionStateRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"transactionsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receiptsRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"proposerV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"proposerR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"proposerS\",\"type\":\"bytes32\"}],\"internalType\":\"structStateBridge.BlockData\",\"name\":\"blockData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawTreasuryBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// SLASHAMOUNT is a free data retrieval call binding the contract method 0x37720606.
//
// Solidity: function SLASH_AMOUNT() view returns(uint256)
func (_Contracts *ContractsCaller) SLASHAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "SLASH_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLASHAMOUNT is a free data retrieval call binding the contract method 0x37720606.
//
// Solidity: function SLASH_AMOUNT() view returns(uint256)
func (_Contracts *ContractsSession) SLASHAMOUNT() (*big.Int, error) {
	return _Contracts.Contract.SLASHAMOUNT(&_Contracts.CallOpts)
}

// SLASHAMOUNT is a free data retrieval call binding the contract method 0x37720606.
//
// Solidity: function SLASH_AMOUNT() view returns(uint256)
func (_Contracts *ContractsCallerSession) SLASHAMOUNT() (*big.Int, error) {
	return _Contracts.Contract.SLASHAMOUNT(&_Contracts.CallOpts)
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() view returns(uint256)
func (_Contracts *ContractsCaller) BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "blockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() view returns(uint256)
func (_Contracts *ContractsSession) BlockNumber() (*big.Int, error) {
	return _Contracts.Contract.BlockNumber(&_Contracts.CallOpts)
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() view returns(uint256)
func (_Contracts *ContractsCallerSession) BlockNumber() (*big.Int, error) {
	return _Contracts.Contract.BlockNumber(&_Contracts.CallOpts)
}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 _blockNumber) view returns((bytes32,uint256,address,bool,bool,bytes32,(bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint8,bytes32,bytes32)))
func (_Contracts *ContractsCaller) GetUpdate(opts *bind.CallOpts, _blockNumber *big.Int) (StateBridgeBlockUpdate, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getUpdate", _blockNumber)

	if err != nil {
		return *new(StateBridgeBlockUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(StateBridgeBlockUpdate)).(*StateBridgeBlockUpdate)

	return out0, err

}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 _blockNumber) view returns((bytes32,uint256,address,bool,bool,bytes32,(bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint8,bytes32,bytes32)))
func (_Contracts *ContractsSession) GetUpdate(_blockNumber *big.Int) (StateBridgeBlockUpdate, error) {
	return _Contracts.Contract.GetUpdate(&_Contracts.CallOpts, _blockNumber)
}

// GetUpdate is a free data retrieval call binding the contract method 0x32cb25be.
//
// Solidity: function getUpdate(uint256 _blockNumber) view returns((bytes32,uint256,address,bool,bool,bytes32,(bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint8,bytes32,bytes32)))
func (_Contracts *ContractsCallerSession) GetUpdate(_blockNumber *big.Int) (StateBridgeBlockUpdate, error) {
	return _Contracts.Contract.GetUpdate(&_Contracts.CallOpts, _blockNumber)
}

// Headers is a free data retrieval call binding the contract method 0x68b1487c.
//
// Solidity: function headers(uint64 slot) view returns((uint64,uint64,bytes32,bytes32,bytes32))
func (_Contracts *ContractsCaller) Headers(opts *bind.CallOpts, slot uint64) (BeaconBlockHeader, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "headers", slot)

	if err != nil {
		return *new(BeaconBlockHeader), err
	}

	out0 := *abi.ConvertType(out[0], new(BeaconBlockHeader)).(*BeaconBlockHeader)

	return out0, err

}

// Headers is a free data retrieval call binding the contract method 0x68b1487c.
//
// Solidity: function headers(uint64 slot) view returns((uint64,uint64,bytes32,bytes32,bytes32))
func (_Contracts *ContractsSession) Headers(slot uint64) (BeaconBlockHeader, error) {
	return _Contracts.Contract.Headers(&_Contracts.CallOpts, slot)
}

// Headers is a free data retrieval call binding the contract method 0x68b1487c.
//
// Solidity: function headers(uint64 slot) view returns((uint64,uint64,bytes32,bytes32,bytes32))
func (_Contracts *ContractsCallerSession) Headers(slot uint64) (BeaconBlockHeader, error) {
	return _Contracts.Contract.Headers(&_Contracts.CallOpts, slot)
}

// LastBlockNumber is a free data retrieval call binding the contract method 0x2552317c.
//
// Solidity: function lastBlockNumber() view returns(uint256)
func (_Contracts *ContractsCaller) LastBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "lastBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBlockNumber is a free data retrieval call binding the contract method 0x2552317c.
//
// Solidity: function lastBlockNumber() view returns(uint256)
func (_Contracts *ContractsSession) LastBlockNumber() (*big.Int, error) {
	return _Contracts.Contract.LastBlockNumber(&_Contracts.CallOpts)
}

// LastBlockNumber is a free data retrieval call binding the contract method 0x2552317c.
//
// Solidity: function lastBlockNumber() view returns(uint256)
func (_Contracts *ContractsCallerSession) LastBlockNumber() (*big.Int, error) {
	return _Contracts.Contract.LastBlockNumber(&_Contracts.CallOpts)
}

// LightClientState is a free data retrieval call binding the contract method 0x1e3b74aa.
//
// Solidity: function lightClientState() view returns(bytes32 genesisValidatorsRoot, uint256 genesisTime, uint256 secondsPerSlot, bytes4 defaultForkVersion, uint64 head)
func (_Contracts *ContractsCaller) LightClientState(opts *bind.CallOpts) (struct {
	GenesisValidatorsRoot [32]byte
	GenesisTime           *big.Int
	SecondsPerSlot        *big.Int
	DefaultForkVersion    [4]byte
	Head                  uint64
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "lightClientState")

	outstruct := new(struct {
		GenesisValidatorsRoot [32]byte
		GenesisTime           *big.Int
		SecondsPerSlot        *big.Int
		DefaultForkVersion    [4]byte
		Head                  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GenesisValidatorsRoot = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.GenesisTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerSlot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DefaultForkVersion = *abi.ConvertType(out[3], new([4]byte)).(*[4]byte)
	outstruct.Head = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

// LightClientState is a free data retrieval call binding the contract method 0x1e3b74aa.
//
// Solidity: function lightClientState() view returns(bytes32 genesisValidatorsRoot, uint256 genesisTime, uint256 secondsPerSlot, bytes4 defaultForkVersion, uint64 head)
func (_Contracts *ContractsSession) LightClientState() (struct {
	GenesisValidatorsRoot [32]byte
	GenesisTime           *big.Int
	SecondsPerSlot        *big.Int
	DefaultForkVersion    [4]byte
	Head                  uint64
}, error) {
	return _Contracts.Contract.LightClientState(&_Contracts.CallOpts)
}

// LightClientState is a free data retrieval call binding the contract method 0x1e3b74aa.
//
// Solidity: function lightClientState() view returns(bytes32 genesisValidatorsRoot, uint256 genesisTime, uint256 secondsPerSlot, bytes4 defaultForkVersion, uint64 head)
func (_Contracts *ContractsCallerSession) LightClientState() (struct {
	GenesisValidatorsRoot [32]byte
	GenesisTime           *big.Int
	SecondsPerSlot        *big.Int
	DefaultForkVersion    [4]byte
	Head                  uint64
}, error) {
	return _Contracts.Contract.LightClientState(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Contracts *ContractsCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Contracts *ContractsSession) Registry() (common.Address, error) {
	return _Contracts.Contract.Registry(&_Contracts.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Contracts *ContractsCallerSession) Registry() (common.Address, error) {
	return _Contracts.Contract.Registry(&_Contracts.CallOpts)
}

// SyncCommitteeRootByPeriod is a free data retrieval call binding the contract method 0x88501d39.
//
// Solidity: function syncCommitteeRootByPeriod(uint256 period) view returns(bytes32)
func (_Contracts *ContractsCaller) SyncCommitteeRootByPeriod(opts *bind.CallOpts, period *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "syncCommitteeRootByPeriod", period)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SyncCommitteeRootByPeriod is a free data retrieval call binding the contract method 0x88501d39.
//
// Solidity: function syncCommitteeRootByPeriod(uint256 period) view returns(bytes32)
func (_Contracts *ContractsSession) SyncCommitteeRootByPeriod(period *big.Int) ([32]byte, error) {
	return _Contracts.Contract.SyncCommitteeRootByPeriod(&_Contracts.CallOpts, period)
}

// SyncCommitteeRootByPeriod is a free data retrieval call binding the contract method 0x88501d39.
//
// Solidity: function syncCommitteeRootByPeriod(uint256 period) view returns(bytes32)
func (_Contracts *ContractsCallerSession) SyncCommitteeRootByPeriod(period *big.Int) ([32]byte, error) {
	return _Contracts.Contract.SyncCommitteeRootByPeriod(&_Contracts.CallOpts, period)
}

// TreasuryBalance is a free data retrieval call binding the contract method 0x313dab20.
//
// Solidity: function treasuryBalance() view returns(uint256)
func (_Contracts *ContractsCaller) TreasuryBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "treasuryBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryBalance is a free data retrieval call binding the contract method 0x313dab20.
//
// Solidity: function treasuryBalance() view returns(uint256)
func (_Contracts *ContractsSession) TreasuryBalance() (*big.Int, error) {
	return _Contracts.Contract.TreasuryBalance(&_Contracts.CallOpts)
}

// TreasuryBalance is a free data retrieval call binding the contract method 0x313dab20.
//
// Solidity: function treasuryBalance() view returns(uint256)
func (_Contracts *ContractsCallerSession) TreasuryBalance() (*big.Int, error) {
	return _Contracts.Contract.TreasuryBalance(&_Contracts.CallOpts)
}

// UpdateDelay is a free data retrieval call binding the contract method 0x554f94db.
//
// Solidity: function updateDelay() view returns(uint256)
func (_Contracts *ContractsCaller) UpdateDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "updateDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateDelay is a free data retrieval call binding the contract method 0x554f94db.
//
// Solidity: function updateDelay() view returns(uint256)
func (_Contracts *ContractsSession) UpdateDelay() (*big.Int, error) {
	return _Contracts.Contract.UpdateDelay(&_Contracts.CallOpts)
}

// UpdateDelay is a free data retrieval call binding the contract method 0x554f94db.
//
// Solidity: function updateDelay() view returns(uint256)
func (_Contracts *ContractsCallerSession) UpdateDelay() (*big.Int, error) {
	return _Contracts.Contract.UpdateDelay(&_Contracts.CallOpts)
}

// Updates is a free data retrieval call binding the contract method 0xb4c2f727.
//
// Solidity: function updates(uint256 ) view returns(bytes32 blockHash, uint256 challengeTimestamp, address proposer, bool challenged, bool isCritical, bytes32 executionStateRoot, (bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint8,bytes32,bytes32) blockData)
func (_Contracts *ContractsCaller) Updates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BlockHash          [32]byte
	ChallengeTimestamp *big.Int
	Proposer           common.Address
	Challenged         bool
	IsCritical         bool
	ExecutionStateRoot [32]byte
	BlockData          StateBridgeBlockData
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "updates", arg0)

	outstruct := new(struct {
		BlockHash          [32]byte
		ChallengeTimestamp *big.Int
		Proposer           common.Address
		Challenged         bool
		IsCritical         bool
		ExecutionStateRoot [32]byte
		BlockData          StateBridgeBlockData
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ChallengeTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Proposer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Challenged = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.IsCritical = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.ExecutionStateRoot = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.BlockData = *abi.ConvertType(out[6], new(StateBridgeBlockData)).(*StateBridgeBlockData)

	return *outstruct, err

}

// Updates is a free data retrieval call binding the contract method 0xb4c2f727.
//
// Solidity: function updates(uint256 ) view returns(bytes32 blockHash, uint256 challengeTimestamp, address proposer, bool challenged, bool isCritical, bytes32 executionStateRoot, (bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint8,bytes32,bytes32) blockData)
func (_Contracts *ContractsSession) Updates(arg0 *big.Int) (struct {
	BlockHash          [32]byte
	ChallengeTimestamp *big.Int
	Proposer           common.Address
	Challenged         bool
	IsCritical         bool
	ExecutionStateRoot [32]byte
	BlockData          StateBridgeBlockData
}, error) {
	return _Contracts.Contract.Updates(&_Contracts.CallOpts, arg0)
}

// Updates is a free data retrieval call binding the contract method 0xb4c2f727.
//
// Solidity: function updates(uint256 ) view returns(bytes32 blockHash, uint256 challengeTimestamp, address proposer, bool challenged, bool isCritical, bytes32 executionStateRoot, (bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint8,bytes32,bytes32) blockData)
func (_Contracts *ContractsCallerSession) Updates(arg0 *big.Int) (struct {
	BlockHash          [32]byte
	ChallengeTimestamp *big.Int
	Proposer           common.Address
	Challenged         bool
	IsCritical         bool
	ExecutionStateRoot [32]byte
	BlockData          StateBridgeBlockData
}, error) {
	return _Contracts.Contract.Updates(&_Contracts.CallOpts, arg0)
}

// ChallengeBlock is a paid mutator transaction binding the contract method 0x5d3af8b5.
//
// Solidity: function challengeBlock(uint256 _blockNumber, (bytes32,bytes32,bytes[]) proof) returns()
func (_Contracts *ContractsTransactor) ChallengeBlock(opts *bind.TransactOpts, _blockNumber *big.Int, proof StateBridgeFraudProof) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "challengeBlock", _blockNumber, proof)
}

// ChallengeBlock is a paid mutator transaction binding the contract method 0x5d3af8b5.
//
// Solidity: function challengeBlock(uint256 _blockNumber, (bytes32,bytes32,bytes[]) proof) returns()
func (_Contracts *ContractsSession) ChallengeBlock(_blockNumber *big.Int, proof StateBridgeFraudProof) (*types.Transaction, error) {
	return _Contracts.Contract.ChallengeBlock(&_Contracts.TransactOpts, _blockNumber, proof)
}

// ChallengeBlock is a paid mutator transaction binding the contract method 0x5d3af8b5.
//
// Solidity: function challengeBlock(uint256 _blockNumber, (bytes32,bytes32,bytes[]) proof) returns()
func (_Contracts *ContractsTransactorSession) ChallengeBlock(_blockNumber *big.Int, proof StateBridgeFraudProof) (*types.Transaction, error) {
	return _Contracts.Contract.ChallengeBlock(&_Contracts.TransactOpts, _blockNumber, proof)
}

// SetSyncCommittee is a paid mutator transaction binding the contract method 0xac47b452.
//
// Solidity: function setSyncCommittee(uint64 period, bytes32 syncCommitteeRoot, bytes[] proof) returns()
func (_Contracts *ContractsTransactor) SetSyncCommittee(opts *bind.TransactOpts, period uint64, syncCommitteeRoot [32]byte, proof [][]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setSyncCommittee", period, syncCommitteeRoot, proof)
}

// SetSyncCommittee is a paid mutator transaction binding the contract method 0xac47b452.
//
// Solidity: function setSyncCommittee(uint64 period, bytes32 syncCommitteeRoot, bytes[] proof) returns()
func (_Contracts *ContractsSession) SetSyncCommittee(period uint64, syncCommitteeRoot [32]byte, proof [][]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetSyncCommittee(&_Contracts.TransactOpts, period, syncCommitteeRoot, proof)
}

// SetSyncCommittee is a paid mutator transaction binding the contract method 0xac47b452.
//
// Solidity: function setSyncCommittee(uint64 period, bytes32 syncCommitteeRoot, bytes[] proof) returns()
func (_Contracts *ContractsTransactorSession) SetSyncCommittee(period uint64, syncCommitteeRoot [32]byte, proof [][]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetSyncCommittee(&_Contracts.TransactOpts, period, syncCommitteeRoot, proof)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xa1d68b0c.
//
// Solidity: function submitBlock(bytes32 blockHash, bytes32 executionStateRoot, bool isCritical, (bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint8,bytes32,bytes32) blockData) returns()
func (_Contracts *ContractsTransactor) SubmitBlock(opts *bind.TransactOpts, blockHash [32]byte, executionStateRoot [32]byte, isCritical bool, blockData StateBridgeBlockData) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "submitBlock", blockHash, executionStateRoot, isCritical, blockData)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xa1d68b0c.
//
// Solidity: function submitBlock(bytes32 blockHash, bytes32 executionStateRoot, bool isCritical, (bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint8,bytes32,bytes32) blockData) returns()
func (_Contracts *ContractsSession) SubmitBlock(blockHash [32]byte, executionStateRoot [32]byte, isCritical bool, blockData StateBridgeBlockData) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitBlock(&_Contracts.TransactOpts, blockHash, executionStateRoot, isCritical, blockData)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xa1d68b0c.
//
// Solidity: function submitBlock(bytes32 blockHash, bytes32 executionStateRoot, bool isCritical, (bytes32,bytes32,bytes32,bytes32,uint256,uint256,uint8,bytes32,bytes32) blockData) returns()
func (_Contracts *ContractsTransactorSession) SubmitBlock(blockHash [32]byte, executionStateRoot [32]byte, isCritical bool, blockData StateBridgeBlockData) (*types.Transaction, error) {
	return _Contracts.Contract.SubmitBlock(&_Contracts.TransactOpts, blockHash, executionStateRoot, isCritical, blockData)
}

// WithdrawTreasuryBalance is a paid mutator transaction binding the contract method 0xb6b10055.
//
// Solidity: function withdrawTreasuryBalance(address to) returns()
func (_Contracts *ContractsTransactor) WithdrawTreasuryBalance(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawTreasuryBalance", to)
}

// WithdrawTreasuryBalance is a paid mutator transaction binding the contract method 0xb6b10055.
//
// Solidity: function withdrawTreasuryBalance(address to) returns()
func (_Contracts *ContractsSession) WithdrawTreasuryBalance(to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawTreasuryBalance(&_Contracts.TransactOpts, to)
}

// WithdrawTreasuryBalance is a paid mutator transaction binding the contract method 0xb6b10055.
//
// Solidity: function withdrawTreasuryBalance(address to) returns()
func (_Contracts *ContractsTransactorSession) WithdrawTreasuryBalance(to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawTreasuryBalance(&_Contracts.TransactOpts, to)
}

// ContractsBlockSubmittedIterator is returned from FilterBlockSubmitted and is used to iterate over the raw logs and unpacked data for BlockSubmitted events raised by the Contracts contract.
type ContractsBlockSubmittedIterator struct {
	Event *ContractsBlockSubmitted // Event containing the contract specifics and raw log

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
func (it *ContractsBlockSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBlockSubmitted)
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
		it.Event = new(ContractsBlockSubmitted)
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
func (it *ContractsBlockSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBlockSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBlockSubmitted represents a BlockSubmitted event raised by the Contracts contract.
type ContractsBlockSubmitted struct {
	Proposer    common.Address
	BlockNumber *big.Int
	IsCritical  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmitted is a free log retrieval operation binding the contract event 0xd39a24ef9ee0e524a005f0106361e3ef125edbdedec93102932e7756c8f8f257.
//
// Solidity: event BlockSubmitted(address indexed proposer, uint256 blockNumber, bool isCritical)
func (_Contracts *ContractsFilterer) FilterBlockSubmitted(opts *bind.FilterOpts, proposer []common.Address) (*ContractsBlockSubmittedIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "BlockSubmitted", proposerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsBlockSubmittedIterator{contract: _Contracts.contract, event: "BlockSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockSubmitted is a free log subscription operation binding the contract event 0xd39a24ef9ee0e524a005f0106361e3ef125edbdedec93102932e7756c8f8f257.
//
// Solidity: event BlockSubmitted(address indexed proposer, uint256 blockNumber, bool isCritical)
func (_Contracts *ContractsFilterer) WatchBlockSubmitted(opts *bind.WatchOpts, sink chan<- *ContractsBlockSubmitted, proposer []common.Address) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "BlockSubmitted", proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBlockSubmitted)
				if err := _Contracts.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
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

// ParseBlockSubmitted is a log parse operation binding the contract event 0xd39a24ef9ee0e524a005f0106361e3ef125edbdedec93102932e7756c8f8f257.
//
// Solidity: event BlockSubmitted(address indexed proposer, uint256 blockNumber, bool isCritical)
func (_Contracts *ContractsFilterer) ParseBlockSubmitted(log types.Log) (*ContractsBlockSubmitted, error) {
	event := new(ContractsBlockSubmitted)
	if err := _Contracts.contract.UnpackLog(event, "BlockSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsChallengedIterator is returned from FilterChallenged and is used to iterate over the raw logs and unpacked data for Challenged events raised by the Contracts contract.
type ContractsChallengedIterator struct {
	Event *ContractsChallenged // Event containing the contract specifics and raw log

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
func (it *ContractsChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsChallenged)
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
		it.Event = new(ContractsChallenged)
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
func (it *ContractsChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsChallenged represents a Challenged event raised by the Contracts contract.
type ContractsChallenged struct {
	BlockNumber *big.Int
	Challenger  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallenged is a free log retrieval operation binding the contract event 0xcfe09ca25f55d949baba5e280f5750c9ba4b9048fca5532f916067d433afe4d7.
//
// Solidity: event Challenged(uint256 indexed blockNumber, address indexed challenger)
func (_Contracts *ContractsFilterer) FilterChallenged(opts *bind.FilterOpts, blockNumber []*big.Int, challenger []common.Address) (*ContractsChallengedIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Challenged", blockNumberRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsChallengedIterator{contract: _Contracts.contract, event: "Challenged", logs: logs, sub: sub}, nil
}

// WatchChallenged is a free log subscription operation binding the contract event 0xcfe09ca25f55d949baba5e280f5750c9ba4b9048fca5532f916067d433afe4d7.
//
// Solidity: event Challenged(uint256 indexed blockNumber, address indexed challenger)
func (_Contracts *ContractsFilterer) WatchChallenged(opts *bind.WatchOpts, sink chan<- *ContractsChallenged, blockNumber []*big.Int, challenger []common.Address) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Challenged", blockNumberRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsChallenged)
				if err := _Contracts.contract.UnpackLog(event, "Challenged", log); err != nil {
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

// ParseChallenged is a log parse operation binding the contract event 0xcfe09ca25f55d949baba5e280f5750c9ba4b9048fca5532f916067d433afe4d7.
//
// Solidity: event Challenged(uint256 indexed blockNumber, address indexed challenger)
func (_Contracts *ContractsFilterer) ParseChallenged(log types.Log) (*ContractsChallenged, error) {
	event := new(ContractsChallenged)
	if err := _Contracts.contract.UnpackLog(event, "Challenged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Contracts contract.
type ContractsSlashedIterator struct {
	Event *ContractsSlashed // Event containing the contract specifics and raw log

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
func (it *ContractsSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSlashed)
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
		it.Event = new(ContractsSlashed)
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
func (it *ContractsSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSlashed represents a Slashed event raised by the Contracts contract.
type ContractsSlashed struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed validator, uint256 amount)
func (_Contracts *ContractsFilterer) FilterSlashed(opts *bind.FilterOpts, validator []common.Address) (*ContractsSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Slashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsSlashedIterator{contract: _Contracts.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed validator, uint256 amount)
func (_Contracts *ContractsFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *ContractsSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Slashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSlashed)
				if err := _Contracts.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed validator, uint256 amount)
func (_Contracts *ContractsFilterer) ParseSlashed(log types.Log) (*ContractsSlashed, error) {
	event := new(ContractsSlashed)
	if err := _Contracts.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTreasuryWithdrawnIterator is returned from FilterTreasuryWithdrawn and is used to iterate over the raw logs and unpacked data for TreasuryWithdrawn events raised by the Contracts contract.
type ContractsTreasuryWithdrawnIterator struct {
	Event *ContractsTreasuryWithdrawn // Event containing the contract specifics and raw log

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
func (it *ContractsTreasuryWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTreasuryWithdrawn)
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
		it.Event = new(ContractsTreasuryWithdrawn)
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
func (it *ContractsTreasuryWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTreasuryWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTreasuryWithdrawn represents a TreasuryWithdrawn event raised by the Contracts contract.
type ContractsTreasuryWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTreasuryWithdrawn is a free log retrieval operation binding the contract event 0x41fdd680478135993bc53fb2ffaf9560951b57ef62ff6badd02b61e018b4f17f.
//
// Solidity: event TreasuryWithdrawn(address indexed to, uint256 amount)
func (_Contracts *ContractsFilterer) FilterTreasuryWithdrawn(opts *bind.FilterOpts, to []common.Address) (*ContractsTreasuryWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "TreasuryWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTreasuryWithdrawnIterator{contract: _Contracts.contract, event: "TreasuryWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTreasuryWithdrawn is a free log subscription operation binding the contract event 0x41fdd680478135993bc53fb2ffaf9560951b57ef62ff6badd02b61e018b4f17f.
//
// Solidity: event TreasuryWithdrawn(address indexed to, uint256 amount)
func (_Contracts *ContractsFilterer) WatchTreasuryWithdrawn(opts *bind.WatchOpts, sink chan<- *ContractsTreasuryWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "TreasuryWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTreasuryWithdrawn)
				if err := _Contracts.contract.UnpackLog(event, "TreasuryWithdrawn", log); err != nil {
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

// ParseTreasuryWithdrawn is a log parse operation binding the contract event 0x41fdd680478135993bc53fb2ffaf9560951b57ef62ff6badd02b61e018b4f17f.
//
// Solidity: event TreasuryWithdrawn(address indexed to, uint256 amount)
func (_Contracts *ContractsFilterer) ParseTreasuryWithdrawn(log types.Log) (*ContractsTreasuryWithdrawn, error) {
	event := new(ContractsTreasuryWithdrawn)
	if err := _Contracts.contract.UnpackLog(event, "TreasuryWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the Contracts contract.
type ContractsValidatorSlashedIterator struct {
	Event *ContractsValidatorSlashed // Event containing the contract specifics and raw log

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
func (it *ContractsValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsValidatorSlashed)
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
		it.Event = new(ContractsValidatorSlashed)
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
func (it *ContractsValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsValidatorSlashed represents a ValidatorSlashed event raised by the Contracts contract.
type ContractsValidatorSlashed struct {
	Validator  common.Address
	Challenger common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0x6c6f87171ccdd4f0f7d6c1a0c543d6eb923914876823d2f8e01fe0988a235ac0.
//
// Solidity: event ValidatorSlashed(address indexed validator, address indexed challenger, uint256 amount)
func (_Contracts *ContractsFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, validator []common.Address, challenger []common.Address) (*ContractsValidatorSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ValidatorSlashed", validatorRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsValidatorSlashedIterator{contract: _Contracts.contract, event: "ValidatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0x6c6f87171ccdd4f0f7d6c1a0c543d6eb923914876823d2f8e01fe0988a235ac0.
//
// Solidity: event ValidatorSlashed(address indexed validator, address indexed challenger, uint256 amount)
func (_Contracts *ContractsFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractsValidatorSlashed, validator []common.Address, challenger []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ValidatorSlashed", validatorRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsValidatorSlashed)
				if err := _Contracts.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
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

// ParseValidatorSlashed is a log parse operation binding the contract event 0x6c6f87171ccdd4f0f7d6c1a0c543d6eb923914876823d2f8e01fe0988a235ac0.
//
// Solidity: event ValidatorSlashed(address indexed validator, address indexed challenger, uint256 amount)
func (_Contracts *ContractsFilterer) ParseValidatorSlashed(log types.Log) (*ContractsValidatorSlashed, error) {
	event := new(ContractsValidatorSlashed)
	if err := _Contracts.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
