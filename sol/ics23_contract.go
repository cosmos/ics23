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
	Valid bool
	Key   []byte
	Value []byte
	Leaf  ICS23LeafOp
	Path  []ICS23InnerOp
}

// ICS23InnerOp is an auto generated low-level Go binding around an user-defined struct.
type ICS23InnerOp struct {
	Valid  bool
	Hash   uint8
	Prefix []byte
	Suffix []byte
}

// ICS23InnerSpec is an auto generated low-level Go binding around an user-defined struct.
type ICS23InnerSpec struct {
	ChildOrder      []*big.Int
	ChildSize       *big.Int
	MinPrefixLength *big.Int
	MaxPrefixLength *big.Int
	EmptyChild      []byte
	Hash            uint8
}

// ICS23LeafOp is an auto generated low-level Go binding around an user-defined struct.
type ICS23LeafOp struct {
	Valid        bool
	Hash         uint8
	PrehashKey   uint8
	PrehashValue uint8
	Len          uint8
	Prefix       []byte
}

// ICS23NonExistenceProof is an auto generated low-level Go binding around an user-defined struct.
type ICS23NonExistenceProof struct {
	Valid bool
	Key   []byte
	Left  ICS23ExistenceProof
	Right ICS23ExistenceProof
}

// ICS23ProofSpec is an auto generated low-level Go binding around an user-defined struct.
type ICS23ProofSpec struct {
	LeafSpec  ICS23LeafOp
	InnerSpec ICS23InnerSpec
	MaxDepth  *big.Int
	MinDepth  *big.Int
}

// ICS23ABI is the input ABI used to generate the binding from.
const ICS23ABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"child\",\"type\":\"bytes\"}],\"name\":\"applyInner\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"applyLeaf\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"calculate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"spec\",\"type\":\"tuple\"}],\"name\":\"checkAgainstSpec\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bz1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bz2\",\"type\":\"bytes\"}],\"name\":\"compareBytes\",\"outputs\":[{\"internalType\":\"enumICS23.Ordering\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"preimage\",\"type\":\"bytes\"}],\"name\":\"doHashOrNoop\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"doLength\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"bz1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bz2\",\"type\":\"bytes\"}],\"name\":\"equalBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"branch\",\"type\":\"uint256\"}],\"name\":\"getPadding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"order\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"branch\",\"type\":\"uint256\"}],\"name\":\"getPosition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp\",\"name\":\"op\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minPrefix\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefix\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"suffix\",\"type\":\"uint256\"}],\"name\":\"hasPadding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"s\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"hasprefix\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"name\":\"isLeftMost\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"left\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"right\",\"type\":\"tuple[]\"}],\"name\":\"isLeftNeighbor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp\",\"name\":\"right\",\"type\":\"tuple\"}],\"name\":\"isLeftStep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"name\":\"isRightMost\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"spec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp\",\"name\":\"op\",\"type\":\"tuple\"}],\"name\":\"orderFromPadding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hashop\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"lengthop\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"prepareLeafData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyExistence\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"verifyMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structICS23.NonExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leafSpec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"innerSpec\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDepth\",\"type\":\"uint256\"}],\"internalType\":\"structICS23.ProofSpec\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"verifyNonExistence\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leafSpec\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"childOrder\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"childSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrefixLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"emptyChild\",\"type\":\"bytes\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"}],\"internalType\":\"structICS23.InnerSpec\",\"name\":\"innerSpec\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDepth\",\"type\":\"uint256\"}],\"internalType\":\"structICS23.ProofSpec\",\"name\":\"spec\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"left\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_key\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"prehash_value\",\"type\":\"uint8\"},{\"internalType\":\"enumICS23.LengthOp\",\"name\":\"len\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.LeafOp\",\"name\":\"leaf\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"enumICS23.HashOp\",\"name\":\"hash\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"suffix\",\"type\":\"bytes\"}],\"internalType\":\"structICS23.InnerOp[]\",\"name\":\"path\",\"type\":\"tuple[]\"}],\"internalType\":\"structICS23.ExistenceProof\",\"name\":\"right\",\"type\":\"tuple\"}],\"internalType\":\"structICS23.NonExistenceProof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"verifyNonMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ICS23FuncSigs maps the 4-byte function signature to its string representation.
var ICS23FuncSigs = map[string]string{
	"72aaf3df": "applyInner((bool,uint8,bytes,bytes),bytes)",
	"af630bb1": "applyLeaf((bool,uint8,uint8,uint8,uint8,bytes),bytes,bytes)",
	"e5c341e2": "calculate((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]))",
	"9d18a115": "checkAgainstSpec((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,uint8,uint8,uint8,uint8,bytes))",
	"d259183d": "compareBytes(bytes,bytes)",
	"d48f1e4f": "doHash(uint8,bytes)",
	"03801174": "doHashOrNoop(uint8,bytes)",
	"67bb8e81": "doLength(uint8,bytes)",
	"4cac70ff": "equalBytes(bytes,bytes)",
	"0d4383f4": "getPadding((uint256[],uint256,uint256,uint256,bytes,uint8),uint256)",
	"1e63e931": "getPosition(uint256[],uint256)",
	"9c854fbe": "hasPadding((bool,uint8,bytes,bytes),uint256,uint256,uint256)",
	"901d0e15": "hasprefix(bytes,bytes)",
	"951c0b90": "isLeftMost((uint256[],uint256,uint256,uint256,bytes,uint8),(bool,uint8,bytes,bytes)[])",
	"2f1cf262": "isLeftNeighbor((uint256[],uint256,uint256,uint256,bytes,uint8),(bool,uint8,bytes,bytes)[],(bool,uint8,bytes,bytes)[])",
	"b4219c6f": "isLeftStep((uint256[],uint256,uint256,uint256,bytes,uint8),(bool,uint8,bytes,bytes),(bool,uint8,bytes,bytes))",
	"83ead07c": "isRightMost((uint256[],uint256,uint256,uint256,bytes,uint8),(bool,uint8,bytes,bytes)[])",
	"356e77ff": "orderFromPadding((uint256[],uint256,uint256,uint256,bytes,uint8),(bool,uint8,bytes,bytes))",
	"fd29e20a": "prepareLeafData(uint8,uint8,bytes)",
	"452e99a3": "verifyExistence((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,uint8,uint8,uint8,uint8,bytes),bytes,bytes,bytes)",
	"3e339c30": "verifyMembership((bool,uint8,uint8,uint8,uint8,bytes),bytes,(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),bytes,bytes)",
	"b6446a5f": "verifyNonExistence((bool,bytes,(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[])),((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256),bytes,bytes)",
	"fbc2674d": "verifyNonMembership(((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256),bytes,(bool,bytes,(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[])),bytes)",
}

// ICS23Bin is the compiled bytecode used for deploying new contracts.
var ICS23Bin = "0x6001608081815260a0829052600060c081905260e083905261010083815261018060405261014084815261016083815261012091909152825460ff1916851761ff00191690911763ffff0000191663010000001760ff60201b191664010000000017825591929091620000749190816200008b565b5050503480156200008457600080fd5b506200016e565b828054620000999062000131565b90600052602060002090601f016020900481019282620000bd576000855562000108565b82601f10620000d857805160ff191683800117855562000108565b8280016001018555821562000108579182015b8281111562000108578251825591602001919060010190620000eb565b50620001169291506200011a565b5090565b5b808211156200011657600081556001016200011b565b6002810460018216806200014657607f821691505b602082108114156200016857634e487b7160e01b600052602260045260246000fd5b50919050565b612330806200017e6000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c8063901d0e15116100c3578063b6446a5f1161007c578063b6446a5f146102e4578063d259183d146102f7578063d48f1e4f14610317578063e5c341e21461032a578063fbc2674d1461033d578063fd29e20a146103505761014d565b8063901d0e1514610272578063951c0b90146102855780639c854fbe146102985780639d18a115146102ab578063af630bb1146102be578063b4219c6f146102d15761014d565b80633e339c30116101155780633e339c3014610200578063452e99a3146102135780634cac70ff1461022657806367bb8e811461023957806372aaf3df1461024c57806383ead07c1461025f5761014d565b806303801174146101525780630d4383f41461017b5780631e63e931146101a95780632f1cf262146101ca578063356e77ff146101ed575b600080fd5b610165610160366004611a48565b610363565b6040516101729190612187565b60405180910390f35b61018e610189366004611e75565b6103a7565b60408051938452602084019290925290820152606001610172565b6101bc6101b736600461198e565b610438565b604051908152602001610172565b6101dd6101d8366004611d2d565b6104a1565b6040519015158152602001610172565b6101bc6101fb366004611da6565b610682565b6101dd61020e366004611f21565b6106db565b6101dd610221366004611b87565b6106f4565b6101dd6102343660046119e8565b61073f565b610165610247366004611ae4565b6107ed565b61016561025a366004611c52565b61096c565b6101dd61026d366004611cd7565b6109c3565b6101dd6102803660046119e8565b610a60565b6101dd610293366004611cd7565b610af2565b6101dd6102a6366004611c86565b610b68565b6101dd6102b9366004611b31565b610ba5565b6101656102cc366004611ea8565b610d8a565b6101dd6102df366004611dfc565b610e19565b6101dd6102f2366004611f9c565b610e4d565b61030a6103053660046119e8565b610fe2565b60405161017291906121ba565b610165610325366004611a48565b611174565b610165610338366004611aff565b6113d1565b6101dd61034b366004612043565b611450565b61016561035e366004611a89565b611467565b6060600083600581111561038757634e487b7160e01b600052602160045260246000fd5b14156103945750806103a1565b61039e8383611174565b90505b92915050565b6000806000806103bb866000015186610438565b905060008660200151826103cf919061224d565b905060008760400151826103e39190612235565b905060008860600151836103f79190612235565b9050600089602001518560018c6000015151610413919061226c565b61041d919061226c565b610427919061224d565b929a91995091975095505050505050565b60008251821061044757600080fd5b60005b835181101561049b578284828151811061047457634e487b7160e01b600052603260045260246000fd5b602002602001015114156104895790506103a1565b80610493816122b3565b91505061044a565b50600080fd5b600060015b61051e848286516104b7919061226c565b815181106104d557634e487b7160e01b600052603260045260246000fd5b602002602001015160400151848386516104ef919061226c565b8151811061050d57634e487b7160e01b600052603260045260246000fd5b60200260200101516040015161073f565b801561059d575061059d84828651610536919061226c565b8151811061055457634e487b7160e01b600052603260045260246000fd5b6020026020010151606001518483865161056e919061226c565b8151811061058c57634e487b7160e01b600052603260045260246000fd5b60200260200101516060015161073f565b156105b457806105ac816122b3565b9150506104a6565b6000848286516105c4919061226c565b815181106105e257634e487b7160e01b600052603260045260246000fd5b602002602001015190506000848386516105fc919061226c565b8151811061061a57634e487b7160e01b600052603260045260246000fd5b6020026020010151905061062f878383610e19565b61063f576000935050505061067b565b61064987876109c3565b610659576000935050505061067b565b6106638786610af2565b610673576000935050505061067b565b600193505050505b9392505050565b815151600090815b8181101561049b5760008060006106a188856103a7565b9250925092506106b387848484610b68565b156106c55783955050505050506103a1565b50505080806106d3906122b3565b91505061068a565b60006106ea84878786866106f4565b9695505050505050565b60006107008686610ba5565b8015610715575061071583876020015161073f565b801561072a575061072a82876040015161073f565b80156106ea57506106ea61073d876113d1565b855b60008151835114610752575060006103a1565b60005b83518110156107e35782818151811061077e57634e487b7160e01b600052603260045260246000fd5b602001015160f81c60f81b6001600160f81b0319168482815181106107b357634e487b7160e01b600052603260045260246000fd5b01602001516001600160f81b031916146107d15760009150506103a1565b806107db816122b3565b915050610755565b5060019392505050565b6060600083600881111561081157634e487b7160e01b600052602160045260246000fd5b141561081e5750806103a1565b600183600881111561084057634e487b7160e01b600052602160045260246000fd5b141561089c578151608081106108895780607f16608017600782901c915081846040516020016108729392919061214a565b6040516020818303038152906040529150506103a1565b808360405160200161087292919061211b565b60078360088111156108be57634e487b7160e01b600052602160045260246000fd5b14156108d95781516020146108d257600080fd5b50806103a1565b60088360088111156108fb57634e487b7160e01b600052602160045260246000fd5b141561090f5781516040146108d257600080fd5b60405162461bcd60e51b815260206004820152602760248201527f696e76616c6964206f7220756e737570706f72746564206c656e677468206f7060448201526632b930ba34b7b760c91b60648201526084015b60405180910390fd5b606081516000141561097d57600080fd5b6000836040015183856060015160405160200161099c939291906120d8565b60405160208183030381529060405290506109bb846020015182611174565b949350505050565b60008060018460000151516109d8919061226c565b905060008060006109e987856103a7565b92509250925060005b8651811015610a5257610a2e878281518110610a1e57634e487b7160e01b600052603260045260246000fd5b6020026020010151858585610b68565b610a40576000955050505050506103a1565b80610a4a816122b3565b9150506109f2565b506001979650505050505050565b6000805b82518110156107e357828181518110610a8d57634e487b7160e01b600052603260045260246000fd5b602001015160f81c60f81b6001600160f81b031916848281518110610ac257634e487b7160e01b600052603260045260246000fd5b01602001516001600160f81b03191614610ae05760009150506103a1565b80610aea816122b3565b915050610a64565b600080600080610b038660006103a7565b92509250925060005b8551811015610b5b57610b38868281518110610a1e57634e487b7160e01b600052603260045260246000fd5b610b495760009450505050506103a1565b80610b53816122b3565b915050610b0c565b5060019695505050505050565b6000838560400151511015610b7f575060006109bb565b828560400151511115610b94575060006109bb565b506060840151518114949350505050565b600081602001516005811115610bcb57634e487b7160e01b600052602160045260246000fd5b8360600151602001516005811115610bf357634e487b7160e01b600052602160045260246000fd5b148015610c49575081604001516005811115610c1f57634e487b7160e01b600052602160045260246000fd5b8360600151604001516005811115610c4757634e487b7160e01b600052602160045260246000fd5b145b8015610c9e575081606001516005811115610c7457634e487b7160e01b600052602160045260246000fd5b8360600151606001516005811115610c9c57634e487b7160e01b600052602160045260246000fd5b145b8015610cf3575081608001516008811115610cc957634e487b7160e01b600052602160045260246000fd5b8360600151608001516008811115610cf157634e487b7160e01b600052602160045260246000fd5b145b8015610d105750610d10836060015160a001518360a00151610a60565b610d1c575060006103a1565b60005b8360800151518110156107e357610d6984608001518281518110610d5357634e487b7160e01b600052603260045260246000fd5b6020026020010151604001518460a00151610a60565b15610d785760009150506103a1565b80610d82816122b3565b915050610d1f565b6060825160001415610d9b57600080fd5b8151610da657600080fd5b6000610dbb8560400151866080015186611467565b90506000610dd28660600151876080015186611467565b905060008660a001518383604051602001610def939291906120d8565b6040516020818303038152906040529050610e0e876020015182611174565b979650505050505050565b600080610e268585610682565b90506000610e348685610682565b905080610e42836001612235565b149695505050505050565b604084015151600090158015610e665750606085015151155b15610e73575060006109bb565b60408501515115610ee6576040808601518551602082015192820151610e9a9387916106f4565b610ea6575060006109bb565b6000610eba83876060015160200151610fe2565b6002811115610ed957634e487b7160e01b600052602160045260246000fd5b14610ee6575060006109bb565b60608501515115610f5b576060850151845160208201516040830151610f0f93929187916106f4565b610f1b575060006109bb565b6002610f2f83876040015160200151610fe2565b6002811115610f4e57634e487b7160e01b600052602160045260246000fd5b14610f5b575060006109bb565b604085015151610f8c57610f7b8460200151866060015160800151610af2565b610f87575060006109bb565b610fd7565b606085015151610fac57610f7b84602001518660400151608001516109c3565b610fcb84602001518660400151608001518760600151608001516104a1565b610fd7575060006109bb565b506001949350505050565b6000805b835181108015610ff65750825181105b156110f55782818151811061101b57634e487b7160e01b600052603260045260246000fd5b602001015160f81c60f81b6001600160f81b03191684828151811061105057634e487b7160e01b600052603260045260246000fd5b01602001516001600160f81b031916101561106f5760009150506103a1565b82818151811061108f57634e487b7160e01b600052603260045260246000fd5b602001015160f81c60f81b6001600160f81b0319168482815181106110c457634e487b7160e01b600052603260045260246000fd5b01602001516001600160f81b03191611156110e35760029150506103a1565b806110ed816122b3565b915050610fe6565b50815183511415611108575060016103a1565b81518351101561111a575060006103a1565b81518351111561112c575060026103a1565b60405162461bcd60e51b815260206004820152601a60248201527f73686f756c64206e6f742072656163682074686973206c696e650000000000006044820152606401610963565b6060600383600581111561119857634e487b7160e01b600052602160045260246000fd5b14156111ce5781805190602001206040516020016111b891815260200190565b60405160208183030381529060405290506103a1565b60018360058111156111f057634e487b7160e01b600052602160045260246000fd5b14156112585760028260405161120691906120bc565b602060405180830381855afa158015611223573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061124691906119d0565b6040516020016111b891815260200190565b600483600581111561127a57634e487b7160e01b600052602160045260246000fd5b14156112d55760038260405161129091906120bc565b602060405180830381855afa1580156112ad573d6000803e3d6000fd5b5050604051805160601b6bffffffffffffffffffffffff1916602082015260340190506111b8565b60058360058111156112f757634e487b7160e01b600052602160045260246000fd5b141561137b57600360028360405161130f91906120bc565b602060405180830381855afa15801561132c573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061134f91906119d0565b60405160200161136191815260200190565b60408051601f1981840301815290829052611290916120bc565b60405162461bcd60e51b815260206004820152602560248201527f696e76616c6964206f7220756e737570706f727465642068617368206f70657260448201526430ba34b7b760d91b6064820152608401610963565b606060006113ec836060015184602001518560400151610d8a565b905060005b836080015151811015611447576114338460800151828151811061142557634e487b7160e01b600052603260045260246000fd5b60200260200101518361096c565b91508061143f816122b3565b9150506113f1565b5090505b919050565b600061145e83868685610e4d565b95945050505050565b606060006114758584610363565b905060006106ea85836107ed565b600082601f830112611493578081fd5b813560206114a86114a383612212565b6121e2565b82815281810190858301855b858110156114dd576114cb898684358b01016116a6565b845292840192908401906001016114b4565b5090979650505050505050565b600082601f8301126114fa578081fd5b8135602061150a6114a383612212565b8281528181019085830183850287018401881015611526578586fd5b855b858110156114dd57813584529284019290840190600101611528565b8035801515811461144b57600080fd5b600082601f830112611564578081fd5b81356001600160401b0381111561157d5761157d6122e4565b611590601f8201601f19166020016121e2565b8181528460208386010111156115a4578283fd5b816020850160208301379081016020019190915292915050565b80356006811061144b57600080fd5b80356009811061144b57600080fd5b600060a082840312156115ed578081fd5b6115f760a06121e2565b905061160282611544565b815260208201356001600160401b038082111561161e57600080fd5b61162a85838601611554565b6020840152604084013591508082111561164357600080fd5b61164f85838601611554565b6040840152606084013591508082111561166857600080fd5b611674858386016117d5565b6060840152608084013591508082111561168d57600080fd5b5061169a84828501611483565b60808301525092915050565b6000608082840312156116b7578081fd5b6116c160806121e2565b90506116cc82611544565b81526116da602083016115be565b602082015260408201356001600160401b03808211156116f957600080fd5b61170585838601611554565b6040840152606084013591508082111561171e57600080fd5b5061172b84828501611554565b60608301525092915050565b600060c08284031215611748578081fd5b61175260c06121e2565b905081356001600160401b038082111561176b57600080fd5b611777858386016114ea565b835260208401356020840152604084013560408401526060840135606084015260808401359150808211156117ab57600080fd5b506117b884828501611554565b6080830152506117ca60a083016115be565b60a082015292915050565b600060c082840312156117e6578081fd5b6117f060c06121e2565b90506117fb82611544565b8152611809602083016115be565b602082015261181a604083016115be565b604082015261182b606083016115be565b606082015261183c608083016115cd565b608082015260a08201356001600160401b0381111561185a57600080fd5b61186684828501611554565b60a08301525092915050565b600060808284031215611883578081fd5b61188d60806121e2565b905061189882611544565b815260208201356001600160401b03808211156118b457600080fd5b6118c085838601611554565b602084015260408401359150808211156118d957600080fd5b6118e5858386016115dc565b604084015260608401359150808211156118fe57600080fd5b5061172b848285016115dc565b60006080828403121561191c578081fd5b61192660806121e2565b905081356001600160401b038082111561193f57600080fd5b61194b858386016117d5565b8352602084013591508082111561196157600080fd5b5061196e84828501611737565b602083015250604082013560408201526060820135606082015292915050565b600080604083850312156119a0578182fd5b82356001600160401b038111156119b5578283fd5b6119c1858286016114ea565b95602094909401359450505050565b6000602082840312156119e1578081fd5b5051919050565b600080604083850312156119fa578182fd5b82356001600160401b0380821115611a10578384fd5b611a1c86838701611554565b93506020850135915080821115611a31578283fd5b50611a3e85828601611554565b9150509250929050565b60008060408385031215611a5a578182fd5b611a63836115be565b915060208301356001600160401b03811115611a7d578182fd5b611a3e85828601611554565b600080600060608486031215611a9d578081fd5b611aa6846115be565b9250611ab4602085016115cd565b915060408401356001600160401b03811115611ace578182fd5b611ada86828701611554565b9150509250925092565b60008060408385031215611af6578182fd5b611a63836115cd565b600060208284031215611b10578081fd5b81356001600160401b03811115611b25578182fd5b6109bb848285016115dc565b60008060408385031215611b43578182fd5b82356001600160401b0380821115611b59578384fd5b611b65868387016115dc565b93506020850135915080821115611b7a578283fd5b50611a3e858286016117d5565b600080600080600060a08688031215611b9e578283fd5b85356001600160401b0380821115611bb4578485fd5b611bc089838a016115dc565b96506020880135915080821115611bd5578485fd5b611be189838a016117d5565b95506040880135915080821115611bf6578485fd5b611c0289838a01611554565b94506060880135915080821115611c17578283fd5b611c2389838a01611554565b93506080880135915080821115611c38578283fd5b50611c4588828901611554565b9150509295509295909350565b60008060408385031215611c64578182fd5b82356001600160401b0380821115611c7a578384fd5b611a1c868387016116a6565b60008060008060808587031215611c9b578182fd5b84356001600160401b03811115611cb0578283fd5b611cbc878288016116a6565b97602087013597506040870135966060013595509350505050565b60008060408385031215611ce9578182fd5b82356001600160401b0380821115611cff578384fd5b611d0b86838701611737565b93506020850135915080821115611d20578283fd5b50611a3e85828601611483565b600080600060608486031215611d41578081fd5b83356001600160401b0380821115611d57578283fd5b611d6387838801611737565b94506020860135915080821115611d78578283fd5b611d8487838801611483565b93506040860135915080821115611d99578283fd5b50611ada86828701611483565b60008060408385031215611db8578182fd5b82356001600160401b0380821115611dce578384fd5b611dda86838701611737565b93506020850135915080821115611def578283fd5b50611a3e858286016116a6565b600080600060608486031215611e10578081fd5b83356001600160401b0380821115611e26578283fd5b611e3287838801611737565b94506020860135915080821115611e47578283fd5b611e53878388016116a6565b93506040860135915080821115611e68578283fd5b50611ada868287016116a6565b60008060408385031215611e87578182fd5b82356001600160401b03811115611e9c578283fd5b6119c185828601611737565b600080600060608486031215611ebc578081fd5b83356001600160401b0380821115611ed2578283fd5b611ede878388016117d5565b94506020860135915080821115611ef3578283fd5b611eff87838801611554565b93506040860135915080821115611f14578283fd5b50611ada86828701611554565b600080600080600060a08688031215611f38578283fd5b85356001600160401b0380821115611f4e578485fd5b611f5a89838a016117d5565b96506020880135915080821115611f6f578485fd5b611f7b89838a01611554565b95506040880135915080821115611f90578485fd5b611c0289838a016115dc565b60008060008060808587031215611fb1578182fd5b84356001600160401b0380821115611fc7578384fd5b611fd388838901611872565b95506020870135915080821115611fe8578384fd5b611ff48883890161190b565b94506040870135915080821115612009578384fd5b61201588838901611554565b9350606087013591508082111561202a578283fd5b5061203787828801611554565b91505092959194509250565b60008060008060808587031215612058578182fd5b84356001600160401b038082111561206e578384fd5b61207a8883890161190b565b9550602087013591508082111561208f578384fd5b61209b88838901611554565b945060408701359150808211156120b0578384fd5b61201588838901611872565b600082516120ce818460208701612283565b9190910192915050565b600084516120ea818460208901612283565b8451908301906120fe818360208901612283565b8451910190612111818360208801612283565b0195945050505050565b600060ff60f81b8460f81b168252825161213c816001850160208701612283565b919091016001019392505050565b600060ff60f81b808660f81b168352808560f81b166001840152508251612178816002850160208701612283565b91909101600201949350505050565b60006020825282518060208401526121a6816040850160208701612283565b601f01601f19169190910160400192915050565b60208101600383106121dc57634e487b7160e01b600052602160045260246000fd5b91905290565b604051601f8201601f191681016001600160401b038111828210171561220a5761220a6122e4565b604052919050565b60006001600160401b0382111561222b5761222b6122e4565b5060209081020190565b60008219821115612248576122486122ce565b500190565b6000816000190483118215151615612267576122676122ce565b500290565b60008282101561227e5761227e6122ce565b500390565b60005b8381101561229e578181015183820152602001612286565b838111156122ad576000848401525b50505050565b60006000198214156122c7576122c76122ce565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea264697066735822122009f5f964c8d1a95c87007ce8be83d929c9b3b7d97d2fe150b7a385d931ce4b2764736f6c63430008020033"

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

// ApplyInner is a free data retrieval call binding the contract method 0x72aaf3df.
//
// Solidity: function applyInner((bool,uint8,bytes,bytes) op, bytes child) pure returns(bytes)
func (_ICS23 *ICS23Caller) ApplyInner(opts *bind.CallOpts, op ICS23InnerOp, child []byte) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "applyInner", op, child)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ApplyInner is a free data retrieval call binding the contract method 0x72aaf3df.
//
// Solidity: function applyInner((bool,uint8,bytes,bytes) op, bytes child) pure returns(bytes)
func (_ICS23 *ICS23Session) ApplyInner(op ICS23InnerOp, child []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyInner(&_ICS23.CallOpts, op, child)
}

// ApplyInner is a free data retrieval call binding the contract method 0x72aaf3df.
//
// Solidity: function applyInner((bool,uint8,bytes,bytes) op, bytes child) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) ApplyInner(op ICS23InnerOp, child []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyInner(&_ICS23.CallOpts, op, child)
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xaf630bb1.
//
// Solidity: function applyLeaf((bool,uint8,uint8,uint8,uint8,bytes) op, bytes key, bytes value) pure returns(bytes)
func (_ICS23 *ICS23Caller) ApplyLeaf(opts *bind.CallOpts, op ICS23LeafOp, key []byte, value []byte) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "applyLeaf", op, key, value)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ApplyLeaf is a free data retrieval call binding the contract method 0xaf630bb1.
//
// Solidity: function applyLeaf((bool,uint8,uint8,uint8,uint8,bytes) op, bytes key, bytes value) pure returns(bytes)
func (_ICS23 *ICS23Session) ApplyLeaf(op ICS23LeafOp, key []byte, value []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyLeaf(&_ICS23.CallOpts, op, key, value)
}

// ApplyLeaf is a free data retrieval call binding the contract method 0xaf630bb1.
//
// Solidity: function applyLeaf((bool,uint8,uint8,uint8,uint8,bytes) op, bytes key, bytes value) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) ApplyLeaf(op ICS23LeafOp, key []byte, value []byte) ([]byte, error) {
	return _ICS23.Contract.ApplyLeaf(&_ICS23.CallOpts, op, key, value)
}

// Calculate is a free data retrieval call binding the contract method 0xe5c341e2.
//
// Solidity: function calculate((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof) pure returns(bytes)
func (_ICS23 *ICS23Caller) Calculate(opts *bind.CallOpts, proof ICS23ExistenceProof) ([]byte, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "calculate", proof)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Calculate is a free data retrieval call binding the contract method 0xe5c341e2.
//
// Solidity: function calculate((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof) pure returns(bytes)
func (_ICS23 *ICS23Session) Calculate(proof ICS23ExistenceProof) ([]byte, error) {
	return _ICS23.Contract.Calculate(&_ICS23.CallOpts, proof)
}

// Calculate is a free data retrieval call binding the contract method 0xe5c341e2.
//
// Solidity: function calculate((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof) pure returns(bytes)
func (_ICS23 *ICS23CallerSession) Calculate(proof ICS23ExistenceProof) ([]byte, error) {
	return _ICS23.Contract.Calculate(&_ICS23.CallOpts, proof)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0x9d18a115.
//
// Solidity: function checkAgainstSpec((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, (bool,uint8,uint8,uint8,uint8,bytes) spec) pure returns(bool)
func (_ICS23 *ICS23Caller) CheckAgainstSpec(opts *bind.CallOpts, proof ICS23ExistenceProof, spec ICS23LeafOp) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "checkAgainstSpec", proof, spec)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0x9d18a115.
//
// Solidity: function checkAgainstSpec((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, (bool,uint8,uint8,uint8,uint8,bytes) spec) pure returns(bool)
func (_ICS23 *ICS23Session) CheckAgainstSpec(proof ICS23ExistenceProof, spec ICS23LeafOp) (bool, error) {
	return _ICS23.Contract.CheckAgainstSpec(&_ICS23.CallOpts, proof, spec)
}

// CheckAgainstSpec is a free data retrieval call binding the contract method 0x9d18a115.
//
// Solidity: function checkAgainstSpec((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, (bool,uint8,uint8,uint8,uint8,bytes) spec) pure returns(bool)
func (_ICS23 *ICS23CallerSession) CheckAgainstSpec(proof ICS23ExistenceProof, spec ICS23LeafOp) (bool, error) {
	return _ICS23.Contract.CheckAgainstSpec(&_ICS23.CallOpts, proof, spec)
}

// CompareBytes is a free data retrieval call binding the contract method 0xd259183d.
//
// Solidity: function compareBytes(bytes bz1, bytes bz2) pure returns(uint8)
func (_ICS23 *ICS23Caller) CompareBytes(opts *bind.CallOpts, bz1 []byte, bz2 []byte) (uint8, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "compareBytes", bz1, bz2)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CompareBytes is a free data retrieval call binding the contract method 0xd259183d.
//
// Solidity: function compareBytes(bytes bz1, bytes bz2) pure returns(uint8)
func (_ICS23 *ICS23Session) CompareBytes(bz1 []byte, bz2 []byte) (uint8, error) {
	return _ICS23.Contract.CompareBytes(&_ICS23.CallOpts, bz1, bz2)
}

// CompareBytes is a free data retrieval call binding the contract method 0xd259183d.
//
// Solidity: function compareBytes(bytes bz1, bytes bz2) pure returns(uint8)
func (_ICS23 *ICS23CallerSession) CompareBytes(bz1 []byte, bz2 []byte) (uint8, error) {
	return _ICS23.Contract.CompareBytes(&_ICS23.CallOpts, bz1, bz2)
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

// GetPadding is a free data retrieval call binding the contract method 0x0d4383f4.
//
// Solidity: function getPadding((uint256[],uint256,uint256,uint256,bytes,uint8) spec, uint256 branch) pure returns(uint256, uint256, uint256)
func (_ICS23 *ICS23Caller) GetPadding(opts *bind.CallOpts, spec ICS23InnerSpec, branch *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "getPadding", spec, branch)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetPadding is a free data retrieval call binding the contract method 0x0d4383f4.
//
// Solidity: function getPadding((uint256[],uint256,uint256,uint256,bytes,uint8) spec, uint256 branch) pure returns(uint256, uint256, uint256)
func (_ICS23 *ICS23Session) GetPadding(spec ICS23InnerSpec, branch *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ICS23.Contract.GetPadding(&_ICS23.CallOpts, spec, branch)
}

// GetPadding is a free data retrieval call binding the contract method 0x0d4383f4.
//
// Solidity: function getPadding((uint256[],uint256,uint256,uint256,bytes,uint8) spec, uint256 branch) pure returns(uint256, uint256, uint256)
func (_ICS23 *ICS23CallerSession) GetPadding(spec ICS23InnerSpec, branch *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ICS23.Contract.GetPadding(&_ICS23.CallOpts, spec, branch)
}

// GetPosition is a free data retrieval call binding the contract method 0x1e63e931.
//
// Solidity: function getPosition(uint256[] order, uint256 branch) pure returns(uint256)
func (_ICS23 *ICS23Caller) GetPosition(opts *bind.CallOpts, order []*big.Int, branch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "getPosition", order, branch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPosition is a free data retrieval call binding the contract method 0x1e63e931.
//
// Solidity: function getPosition(uint256[] order, uint256 branch) pure returns(uint256)
func (_ICS23 *ICS23Session) GetPosition(order []*big.Int, branch *big.Int) (*big.Int, error) {
	return _ICS23.Contract.GetPosition(&_ICS23.CallOpts, order, branch)
}

// GetPosition is a free data retrieval call binding the contract method 0x1e63e931.
//
// Solidity: function getPosition(uint256[] order, uint256 branch) pure returns(uint256)
func (_ICS23 *ICS23CallerSession) GetPosition(order []*big.Int, branch *big.Int) (*big.Int, error) {
	return _ICS23.Contract.GetPosition(&_ICS23.CallOpts, order, branch)
}

// HasPadding is a free data retrieval call binding the contract method 0x9c854fbe.
//
// Solidity: function hasPadding((bool,uint8,bytes,bytes) op, uint256 minPrefix, uint256 maxPrefix, uint256 suffix) pure returns(bool)
func (_ICS23 *ICS23Caller) HasPadding(opts *bind.CallOpts, op ICS23InnerOp, minPrefix *big.Int, maxPrefix *big.Int, suffix *big.Int) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "hasPadding", op, minPrefix, maxPrefix, suffix)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPadding is a free data retrieval call binding the contract method 0x9c854fbe.
//
// Solidity: function hasPadding((bool,uint8,bytes,bytes) op, uint256 minPrefix, uint256 maxPrefix, uint256 suffix) pure returns(bool)
func (_ICS23 *ICS23Session) HasPadding(op ICS23InnerOp, minPrefix *big.Int, maxPrefix *big.Int, suffix *big.Int) (bool, error) {
	return _ICS23.Contract.HasPadding(&_ICS23.CallOpts, op, minPrefix, maxPrefix, suffix)
}

// HasPadding is a free data retrieval call binding the contract method 0x9c854fbe.
//
// Solidity: function hasPadding((bool,uint8,bytes,bytes) op, uint256 minPrefix, uint256 maxPrefix, uint256 suffix) pure returns(bool)
func (_ICS23 *ICS23CallerSession) HasPadding(op ICS23InnerOp, minPrefix *big.Int, maxPrefix *big.Int, suffix *big.Int) (bool, error) {
	return _ICS23.Contract.HasPadding(&_ICS23.CallOpts, op, minPrefix, maxPrefix, suffix)
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

// IsLeftMost is a free data retrieval call binding the contract method 0x951c0b90.
//
// Solidity: function isLeftMost((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes)[] path) pure returns(bool)
func (_ICS23 *ICS23Caller) IsLeftMost(opts *bind.CallOpts, spec ICS23InnerSpec, path []ICS23InnerOp) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "isLeftMost", spec, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeftMost is a free data retrieval call binding the contract method 0x951c0b90.
//
// Solidity: function isLeftMost((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes)[] path) pure returns(bool)
func (_ICS23 *ICS23Session) IsLeftMost(spec ICS23InnerSpec, path []ICS23InnerOp) (bool, error) {
	return _ICS23.Contract.IsLeftMost(&_ICS23.CallOpts, spec, path)
}

// IsLeftMost is a free data retrieval call binding the contract method 0x951c0b90.
//
// Solidity: function isLeftMost((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes)[] path) pure returns(bool)
func (_ICS23 *ICS23CallerSession) IsLeftMost(spec ICS23InnerSpec, path []ICS23InnerOp) (bool, error) {
	return _ICS23.Contract.IsLeftMost(&_ICS23.CallOpts, spec, path)
}

// IsLeftNeighbor is a free data retrieval call binding the contract method 0x2f1cf262.
//
// Solidity: function isLeftNeighbor((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes)[] left, (bool,uint8,bytes,bytes)[] right) pure returns(bool)
func (_ICS23 *ICS23Caller) IsLeftNeighbor(opts *bind.CallOpts, spec ICS23InnerSpec, left []ICS23InnerOp, right []ICS23InnerOp) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "isLeftNeighbor", spec, left, right)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeftNeighbor is a free data retrieval call binding the contract method 0x2f1cf262.
//
// Solidity: function isLeftNeighbor((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes)[] left, (bool,uint8,bytes,bytes)[] right) pure returns(bool)
func (_ICS23 *ICS23Session) IsLeftNeighbor(spec ICS23InnerSpec, left []ICS23InnerOp, right []ICS23InnerOp) (bool, error) {
	return _ICS23.Contract.IsLeftNeighbor(&_ICS23.CallOpts, spec, left, right)
}

// IsLeftNeighbor is a free data retrieval call binding the contract method 0x2f1cf262.
//
// Solidity: function isLeftNeighbor((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes)[] left, (bool,uint8,bytes,bytes)[] right) pure returns(bool)
func (_ICS23 *ICS23CallerSession) IsLeftNeighbor(spec ICS23InnerSpec, left []ICS23InnerOp, right []ICS23InnerOp) (bool, error) {
	return _ICS23.Contract.IsLeftNeighbor(&_ICS23.CallOpts, spec, left, right)
}

// IsLeftStep is a free data retrieval call binding the contract method 0xb4219c6f.
//
// Solidity: function isLeftStep((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes) left, (bool,uint8,bytes,bytes) right) pure returns(bool)
func (_ICS23 *ICS23Caller) IsLeftStep(opts *bind.CallOpts, spec ICS23InnerSpec, left ICS23InnerOp, right ICS23InnerOp) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "isLeftStep", spec, left, right)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeftStep is a free data retrieval call binding the contract method 0xb4219c6f.
//
// Solidity: function isLeftStep((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes) left, (bool,uint8,bytes,bytes) right) pure returns(bool)
func (_ICS23 *ICS23Session) IsLeftStep(spec ICS23InnerSpec, left ICS23InnerOp, right ICS23InnerOp) (bool, error) {
	return _ICS23.Contract.IsLeftStep(&_ICS23.CallOpts, spec, left, right)
}

// IsLeftStep is a free data retrieval call binding the contract method 0xb4219c6f.
//
// Solidity: function isLeftStep((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes) left, (bool,uint8,bytes,bytes) right) pure returns(bool)
func (_ICS23 *ICS23CallerSession) IsLeftStep(spec ICS23InnerSpec, left ICS23InnerOp, right ICS23InnerOp) (bool, error) {
	return _ICS23.Contract.IsLeftStep(&_ICS23.CallOpts, spec, left, right)
}

// IsRightMost is a free data retrieval call binding the contract method 0x83ead07c.
//
// Solidity: function isRightMost((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes)[] path) pure returns(bool)
func (_ICS23 *ICS23Caller) IsRightMost(opts *bind.CallOpts, spec ICS23InnerSpec, path []ICS23InnerOp) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "isRightMost", spec, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRightMost is a free data retrieval call binding the contract method 0x83ead07c.
//
// Solidity: function isRightMost((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes)[] path) pure returns(bool)
func (_ICS23 *ICS23Session) IsRightMost(spec ICS23InnerSpec, path []ICS23InnerOp) (bool, error) {
	return _ICS23.Contract.IsRightMost(&_ICS23.CallOpts, spec, path)
}

// IsRightMost is a free data retrieval call binding the contract method 0x83ead07c.
//
// Solidity: function isRightMost((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes)[] path) pure returns(bool)
func (_ICS23 *ICS23CallerSession) IsRightMost(spec ICS23InnerSpec, path []ICS23InnerOp) (bool, error) {
	return _ICS23.Contract.IsRightMost(&_ICS23.CallOpts, spec, path)
}

// OrderFromPadding is a free data retrieval call binding the contract method 0x356e77ff.
//
// Solidity: function orderFromPadding((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes) op) pure returns(uint256)
func (_ICS23 *ICS23Caller) OrderFromPadding(opts *bind.CallOpts, spec ICS23InnerSpec, op ICS23InnerOp) (*big.Int, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "orderFromPadding", spec, op)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderFromPadding is a free data retrieval call binding the contract method 0x356e77ff.
//
// Solidity: function orderFromPadding((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes) op) pure returns(uint256)
func (_ICS23 *ICS23Session) OrderFromPadding(spec ICS23InnerSpec, op ICS23InnerOp) (*big.Int, error) {
	return _ICS23.Contract.OrderFromPadding(&_ICS23.CallOpts, spec, op)
}

// OrderFromPadding is a free data retrieval call binding the contract method 0x356e77ff.
//
// Solidity: function orderFromPadding((uint256[],uint256,uint256,uint256,bytes,uint8) spec, (bool,uint8,bytes,bytes) op) pure returns(uint256)
func (_ICS23 *ICS23CallerSession) OrderFromPadding(spec ICS23InnerSpec, op ICS23InnerOp) (*big.Int, error) {
	return _ICS23.Contract.OrderFromPadding(&_ICS23.CallOpts, spec, op)
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

// VerifyExistence is a free data retrieval call binding the contract method 0x452e99a3.
//
// Solidity: function verifyExistence((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, (bool,uint8,uint8,uint8,uint8,bytes) spec, bytes root, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23Caller) VerifyExistence(opts *bind.CallOpts, proof ICS23ExistenceProof, spec ICS23LeafOp, root []byte, key []byte, value []byte) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "verifyExistence", proof, spec, root, key, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyExistence is a free data retrieval call binding the contract method 0x452e99a3.
//
// Solidity: function verifyExistence((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, (bool,uint8,uint8,uint8,uint8,bytes) spec, bytes root, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23Session) VerifyExistence(proof ICS23ExistenceProof, spec ICS23LeafOp, root []byte, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyExistence(&_ICS23.CallOpts, proof, spec, root, key, value)
}

// VerifyExistence is a free data retrieval call binding the contract method 0x452e99a3.
//
// Solidity: function verifyExistence((bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, (bool,uint8,uint8,uint8,uint8,bytes) spec, bytes root, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23CallerSession) VerifyExistence(proof ICS23ExistenceProof, spec ICS23LeafOp, root []byte, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyExistence(&_ICS23.CallOpts, proof, spec, root, key, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x3e339c30.
//
// Solidity: function verifyMembership((bool,uint8,uint8,uint8,uint8,bytes) spec, bytes root, (bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23Caller) VerifyMembership(opts *bind.CallOpts, spec ICS23LeafOp, root []byte, proof ICS23ExistenceProof, key []byte, value []byte) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "verifyMembership", spec, root, proof, key, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMembership is a free data retrieval call binding the contract method 0x3e339c30.
//
// Solidity: function verifyMembership((bool,uint8,uint8,uint8,uint8,bytes) spec, bytes root, (bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23Session) VerifyMembership(spec ICS23LeafOp, root []byte, proof ICS23ExistenceProof, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyMembership(&_ICS23.CallOpts, spec, root, proof, key, value)
}

// VerifyMembership is a free data retrieval call binding the contract method 0x3e339c30.
//
// Solidity: function verifyMembership((bool,uint8,uint8,uint8,uint8,bytes) spec, bytes root, (bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]) proof, bytes key, bytes value) pure returns(bool)
func (_ICS23 *ICS23CallerSession) VerifyMembership(spec ICS23LeafOp, root []byte, proof ICS23ExistenceProof, key []byte, value []byte) (bool, error) {
	return _ICS23.Contract.VerifyMembership(&_ICS23.CallOpts, spec, root, proof, key, value)
}

// VerifyNonExistence is a free data retrieval call binding the contract method 0xb6446a5f.
//
// Solidity: function verifyNonExistence((bool,bytes,(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[])) proof, ((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256) spec, bytes root, bytes key) pure returns(bool)
func (_ICS23 *ICS23Caller) VerifyNonExistence(opts *bind.CallOpts, proof ICS23NonExistenceProof, spec ICS23ProofSpec, root []byte, key []byte) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "verifyNonExistence", proof, spec, root, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyNonExistence is a free data retrieval call binding the contract method 0xb6446a5f.
//
// Solidity: function verifyNonExistence((bool,bytes,(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[])) proof, ((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256) spec, bytes root, bytes key) pure returns(bool)
func (_ICS23 *ICS23Session) VerifyNonExistence(proof ICS23NonExistenceProof, spec ICS23ProofSpec, root []byte, key []byte) (bool, error) {
	return _ICS23.Contract.VerifyNonExistence(&_ICS23.CallOpts, proof, spec, root, key)
}

// VerifyNonExistence is a free data retrieval call binding the contract method 0xb6446a5f.
//
// Solidity: function verifyNonExistence((bool,bytes,(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[])) proof, ((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256) spec, bytes root, bytes key) pure returns(bool)
func (_ICS23 *ICS23CallerSession) VerifyNonExistence(proof ICS23NonExistenceProof, spec ICS23ProofSpec, root []byte, key []byte) (bool, error) {
	return _ICS23.Contract.VerifyNonExistence(&_ICS23.CallOpts, proof, spec, root, key)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfbc2674d.
//
// Solidity: function verifyNonMembership(((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256) spec, bytes root, (bool,bytes,(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[])) proof, bytes key) pure returns(bool)
func (_ICS23 *ICS23Caller) VerifyNonMembership(opts *bind.CallOpts, spec ICS23ProofSpec, root []byte, proof ICS23NonExistenceProof, key []byte) (bool, error) {
	var out []interface{}
	err := _ICS23.contract.Call(opts, &out, "verifyNonMembership", spec, root, proof, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfbc2674d.
//
// Solidity: function verifyNonMembership(((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256) spec, bytes root, (bool,bytes,(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[])) proof, bytes key) pure returns(bool)
func (_ICS23 *ICS23Session) VerifyNonMembership(spec ICS23ProofSpec, root []byte, proof ICS23NonExistenceProof, key []byte) (bool, error) {
	return _ICS23.Contract.VerifyNonMembership(&_ICS23.CallOpts, spec, root, proof, key)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfbc2674d.
//
// Solidity: function verifyNonMembership(((bool,uint8,uint8,uint8,uint8,bytes),(uint256[],uint256,uint256,uint256,bytes,uint8),uint256,uint256) spec, bytes root, (bool,bytes,(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[]),(bool,bytes,bytes,(bool,uint8,uint8,uint8,uint8,bytes),(bool,uint8,bytes,bytes)[])) proof, bytes key) pure returns(bool)
func (_ICS23 *ICS23CallerSession) VerifyNonMembership(spec ICS23ProofSpec, root []byte, proof ICS23NonExistenceProof, key []byte) (bool, error) {
	return _ICS23.Contract.VerifyNonMembership(&_ICS23.CallOpts, spec, root, proof, key)
}
