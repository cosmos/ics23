package ics23

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type ExistenceProofTestStruct struct {
	Proof    *ExistenceProof
	IsErr    bool
	Expected []byte
}

func ExistenceProofTestData(t *testing.T) map[string]ExistenceProofTestStruct {
	fname := filepath.Join("..", "testdata", "TestExistenceProofData.json")
	ffile, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	var cases map[string]ExistenceProofTestStruct
	jsonDecoder := json.NewDecoder(ffile)
	err = jsonDecoder.Decode(&cases)
	if err != nil {
		t.Fatal(err)
	}
	return cases
}

type CheckLeafTestStruct struct {
	Leaf  *LeafOp
	Spec  *LeafOp
	IsErr bool
}

func CheckLeafTestData(t *testing.T) map[string]CheckLeafTestStruct {
	fname := filepath.Join("..", "testdata", "TestCheckLeafData.json")
	ffile, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	var cases map[string]CheckLeafTestStruct
	jsonDecoder := json.NewDecoder(ffile)
	err = jsonDecoder.Decode(&cases)
	if err != nil {
		t.Fatal(err)
	}
	return cases
}

type CheckAgainstSpecTestStruct struct {
	Proof *ExistenceProof
	Spec  *ProofSpec
	IsErr bool
}

func CheckAgainstSpecTestData(t *testing.T) map[string]CheckAgainstSpecTestStruct {
	fname := filepath.Join("..", "testdata", "TestCheckAgainstSpecData.json")
	ffile, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	var cases map[string]CheckAgainstSpecTestStruct
	jsonDecoder := json.NewDecoder(ffile)
	err = jsonDecoder.Decode(&cases)
	if err != nil {
		t.Fatal(err)
	}
	return cases
}

var SpecWithEmptyChild = &ProofSpec{
	LeafSpec: &LeafOp{
		Prefix:       []byte{0},
		Hash:         HashOp_SHA256,
		PrehashValue: HashOp_SHA256,
	},
	InnerSpec: &InnerSpec{
		ChildOrder:      []int32{0, 1},
		ChildSize:       32,
		MinPrefixLength: 1,
		MaxPrefixLength: 1,
		EmptyChild:      []byte("32_empty_child_placeholder_bytes"),
		Hash:            HashOp_SHA256,
	},
}

type EmptyBranchTestStruct struct {
	Op      *InnerOp
	Spec    *ProofSpec
	IsLeft  bool
	IsRight bool
}

func EmptyBranchTestData(t *testing.T) []EmptyBranchTestStruct {
	emptyChild := SpecWithEmptyChild.InnerSpec.EmptyChild

	return []EmptyBranchTestStruct{
		{
			Op: &InnerOp{
				Prefix: append([]byte{1}, emptyChild...),
				Suffix: nil,
				Hash:   SpecWithEmptyChild.InnerSpec.Hash,
			},
			Spec:    SpecWithEmptyChild,
			IsLeft:  true,
			IsRight: false,
		},
		{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: emptyChild,
				Hash:   SpecWithEmptyChild.InnerSpec.Hash,
			},
			Spec:    SpecWithEmptyChild,
			IsLeft:  false,
			IsRight: true,
		},
		// non-empty cases
		{
			Op: &InnerOp{
				Prefix: append([]byte{1}, make([]byte, 32)...),
				Suffix: nil,
				Hash:   SpecWithEmptyChild.InnerSpec.Hash,
			},
			Spec:    SpecWithEmptyChild,
			IsLeft:  false,
			IsRight: false,
		},
		{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: make([]byte, 32),
				Hash:   SpecWithEmptyChild.InnerSpec.Hash,
			},
			Spec:    SpecWithEmptyChild,
			IsLeft:  false,
			IsRight: false,
		},
		{
			Op: &InnerOp{
				Prefix: append(append([]byte{1}, emptyChild[0:28]...), []byte("xxxx")...),
				Suffix: nil,
				Hash:   SpecWithEmptyChild.InnerSpec.Hash,
			},
			Spec:    SpecWithEmptyChild,
			IsLeft:  false,
			IsRight: false,
		},
		{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: append(append([]byte(nil), emptyChild[0:28]...), []byte("xxxx")...),
				Hash:   SpecWithEmptyChild.InnerSpec.Hash,
			},
			Spec:    SpecWithEmptyChild,
			IsLeft:  false,
			IsRight: false,
		},
		// some cases using a spec with no empty child
		{
			Op: &InnerOp{
				Prefix: append([]byte{1}, make([]byte, 32)...),
				Suffix: nil,
				Hash:   TendermintSpec.InnerSpec.Hash,
			},
			Spec:    TendermintSpec,
			IsLeft:  false,
			IsRight: false,
		},
		{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: make([]byte, 32),
				Hash:   TendermintSpec.InnerSpec.Hash,
			},
			Spec:    TendermintSpec,
			IsLeft:  false,
			IsRight: false,
		},
	}
}
