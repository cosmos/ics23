// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ics23_sol

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

// BatchEntryData is an auto generated low-level Go binding around an user-defined struct.
type BatchEntryData struct {
	Exist    ExistenceProofData
	Nonexist NonExistenceProofData
}

// BatchProofData is an auto generated low-level Go binding around an user-defined struct.
type BatchProofData struct {
	Entries []BatchEntryData
}

// CommitmentProofData is an auto generated low-level Go binding around an user-defined struct.
type CommitmentProofData struct {
	Exist      ExistenceProofData
	Nonexist   NonExistenceProofData
	Batch      BatchProofData
	Compressed CompressedBatchProofData
}

// CompressedBatchEntryData is an auto generated low-level Go binding around an user-defined struct.
type CompressedBatchEntryData struct {
	Exist    CompressedExistenceProofData
	Nonexist CompressedNonExistenceProofData
}

// CompressedBatchProofData is an auto generated low-level Go binding around an user-defined struct.
type CompressedBatchProofData struct {
	Entries      []CompressedBatchEntryData
	LookupInners []InnerOpData
}

// CompressedExistenceProofData is an auto generated low-level Go binding around an user-defined struct.
type CompressedExistenceProofData struct {
	Key   []byte
	Value []byte
	Leaf  LeafOpData
	Path  []int32
}

// CompressedNonExistenceProofData is an auto generated low-level Go binding around an user-defined struct.
type CompressedNonExistenceProofData struct {
	Key   []byte
	Left  CompressedExistenceProofData
	Right CompressedExistenceProofData
}

// ExistenceProofData is an auto generated low-level Go binding around an user-defined struct.
type ExistenceProofData struct {
	Key   []byte
	Value []byte
	Leaf  LeafOpData
	Path  []InnerOpData
}

// Ics23BatchItem is an auto generated low-level Go binding around an user-defined struct.
type Ics23BatchItem struct {
	Key   []byte
	Value []byte
}

// InnerOpData is an auto generated low-level Go binding around an user-defined struct.
type InnerOpData struct {
	Hash   uint8
	Prefix []byte
	Suffix []byte
}

// InnerSpecData is an auto generated low-level Go binding around an user-defined struct.
type InnerSpecData struct {
	ChildOrder      []int32
	ChildSize       int32
	MinPrefixLength int32
	MaxPrefixLength int32
	EmptyChild      []byte
	Hash            uint8
}

// LeafOpData is an auto generated low-level Go binding around an user-defined struct.
type LeafOpData struct {
	Hash         uint8
	PrehashKey   uint8
	PrehashValue uint8
	Length       uint8
	Prefix       []byte
}

// NonExistenceProofData is an auto generated low-level Go binding around an user-defined struct.
type NonExistenceProofData struct {
	Key   []byte
	Left  ExistenceProofData
	Right ExistenceProofData
}

// ProofSpecData is an auto generated low-level Go binding around an user-defined struct.
type ProofSpecData struct {
	LeafSpec  LeafOpData
	InnerSpec InnerSpecData
	MaxDepth  int32
	MinDepth  int32
}

// BatchEntryABI is the input ABI used to generate the binding from.
const BatchEntryABI = "[]"

// BatchEntryBin is the compiled bytecode used for deploying new contracts.
var BatchEntryBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203742061f080b73860e3b8cc966ddf9bcef8c4712f05f61cbb3da2307bb066b7f64736f6c63430008090033"

// DeployBatchEntry deploys a new Ethereum contract, binding an instance of BatchEntry to it.
func DeployBatchEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BatchEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchEntry{BatchEntryCaller: BatchEntryCaller{contract: contract}, BatchEntryTransactor: BatchEntryTransactor{contract: contract}, BatchEntryFilterer: BatchEntryFilterer{contract: contract}}, nil
}

// BatchEntry is an auto generated Go binding around an Ethereum contract.
type BatchEntry struct {
	BatchEntryCaller     // Read-only binding to the contract
	BatchEntryTransactor // Write-only binding to the contract
	BatchEntryFilterer   // Log filterer for contract events
}

// BatchEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchEntrySession struct {
	Contract     *BatchEntry       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchEntryCallerSession struct {
	Contract *BatchEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BatchEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchEntryTransactorSession struct {
	Contract     *BatchEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BatchEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchEntryRaw struct {
	Contract *BatchEntry // Generic contract binding to access the raw methods on
}

// BatchEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchEntryCallerRaw struct {
	Contract *BatchEntryCaller // Generic read-only contract binding to access the raw methods on
}

// BatchEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchEntryTransactorRaw struct {
	Contract *BatchEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchEntry creates a new instance of BatchEntry, bound to a specific deployed contract.
func NewBatchEntry(address common.Address, backend bind.ContractBackend) (*BatchEntry, error) {
	contract, err := bindBatchEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchEntry{BatchEntryCaller: BatchEntryCaller{contract: contract}, BatchEntryTransactor: BatchEntryTransactor{contract: contract}, BatchEntryFilterer: BatchEntryFilterer{contract: contract}}, nil
}

// NewBatchEntryCaller creates a new read-only instance of BatchEntry, bound to a specific deployed contract.
func NewBatchEntryCaller(address common.Address, caller bind.ContractCaller) (*BatchEntryCaller, error) {
	contract, err := bindBatchEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchEntryCaller{contract: contract}, nil
}

// NewBatchEntryTransactor creates a new write-only instance of BatchEntry, bound to a specific deployed contract.
func NewBatchEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchEntryTransactor, error) {
	contract, err := bindBatchEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchEntryTransactor{contract: contract}, nil
}

// NewBatchEntryFilterer creates a new log filterer instance of BatchEntry, bound to a specific deployed contract.
func NewBatchEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchEntryFilterer, error) {
	contract, err := bindBatchEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchEntryFilterer{contract: contract}, nil
}

// bindBatchEntry binds a generic wrapper to an already deployed contract.
func bindBatchEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchEntry *BatchEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchEntry.Contract.BatchEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchEntry *BatchEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchEntry.Contract.BatchEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchEntry *BatchEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchEntry.Contract.BatchEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchEntry *BatchEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchEntry *BatchEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchEntry *BatchEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchEntry.Contract.contract.Transact(opts, method, params...)
}

// BatchProofABI is the input ABI used to generate the binding from.
const BatchProofABI = "[]"

// BatchProofBin is the compiled bytecode used for deploying new contracts.
var BatchProofBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e3fb702ca13ae2e51a6651e6edeb36d023e514aeb6d1cd7fcefa234323c0174a64736f6c63430008090033"

// DeployBatchProof deploys a new Ethereum contract, binding an instance of BatchProof to it.
func DeployBatchProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchProof, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BatchProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchProof{BatchProofCaller: BatchProofCaller{contract: contract}, BatchProofTransactor: BatchProofTransactor{contract: contract}, BatchProofFilterer: BatchProofFilterer{contract: contract}}, nil
}

// BatchProof is an auto generated Go binding around an Ethereum contract.
type BatchProof struct {
	BatchProofCaller     // Read-only binding to the contract
	BatchProofTransactor // Write-only binding to the contract
	BatchProofFilterer   // Log filterer for contract events
}

// BatchProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchProofSession struct {
	Contract     *BatchProof       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchProofCallerSession struct {
	Contract *BatchProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BatchProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchProofTransactorSession struct {
	Contract     *BatchProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BatchProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchProofRaw struct {
	Contract *BatchProof // Generic contract binding to access the raw methods on
}

// BatchProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchProofCallerRaw struct {
	Contract *BatchProofCaller // Generic read-only contract binding to access the raw methods on
}

// BatchProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchProofTransactorRaw struct {
	Contract *BatchProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchProof creates a new instance of BatchProof, bound to a specific deployed contract.
func NewBatchProof(address common.Address, backend bind.ContractBackend) (*BatchProof, error) {
	contract, err := bindBatchProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchProof{BatchProofCaller: BatchProofCaller{contract: contract}, BatchProofTransactor: BatchProofTransactor{contract: contract}, BatchProofFilterer: BatchProofFilterer{contract: contract}}, nil
}

// NewBatchProofCaller creates a new read-only instance of BatchProof, bound to a specific deployed contract.
func NewBatchProofCaller(address common.Address, caller bind.ContractCaller) (*BatchProofCaller, error) {
	contract, err := bindBatchProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchProofCaller{contract: contract}, nil
}

// NewBatchProofTransactor creates a new write-only instance of BatchProof, bound to a specific deployed contract.
func NewBatchProofTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchProofTransactor, error) {
	contract, err := bindBatchProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchProofTransactor{contract: contract}, nil
}

// NewBatchProofFilterer creates a new log filterer instance of BatchProof, bound to a specific deployed contract.
func NewBatchProofFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchProofFilterer, error) {
	contract, err := bindBatchProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchProofFilterer{contract: contract}, nil
}

// bindBatchProof binds a generic wrapper to an already deployed contract.
func bindBatchProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchProof *BatchProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchProof.Contract.BatchProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchProof *BatchProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchProof.Contract.BatchProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchProof *BatchProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchProof.Contract.BatchProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchProof *BatchProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchProof *BatchProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchProof *BatchProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchProof.Contract.contract.Transact(opts, method, params...)
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e46b091769dbf621d1aee0439bea552bbe168a1df68c77864119a8f1b8db5c4e64736f6c63430008090033"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// CommitmentProofABI is the input ABI used to generate the binding from.
const CommitmentProofABI = "[]"

// CommitmentProofBin is the compiled bytecode used for deploying new contracts.
var CommitmentProofBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220631c74804bd77e4758be67d9ec180585eaf4ee21ec14cb399aea4a47d864648f64736f6c63430008090033"

// DeployCommitmentProof deploys a new Ethereum contract, binding an instance of CommitmentProof to it.
func DeployCommitmentProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CommitmentProof, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitmentProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CommitmentProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CommitmentProof{CommitmentProofCaller: CommitmentProofCaller{contract: contract}, CommitmentProofTransactor: CommitmentProofTransactor{contract: contract}, CommitmentProofFilterer: CommitmentProofFilterer{contract: contract}}, nil
}

// CommitmentProof is an auto generated Go binding around an Ethereum contract.
type CommitmentProof struct {
	CommitmentProofCaller     // Read-only binding to the contract
	CommitmentProofTransactor // Write-only binding to the contract
	CommitmentProofFilterer   // Log filterer for contract events
}

// CommitmentProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type CommitmentProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitmentProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CommitmentProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitmentProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CommitmentProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CommitmentProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CommitmentProofSession struct {
	Contract     *CommitmentProof  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CommitmentProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CommitmentProofCallerSession struct {
	Contract *CommitmentProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CommitmentProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CommitmentProofTransactorSession struct {
	Contract     *CommitmentProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CommitmentProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type CommitmentProofRaw struct {
	Contract *CommitmentProof // Generic contract binding to access the raw methods on
}

// CommitmentProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CommitmentProofCallerRaw struct {
	Contract *CommitmentProofCaller // Generic read-only contract binding to access the raw methods on
}

// CommitmentProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CommitmentProofTransactorRaw struct {
	Contract *CommitmentProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCommitmentProof creates a new instance of CommitmentProof, bound to a specific deployed contract.
func NewCommitmentProof(address common.Address, backend bind.ContractBackend) (*CommitmentProof, error) {
	contract, err := bindCommitmentProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CommitmentProof{CommitmentProofCaller: CommitmentProofCaller{contract: contract}, CommitmentProofTransactor: CommitmentProofTransactor{contract: contract}, CommitmentProofFilterer: CommitmentProofFilterer{contract: contract}}, nil
}

// NewCommitmentProofCaller creates a new read-only instance of CommitmentProof, bound to a specific deployed contract.
func NewCommitmentProofCaller(address common.Address, caller bind.ContractCaller) (*CommitmentProofCaller, error) {
	contract, err := bindCommitmentProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CommitmentProofCaller{contract: contract}, nil
}

// NewCommitmentProofTransactor creates a new write-only instance of CommitmentProof, bound to a specific deployed contract.
func NewCommitmentProofTransactor(address common.Address, transactor bind.ContractTransactor) (*CommitmentProofTransactor, error) {
	contract, err := bindCommitmentProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CommitmentProofTransactor{contract: contract}, nil
}

// NewCommitmentProofFilterer creates a new log filterer instance of CommitmentProof, bound to a specific deployed contract.
func NewCommitmentProofFilterer(address common.Address, filterer bind.ContractFilterer) (*CommitmentProofFilterer, error) {
	contract, err := bindCommitmentProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CommitmentProofFilterer{contract: contract}, nil
}

// bindCommitmentProof binds a generic wrapper to an already deployed contract.
func bindCommitmentProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CommitmentProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitmentProof *CommitmentProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitmentProof.Contract.CommitmentProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitmentProof *CommitmentProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitmentProof.Contract.CommitmentProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitmentProof *CommitmentProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitmentProof.Contract.CommitmentProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CommitmentProof *CommitmentProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CommitmentProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CommitmentProof *CommitmentProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CommitmentProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CommitmentProof *CommitmentProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CommitmentProof.Contract.contract.Transact(opts, method, params...)
}

// CompressABI is the input ABI used to generate the binding from.
const CompressABI = "[]"

// CompressBin is the compiled bytecode used for deploying new contracts.
var CompressBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b471dd530ee39a7a3901cffd90b540c8de0e8be849620b82fd1fab5dd6772f2864736f6c63430008090033"

// DeployCompress deploys a new Ethereum contract, binding an instance of Compress to it.
func DeployCompress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Compress, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Compress{CompressCaller: CompressCaller{contract: contract}, CompressTransactor: CompressTransactor{contract: contract}, CompressFilterer: CompressFilterer{contract: contract}}, nil
}

// Compress is an auto generated Go binding around an Ethereum contract.
type Compress struct {
	CompressCaller     // Read-only binding to the contract
	CompressTransactor // Write-only binding to the contract
	CompressFilterer   // Log filterer for contract events
}

// CompressCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressSession struct {
	Contract     *Compress         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressCallerSession struct {
	Contract *CompressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CompressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressTransactorSession struct {
	Contract     *CompressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CompressRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressRaw struct {
	Contract *Compress // Generic contract binding to access the raw methods on
}

// CompressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressCallerRaw struct {
	Contract *CompressCaller // Generic read-only contract binding to access the raw methods on
}

// CompressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressTransactorRaw struct {
	Contract *CompressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompress creates a new instance of Compress, bound to a specific deployed contract.
func NewCompress(address common.Address, backend bind.ContractBackend) (*Compress, error) {
	contract, err := bindCompress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compress{CompressCaller: CompressCaller{contract: contract}, CompressTransactor: CompressTransactor{contract: contract}, CompressFilterer: CompressFilterer{contract: contract}}, nil
}

// NewCompressCaller creates a new read-only instance of Compress, bound to a specific deployed contract.
func NewCompressCaller(address common.Address, caller bind.ContractCaller) (*CompressCaller, error) {
	contract, err := bindCompress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressCaller{contract: contract}, nil
}

// NewCompressTransactor creates a new write-only instance of Compress, bound to a specific deployed contract.
func NewCompressTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressTransactor, error) {
	contract, err := bindCompress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressTransactor{contract: contract}, nil
}

// NewCompressFilterer creates a new log filterer instance of Compress, bound to a specific deployed contract.
func NewCompressFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressFilterer, error) {
	contract, err := bindCompress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressFilterer{contract: contract}, nil
}

// bindCompress binds a generic wrapper to an already deployed contract.
func bindCompress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compress *CompressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compress.Contract.CompressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compress *CompressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compress.Contract.CompressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compress *CompressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compress.Contract.CompressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compress *CompressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compress *CompressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compress *CompressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compress.Contract.contract.Transact(opts, method, params...)
}

// CompressUnitTestABI is the input ABI used to generate the binding from.
const CompressUnitTestABI = "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structBatchProof.Data\",\"name\":\"batch\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"lookup_inners\",\"type\":\"tuple[]\"}],\"internalType\":\"structCompressedBatchProof.Data\",\"name\":\"compressed\",\"type\":\"tuple\"}],\"internalType\":\"structCommitmentProof.Data\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"decompress\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structBatchProof.Data\",\"name\":\"batch\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"lookup_inners\",\"type\":\"tuple[]\"}],\"internalType\":\"structCompressedBatchProof.Data\",\"name\":\"compressed\",\"type\":\"tuple\"}],\"internalType\":\"structCommitmentProof.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// CompressUnitTestBin is the compiled bytecode used for deploying new contracts.
var CompressUnitTestBin = "0x608060405234801561001057600080fd5b506115b7806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806398d6118314610030575b600080fd5b61004361003e366004610ffa565b610059565b6040516100509190611468565b60405180910390f35b6100616106a0565b600061006c836101c0565b9050806100825761007b6106a0565b9392505050565b805161009357610090610702565b81525b60208101516100aa576100a4610730565b60208201525b60608101516100cd57604080518082019091526060808252602082015260608201525b60005b604082015151518110156101b95761010b82604001516000015182815181106100fb576100fb611542565b6020026020010151600001511590565b1561014157610118610702565b80836040015160000151838151811061013357610133611542565b602090810291909101015152505b61016e826040015160000151828151811061015e5761015e611542565b6020026020010151602001511590565b156101a75761017b610730565b80836040015160000151838151811061019657610196611542565b602002602001015160200181905250505b806101b181611558565b9150506100d0565b5092915050565b6101c86106a0565b6101d58260600151610237565b1515600114156101e3575090565b60405180608001604052806101f6610265565b8152602001610203610273565b81526020016040518060200160405280610220866060015161027b565b9052815260200161022f610340565b905292915050565b8051516000901561024a57506000919050565b6020820151511561025d57506000919050565b506001919050565b61026d610702565b50600090565b61026d610730565b606060008260000151516001600160401b0381111561029c5761029c61079f565b6040519080825280602002602001820160405280156102d557816020015b6102c2610757565b8152602001906001900390816102ba5790505b50905060005b8351518110156101b957610310846000015182815181106102fe576102fe611542565b60200260200101518560200151610359565b82828151811061032257610322611542565b6020026020010181905250808061033890611558565b9150506102db565b604080518082019091526060808252602082015261026d565b610361610757565b825161036c90610406565b61039f57604051806040016040528061038985600001518561043f565b8152602001610396610273565b90529050610400565b60405180604001604052806103b2610265565b8152602001604051806060016040528086602001516000015181526020016103e28760200151602001518761043f565b81526020016103f98760200151604001518761043f565b9052905290505b92915050565b8051516000901561041957506000919050565b6020820151511561042c57506000919050565b6060820151511561025d57506000919050565b610447610702565b61045083610406565b156104645761045d610265565b9050610400565b600060405180608001604052808560000151815260200185602001518152602001856040015181526020018560600151516001600160401b038111156104ac576104ac61079f565b60405190808252806020026020018201604052801561050257816020015b6104ef6040805160608101909152806000815260200160608152602001606081525090565b8152602001906001900390816104ca5790505b509052905060005b8460600151518110156106425760008560600151828151811061052f5761052f611542565b602002602001015160030b121561057e5760405162461bcd60e51b815260206004820152600e60248201526d070726f6f662e70617468203c20360941b60448201526064015b60405180910390fd5b60006105a98660600151838151811061059957610599611542565b602002602001015160030b61064a565b9050845181106105f35760405162461bcd60e51b81526020600482015260156024820152740e6e8cae0407c7a40d8deded6eae05cd8cadccee8d605b1b6044820152606401610575565b84818151811061060557610605611542565b60200260200101518360600151838151811061062357610623611542565b602002602001018190525050808061063a90611558565b91505061050a565b509392505050565b60008082121561069c5760405162461bcd60e51b815260206004820181905260248201527f53616665436173743a2076616c7565206d75737420626520706f7369746976656044820152606401610575565b5090565b60405180608001604052806106b3610702565b81526020016106c0610730565b81526020016106db6040518060200160405280606081525090565b81526020016106fd604051806040016040528060608152602001606081525090565b905290565b60405180608001604052806060815260200160608152602001610723610777565b8152602001606081525090565b60405180606001604052806060815260200161074a610702565b81526020016106fd610702565b604051806040016040528061076a610702565b81526020016106fd610730565b6040805160a08101909152806000815260200160008152602001600081526020016000610723565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156107d7576107d761079f565b60405290565b604051608081016001600160401b03811182821017156107d7576107d761079f565b604051602081016001600160401b03811182821017156107d7576107d761079f565b604080519081016001600160401b03811182821017156107d7576107d761079f565b604051601f8201601f191681016001600160401b038111828210171561086b5761086b61079f565b604052919050565b600082601f83011261088457600080fd5b81356001600160401b0381111561089d5761089d61079f565b6108b0601f8201601f1916602001610843565b8181528460208386010111156108c557600080fd5b816020850160208301376000918101602001919091529392505050565b8035600781106108f157600080fd5b919050565b600060a0828403121561090857600080fd5b60405160a081016001600160401b03828210818311171561092b5761092b61079f565b8160405282935061093b856108e2565b8352610949602086016108e2565b602084015261095a604086016108e2565b6040840152606085013591506009821061097357600080fd5b816060840152608085013591508082111561098d57600080fd5b5061099a85828601610873565b6080830152505092915050565b60006001600160401b038211156109c0576109c061079f565b5060051b60200190565b600082601f8301126109db57600080fd5b813560206109f06109eb836109a7565b610843565b82815260059290921b84018101918181019086841115610a0f57600080fd5b8286015b84811015610abe5780356001600160401b0380821115610a335760008081fd5b908801906060828b03601f1901811315610a4d5760008081fd5b610a556107b5565b610a608885016108e2565b815260408085013584811115610a765760008081fd5b610a848e8b83890101610873565b838b015250918401359183831115610a9c5760008081fd5b610aaa8d8a85880101610873565b908201528652505050918301918301610a13565b509695505050505050565b600060808284031215610adb57600080fd5b610ae36107dd565b905081356001600160401b0380821115610afc57600080fd5b610b0885838601610873565b83526020840135915080821115610b1e57600080fd5b610b2a85838601610873565b60208401526040840135915080821115610b4357600080fd5b610b4f858386016108f6565b60408401526060840135915080821115610b6857600080fd5b50610b75848285016109ca565b60608301525092915050565b600060608284031215610b9357600080fd5b610b9b6107b5565b905081356001600160401b0380821115610bb457600080fd5b610bc085838601610873565b83526020840135915080821115610bd657600080fd5b610be285838601610ac9565b60208401526040840135915080821115610bfb57600080fd5b50610c0884828501610ac9565b60408301525092915050565b60006020808385031215610c2757600080fd5b610c2f6107ff565b915082356001600160401b0380821115610c4857600080fd5b818501915085601f830112610c5c57600080fd5b8135610c6a6109eb826109a7565b81815260059190911b83018401908481019088831115610c8957600080fd5b8585015b83811015610d1c57803585811115610ca55760008081fd5b86016040818c03601f1901811315610cbd5760008081fd5b610cc5610821565b8983013588811115610cd75760008081fd5b610ce58e8c83870101610ac9565b825250908201359087821115610cfb5760008081fd5b610d098d8b84860101610b81565b818b015285525050918601918601610c8d565b50865250939695505050505050565b600082601f830112610d3c57600080fd5b81356020610d4c6109eb836109a7565b82815260059290921b84018101918181019086841115610d6b57600080fd5b8286015b84811015610abe5780358060030b8114610d895760008081fd5b8352918301918301610d6f565b600060808284031215610da857600080fd5b610db06107dd565b905081356001600160401b0380821115610dc957600080fd5b610dd585838601610873565b83526020840135915080821115610deb57600080fd5b610df785838601610873565b60208401526040840135915080821115610e1057600080fd5b610e1c858386016108f6565b60408401526060840135915080821115610e3557600080fd5b50610b7584828501610d2b565b600060408284031215610e5457600080fd5b610e5c610821565b905081356001600160401b0380821115610e7557600080fd5b818401915084601f830112610e8957600080fd5b81356020610e996109eb836109a7565b82815260059290921b84018101918181019088841115610eb857600080fd5b8286015b84811015610fc957803586811115610ed357600080fd5b8701601f196040828d0382011215610eea57600080fd5b610ef2610821565b8683013589811115610f0357600080fd5b610f118e8983870101610d96565b825250604083013589811115610f2657600080fd5b92909201916060838e0383011215610f3d57600080fd5b610f456107b5565b91508683013589811115610f5857600080fd5b610f668e8983870101610873565b835250604083013589811115610f7b57600080fd5b610f898e8983870101610d96565b8884015250606083013589811115610fa057600080fd5b610fae8e8983870101610d96565b60408401525080870191909152845250918301918301610ebc565b5086525085810135935082841115610fe057600080fd5b610fec878588016109ca565b818601525050505092915050565b60006020828403121561100c57600080fd5b81356001600160401b038082111561102357600080fd5b908301906080828603121561103757600080fd5b61103f6107dd565b82358281111561104e57600080fd5b61105a87828601610ac9565b82525060208301358281111561106f57600080fd5b61107b87828601610b81565b60208301525060408301358281111561109357600080fd5b61109f87828601610c14565b6040830152506060830135828111156110b757600080fd5b6110c387828601610e42565b60608301525095945050505050565b6000815180845260005b818110156110f8576020818501810151868301820152016110dc565b8181111561110a576000602083870101525b50601f01601f19169290920160200192915050565b634e487b7160e01b600052602160045260246000fd5b600781106111455761114561111f565b9052565b611154828251611135565b600060208201516111686020850182611135565b50604082015161117b6040850182611135565b506060820151600981106111915761119161111f565b80606085015250608082015160a060808501526111b160a08501826110d2565b949350505050565b600081518084526020808501808196508360051b8101915082860160005b85811015611239578284038952815160606111f3868351611135565b868201518188880152611208828801826110d2565b9150506040808301519250868203818801525061122581836110d2565b9a87019a95505050908401906001016111d7565b5091979650505050505050565b600081516080845261125b60808501826110d2565b90506020830151848203602086015261127482826110d2565b9150506040830151848203604086015261128e8282611149565b915050606083015184820360608601526112a882826111b9565b95945050505050565b60008151606084526112c660608501826110d2565b9050602083015184820360208601526112df8282611246565b915050604083015184820360408601526112a88282611246565b600081516080845261130e60808501826110d2565b90506020808401518583038287015261132783826110d2565b925050604084015185830360408701526113418382611149565b606086810151888303918901919091528051808352908401945060009250908301905b80831015610abe57845160030b8252938301936001929092019190830190611364565b600060408084018351828652818151808452606093508388019150838160051b8901016020808501945060005b83811015611442578a8303605f19018552855180518985526113d88a8601826112f9565b90508382015191508481038486015281518982526113f88a8301826110d2565b9050848301518282038684015261140f82826112f9565b9150508a83015192508181038b83015261142981846112f9565b98850198978501979550505060019190910190506113b4565b50808901519650898203818b01525061145b81876111b9565b9998505050505050505050565b60006020808352835160808285015261148460a0850182611246565b905081850151601f1960408187850301818801526114a284846112b1565b88820151888203840160608a015251868252805187830181905291955086019350600581901b85018201908286019060005b8181101561152457878403603f19018352865180518686526114f887870182611246565b918b0151868303878d015291905061151081836112b1565b988b019895505050918801916001016114d4565b505050606089015195508288820301608089015261145b8187611387565b634e487b7160e01b600052603260045260246000fd5b600060001982141561157a57634e487b7160e01b600052601160045260246000fd5b506001019056fea26469706673582212202c33f051feccb7bdb8f70ee4a1866dd4d222ccc941a996183ec965a69474129f64736f6c63430008090033"

// DeployCompressUnitTest deploys a new Ethereum contract, binding an instance of CompressUnitTest to it.
func DeployCompressUnitTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompressUnitTest, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressUnitTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompressUnitTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompressUnitTest{CompressUnitTestCaller: CompressUnitTestCaller{contract: contract}, CompressUnitTestTransactor: CompressUnitTestTransactor{contract: contract}, CompressUnitTestFilterer: CompressUnitTestFilterer{contract: contract}}, nil
}

// CompressUnitTest is an auto generated Go binding around an Ethereum contract.
type CompressUnitTest struct {
	CompressUnitTestCaller     // Read-only binding to the contract
	CompressUnitTestTransactor // Write-only binding to the contract
	CompressUnitTestFilterer   // Log filterer for contract events
}

// CompressUnitTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressUnitTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressUnitTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressUnitTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressUnitTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressUnitTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressUnitTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressUnitTestSession struct {
	Contract     *CompressUnitTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompressUnitTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressUnitTestCallerSession struct {
	Contract *CompressUnitTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CompressUnitTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressUnitTestTransactorSession struct {
	Contract     *CompressUnitTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CompressUnitTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressUnitTestRaw struct {
	Contract *CompressUnitTest // Generic contract binding to access the raw methods on
}

// CompressUnitTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressUnitTestCallerRaw struct {
	Contract *CompressUnitTestCaller // Generic read-only contract binding to access the raw methods on
}

// CompressUnitTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressUnitTestTransactorRaw struct {
	Contract *CompressUnitTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompressUnitTest creates a new instance of CompressUnitTest, bound to a specific deployed contract.
func NewCompressUnitTest(address common.Address, backend bind.ContractBackend) (*CompressUnitTest, error) {
	contract, err := bindCompressUnitTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompressUnitTest{CompressUnitTestCaller: CompressUnitTestCaller{contract: contract}, CompressUnitTestTransactor: CompressUnitTestTransactor{contract: contract}, CompressUnitTestFilterer: CompressUnitTestFilterer{contract: contract}}, nil
}

// NewCompressUnitTestCaller creates a new read-only instance of CompressUnitTest, bound to a specific deployed contract.
func NewCompressUnitTestCaller(address common.Address, caller bind.ContractCaller) (*CompressUnitTestCaller, error) {
	contract, err := bindCompressUnitTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressUnitTestCaller{contract: contract}, nil
}

// NewCompressUnitTestTransactor creates a new write-only instance of CompressUnitTest, bound to a specific deployed contract.
func NewCompressUnitTestTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressUnitTestTransactor, error) {
	contract, err := bindCompressUnitTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressUnitTestTransactor{contract: contract}, nil
}

// NewCompressUnitTestFilterer creates a new log filterer instance of CompressUnitTest, bound to a specific deployed contract.
func NewCompressUnitTestFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressUnitTestFilterer, error) {
	contract, err := bindCompressUnitTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressUnitTestFilterer{contract: contract}, nil
}

// bindCompressUnitTest binds a generic wrapper to an already deployed contract.
func bindCompressUnitTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressUnitTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressUnitTest *CompressUnitTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressUnitTest.Contract.CompressUnitTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressUnitTest *CompressUnitTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressUnitTest.Contract.CompressUnitTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressUnitTest *CompressUnitTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressUnitTest.Contract.CompressUnitTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressUnitTest *CompressUnitTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressUnitTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressUnitTest *CompressUnitTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressUnitTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressUnitTest *CompressUnitTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressUnitTest.Contract.contract.Transact(opts, method, params...)
}

// Decompress is a free data retrieval call binding the contract method 0x98d61183.
//
// Solidity: function decompress(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof) pure returns(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])))
func (_CompressUnitTest *CompressUnitTestCaller) Decompress(opts *bind.CallOpts, proof CommitmentProofData) (CommitmentProofData, error) {
	var out []interface{}
	err := _CompressUnitTest.contract.Call(opts, &out, "decompress", proof)

	if err != nil {
		return *new(CommitmentProofData), err
	}

	out0 := *abi.ConvertType(out[0], new(CommitmentProofData)).(*CommitmentProofData)

	return out0, err

}

// Decompress is a free data retrieval call binding the contract method 0x98d61183.
//
// Solidity: function decompress(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof) pure returns(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])))
func (_CompressUnitTest *CompressUnitTestSession) Decompress(proof CommitmentProofData) (CommitmentProofData, error) {
	return _CompressUnitTest.Contract.Decompress(&_CompressUnitTest.CallOpts, proof)
}

// Decompress is a free data retrieval call binding the contract method 0x98d61183.
//
// Solidity: function decompress(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof) pure returns(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])))
func (_CompressUnitTest *CompressUnitTestCallerSession) Decompress(proof CommitmentProofData) (CommitmentProofData, error) {
	return _CompressUnitTest.Contract.Decompress(&_CompressUnitTest.CallOpts, proof)
}

// CompressedBatchEntryABI is the input ABI used to generate the binding from.
const CompressedBatchEntryABI = "[]"

// CompressedBatchEntryBin is the compiled bytecode used for deploying new contracts.
var CompressedBatchEntryBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220da223f7b3f52087910beef3c77d9ebc312d19ddab9235150872cc14deaa5b52a64736f6c63430008090033"

// DeployCompressedBatchEntry deploys a new Ethereum contract, binding an instance of CompressedBatchEntry to it.
func DeployCompressedBatchEntry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompressedBatchEntry, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressedBatchEntryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompressedBatchEntryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompressedBatchEntry{CompressedBatchEntryCaller: CompressedBatchEntryCaller{contract: contract}, CompressedBatchEntryTransactor: CompressedBatchEntryTransactor{contract: contract}, CompressedBatchEntryFilterer: CompressedBatchEntryFilterer{contract: contract}}, nil
}

// CompressedBatchEntry is an auto generated Go binding around an Ethereum contract.
type CompressedBatchEntry struct {
	CompressedBatchEntryCaller     // Read-only binding to the contract
	CompressedBatchEntryTransactor // Write-only binding to the contract
	CompressedBatchEntryFilterer   // Log filterer for contract events
}

// CompressedBatchEntryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressedBatchEntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchEntryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressedBatchEntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchEntryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressedBatchEntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchEntrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressedBatchEntrySession struct {
	Contract     *CompressedBatchEntry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CompressedBatchEntryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressedBatchEntryCallerSession struct {
	Contract *CompressedBatchEntryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CompressedBatchEntryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressedBatchEntryTransactorSession struct {
	Contract     *CompressedBatchEntryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CompressedBatchEntryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressedBatchEntryRaw struct {
	Contract *CompressedBatchEntry // Generic contract binding to access the raw methods on
}

// CompressedBatchEntryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressedBatchEntryCallerRaw struct {
	Contract *CompressedBatchEntryCaller // Generic read-only contract binding to access the raw methods on
}

// CompressedBatchEntryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressedBatchEntryTransactorRaw struct {
	Contract *CompressedBatchEntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompressedBatchEntry creates a new instance of CompressedBatchEntry, bound to a specific deployed contract.
func NewCompressedBatchEntry(address common.Address, backend bind.ContractBackend) (*CompressedBatchEntry, error) {
	contract, err := bindCompressedBatchEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchEntry{CompressedBatchEntryCaller: CompressedBatchEntryCaller{contract: contract}, CompressedBatchEntryTransactor: CompressedBatchEntryTransactor{contract: contract}, CompressedBatchEntryFilterer: CompressedBatchEntryFilterer{contract: contract}}, nil
}

// NewCompressedBatchEntryCaller creates a new read-only instance of CompressedBatchEntry, bound to a specific deployed contract.
func NewCompressedBatchEntryCaller(address common.Address, caller bind.ContractCaller) (*CompressedBatchEntryCaller, error) {
	contract, err := bindCompressedBatchEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchEntryCaller{contract: contract}, nil
}

// NewCompressedBatchEntryTransactor creates a new write-only instance of CompressedBatchEntry, bound to a specific deployed contract.
func NewCompressedBatchEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressedBatchEntryTransactor, error) {
	contract, err := bindCompressedBatchEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchEntryTransactor{contract: contract}, nil
}

// NewCompressedBatchEntryFilterer creates a new log filterer instance of CompressedBatchEntry, bound to a specific deployed contract.
func NewCompressedBatchEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressedBatchEntryFilterer, error) {
	contract, err := bindCompressedBatchEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchEntryFilterer{contract: contract}, nil
}

// bindCompressedBatchEntry binds a generic wrapper to an already deployed contract.
func bindCompressedBatchEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressedBatchEntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedBatchEntry *CompressedBatchEntryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedBatchEntry.Contract.CompressedBatchEntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedBatchEntry *CompressedBatchEntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedBatchEntry.Contract.CompressedBatchEntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedBatchEntry *CompressedBatchEntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedBatchEntry.Contract.CompressedBatchEntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedBatchEntry *CompressedBatchEntryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedBatchEntry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedBatchEntry *CompressedBatchEntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedBatchEntry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedBatchEntry *CompressedBatchEntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedBatchEntry.Contract.contract.Transact(opts, method, params...)
}

// CompressedBatchProofABI is the input ABI used to generate the binding from.
const CompressedBatchProofABI = "[]"

// CompressedBatchProofBin is the compiled bytecode used for deploying new contracts.
var CompressedBatchProofBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122077777c0034808ded8911864499213a73e39cd40a1f4d4e2e20dde25297359c0064736f6c63430008090033"

// DeployCompressedBatchProof deploys a new Ethereum contract, binding an instance of CompressedBatchProof to it.
func DeployCompressedBatchProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompressedBatchProof, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressedBatchProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompressedBatchProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompressedBatchProof{CompressedBatchProofCaller: CompressedBatchProofCaller{contract: contract}, CompressedBatchProofTransactor: CompressedBatchProofTransactor{contract: contract}, CompressedBatchProofFilterer: CompressedBatchProofFilterer{contract: contract}}, nil
}

// CompressedBatchProof is an auto generated Go binding around an Ethereum contract.
type CompressedBatchProof struct {
	CompressedBatchProofCaller     // Read-only binding to the contract
	CompressedBatchProofTransactor // Write-only binding to the contract
	CompressedBatchProofFilterer   // Log filterer for contract events
}

// CompressedBatchProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressedBatchProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressedBatchProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressedBatchProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedBatchProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressedBatchProofSession struct {
	Contract     *CompressedBatchProof // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CompressedBatchProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressedBatchProofCallerSession struct {
	Contract *CompressedBatchProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CompressedBatchProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressedBatchProofTransactorSession struct {
	Contract     *CompressedBatchProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CompressedBatchProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressedBatchProofRaw struct {
	Contract *CompressedBatchProof // Generic contract binding to access the raw methods on
}

// CompressedBatchProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressedBatchProofCallerRaw struct {
	Contract *CompressedBatchProofCaller // Generic read-only contract binding to access the raw methods on
}

// CompressedBatchProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressedBatchProofTransactorRaw struct {
	Contract *CompressedBatchProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompressedBatchProof creates a new instance of CompressedBatchProof, bound to a specific deployed contract.
func NewCompressedBatchProof(address common.Address, backend bind.ContractBackend) (*CompressedBatchProof, error) {
	contract, err := bindCompressedBatchProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchProof{CompressedBatchProofCaller: CompressedBatchProofCaller{contract: contract}, CompressedBatchProofTransactor: CompressedBatchProofTransactor{contract: contract}, CompressedBatchProofFilterer: CompressedBatchProofFilterer{contract: contract}}, nil
}

// NewCompressedBatchProofCaller creates a new read-only instance of CompressedBatchProof, bound to a specific deployed contract.
func NewCompressedBatchProofCaller(address common.Address, caller bind.ContractCaller) (*CompressedBatchProofCaller, error) {
	contract, err := bindCompressedBatchProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchProofCaller{contract: contract}, nil
}

// NewCompressedBatchProofTransactor creates a new write-only instance of CompressedBatchProof, bound to a specific deployed contract.
func NewCompressedBatchProofTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressedBatchProofTransactor, error) {
	contract, err := bindCompressedBatchProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchProofTransactor{contract: contract}, nil
}

// NewCompressedBatchProofFilterer creates a new log filterer instance of CompressedBatchProof, bound to a specific deployed contract.
func NewCompressedBatchProofFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressedBatchProofFilterer, error) {
	contract, err := bindCompressedBatchProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressedBatchProofFilterer{contract: contract}, nil
}

// bindCompressedBatchProof binds a generic wrapper to an already deployed contract.
func bindCompressedBatchProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressedBatchProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedBatchProof *CompressedBatchProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedBatchProof.Contract.CompressedBatchProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedBatchProof *CompressedBatchProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedBatchProof.Contract.CompressedBatchProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedBatchProof *CompressedBatchProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedBatchProof.Contract.CompressedBatchProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedBatchProof *CompressedBatchProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedBatchProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedBatchProof *CompressedBatchProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedBatchProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedBatchProof *CompressedBatchProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedBatchProof.Contract.contract.Transact(opts, method, params...)
}

// CompressedExistenceProofABI is the input ABI used to generate the binding from.
const CompressedExistenceProofABI = "[]"

// CompressedExistenceProofBin is the compiled bytecode used for deploying new contracts.
var CompressedExistenceProofBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201808a43d6df6525ff23f01481c4949b8e9e7387835da7cda5b28284adb687e7b64736f6c63430008090033"

// DeployCompressedExistenceProof deploys a new Ethereum contract, binding an instance of CompressedExistenceProof to it.
func DeployCompressedExistenceProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompressedExistenceProof, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressedExistenceProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompressedExistenceProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompressedExistenceProof{CompressedExistenceProofCaller: CompressedExistenceProofCaller{contract: contract}, CompressedExistenceProofTransactor: CompressedExistenceProofTransactor{contract: contract}, CompressedExistenceProofFilterer: CompressedExistenceProofFilterer{contract: contract}}, nil
}

// CompressedExistenceProof is an auto generated Go binding around an Ethereum contract.
type CompressedExistenceProof struct {
	CompressedExistenceProofCaller     // Read-only binding to the contract
	CompressedExistenceProofTransactor // Write-only binding to the contract
	CompressedExistenceProofFilterer   // Log filterer for contract events
}

// CompressedExistenceProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressedExistenceProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedExistenceProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressedExistenceProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedExistenceProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressedExistenceProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedExistenceProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressedExistenceProofSession struct {
	Contract     *CompressedExistenceProof // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CompressedExistenceProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressedExistenceProofCallerSession struct {
	Contract *CompressedExistenceProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// CompressedExistenceProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressedExistenceProofTransactorSession struct {
	Contract     *CompressedExistenceProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// CompressedExistenceProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressedExistenceProofRaw struct {
	Contract *CompressedExistenceProof // Generic contract binding to access the raw methods on
}

// CompressedExistenceProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressedExistenceProofCallerRaw struct {
	Contract *CompressedExistenceProofCaller // Generic read-only contract binding to access the raw methods on
}

// CompressedExistenceProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressedExistenceProofTransactorRaw struct {
	Contract *CompressedExistenceProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompressedExistenceProof creates a new instance of CompressedExistenceProof, bound to a specific deployed contract.
func NewCompressedExistenceProof(address common.Address, backend bind.ContractBackend) (*CompressedExistenceProof, error) {
	contract, err := bindCompressedExistenceProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompressedExistenceProof{CompressedExistenceProofCaller: CompressedExistenceProofCaller{contract: contract}, CompressedExistenceProofTransactor: CompressedExistenceProofTransactor{contract: contract}, CompressedExistenceProofFilterer: CompressedExistenceProofFilterer{contract: contract}}, nil
}

// NewCompressedExistenceProofCaller creates a new read-only instance of CompressedExistenceProof, bound to a specific deployed contract.
func NewCompressedExistenceProofCaller(address common.Address, caller bind.ContractCaller) (*CompressedExistenceProofCaller, error) {
	contract, err := bindCompressedExistenceProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedExistenceProofCaller{contract: contract}, nil
}

// NewCompressedExistenceProofTransactor creates a new write-only instance of CompressedExistenceProof, bound to a specific deployed contract.
func NewCompressedExistenceProofTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressedExistenceProofTransactor, error) {
	contract, err := bindCompressedExistenceProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedExistenceProofTransactor{contract: contract}, nil
}

// NewCompressedExistenceProofFilterer creates a new log filterer instance of CompressedExistenceProof, bound to a specific deployed contract.
func NewCompressedExistenceProofFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressedExistenceProofFilterer, error) {
	contract, err := bindCompressedExistenceProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressedExistenceProofFilterer{contract: contract}, nil
}

// bindCompressedExistenceProof binds a generic wrapper to an already deployed contract.
func bindCompressedExistenceProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressedExistenceProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedExistenceProof *CompressedExistenceProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedExistenceProof.Contract.CompressedExistenceProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedExistenceProof *CompressedExistenceProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedExistenceProof.Contract.CompressedExistenceProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedExistenceProof *CompressedExistenceProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedExistenceProof.Contract.CompressedExistenceProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedExistenceProof *CompressedExistenceProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedExistenceProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedExistenceProof *CompressedExistenceProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedExistenceProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedExistenceProof *CompressedExistenceProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedExistenceProof.Contract.contract.Transact(opts, method, params...)
}

// CompressedNonExistenceProofABI is the input ABI used to generate the binding from.
const CompressedNonExistenceProofABI = "[]"

// CompressedNonExistenceProofBin is the compiled bytecode used for deploying new contracts.
var CompressedNonExistenceProofBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e085c5abb686e514f849dca6af14c2e43e649a6f90204f50ea11269fcb70bb7864736f6c63430008090033"

// DeployCompressedNonExistenceProof deploys a new Ethereum contract, binding an instance of CompressedNonExistenceProof to it.
func DeployCompressedNonExistenceProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CompressedNonExistenceProof, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressedNonExistenceProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CompressedNonExistenceProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CompressedNonExistenceProof{CompressedNonExistenceProofCaller: CompressedNonExistenceProofCaller{contract: contract}, CompressedNonExistenceProofTransactor: CompressedNonExistenceProofTransactor{contract: contract}, CompressedNonExistenceProofFilterer: CompressedNonExistenceProofFilterer{contract: contract}}, nil
}

// CompressedNonExistenceProof is an auto generated Go binding around an Ethereum contract.
type CompressedNonExistenceProof struct {
	CompressedNonExistenceProofCaller     // Read-only binding to the contract
	CompressedNonExistenceProofTransactor // Write-only binding to the contract
	CompressedNonExistenceProofFilterer   // Log filterer for contract events
}

// CompressedNonExistenceProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressedNonExistenceProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedNonExistenceProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressedNonExistenceProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedNonExistenceProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressedNonExistenceProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressedNonExistenceProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressedNonExistenceProofSession struct {
	Contract     *CompressedNonExistenceProof // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CompressedNonExistenceProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressedNonExistenceProofCallerSession struct {
	Contract *CompressedNonExistenceProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// CompressedNonExistenceProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressedNonExistenceProofTransactorSession struct {
	Contract     *CompressedNonExistenceProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// CompressedNonExistenceProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressedNonExistenceProofRaw struct {
	Contract *CompressedNonExistenceProof // Generic contract binding to access the raw methods on
}

// CompressedNonExistenceProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressedNonExistenceProofCallerRaw struct {
	Contract *CompressedNonExistenceProofCaller // Generic read-only contract binding to access the raw methods on
}

// CompressedNonExistenceProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressedNonExistenceProofTransactorRaw struct {
	Contract *CompressedNonExistenceProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompressedNonExistenceProof creates a new instance of CompressedNonExistenceProof, bound to a specific deployed contract.
func NewCompressedNonExistenceProof(address common.Address, backend bind.ContractBackend) (*CompressedNonExistenceProof, error) {
	contract, err := bindCompressedNonExistenceProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CompressedNonExistenceProof{CompressedNonExistenceProofCaller: CompressedNonExistenceProofCaller{contract: contract}, CompressedNonExistenceProofTransactor: CompressedNonExistenceProofTransactor{contract: contract}, CompressedNonExistenceProofFilterer: CompressedNonExistenceProofFilterer{contract: contract}}, nil
}

// NewCompressedNonExistenceProofCaller creates a new read-only instance of CompressedNonExistenceProof, bound to a specific deployed contract.
func NewCompressedNonExistenceProofCaller(address common.Address, caller bind.ContractCaller) (*CompressedNonExistenceProofCaller, error) {
	contract, err := bindCompressedNonExistenceProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedNonExistenceProofCaller{contract: contract}, nil
}

// NewCompressedNonExistenceProofTransactor creates a new write-only instance of CompressedNonExistenceProof, bound to a specific deployed contract.
func NewCompressedNonExistenceProofTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressedNonExistenceProofTransactor, error) {
	contract, err := bindCompressedNonExistenceProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressedNonExistenceProofTransactor{contract: contract}, nil
}

// NewCompressedNonExistenceProofFilterer creates a new log filterer instance of CompressedNonExistenceProof, bound to a specific deployed contract.
func NewCompressedNonExistenceProofFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressedNonExistenceProofFilterer, error) {
	contract, err := bindCompressedNonExistenceProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressedNonExistenceProofFilterer{contract: contract}, nil
}

// bindCompressedNonExistenceProof binds a generic wrapper to an already deployed contract.
func bindCompressedNonExistenceProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CompressedNonExistenceProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedNonExistenceProof.Contract.CompressedNonExistenceProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedNonExistenceProof.Contract.CompressedNonExistenceProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedNonExistenceProof.Contract.CompressedNonExistenceProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CompressedNonExistenceProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CompressedNonExistenceProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CompressedNonExistenceProof *CompressedNonExistenceProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CompressedNonExistenceProof.Contract.contract.Transact(opts, method, params...)
}

// ExistenceProofABI is the input ABI used to generate the binding from.
const ExistenceProofABI = "[]"

// ExistenceProofBin is the compiled bytecode used for deploying new contracts.
var ExistenceProofBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122087bb076404d6ee26b14ad133799989bc45999c6d225b837e6b57d2395a6642fd64736f6c63430008090033"

// DeployExistenceProof deploys a new Ethereum contract, binding an instance of ExistenceProof to it.
func DeployExistenceProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExistenceProof, error) {
	parsed, err := abi.JSON(strings.NewReader(ExistenceProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExistenceProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExistenceProof{ExistenceProofCaller: ExistenceProofCaller{contract: contract}, ExistenceProofTransactor: ExistenceProofTransactor{contract: contract}, ExistenceProofFilterer: ExistenceProofFilterer{contract: contract}}, nil
}

// ExistenceProof is an auto generated Go binding around an Ethereum contract.
type ExistenceProof struct {
	ExistenceProofCaller     // Read-only binding to the contract
	ExistenceProofTransactor // Write-only binding to the contract
	ExistenceProofFilterer   // Log filterer for contract events
}

// ExistenceProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExistenceProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistenceProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExistenceProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistenceProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExistenceProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistenceProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExistenceProofSession struct {
	Contract     *ExistenceProof   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExistenceProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExistenceProofCallerSession struct {
	Contract *ExistenceProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ExistenceProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExistenceProofTransactorSession struct {
	Contract     *ExistenceProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ExistenceProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExistenceProofRaw struct {
	Contract *ExistenceProof // Generic contract binding to access the raw methods on
}

// ExistenceProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExistenceProofCallerRaw struct {
	Contract *ExistenceProofCaller // Generic read-only contract binding to access the raw methods on
}

// ExistenceProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExistenceProofTransactorRaw struct {
	Contract *ExistenceProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExistenceProof creates a new instance of ExistenceProof, bound to a specific deployed contract.
func NewExistenceProof(address common.Address, backend bind.ContractBackend) (*ExistenceProof, error) {
	contract, err := bindExistenceProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExistenceProof{ExistenceProofCaller: ExistenceProofCaller{contract: contract}, ExistenceProofTransactor: ExistenceProofTransactor{contract: contract}, ExistenceProofFilterer: ExistenceProofFilterer{contract: contract}}, nil
}

// NewExistenceProofCaller creates a new read-only instance of ExistenceProof, bound to a specific deployed contract.
func NewExistenceProofCaller(address common.Address, caller bind.ContractCaller) (*ExistenceProofCaller, error) {
	contract, err := bindExistenceProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExistenceProofCaller{contract: contract}, nil
}

// NewExistenceProofTransactor creates a new write-only instance of ExistenceProof, bound to a specific deployed contract.
func NewExistenceProofTransactor(address common.Address, transactor bind.ContractTransactor) (*ExistenceProofTransactor, error) {
	contract, err := bindExistenceProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExistenceProofTransactor{contract: contract}, nil
}

// NewExistenceProofFilterer creates a new log filterer instance of ExistenceProof, bound to a specific deployed contract.
func NewExistenceProofFilterer(address common.Address, filterer bind.ContractFilterer) (*ExistenceProofFilterer, error) {
	contract, err := bindExistenceProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExistenceProofFilterer{contract: contract}, nil
}

// bindExistenceProof binds a generic wrapper to an already deployed contract.
func bindExistenceProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExistenceProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExistenceProof *ExistenceProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExistenceProof.Contract.ExistenceProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExistenceProof *ExistenceProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExistenceProof.Contract.ExistenceProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExistenceProof *ExistenceProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExistenceProof.Contract.ExistenceProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExistenceProof *ExistenceProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExistenceProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExistenceProof *ExistenceProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExistenceProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExistenceProof *ExistenceProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExistenceProof.Contract.contract.Transact(opts, method, params...)
}

// GoogleProtobufAnyABI is the input ABI used to generate the binding from.
const GoogleProtobufAnyABI = "[]"

// GoogleProtobufAnyBin is the compiled bytecode used for deploying new contracts.
var GoogleProtobufAnyBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220836af561efc39dd1d35461705e188b8fdef523fbf782a7ec550da02312f3136a64736f6c63430008090033"

// DeployGoogleProtobufAny deploys a new Ethereum contract, binding an instance of GoogleProtobufAny to it.
func DeployGoogleProtobufAny(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GoogleProtobufAny, error) {
	parsed, err := abi.JSON(strings.NewReader(GoogleProtobufAnyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GoogleProtobufAnyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GoogleProtobufAny{GoogleProtobufAnyCaller: GoogleProtobufAnyCaller{contract: contract}, GoogleProtobufAnyTransactor: GoogleProtobufAnyTransactor{contract: contract}, GoogleProtobufAnyFilterer: GoogleProtobufAnyFilterer{contract: contract}}, nil
}

// GoogleProtobufAny is an auto generated Go binding around an Ethereum contract.
type GoogleProtobufAny struct {
	GoogleProtobufAnyCaller     // Read-only binding to the contract
	GoogleProtobufAnyTransactor // Write-only binding to the contract
	GoogleProtobufAnyFilterer   // Log filterer for contract events
}

// GoogleProtobufAnyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoogleProtobufAnyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoogleProtobufAnyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoogleProtobufAnyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoogleProtobufAnyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoogleProtobufAnyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoogleProtobufAnySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoogleProtobufAnySession struct {
	Contract     *GoogleProtobufAny // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GoogleProtobufAnyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoogleProtobufAnyCallerSession struct {
	Contract *GoogleProtobufAnyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GoogleProtobufAnyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoogleProtobufAnyTransactorSession struct {
	Contract     *GoogleProtobufAnyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GoogleProtobufAnyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoogleProtobufAnyRaw struct {
	Contract *GoogleProtobufAny // Generic contract binding to access the raw methods on
}

// GoogleProtobufAnyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoogleProtobufAnyCallerRaw struct {
	Contract *GoogleProtobufAnyCaller // Generic read-only contract binding to access the raw methods on
}

// GoogleProtobufAnyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoogleProtobufAnyTransactorRaw struct {
	Contract *GoogleProtobufAnyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoogleProtobufAny creates a new instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAny(address common.Address, backend bind.ContractBackend) (*GoogleProtobufAny, error) {
	contract, err := bindGoogleProtobufAny(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAny{GoogleProtobufAnyCaller: GoogleProtobufAnyCaller{contract: contract}, GoogleProtobufAnyTransactor: GoogleProtobufAnyTransactor{contract: contract}, GoogleProtobufAnyFilterer: GoogleProtobufAnyFilterer{contract: contract}}, nil
}

// NewGoogleProtobufAnyCaller creates a new read-only instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAnyCaller(address common.Address, caller bind.ContractCaller) (*GoogleProtobufAnyCaller, error) {
	contract, err := bindGoogleProtobufAny(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAnyCaller{contract: contract}, nil
}

// NewGoogleProtobufAnyTransactor creates a new write-only instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAnyTransactor(address common.Address, transactor bind.ContractTransactor) (*GoogleProtobufAnyTransactor, error) {
	contract, err := bindGoogleProtobufAny(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAnyTransactor{contract: contract}, nil
}

// NewGoogleProtobufAnyFilterer creates a new log filterer instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAnyFilterer(address common.Address, filterer bind.ContractFilterer) (*GoogleProtobufAnyFilterer, error) {
	contract, err := bindGoogleProtobufAny(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAnyFilterer{contract: contract}, nil
}

// bindGoogleProtobufAny binds a generic wrapper to an already deployed contract.
func bindGoogleProtobufAny(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GoogleProtobufAnyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoogleProtobufAny *GoogleProtobufAnyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoogleProtobufAny.Contract.GoogleProtobufAnyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoogleProtobufAny *GoogleProtobufAnyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.GoogleProtobufAnyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoogleProtobufAny *GoogleProtobufAnyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.GoogleProtobufAnyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoogleProtobufAny *GoogleProtobufAnyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoogleProtobufAny.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoogleProtobufAny *GoogleProtobufAnyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoogleProtobufAny *GoogleProtobufAnyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.contract.Transact(opts, method, params...)
}

// ICS23UnitTestABI is the input ABI used to generate the binding from.
const ICS23UnitTestABI = "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf_spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int32[]\",\"name\":\"child_order\",\"type\":\"int32[]\"},{\"internalType\":\"int32\",\"name\":\"child_size\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"max_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"empty_child\",\"type\":\"bytes\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structInnerSpec.Data\",\"name\":\"inner_spec\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"max_depth\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_depth\",\"type\":\"int32\"}],\"internalType\":\"structProofSpec.Data\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"commitmentRoot\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structBatchProof.Data\",\"name\":\"batch\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"lookup_inners\",\"type\":\"tuple[]\"}],\"internalType\":\"structCompressedBatchProof.Data\",\"name\":\"compressed\",\"type\":\"tuple\"}],\"internalType\":\"structCommitmentProof.Data\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structIcs23.BatchItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"name\":\"batchVerifyMembership\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf_spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int32[]\",\"name\":\"child_order\",\"type\":\"int32[]\"},{\"internalType\":\"int32\",\"name\":\"child_size\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"max_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"empty_child\",\"type\":\"bytes\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structInnerSpec.Data\",\"name\":\"inner_spec\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"max_depth\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_depth\",\"type\":\"int32\"}],\"internalType\":\"structProofSpec.Data\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"commitmentRoot\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structBatchProof.Data\",\"name\":\"batch\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"lookup_inners\",\"type\":\"tuple[]\"}],\"internalType\":\"structCompressedBatchProof.Data\",\"name\":\"compressed\",\"type\":\"tuple\"}],\"internalType\":\"structCommitmentProof.Data\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"keys\",\"type\":\"bytes[]\"}],\"name\":\"batchVerifyNonMembership\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf_spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int32[]\",\"name\":\"child_order\",\"type\":\"int32[]\"},{\"internalType\":\"int32\",\"name\":\"child_size\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"max_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"empty_child\",\"type\":\"bytes\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structInnerSpec.Data\",\"name\":\"inner_spec\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"max_depth\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_depth\",\"type\":\"int32\"}],\"internalType\":\"structProofSpec.Data\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"commitmentRoot\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structBatchProof.Data\",\"name\":\"batch\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"lookup_inners\",\"type\":\"tuple[]\"}],\"internalType\":\"structCompressedBatchProof.Data\",\"name\":\"compressed\",\"type\":\"tuple\"}],\"internalType\":\"structCommitmentProof.Data\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyMembership\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf_spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int32[]\",\"name\":\"child_order\",\"type\":\"int32[]\"},{\"internalType\":\"int32\",\"name\":\"child_size\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"max_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"empty_child\",\"type\":\"bytes\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structInnerSpec.Data\",\"name\":\"inner_spec\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"max_depth\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_depth\",\"type\":\"int32\"}],\"internalType\":\"structProofSpec.Data\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"commitmentRoot\",\"type\":\"bytes\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structBatchProof.Data\",\"name\":\"batch\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"lookup_inners\",\"type\":\"tuple[]\"}],\"internalType\":\"structCompressedBatchProof.Data\",\"name\":\"compressed\",\"type\":\"tuple\"}],\"internalType\":\"structCommitmentProof.Data\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"verifyNonMembership\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ICS23UnitTestBin is the compiled bytecode used for deploying new contracts.
var ICS23UnitTestBin = "0x608060405234801561001057600080fd5b50613908806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633aaa08eb146100515780636b6c681814610066578063b24b164b14610079578063c4cd136f1461008c575b600080fd5b61006461005f366004613326565b61009f565b005b6100646100743660046134a0565b6100b1565b610064610087366004613571565b6100c5565b61006461009a36600461368a565b6100d1565b6100ab848484846100dd565b50505050565b6100be8585858585610157565b5050505050565b6100ab848484846101e4565b6100ab84848484610234565b60006100e8836102b4565b905060005b825181101561014f5761013d86868486858151811061010e5761010e613736565b60200260200101516000015187868151811061012c5761012c613736565b602002602001015160200151610157565b8061014781613762565b9150506100ed565b505050505050565b6000610162846102b4565b90506000610170828561032b565b9050806101ce5760405162461bcd60e51b815260206004820152602160248201527f676574457869737450726f6f66466f724b6579206e6f7420617661696c61626c6044820152606560f81b60648201526084015b60405180910390fd5b6101db8188888787610437565b50505050505050565b60006101ef836102b4565b905060005b825181101561014f5761022286868486858151811061021557610215613736565b6020026020010151610234565b8061022c81613762565b9150506101f4565b600061023f836102b4565b9050600061024d8284610576565b9050806102a85760405162461bcd60e51b8152602060048201526024808201527f6765744e6f6e457869737450726f6f66466f724b6579206e6f7420617661696c60448201526361626c6560e01b60648201526084016101c5565b61014f818787866106d2565b6102bc6127ce565b6102c98260600151610a16565b1515600114156102d7575090565b60405180608001604052806102ea610a44565b81526020016102f7610a52565b815260200160405180602001604052806103148660600151610a5a565b90528152602001610323610b26565b905292915050565b610333612830565b82511561035e578251516103479083610b3f565b15156001141561035957508151610431565b610426565b6040830151156104265760005b60408401515151811015610424576103a6846040015160000151828151811061039657610396613736565b6020026020010151600001511590565b1580156103e157506103e184604001516000015182815181106103cb576103cb613736565b6020026020010151600001516000015184610b3f565b15610412576040840151518051829081106103fe576103fe613736565b602002602001015160000151915050610431565b8061041c81613762565b91505061036b565b505b61042e610a44565b90505b92915050565b84516104439083610b3f565b61048f5760405162461bcd60e51b815260206004820181905260248201527f50726f7669646564206b657920646f65736e2774206d617463682070726f6f6660448201526064016101c5565b61049d856020015182610b3f565b6104f45760405162461bcd60e51b815260206004820152602260248201527f50726f76696465642076616c756520646f65736e2774206d617463682070726f60448201526137b360f11b60648201526084016101c5565b6104fe8585610ba3565b600061050986610d1f565b90506105158185610b3f565b61014f5760405162461bcd60e51b815260206004820152602c60248201527f43616c63756c636174656420726f6f7420646f65736e2774206d61746368207060448201526b1c9bdd9a591959081c9bdbdd60a21b60648201526084016101c5565b61057e61285e565b6020830151156105c85761059a83602001516020015183610db0565b80156105b357506105b383602001516040015183610dd9565b156105c357506020820151610431565b6106ca565b6040830151156106ca5760005b604084015151518110156106c857610610846040015160000151828151811061060057610600613736565b6020026020010151602001511590565b15801561064b575061064b846040015160000151828151811061063557610635613736565b6020026020010151602001516020015184610db0565b80156106855750610685846040015160000151828151811061066f5761066f613736565b6020026020010151602001516040015184610dd9565b156106b6576040840151518051829081106106a2576106a2613736565b602002602001015160200151915050610431565b806106c081613762565b9150506105d5565b505b61042e610a52565b6060806106e28660200151610e02565b61070957602080870151805191810151610700928891889190610437565b60208601515191505b6107168660400151610e02565b61073d5760408601518051602082015161073592918891889190610437565b506040850151515b60008251118061074e575060008151115b6107a55760405162461bcd60e51b815260206004820152602260248201527f626f7468206c65667420616e642072696768742070726f6f6673206d697373696044820152616e6760f01b60648201526084016101c5565b8051156108055760006107b88483610e3b565b126108055760405162461bcd60e51b815260206004820152601e60248201527f6b6579206973206e6f74206c656674206f662072696768742070726f6f66000060448201526064016101c5565b8151156108655760006108188484610e3b565b136108655760405162461bcd60e51b815260206004820152601e60248201527f6b6579206973206e6f74207269676874206f66206c6566742070726f6f66000060448201526064016101c5565b81516108ec576108818560200151876040015160600151610f38565b6108e75760405162461bcd60e51b815260206004820152603160248201527f6c6566742070726f6f66206d697373696e672c2072696768742070726f6f66206044820152701b5d5cdd081899481b19599d0b5b5bdcdd607a1b60648201526084016101c5565b61014f565b805161097a576109088560200151876020015160600151610fb0565b6108e75760405162461bcd60e51b815260206004820152603f60248201527f697352696768744d6f73743a2072696768742070726f6f66206d697373696e6760448201527f2c206c6566742070726f6f66206d7573742062652072696768742d6d6f73740060648201526084016101c5565b610999856020015187602001516060015188604001516060015161102f565b61014f5760405162461bcd60e51b815260206004820152604260248201527f69734c6566744e65696768626f723a2072696768742070726f6f66206d69737360448201527f696e672c206c6566742070726f6f66206d7573742062652072696768742d6d6f6064820152611cdd60f21b608482015260a4016101c5565b80515160009015610a2957506000919050565b60208201515115610a3c57506000919050565b506001919050565b610a4c612830565b50600090565b610a4c61285e565b606060008260000151516001600160401b03811115610a7b57610a7b6128cd565b604051908082528060200260200182016040528015610ab457816020015b610aa1612885565b815260200190600190039081610a995790505b50905060005b835151811015610b1f57610aef84600001518281518110610add57610add613736565b602002602001015185602001516111aa565b828281518110610b0157610b01613736565b60200260200101819052508080610b1790613762565b915050610aba565b5092915050565b6040805180820190915260608082526020820152610a4c565b815181516000916001918114808314610b5b5760009250610b99565b600160208701838101602088015b600284838510011415610b94578051835114610b885760009650600093505b60209283019201610b69565b505050505b5090949350505050565b610bb08260400151611255565b15610bcd5760405162461bcd60e51b81526004016101c59061377d565b610bdb8260400151826112f4565b6000816060015160030b1315610c56576000610bfd826060015160030b611587565b8360600151511015905080610c545760405162461bcd60e51b815260206004820152601860248201527f496e6e65724f707320646570746820746f6f2073686f7274000000000000000060448201526064016101c5565b505b6000816040015160030b1315610cd1576000610c78826040015160030b611587565b8360600151511115905080610ccf5760405162461bcd60e51b815260206004820152601760248201527f496e6e65724f707320646570746820746f6f206c6f6e6700000000000000000060448201526064016101c5565b505b60005b826060015151811015610d1a57610d0883606001518281518110610cfa57610cfa613736565b6020026020010151836115dd565b80610d1281613762565b915050610cd4565b505050565b6060610d2d82604001511590565b15610d4a5760405162461bcd60e51b81526004016101c59061377d565b6000610d6383604001518460000151856020015161180e565b905060005b836060015151811015610b1f57610d9c84606001518281518110610d8e57610d8e613736565b60200260200101518361190f565b915080610da881613762565b915050610d68565b6000610dbb83610e02565b8061042e57506000610dd1846000015184610e3b565b129392505050565b6000610de483610e02565b8061042e57506000610dfa846000015184610e3b565b139392505050565b80515160009015610e1557506000919050565b60208201515115610e2857506000919050565b60608201515115610a3c57506000919050565b600080610e4a845184516119a8565b905060005b81811015610f0657838181518110610e6957610e69613736565b602001015160f81c60f81b60f81c60ff16858281518110610e8c57610e8c613736565b016020015160f81c1015610ea65760001992505050610431565b838181518110610eb857610eb8613736565b602001015160f81c60f81b60f81c60ff16858281518110610edb57610edb613736565b016020015160f81c1115610ef457600192505050610431565b80610efe81613762565b915050610e4f565b508084511115610f1a576001915050610431565b8083511115610f2e57600019915050610431565b5060009392505050565b600080600080610f498660006119be565b92509250925060005b8551811015610fa357610f80868281518110610f7057610f70613736565b6020026020010151858585611a62565b610f91576000945050505050610431565b80610f9b81613762565b915050610f52565b5060019695505050505050565b6000806001846000015151610fc591906137c1565b90506000806000610fd687856119be565b92509250925060005b865181101561102157610ffd878281518110610f7057610f70613736565b61100f57600095505050505050610431565b8061101981613762565b915050610fdf565b506001979650505050505050565b6000806001845161104091906137c1565b905060006001845161105291906137c1565b90505b61109985838151811061106a5761106a613736565b60200260200101516020015185838151811061108857611088613736565b602002602001015160200151610b3f565b80156110e457506110e48583815181106110b5576110b5613736565b6020026020010151604001518583815181106110d3576110d3613736565b602002602001015160400151610b3f565b15611108576110f46001836137c1565b91506111016001826137c1565b9050611055565b6111458686848151811061111e5761111e613736565b602002602001015186848151811061113857611138613736565b6020026020010151611a9f565b611154576000925050506111a3565b6111698661116487600086611ad1565b610fb0565b611178576000925050506111a3565b61118d8661118886600085611ad1565b610f38565b61119c576000925050506111a3565b6001925050505b9392505050565b6111b2612885565b82516111bd90610e02565b6111f05760405180604001604052806111da856000015185611bac565b81526020016111e7610a52565b90529050610431565b6040518060400160405280611203610a44565b81526020016040518060600160405280866020015160000151815260200161123387602001516020015187611bac565b815260200161124a87602001516040015187611bac565b905290529392505050565b8051600090600681111561126b5761126b6137d8565b1561127857506000919050565b8160200151600681111561128e5761128e6137d8565b1561129b57506000919050565b816040015160068111156112b1576112b16137d8565b156112be57506000919050565b816060015160088111156112d4576112d46137d8565b156112e157506000919050565b60808201515115610a3c57506000919050565b8051516006811115611308576113086137d8565b8251600681111561131b5761131b6137d8565b1461136e5760405162461bcd60e51b815260206004820152602f60248201526000805160206138b383398151915260448201526e0657870656374656420486173684f7608c1b60648201526084016101c5565b8051602001516006811115611385576113856137d8565b8260200151600681111561139b5761139b6137d8565b146113f25760405162461bcd60e51b815260206004820152603360248201526000805160206138b3833981519152604482015272657870656374656420507265686173684b657960681b60648201526084016101c5565b8051604001516006811115611409576114096137d8565b8260400151600681111561141f5761141f6137d8565b146114785760405162461bcd60e51b815260206004820152603560248201526000805160206138b38339815191526044820152746578706563746564205072656861736856616c756560581b60648201526084016101c5565b805160600151600881111561148f5761148f6137d8565b826060015160088111156114a5576114a56137d8565b146114fa5760405162461bcd60e51b815260206004820152603160248201526000805160206138b383398151915260448201527006578706563746564206c656e6774684f7607c1b60648201526084016101c5565b60006115128360800151836000015160800151611db2565b905080610d1a5760405162461bcd60e51b815260206004820152603c60248201527f636865636b416761696e73745370656320666f72204c6561664f70202d204c6560448201527f61662050726566697820646f65736e277420737461727420776974680000000060648201526084016101c5565b6000808212156115d95760405162461bcd60e51b815260206004820181905260248201527f53616665436173743a2076616c7565206d75737420626520706f73697469766560448201526064016101c5565b5090565b806020015160a0015160068111156115f7576115f76137d8565b8251600681111561160a5761160a6137d8565b146116705760405162461bcd60e51b815260206004820152603060248201527f636865636b416761696e73745370656320666f7220496e6e65724f70202d205560448201526f06e657870656374656420486173684f760841b60648201526084016101c5565b600061168682602001516040015160030b611587565b90508083602001515110156116dd5760405162461bcd60e51b815260206004820152601860248201527f496e6e65724f702070726566697820746f6f2073686f7274000000000000000060448201526064016101c5565b81516080015160208401516000906116f59083611db2565b905080156117515760405162461bcd60e51b8152602060048201526024808201527f496e6e6572205072656669782073746172747320776974682077726f6e672076604482015263616c756560e01b60648201526084016101c5565b600061176785602001516020015160030b611587565b905060008160018760200151600001515161178291906137c1565b61178c91906137ee565b905060006117a487602001516060015160030b611587565b90506117b0828261380d565b88602001515111156118045760405162461bcd60e51b815260206004820152601760248201527f496e6e65724f702070726566697820746f6f206c6f6e6700000000000000000060448201526064016101c5565b5050505050505050565b606060008351116118555760405162461bcd60e51b81526020600482015260116024820152704c656166206f70206e65656473206b657960781b60448201526064016101c5565b600082511161189c5760405162461bcd60e51b81526020600482015260136024820152724c656166206f70206e656564732076616c756560681b60448201526064016101c5565b60006118b18560200151866060015186611df3565b905060006118c88660400151876060015186611df3565b90506000866080015183836040516020016118e593929190613860565b6040516020818303038152906040529050611904876000015182611e16565b979650505050505050565b606060008251116119625760405162461bcd60e51b815260206004820152601a60248201527f496e6e6572206f70206e65656473206368696c642076616c756500000000000060448201526064016101c5565b6000836020015183856040015160405160200161198193929190613860565b60405160208183030381529060405290506119a0846000015182611e16565b949350505050565b60008183106119b7578161042e565b5090919050565b6000806000806119d4866020015160030b611587565b905060006119e687600001518761213b565b905060006119f483836137ee565b9050611a06886040015160030b611587565b611a10908261380d565b9550611a22886060015160030b611587565b611a2c908261380d565b9450828260018a6000015151611a4291906137c1565b611a4c91906137c1565b611a5691906137ee565b93505050509250925092565b6000838560200151511015611a79575060006119a0565b828560200151511115611a8e575060006119a0565b506040840151518114949350505050565b600080611aac8585612208565b90506000611aba8685612208565b9050611ac782600161380d565b1495945050505050565b60606000611adf84846137c1565b6001600160401b03811115611af657611af66128cd565b604051908082528060200260200182016040528015611b4c57816020015b611b396040805160608101909152806000815260200160608152602001606081525090565b815260200190600190039081611b145790505b509050835b83811015611ba357858181518110611b6b57611b6b613736565b6020026020010151828281518110611b8557611b85613736565b60200260200101819052508080611b9b90613762565b915050611b51565b50949350505050565b611bb4612830565b611bbd83610e02565b15611bd157611bca610a44565b9050610431565b600060405180608001604052808560000151815260200185602001518152602001856040015181526020018560600151516001600160401b03811115611c1957611c196128cd565b604051908082528060200260200182016040528015611c6f57816020015b611c5c6040805160608101909152806000815260200160608152602001606081525090565b815260200190600190039081611c375790505b509052905060005b846060015151811015611daa57600085606001518281518110611c9c57611c9c613736565b602002602001015160030b1215611ce65760405162461bcd60e51b815260206004820152600e60248201526d070726f6f662e70617468203c20360941b60448201526064016101c5565b6000611d1186606001518381518110611d0157611d01613736565b602002602001015160030b611587565b905084518110611d5b5760405162461bcd60e51b81526020600482015260156024820152740e6e8cae0407c7a40d8deded6eae05cd8cadccee8d605b1b60448201526064016101c5565b848181518110611d6d57611d6d613736565b602002602001015183606001518381518110611d8b57611d8b613736565b6020026020010181905250508080611da290613762565b915050611c77565b509392505050565b6000815160001415611dc657506001610431565b825182511115611dd857506000610431565b6000611de784600085516122c3565b90506119a08382610b3f565b60606000611e0185846123d0565b9050611e0d84826123fd565b95945050505050565b60606001836006811115611e2c57611e2c6137d8565b1415611eaa57600282604051611e42919061387e565b602060405180830381855afa158015611e5f573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190611e82919061388a565b604051602001611e9491815260200190565b6040516020818303038152906040529050610431565b6003836006811115611ebe57611ebe6137d8565b1415611ede578180519060200120604051602001611e9491815260200190565b6004836006811115611ef257611ef26137d8565b1415611f4d57600382604051611f08919061387e565b602060405180830381855afa158015611f25573d6000803e3d6000fd5b5050604051805160601b6bffffffffffffffffffffffff191660208201526034019050611e94565b6005836006811115611f6157611f616137d8565b1415612043576000600283604051611f79919061387e565b602060405180830381855afa158015611f96573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190611fb9919061388a565b604051602001611fcb91815260200190565b6040516020818303038152906040529050600381604051611fec919061387e565b602060405180830381855afa158015612009573d6000803e3d6000fd5b5050604051805160601b6bffffffffffffffffffffffff191660208201526034019050604051602081830303815290604052915050610431565b6002836006811115612057576120576137d8565b141561209c5760405162461bcd60e51b815260206004820152601460248201527314d2104d4c4c881b9bdd081cdd5c1c1bdc9d195960621b60448201526064016101c5565b60068360068111156120b0576120b06137d8565b14156120fe5760405162461bcd60e51b815260206004820152601860248201527f5348413531325f323536206e6f7420737570706f72746564000000000000000060448201526064016101c5565b60405162461bcd60e51b81526020600482015260126024820152710556e737570706f7274656420686173684f760741b60448201526064016101c5565b60008251821061217e5760405162461bcd60e51b815260206004820152600e60248201526d0d2dcecc2d8d2c840c4e4c2dcc6d60931b60448201526064016101c5565b60005b83518110156121bf57826121a0858381518110611d0157611d01613736565b14156121ad579050610431565b806121b781613762565b915050612181565b5060405162461bcd60e51b815260206004820152601960248201527f6272616e6368206e6f7420666f756e6420696e206f726465720000000000000060448201526064016101c5565b815151600090815b8181101561226657600080600061222788856119be565b92509250925061223987848484611a62565b151560011415612250578395505050505050610431565b505050808061225e90613762565b915050612210565b5060405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f742066696e6420616e792076616c69642073706163696e6720666f60448201526a722074686973206e6f646560a81b60648201526084016101c5565b6060816122d181601f61380d565b10156123105760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b60448201526064016101c5565b61231a828461380d565b8451101561235e5760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b60448201526064016101c5565b60608215801561237d5760405191506000825260208201604052611ba3565b6040519150601f8416801560200281840101858101878315602002848b0101015b818310156123b657805183526020928301920161239e565b5050858452601f01601f1916604052505090509392505050565b606060008360068111156123e6576123e66137d8565b14156123f3575080610431565b61042e8383611e16565b60606000836008811115612413576124136137d8565b1415612420575080610431565b6001836008811115612434576124346137d8565b14156124c85760006124468351612709565b90506000816001600160401b03811115612462576124626128cd565b6040519080825280601f01601f19166020018201604052801561248c576020820181803683370190505b50905061249c8451602083612726565b5080846040516020016124b09291906138a3565b60405160208183030381529060405292505050610431565b60078360088111156124dc576124dc6137d8565b141561252e5781516020146125275760405162461bcd60e51b81526020600482015260116024820152703230ba30973632b733ba3410109e90199960791b60448201526064016101c5565b5080610431565b6008836008811115612542576125426137d8565b141561258d5781516040146125275760405162461bcd60e51b815260206004820152601160248201527019185d184b9b195b99dd1a08084f480d8d607a1b60448201526064016101c5565b60048360088111156125a1576125a16137d8565b14156126cd5760006125b38351612769565b60408051600480825281830190925291925060e083901b916000916020820181803683370190505090508160031a60f81b816000815181106125f7576125f7613736565b60200101906001600160f81b031916908160001a9053508160021a60f81b8160018151811061262857612628613736565b60200101906001600160f81b031916908160001a9053508160011a60f81b8160028151811061265957612659613736565b60200101906001600160f81b031916908160001a9053508160001a60f81b8160038151811061268a5761268a613736565b60200101906001600160f81b031916908160001a90535080856040516020016126b49291906138a3565b6040516020818303038152906040529350505050610431565b60405162461bcd60e51b81526020600482015260116024820152700556e737570706f72746564206c656e4f7607c1b60448201526064016101c5565b60071c600060015b82156104315760079290921c91600101612711565b600080828401607f86165b600787901c15612759578060801782535060079590951c9460019182019101607f8616612731565b8082535050600101949350505050565b600063ffffffff8211156115d95760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201526532206269747360d01b60648201526084016101c5565b60405180608001604052806127e1612830565b81526020016127ee61285e565b81526020016128096040518060200160405280606081525090565b815260200161282b604051806040016040528060608152602001606081525090565b905290565b604051806080016040528060608152602001606081526020016128516128a5565b8152602001606081525090565b604051806060016040528060608152602001612878612830565b815260200161282b612830565b6040518060400160405280612898612830565b815260200161282b61285e565b6040805160a08101909152806000815260200160008152602001600081526020016000612851565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b0381118282101715612905576129056128cd565b60405290565b60405160c081016001600160401b0381118282101715612905576129056128cd565b604051606081016001600160401b0381118282101715612905576129056128cd565b604051602081016001600160401b0381118282101715612905576129056128cd565b604080519081016001600160401b0381118282101715612905576129056128cd565b604051601f8201601f191681016001600160401b03811182821017156129bb576129bb6128cd565b604052919050565b8035600781106129d257600080fd5b919050565b600082601f8301126129e857600080fd5b81356001600160401b03811115612a0157612a016128cd565b612a14601f8201601f1916602001612993565b818152846020838601011115612a2957600080fd5b816020850160208301376000918101602001919091529392505050565b600060a08284031215612a5857600080fd5b60405160a081016001600160401b038282108183111715612a7b57612a7b6128cd565b81604052829350612a8b856129c3565b8352612a99602086016129c3565b6020840152612aaa604086016129c3565b60408401526060850135915060098210612ac357600080fd5b8160608401526080850135915080821115612add57600080fd5b50612aea858286016129d7565b6080830152505092915050565b60006001600160401b03821115612b1057612b106128cd565b5060051b60200190565b8035600381900b81146129d257600080fd5b600082601f830112612b3d57600080fd5b81356020612b52612b4d83612af7565b612993565b82815260059290921b84018101918181019086841115612b7157600080fd5b8286015b84811015612b9357612b8681612b1a565b8352918301918301612b75565b509695505050505050565b600060808284031215612bb057600080fd5b612bb86128e3565b905081356001600160401b0380821115612bd157600080fd5b612bdd85838601612a46565b83526020840135915080821115612bf357600080fd5b9083019060c08286031215612c0757600080fd5b612c0f61290b565b823582811115612c1e57600080fd5b612c2a87828601612b2c565b825250612c3960208401612b1a565b6020820152612c4a60408401612b1a565b6040820152612c5b60608401612b1a565b6060820152608083013582811115612c7257600080fd5b612c7e878286016129d7565b608083015250612c9060a084016129c3565b60a0820152602084015250612ca9905060408301612b1a565b6040820152612cba60608301612b1a565b606082015292915050565b600082601f830112612cd657600080fd5b81356020612ce6612b4d83612af7565b82815260059290921b84018101918181019086841115612d0557600080fd5b8286015b84811015612b935780356001600160401b0380821115612d295760008081fd5b908801906060828b03601f1901811315612d435760008081fd5b612d4b61292d565b612d568885016129c3565b815260408085013584811115612d6c5760008081fd5b612d7a8e8b838901016129d7565b838b015250918401359183831115612d925760008081fd5b612da08d8a858801016129d7565b908201528652505050918301918301612d09565b600060808284031215612dc657600080fd5b612dce6128e3565b905081356001600160401b0380821115612de757600080fd5b612df3858386016129d7565b83526020840135915080821115612e0957600080fd5b612e15858386016129d7565b60208401526040840135915080821115612e2e57600080fd5b612e3a85838601612a46565b60408401526060840135915080821115612e5357600080fd5b50612e6084828501612cc5565b60608301525092915050565b600060608284031215612e7e57600080fd5b612e8661292d565b905081356001600160401b0380821115612e9f57600080fd5b612eab858386016129d7565b83526020840135915080821115612ec157600080fd5b612ecd85838601612db4565b60208401526040840135915080821115612ee657600080fd5b50612ef384828501612db4565b60408301525092915050565b60006020808385031215612f1257600080fd5b612f1a61294f565b915082356001600160401b0380821115612f3357600080fd5b818501915085601f830112612f4757600080fd5b8135612f55612b4d82612af7565b81815260059190911b83018401908481019088831115612f7457600080fd5b8585015b8381101561300757803585811115612f905760008081fd5b86016040818c03601f1901811315612fa85760008081fd5b612fb0612971565b8983013588811115612fc25760008081fd5b612fd08e8c83870101612db4565b825250908201359087821115612fe65760008081fd5b612ff48d8b84860101612e6c565b818b015285525050918601918601612f78565b50865250939695505050505050565b60006080828403121561302857600080fd5b6130306128e3565b905081356001600160401b038082111561304957600080fd5b613055858386016129d7565b8352602084013591508082111561306b57600080fd5b613077858386016129d7565b6020840152604084013591508082111561309057600080fd5b61309c85838601612a46565b604084015260608401359150808211156130b557600080fd5b50612e6084828501612b2c565b6000604082840312156130d457600080fd5b6130dc612971565b905081356001600160401b03808211156130f557600080fd5b818401915084601f83011261310957600080fd5b81356020613119612b4d83612af7565b82815260059290921b8401810191818101908884111561313857600080fd5b8286015b848110156132495780358681111561315357600080fd5b8701601f196040828d038201121561316a57600080fd5b613172612971565b868301358981111561318357600080fd5b6131918e8983870101613016565b8252506040830135898111156131a657600080fd5b92909201916060838e03830112156131bd57600080fd5b6131c561292d565b915086830135898111156131d857600080fd5b6131e68e89838701016129d7565b8352506040830135898111156131fb57600080fd5b6132098e8983870101613016565b888401525060608301358981111561322057600080fd5b61322e8e8983870101613016565b6040840152508087019190915284525091830191830161313c565b508652508581013593508284111561326057600080fd5b61326c87858801612cc5565b818601525050505092915050565b60006080828403121561328c57600080fd5b6132946128e3565b905081356001600160401b03808211156132ad57600080fd5b6132b985838601612db4565b835260208401359150808211156132cf57600080fd5b6132db85838601612e6c565b602084015260408401359150808211156132f457600080fd5b61330085838601612eff565b6040840152606084013591508082111561331957600080fd5b50612e60848285016130c2565b6000806000806080858703121561333c57600080fd5b6001600160401b03808635111561335257600080fd5b61335f8787358801612b9e565b945060208601358181111561337357600080fd5b61337f888289016129d7565b94505060408601358181111561339457600080fd5b6133a08882890161327a565b9350506060860135818111156133b557600080fd5b8601601f810188136133c657600080fd5b6133d3612b4d8235612af7565b81358082526020808301929160051b8401018a8111156133f257600080fd5b602084015b8181101561348e57858135111561340d57600080fd5b803585016040818e03601f1901121561342557600080fd5b61342d612971565b60208201358881111561343f57600080fd5b61344e8f6020838601016129d7565b82525060408201358881111561346357600080fd5b6134728f6020838601016129d7565b60208301525080865250506020840193506020810190506133f7565b50508094505050505092959194509250565b600080600080600060a086880312156134b857600080fd5b85356001600160401b03808211156134cf57600080fd5b6134db89838a01612b9e565b965060208801359150808211156134f157600080fd5b6134fd89838a016129d7565b9550604088013591508082111561351357600080fd5b61351f89838a0161327a565b9450606088013591508082111561353557600080fd5b61354189838a016129d7565b9350608088013591508082111561355757600080fd5b50613564888289016129d7565b9150509295509295909350565b6000806000806080858703121561358757600080fd5b84356001600160401b038082111561359e57600080fd5b6135aa88838901612b9e565b95506020915081870135818111156135c157600080fd5b6135cd89828a016129d7565b9550506040870135818111156135e257600080fd5b6135ee89828a0161327a565b94505060608701358181111561360357600080fd5b8701601f8101891361361457600080fd5b8035613622612b4d82612af7565b81815260059190911b8201840190848101908b83111561364157600080fd5b8584015b838110156136795780358681111561365d5760008081fd5b61366b8e89838901016129d7565b845250918601918601613645565b50989b979a50959850505050505050565b600080600080608085870312156136a057600080fd5b84356001600160401b03808211156136b757600080fd5b6136c388838901612b9e565b955060208701359150808211156136d957600080fd5b6136e5888389016129d7565b945060408701359150808211156136fb57600080fd5b6137078883890161327a565b9350606087013591508082111561371d57600080fd5b5061372a878288016129d7565b91505092959194509250565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156137765761377661374c565b5060010190565b60208082526024908201527f4578697374656e63652050726f6f66206e6565647320646566696e6564204c65604082015263061664f760e41b606082015260800190565b6000828210156137d3576137d361374c565b500390565b634e487b7160e01b600052602160045260246000fd5b60008160001904831182151516156138085761380861374c565b500290565b600082198211156138205761382061374c565b500190565b6000815160005b81811015613846576020818501810151868301520161382c565b81811115613855576000828601525b509290920192915050565b6000611e0d6138786138728488613825565b86613825565b84613825565b600061042e8284613825565b60006020828403121561389c57600080fd5b5051919050565b60006119a0613878838661382556fe636865636b416761696e73745370656320666f72204c6561664f70202d20556ea2646970667358221220f9c7e74d71d1743a7bf76ba76ad4b180ad150d07f1a37c46e10ca7be9c19f0d364736f6c63430008090033"

// DeployICS23UnitTest deploys a new Ethereum contract, binding an instance of ICS23UnitTest to it.
func DeployICS23UnitTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ICS23UnitTest, error) {
	parsed, err := abi.JSON(strings.NewReader(ICS23UnitTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ICS23UnitTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ICS23UnitTest{ICS23UnitTestCaller: ICS23UnitTestCaller{contract: contract}, ICS23UnitTestTransactor: ICS23UnitTestTransactor{contract: contract}, ICS23UnitTestFilterer: ICS23UnitTestFilterer{contract: contract}}, nil
}

// ICS23UnitTest is an auto generated Go binding around an Ethereum contract.
type ICS23UnitTest struct {
	ICS23UnitTestCaller     // Read-only binding to the contract
	ICS23UnitTestTransactor // Write-only binding to the contract
	ICS23UnitTestFilterer   // Log filterer for contract events
}

// ICS23UnitTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICS23UnitTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS23UnitTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICS23UnitTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS23UnitTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICS23UnitTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS23UnitTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICS23UnitTestSession struct {
	Contract     *ICS23UnitTest    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICS23UnitTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICS23UnitTestCallerSession struct {
	Contract *ICS23UnitTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ICS23UnitTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICS23UnitTestTransactorSession struct {
	Contract     *ICS23UnitTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ICS23UnitTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICS23UnitTestRaw struct {
	Contract *ICS23UnitTest // Generic contract binding to access the raw methods on
}

// ICS23UnitTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICS23UnitTestCallerRaw struct {
	Contract *ICS23UnitTestCaller // Generic read-only contract binding to access the raw methods on
}

// ICS23UnitTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICS23UnitTestTransactorRaw struct {
	Contract *ICS23UnitTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICS23UnitTest creates a new instance of ICS23UnitTest, bound to a specific deployed contract.
func NewICS23UnitTest(address common.Address, backend bind.ContractBackend) (*ICS23UnitTest, error) {
	contract, err := bindICS23UnitTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICS23UnitTest{ICS23UnitTestCaller: ICS23UnitTestCaller{contract: contract}, ICS23UnitTestTransactor: ICS23UnitTestTransactor{contract: contract}, ICS23UnitTestFilterer: ICS23UnitTestFilterer{contract: contract}}, nil
}

// NewICS23UnitTestCaller creates a new read-only instance of ICS23UnitTest, bound to a specific deployed contract.
func NewICS23UnitTestCaller(address common.Address, caller bind.ContractCaller) (*ICS23UnitTestCaller, error) {
	contract, err := bindICS23UnitTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICS23UnitTestCaller{contract: contract}, nil
}

// NewICS23UnitTestTransactor creates a new write-only instance of ICS23UnitTest, bound to a specific deployed contract.
func NewICS23UnitTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ICS23UnitTestTransactor, error) {
	contract, err := bindICS23UnitTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICS23UnitTestTransactor{contract: contract}, nil
}

// NewICS23UnitTestFilterer creates a new log filterer instance of ICS23UnitTest, bound to a specific deployed contract.
func NewICS23UnitTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ICS23UnitTestFilterer, error) {
	contract, err := bindICS23UnitTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICS23UnitTestFilterer{contract: contract}, nil
}

// bindICS23UnitTest binds a generic wrapper to an already deployed contract.
func bindICS23UnitTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICS23UnitTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS23UnitTest *ICS23UnitTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS23UnitTest.Contract.ICS23UnitTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS23UnitTest *ICS23UnitTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS23UnitTest.Contract.ICS23UnitTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS23UnitTest *ICS23UnitTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS23UnitTest.Contract.ICS23UnitTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS23UnitTest *ICS23UnitTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS23UnitTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS23UnitTest *ICS23UnitTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS23UnitTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS23UnitTest *ICS23UnitTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS23UnitTest.Contract.contract.Transact(opts, method, params...)
}

// BatchVerifyMembership is a free data retrieval call binding the contract method 0x3aaa08eb.
//
// Solidity: function batchVerifyMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, (bytes,bytes)[] items) pure returns()
func (_ICS23UnitTest *ICS23UnitTestCaller) BatchVerifyMembership(opts *bind.CallOpts, spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, items []Ics23BatchItem) error {
	var out []interface{}
	err := _ICS23UnitTest.contract.Call(opts, &out, "batchVerifyMembership", spec, commitmentRoot, proof, items)

	if err != nil {
		return err
	}

	return err

}

// BatchVerifyMembership is a free data retrieval call binding the contract method 0x3aaa08eb.
//
// Solidity: function batchVerifyMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, (bytes,bytes)[] items) pure returns()
func (_ICS23UnitTest *ICS23UnitTestSession) BatchVerifyMembership(spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, items []Ics23BatchItem) error {
	return _ICS23UnitTest.Contract.BatchVerifyMembership(&_ICS23UnitTest.CallOpts, spec, commitmentRoot, proof, items)
}

// BatchVerifyMembership is a free data retrieval call binding the contract method 0x3aaa08eb.
//
// Solidity: function batchVerifyMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, (bytes,bytes)[] items) pure returns()
func (_ICS23UnitTest *ICS23UnitTestCallerSession) BatchVerifyMembership(spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, items []Ics23BatchItem) error {
	return _ICS23UnitTest.Contract.BatchVerifyMembership(&_ICS23UnitTest.CallOpts, spec, commitmentRoot, proof, items)
}

// BatchVerifyNonMembership is a free data retrieval call binding the contract method 0xb24b164b.
//
// Solidity: function batchVerifyNonMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, bytes[] keys) pure returns()
func (_ICS23UnitTest *ICS23UnitTestCaller) BatchVerifyNonMembership(opts *bind.CallOpts, spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, keys [][]byte) error {
	var out []interface{}
	err := _ICS23UnitTest.contract.Call(opts, &out, "batchVerifyNonMembership", spec, commitmentRoot, proof, keys)

	if err != nil {
		return err
	}

	return err

}

// BatchVerifyNonMembership is a free data retrieval call binding the contract method 0xb24b164b.
//
// Solidity: function batchVerifyNonMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, bytes[] keys) pure returns()
func (_ICS23UnitTest *ICS23UnitTestSession) BatchVerifyNonMembership(spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, keys [][]byte) error {
	return _ICS23UnitTest.Contract.BatchVerifyNonMembership(&_ICS23UnitTest.CallOpts, spec, commitmentRoot, proof, keys)
}

// BatchVerifyNonMembership is a free data retrieval call binding the contract method 0xb24b164b.
//
// Solidity: function batchVerifyNonMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, bytes[] keys) pure returns()
func (_ICS23UnitTest *ICS23UnitTestCallerSession) BatchVerifyNonMembership(spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, keys [][]byte) error {
	return _ICS23UnitTest.Contract.BatchVerifyNonMembership(&_ICS23UnitTest.CallOpts, spec, commitmentRoot, proof, keys)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x6b6c6818.
//
// Solidity: function verifyMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, bytes key, bytes value) pure returns()
func (_ICS23UnitTest *ICS23UnitTestCaller) VerifyMembership(opts *bind.CallOpts, spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, key []byte, value []byte) error {
	var out []interface{}
	err := _ICS23UnitTest.contract.Call(opts, &out, "verifyMembership", spec, commitmentRoot, proof, key, value)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0x6b6c6818.
//
// Solidity: function verifyMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, bytes key, bytes value) pure returns()
func (_ICS23UnitTest *ICS23UnitTestSession) VerifyMembership(spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, key []byte, value []byte) error {
	return _ICS23UnitTest.Contract.VerifyMembership(&_ICS23UnitTest.CallOpts, spec, commitmentRoot, proof, key, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x6b6c6818.
//
// Solidity: function verifyMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, bytes key, bytes value) pure returns()
func (_ICS23UnitTest *ICS23UnitTestCallerSession) VerifyMembership(spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, key []byte, value []byte) error {
	return _ICS23UnitTest.Contract.VerifyMembership(&_ICS23UnitTest.CallOpts, spec, commitmentRoot, proof, key, value)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xc4cd136f.
//
// Solidity: function verifyNonMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, bytes key) pure returns()
func (_ICS23UnitTest *ICS23UnitTestCaller) VerifyNonMembership(opts *bind.CallOpts, spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, key []byte) error {
	var out []interface{}
	err := _ICS23UnitTest.contract.Call(opts, &out, "verifyNonMembership", spec, commitmentRoot, proof, key)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xc4cd136f.
//
// Solidity: function verifyNonMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, bytes key) pure returns()
func (_ICS23UnitTest *ICS23UnitTestSession) VerifyNonMembership(spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, key []byte) error {
	return _ICS23UnitTest.Contract.VerifyNonMembership(&_ICS23UnitTest.CallOpts, spec, commitmentRoot, proof, key)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xc4cd136f.
//
// Solidity: function verifyNonMembership(((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec, bytes commitmentRoot, ((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof, bytes key) pure returns()
func (_ICS23UnitTest *ICS23UnitTestCallerSession) VerifyNonMembership(spec ProofSpecData, commitmentRoot []byte, proof CommitmentProofData, key []byte) error {
	return _ICS23UnitTest.Contract.VerifyNonMembership(&_ICS23UnitTest.CallOpts, spec, commitmentRoot, proof, key)
}

// Ics23ABI is the input ABI used to generate the binding from.
const Ics23ABI = "[]"

// Ics23Bin is the compiled bytecode used for deploying new contracts.
var Ics23Bin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122082d6a8d3fe3355fe9d4229ac650e7cc86873fa1b671c4285717fcd23ed2b668364736f6c63430008090033"

// DeployIcs23 deploys a new Ethereum contract, binding an instance of Ics23 to it.
func DeployIcs23(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ics23, error) {
	parsed, err := abi.JSON(strings.NewReader(Ics23ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Ics23Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ics23{Ics23Caller: Ics23Caller{contract: contract}, Ics23Transactor: Ics23Transactor{contract: contract}, Ics23Filterer: Ics23Filterer{contract: contract}}, nil
}

// Ics23 is an auto generated Go binding around an Ethereum contract.
type Ics23 struct {
	Ics23Caller     // Read-only binding to the contract
	Ics23Transactor // Write-only binding to the contract
	Ics23Filterer   // Log filterer for contract events
}

// Ics23Caller is an auto generated read-only Go binding around an Ethereum contract.
type Ics23Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics23Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Ics23Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics23Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ics23Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ics23Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ics23Session struct {
	Contract     *Ics23            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ics23CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ics23CallerSession struct {
	Contract *Ics23Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Ics23TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ics23TransactorSession struct {
	Contract     *Ics23Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ics23Raw is an auto generated low-level Go binding around an Ethereum contract.
type Ics23Raw struct {
	Contract *Ics23 // Generic contract binding to access the raw methods on
}

// Ics23CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ics23CallerRaw struct {
	Contract *Ics23Caller // Generic read-only contract binding to access the raw methods on
}

// Ics23TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ics23TransactorRaw struct {
	Contract *Ics23Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIcs23 creates a new instance of Ics23, bound to a specific deployed contract.
func NewIcs23(address common.Address, backend bind.ContractBackend) (*Ics23, error) {
	contract, err := bindIcs23(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ics23{Ics23Caller: Ics23Caller{contract: contract}, Ics23Transactor: Ics23Transactor{contract: contract}, Ics23Filterer: Ics23Filterer{contract: contract}}, nil
}

// NewIcs23Caller creates a new read-only instance of Ics23, bound to a specific deployed contract.
func NewIcs23Caller(address common.Address, caller bind.ContractCaller) (*Ics23Caller, error) {
	contract, err := bindIcs23(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ics23Caller{contract: contract}, nil
}

// NewIcs23Transactor creates a new write-only instance of Ics23, bound to a specific deployed contract.
func NewIcs23Transactor(address common.Address, transactor bind.ContractTransactor) (*Ics23Transactor, error) {
	contract, err := bindIcs23(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ics23Transactor{contract: contract}, nil
}

// NewIcs23Filterer creates a new log filterer instance of Ics23, bound to a specific deployed contract.
func NewIcs23Filterer(address common.Address, filterer bind.ContractFilterer) (*Ics23Filterer, error) {
	contract, err := bindIcs23(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ics23Filterer{contract: contract}, nil
}

// bindIcs23 binds a generic wrapper to an already deployed contract.
func bindIcs23(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ics23ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics23 *Ics23Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics23.Contract.Ics23Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics23 *Ics23Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics23.Contract.Ics23Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics23 *Ics23Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics23.Contract.Ics23Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ics23 *Ics23CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ics23.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ics23 *Ics23TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ics23.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ics23 *Ics23TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ics23.Contract.contract.Transact(opts, method, params...)
}

// InnerOpABI is the input ABI used to generate the binding from.
const InnerOpABI = "[]"

// InnerOpBin is the compiled bytecode used for deploying new contracts.
var InnerOpBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c50cb7b7031f3c2707a8893432bd7d26e3c0fb906190866b34bfc5c338b9f89d64736f6c63430008090033"

// DeployInnerOp deploys a new Ethereum contract, binding an instance of InnerOp to it.
func DeployInnerOp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InnerOp, error) {
	parsed, err := abi.JSON(strings.NewReader(InnerOpABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InnerOpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InnerOp{InnerOpCaller: InnerOpCaller{contract: contract}, InnerOpTransactor: InnerOpTransactor{contract: contract}, InnerOpFilterer: InnerOpFilterer{contract: contract}}, nil
}

// InnerOp is an auto generated Go binding around an Ethereum contract.
type InnerOp struct {
	InnerOpCaller     // Read-only binding to the contract
	InnerOpTransactor // Write-only binding to the contract
	InnerOpFilterer   // Log filterer for contract events
}

// InnerOpCaller is an auto generated read-only Go binding around an Ethereum contract.
type InnerOpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerOpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InnerOpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerOpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InnerOpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerOpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InnerOpSession struct {
	Contract     *InnerOp          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InnerOpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InnerOpCallerSession struct {
	Contract *InnerOpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// InnerOpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InnerOpTransactorSession struct {
	Contract     *InnerOpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InnerOpRaw is an auto generated low-level Go binding around an Ethereum contract.
type InnerOpRaw struct {
	Contract *InnerOp // Generic contract binding to access the raw methods on
}

// InnerOpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InnerOpCallerRaw struct {
	Contract *InnerOpCaller // Generic read-only contract binding to access the raw methods on
}

// InnerOpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InnerOpTransactorRaw struct {
	Contract *InnerOpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInnerOp creates a new instance of InnerOp, bound to a specific deployed contract.
func NewInnerOp(address common.Address, backend bind.ContractBackend) (*InnerOp, error) {
	contract, err := bindInnerOp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InnerOp{InnerOpCaller: InnerOpCaller{contract: contract}, InnerOpTransactor: InnerOpTransactor{contract: contract}, InnerOpFilterer: InnerOpFilterer{contract: contract}}, nil
}

// NewInnerOpCaller creates a new read-only instance of InnerOp, bound to a specific deployed contract.
func NewInnerOpCaller(address common.Address, caller bind.ContractCaller) (*InnerOpCaller, error) {
	contract, err := bindInnerOp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InnerOpCaller{contract: contract}, nil
}

// NewInnerOpTransactor creates a new write-only instance of InnerOp, bound to a specific deployed contract.
func NewInnerOpTransactor(address common.Address, transactor bind.ContractTransactor) (*InnerOpTransactor, error) {
	contract, err := bindInnerOp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InnerOpTransactor{contract: contract}, nil
}

// NewInnerOpFilterer creates a new log filterer instance of InnerOp, bound to a specific deployed contract.
func NewInnerOpFilterer(address common.Address, filterer bind.ContractFilterer) (*InnerOpFilterer, error) {
	contract, err := bindInnerOp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InnerOpFilterer{contract: contract}, nil
}

// bindInnerOp binds a generic wrapper to an already deployed contract.
func bindInnerOp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InnerOpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InnerOp *InnerOpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InnerOp.Contract.InnerOpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InnerOp *InnerOpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InnerOp.Contract.InnerOpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InnerOp *InnerOpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InnerOp.Contract.InnerOpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InnerOp *InnerOpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InnerOp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InnerOp *InnerOpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InnerOp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InnerOp *InnerOpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InnerOp.Contract.contract.Transact(opts, method, params...)
}

// InnerSpecABI is the input ABI used to generate the binding from.
const InnerSpecABI = "[]"

// InnerSpecBin is the compiled bytecode used for deploying new contracts.
var InnerSpecBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220841913e9bd3963243f4e119ed6f4dd88e75e43824fd25d15e1adb3c004676a7164736f6c63430008090033"

// DeployInnerSpec deploys a new Ethereum contract, binding an instance of InnerSpec to it.
func DeployInnerSpec(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InnerSpec, error) {
	parsed, err := abi.JSON(strings.NewReader(InnerSpecABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InnerSpecBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InnerSpec{InnerSpecCaller: InnerSpecCaller{contract: contract}, InnerSpecTransactor: InnerSpecTransactor{contract: contract}, InnerSpecFilterer: InnerSpecFilterer{contract: contract}}, nil
}

// InnerSpec is an auto generated Go binding around an Ethereum contract.
type InnerSpec struct {
	InnerSpecCaller     // Read-only binding to the contract
	InnerSpecTransactor // Write-only binding to the contract
	InnerSpecFilterer   // Log filterer for contract events
}

// InnerSpecCaller is an auto generated read-only Go binding around an Ethereum contract.
type InnerSpecCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerSpecTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InnerSpecTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerSpecFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InnerSpecFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InnerSpecSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InnerSpecSession struct {
	Contract     *InnerSpec        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InnerSpecCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InnerSpecCallerSession struct {
	Contract *InnerSpecCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// InnerSpecTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InnerSpecTransactorSession struct {
	Contract     *InnerSpecTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InnerSpecRaw is an auto generated low-level Go binding around an Ethereum contract.
type InnerSpecRaw struct {
	Contract *InnerSpec // Generic contract binding to access the raw methods on
}

// InnerSpecCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InnerSpecCallerRaw struct {
	Contract *InnerSpecCaller // Generic read-only contract binding to access the raw methods on
}

// InnerSpecTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InnerSpecTransactorRaw struct {
	Contract *InnerSpecTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInnerSpec creates a new instance of InnerSpec, bound to a specific deployed contract.
func NewInnerSpec(address common.Address, backend bind.ContractBackend) (*InnerSpec, error) {
	contract, err := bindInnerSpec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InnerSpec{InnerSpecCaller: InnerSpecCaller{contract: contract}, InnerSpecTransactor: InnerSpecTransactor{contract: contract}, InnerSpecFilterer: InnerSpecFilterer{contract: contract}}, nil
}

// NewInnerSpecCaller creates a new read-only instance of InnerSpec, bound to a specific deployed contract.
func NewInnerSpecCaller(address common.Address, caller bind.ContractCaller) (*InnerSpecCaller, error) {
	contract, err := bindInnerSpec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InnerSpecCaller{contract: contract}, nil
}

// NewInnerSpecTransactor creates a new write-only instance of InnerSpec, bound to a specific deployed contract.
func NewInnerSpecTransactor(address common.Address, transactor bind.ContractTransactor) (*InnerSpecTransactor, error) {
	contract, err := bindInnerSpec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InnerSpecTransactor{contract: contract}, nil
}

// NewInnerSpecFilterer creates a new log filterer instance of InnerSpec, bound to a specific deployed contract.
func NewInnerSpecFilterer(address common.Address, filterer bind.ContractFilterer) (*InnerSpecFilterer, error) {
	contract, err := bindInnerSpec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InnerSpecFilterer{contract: contract}, nil
}

// bindInnerSpec binds a generic wrapper to an already deployed contract.
func bindInnerSpec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InnerSpecABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InnerSpec *InnerSpecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InnerSpec.Contract.InnerSpecCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InnerSpec *InnerSpecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InnerSpec.Contract.InnerSpecTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InnerSpec *InnerSpecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InnerSpec.Contract.InnerSpecTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InnerSpec *InnerSpecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InnerSpec.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InnerSpec *InnerSpecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InnerSpec.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InnerSpec *InnerSpecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InnerSpec.Contract.contract.Transact(opts, method, params...)
}

// LeafOpABI is the input ABI used to generate the binding from.
const LeafOpABI = "[]"

// LeafOpBin is the compiled bytecode used for deploying new contracts.
var LeafOpBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220081e16d19bc06cf99ee7a53ddb45fc227a6abc974454a66e0d9b332d4d094bf364736f6c63430008090033"

// DeployLeafOp deploys a new Ethereum contract, binding an instance of LeafOp to it.
func DeployLeafOp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LeafOp, error) {
	parsed, err := abi.JSON(strings.NewReader(LeafOpABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LeafOpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LeafOp{LeafOpCaller: LeafOpCaller{contract: contract}, LeafOpTransactor: LeafOpTransactor{contract: contract}, LeafOpFilterer: LeafOpFilterer{contract: contract}}, nil
}

// LeafOp is an auto generated Go binding around an Ethereum contract.
type LeafOp struct {
	LeafOpCaller     // Read-only binding to the contract
	LeafOpTransactor // Write-only binding to the contract
	LeafOpFilterer   // Log filterer for contract events
}

// LeafOpCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeafOpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeafOpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeafOpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeafOpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeafOpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeafOpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeafOpSession struct {
	Contract     *LeafOp           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeafOpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeafOpCallerSession struct {
	Contract *LeafOpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LeafOpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeafOpTransactorSession struct {
	Contract     *LeafOpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeafOpRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeafOpRaw struct {
	Contract *LeafOp // Generic contract binding to access the raw methods on
}

// LeafOpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeafOpCallerRaw struct {
	Contract *LeafOpCaller // Generic read-only contract binding to access the raw methods on
}

// LeafOpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeafOpTransactorRaw struct {
	Contract *LeafOpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeafOp creates a new instance of LeafOp, bound to a specific deployed contract.
func NewLeafOp(address common.Address, backend bind.ContractBackend) (*LeafOp, error) {
	contract, err := bindLeafOp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LeafOp{LeafOpCaller: LeafOpCaller{contract: contract}, LeafOpTransactor: LeafOpTransactor{contract: contract}, LeafOpFilterer: LeafOpFilterer{contract: contract}}, nil
}

// NewLeafOpCaller creates a new read-only instance of LeafOp, bound to a specific deployed contract.
func NewLeafOpCaller(address common.Address, caller bind.ContractCaller) (*LeafOpCaller, error) {
	contract, err := bindLeafOp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeafOpCaller{contract: contract}, nil
}

// NewLeafOpTransactor creates a new write-only instance of LeafOp, bound to a specific deployed contract.
func NewLeafOpTransactor(address common.Address, transactor bind.ContractTransactor) (*LeafOpTransactor, error) {
	contract, err := bindLeafOp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeafOpTransactor{contract: contract}, nil
}

// NewLeafOpFilterer creates a new log filterer instance of LeafOp, bound to a specific deployed contract.
func NewLeafOpFilterer(address common.Address, filterer bind.ContractFilterer) (*LeafOpFilterer, error) {
	contract, err := bindLeafOp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeafOpFilterer{contract: contract}, nil
}

// bindLeafOp binds a generic wrapper to an already deployed contract.
func bindLeafOp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LeafOpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeafOp *LeafOpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeafOp.Contract.LeafOpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeafOp *LeafOpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeafOp.Contract.LeafOpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeafOp *LeafOpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeafOp.Contract.LeafOpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeafOp *LeafOpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeafOp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeafOp *LeafOpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeafOp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeafOp *LeafOpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeafOp.Contract.contract.Transact(opts, method, params...)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
var MathBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c0b7fa1476bcc4110f58f2af7f2f2b434ca0e067d749bf759fcdb20a138c419864736f6c63430008090033"

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// NonExistenceProofABI is the input ABI used to generate the binding from.
const NonExistenceProofABI = "[]"

// NonExistenceProofBin is the compiled bytecode used for deploying new contracts.
var NonExistenceProofBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b0871376b5278f80453269ac44a0fc1a6c9595cdc792382f7e9509cac44fcbd364736f6c63430008090033"

// DeployNonExistenceProof deploys a new Ethereum contract, binding an instance of NonExistenceProof to it.
func DeployNonExistenceProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NonExistenceProof, error) {
	parsed, err := abi.JSON(strings.NewReader(NonExistenceProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NonExistenceProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NonExistenceProof{NonExistenceProofCaller: NonExistenceProofCaller{contract: contract}, NonExistenceProofTransactor: NonExistenceProofTransactor{contract: contract}, NonExistenceProofFilterer: NonExistenceProofFilterer{contract: contract}}, nil
}

// NonExistenceProof is an auto generated Go binding around an Ethereum contract.
type NonExistenceProof struct {
	NonExistenceProofCaller     // Read-only binding to the contract
	NonExistenceProofTransactor // Write-only binding to the contract
	NonExistenceProofFilterer   // Log filterer for contract events
}

// NonExistenceProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type NonExistenceProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonExistenceProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NonExistenceProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonExistenceProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NonExistenceProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NonExistenceProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NonExistenceProofSession struct {
	Contract     *NonExistenceProof // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NonExistenceProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NonExistenceProofCallerSession struct {
	Contract *NonExistenceProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NonExistenceProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NonExistenceProofTransactorSession struct {
	Contract     *NonExistenceProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NonExistenceProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type NonExistenceProofRaw struct {
	Contract *NonExistenceProof // Generic contract binding to access the raw methods on
}

// NonExistenceProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NonExistenceProofCallerRaw struct {
	Contract *NonExistenceProofCaller // Generic read-only contract binding to access the raw methods on
}

// NonExistenceProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NonExistenceProofTransactorRaw struct {
	Contract *NonExistenceProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNonExistenceProof creates a new instance of NonExistenceProof, bound to a specific deployed contract.
func NewNonExistenceProof(address common.Address, backend bind.ContractBackend) (*NonExistenceProof, error) {
	contract, err := bindNonExistenceProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NonExistenceProof{NonExistenceProofCaller: NonExistenceProofCaller{contract: contract}, NonExistenceProofTransactor: NonExistenceProofTransactor{contract: contract}, NonExistenceProofFilterer: NonExistenceProofFilterer{contract: contract}}, nil
}

// NewNonExistenceProofCaller creates a new read-only instance of NonExistenceProof, bound to a specific deployed contract.
func NewNonExistenceProofCaller(address common.Address, caller bind.ContractCaller) (*NonExistenceProofCaller, error) {
	contract, err := bindNonExistenceProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NonExistenceProofCaller{contract: contract}, nil
}

// NewNonExistenceProofTransactor creates a new write-only instance of NonExistenceProof, bound to a specific deployed contract.
func NewNonExistenceProofTransactor(address common.Address, transactor bind.ContractTransactor) (*NonExistenceProofTransactor, error) {
	contract, err := bindNonExistenceProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NonExistenceProofTransactor{contract: contract}, nil
}

// NewNonExistenceProofFilterer creates a new log filterer instance of NonExistenceProof, bound to a specific deployed contract.
func NewNonExistenceProofFilterer(address common.Address, filterer bind.ContractFilterer) (*NonExistenceProofFilterer, error) {
	contract, err := bindNonExistenceProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NonExistenceProofFilterer{contract: contract}, nil
}

// bindNonExistenceProof binds a generic wrapper to an already deployed contract.
func bindNonExistenceProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NonExistenceProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonExistenceProof *NonExistenceProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonExistenceProof.Contract.NonExistenceProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonExistenceProof *NonExistenceProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonExistenceProof.Contract.NonExistenceProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonExistenceProof *NonExistenceProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonExistenceProof.Contract.NonExistenceProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NonExistenceProof *NonExistenceProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NonExistenceProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NonExistenceProof *NonExistenceProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NonExistenceProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NonExistenceProof *NonExistenceProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NonExistenceProof.Contract.contract.Transact(opts, method, params...)
}

// OpsABI is the input ABI used to generate the binding from.
const OpsABI = "[]"

// OpsBin is the compiled bytecode used for deploying new contracts.
var OpsBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200ac63aa3542ac73ce637671cd140c6534398da947e2d46ef115cec16635cd89564736f6c63430008090033"

// DeployOps deploys a new Ethereum contract, binding an instance of Ops to it.
func DeployOps(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ops, error) {
	parsed, err := abi.JSON(strings.NewReader(OpsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OpsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ops{OpsCaller: OpsCaller{contract: contract}, OpsTransactor: OpsTransactor{contract: contract}, OpsFilterer: OpsFilterer{contract: contract}}, nil
}

// Ops is an auto generated Go binding around an Ethereum contract.
type Ops struct {
	OpsCaller     // Read-only binding to the contract
	OpsTransactor // Write-only binding to the contract
	OpsFilterer   // Log filterer for contract events
}

// OpsCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpsSession struct {
	Contract     *Ops              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpsCallerSession struct {
	Contract *OpsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OpsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpsTransactorSession struct {
	Contract     *OpsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpsRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpsRaw struct {
	Contract *Ops // Generic contract binding to access the raw methods on
}

// OpsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpsCallerRaw struct {
	Contract *OpsCaller // Generic read-only contract binding to access the raw methods on
}

// OpsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpsTransactorRaw struct {
	Contract *OpsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOps creates a new instance of Ops, bound to a specific deployed contract.
func NewOps(address common.Address, backend bind.ContractBackend) (*Ops, error) {
	contract, err := bindOps(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ops{OpsCaller: OpsCaller{contract: contract}, OpsTransactor: OpsTransactor{contract: contract}, OpsFilterer: OpsFilterer{contract: contract}}, nil
}

// NewOpsCaller creates a new read-only instance of Ops, bound to a specific deployed contract.
func NewOpsCaller(address common.Address, caller bind.ContractCaller) (*OpsCaller, error) {
	contract, err := bindOps(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpsCaller{contract: contract}, nil
}

// NewOpsTransactor creates a new write-only instance of Ops, bound to a specific deployed contract.
func NewOpsTransactor(address common.Address, transactor bind.ContractTransactor) (*OpsTransactor, error) {
	contract, err := bindOps(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpsTransactor{contract: contract}, nil
}

// NewOpsFilterer creates a new log filterer instance of Ops, bound to a specific deployed contract.
func NewOpsFilterer(address common.Address, filterer bind.ContractFilterer) (*OpsFilterer, error) {
	contract, err := bindOps(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpsFilterer{contract: contract}, nil
}

// bindOps binds a generic wrapper to an already deployed contract.
func bindOps(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ops *OpsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ops.Contract.OpsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ops *OpsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ops.Contract.OpsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ops *OpsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ops.Contract.OpsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ops *OpsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ops.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ops *OpsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ops.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ops *OpsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ops.Contract.contract.Transact(opts, method, params...)
}

// OpsUnitTestABI is the input ABI used to generate the binding from.
const OpsUnitTestABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"applyOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data\",\"name\":\"inner\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"child\",\"type\":\"bytes\"}],\"name\":\"applyOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"op\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf_spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int32[]\",\"name\":\"child_order\",\"type\":\"int32[]\"},{\"internalType\":\"int32\",\"name\":\"child_size\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"max_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"empty_child\",\"type\":\"bytes\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structInnerSpec.Data\",\"name\":\"inner_spec\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"max_depth\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_depth\",\"type\":\"int32\"}],\"internalType\":\"structProofSpec.Data\",\"name\":\"spec\",\"type\":\"tuple\"}],\"name\":\"checkAgainstSpec\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data\",\"name\":\"op\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf_spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int32[]\",\"name\":\"child_order\",\"type\":\"int32[]\"},{\"internalType\":\"int32\",\"name\":\"child_size\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"max_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"empty_child\",\"type\":\"bytes\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structInnerSpec.Data\",\"name\":\"inner_spec\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"max_depth\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_depth\",\"type\":\"int32\"}],\"internalType\":\"structProofSpec.Data\",\"name\":\"spec\",\"type\":\"tuple\"}],\"name\":\"checkAgainstSpec\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"compare\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hashOp\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"preImage\",\"type\":\"bytes\"}],\"name\":\"doHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OpsUnitTestBin is the compiled bytecode used for deploying new contracts.
var OpsUnitTestBin = "0x608060405234801561001057600080fd5b50611a21806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80635e7b22fa1461006757806381e79ac61461007c5780639e7fe568146100a2578063bd2a7e7f146100c2578063d48f1e4f146100d5578063f7717f42146100e8575b600080fd5b61007a61007536600461158a565b6100fb565b005b61008f61008a3660046115ed565b610109565b6040519081526020015b60405180910390f35b6100b56100b0366004611646565b61011e565b60405161009991906116fd565b6100b56100d03660046117ca565b610133565b6100b56100e3366004611800565b61013f565b61007a6100f6366004611843565b61014b565b6101058282610155565b5050565b600061011583836103f2565b90505b92915050565b606061012b8484846104ef565b949350505050565b606061011583836105f0565b6060610115838361067d565b61010582826109a2565b805151600681111561016957610169611879565b8251600681111561017c5761017c611879565b146101d45760405162461bcd60e51b815260206004820152602f60248201526000805160206119cc83398151915260448201526e0657870656374656420486173684f7608c1b60648201526084015b60405180910390fd5b80516020015160068111156101eb576101eb611879565b8260200151600681111561020157610201611879565b146102585760405162461bcd60e51b815260206004820152603360248201526000805160206119cc833981519152604482015272657870656374656420507265686173684b657960681b60648201526084016101cb565b805160400151600681111561026f5761026f611879565b8260400151600681111561028557610285611879565b146102de5760405162461bcd60e51b815260206004820152603560248201526000805160206119cc8339815191526044820152746578706563746564205072656861736856616c756560581b60648201526084016101cb565b80516060015160088111156102f5576102f5611879565b8260600151600881111561030b5761030b611879565b146103605760405162461bcd60e51b815260206004820152603160248201526000805160206119cc83398151915260448201527006578706563746564206c656e6774684f7607c1b60648201526084016101cb565b60006103788360800151836000015160800151610bd3565b9050806103ed5760405162461bcd60e51b815260206004820152603c60248201527f636865636b416761696e73745370656320666f72204c6561664f70202d204c6560448201527f61662050726566697820646f65736e277420737461727420776974680000000060648201526084016101cb565b505050565b60008061040184518451610c14565b905060005b818110156104bd578381815181106104205761042061188f565b602001015160f81c60f81b60f81c60ff168582815181106104435761044361188f565b016020015160f81c101561045d5760001992505050610118565b83818151811061046f5761046f61188f565b602001015160f81c60f81b60f81c60ff168582815181106104925761049261188f565b016020015160f81c11156104ab57600192505050610118565b806104b5816118bb565b915050610406565b5080845111156104d1576001915050610118565b80835111156104e557600019915050610118565b5060009392505050565b606060008351116105365760405162461bcd60e51b81526020600482015260116024820152704c656166206f70206e65656473206b657960781b60448201526064016101cb565b600082511161057d5760405162461bcd60e51b81526020600482015260136024820152724c656166206f70206e656564732076616c756560681b60448201526064016101cb565b60006105928560200151866060015186610c2a565b905060006105a98660400151876060015186610c2a565b90506000866080015183836040516020016105c6939291906118d6565b60405160208183030381529060405290506105e587600001518261067d565b979650505050505050565b606060008251116106435760405162461bcd60e51b815260206004820152601a60248201527f496e6e6572206f70206e65656473206368696c642076616c756500000000000060448201526064016101cb565b60008360200151838560400151604051602001610662939291906118d6565b604051602081830303815290604052905061012b8460000151825b6060600183600681111561069357610693611879565b1415610711576002826040516106a99190611919565b602060405180830381855afa1580156106c6573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906106e99190611935565b6040516020016106fb91815260200190565b6040516020818303038152906040529050610118565b600383600681111561072557610725611879565b14156107455781805190602001206040516020016106fb91815260200190565b600483600681111561075957610759611879565b14156107b45760038260405161076f9190611919565b602060405180830381855afa15801561078c573d6000803e3d6000fd5b5050604051805160601b6bffffffffffffffffffffffff1916602082015260340190506106fb565b60058360068111156107c8576107c8611879565b14156108aa5760006002836040516107e09190611919565b602060405180830381855afa1580156107fd573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906108209190611935565b60405160200161083291815260200190565b60405160208183030381529060405290506003816040516108539190611919565b602060405180830381855afa158015610870573d6000803e3d6000fd5b5050604051805160601b6bffffffffffffffffffffffff191660208201526034019050604051602081830303815290604052915050610118565b60028360068111156108be576108be611879565b14156109035760405162461bcd60e51b815260206004820152601460248201527314d2104d4c4c881b9bdd081cdd5c1c1bdc9d195960621b60448201526064016101cb565b600683600681111561091757610917611879565b14156109655760405162461bcd60e51b815260206004820152601860248201527f5348413531325f323536206e6f7420737570706f72746564000000000000000060448201526064016101cb565b60405162461bcd60e51b81526020600482015260126024820152710556e737570706f7274656420686173684f760741b60448201526064016101cb565b806020015160a0015160068111156109bc576109bc611879565b825160068111156109cf576109cf611879565b14610a355760405162461bcd60e51b815260206004820152603060248201527f636865636b416761696e73745370656320666f7220496e6e65724f70202d205560448201526f06e657870656374656420486173684f760841b60648201526084016101cb565b6000610a4b82602001516040015160030b610c4d565b9050808360200151511015610aa25760405162461bcd60e51b815260206004820152601860248201527f496e6e65724f702070726566697820746f6f2073686f7274000000000000000060448201526064016101cb565b8151608001516020840151600090610aba9083610bd3565b90508015610b165760405162461bcd60e51b8152602060048201526024808201527f496e6e6572205072656669782073746172747320776974682077726f6e672076604482015263616c756560e01b60648201526084016101cb565b6000610b2c85602001516020015160030b610c4d565b9050600081600187602001516000015151610b47919061194e565b610b519190611965565b90506000610b6987602001516060015160030b610c4d565b9050610b758282611984565b8860200151511115610bc95760405162461bcd60e51b815260206004820152601760248201527f496e6e65724f702070726566697820746f6f206c6f6e6700000000000000000060448201526064016101cb565b5050505050505050565b6000815160001415610be757506001610118565b825182511115610bf957506000610118565b6000610c088460008551610ca3565b905061012b8382610db0565b6000818310610c235781610115565b5090919050565b60606000610c388584610e14565b9050610c448482610e41565b95945050505050565b600080821215610c9f5760405162461bcd60e51b815260206004820181905260248201527f53616665436173743a2076616c7565206d75737420626520706f73697469766560448201526064016101cb565b5090565b606081610cb181601f611984565b1015610cf05760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b60448201526064016101cb565b610cfa8284611984565b84511015610d3e5760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b60448201526064016101cb565b606082158015610d5d5760405191506000825260208201604052610da7565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015610d96578051835260209283019201610d7e565b5050858452601f01601f1916604052505b50949350505050565b815181516000916001918114808314610dcc5760009250610e0a565b600160208701838101602088015b600284838510011415610e05578051835114610df95760009650600093505b60209283019201610dda565b505050505b5090949350505050565b60606000836006811115610e2a57610e2a611879565b1415610e37575080610118565b610115838361067d565b60606000836008811115610e5757610e57611879565b1415610e64575080610118565b6001836008811115610e7857610e78611879565b1415610f0c576000610e8a835161114d565b90506000816001600160401b03811115610ea657610ea6611212565b6040519080825280601f01601f191660200182016040528015610ed0576020820181803683370190505b509050610ee0845160208361116a565b508084604051602001610ef492919061199c565b60405160208183030381529060405292505050610118565b6007836008811115610f2057610f20611879565b1415610f72578151602014610f6b5760405162461bcd60e51b81526020600482015260116024820152703230ba30973632b733ba3410109e90199960791b60448201526064016101cb565b5080610118565b6008836008811115610f8657610f86611879565b1415610fd1578151604014610f6b5760405162461bcd60e51b815260206004820152601160248201527019185d184b9b195b99dd1a08084f480d8d607a1b60448201526064016101cb565b6004836008811115610fe557610fe5611879565b1415611111576000610ff783516111ad565b60408051600480825281830190925291925060e083901b916000916020820181803683370190505090508160031a60f81b8160008151811061103b5761103b61188f565b60200101906001600160f81b031916908160001a9053508160021a60f81b8160018151811061106c5761106c61188f565b60200101906001600160f81b031916908160001a9053508160011a60f81b8160028151811061109d5761109d61188f565b60200101906001600160f81b031916908160001a9053508160001a60f81b816003815181106110ce576110ce61188f565b60200101906001600160f81b031916908160001a90535080856040516020016110f892919061199c565b6040516020818303038152906040529350505050610118565b60405162461bcd60e51b81526020600482015260116024820152700556e737570706f72746564206c656e4f7607c1b60448201526064016101cb565b60071c600060015b82156101185760079290921c91600101611155565b600080828401607f86165b600787901c1561119d578060801782535060079590951c9460019182019101607f8616611175565b8082535050600101949350505050565b600063ffffffff821115610c9f5760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201526532206269747360d01b60648201526084016101cb565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b038111828210171561124a5761124a611212565b60405290565b60405160c081016001600160401b038111828210171561124a5761124a611212565b604051601f8201601f191681016001600160401b038111828210171561129a5761129a611212565b604052919050565b8035600781106112b157600080fd5b919050565b600082601f8301126112c757600080fd5b81356001600160401b038111156112e0576112e0611212565b6112f3601f8201601f1916602001611272565b81815284602083860101111561130857600080fd5b816020850160208301376000918101602001919091529392505050565b600060a0828403121561133757600080fd5b60405160a081016001600160401b03828210818311171561135a5761135a611212565b8160405282935061136a856112a2565b8352611378602086016112a2565b6020840152611389604086016112a2565b604084015260608501359150600982106113a257600080fd5b81606084015260808501359150808211156113bc57600080fd5b506113c9858286016112b6565b6080830152505092915050565b8035600381900b81146112b157600080fd5b600082601f8301126113f957600080fd5b813560206001600160401b0382111561141457611414611212565b8160051b611423828201611272565b928352848101820192828101908785111561143d57600080fd5b83870192505b848310156105e557611454836113d6565b82529183019190830190611443565b60006080828403121561147557600080fd5b61147d611228565b905081356001600160401b038082111561149657600080fd5b6114a285838601611325565b835260208401359150808211156114b857600080fd5b9083019060c082860312156114cc57600080fd5b6114d4611250565b8235828111156114e357600080fd5b6114ef878286016113e8565b8252506114fe602084016113d6565b602082015261150f604084016113d6565b6040820152611520606084016113d6565b606082015260808301358281111561153757600080fd5b611543878286016112b6565b60808301525061155560a084016112a2565b60a082015260208401525061156e9050604083016113d6565b604082015261157f606083016113d6565b606082015292915050565b6000806040838503121561159d57600080fd5b82356001600160401b03808211156115b457600080fd5b6115c086838701611325565b935060208501359150808211156115d657600080fd5b506115e385828601611463565b9150509250929050565b6000806040838503121561160057600080fd5b82356001600160401b038082111561161757600080fd5b611623868387016112b6565b9350602085013591508082111561163957600080fd5b506115e3858286016112b6565b60008060006060848603121561165b57600080fd5b83356001600160401b038082111561167257600080fd5b61167e87838801611325565b9450602086013591508082111561169457600080fd5b6116a0878388016112b6565b935060408601359150808211156116b657600080fd5b506116c3868287016112b6565b9150509250925092565b60005b838110156116e85781810151838201526020016116d0565b838111156116f7576000848401525b50505050565b602081526000825180602084015261171c8160408501602087016116cd565b601f01601f19169190910160400192915050565b60006060828403121561174257600080fd5b604051606081016001600160401b03828210818311171561176557611765611212565b81604052829350611775856112a2565b8352602085013591508082111561178b57600080fd5b611797868387016112b6565b602084015260408501359150808211156117b057600080fd5b506117bd858286016112b6565b6040830152505092915050565b600080604083850312156117dd57600080fd5b82356001600160401b03808211156117f457600080fd5b61162386838701611730565b6000806040838503121561181357600080fd5b61181c836112a2565b915060208301356001600160401b0381111561183757600080fd5b6115e3858286016112b6565b6000806040838503121561185657600080fd5b82356001600160401b038082111561186d57600080fd5b6115c086838701611730565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156118cf576118cf6118a5565b5060010190565b600084516118e88184602089016116cd565b8451908301906118fc8183602089016116cd565b845191019061190f8183602088016116cd565b0195945050505050565b6000825161192b8184602087016116cd565b9190910192915050565b60006020828403121561194757600080fd5b5051919050565b600082821015611960576119606118a5565b500390565b600081600019048311821515161561197f5761197f6118a5565b500290565b60008219821115611997576119976118a5565b500190565b600083516119ae8184602088016116cd565b8351908301906119c28183602088016116cd565b0194935050505056fe636865636b416761696e73745370656320666f72204c6561664f70202d20556ea264697066735822122018137ff09506aead04367c684948e74a2978c810b152d742f7078bf1fc350eb364736f6c63430008090033"

// DeployOpsUnitTest deploys a new Ethereum contract, binding an instance of OpsUnitTest to it.
func DeployOpsUnitTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OpsUnitTest, error) {
	parsed, err := abi.JSON(strings.NewReader(OpsUnitTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OpsUnitTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OpsUnitTest{OpsUnitTestCaller: OpsUnitTestCaller{contract: contract}, OpsUnitTestTransactor: OpsUnitTestTransactor{contract: contract}, OpsUnitTestFilterer: OpsUnitTestFilterer{contract: contract}}, nil
}

// OpsUnitTest is an auto generated Go binding around an Ethereum contract.
type OpsUnitTest struct {
	OpsUnitTestCaller     // Read-only binding to the contract
	OpsUnitTestTransactor // Write-only binding to the contract
	OpsUnitTestFilterer   // Log filterer for contract events
}

// OpsUnitTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpsUnitTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsUnitTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpsUnitTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsUnitTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpsUnitTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpsUnitTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpsUnitTestSession struct {
	Contract     *OpsUnitTest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpsUnitTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpsUnitTestCallerSession struct {
	Contract *OpsUnitTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OpsUnitTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpsUnitTestTransactorSession struct {
	Contract     *OpsUnitTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OpsUnitTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpsUnitTestRaw struct {
	Contract *OpsUnitTest // Generic contract binding to access the raw methods on
}

// OpsUnitTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpsUnitTestCallerRaw struct {
	Contract *OpsUnitTestCaller // Generic read-only contract binding to access the raw methods on
}

// OpsUnitTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpsUnitTestTransactorRaw struct {
	Contract *OpsUnitTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOpsUnitTest creates a new instance of OpsUnitTest, bound to a specific deployed contract.
func NewOpsUnitTest(address common.Address, backend bind.ContractBackend) (*OpsUnitTest, error) {
	contract, err := bindOpsUnitTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OpsUnitTest{OpsUnitTestCaller: OpsUnitTestCaller{contract: contract}, OpsUnitTestTransactor: OpsUnitTestTransactor{contract: contract}, OpsUnitTestFilterer: OpsUnitTestFilterer{contract: contract}}, nil
}

// NewOpsUnitTestCaller creates a new read-only instance of OpsUnitTest, bound to a specific deployed contract.
func NewOpsUnitTestCaller(address common.Address, caller bind.ContractCaller) (*OpsUnitTestCaller, error) {
	contract, err := bindOpsUnitTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpsUnitTestCaller{contract: contract}, nil
}

// NewOpsUnitTestTransactor creates a new write-only instance of OpsUnitTest, bound to a specific deployed contract.
func NewOpsUnitTestTransactor(address common.Address, transactor bind.ContractTransactor) (*OpsUnitTestTransactor, error) {
	contract, err := bindOpsUnitTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpsUnitTestTransactor{contract: contract}, nil
}

// NewOpsUnitTestFilterer creates a new log filterer instance of OpsUnitTest, bound to a specific deployed contract.
func NewOpsUnitTestFilterer(address common.Address, filterer bind.ContractFilterer) (*OpsUnitTestFilterer, error) {
	contract, err := bindOpsUnitTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpsUnitTestFilterer{contract: contract}, nil
}

// bindOpsUnitTest binds a generic wrapper to an already deployed contract.
func bindOpsUnitTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpsUnitTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpsUnitTest *OpsUnitTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpsUnitTest.Contract.OpsUnitTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpsUnitTest *OpsUnitTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpsUnitTest.Contract.OpsUnitTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpsUnitTest *OpsUnitTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpsUnitTest.Contract.OpsUnitTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OpsUnitTest *OpsUnitTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OpsUnitTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OpsUnitTest *OpsUnitTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OpsUnitTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OpsUnitTest *OpsUnitTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OpsUnitTest.Contract.contract.Transact(opts, method, params...)
}

// ApplyOp is a free data retrieval call binding the contract method 0x9e7fe568.
//
// Solidity: function applyOp((uint8,uint8,uint8,uint8,bytes) leaf, bytes key, bytes value) pure returns(bytes)
func (_OpsUnitTest *OpsUnitTestCaller) ApplyOp(opts *bind.CallOpts, leaf LeafOpData, key []byte, value []byte) ([]byte, error) {
	var out []interface{}
	err := _OpsUnitTest.contract.Call(opts, &out, "applyOp", leaf, key, value)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ApplyOp is a free data retrieval call binding the contract method 0x9e7fe568.
//
// Solidity: function applyOp((uint8,uint8,uint8,uint8,bytes) leaf, bytes key, bytes value) pure returns(bytes)
func (_OpsUnitTest *OpsUnitTestSession) ApplyOp(leaf LeafOpData, key []byte, value []byte) ([]byte, error) {
	return _OpsUnitTest.Contract.ApplyOp(&_OpsUnitTest.CallOpts, leaf, key, value)
}

// ApplyOp is a free data retrieval call binding the contract method 0x9e7fe568.
//
// Solidity: function applyOp((uint8,uint8,uint8,uint8,bytes) leaf, bytes key, bytes value) pure returns(bytes)
func (_OpsUnitTest *OpsUnitTestCallerSession) ApplyOp(leaf LeafOpData, key []byte, value []byte) ([]byte, error) {
	return _OpsUnitTest.Contract.ApplyOp(&_OpsUnitTest.CallOpts, leaf, key, value)
}

// ApplyOp0 is a free data retrieval call binding the contract method 0xbd2a7e7f.
//
// Solidity: function applyOp((uint8,bytes,bytes) inner, bytes child) pure returns(bytes)
func (_OpsUnitTest *OpsUnitTestCaller) ApplyOp0(opts *bind.CallOpts, inner InnerOpData, child []byte) ([]byte, error) {
	var out []interface{}
	err := _OpsUnitTest.contract.Call(opts, &out, "applyOp0", inner, child)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ApplyOp0 is a free data retrieval call binding the contract method 0xbd2a7e7f.
//
// Solidity: function applyOp((uint8,bytes,bytes) inner, bytes child) pure returns(bytes)
func (_OpsUnitTest *OpsUnitTestSession) ApplyOp0(inner InnerOpData, child []byte) ([]byte, error) {
	return _OpsUnitTest.Contract.ApplyOp0(&_OpsUnitTest.CallOpts, inner, child)
}

// ApplyOp0 is a free data retrieval call binding the contract method 0xbd2a7e7f.
//
// Solidity: function applyOp((uint8,bytes,bytes) inner, bytes child) pure returns(bytes)
func (_OpsUnitTest *OpsUnitTestCallerSession) ApplyOp0(inner InnerOpData, child []byte) ([]byte, error) {
	return _OpsUnitTest.Contract.ApplyOp0(&_OpsUnitTest.CallOpts, inner, child)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0x5e7b22fa.
//
// Solidity: function checkAgainstSpec((uint8,uint8,uint8,uint8,bytes) op, ((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec) pure returns()
func (_OpsUnitTest *OpsUnitTestCaller) CheckAgainstSpec(opts *bind.CallOpts, op LeafOpData, spec ProofSpecData) error {
	var out []interface{}
	err := _OpsUnitTest.contract.Call(opts, &out, "checkAgainstSpec", op, spec)

	if err != nil {
		return err
	}

	return err

}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0x5e7b22fa.
//
// Solidity: function checkAgainstSpec((uint8,uint8,uint8,uint8,bytes) op, ((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec) pure returns()
func (_OpsUnitTest *OpsUnitTestSession) CheckAgainstSpec(op LeafOpData, spec ProofSpecData) error {
	return _OpsUnitTest.Contract.CheckAgainstSpec(&_OpsUnitTest.CallOpts, op, spec)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0x5e7b22fa.
//
// Solidity: function checkAgainstSpec((uint8,uint8,uint8,uint8,bytes) op, ((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec) pure returns()
func (_OpsUnitTest *OpsUnitTestCallerSession) CheckAgainstSpec(op LeafOpData, spec ProofSpecData) error {
	return _OpsUnitTest.Contract.CheckAgainstSpec(&_OpsUnitTest.CallOpts, op, spec)
}

// CheckAgainstSpec0 is a free data retrieval call binding the contract method 0xf7717f42.
//
// Solidity: function checkAgainstSpec((uint8,bytes,bytes) op, ((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec) pure returns()
func (_OpsUnitTest *OpsUnitTestCaller) CheckAgainstSpec0(opts *bind.CallOpts, op InnerOpData, spec ProofSpecData) error {
	var out []interface{}
	err := _OpsUnitTest.contract.Call(opts, &out, "checkAgainstSpec0", op, spec)

	if err != nil {
		return err
	}

	return err

}

// CheckAgainstSpec0 is a free data retrieval call binding the contract method 0xf7717f42.
//
// Solidity: function checkAgainstSpec((uint8,bytes,bytes) op, ((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec) pure returns()
func (_OpsUnitTest *OpsUnitTestSession) CheckAgainstSpec0(op InnerOpData, spec ProofSpecData) error {
	return _OpsUnitTest.Contract.CheckAgainstSpec0(&_OpsUnitTest.CallOpts, op, spec)
}

// CheckAgainstSpec0 is a free data retrieval call binding the contract method 0xf7717f42.
//
// Solidity: function checkAgainstSpec((uint8,bytes,bytes) op, ((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec) pure returns()
func (_OpsUnitTest *OpsUnitTestCallerSession) CheckAgainstSpec0(op InnerOpData, spec ProofSpecData) error {
	return _OpsUnitTest.Contract.CheckAgainstSpec0(&_OpsUnitTest.CallOpts, op, spec)
}

// Compare is a free data retrieval call binding the contract method 0x81e79ac6.
//
// Solidity: function compare(bytes a, bytes b) pure returns(int256)
func (_OpsUnitTest *OpsUnitTestCaller) Compare(opts *bind.CallOpts, a []byte, b []byte) (*big.Int, error) {
	var out []interface{}
	err := _OpsUnitTest.contract.Call(opts, &out, "compare", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Compare is a free data retrieval call binding the contract method 0x81e79ac6.
//
// Solidity: function compare(bytes a, bytes b) pure returns(int256)
func (_OpsUnitTest *OpsUnitTestSession) Compare(a []byte, b []byte) (*big.Int, error) {
	return _OpsUnitTest.Contract.Compare(&_OpsUnitTest.CallOpts, a, b)
}

// Compare is a free data retrieval call binding the contract method 0x81e79ac6.
//
// Solidity: function compare(bytes a, bytes b) pure returns(int256)
func (_OpsUnitTest *OpsUnitTestCallerSession) Compare(a []byte, b []byte) (*big.Int, error) {
	return _OpsUnitTest.Contract.Compare(&_OpsUnitTest.CallOpts, a, b)
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 hashOp, bytes preImage) pure returns(bytes)
func (_OpsUnitTest *OpsUnitTestCaller) DoHash(opts *bind.CallOpts, hashOp uint8, preImage []byte) ([]byte, error) {
	var out []interface{}
	err := _OpsUnitTest.contract.Call(opts, &out, "doHash", hashOp, preImage)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 hashOp, bytes preImage) pure returns(bytes)
func (_OpsUnitTest *OpsUnitTestSession) DoHash(hashOp uint8, preImage []byte) ([]byte, error) {
	return _OpsUnitTest.Contract.DoHash(&_OpsUnitTest.CallOpts, hashOp, preImage)
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 hashOp, bytes preImage) pure returns(bytes)
func (_OpsUnitTest *OpsUnitTestCallerSession) DoHash(hashOp uint8, preImage []byte) ([]byte, error) {
	return _OpsUnitTest.Contract.DoHash(&_OpsUnitTest.CallOpts, hashOp, preImage)
}

// PROOFSPROTOGLOBALENUMSABI is the input ABI used to generate the binding from.
const PROOFSPROTOGLOBALENUMSABI = "[]"

// PROOFSPROTOGLOBALENUMSBin is the compiled bytecode used for deploying new contracts.
var PROOFSPROTOGLOBALENUMSBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220803567e438c2b32cc521747721a4b4cc79e9d6055019b5d6d80295ca96c08f5764736f6c63430008090033"

// DeployPROOFSPROTOGLOBALENUMS deploys a new Ethereum contract, binding an instance of PROOFSPROTOGLOBALENUMS to it.
func DeployPROOFSPROTOGLOBALENUMS(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PROOFSPROTOGLOBALENUMS, error) {
	parsed, err := abi.JSON(strings.NewReader(PROOFSPROTOGLOBALENUMSABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PROOFSPROTOGLOBALENUMSBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PROOFSPROTOGLOBALENUMS{PROOFSPROTOGLOBALENUMSCaller: PROOFSPROTOGLOBALENUMSCaller{contract: contract}, PROOFSPROTOGLOBALENUMSTransactor: PROOFSPROTOGLOBALENUMSTransactor{contract: contract}, PROOFSPROTOGLOBALENUMSFilterer: PROOFSPROTOGLOBALENUMSFilterer{contract: contract}}, nil
}

// PROOFSPROTOGLOBALENUMS is an auto generated Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMS struct {
	PROOFSPROTOGLOBALENUMSCaller     // Read-only binding to the contract
	PROOFSPROTOGLOBALENUMSTransactor // Write-only binding to the contract
	PROOFSPROTOGLOBALENUMSFilterer   // Log filterer for contract events
}

// PROOFSPROTOGLOBALENUMSCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROOFSPROTOGLOBALENUMSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROOFSPROTOGLOBALENUMSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROOFSPROTOGLOBALENUMSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROOFSPROTOGLOBALENUMSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROOFSPROTOGLOBALENUMSSession struct {
	Contract     *PROOFSPROTOGLOBALENUMS // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PROOFSPROTOGLOBALENUMSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROOFSPROTOGLOBALENUMSCallerSession struct {
	Contract *PROOFSPROTOGLOBALENUMSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// PROOFSPROTOGLOBALENUMSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROOFSPROTOGLOBALENUMSTransactorSession struct {
	Contract     *PROOFSPROTOGLOBALENUMSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// PROOFSPROTOGLOBALENUMSRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSRaw struct {
	Contract *PROOFSPROTOGLOBALENUMS // Generic contract binding to access the raw methods on
}

// PROOFSPROTOGLOBALENUMSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSCallerRaw struct {
	Contract *PROOFSPROTOGLOBALENUMSCaller // Generic read-only contract binding to access the raw methods on
}

// PROOFSPROTOGLOBALENUMSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROOFSPROTOGLOBALENUMSTransactorRaw struct {
	Contract *PROOFSPROTOGLOBALENUMSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROOFSPROTOGLOBALENUMS creates a new instance of PROOFSPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewPROOFSPROTOGLOBALENUMS(address common.Address, backend bind.ContractBackend) (*PROOFSPROTOGLOBALENUMS, error) {
	contract, err := bindPROOFSPROTOGLOBALENUMS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROOFSPROTOGLOBALENUMS{PROOFSPROTOGLOBALENUMSCaller: PROOFSPROTOGLOBALENUMSCaller{contract: contract}, PROOFSPROTOGLOBALENUMSTransactor: PROOFSPROTOGLOBALENUMSTransactor{contract: contract}, PROOFSPROTOGLOBALENUMSFilterer: PROOFSPROTOGLOBALENUMSFilterer{contract: contract}}, nil
}

// NewPROOFSPROTOGLOBALENUMSCaller creates a new read-only instance of PROOFSPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewPROOFSPROTOGLOBALENUMSCaller(address common.Address, caller bind.ContractCaller) (*PROOFSPROTOGLOBALENUMSCaller, error) {
	contract, err := bindPROOFSPROTOGLOBALENUMS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROOFSPROTOGLOBALENUMSCaller{contract: contract}, nil
}

// NewPROOFSPROTOGLOBALENUMSTransactor creates a new write-only instance of PROOFSPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewPROOFSPROTOGLOBALENUMSTransactor(address common.Address, transactor bind.ContractTransactor) (*PROOFSPROTOGLOBALENUMSTransactor, error) {
	contract, err := bindPROOFSPROTOGLOBALENUMS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROOFSPROTOGLOBALENUMSTransactor{contract: contract}, nil
}

// NewPROOFSPROTOGLOBALENUMSFilterer creates a new log filterer instance of PROOFSPROTOGLOBALENUMS, bound to a specific deployed contract.
func NewPROOFSPROTOGLOBALENUMSFilterer(address common.Address, filterer bind.ContractFilterer) (*PROOFSPROTOGLOBALENUMSFilterer, error) {
	contract, err := bindPROOFSPROTOGLOBALENUMS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROOFSPROTOGLOBALENUMSFilterer{contract: contract}, nil
}

// bindPROOFSPROTOGLOBALENUMS binds a generic wrapper to an already deployed contract.
func bindPROOFSPROTOGLOBALENUMS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PROOFSPROTOGLOBALENUMSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROOFSPROTOGLOBALENUMS.Contract.PROOFSPROTOGLOBALENUMSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROOFSPROTOGLOBALENUMS.Contract.PROOFSPROTOGLOBALENUMSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROOFSPROTOGLOBALENUMS.Contract.PROOFSPROTOGLOBALENUMSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROOFSPROTOGLOBALENUMS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROOFSPROTOGLOBALENUMS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROOFSPROTOGLOBALENUMS *PROOFSPROTOGLOBALENUMSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROOFSPROTOGLOBALENUMS.Contract.contract.Transact(opts, method, params...)
}

// ProofABI is the input ABI used to generate the binding from.
const ProofABI = "[]"

// ProofBin is the compiled bytecode used for deploying new contracts.
var ProofBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122082351082fff40436b618ee49b9260c10a2877162fd7805173023b74b1728561064736f6c63430008090033"

// DeployProof deploys a new Ethereum contract, binding an instance of Proof to it.
func DeployProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Proof, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proof{ProofCaller: ProofCaller{contract: contract}, ProofTransactor: ProofTransactor{contract: contract}, ProofFilterer: ProofFilterer{contract: contract}}, nil
}

// Proof is an auto generated Go binding around an Ethereum contract.
type Proof struct {
	ProofCaller     // Read-only binding to the contract
	ProofTransactor // Write-only binding to the contract
	ProofFilterer   // Log filterer for contract events
}

// ProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofSession struct {
	Contract     *Proof            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofCallerSession struct {
	Contract *ProofCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofTransactorSession struct {
	Contract     *ProofTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofRaw struct {
	Contract *Proof // Generic contract binding to access the raw methods on
}

// ProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofCallerRaw struct {
	Contract *ProofCaller // Generic read-only contract binding to access the raw methods on
}

// ProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofTransactorRaw struct {
	Contract *ProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProof creates a new instance of Proof, bound to a specific deployed contract.
func NewProof(address common.Address, backend bind.ContractBackend) (*Proof, error) {
	contract, err := bindProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proof{ProofCaller: ProofCaller{contract: contract}, ProofTransactor: ProofTransactor{contract: contract}, ProofFilterer: ProofFilterer{contract: contract}}, nil
}

// NewProofCaller creates a new read-only instance of Proof, bound to a specific deployed contract.
func NewProofCaller(address common.Address, caller bind.ContractCaller) (*ProofCaller, error) {
	contract, err := bindProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofCaller{contract: contract}, nil
}

// NewProofTransactor creates a new write-only instance of Proof, bound to a specific deployed contract.
func NewProofTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofTransactor, error) {
	contract, err := bindProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofTransactor{contract: contract}, nil
}

// NewProofFilterer creates a new log filterer instance of Proof, bound to a specific deployed contract.
func NewProofFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofFilterer, error) {
	contract, err := bindProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofFilterer{contract: contract}, nil
}

// bindProof binds a generic wrapper to an already deployed contract.
func bindProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proof *ProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proof.Contract.ProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proof *ProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.Contract.ProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proof *ProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proof.Contract.ProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proof *ProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proof *ProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proof *ProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proof.Contract.contract.Transact(opts, method, params...)
}

// ProofSpecABI is the input ABI used to generate the binding from.
const ProofSpecABI = "[]"

// ProofSpecBin is the compiled bytecode used for deploying new contracts.
var ProofSpecBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205305f069d3f93c035b449ac93d44ac0b2e4784d742ce3a5ffdeca4ea8c6d420a64736f6c63430008090033"

// DeployProofSpec deploys a new Ethereum contract, binding an instance of ProofSpec to it.
func DeployProofSpec(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProofSpec, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofSpecABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProofSpecBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProofSpec{ProofSpecCaller: ProofSpecCaller{contract: contract}, ProofSpecTransactor: ProofSpecTransactor{contract: contract}, ProofSpecFilterer: ProofSpecFilterer{contract: contract}}, nil
}

// ProofSpec is an auto generated Go binding around an Ethereum contract.
type ProofSpec struct {
	ProofSpecCaller     // Read-only binding to the contract
	ProofSpecTransactor // Write-only binding to the contract
	ProofSpecFilterer   // Log filterer for contract events
}

// ProofSpecCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofSpecCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSpecTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofSpecTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSpecFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofSpecFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSpecSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofSpecSession struct {
	Contract     *ProofSpec        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofSpecCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofSpecCallerSession struct {
	Contract *ProofSpecCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ProofSpecTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofSpecTransactorSession struct {
	Contract     *ProofSpecTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ProofSpecRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofSpecRaw struct {
	Contract *ProofSpec // Generic contract binding to access the raw methods on
}

// ProofSpecCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofSpecCallerRaw struct {
	Contract *ProofSpecCaller // Generic read-only contract binding to access the raw methods on
}

// ProofSpecTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofSpecTransactorRaw struct {
	Contract *ProofSpecTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProofSpec creates a new instance of ProofSpec, bound to a specific deployed contract.
func NewProofSpec(address common.Address, backend bind.ContractBackend) (*ProofSpec, error) {
	contract, err := bindProofSpec(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProofSpec{ProofSpecCaller: ProofSpecCaller{contract: contract}, ProofSpecTransactor: ProofSpecTransactor{contract: contract}, ProofSpecFilterer: ProofSpecFilterer{contract: contract}}, nil
}

// NewProofSpecCaller creates a new read-only instance of ProofSpec, bound to a specific deployed contract.
func NewProofSpecCaller(address common.Address, caller bind.ContractCaller) (*ProofSpecCaller, error) {
	contract, err := bindProofSpec(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofSpecCaller{contract: contract}, nil
}

// NewProofSpecTransactor creates a new write-only instance of ProofSpec, bound to a specific deployed contract.
func NewProofSpecTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofSpecTransactor, error) {
	contract, err := bindProofSpec(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofSpecTransactor{contract: contract}, nil
}

// NewProofSpecFilterer creates a new log filterer instance of ProofSpec, bound to a specific deployed contract.
func NewProofSpecFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofSpecFilterer, error) {
	contract, err := bindProofSpec(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofSpecFilterer{contract: contract}, nil
}

// bindProofSpec binds a generic wrapper to an already deployed contract.
func bindProofSpec(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofSpecABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofSpec *ProofSpecRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofSpec.Contract.ProofSpecCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofSpec *ProofSpecRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofSpec.Contract.ProofSpecTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofSpec *ProofSpecRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofSpec.Contract.ProofSpecTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofSpec *ProofSpecCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofSpec.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofSpec *ProofSpecTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofSpec.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofSpec *ProofSpecTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofSpec.Contract.contract.Transact(opts, method, params...)
}

// ProofUnitTestABI is the input ABI used to generate the binding from.
const ProofUnitTestABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structBatchProof.Data\",\"name\":\"batch\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"exist\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"internalType\":\"int32[]\",\"name\":\"path\",\"type\":\"int32[]\"}],\"internalType\":\"structCompressedExistenceProof.Data\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedNonExistenceProof.Data\",\"name\":\"nonexist\",\"type\":\"tuple\"}],\"internalType\":\"structCompressedBatchEntry.Data[]\",\"name\":\"entries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"lookup_inners\",\"type\":\"tuple[]\"}],\"internalType\":\"structCompressedBatchProof.Data\",\"name\":\"compressed\",\"type\":\"tuple\"}],\"internalType\":\"structCommitmentProof.Data\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structInnerOp.Data[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structExistenceProof.Data\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.LengthOp\",\"name\":\"length\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structLeafOp.Data\",\"name\":\"leaf_spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int32[]\",\"name\":\"child_order\",\"type\":\"int32[]\"},{\"internalType\":\"int32\",\"name\":\"child_size\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"max_prefix_length\",\"type\":\"int32\"},{\"internalType\":\"bytes\",\"name\":\"empty_child\",\"type\":\"bytes\"},{\"internalType\":\"enumPROOFS_PROTO_GLOBAL_ENUMS.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structInnerSpec.Data\",\"name\":\"inner_spec\",\"type\":\"tuple\"},{\"internalType\":\"int32\",\"name\":\"max_depth\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"min_depth\",\"type\":\"int32\"}],\"internalType\":\"structProofSpec.Data\",\"name\":\"spec\",\"type\":\"tuple\"}],\"name\":\"checkAgainstSpec\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProofUnitTestBin is the compiled bytecode used for deploying new contracts.
var ProofUnitTestBin = "0x608060405234801561001057600080fd5b506128ff806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063afc2128d14610046578063b5b7deb71461006f578063e6af408e14610082575b600080fd5b610059610054366004611fb5565b610097565b6040516100669190612019565b60405180910390f35b61005961007d3660046124ce565b6100a8565b6100956100903660046125a6565b6100b3565b005b60606100a2826100c1565b92915050565b60606100a282610162565b6100bd82826103b6565b5050565b60606100cf82604001511590565b156100f55760405162461bcd60e51b81526004016100ec90612713565b60405180910390fd5b600061010e836040015184600001518560200151610532565b905060005b83606001515181101561015b576101478460600151828151811061013957610139612757565b602002602001015183610633565b91508061015381612783565b915050610113565b5092915050565b606061017182600001516106cc565b6101805781516100a2906100c1565b61018d8260200151610705565b61019e576100a28260200151610718565b6101ab8260400151610705565b61033957604082015151516102025760405162461bcd60e51b815260206004820152601860248201527f62617463682070726f6f6620686173206e6f20656e747279000000000000000060448201526064016100ec565b61022882604001516000015160008151811061022057610220612757565b506001919050565b156102755760405162461bcd60e51b815260206004820152601b60248201527f62617463682070726f6f662068617320656d70747920656e747279000000000060448201526064016100ec565b6102a482604001516000015160008151811061029357610293612757565b6020026020010151600001516106cc565b6102d7576100a28260400151600001516000815181106102c6576102c6612757565b6020026020010151600001516100c1565b6103068260400151600001516000815181106102f5576102f5612757565b602002602001015160200151610705565b610339576100a282604001516000015160008151811061032857610328612757565b602002602001015160200151610718565b61034682606001516107b8565b61035b576100a2610356836107de565b610162565b60405162461bcd60e51b815260206004820152602a60248201527f63616c63756c617465526f6f7428436f6d6d69746d656e7450726f6f6629206560448201526936b83a3c90383937b7b360b11b60648201526084016100ec565b6103c38260400151610855565b156103e05760405162461bcd60e51b81526004016100ec90612713565b6103ee8260400151826108f4565b6000816060015160030b1315610469576000610410826060015160030b610b87565b83606001515110159050806104675760405162461bcd60e51b815260206004820152601860248201527f496e6e65724f707320646570746820746f6f2073686f7274000000000000000060448201526064016100ec565b505b6000816040015160030b13156104e457600061048b826040015160030b610b87565b83606001515111159050806104e25760405162461bcd60e51b815260206004820152601760248201527f496e6e65724f707320646570746820746f6f206c6f6e6700000000000000000060448201526064016100ec565b505b60005b82606001515181101561052d5761051b8360600151828151811061050d5761050d612757565b602002602001015183610bdd565b8061052581612783565b9150506104e7565b505050565b606060008351116105795760405162461bcd60e51b81526020600482015260116024820152704c656166206f70206e65656473206b657960781b60448201526064016100ec565b60008251116105c05760405162461bcd60e51b81526020600482015260136024820152724c656166206f70206e656564732076616c756560681b60448201526064016100ec565b60006105d58560200151866060015186610e0e565b905060006105ec8660400151876060015186610e0e565b90506000866080015183836040516020016106099392919061279e565b6040516020818303038152906040529050610628876000015182610e31565b979650505050505050565b606060008251116106865760405162461bcd60e51b815260206004820152601a60248201527f496e6e6572206f70206e65656473206368696c642076616c756500000000000060448201526064016100ec565b600083602001518385604001516040516020016106a59392919061279e565b60405160208183030381529060405290506106c4846000015182610e31565b949350505050565b805151600090156106df57506000919050565b602082015151156106f257506000919050565b6060820151511561022057506000919050565b8051516000901561022057506000919050565b606061072782602001516106cc565b610738576100a282602001516100c1565b61074582604001516106cc565b610756576100a282604001516100c1565b60405162461bcd60e51b815260206004820152603160248201527f4e6f6e6578697374656e63652070726f6f662068617320656d707479204c65666044820152703a1030b732102934b3b43a10383937b7b360791b60648201526084016100ec565b805151600090156107cb57506000919050565b6020820151511561022057506000919050565b6107e6611ab2565b6107f382606001516107b8565b151560011415610801575090565b6040518060800160405280610814611156565b8152602001610821611164565b8152602001604051806020016040528061083e866060015161116c565b9052815260200161084d611231565b905292915050565b8051600090600681111561086b5761086b6127e1565b1561087857506000919050565b8160200151600681111561088e5761088e6127e1565b1561089b57506000919050565b816040015160068111156108b1576108b16127e1565b156108be57506000919050565b816060015160088111156108d4576108d46127e1565b156108e157506000919050565b6080820151511561022057506000919050565b8051516006811115610908576109086127e1565b8251600681111561091b5761091b6127e1565b1461096e5760405162461bcd60e51b815260206004820152602f60248201526000805160206128aa83398151915260448201526e0657870656374656420486173684f7608c1b60648201526084016100ec565b8051602001516006811115610985576109856127e1565b8260200151600681111561099b5761099b6127e1565b146109f25760405162461bcd60e51b815260206004820152603360248201526000805160206128aa833981519152604482015272657870656374656420507265686173684b657960681b60648201526084016100ec565b8051604001516006811115610a0957610a096127e1565b82604001516006811115610a1f57610a1f6127e1565b14610a785760405162461bcd60e51b815260206004820152603560248201526000805160206128aa8339815191526044820152746578706563746564205072656861736856616c756560581b60648201526084016100ec565b8051606001516008811115610a8f57610a8f6127e1565b82606001516008811115610aa557610aa56127e1565b14610afa5760405162461bcd60e51b815260206004820152603160248201526000805160206128aa83398151915260448201527006578706563746564206c656e6774684f7607c1b60648201526084016100ec565b6000610b12836080015183600001516080015161124a565b90508061052d5760405162461bcd60e51b815260206004820152603c60248201527f636865636b416761696e73745370656320666f72204c6561664f70202d204c6560448201527f61662050726566697820646f65736e277420737461727420776974680000000060648201526084016100ec565b600080821215610bd95760405162461bcd60e51b815260206004820181905260248201527f53616665436173743a2076616c7565206d75737420626520706f73697469766560448201526064016100ec565b5090565b806020015160a001516006811115610bf757610bf76127e1565b82516006811115610c0a57610c0a6127e1565b14610c705760405162461bcd60e51b815260206004820152603060248201527f636865636b416761696e73745370656320666f7220496e6e65724f70202d205560448201526f06e657870656374656420486173684f760841b60648201526084016100ec565b6000610c8682602001516040015160030b610b87565b9050808360200151511015610cdd5760405162461bcd60e51b815260206004820152601860248201527f496e6e65724f702070726566697820746f6f2073686f7274000000000000000060448201526064016100ec565b8151608001516020840151600090610cf5908361124a565b90508015610d515760405162461bcd60e51b8152602060048201526024808201527f496e6e6572205072656669782073746172747320776974682077726f6e672076604482015263616c756560e01b60648201526084016100ec565b6000610d6785602001516020015160030b610b87565b9050600081600187602001516000015151610d8291906127f7565b610d8c919061280e565b90506000610da487602001516060015160030b610b87565b9050610db0828261282d565b8860200151511115610e045760405162461bcd60e51b815260206004820152601760248201527f496e6e65724f702070726566697820746f6f206c6f6e6700000000000000000060448201526064016100ec565b5050505050505050565b60606000610e1c858461128b565b9050610e2884826112bf565b95945050505050565b60606001836006811115610e4757610e476127e1565b1415610ec557600282604051610e5d9190612845565b602060405180830381855afa158015610e7a573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610e9d9190612861565b604051602001610eaf91815260200190565b60405160208183030381529060405290506100a2565b6003836006811115610ed957610ed96127e1565b1415610ef9578180519060200120604051602001610eaf91815260200190565b6004836006811115610f0d57610f0d6127e1565b1415610f6857600382604051610f239190612845565b602060405180830381855afa158015610f40573d6000803e3d6000fd5b5050604051805160601b6bffffffffffffffffffffffff191660208201526034019050610eaf565b6005836006811115610f7c57610f7c6127e1565b141561105e576000600283604051610f949190612845565b602060405180830381855afa158015610fb1573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610fd49190612861565b604051602001610fe691815260200190565b60405160208183030381529060405290506003816040516110079190612845565b602060405180830381855afa158015611024573d6000803e3d6000fd5b5050604051805160601b6bffffffffffffffffffffffff1916602082015260340190506040516020818303038152906040529150506100a2565b6002836006811115611072576110726127e1565b14156110b75760405162461bcd60e51b815260206004820152601460248201527314d2104d4c4c881b9bdd081cdd5c1c1bdc9d195960621b60448201526064016100ec565b60068360068111156110cb576110cb6127e1565b14156111195760405162461bcd60e51b815260206004820152601860248201527f5348413531325f323536206e6f7420737570706f72746564000000000000000060448201526064016100ec565b60405162461bcd60e51b81526020600482015260126024820152710556e737570706f7274656420686173684f760741b60448201526064016100ec565b61115e611b14565b50600090565b61115e611b42565b606060008260000151516001600160401b0381111561118d5761118d611bb1565b6040519080825280602002602001820160405280156111c657816020015b6111b3611b69565b8152602001906001900390816111ab5790505b50905060005b83515181101561015b57611201846000015182815181106111ef576111ef612757565b602002602001015185602001516115cb565b82828151811061121357611213612757565b6020026020010181905250808061122990612783565b9150506111cc565b604080518082019091526060808252602082015261115e565b600081516000141561125e575060016100a2565b825182511115611270575060006100a2565b600061127f8460008551611676565b90506106c48382611783565b606060008360068111156112a1576112a16127e1565b14156112ae5750806100a2565b6112b88383610e31565b9392505050565b606060008360088111156112d5576112d56127e1565b14156112e25750806100a2565b60018360088111156112f6576112f66127e1565b141561138a57600061130883516117e7565b90506000816001600160401b0381111561132457611324611bb1565b6040519080825280601f01601f19166020018201604052801561134e576020820181803683370190505b50905061135e8451602083611804565b50808460405160200161137292919061287a565b604051602081830303815290604052925050506100a2565b600783600881111561139e5761139e6127e1565b14156113f05781516020146113e95760405162461bcd60e51b81526020600482015260116024820152703230ba30973632b733ba3410109e90199960791b60448201526064016100ec565b50806100a2565b6008836008811115611404576114046127e1565b141561144f5781516040146113e95760405162461bcd60e51b815260206004820152601160248201527019185d184b9b195b99dd1a08084f480d8d607a1b60448201526064016100ec565b6004836008811115611463576114636127e1565b141561158f5760006114758351611847565b60408051600480825281830190925291925060e083901b916000916020820181803683370190505090508160031a60f81b816000815181106114b9576114b9612757565b60200101906001600160f81b031916908160001a9053508160021a60f81b816001815181106114ea576114ea612757565b60200101906001600160f81b031916908160001a9053508160011a60f81b8160028151811061151b5761151b612757565b60200101906001600160f81b031916908160001a9053508160001a60f81b8160038151811061154c5761154c612757565b60200101906001600160f81b031916908160001a905350808560405160200161157692919061287a565b60405160208183030381529060405293505050506100a2565b60405162461bcd60e51b81526020600482015260116024820152700556e737570706f72746564206c656e4f7607c1b60448201526064016100ec565b6115d3611b69565b82516115de906106cc565b6116115760405180604001604052806115fb8560000151856118ac565b8152602001611608611164565b905290506100a2565b6040518060400160405280611624611156565b815260200160405180606001604052808660200151600001518152602001611654876020015160200151876118ac565b815260200161166b876020015160400151876118ac565b905290529392505050565b60608161168481601f61282d565b10156116c35760405162461bcd60e51b815260206004820152600e60248201526d736c6963655f6f766572666c6f7760901b60448201526064016100ec565b6116cd828461282d565b845110156117115760405162461bcd60e51b8152602060048201526011602482015270736c6963655f6f75744f66426f756e647360781b60448201526064016100ec565b606082158015611730576040519150600082526020820160405261177a565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015611769578051835260209283019201611751565b5050858452601f01601f1916604052505b50949350505050565b81518151600091600191811480831461179f57600092506117dd565b600160208701838101602088015b6002848385100114156117d85780518351146117cc5760009650600093505b602092830192016117ad565b505050505b5090949350505050565b60071c600060015b82156100a25760079290921c916001016117ef565b600080828401607f86165b600787901c15611837578060801782535060079590951c9460019182019101607f861661180f565b8082535050600101949350505050565b600063ffffffff821115610bd95760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201526532206269747360d01b60648201526084016100ec565b6118b4611b14565b6118bd836106cc565b156118d1576118ca611156565b90506100a2565b600060405180608001604052808560000151815260200185602001518152602001856040015181526020018560600151516001600160401b0381111561191957611919611bb1565b60405190808252806020026020018201604052801561196f57816020015b61195c6040805160608101909152806000815260200160608152602001606081525090565b8152602001906001900390816119375790505b509052905060005b846060015151811015611aaa5760008560600151828151811061199c5761199c612757565b602002602001015160030b12156119e65760405162461bcd60e51b815260206004820152600e60248201526d070726f6f662e70617468203c20360941b60448201526064016100ec565b6000611a1186606001518381518110611a0157611a01612757565b602002602001015160030b610b87565b905084518110611a5b5760405162461bcd60e51b81526020600482015260156024820152740e6e8cae0407c7a40d8deded6eae05cd8cadccee8d605b1b60448201526064016100ec565b848181518110611a6d57611a6d612757565b602002602001015183606001518381518110611a8b57611a8b612757565b6020026020010181905250508080611aa290612783565b915050611977565b509392505050565b6040518060800160405280611ac5611b14565b8152602001611ad2611b42565b8152602001611aed6040518060200160405280606081525090565b8152602001611b0f604051806040016040528060608152602001606081525090565b905290565b60405180608001604052806060815260200160608152602001611b35611b89565b8152602001606081525090565b604051806060016040528060608152602001611b5c611b14565b8152602001611b0f611b14565b6040518060400160405280611b7c611b14565b8152602001611b0f611b42565b6040805160a08101909152806000815260200160008152602001600081526020016000611b35565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715611be957611be9611bb1565b60405290565b604051608081016001600160401b0381118282101715611be957611be9611bb1565b604051602081016001600160401b0381118282101715611be957611be9611bb1565b604080519081016001600160401b0381118282101715611be957611be9611bb1565b60405160c081016001600160401b0381118282101715611be957611be9611bb1565b604051601f8201601f191681016001600160401b0381118282101715611c9f57611c9f611bb1565b604052919050565b600082601f830112611cb857600080fd5b81356001600160401b03811115611cd157611cd1611bb1565b611ce4601f8201601f1916602001611c77565b818152846020838601011115611cf957600080fd5b816020850160208301376000918101602001919091529392505050565b803560078110611d2557600080fd5b919050565b600060a08284031215611d3c57600080fd5b60405160a081016001600160401b038282108183111715611d5f57611d5f611bb1565b81604052829350611d6f85611d16565b8352611d7d60208601611d16565b6020840152611d8e60408601611d16565b60408401526060850135915060098210611da757600080fd5b8160608401526080850135915080821115611dc157600080fd5b50611dce85828601611ca7565b6080830152505092915050565b60006001600160401b03821115611df457611df4611bb1565b5060051b60200190565b600082601f830112611e0f57600080fd5b81356020611e24611e1f83611ddb565b611c77565b82815260059290921b84018101918181019086841115611e4357600080fd5b8286015b84811015611ef25780356001600160401b0380821115611e675760008081fd5b908801906060828b03601f1901811315611e815760008081fd5b611e89611bc7565b611e94888501611d16565b815260408085013584811115611eaa5760008081fd5b611eb88e8b83890101611ca7565b838b015250918401359183831115611ed05760008081fd5b611ede8d8a85880101611ca7565b908201528652505050918301918301611e47565b509695505050505050565b600060808284031215611f0f57600080fd5b611f17611bef565b905081356001600160401b0380821115611f3057600080fd5b611f3c85838601611ca7565b83526020840135915080821115611f5257600080fd5b611f5e85838601611ca7565b60208401526040840135915080821115611f7757600080fd5b611f8385838601611d2a565b60408401526060840135915080821115611f9c57600080fd5b50611fa984828501611dfe565b60608301525092915050565b600060208284031215611fc757600080fd5b81356001600160401b03811115611fdd57600080fd5b6106c484828501611efd565b60005b83811015612004578181015183820152602001611fec565b83811115612013576000848401525b50505050565b6020815260008251806020840152612038816040850160208701611fe9565b601f01601f19169190910160400192915050565b60006060828403121561205e57600080fd5b612066611bc7565b905081356001600160401b038082111561207f57600080fd5b61208b85838601611ca7565b835260208401359150808211156120a157600080fd5b6120ad85838601611efd565b602084015260408401359150808211156120c657600080fd5b506120d384828501611efd565b60408301525092915050565b600060208083850312156120f257600080fd5b6120fa611c11565b915082356001600160401b038082111561211357600080fd5b818501915085601f83011261212757600080fd5b8135612135611e1f82611ddb565b81815260059190911b8301840190848101908883111561215457600080fd5b8585015b838110156121e7578035858111156121705760008081fd5b86016040818c03601f19018113156121885760008081fd5b612190611c33565b89830135888111156121a25760008081fd5b6121b08e8c83870101611efd565b8252509082013590878211156121c65760008081fd5b6121d48d8b8486010161204c565b818b015285525050918601918601612158565b50865250939695505050505050565b8035600381900b8114611d2557600080fd5b600082601f83011261221957600080fd5b81356020612229611e1f83611ddb565b82815260059290921b8401810191818101908684111561224857600080fd5b8286015b84811015611ef25761225d816121f6565b835291830191830161224c565b60006080828403121561227c57600080fd5b612284611bef565b905081356001600160401b038082111561229d57600080fd5b6122a985838601611ca7565b835260208401359150808211156122bf57600080fd5b6122cb85838601611ca7565b602084015260408401359150808211156122e457600080fd5b6122f085838601611d2a565b6040840152606084013591508082111561230957600080fd5b50611fa984828501612208565b60006040828403121561232857600080fd5b612330611c33565b905081356001600160401b038082111561234957600080fd5b818401915084601f83011261235d57600080fd5b8135602061236d611e1f83611ddb565b82815260059290921b8401810191818101908884111561238c57600080fd5b8286015b8481101561249d578035868111156123a757600080fd5b8701601f196040828d03820112156123be57600080fd5b6123c6611c33565b86830135898111156123d757600080fd5b6123e58e898387010161226a565b8252506040830135898111156123fa57600080fd5b92909201916060838e038301121561241157600080fd5b612419611bc7565b9150868301358981111561242c57600080fd5b61243a8e8983870101611ca7565b83525060408301358981111561244f57600080fd5b61245d8e898387010161226a565b888401525060608301358981111561247457600080fd5b6124828e898387010161226a565b60408401525080870191909152845250918301918301612390565b50865250858101359350828411156124b457600080fd5b6124c087858801611dfe565b818601525050505092915050565b6000602082840312156124e057600080fd5b81356001600160401b03808211156124f757600080fd5b908301906080828603121561250b57600080fd5b612513611bef565b82358281111561252257600080fd5b61252e87828601611efd565b82525060208301358281111561254357600080fd5b61254f8782860161204c565b60208301525060408301358281111561256757600080fd5b612573878286016120df565b60408301525060608301358281111561258b57600080fd5b61259787828601612316565b60608301525095945050505050565b600080604083850312156125b957600080fd5b82356001600160401b03808211156125d057600080fd5b6125dc86838701611efd565b935060208501359150808211156125f257600080fd5b908401906080828703121561260657600080fd5b61260e611bef565b82358281111561261d57600080fd5b61262988828601611d2a565b82525060208301358281111561263e57600080fd5b830160c0818903121561265057600080fd5b612658611c55565b81358481111561266757600080fd5b6126738a828501612208565b825250612682602083016121f6565b6020820152612693604083016121f6565b60408201526126a4606083016121f6565b60608201526080820135848111156126bb57600080fd5b6126c78a828501611ca7565b6080830152506126d960a08301611d16565b60a08201526020830152506126f0604084016121f6565b6040820152612701606084016121f6565b60608201528093505050509250929050565b60208082526024908201527f4578697374656e63652050726f6f66206e6565647320646566696e6564204c65604082015263061664f760e41b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006000198214156127975761279761276d565b5060010190565b600084516127b0818460208901611fe9565b8451908301906127c4818360208901611fe9565b84519101906127d7818360208801611fe9565b0195945050505050565b634e487b7160e01b600052602160045260246000fd5b6000828210156128095761280961276d565b500390565b60008160001904831182151516156128285761282861276d565b500290565b600082198211156128405761284061276d565b500190565b60008251612857818460208701611fe9565b9190910192915050565b60006020828403121561287357600080fd5b5051919050565b6000835161288c818460208801611fe9565b8351908301906128a0818360208801611fe9565b0194935050505056fe636865636b416761696e73745370656320666f72204c6561664f70202d20556ea2646970667358221220a2b1b5976fd9c7dee6e39668707da8f712afd8020fa03a74e9cd9cc340e130ee64736f6c63430008090033"

// DeployProofUnitTest deploys a new Ethereum contract, binding an instance of ProofUnitTest to it.
func DeployProofUnitTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProofUnitTest, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofUnitTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProofUnitTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProofUnitTest{ProofUnitTestCaller: ProofUnitTestCaller{contract: contract}, ProofUnitTestTransactor: ProofUnitTestTransactor{contract: contract}, ProofUnitTestFilterer: ProofUnitTestFilterer{contract: contract}}, nil
}

// ProofUnitTest is an auto generated Go binding around an Ethereum contract.
type ProofUnitTest struct {
	ProofUnitTestCaller     // Read-only binding to the contract
	ProofUnitTestTransactor // Write-only binding to the contract
	ProofUnitTestFilterer   // Log filterer for contract events
}

// ProofUnitTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofUnitTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofUnitTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofUnitTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofUnitTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofUnitTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofUnitTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofUnitTestSession struct {
	Contract     *ProofUnitTest    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofUnitTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofUnitTestCallerSession struct {
	Contract *ProofUnitTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProofUnitTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofUnitTestTransactorSession struct {
	Contract     *ProofUnitTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProofUnitTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofUnitTestRaw struct {
	Contract *ProofUnitTest // Generic contract binding to access the raw methods on
}

// ProofUnitTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofUnitTestCallerRaw struct {
	Contract *ProofUnitTestCaller // Generic read-only contract binding to access the raw methods on
}

// ProofUnitTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofUnitTestTransactorRaw struct {
	Contract *ProofUnitTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProofUnitTest creates a new instance of ProofUnitTest, bound to a specific deployed contract.
func NewProofUnitTest(address common.Address, backend bind.ContractBackend) (*ProofUnitTest, error) {
	contract, err := bindProofUnitTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProofUnitTest{ProofUnitTestCaller: ProofUnitTestCaller{contract: contract}, ProofUnitTestTransactor: ProofUnitTestTransactor{contract: contract}, ProofUnitTestFilterer: ProofUnitTestFilterer{contract: contract}}, nil
}

// NewProofUnitTestCaller creates a new read-only instance of ProofUnitTest, bound to a specific deployed contract.
func NewProofUnitTestCaller(address common.Address, caller bind.ContractCaller) (*ProofUnitTestCaller, error) {
	contract, err := bindProofUnitTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofUnitTestCaller{contract: contract}, nil
}

// NewProofUnitTestTransactor creates a new write-only instance of ProofUnitTest, bound to a specific deployed contract.
func NewProofUnitTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofUnitTestTransactor, error) {
	contract, err := bindProofUnitTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofUnitTestTransactor{contract: contract}, nil
}

// NewProofUnitTestFilterer creates a new log filterer instance of ProofUnitTest, bound to a specific deployed contract.
func NewProofUnitTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofUnitTestFilterer, error) {
	contract, err := bindProofUnitTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofUnitTestFilterer{contract: contract}, nil
}

// bindProofUnitTest binds a generic wrapper to an already deployed contract.
func bindProofUnitTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofUnitTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofUnitTest *ProofUnitTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofUnitTest.Contract.ProofUnitTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofUnitTest *ProofUnitTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofUnitTest.Contract.ProofUnitTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofUnitTest *ProofUnitTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofUnitTest.Contract.ProofUnitTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProofUnitTest *ProofUnitTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProofUnitTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProofUnitTest *ProofUnitTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProofUnitTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProofUnitTest *ProofUnitTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProofUnitTest.Contract.contract.Transact(opts, method, params...)
}

// CalculateRoot is a free data retrieval call binding the contract method 0xafc2128d.
//
// Solidity: function calculateRoot((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof) pure returns(bytes)
func (_ProofUnitTest *ProofUnitTestCaller) CalculateRoot(opts *bind.CallOpts, proof ExistenceProofData) ([]byte, error) {
	var out []interface{}
	err := _ProofUnitTest.contract.Call(opts, &out, "calculateRoot", proof)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0xafc2128d.
//
// Solidity: function calculateRoot((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof) pure returns(bytes)
func (_ProofUnitTest *ProofUnitTestSession) CalculateRoot(proof ExistenceProofData) ([]byte, error) {
	return _ProofUnitTest.Contract.CalculateRoot(&_ProofUnitTest.CallOpts, proof)
}

// CalculateRoot is a free data retrieval call binding the contract method 0xafc2128d.
//
// Solidity: function calculateRoot((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof) pure returns(bytes)
func (_ProofUnitTest *ProofUnitTestCallerSession) CalculateRoot(proof ExistenceProofData) ([]byte, error) {
	return _ProofUnitTest.Contract.CalculateRoot(&_ProofUnitTest.CallOpts, proof)
}

// CalculateRoot0 is a free data retrieval call binding the contract method 0xb5b7deb7.
//
// Solidity: function calculateRoot(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof) pure returns(bytes)
func (_ProofUnitTest *ProofUnitTestCaller) CalculateRoot0(opts *bind.CallOpts, proof CommitmentProofData) ([]byte, error) {
	var out []interface{}
	err := _ProofUnitTest.contract.Call(opts, &out, "calculateRoot0", proof)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CalculateRoot0 is a free data retrieval call binding the contract method 0xb5b7deb7.
//
// Solidity: function calculateRoot(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof) pure returns(bytes)
func (_ProofUnitTest *ProofUnitTestSession) CalculateRoot0(proof CommitmentProofData) ([]byte, error) {
	return _ProofUnitTest.Contract.CalculateRoot0(&_ProofUnitTest.CallOpts, proof)
}

// CalculateRoot0 is a free data retrieval call binding the contract method 0xb5b7deb7.
//
// Solidity: function calculateRoot(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[])))[]),(((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[]),(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),int32[])))[],(uint8,bytes,bytes)[])) proof) pure returns(bytes)
func (_ProofUnitTest *ProofUnitTestCallerSession) CalculateRoot0(proof CommitmentProofData) ([]byte, error) {
	return _ProofUnitTest.Contract.CalculateRoot0(&_ProofUnitTest.CallOpts, proof)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xe6af408e.
//
// Solidity: function checkAgainstSpec((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, ((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec) pure returns()
func (_ProofUnitTest *ProofUnitTestCaller) CheckAgainstSpec(opts *bind.CallOpts, proof ExistenceProofData, spec ProofSpecData) error {
	var out []interface{}
	err := _ProofUnitTest.contract.Call(opts, &out, "checkAgainstSpec", proof, spec)

	if err != nil {
		return err
	}

	return err

}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xe6af408e.
//
// Solidity: function checkAgainstSpec((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, ((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec) pure returns()
func (_ProofUnitTest *ProofUnitTestSession) CheckAgainstSpec(proof ExistenceProofData, spec ProofSpecData) error {
	return _ProofUnitTest.Contract.CheckAgainstSpec(&_ProofUnitTest.CallOpts, proof, spec)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xe6af408e.
//
// Solidity: function checkAgainstSpec((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, ((uint8,uint8,uint8,uint8,bytes),(int32[],int32,int32,int32,bytes,uint8),int32,int32) spec) pure returns()
func (_ProofUnitTest *ProofUnitTestCallerSession) CheckAgainstSpec(proof ExistenceProofData, spec ProofSpecData) error {
	return _ProofUnitTest.Contract.CheckAgainstSpec(&_ProofUnitTest.CallOpts, proof, spec)
}

// ProtoBufRuntimeABI is the input ABI used to generate the binding from.
const ProtoBufRuntimeABI = "[]"

// ProtoBufRuntimeBin is the compiled bytecode used for deploying new contracts.
var ProtoBufRuntimeBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209b3e53136cc68e7f54670a15cca16642f2601d010cbcfd3dd9ca5ae521df731564736f6c63430008090033"

// DeployProtoBufRuntime deploys a new Ethereum contract, binding an instance of ProtoBufRuntime to it.
func DeployProtoBufRuntime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProtoBufRuntime, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtoBufRuntimeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProtoBufRuntimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProtoBufRuntime{ProtoBufRuntimeCaller: ProtoBufRuntimeCaller{contract: contract}, ProtoBufRuntimeTransactor: ProtoBufRuntimeTransactor{contract: contract}, ProtoBufRuntimeFilterer: ProtoBufRuntimeFilterer{contract: contract}}, nil
}

// ProtoBufRuntime is an auto generated Go binding around an Ethereum contract.
type ProtoBufRuntime struct {
	ProtoBufRuntimeCaller     // Read-only binding to the contract
	ProtoBufRuntimeTransactor // Write-only binding to the contract
	ProtoBufRuntimeFilterer   // Log filterer for contract events
}

// ProtoBufRuntimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtoBufRuntimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoBufRuntimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtoBufRuntimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoBufRuntimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtoBufRuntimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoBufRuntimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtoBufRuntimeSession struct {
	Contract     *ProtoBufRuntime  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtoBufRuntimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtoBufRuntimeCallerSession struct {
	Contract *ProtoBufRuntimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProtoBufRuntimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtoBufRuntimeTransactorSession struct {
	Contract     *ProtoBufRuntimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProtoBufRuntimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtoBufRuntimeRaw struct {
	Contract *ProtoBufRuntime // Generic contract binding to access the raw methods on
}

// ProtoBufRuntimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtoBufRuntimeCallerRaw struct {
	Contract *ProtoBufRuntimeCaller // Generic read-only contract binding to access the raw methods on
}

// ProtoBufRuntimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtoBufRuntimeTransactorRaw struct {
	Contract *ProtoBufRuntimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtoBufRuntime creates a new instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntime(address common.Address, backend bind.ContractBackend) (*ProtoBufRuntime, error) {
	contract, err := bindProtoBufRuntime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntime{ProtoBufRuntimeCaller: ProtoBufRuntimeCaller{contract: contract}, ProtoBufRuntimeTransactor: ProtoBufRuntimeTransactor{contract: contract}, ProtoBufRuntimeFilterer: ProtoBufRuntimeFilterer{contract: contract}}, nil
}

// NewProtoBufRuntimeCaller creates a new read-only instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntimeCaller(address common.Address, caller bind.ContractCaller) (*ProtoBufRuntimeCaller, error) {
	contract, err := bindProtoBufRuntime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntimeCaller{contract: contract}, nil
}

// NewProtoBufRuntimeTransactor creates a new write-only instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntimeTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtoBufRuntimeTransactor, error) {
	contract, err := bindProtoBufRuntime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntimeTransactor{contract: contract}, nil
}

// NewProtoBufRuntimeFilterer creates a new log filterer instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntimeFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtoBufRuntimeFilterer, error) {
	contract, err := bindProtoBufRuntime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntimeFilterer{contract: contract}, nil
}

// bindProtoBufRuntime binds a generic wrapper to an already deployed contract.
func bindProtoBufRuntime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtoBufRuntimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoBufRuntime *ProtoBufRuntimeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoBufRuntime.Contract.ProtoBufRuntimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoBufRuntime *ProtoBufRuntimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.ProtoBufRuntimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoBufRuntime *ProtoBufRuntimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.ProtoBufRuntimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoBufRuntime *ProtoBufRuntimeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoBufRuntime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoBufRuntime *ProtoBufRuntimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoBufRuntime *ProtoBufRuntimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.contract.Transact(opts, method, params...)
}

// SafeCastABI is the input ABI used to generate the binding from.
const SafeCastABI = "[]"

// SafeCastBin is the compiled bytecode used for deploying new contracts.
var SafeCastBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201b4679b5db9dcf8aa501c15c1f6cd465d2e48fab3466c77d230222952754346764736f6c63430008090033"

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeCastABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeCastABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}
