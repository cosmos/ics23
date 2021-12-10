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

type EmptyBranchTestStruct struct {
	Op     *InnerOp
	Spec   *InnerSpec
	IsTrue bool
	IsLeft bool
}

var InnerSpecWithEmptyChild = InnerSpec{
	ChildOrder:      []int32{0, 1},
	ChildSize:       32,
	MinPrefixLength: 1,
	MaxPrefixLength: 1,
	EmptyChild:      []byte("32_empty_child_placeholder_bytes"),
	Hash:            HashOp_SHA256,
}

func EmptyBranchTestData(t *testing.T) []EmptyBranchTestStruct {
	return []EmptyBranchTestStruct{
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: InnerSpecWithEmptyChild.EmptyChild,
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: true,
			IsLeft: false,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: append([]byte{1}, make([]byte, 32)...),
				Suffix: InnerSpecWithEmptyChild.EmptyChild,
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: true,
			IsLeft: false,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: make([]byte, 32),
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: false,
			IsLeft: false,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: nil,
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: false,
			IsLeft: false,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: append(InnerSpecWithEmptyChild.EmptyChild, []byte("xxxx")...),
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: false,
			IsLeft: false,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: nil,
			},
			Spec:   IavlSpec.InnerSpec,
			IsTrue: false,
			IsLeft: false,
		},

		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: append([]byte{1}, InnerSpecWithEmptyChild.EmptyChild...),
				Suffix: nil,
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: true,
			IsLeft: true,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: append([]byte{1}, InnerSpecWithEmptyChild.EmptyChild...),
				Suffix: make([]byte, 32),
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: true,
			IsLeft: true,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: append([]byte{1}, make([]byte, 32)...),
				Suffix: nil,
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: false,
			IsLeft: true,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: nil,
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: false,
			IsLeft: true,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: append(
					append([]byte{1}, InnerSpecWithEmptyChild.EmptyChild...),
					[]byte("xxxx")...),
				Suffix: nil,
			},
			Spec:   &InnerSpecWithEmptyChild,
			IsTrue: false,
			IsLeft: true,
		},
		EmptyBranchTestStruct{
			Op: &InnerOp{
				Prefix: []byte{1},
				Suffix: nil,
			},
			Spec:   IavlSpec.InnerSpec,
			IsTrue: false,
			IsLeft: true,
		},
	}
}
