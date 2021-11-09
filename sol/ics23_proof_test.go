package ics23_sol

import (
	"bytes"
	"testing"

	ics23 "github.com/confio/ics23/go"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestExistenceProof(t *testing.T) {
	cases := ics23.ExistenceProofTestData()
	deployer, callers, client := setup(t, len(cases))
	_, _, entry, err := DeployProofUnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployProofUnitTest", err)
	}
	client.Commit()

	idx := 0
	for name, tc := range cases {
		testCaller := callers[idx]
		idx++
		t.Run(name, func(t *testing.T) {
			data := ExistenceProofData{
				Key:   tc.Proof.Key,
				Path:  translateInnerOps(tc.Proof.Path...),
				Leaf:  translateLeafOp(tc.Proof.Leaf),
				Value: tc.Proof.Value,
			}
			res, err := entry.CalculateRoot(&testCaller, data)
			// short-circuit with error case
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got none")
				return
			}
			if tc.IsErr == false && err != nil {
				t.Fatal(err)
			}
			if bytes.Equal(res, tc.Expected) == false {
				t.Errorf("Bad result: %s vs %s", hexutil.Encode(res), hexutil.Encode(tc.Expected))
			}
		})
	}
}

func TestCheckAgainstSpec(t *testing.T) {
	cases := ics23.CheckAgainstSpecTestData()
	deployer, callers, client := setup(t, len(cases))
	_, _, entry, err := DeployProofUnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployProofUnitTest", err)
	}
	client.Commit()

	idx := 0
	for name, tc := range cases {
		testCaller := callers[idx]
		idx++
		t.Run(name, func(t *testing.T) {
			specs := translateProofSpec(tc.Spec)
			proof := translateExistenceProof(tc.Proof)
			err := entry.CheckAgainstSpec(&testCaller, proof, specs)
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got nil")
			} else if tc.IsErr == false && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}
