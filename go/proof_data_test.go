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
