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
const ProofsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"op\",\"type\":\"uint8\"},{\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHashOrNoop\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\"}],\"name\":\"path\",\"type\":\"tuple[]\"}],\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"spec\",\"type\":\"tuple\"},{\"name\":\"root\",\"type\":\"bytes\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyExistence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\"}],\"name\":\"op\",\"type\":\"tuple\"},{\"name\":\"child\",\"type\":\"bytes\"}],\"name\":\"applyInner\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"bz1\",\"type\":\"bytes\"},{\"name\":\"bz2\",\"type\":\"bytes\"}],\"name\":\"equalBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"op\",\"type\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"doLength\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\"}],\"name\":\"path\",\"type\":\"tuple[]\"}],\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"calculate\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"s\",\"type\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"hasprefix\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"op\",\"type\":\"uint8\"},{\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\"}],\"name\":\"path\",\"type\":\"tuple[]\"}],\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"spec\",\"type\":\"tuple\"}],\"name\":\"checkAgainstSpec\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"op\",\"type\":\"tuple\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"applyLeaf\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashop\",\"type\":\"uint8\"},{\"name\":\"lengthop\",\"type\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"prepareLeafData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ProofsFuncSigs maps the 4-byte function signature to its string representation.
var ProofsFuncSigs = map[string]string{
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
}

// ProofsBin is the compiled bytecode used for deploying new contracts.
var ProofsBin = "0x60016080818152600060a081905260c083905260e083905261016060405261012083815261014082815261010091909152815460ff1916841762ffff001916620100001763ff0000001916630100000017825591929091620000639190816200007a565b5050503480156200007357600080fd5b506200011f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000bd57805160ff1916838001178555620000ed565b82800160010185558215620000ed579182015b82811115620000ed578251825591602001919060010190620000d0565b50620000fb929150620000ff565b5090565b6200011c91905b80821115620000fb576000815560010162000106565b90565b6111bf806200012f6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80637e1fd3bc116100715780637e1fd3bc14610130578063901d0e1514610143578063d48f1e4f14610156578063f10e9a9c14610169578063f6747d821461017c578063fd29e20a1461018f576100a9565b806303801174146100ae5780632e3098a9146100d75780633d4a397d146100f75780634cac70ff1461010a57806367bb8e811461011d575b600080fd5b6100c16100bc366004610b82565b6101a2565b6040516100ce9190611024565b60405180910390f35b6100ea6100e5366004610cb5565b6101d2565b6040516100ce9190611016565b6100c1610105366004610d9c565b61022b565b6100ea610118366004610b1b565b610282565b6100c161012b366004610c05565b6102fe565b6100c161013e366004610c24565b6103f1565b6100ea610151366004610b1b565b61044f565b6100c1610164366004610b82565b6104af565b6100ea610177366004610c58565b610668565b6100c161018a366004610dd1565b61078a565b6100c161019d366004610ba1565b610819565b606060008360058111156101b257fe5b14156101bf5750806101cc565b6101c983836104af565b90505b92915050565b60006101de8686610668565b80156101f357506101f3838760000151610282565b80156102085750610208828760200151610282565b8015610221575061022161021b876103f1565b85610282565b9695505050505050565b606081516000141561023c57600080fd5b6060836020015183856040015160405160200161025b93929190610fa1565b604051602081830303815290604052905061027a8460000151826104af565b949350505050565b60008151835114610295575060006101cc565b60005b83518110156102f4578281815181106102ad57fe5b602001015160f81c60f81b6001600160f81b0319168482815181106102ce57fe5b01602001516001600160f81b031916146102ec5760009150506101cc565b600101610298565b5060019392505050565b6060600083600881111561030e57fe5b141561031b5750806101cc565b600183600881111561032957fe5b1415610385578151608081106103725780607f16608017600782901c9150818460405160200161035b93929190610fea565b6040516020818303038152906040529150506101cc565b808360405160200161035b929190610fce565b600783600881111561039357fe5b14156103ae5781516020146103a757600080fd5b50806101cc565b60088360088111156103bc57fe5b14156103d05781516040146103a757600080fd5b60405162461bcd60e51b81526004016103e890611035565b60405180910390fd5b60608061040b83604001518460000151856020015161078a565b905060005b8360600151518110156104485761043e8460600151828151811061043057fe5b60200260200101518361022b565b9150600101610410565b5092915050565b6000805b82518110156102f45782818151811061046857fe5b602001015160f81c60f81b6001600160f81b03191684828151811061048957fe5b01602001516001600160f81b031916146104a75760009150506101cc565b600101610453565b606060038360058111156104bf57fe5b14156104f35781805190602001206040516020016104dd9190610f79565b60405160208183030381529060405290506101cc565b600183600581111561050157fe5b1415610567576002826040516105179190610f8e565b602060405180830381855afa158015610534573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052506105579190810190610afd565b6040516020016104dd9190610f79565b600483600581111561057557fe5b14156105c05760038260405161058b9190610f8e565b602060405180830381855afa1580156105a8573d6000803e3d6000fd5b505060405180516104dd925060601b90602001610f64565b60058360058111156105ce57fe5b14156106505760036002836040516105e69190610f8e565b602060405180830381855afa158015610603573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052506106269190810190610afd565b6040516020016106369190610f79565b60408051601f198184030181529082905261058b91610f8e565b60405162461bcd60e51b81526004016103e890611045565b8051600090600581111561067857fe5b604084015151600581111561068957fe5b1480156106b75750816020015160058111156106a157fe5b83604001516020015160058111156106b557fe5b145b80156106e45750816040015160058111156106ce57fe5b83604001516040015160058111156106e257fe5b145b80156107115750816060015160088111156106fb57fe5b836040015160600151600881111561070f57fe5b145b801561072e575061072e836040015160800151836080015161044f565b61073a575060006101cc565b60005b8360600151518110156102f4576107738460600151828151811061075d57fe5b602002602001015160200151846080015161044f565b156107825760009150506101cc565b60010161073d565b606082516000141561079b57600080fd5b81516107a657600080fd5b60606107bb8560200151866060015186610819565b905060606107d28660400151876060015186610819565b90506060866080015183836040516020016107ef93929190610fa1565b604051602081830303815290604052905061080e8760000151826104af565b979650505050505050565b60608061082685846101a2565b9050606061022185836102fe565b600082601f83011261084557600080fd5b81356108586108538261107b565b611055565b81815260209384019390925082018360005b83811015610896578135860161088088826109da565b845250602092830192919091019060010161086a565b5050505092915050565b80516101cc8161114b565b600082601f8301126108bc57600080fd5b81356108ca6108538261109b565b915080825260208301602083018583830111156108e657600080fd5b6108f18382846110f4565b50505092915050565b80356101cc81611162565b80356101cc8161116f565b60006080828403121561092257600080fd5b61092c6080611055565b905081356001600160401b0381111561094457600080fd5b610950848285016108ab565b82525060208201356001600160401b0381111561096c57600080fd5b610978848285016108ab565b60208301525060408201356001600160401b0381111561099757600080fd5b6109a384828501610a63565b60408301525060608201356001600160401b038111156109c257600080fd5b6109ce84828501610834565b60608301525092915050565b6000606082840312156109ec57600080fd5b6109f66060611055565b90506000610a0484846108fa565b82525060208201356001600160401b03811115610a2057600080fd5b610a2c848285016108ab565b60208301525060408201356001600160401b03811115610a4b57600080fd5b610a57848285016108ab565b60408301525092915050565b600060a08284031215610a7557600080fd5b610a7f60a0611055565b90506000610a8d84846108fa565b8252506020610a9e848483016108fa565b6020830152506040610ab2848285016108fa565b6040830152506060610ac684828501610905565b60608301525060808201356001600160401b03811115610ae557600080fd5b610af1848285016108ab565b60808301525092915050565b600060208284031215610b0f57600080fd5b600061027a84846108a0565b60008060408385031215610b2e57600080fd5b82356001600160401b03811115610b4457600080fd5b610b50858286016108ab565b92505060208301356001600160401b03811115610b6c57600080fd5b610b78858286016108ab565b9150509250929050565b60008060408385031215610b9557600080fd5b6000610b5085856108fa565b600080600060608486031215610bb657600080fd5b6000610bc286866108fa565b9350506020610bd386828701610905565b92505060408401356001600160401b03811115610bef57600080fd5b610bfb868287016108ab565b9150509250925092565b60008060408385031215610c1857600080fd5b6000610b508585610905565b600060208284031215610c3657600080fd5b81356001600160401b03811115610c4c57600080fd5b61027a84828501610910565b60008060408385031215610c6b57600080fd5b82356001600160401b03811115610c8157600080fd5b610c8d85828601610910565b92505060208301356001600160401b03811115610ca957600080fd5b610b7885828601610a63565b600080600080600060a08688031215610ccd57600080fd5b85356001600160401b03811115610ce357600080fd5b610cef88828901610910565b95505060208601356001600160401b03811115610d0b57600080fd5b610d1788828901610a63565b94505060408601356001600160401b03811115610d3357600080fd5b610d3f888289016108ab565b93505060608601356001600160401b03811115610d5b57600080fd5b610d67888289016108ab565b92505060808601356001600160401b03811115610d8357600080fd5b610d8f888289016108ab565b9150509295509295909350565b60008060408385031215610daf57600080fd5b82356001600160401b03811115610dc557600080fd5b610b50858286016109da565b600080600060608486031215610de657600080fd5b83356001600160401b03811115610dfc57600080fd5b610e0886828701610a63565b93505060208401356001600160401b03811115610e2457600080fd5b610bd3868287016108ab565b610e39816110d4565b82525050565b610e39610e4b826110d9565b6110eb565b610e39610e4b826110eb565b6000610e67826110c2565b610e7181856110c6565b9350610e81818560208601611100565b610e8a8161113b565b9093019392505050565b6000610e9f826110c2565b610ea981856110cf565b9350610eb9818560208601611100565b9290920192915050565b6000610ed06027836110c6565b7f696e76616c6964206f7220756e737570706f72746564206c656e677468206f7081526632b930ba34b7b760c91b602082015260400192915050565b6000610f196025836110c6565b7f696e76616c6964206f7220756e737570706f727465642068617368206f70657281526430ba34b7b760d91b602082015260400192915050565b610e39610f5f826110ee565b611130565b6000610f708284610e3f565b50601401919050565b6000610f858284610e50565b50602001919050565b6000610f9a8284610e94565b9392505050565b6000610fad8286610e94565b9150610fb98285610e94565b9150610fc58284610e94565b95945050505050565b6000610fda8285610f53565b60018201915061027a8284610e94565b6000610ff68286610f53565b6001820191506110068285610f53565b600182019150610fc58284610e94565b602081016101cc8284610e30565b602080825281016101c98184610e5c565b602080825281016101cc81610ec3565b602080825281016101cc81610f0c565b6040518181016001600160401b038111828210171561107357600080fd5b604052919050565b60006001600160401b0382111561109157600080fd5b5060209081020190565b60006001600160401b038211156110b157600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b919050565b151590565b6bffffffffffffffffffffffff191690565b90565b60ff1690565b82818337506000910152565b60005b8381101561111b578181015183820152602001611103565b8381111561112a576000848401525b50505050565b60006101cc82611145565b601f01601f191690565b60f81b90565b611154816110eb565b811461115f57600080fd5b50565b6006811061115f57600080fd5b6009811061115f57600080fdfea365627a7a723058207cad5ea356d3ce2231365cdc70196253a06927f855a325198ead8ef2746d58136c6578706572696d656e74616cf564736f6c63430005090040"

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

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	Key   []byte
	Value []byte
	Leaf  Struct0
	Path  []Struct1
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Hash   uint8
	Prefix []byte
	Suffix []byte
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Hash         uint8
	PrehashKey   uint8
	PrehashValue uint8
	Len          uint8
	Prefix       []byte
}

// ApplyInner is a free data retrieval call binding the contract method 0x3d4a397d.
//
// Solidity: function applyInner(Struct1 op, bytes child) constant returns(bytes)
func (_Proofs *ProofsCaller) ApplyInner(opts *bind.CallOpts, op Struct1, child []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "applyInner", op, child)
	return *ret0, err
}

// ApplyInner is a free data retrieval call binding the contract method 0x3d4a397d.
//
// Solidity: function applyInner(Struct1 op, bytes child) constant returns(bytes)
func (_Proofs *ProofsSession) ApplyInner(op Struct1, child []byte) ([]byte, error) {
	return _Proofs.Contract.ApplyInner(&_Proofs.CallOpts, op, child)
}

// ApplyInner is a free data retrieval call binding the contract method 0x3d4a397d.
//
// Solidity: function applyInner(Struct1 op, bytes child) constant returns(bytes)
func (_Proofs *ProofsCallerSession) ApplyInner(op Struct1, child []byte) ([]byte, error) {
	return _Proofs.Contract.ApplyInner(&_Proofs.CallOpts, op, child)
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xf6747d82.
//
// Solidity: function applyLeaf(Struct0 op, bytes key, bytes value) constant returns(bytes)
func (_Proofs *ProofsCaller) ApplyLeaf(opts *bind.CallOpts, op Struct0, key []byte, value []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "applyLeaf", op, key, value)
	return *ret0, err
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xf6747d82.
//
// Solidity: function applyLeaf(Struct0 op, bytes key, bytes value) constant returns(bytes)
func (_Proofs *ProofsSession) ApplyLeaf(op Struct0, key []byte, value []byte) ([]byte, error) {
	return _Proofs.Contract.ApplyLeaf(&_Proofs.CallOpts, op, key, value)
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xf6747d82.
//
// Solidity: function applyLeaf(Struct0 op, bytes key, bytes value) constant returns(bytes)
func (_Proofs *ProofsCallerSession) ApplyLeaf(op Struct0, key []byte, value []byte) ([]byte, error) {
	return _Proofs.Contract.ApplyLeaf(&_Proofs.CallOpts, op, key, value)
}

// Calculate is a free data retrieval call binding the contract method 0x7e1fd3bc.
//
// Solidity: function calculate(Struct2 proof) constant returns(bytes)
func (_Proofs *ProofsCaller) Calculate(opts *bind.CallOpts, proof Struct2) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "calculate", proof)
	return *ret0, err
}

// Calculate is a free data retrieval call binding the contract method 0x7e1fd3bc.
//
// Solidity: function calculate(Struct2 proof) constant returns(bytes)
func (_Proofs *ProofsSession) Calculate(proof Struct2) ([]byte, error) {
	return _Proofs.Contract.Calculate(&_Proofs.CallOpts, proof)
}

// Calculate is a free data retrieval call binding the contract method 0x7e1fd3bc.
//
// Solidity: function calculate(Struct2 proof) constant returns(bytes)
func (_Proofs *ProofsCallerSession) Calculate(proof Struct2) ([]byte, error) {
	return _Proofs.Contract.Calculate(&_Proofs.CallOpts, proof)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xf10e9a9c.
//
// Solidity: function checkAgainstSpec(Struct2 proof, Struct0 spec) constant returns(bool)
func (_Proofs *ProofsCaller) CheckAgainstSpec(opts *bind.CallOpts, proof Struct2, spec Struct0) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "checkAgainstSpec", proof, spec)
	return *ret0, err
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xf10e9a9c.
//
// Solidity: function checkAgainstSpec(Struct2 proof, Struct0 spec) constant returns(bool)
func (_Proofs *ProofsSession) CheckAgainstSpec(proof Struct2, spec Struct0) (bool, error) {
	return _Proofs.Contract.CheckAgainstSpec(&_Proofs.CallOpts, proof, spec)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xf10e9a9c.
//
// Solidity: function checkAgainstSpec(Struct2 proof, Struct0 spec) constant returns(bool)
func (_Proofs *ProofsCallerSession) CheckAgainstSpec(proof Struct2, spec Struct0) (bool, error) {
	return _Proofs.Contract.CheckAgainstSpec(&_Proofs.CallOpts, proof, spec)
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

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes bz1, bytes bz2) constant returns(bool)
func (_Proofs *ProofsCaller) EqualBytes(opts *bind.CallOpts, bz1 []byte, bz2 []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "equalBytes", bz1, bz2)
	return *ret0, err
}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes bz1, bytes bz2) constant returns(bool)
func (_Proofs *ProofsSession) EqualBytes(bz1 []byte, bz2 []byte) (bool, error) {
	return _Proofs.Contract.EqualBytes(&_Proofs.CallOpts, bz1, bz2)
}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes bz1, bytes bz2) constant returns(bool)
func (_Proofs *ProofsCallerSession) EqualBytes(bz1 []byte, bz2 []byte) (bool, error) {
	return _Proofs.Contract.EqualBytes(&_Proofs.CallOpts, bz1, bz2)
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

// VerifyExistence is a free data retrieval call binding the contract method 0x2e3098a9.
//
// Solidity: function verifyExistence(Struct2 proof, Struct0 spec, bytes root, bytes key, bytes value) constant returns(bool)
func (_Proofs *ProofsCaller) VerifyExistence(opts *bind.CallOpts, proof Struct2, spec Struct0, root []byte, key []byte, value []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Proofs.contract.Call(opts, out, "verifyExistence", proof, spec, root, key, value)
	return *ret0, err
}

// VerifyExistence is a free data retrieval call binding the contract method 0x2e3098a9.
//
// Solidity: function verifyExistence(Struct2 proof, Struct0 spec, bytes root, bytes key, bytes value) constant returns(bool)
func (_Proofs *ProofsSession) VerifyExistence(proof Struct2, spec Struct0, root []byte, key []byte, value []byte) (bool, error) {
	return _Proofs.Contract.VerifyExistence(&_Proofs.CallOpts, proof, spec, root, key, value)
}

// VerifyExistence is a free data retrieval call binding the contract method 0x2e3098a9.
//
// Solidity: function verifyExistence(Struct2 proof, Struct0 spec, bytes root, bytes key, bytes value) constant returns(bool)
func (_Proofs *ProofsCallerSession) VerifyExistence(proof Struct2, spec Struct0, root []byte, key []byte, value []byte) (bool, error) {
	return _Proofs.Contract.VerifyExistence(&_Proofs.CallOpts, proof, spec, root, key, value)
}
