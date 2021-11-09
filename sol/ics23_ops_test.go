package ics23_sol

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"

	ics23 "github.com/confio/ics23/go"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestLeafOp(t *testing.T) {
	cases := ics23.LeafOpTestData()
	deployer, callers, client := setup(t, len(cases))
	_, _, entry, err := DeployOpsUnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployIcs23ExistenceProof", err)
	}
	client.Commit()

	idx := 0
	for name, tc := range cases {
		//WARNING: sha-512 in solidity not yet supported
		if strings.Contains(name, "sha-512") == true {
			continue
		}
		testCaller := callers[idx]
		idx++
		t.Run(name, func(t *testing.T) {
			leafOp := translateLeafOp(tc.Op)
			res, err := entry.ApplyOp(&testCaller, leafOp, tc.Key, tc.Value)
			// short-circuit with error case
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got none")
				return
			}
			if tc.IsErr == false && err != nil {
				t.Fatal(err, tc.Op)
			}
			if bytes.Equal(res, tc.Expected) == false {
				t.Errorf("Bad result: %s vs %s", hexutil.Encode(res), hexutil.Encode(tc.Expected))
			}
		})
	}
}

func TestCheckLeaf(t *testing.T) {
	cases := ics23.CheckLeafTestData()
	deployer, callers, client := setup(t, len(cases))
	_, _, entry, err := DeployOpsUnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployOpsUnitTest", err)
	}
	client.Commit()

	idx := 0
	for name, tc := range cases {
		testCaller := callers[idx]
		idx++
		t.Run(name, func(t *testing.T) {
			proof := ProofSpecData{LeafSpec: translateLeafOp(tc.Spec)}
			leafOp := translateLeafOp(tc.Leaf)
			err := entry.CheckAgainstSpec(&testCaller, leafOp, proof)
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got nil")
			} else if tc.IsErr == false && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}

func TestInnerOp(t *testing.T) {
	cases := ics23.InnerOpTestData()
	deployer, callers, client := setup(t, len(cases))
	_, _, entry, err := DeployOpsUnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployOpsUnitTest", err)
	}
	client.Commit()

	idx := 0
	for name, tc := range cases {
		testCaller := callers[idx]
		idx++
		t.Run(name, func(t *testing.T) {
			innerOp := translateInnerOp(tc.Op)
			res, err := entry.ApplyOp0(&testCaller, innerOp, tc.Child)
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

func TestDoHash(t *testing.T) {
	cases := ics23.DoHashTestData()
	deployer, callers, client := setup(t, len(cases))
	_, _, entry, err := DeployOpsUnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployOpsUnitTest", err)
	}
	client.Commit()

	idx := 0
	for name, tc := range cases {
		if strings.Contains(name, "512") {
			continue
		}
		testCaller := callers[idx]
		idx++
		t.Run(name, func(t *testing.T) {
			hashOp := translateHashOp(tc.HashOp)
			res, err := entry.DoHash(&testCaller, hashOp, []byte(tc.Preimage))
			if err != nil {
				t.Fatal(err)
			}
			hexRes := hex.EncodeToString(res)
			if hexRes != tc.ExpectedHash {
				t.Fatalf("Expected %s got %s", tc.ExpectedHash, hexRes)
			}
		})
	}
}
