// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proofs

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ProofsABI is the input ABI used to generate the binding from.
const ProofsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"op\",\"type\":\"uint8\"},{\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHashOrNoop\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"op\",\"type\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"doLength\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"s\",\"type\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"hasprefix\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"op\",\"type\":\"uint8\"},{\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashop\",\"type\":\"uint8\"},{\"name\":\"lengthop\",\"type\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"prepareLeafData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProofsFuncSigs maps the 4-byte function signature to its string representation.
var ProofsFuncSigs = map[string]string{
	"d48f1e4f": "doHash(uint8,bytes)",
	"03801174": "doHashOrNoop(uint8,bytes)",
	"67bb8e81": "doLength(uint8,bytes)",
	"901d0e15": "hasprefix(bytes,bytes)",
	"fd29e20a": "prepareLeafData(uint8,uint8,bytes)",
}

// ProofsBin is the compiled bytecode used for deploying new contracts.
var ProofsBin = "0x60016080818152600060a081905260c083905260e083905261016060405261012083815261014082815261010091909152815460ff1916841762ffff001916620100001763ff0000001916630100000017825591929091610061919081610076565b50505034801561007057600080fd5b50610111565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100b757805160ff19168380011785556100e4565b828001600101855582156100e4579182015b828111156100e45782518255916020019190600101906100c9565b506100f09291506100f4565b5090565b61010e91905b808211156100f057600081556001016100fa565b90565b610ac9806101206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063038011741461005c57806367bb8e811461017f578063901d0e151461022d578063d48f1e4f1461036a578063fd29e20a14610418575b600080fd5b61010a6004803603604081101561007257600080fd5b60ff8235169190810190604081016020820135600160201b81111561009657600080fd5b8201836020820111156100a857600080fd5b803590602001918460018302840111600160201b831117156100c957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104cf945050505050565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014457818101518382015260200161012c565b50505050905090810190601f1680156101715780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61010a6004803603604081101561019557600080fd5b60ff8235169190810190604081016020820135600160201b8111156101b957600080fd5b8201836020820111156101cb57600080fd5b803590602001918460018302840111600160201b831117156101ec57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104ff945050505050565b6103566004803603604081101561024357600080fd5b810190602081018135600160201b81111561025d57600080fd5b82018360208201111561026f57600080fd5b803590602001918460018302840111600160201b8311171561029057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156102e257600080fd5b8201836020820111156102f457600080fd5b803590602001918460018302840111600160201b8311171561031557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106da945050505050565b604080519115158252519081900360200190f35b61010a6004803603604081101561038057600080fd5b60ff8235169190810190604081016020820135600160201b8111156103a457600080fd5b8201836020820111156103b657600080fd5b803590602001918460018302840111600160201b831117156103d757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610744945050505050565b61010a6004803603606081101561042e57600080fd5b60ff8235811692602081013590911691810190606081016040820135600160201b81111561045b57600080fd5b82018360208201111561046d57600080fd5b803590602001918460018302840111600160201b8311171561048e57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a23945050505050565b606060008360058111156104df57fe5b14156104ec5750806104f9565b6104f68383610744565b90505b92915050565b6060600083600881111561050f57fe5b141561051c5750806104f9565b600183600881111561052a57fe5b1415610658578151608081106105d85780607f16608017600782901c91508184604051602001808460ff1660ff1660f81b81526001018360ff1660ff1660f81b815260010182805190602001908083835b6020831061059a5780518252601f19909201916020918201910161057b565b6001836020036101000a03801982511681845116808217855250505050505090500193505050506040516020818303038152906040529150506104f9565b8083604051602001808360ff1660ff1660f81b815260010182805190602001908083835b6020831061061b5780518252601f1990920191602091820191016105fc565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529150506104f9565b600783600881111561066657fe5b141561068157815160201461067a57600080fd5b50806104f9565b600883600881111561068f57fe5b14156106a357815160401461067a57600080fd5b60405162461bcd60e51b8152600401808060200182810382526027815260200180610a496027913960400191505060405180910390fd5b6000805b825181101561073a578281815181106106f357fe5b602001015160f81c60f81b6001600160f81b03191684828151811061071457fe5b01602001516001600160f81b031916146107325760009150506104f9565b6001016106de565b5060019392505050565b6060600383600581111561075457fe5b14156107885781805190602001206040516020018082815260200191505060405160208183030381529060405290506104f9565b600183600581111561079657fe5b1415610846576002826040518082805190602001908083835b602083106107ce5780518252601f1990920191602091820191016107af565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561080d573d6000803e3d6000fd5b5050506040513d602081101561082257600080fd5b505160408051602081810193909352815180820390930183528101905290506104f9565b600483600581111561085457fe5b1415610905576003826040518082805190602001908083835b6020831061088c5780518252601f19909201916020918201910161086d565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa1580156108cb573d6000803e3d6000fd5b505060408051805160601b6bffffffffffffffffffffffff1916602082015281516014818303018152603490910190915291506104f99050565b600583600581111561091357fe5b14156109ec5760036002836040518082805190602001908083835b6020831061094d5780518252601f19909201916020918201910161092e565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801561098c573d6000803e3d6000fd5b5050506040513d60208110156109a157600080fd5b50516040805160208181019390935281518082038401815290820191829052805190928291908401908083836020831061088c5780518252601f19909201916020918201910161086d565b60405162461bcd60e51b8152600401808060200182810382526025815260200180610a706025913960400191505060405180910390fd5b606080610a3085846104cf565b90506060610a3e85836104ff565b969550505050505056fe696e76616c6964206f7220756e737570706f72746564206c656e677468206f7065726174696f6e696e76616c6964206f7220756e737570706f727465642068617368206f7065726174696f6ea265627a7a72305820c33c19623256e13565ca6123c410ad5b2fc0206db20620c463f4a62233e0cb7564736f6c63430005090032"

// DeployProofs deploys a new Ethereum contract, binding an instance of Proofs to it.
func DeployProofs(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Proofs, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProofsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proofs{ProofsCaller: ProofsCaller{contract: contract}, ProofsTransactor: ProofsTransactor{contract: contract}, ProofsFilterer: ProofsFilterer{contract: contract}}, nil
}

// Proofs is an auto generated Go binding around an Ethereum contract.
type Proofs struct {
	ProofsCaller     // Read-only binding to the contract
	ProofsTransactor // Write-only binding to the contract
	ProofsFilterer   // Log filterer for contract events
}

// ProofsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofsSession struct {
	Contract     *Proofs           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofsCallerSession struct {
	Contract *ProofsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProofsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofsTransactorSession struct {
	Contract     *ProofsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofsRaw struct {
	Contract *Proofs // Generic contract binding to access the raw methods on
}

// ProofsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofsCallerRaw struct {
	Contract *ProofsCaller // Generic read-only contract binding to access the raw methods on
}

// ProofsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofsTransactorRaw struct {
	Contract *ProofsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProofs creates a new instance of Proofs, bound to a specific deployed contract.
func NewProofs(address common.Address, backend bind.ContractBackend) (*Proofs, error) {
	contract, err := bindProofs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proofs{ProofsCaller: ProofsCaller{contract: contract}, ProofsTransactor: ProofsTransactor{contract: contract}, ProofsFilterer: ProofsFilterer{contract: contract}}, nil
}

// NewProofsCaller creates a new read-only instance of Proofs, bound to a specific deployed contract.
func NewProofsCaller(address common.Address, caller bind.ContractCaller) (*ProofsCaller, error) {
	contract, err := bindProofs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofsCaller{contract: contract}, nil
}

// NewProofsTransactor creates a new write-only instance of Proofs, bound to a specific deployed contract.
func NewProofsTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofsTransactor, error) {
	contract, err := bindProofs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofsTransactor{contract: contract}, nil
}

// NewProofsFilterer creates a new log filterer instance of Proofs, bound to a specific deployed contract.
func NewProofsFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofsFilterer, error) {
	contract, err := bindProofs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofsFilterer{contract: contract}, nil
}

// bindProofs binds a generic wrapper to an already deployed contract.
func bindProofs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proofs *ProofsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proofs.Contract.ProofsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proofs *ProofsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proofs.Contract.ProofsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proofs *ProofsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proofs.Contract.ProofsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proofs *ProofsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proofs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proofs *ProofsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proofs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proofs *ProofsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proofs.Contract.contract.Transact(opts, method, params...)
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 op, bytes preimage) constant returns(bytes)
func (_Proofs *ProofsCaller) DoHash(opts *bind.CallOpts, op uint8, preimage []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "doHash", op, preimage)
	return *ret0, err
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 op, bytes preimage) constant returns(bytes)
func (_Proofs *ProofsSession) DoHash(op uint8, preimage []byte) ([]byte, error) {
	return _Proofs.Contract.DoHash(&_Proofs.CallOpts, op, preimage)
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 op, bytes preimage) constant returns(bytes)
func (_Proofs *ProofsCallerSession) DoHash(op uint8, preimage []byte) ([]byte, error) {
	return _Proofs.Contract.DoHash(&_Proofs.CallOpts, op, preimage)
}

// DoHashOrNoop is a free data retrieval call binding the contract method 0x03801174.
//
// Solidity: function doHashOrNoop(uint8 op, bytes preimage) constant returns(bytes)
func (_Proofs *ProofsCaller) DoHashOrNoop(opts *bind.CallOpts, op uint8, preimage []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "doHashOrNoop", op, preimage)
	return *ret0, err
}

// DoHashOrNoop is a free data retrieval call binding the contract method 0x03801174.
//
// Solidity: function doHashOrNoop(uint8 op, bytes preimage) constant returns(bytes)
func (_Proofs *ProofsSession) DoHashOrNoop(op uint8, preimage []byte) ([]byte, error) {
	return _Proofs.Contract.DoHashOrNoop(&_Proofs.CallOpts, op, preimage)
}

// DoHashOrNoop is a free data retrieval call binding the contract method 0x03801174.
//
// Solidity: function doHashOrNoop(uint8 op, bytes preimage) constant returns(bytes)
func (_Proofs *ProofsCallerSession) DoHashOrNoop(op uint8, preimage []byte) ([]byte, error) {
	return _Proofs.Contract.DoHashOrNoop(&_Proofs.CallOpts, op, preimage)
}

// DoLength is a free data retrieval call binding the contract method 0x67bb8e81.
//
// Solidity: function doLength(uint8 op, bytes data) constant returns(bytes)
func (_Proofs *ProofsCaller) DoLength(opts *bind.CallOpts, op uint8, data []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "doLength", op, data)
	return *ret0, err
}

// DoLength is a free data retrieval call binding the contract method 0x67bb8e81.
//
// Solidity: function doLength(uint8 op, bytes data) constant returns(bytes)
func (_Proofs *ProofsSession) DoLength(op uint8, data []byte) ([]byte, error) {
	return _Proofs.Contract.DoLength(&_Proofs.CallOpts, op, data)
}

// DoLength is a free data retrieval call binding the contract method 0x67bb8e81.
//
// Solidity: function doLength(uint8 op, bytes data) constant returns(bytes)
func (_Proofs *ProofsCallerSession) DoLength(op uint8, data []byte) ([]byte, error) {
	return _Proofs.Contract.DoLength(&_Proofs.CallOpts, op, data)
}

// Hasprefix is a free data retrieval call binding the contract method 0x901d0e15.
//
// Solidity: function hasprefix(bytes s, bytes prefix) constant returns(bool)
func (_Proofs *ProofsCaller) Hasprefix(opts *bind.CallOpts, s []byte, prefix []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "hasprefix", s, prefix)
	return *ret0, err
}

// Hasprefix is a free data retrieval call binding the contract method 0x901d0e15.
//
// Solidity: function hasprefix(bytes s, bytes prefix) constant returns(bool)
func (_Proofs *ProofsSession) Hasprefix(s []byte, prefix []byte) (bool, error) {
	return _Proofs.Contract.Hasprefix(&_Proofs.CallOpts, s, prefix)
}

// Hasprefix is a free data retrieval call binding the contract method 0x901d0e15.
//
// Solidity: function hasprefix(bytes s, bytes prefix) constant returns(bool)
func (_Proofs *ProofsCallerSession) Hasprefix(s []byte, prefix []byte) (bool, error) {
	return _Proofs.Contract.Hasprefix(&_Proofs.CallOpts, s, prefix)
}

// PrepareLeafData is a free data retrieval call binding the contract method 0xfd29e20a.
//
// Solidity: function prepareLeafData(uint8 hashop, uint8 lengthop, bytes data) constant returns(bytes)
func (_Proofs *ProofsCaller) PrepareLeafData(opts *bind.CallOpts, hashop uint8, lengthop uint8, data []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "prepareLeafData", hashop, lengthop, data)
	return *ret0, err
}

// PrepareLeafData is a free data retrieval call binding the contract method 0xfd29e20a.
//
// Solidity: function prepareLeafData(uint8 hashop, uint8 lengthop, bytes data) constant returns(bytes)
func (_Proofs *ProofsSession) PrepareLeafData(hashop uint8, lengthop uint8, data []byte) ([]byte, error) {
	return _Proofs.Contract.PrepareLeafData(&_Proofs.CallOpts, hashop, lengthop, data)
}

// PrepareLeafData is a free data retrieval call binding the contract method 0xfd29e20a.
//
// Solidity: function prepareLeafData(uint8 hashop, uint8 lengthop, bytes data) constant returns(bytes)
func (_Proofs *ProofsCallerSession) PrepareLeafData(hashop uint8, lengthop uint8, data []byte) ([]byte, error) {
	return _Proofs.Contract.PrepareLeafData(&_Proofs.CallOpts, hashop, lengthop, data)
}
