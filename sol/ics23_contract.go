// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ics23

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

// ICS23ExistenceProof is an auto generated low-level Go binding around an user-defined struct.
type ICS23ExistenceProof struct {
	Key   []byte
	Value []byte
	Leaf  ICS23LeafOp
	Path  []ICS23InnerOp
}

// ICS23InnerOp is an auto generated low-level Go binding around an user-defined struct.
type ICS23InnerOp struct {
	Hash   uint8
	Prefix []byte
	Suffix []byte
}

// ICS23LeafOp is an auto generated low-level Go binding around an user-defined struct.
type ICS23LeafOp struct {
	Hash         uint8
	PrehashKey   uint8
	PrehashValue uint8
	Len          uint8
	Prefix       []byte
}

// ICS23ABI is the input ABI used to generate the binding from.
const ICS23ABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"child\",\"type\":\"bytes\"}],\"name\":\"applyInner\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"applyLeaf\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"spec\",\"type\":\"tuple\"}],\"name\":\"checkAgainstSpec\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHashOrNoop\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"doLength\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bz1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bz2\",\"type\":\"bytes\"}],\"name\":\"equalBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"s\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"hasprefix\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hashop\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"lengthop\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"prepareLeafData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyExistence\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ICS23FuncSigs maps the 4-byte function signature to its string representation.
var ICS23FuncSigs = map[string]string{
	"3d4a397d": "applyInner((uint8,bytes,bytes),bytes)",
	"f6747d82": "applyLeaf((uint8,uint8,uint8,uint8,bytes),bytes,bytes)",
	"7e1fd3bc": "calculate((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]))",
	"f10e9a9c": "checkAgainstSpec((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(uint8,uint8,uint8,uint8,bytes))",
	"d48f1e4f": "doHash(uint8,bytes)",
	"03801174": "doHashOrNoop(uint8,bytes)",
	"67bb8e81": "doLength(uint8,bytes)",
	"4cac70ff": "equalBytes(bytes,bytes)",
	"901d0e15": "hasprefix(bytes,bytes)",
	"fd29e20a": "prepareLeafData(uint8,uint8,bytes)",
	"2e3098a9": "verifyExistence((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),(uint8,uint8,uint8,uint8,bytes),bytes,bytes,bytes)",
	"a04f1008": "verifyMembership((uint8,uint8,uint8,uint8,bytes),bytes,(bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]),bytes,bytes)",
}

// ICS23Bin is the compiled bytecode used for deploying new contracts.
var ICS23Bin = "0x60016080818152600060a081905260c083905260e083905261016060405261012083815261014082815261010091909152815460ff1916841762ffff001916620100001763ff0000001916630100000017825591929091620000639190816200007a565b5050503480156200007357600080fd5b506200015d565b828054620000889062000120565b90600052602060002090601f016020900481019282620000ac5760008555620000f7565b82601f10620000c757805160ff1916838001178555620000f7565b82800160010185558215620000f7579182015b82811115620000f7578251825591602001919060010190620000da565b506200010592915062000109565b5090565b5b808211156200010557600081556001016200010a565b6002810460018216806200013557607f821691505b602082108114156200015757634e487b7160e01b600052602260045260246000fd5b50919050565b61132e806200016d6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063901d0e1511610071578063901d0e1514610151578063a04f100814610164578063d48f1e4f14610177578063f10e9a9c1461018a578063f6747d821461019d578063fd29e20a146101b0576100b4565b806303801174146100b95780632e3098a9146100e25780633d4a397d146101055780634cac70ff1461011857806367bb8e811461012b5780637e1fd3bc1461013e575b600080fd5b6100cc6100c7366004610e2b565b6101c3565b6040516100d99190611228565b60405180910390f35b6100f56100f0366004610f6a565b610207565b60405190151581526020016100d9565b6100cc610113366004611035565b610260565b6100f5610126366004610dcb565b6102b7565b6100cc610139366004610ec7565b610365565b6100cc61014c366004610ee2565b6104e4565b6100f561015f366004610dcb565b610563565b6100f56101723660046110e2565b6105f5565b6100cc610185366004610e2b565b610604565b6100f5610198366004610f14565b610861565b6100cc6101ab366004611069565b610a41565b6100cc6101be366004610e6c565b610ad0565b606060008360058111156101e757634e487b7160e01b600052602160045260246000fd5b14156101f4575080610201565b6101fe8383610604565b90505b92915050565b60006102138686610861565b801561022857506102288387600001516102b7565b801561023d575061023d8287602001516102b7565b80156102565750610256610250876104e4565b856102b7565b9695505050505050565b606081516000141561027157600080fd5b6000836020015183856040015160405160200161029093929190611179565b60405160208183030381529060405290506102af846000015182610604565b949350505050565b600081518351146102ca57506000610201565b60005b835181101561035b578281815181106102f657634e487b7160e01b600052603260045260246000fd5b602001015160f81c60f81b6001600160f81b03191684828151811061032b57634e487b7160e01b600052603260045260246000fd5b01602001516001600160f81b03191614610349576000915050610201565b80610353816112bb565b9150506102cd565b5060019392505050565b6060600083600881111561038957634e487b7160e01b600052602160045260246000fd5b1415610396575080610201565b60018360088111156103b857634e487b7160e01b600052602160045260246000fd5b1415610414578151608081106104015780607f16608017600782901c915081846040516020016103ea939291906111eb565b604051602081830303815290604052915050610201565b80836040516020016103ea9291906111bc565b600783600881111561043657634e487b7160e01b600052602160045260246000fd5b141561045157815160201461044a57600080fd5b5080610201565b600883600881111561047357634e487b7160e01b600052602160045260246000fd5b141561048757815160401461044a57600080fd5b60405162461bcd60e51b815260206004820152602760248201527f696e76616c6964206f7220756e737570706f72746564206c656e677468206f7060448201526632b930ba34b7b760c91b60648201526084015b60405180910390fd5b606060006104ff836040015184600001518560200151610a41565b905060005b83606001515181101561055a576105468460600151828151811061053857634e487b7160e01b600052603260045260246000fd5b602002602001015183610260565b915080610552816112bb565b915050610504565b5090505b919050565b6000805b825181101561035b5782818151811061059057634e487b7160e01b600052603260045260246000fd5b602001015160f81c60f81b6001600160f81b0319168482815181106105c557634e487b7160e01b600052603260045260246000fd5b01602001516001600160f81b031916146105e3576000915050610201565b806105ed816112bb565b915050610567565b60006102568487878686610207565b6060600383600581111561062857634e487b7160e01b600052602160045260246000fd5b141561065e57818051906020012060405160200161064891815260200190565b6040516020818303038152906040529050610201565b600183600581111561068057634e487b7160e01b600052602160045260246000fd5b14156106e857600282604051610696919061115d565b602060405180830381855afa1580156106b3573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906106d69190610db3565b60405160200161064891815260200190565b600483600581111561070a57634e487b7160e01b600052602160045260246000fd5b141561076557600382604051610720919061115d565b602060405180830381855afa15801561073d573d6000803e3d6000fd5b5050604051805160601b6bffffffffffffffffffffffff191660208201526034019050610648565b600583600581111561078757634e487b7160e01b600052602160045260246000fd5b141561080b57600360028360405161079f919061115d565b602060405180830381855afa1580156107bc573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906107df9190610db3565b6040516020016107f191815260200190565b60408051601f19818403018152908290526107209161115d565b60405162461bcd60e51b815260206004820152602560248201527f696e76616c6964206f7220756e737570706f727465642068617368206f70657260448201526430ba34b7b760d91b60648201526084016104db565b8051600090600581111561088557634e487b7160e01b600052602160045260246000fd5b60408401515160058111156108aa57634e487b7160e01b600052602160045260246000fd5b1480156109005750816020015160058111156108d657634e487b7160e01b600052602160045260246000fd5b83604001516020015160058111156108fe57634e487b7160e01b600052602160045260246000fd5b145b801561095557508160400151600581111561092b57634e487b7160e01b600052602160045260246000fd5b836040015160400151600581111561095357634e487b7160e01b600052602160045260246000fd5b145b80156109aa57508160600151600881111561098057634e487b7160e01b600052602160045260246000fd5b83604001516060015160088111156109a857634e487b7160e01b600052602160045260246000fd5b145b80156109c757506109c78360400151608001518360800151610563565b6109d357506000610201565b60005b83606001515181101561035b57610a2084606001518281518110610a0a57634e487b7160e01b600052603260045260246000fd5b6020026020010151602001518460800151610563565b15610a2f576000915050610201565b80610a39816112bb565b9150506109d6565b6060825160001415610a5257600080fd5b8151610a5d57600080fd5b6000610a728560200151866060015186610ad0565b90506000610a898660400151876060015186610ad0565b9050600086608001518383604051602001610aa693929190611179565b6040516020818303038152906040529050610ac5876000015182610604565b979650505050505050565b60606000610ade85846101c3565b905060006102568583610365565b600082601f830112610afc578081fd5b813560206001600160401b03821115610b1757610b176112e2565b610b24818284020161125b565b82815281810190858301855b85811015610b5957610b47898684358b0101610ca7565b84529284019290840190600101610b30565b5090979650505050505050565b600082601f830112610b76578081fd5b81356001600160401b03811115610b8f57610b8f6112e2565b610ba2601f8201601f191660200161125b565b818152846020838601011115610bb6578283fd5b816020850160208301379081016020019190915292915050565b80356006811061055e57600080fd5b80356009811061055e57600080fd5b600060808284031215610bff578081fd5b610c09608061125b565b905081356001600160401b0380821115610c2257600080fd5b610c2e85838601610b66565b83526020840135915080821115610c4457600080fd5b610c5085838601610b66565b60208401526040840135915080821115610c6957600080fd5b610c7585838601610d27565b60408401526060840135915080821115610c8e57600080fd5b50610c9b84828501610aec565b60608301525092915050565b600060608284031215610cb8578081fd5b610cc2606061125b565b9050610ccd82610bd0565b815260208201356001600160401b0380821115610ce957600080fd5b610cf585838601610b66565b60208401526040840135915080821115610d0e57600080fd5b50610d1b84828501610b66565b60408301525092915050565b600060a08284031215610d38578081fd5b610d4260a061125b565b9050610d4d82610bd0565b8152610d5b60208301610bd0565b6020820152610d6c60408301610bd0565b6040820152610d7d60608301610bdf565b606082015260808201356001600160401b03811115610d9b57600080fd5b610da784828501610b66565b60808301525092915050565b600060208284031215610dc4578081fd5b5051919050565b60008060408385031215610ddd578081fd5b82356001600160401b0380821115610df3578283fd5b610dff86838701610b66565b93506020850135915080821115610e14578283fd5b50610e2185828601610b66565b9150509250929050565b60008060408385031215610e3d578182fd5b610e4683610bd0565b915060208301356001600160401b03811115610e60578182fd5b610e2185828601610b66565b600080600060608486031215610e80578081fd5b610e8984610bd0565b9250610e9760208501610bdf565b915060408401356001600160401b03811115610eb1578182fd5b610ebd86828701610b66565b9150509250925092565b60008060408385031215610ed9578182fd5b610e4683610bdf565b600060208284031215610ef3578081fd5b81356001600160401b03811115610f08578182fd5b6102af84828501610bee565b60008060408385031215610f26578182fd5b82356001600160401b0380821115610f3c578384fd5b610f4886838701610bee565b93506020850135915080821115610f5d578283fd5b50610e2185828601610d27565b600080600080600060a08688031215610f81578081fd5b85356001600160401b0380821115610f97578283fd5b610fa389838a01610bee565b96506020880135915080821115610fb8578283fd5b610fc489838a01610d27565b95506040880135915080821115610fd9578283fd5b610fe589838a01610b66565b94506060880135915080821115610ffa578283fd5b61100689838a01610b66565b9350608088013591508082111561101b578283fd5b5061102888828901610b66565b9150509295509295909350565b60008060408385031215611047578182fd5b82356001600160401b038082111561105d578384fd5b610dff86838701610ca7565b60008060006060848603121561107d578283fd5b83356001600160401b0380821115611093578485fd5b61109f87838801610d27565b945060208601359150808211156110b4578384fd5b6110c087838801610b66565b935060408601359150808211156110d5578283fd5b50610ebd86828701610b66565b600080600080600060a086880312156110f9578283fd5b85356001600160401b038082111561110f578485fd5b61111b89838a01610d27565b96506020880135915080821115611130578485fd5b61113c89838a01610b66565b95506040880135915080821115611151578485fd5b610fe589838a01610bee565b6000825161116f81846020870161128b565b9190910192915050565b6000845161118b81846020890161128b565b84519083019061119f81836020890161128b565b84519101906111b281836020880161128b565b0195945050505050565b600060ff60f81b8460f81b16825282516111dd81600185016020870161128b565b919091016001019392505050565b600060ff60f81b808660f81b168352808560f81b16600184015250825161121981600285016020870161128b565b91909101600201949350505050565b600060208252825180602084015261124781604085016020870161128b565b601f01601f19169190910160400192915050565b604051601f8201601f191681016001600160401b0381118282101715611283576112836112e2565b604052919050565b60005b838110156112a657818101518382015260200161128e565b838111156112b5576000848401525b50505050565b60006000198214156112db57634e487b7160e01b81526011600452602481fd5b5060010190565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220bef6fd28c4a83952ebfa43369489377a0393358642b37d0a4eb2434c06a3d90964736f6c63430008020033"

// DeployICS23 deploys a new Ethereum contract, binding an instance of ICS23 to it.
func DeployICS23(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ICS23, error) {
	parsed, err := abi.JSON(strings.NewReader(ICS23ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ICS23Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ICS23{ICS23Caller: ICS23Caller{contract: contract}, ICS23Transactor: ICS23Transactor{contract: contract}, ICS23Filterer: ICS23Filterer{contract: contract}}, nil
}

// ICS23 is an auto generated Go binding around an Ethereum contract.
type ICS23 struct {
	ICS23Caller     // Read-only binding to the contract
	ICS23Transactor // Write-only binding to the contract
	ICS23Filterer   // Log filterer for contract events
}

// ICS23Caller is an auto generated read-only Go binding around an Ethereum contract.
type ICS23Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS23Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ICS23Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS23Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICS23Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICS23Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICS23Session struct {
	Contract     *ICS23            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICS23CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICS23CallerSession struct {
	Contract *ICS23Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ICS23TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICS23TransactorSession struct {
	Contract     *ICS23Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICS23Raw is an auto generated low-level Go binding around an Ethereum contract.
type ICS23Raw struct {
	Contract *ICS23 // Generic contract binding to access the raw methods on
}

// ICS23CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICS23CallerRaw struct {
	Contract *ICS23Caller // Generic read-only contract binding to access the raw methods on
}

// ICS23TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICS23TransactorRaw struct {
	Contract *ICS23Transactor // Generic write-only contract binding to access the raw methods on
}

// NewICS23 creates a new instance of ICS23, bound to a specific deployed contract.
func NewICS23(address common.Address, backend bind.ContractBackend) (*ICS23, error) {
	contract, err := bindICS23(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICS23{ICS23Caller: ICS23Caller{contract: contract}, ICS23Transactor: ICS23Transactor{contract: contract}, ICS23Filterer: ICS23Filterer{contract: contract}}, nil
}

// NewICS23Caller creates a new read-only instance of ICS23, bound to a specific deployed contract.
func NewICS23Caller(address common.Address, caller bind.ContractCaller) (*ICS23Caller, error) {
	contract, err := bindICS23(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICS23Caller{contract: contract}, nil
}

// NewICS23Transactor creates a new write-only instance of ICS23, bound to a specific deployed contract.
func NewICS23Transactor(address common.Address, transactor bind.ContractTransactor) (*ICS23Transactor, error) {
	contract, err := bindICS23(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICS23Transactor{contract: contract}, nil
}

// NewICS23Filterer creates a new log filterer instance of ICS23, bound to a specific deployed contract.
func NewICS23Filterer(address common.Address, filterer bind.ContractFilterer) (*ICS23Filterer, error) {
	contract, err := bindICS23(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICS23Filterer{contract: contract}, nil
}

// bindICS23 binds a generic wrapper to an already deployed contract.
func bindICS23(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ICS23ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS23 *ICS23Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS23.Contract.ICS23Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS23 *ICS23Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS23.Contract.ICS23Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS23 *ICS23Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS23.Contract.ICS23Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICS23 *ICS23CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICS23.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICS23 *ICS23TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICS23.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICS23 *ICS23TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICS23.Contract.contract.Transact(opts, method, params...)
}

// ApplyInner is a free data retrieval call binding the contract method 0x3d4a397d.
//
// Solidity: function applyInner((uint8,bytes,bytes) op, bytes child) pure returns(bytes)
func (_ICS23 *ICS23Caller) ApplyInner(opts *bind.CallOpts, op ICS23InnerOp, child []byte) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "applyInner", op, child)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ApplyInner is a free data retrieval call binding the contract method 0x3d4a397d.
//
// Solidity: function applyInner((uint8,bytes,bytes) op, bytes child) pure returns(bytes)
func (_ICS23 *ICS23Session) ApplyInner(op ICS23InnerOp, child []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyInner(&_ICS23.CallOpts, op, child)
}

// ApplyInner is a free data retrieval call binding the contract method 0x3d4a397d.
//
// Solidity: function applyInner((uint8,bytes,bytes) op, bytes child) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) ApplyInner(op ICS23InnerOp, child []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyInner(&_ICS23.CallOpts, op, child)
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xf6747d82.
//
// Solidity: function applyLeaf((uint8,uint8,uint8,uint8,bytes) op, bytes key, bytes value) pure returns(bytes)
func (_ICS23 *ICS23Caller) ApplyLeaf(opts *bind.CallOpts, op ICS23LeafOp, key []byte, value []byte) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "applyLeaf", op, key, value)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ApplyLeaf is a free data retrieval call binding the contract method 0xf6747d82.
//
// Solidity: function applyLeaf((uint8,uint8,uint8,uint8,bytes) op, bytes key, bytes value) pure returns(bytes)
func (_ICS23 *ICS23Session) ApplyLeaf(op ICS23LeafOp, key []byte, value []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyLeaf(&_ICS23.CallOpts, op, key, value)
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xf6747d82.
//
// Solidity: function applyLeaf((uint8,uint8,uint8,uint8,bytes) op, bytes key, bytes value) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) ApplyLeaf(op ICS23LeafOp, key []byte, value []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyLeaf(&_ICS23.CallOpts, op, key, value)
}

// Calculate is a free data retrieval call binding the contract method 0x7e1fd3bc.
//
// Solidity: function calculate((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof) pure returns(bytes)
func (_ICS23 *ICS23Caller) Calculate(opts *bind.CallOpts, proof ICS23ExistenceProof) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "calculate", proof)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Calculate is a free data retrieval call binding the contract method 0x7e1fd3bc.
//
// Solidity: function calculate((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof) pure returns(bytes)
func (_ICS23 *ICS23Session) Calculate(proof ICS23ExistenceProof) ([]byte, error) {
	return _ICS23.Contract.Calculate(&_ICS23.CallOpts, proof)
}

// Calculate is a free data retrieval call binding the contract method 0x7e1fd3bc.
//
// Solidity: function calculate((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) Calculate(proof ICS23ExistenceProof) ([]byte, error) {
	return _ICS23.Contract.Calculate(&_ICS23.CallOpts, proof)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xf10e9a9c.
//
// Solidity: function checkAgainstSpec((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, (uint8,uint8,uint8,uint8,bytes) spec) pure returns(bool)
func (_ICS23 *ICS23Caller) CheckAgainstSpec(opts *bind.CallOpts, proof ICS23ExistenceProof, spec ICS23LeafOp) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "checkAgainstSpec", proof, spec)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xf10e9a9c.
//
// Solidity: function checkAgainstSpec((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, (uint8,uint8,uint8,uint8,bytes) spec) pure returns(bool)
func (_ICS23 *ICS23Session) CheckAgainstSpec(proof ICS23ExistenceProof, spec ICS23LeafOp) (bool, error) {
	return _ICS23.Contract.CheckAgainstSpec(&_ICS23.CallOpts, proof, spec)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xf10e9a9c.
//
// Solidity: function checkAgainstSpec((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, (uint8,uint8,uint8,uint8,bytes) spec) pure returns(bool)
func (_ICS23 *ICS23CallerSession) CheckAgainstSpec(proof ICS23ExistenceProof, spec ICS23LeafOp) (bool, error) {
	return _ICS23.Contract.CheckAgainstSpec(&_ICS23.CallOpts, proof, spec)
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 op, bytes preimage) pure returns(bytes)
func (_ICS23 *ICS23Caller) DoHash(opts *bind.CallOpts, op uint8, preimage []byte) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "doHash", op, preimage)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 op, bytes preimage) pure returns(bytes)
func (_ICS23 *ICS23Session) DoHash(op uint8, preimage []byte) ([]byte, error) {
	return _ICS23.Contract.DoHash(&_ICS23.CallOpts, op, preimage)
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 op, bytes preimage) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) DoHash(op uint8, preimage []byte) ([]byte, error) {
	return _ICS23.Contract.DoHash(&_ICS23.CallOpts, op, preimage)
}

// DoHashOrNoop is a free data retrieval call binding the contract method 0x03801174.
//
// Solidity: function doHashOrNoop(uint8 op, bytes preimage) pure returns(bytes)
func (_ICS23 *ICS23Caller) DoHashOrNoop(opts *bind.CallOpts, op uint8, preimage []byte) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "doHashOrNoop", op, preimage)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DoHashOrNoop is a free data retrieval call binding the contract method 0x03801174.
//
// Solidity: function doHashOrNoop(uint8 op, bytes preimage) pure returns(bytes)
func (_ICS23 *ICS23Session) DoHashOrNoop(op uint8, preimage []byte) ([]byte, error) {
	return _ICS23.Contract.DoHashOrNoop(&_ICS23.CallOpts, op, preimage)
}

// DoHashOrNoop is a free data retrieval call binding the contract method 0x03801174.
//
// Solidity: function doHashOrNoop(uint8 op, bytes preimage) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) DoHashOrNoop(op uint8, preimage []byte) ([]byte, error) {
	return _ICS23.Contract.DoHashOrNoop(&_ICS23.CallOpts, op, preimage)
}

// DoLength is a free data retrieval call binding the contract method 0x67bb8e81.
//
// Solidity: function doLength(uint8 op, bytes data) pure returns(bytes)
func (_ICS23 *ICS23Caller) DoLength(opts *bind.CallOpts, op uint8, data []byte) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "doLength", op, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DoLength is a free data retrieval call binding the contract method 0x67bb8e81.
//
// Solidity: function doLength(uint8 op, bytes data) pure returns(bytes)
func (_ICS23 *ICS23Session) DoLength(op uint8, data []byte) ([]byte, error) {
	return _ICS23.Contract.DoLength(&_ICS23.CallOpts, op, data)
}

// DoLength is a free data retrieval call binding the contract method 0x67bb8e81.
//
// Solidity: function doLength(uint8 op, bytes data) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) DoLength(op uint8, data []byte) ([]byte, error) {
	return _ICS23.Contract.DoLength(&_ICS23.CallOpts, op, data)
}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes bz1, bytes bz2) pure returns(bool)
func (_ICS23 *ICS23Caller) EqualBytes(opts *bind.CallOpts, bz1 []byte, bz2 []byte) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "equalBytes", bz1, bz2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes bz1, bytes bz2) pure returns(bool)
func (_ICS23 *ICS23Session) EqualBytes(bz1 []byte, bz2 []byte) (bool, error) {
	return _ICS23.Contract.EqualBytes(&_ICS23.CallOpts, bz1, bz2)
}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes bz1, bytes bz2) pure returns(bool)
func (_ICS23 *ICS23CallerSession) EqualBytes(bz1 []byte, bz2 []byte) (bool, error) {
	return _ICS23.Contract.EqualBytes(&_ICS23.CallOpts, bz1, bz2)
}

// Hasprefix is a free data retrieval call binding the contract method 0x901d0e15.
//
// Solidity: function hasprefix(bytes s, bytes prefix) pure returns(bool)
func (_ICS23 *ICS23Caller) Hasprefix(opts *bind.CallOpts, s []byte, prefix []byte) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "hasprefix", s, prefix)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Hasprefix is a free data retrieval call binding the contract method 0x901d0e15.
//
// Solidity: function hasprefix(bytes s, bytes prefix) pure returns(bool)
func (_ICS23 *ICS23Session) Hasprefix(s []byte, prefix []byte) (bool, error) {
	return _ICS23.Contract.Hasprefix(&_ICS23.CallOpts, s, prefix)
}

// Hasprefix is a free data retrieval call binding the contract method 0x901d0e15.
//
// Solidity: function hasprefix(bytes s, bytes prefix) pure returns(bool)
func (_ICS23 *ICS23CallerSession) Hasprefix(s []byte, prefix []byte) (bool, error) {
	return _ICS23.Contract.Hasprefix(&_ICS23.CallOpts, s, prefix)
}

// PrepareLeafData is a free data retrieval call binding the contract method 0xfd29e20a.
//
// Solidity: function prepareLeafData(uint8 hashop, uint8 lengthop, bytes data) pure returns(bytes)
func (_ICS23 *ICS23Caller) PrepareLeafData(opts *bind.CallOpts, hashop uint8, lengthop uint8, data []byte) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "prepareLeafData", hashop, lengthop, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PrepareLeafData is a free data retrieval call binding the contract method 0xfd29e20a.
//
// Solidity: function prepareLeafData(uint8 hashop, uint8 lengthop, bytes data) pure returns(bytes)
func (_ICS23 *ICS23Session) PrepareLeafData(hashop uint8, lengthop uint8, data []byte) ([]byte, error) {
	return _ICS23.Contract.PrepareLeafData(&_ICS23.CallOpts, hashop, lengthop, data)
}

// PrepareLeafData is a free data retrieval call binding the contract method 0xfd29e20a.
//
// Solidity: function prepareLeafData(uint8 hashop, uint8 lengthop, bytes data) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) PrepareLeafData(hashop uint8, lengthop uint8, data []byte) ([]byte, error) {
	return _ICS23.Contract.PrepareLeafData(&_ICS23.CallOpts, hashop, lengthop, data)
}

// VerifyExistence is a free data retrieval call binding the contract method 0x2e3098a9.
//
// Solidity: function verifyExistence((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, (uint8,uint8,uint8,uint8,bytes) spec, bytes root, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23Caller) VerifyExistence(opts *bind.CallOpts, proof ICS23ExistenceProof, spec ICS23LeafOp, root []byte, key []byte, value []byte) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "verifyExistence", proof, spec, root, key, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyExistence is a free data retrieval call binding the contract method 0x2e3098a9.
//
// Solidity: function verifyExistence((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, (uint8,uint8,uint8,uint8,bytes) spec, bytes root, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23Session) VerifyExistence(proof ICS23ExistenceProof, spec ICS23LeafOp, root []byte, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyExistence(&_ICS23.CallOpts, proof, spec, root, key, value)
}

// VerifyExistence is a free data retrieval call binding the contract method 0x2e3098a9.
//
// Solidity: function verifyExistence((bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, (uint8,uint8,uint8,uint8,bytes) spec, bytes root, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23CallerSession) VerifyExistence(proof ICS23ExistenceProof, spec ICS23LeafOp, root []byte, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyExistence(&_ICS23.CallOpts, proof, spec, root, key, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xa04f1008.
//
// Solidity: function verifyMembership((uint8,uint8,uint8,uint8,bytes) spec, bytes root, (bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23Caller) VerifyMembership(opts *bind.CallOpts, spec ICS23LeafOp, root []byte, proof ICS23ExistenceProof, key []byte, value []byte) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "verifyMembership", spec, root, proof, key, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xa04f1008.
//
// Solidity: function verifyMembership((uint8,uint8,uint8,uint8,bytes) spec, bytes root, (bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23Session) VerifyMembership(spec ICS23LeafOp, root []byte, proof ICS23ExistenceProof, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyMembership(&_ICS23.CallOpts, spec, root, proof, key, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xa04f1008.
//
// Solidity: function verifyMembership((uint8,uint8,uint8,uint8,bytes) spec, bytes root, (bytes,bytes,(uint8,uint8,uint8,uint8,bytes),(uint8,bytes,bytes)[]) proof, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23CallerSession) VerifyMembership(spec ICS23LeafOp, root []byte, proof ICS23ExistenceProof, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyMembership(&_ICS23.CallOpts, spec, root, proof, key, value)
}
