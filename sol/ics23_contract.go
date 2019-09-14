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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ICS23ABI is the input ABI used to generate the binding from.
const ICS23ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"op\",\"type\":\"uint8\"},{\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHashOrNoop\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\"}],\"name\":\"path\",\"type\":\"tuple[]\"}],\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"spec\",\"type\":\"tuple\"},{\"name\":\"root\",\"type\":\"bytes\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyExistence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\"}],\"name\":\"op\",\"type\":\"tuple\"},{\"name\":\"child\",\"type\":\"bytes\"}],\"name\":\"applyInner\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"bz1\",\"type\":\"bytes\"},{\"name\":\"bz2\",\"type\":\"bytes\"}],\"name\":\"equalBytes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"op\",\"type\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"doLength\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\"}],\"name\":\"path\",\"type\":\"tuple[]\"}],\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"calculate\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"s\",\"type\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"hasprefix\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"spec\",\"type\":\"tuple\"},{\"name\":\"root\",\"type\":\"bytes\"},{\"components\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\"}],\"name\":\"path\",\"type\":\"tuple[]\"}],\"name\":\"proof\",\"type\":\"tuple\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyMembership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"op\",\"type\":\"uint8\"},{\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\"}],\"name\":\"path\",\"type\":\"tuple[]\"}],\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"spec\",\"type\":\"tuple\"}],\"name\":\"checkAgainstSpec\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"components\":[{\"name\":\"hash\",\"type\":\"uint8\"},{\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"name\":\"len\",\"type\":\"uint8\"},{\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"op\",\"type\":\"tuple\"},{\"name\":\"key\",\"type\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"applyLeaf\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hashop\",\"type\":\"uint8\"},{\"name\":\"lengthop\",\"type\":\"uint8\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"prepareLeafData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

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
var ICS23Bin = "0x60016080818152600060a081905260c083905260e083905261016060405261012083815261014082815261010091909152815460ff1916841762ffff001916620100001763ff0000001916630100000017825591929091620000639190816200007a565b5050503480156200007357600080fd5b506200011f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000bd57805160ff1916838001178555620000ed565b82800160010185558215620000ed579182015b82811115620000ed578251825591602001919060010190620000d0565b50620000fb929150620000ff565b5090565b6200011c91905b80821115620000fb576000815560010162000106565b90565b611276806200012f6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063901d0e1511610071578063901d0e151461014e578063a04f100814610161578063d48f1e4f14610174578063f10e9a9c14610187578063f6747d821461019a578063fd29e20a146101ad576100b4565b806303801174146100b95780632e3098a9146100e25780633d4a397d146101025780634cac70ff1461011557806367bb8e81146101285780637e1fd3bc1461013b575b600080fd5b6100cc6100c7366004610baf565b6101c0565b6040516100d991906110db565b60405180910390f35b6100f56100f0366004610ce2565b6101f0565b6040516100d991906110cd565b6100cc610110366004610dc9565b610249565b6100f5610123366004610b48565b6102a0565b6100cc610136366004610c32565b61031c565b6100cc610149366004610c51565b61040f565b6100f561015c366004610b48565b61046d565b6100f561016f366004610e5d565b6104cd565b6100cc610182366004610baf565b6104dc565b6100f5610195366004610c85565b610695565b6100cc6101a8366004610dfe565b6107b7565b6100cc6101bb366004610bce565b610846565b606060008360058111156101d057fe5b14156101dd5750806101ea565b6101e783836104dc565b90505b92915050565b60006101fc8686610695565b801561021157506102118387600001516102a0565b801561022657506102268287602001516102a0565b801561023f575061023f6102398761040f565b856102a0565b9695505050505050565b606081516000141561025a57600080fd5b6060836020015183856040015160405160200161027993929190611058565b60405160208183030381529060405290506102988460000151826104dc565b949350505050565b600081518351146102b3575060006101ea565b60005b8351811015610312578281815181106102cb57fe5b602001015160f81c60f81b6001600160f81b0319168482815181106102ec57fe5b01602001516001600160f81b0319161461030a5760009150506101ea565b6001016102b6565b5060019392505050565b6060600083600881111561032c57fe5b14156103395750806101ea565b600183600881111561034757fe5b14156103a3578151608081106103905780607f16608017600782901c91508184604051602001610379939291906110a1565b6040516020818303038152906040529150506101ea565b8083604051602001610379929190611085565b60078360088111156103b157fe5b14156103cc5781516020146103c557600080fd5b50806101ea565b60088360088111156103da57fe5b14156103ee5781516040146103c557600080fd5b60405162461bcd60e51b8152600401610406906110ec565b60405180910390fd5b6060806104298360400151846000015185602001516107b7565b905060005b8360600151518110156104665761045c8460600151828151811061044e57fe5b602002602001015183610249565b915060010161042e565b5092915050565b6000805b82518110156103125782818151811061048657fe5b602001015160f81c60f81b6001600160f81b0319168482815181106104a757fe5b01602001516001600160f81b031916146104c55760009150506101ea565b600101610471565b600061023f84878786866101f0565b606060038360058111156104ec57fe5b141561052057818051906020012060405160200161050a9190611030565b60405160208183030381529060405290506101ea565b600183600581111561052e57fe5b1415610594576002826040516105449190611045565b602060405180830381855afa158015610561573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052506105849190810190610b2a565b60405160200161050a9190611030565b60048360058111156105a257fe5b14156105ed576003826040516105b89190611045565b602060405180830381855afa1580156105d5573d6000803e3d6000fd5b5050604051805161050a925060601b9060200161101b565b60058360058111156105fb57fe5b141561067d5760036002836040516106139190611045565b602060405180830381855afa158015610630573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052506106539190810190610b2a565b6040516020016106639190611030565b60408051601f19818403018152908290526105b891611045565b60405162461bcd60e51b8152600401610406906110fc565b805160009060058111156106a557fe5b60408401515160058111156106b657fe5b1480156106e45750816020015160058111156106ce57fe5b83604001516020015160058111156106e257fe5b145b80156107115750816040015160058111156106fb57fe5b836040015160400151600581111561070f57fe5b145b801561073e57508160600151600881111561072857fe5b836040015160600151600881111561073c57fe5b145b801561075b575061075b836040015160800151836080015161046d565b610767575060006101ea565b60005b836060015151811015610312576107a08460600151828151811061078a57fe5b602002602001015160200151846080015161046d565b156107af5760009150506101ea565b60010161076a565b60608251600014156107c857600080fd5b81516107d357600080fd5b60606107e88560200151866060015186610846565b905060606107ff8660400151876060015186610846565b905060608660800151838360405160200161081c93929190611058565b604051602081830303815290604052905061083b8760000151826104dc565b979650505050505050565b60608061085385846101c0565b9050606061023f858361031c565b600082601f83011261087257600080fd5b813561088561088082611132565b61110c565b81815260209384019390925082018360005b838110156108c357813586016108ad8882610a07565b8452506020928301929190910190600101610897565b5050505092915050565b80516101ea81611202565b600082601f8301126108e957600080fd5b81356108f761088082611152565b9150808252602083016020830185838301111561091357600080fd5b61091e8382846111ab565b50505092915050565b80356101ea81611219565b80356101ea81611226565b60006080828403121561094f57600080fd5b610959608061110c565b905081356001600160401b0381111561097157600080fd5b61097d848285016108d8565b82525060208201356001600160401b0381111561099957600080fd5b6109a5848285016108d8565b60208301525060408201356001600160401b038111156109c457600080fd5b6109d084828501610a90565b60408301525060608201356001600160401b038111156109ef57600080fd5b6109fb84828501610861565b60608301525092915050565b600060608284031215610a1957600080fd5b610a23606061110c565b90506000610a318484610927565b82525060208201356001600160401b03811115610a4d57600080fd5b610a59848285016108d8565b60208301525060408201356001600160401b03811115610a7857600080fd5b610a84848285016108d8565b60408301525092915050565b600060a08284031215610aa257600080fd5b610aac60a061110c565b90506000610aba8484610927565b8252506020610acb84848301610927565b6020830152506040610adf84828501610927565b6040830152506060610af384828501610932565b60608301525060808201356001600160401b03811115610b1257600080fd5b610b1e848285016108d8565b60808301525092915050565b600060208284031215610b3c57600080fd5b600061029884846108cd565b60008060408385031215610b5b57600080fd5b82356001600160401b03811115610b7157600080fd5b610b7d858286016108d8565b92505060208301356001600160401b03811115610b9957600080fd5b610ba5858286016108d8565b9150509250929050565b60008060408385031215610bc257600080fd5b6000610b7d8585610927565b600080600060608486031215610be357600080fd5b6000610bef8686610927565b9350506020610c0086828701610932565b92505060408401356001600160401b03811115610c1c57600080fd5b610c28868287016108d8565b9150509250925092565b60008060408385031215610c4557600080fd5b6000610b7d8585610932565b600060208284031215610c6357600080fd5b81356001600160401b03811115610c7957600080fd5b6102988482850161093d565b60008060408385031215610c9857600080fd5b82356001600160401b03811115610cae57600080fd5b610cba8582860161093d565b92505060208301356001600160401b03811115610cd657600080fd5b610ba585828601610a90565b600080600080600060a08688031215610cfa57600080fd5b85356001600160401b03811115610d1057600080fd5b610d1c8882890161093d565b95505060208601356001600160401b03811115610d3857600080fd5b610d4488828901610a90565b94505060408601356001600160401b03811115610d6057600080fd5b610d6c888289016108d8565b93505060608601356001600160401b03811115610d8857600080fd5b610d94888289016108d8565b92505060808601356001600160401b03811115610db057600080fd5b610dbc888289016108d8565b9150509295509295909350565b60008060408385031215610ddc57600080fd5b82356001600160401b03811115610df257600080fd5b610b7d85828601610a07565b600080600060608486031215610e1357600080fd5b83356001600160401b03811115610e2957600080fd5b610e3586828701610a90565b93505060208401356001600160401b03811115610e5157600080fd5b610c00868287016108d8565b600080600080600060a08688031215610e7557600080fd5b85356001600160401b03811115610e8b57600080fd5b610e9788828901610a90565b95505060208601356001600160401b03811115610eb357600080fd5b610ebf888289016108d8565b94505060408601356001600160401b03811115610edb57600080fd5b610d6c8882890161093d565b610ef08161118b565b82525050565b610ef0610f0282611190565b6111a2565b610ef0610f02826111a2565b6000610f1e82611179565b610f28818561117d565b9350610f388185602086016111b7565b610f41816111f2565b9093019392505050565b6000610f5682611179565b610f608185611186565b9350610f708185602086016111b7565b9290920192915050565b6000610f8760278361117d565b7f696e76616c6964206f7220756e737570706f72746564206c656e677468206f7081526632b930ba34b7b760c91b602082015260400192915050565b6000610fd060258361117d565b7f696e76616c6964206f7220756e737570706f727465642068617368206f70657281526430ba34b7b760d91b602082015260400192915050565b610ef0611016826111a5565b6111e7565b60006110278284610ef6565b50601401919050565b600061103c8284610f07565b50602001919050565b60006110518284610f4b565b9392505050565b60006110648286610f4b565b91506110708285610f4b565b915061107c8284610f4b565b95945050505050565b6000611091828561100a565b6001820191506102988284610f4b565b60006110ad828661100a565b6001820191506110bd828561100a565b60018201915061107c8284610f4b565b602081016101ea8284610ee7565b602080825281016101e78184610f13565b602080825281016101ea81610f7a565b602080825281016101ea81610fc3565b6040518181016001600160401b038111828210171561112a57600080fd5b604052919050565b60006001600160401b0382111561114857600080fd5b5060209081020190565b60006001600160401b0382111561116857600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b919050565b151590565b6bffffffffffffffffffffffff191690565b90565b60ff1690565b82818337506000910152565b60005b838110156111d25781810151838201526020016111ba565b838111156111e1576000848401525b50505050565b60006101ea826111fc565b601f01601f191690565b60f81b90565b61120b816111a2565b811461121657600080fd5b50565b6006811061121657600080fd5b6009811061121657600080fdfea365627a7a7230582021e845f87264315bcfe13337b935b2d08290e3ad34917bca35059566373a87896c6578706572696d656e74616cf564736f6c63430005090040"

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
func (_ICS23 *ICS23Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ICS23 *ICS23CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_ICS23 *ICS23Caller) ApplyInner(opts *bind.CallOpts, op Struct1, child []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "applyInner", op, child)
	return *ret0, err
}

// ApplyInner is a free data retrieval call binding the contract method 0x3d4a397d.
//
// Solidity: function applyInner(Struct1 op, bytes child) constant returns(bytes)
func (_ICS23 *ICS23Session) ApplyInner(op Struct1, child []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyInner(&_ICS23.CallOpts, op, child)
}

// ApplyInner is a free data retrieval call binding the contract method 0x3d4a397d.
//
// Solidity: function applyInner(Struct1 op, bytes child) constant returns(bytes)
func (_ICS23 *ICS23CallerSession) ApplyInner(op Struct1, child []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyInner(&_ICS23.CallOpts, op, child)
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xf6747d82.
//
// Solidity: function applyLeaf(Struct0 op, bytes key, bytes value) constant returns(bytes)
func (_ICS23 *ICS23Caller) ApplyLeaf(opts *bind.CallOpts, op Struct0, key []byte, value []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "applyLeaf", op, key, value)
	return *ret0, err
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xf6747d82.
//
// Solidity: function applyLeaf(Struct0 op, bytes key, bytes value) constant returns(bytes)
func (_ICS23 *ICS23Session) ApplyLeaf(op Struct0, key []byte, value []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyLeaf(&_ICS23.CallOpts, op, key, value)
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xf6747d82.
//
// Solidity: function applyLeaf(Struct0 op, bytes key, bytes value) constant returns(bytes)
func (_ICS23 *ICS23CallerSession) ApplyLeaf(op Struct0, key []byte, value []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyLeaf(&_ICS23.CallOpts, op, key, value)
}

// Calculate is a free data retrieval call binding the contract method 0x7e1fd3bc.
//
// Solidity: function calculate(Struct2 proof) constant returns(bytes)
func (_ICS23 *ICS23Caller) Calculate(opts *bind.CallOpts, proof Struct2) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "calculate", proof)
	return *ret0, err
}

// Calculate is a free data retrieval call binding the contract method 0x7e1fd3bc.
//
// Solidity: function calculate(Struct2 proof) constant returns(bytes)
func (_ICS23 *ICS23Session) Calculate(proof Struct2) ([]byte, error) {
	return _ICS23.Contract.Calculate(&_ICS23.CallOpts, proof)
}

// Calculate is a free data retrieval call binding the contract method 0x7e1fd3bc.
//
// Solidity: function calculate(Struct2 proof) constant returns(bytes)
func (_ICS23 *ICS23CallerSession) Calculate(proof Struct2) ([]byte, error) {
	return _ICS23.Contract.Calculate(&_ICS23.CallOpts, proof)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xf10e9a9c.
//
// Solidity: function checkAgainstSpec(Struct2 proof, Struct0 spec) constant returns(bool)
func (_ICS23 *ICS23Caller) CheckAgainstSpec(opts *bind.CallOpts, proof Struct2, spec Struct0) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "checkAgainstSpec", proof, spec)
	return *ret0, err
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xf10e9a9c.
//
// Solidity: function checkAgainstSpec(Struct2 proof, Struct0 spec) constant returns(bool)
func (_ICS23 *ICS23Session) CheckAgainstSpec(proof Struct2, spec Struct0) (bool, error) {
	return _ICS23.Contract.CheckAgainstSpec(&_ICS23.CallOpts, proof, spec)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0xf10e9a9c.
//
// Solidity: function checkAgainstSpec(Struct2 proof, Struct0 spec) constant returns(bool)
func (_ICS23 *ICS23CallerSession) CheckAgainstSpec(proof Struct2, spec Struct0) (bool, error) {
	return _ICS23.Contract.CheckAgainstSpec(&_ICS23.CallOpts, proof, spec)
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 op, bytes preimage) constant returns(bytes)
func (_ICS23 *ICS23Caller) DoHash(opts *bind.CallOpts, op uint8, preimage []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "doHash", op, preimage)
	return *ret0, err
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 op, bytes preimage) constant returns(bytes)
func (_ICS23 *ICS23Session) DoHash(op uint8, preimage []byte) ([]byte, error) {
	return _ICS23.Contract.DoHash(&_ICS23.CallOpts, op, preimage)
}

// DoHash is a free data retrieval call binding the contract method 0xd48f1e4f.
//
// Solidity: function doHash(uint8 op, bytes preimage) constant returns(bytes)
func (_ICS23 *ICS23CallerSession) DoHash(op uint8, preimage []byte) ([]byte, error) {
	return _ICS23.Contract.DoHash(&_ICS23.CallOpts, op, preimage)
}

// DoHashOrNoop is a free data retrieval call binding the contract method 0x03801174.
//
// Solidity: function doHashOrNoop(uint8 op, bytes preimage) constant returns(bytes)
func (_ICS23 *ICS23Caller) DoHashOrNoop(opts *bind.CallOpts, op uint8, preimage []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "doHashOrNoop", op, preimage)
	return *ret0, err
}

// DoHashOrNoop is a free data retrieval call binding the contract method 0x03801174.
//
// Solidity: function doHashOrNoop(uint8 op, bytes preimage) constant returns(bytes)
func (_ICS23 *ICS23Session) DoHashOrNoop(op uint8, preimage []byte) ([]byte, error) {
	return _ICS23.Contract.DoHashOrNoop(&_ICS23.CallOpts, op, preimage)
}

// DoHashOrNoop is a free data retrieval call binding the contract method 0x03801174.
//
// Solidity: function doHashOrNoop(uint8 op, bytes preimage) constant returns(bytes)
func (_ICS23 *ICS23CallerSession) DoHashOrNoop(op uint8, preimage []byte) ([]byte, error) {
	return _ICS23.Contract.DoHashOrNoop(&_ICS23.CallOpts, op, preimage)
}

// DoLength is a free data retrieval call binding the contract method 0x67bb8e81.
//
// Solidity: function doLength(uint8 op, bytes data) constant returns(bytes)
func (_ICS23 *ICS23Caller) DoLength(opts *bind.CallOpts, op uint8, data []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "doLength", op, data)
	return *ret0, err
}

// DoLength is a free data retrieval call binding the contract method 0x67bb8e81.
//
// Solidity: function doLength(uint8 op, bytes data) constant returns(bytes)
func (_ICS23 *ICS23Session) DoLength(op uint8, data []byte) ([]byte, error) {
	return _ICS23.Contract.DoLength(&_ICS23.CallOpts, op, data)
}

// DoLength is a free data retrieval call binding the contract method 0x67bb8e81.
//
// Solidity: function doLength(uint8 op, bytes data) constant returns(bytes)
func (_ICS23 *ICS23CallerSession) DoLength(op uint8, data []byte) ([]byte, error) {
	return _ICS23.Contract.DoLength(&_ICS23.CallOpts, op, data)
}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes bz1, bytes bz2) constant returns(bool)
func (_ICS23 *ICS23Caller) EqualBytes(opts *bind.CallOpts, bz1 []byte, bz2 []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "equalBytes", bz1, bz2)
	return *ret0, err
}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes bz1, bytes bz2) constant returns(bool)
func (_ICS23 *ICS23Session) EqualBytes(bz1 []byte, bz2 []byte) (bool, error) {
	return _ICS23.Contract.EqualBytes(&_ICS23.CallOpts, bz1, bz2)
}

// EqualBytes is a free data retrieval call binding the contract method 0x4cac70ff.
//
// Solidity: function equalBytes(bytes bz1, bytes bz2) constant returns(bool)
func (_ICS23 *ICS23CallerSession) EqualBytes(bz1 []byte, bz2 []byte) (bool, error) {
	return _ICS23.Contract.EqualBytes(&_ICS23.CallOpts, bz1, bz2)
}

// Hasprefix is a free data retrieval call binding the contract method 0x901d0e15.
//
// Solidity: function hasprefix(bytes s, bytes prefix) constant returns(bool)
func (_ICS23 *ICS23Caller) Hasprefix(opts *bind.CallOpts, s []byte, prefix []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "hasprefix", s, prefix)
	return *ret0, err
}

// Hasprefix is a free data retrieval call binding the contract method 0x901d0e15.
//
// Solidity: function hasprefix(bytes s, bytes prefix) constant returns(bool)
func (_ICS23 *ICS23Session) Hasprefix(s []byte, prefix []byte) (bool, error) {
	return _ICS23.Contract.Hasprefix(&_ICS23.CallOpts, s, prefix)
}

// Hasprefix is a free data retrieval call binding the contract method 0x901d0e15.
//
// Solidity: function hasprefix(bytes s, bytes prefix) constant returns(bool)
func (_ICS23 *ICS23CallerSession) Hasprefix(s []byte, prefix []byte) (bool, error) {
	return _ICS23.Contract.Hasprefix(&_ICS23.CallOpts, s, prefix)
}

// PrepareLeafData is a free data retrieval call binding the contract method 0xfd29e20a.
//
// Solidity: function prepareLeafData(uint8 hashop, uint8 lengthop, bytes data) constant returns(bytes)
func (_ICS23 *ICS23Caller) PrepareLeafData(opts *bind.CallOpts, hashop uint8, lengthop uint8, data []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "prepareLeafData", hashop, lengthop, data)
	return *ret0, err
}

// PrepareLeafData is a free data retrieval call binding the contract method 0xfd29e20a.
//
// Solidity: function prepareLeafData(uint8 hashop, uint8 lengthop, bytes data) constant returns(bytes)
func (_ICS23 *ICS23Session) PrepareLeafData(hashop uint8, lengthop uint8, data []byte) ([]byte, error) {
	return _ICS23.Contract.PrepareLeafData(&_ICS23.CallOpts, hashop, lengthop, data)
}

// PrepareLeafData is a free data retrieval call binding the contract method 0xfd29e20a.
//
// Solidity: function prepareLeafData(uint8 hashop, uint8 lengthop, bytes data) constant returns(bytes)
func (_ICS23 *ICS23CallerSession) PrepareLeafData(hashop uint8, lengthop uint8, data []byte) ([]byte, error) {
	return _ICS23.Contract.PrepareLeafData(&_ICS23.CallOpts, hashop, lengthop, data)
}

// VerifyExistence is a free data retrieval call binding the contract method 0x2e3098a9.
//
// Solidity: function verifyExistence(Struct2 proof, Struct0 spec, bytes root, bytes key, bytes value) constant returns(bool)
func (_ICS23 *ICS23Caller) VerifyExistence(opts *bind.CallOpts, proof Struct2, spec Struct0, root []byte, key []byte, value []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "verifyExistence", proof, spec, root, key, value)
	return *ret0, err
}

// VerifyExistence is a free data retrieval call binding the contract method 0x2e3098a9.
//
// Solidity: function verifyExistence(Struct2 proof, Struct0 spec, bytes root, bytes key, bytes value) constant returns(bool)
func (_ICS23 *ICS23Session) VerifyExistence(proof Struct2, spec Struct0, root []byte, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyExistence(&_ICS23.CallOpts, proof, spec, root, key, value)
}

// VerifyExistence is a free data retrieval call binding the contract method 0x2e3098a9.
//
// Solidity: function verifyExistence(Struct2 proof, Struct0 spec, bytes root, bytes key, bytes value) constant returns(bool)
func (_ICS23 *ICS23CallerSession) VerifyExistence(proof Struct2, spec Struct0, root []byte, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyExistence(&_ICS23.CallOpts, proof, spec, root, key, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xa04f1008.
//
// Solidity: function verifyMembership(Struct0 spec, bytes root, Struct2 proof, bytes key, bytes value) constant returns(bool)
func (_ICS23 *ICS23Caller) VerifyMembership(opts *bind.CallOpts, spec Struct0, root []byte, proof Struct2, key []byte, value []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ICS23.contract.Call(opts, out, "verifyMembership", spec, root, proof, key, value)
	return *ret0, err
}

// VerifyMembership is a free data retrieval call binding the contract method 0xa04f1008.
//
// Solidity: function verifyMembership(Struct0 spec, bytes root, Struct2 proof, bytes key, bytes value) constant returns(bool)
func (_ICS23 *ICS23Session) VerifyMembership(spec Struct0, root []byte, proof Struct2, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyMembership(&_ICS23.CallOpts, spec, root, proof, key, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xa04f1008.
//
// Solidity: function verifyMembership(Struct0 spec, bytes root, Struct2 proof, bytes key, bytes value) constant returns(bool)
func (_ICS23 *ICS23CallerSession) VerifyMembership(spec Struct0, root []byte, proof Struct2, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyMembership(&_ICS23.CallOpts, spec, root, proof, key, value)
}
