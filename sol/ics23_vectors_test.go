package ics23_sol

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"

	ics23 "github.com/confio/ics23/go"
)

func TestCompare(t *testing.T) {
	deployer, callers, client := setup(t, 100)
	_, _, opsEntry, err := DeployOpsUnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployOpsUnitTest", err)
	}
	client.Commit()

	for idx := 0; idx < 100; idx++ {
		szA := rand.Intn(100)
		szB := rand.Intn(100)
		a := make([]byte, szA)
		b := make([]byte, szB)
		rand.Read(a)
		rand.Read(b)
		goresult := bytes.Compare(a, b)
		solresult, err := opsEntry.Compare(&callers[idx], a, b)
		if err != nil {
			t.Fatal(err)
		}
		if int64(goresult) != solresult.Int64() {
			t.Fatal(a, b, goresult, solresult)
		}
	}
}

func TestVectors(t *testing.T) {
	cases := ics23.TestVectorsData()
	deployer, callers, client := setup(t, len(cases))
	_, _, ics23UnitTest, err := DeployICS23UnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployICS23UnitTest", err)
	}
	_, _, ics23ProofUnitTest, err := DeployProofUnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployProofUnit;Test", err)
	}
	client.Commit()

	idx := 0
	for _, tc := range cases {
		testCaller := callers[idx]
		idx++
		name := fmt.Sprintf("%s/%s", tc.Dir, tc.Filename)
		t.Run(name, func(t *testing.T) {
			proof, ref := ics23.LoadFile(t, tc.Dir, tc.Filename)
			proofData := translateCommitmentProof(proof)
			calculatedRoot, err := ics23ProofUnitTest.CalculateRoot0(&testCaller, proofData)
			if err != nil {
				t.Fatal("proof.Calculate() returned error", err)
			}
			if !bytes.Equal(ref.RootHash, calculatedRoot) {
				t.Fatalf("calculated root: %X did not match expected root: %X", calculatedRoot, ref.RootHash)
			}
			spec := translateProofSpec(tc.Spec)
			if ref.Value == nil {
				valid := ics23UnitTest.VerifyNonMembership(&testCaller, spec, ref.RootHash, proofData, ref.Key)
				if valid != nil {
					t.Fatal("Invalid NonMembership proof:", valid)
				}
			} else {
				valid := ics23UnitTest.VerifyMembership(&testCaller, spec, ref.RootHash, proofData, ref.Key, ref.Value)
				if valid != nil {
					t.Fatal("Invalid Membership proof:", valid)
				}
			}
		})
	}
}

func TestBatchVectors(t *testing.T) {
	cases := ics23.TestBatchVectorsData(t)
	deployer, callers, client := setup(t, len(cases))
	_, _, ics23UnitTest, err := DeployICS23UnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployICS23UnitTest", err)
	}
	client.Commit()

	idx := 0
	for name, tc := range cases {
		testCaller := callers[idx]
		idx++

		t.Run(name, func(t *testing.T) {
			// try one proof
			spec := translateProofSpec(tc.Spec)
			proof := translateCommitmentProof(tc.Proof)
			if tc.Ref.Value == nil {
				// non-existence
				verifyError := ics23UnitTest.VerifyNonMembership(&testCaller, spec, tc.Ref.RootHash, proof, tc.Ref.Key)
				valid := verifyError == nil
				//valid := VerifyNonMembership(tc.Spec, tc.Ref.RootHash, tc.Proof, tc.Ref.Key)
				if valid == tc.Invalid {
					t.Fatalf("Expected proof validity: %t, error: %v", !tc.Invalid, verifyError)
				}
				keys := [][]byte{tc.Ref.Key}
				verifyError = ics23UnitTest.BatchVerifyNonMembership(&testCaller, spec, tc.Ref.RootHash, proof, keys)
				valid = verifyError == nil
				if valid == tc.Invalid {
					t.Fatalf("Expected batch proof validity: %t, error: %v", !tc.Invalid, verifyError)
				}
			} else {
				verifyError := ics23UnitTest.VerifyMembership(&testCaller, spec, tc.Ref.RootHash, proof, tc.Ref.Key, tc.Ref.Value)
				valid := verifyError == nil
				if valid == tc.Invalid {
					t.Fatalf("Expected proof validity: %t, error: %v", !tc.Invalid, verifyError)
				}
				items := []Ics23BatchItem{{Key: tc.Ref.Key, Value: tc.Ref.Value}}
				verifyError = ics23UnitTest.BatchVerifyMembership(&testCaller, spec, tc.Ref.RootHash, proof, items)
				valid = verifyError == nil
				if valid == tc.Invalid {
					t.Fatalf("Expected batch proof validity: %t, error: %v", !tc.Invalid, verifyError)
				}
			}
		})
	}
}

func TestDecompressBatchVectors(t *testing.T) {
	cases := ics23.TestDecompressBatchVectorsData(t)
	deployer, callers, client := setup(t, len(cases))
	_, _, entry, err := DeployCompressUnitTest(deployer, client)
	if err != nil {
		t.Fatal("DeployICS23UnitTest", err)
	}
	client.Commit()

	idx := 0
	for name, tc := range cases {
		testCaller := callers[idx]
		idx++
		t.Run(name, func(t *testing.T) {
			if err != nil {
				t.Fatalf("Marshal batch %v", err)
			}
			proof := translateCommitmentProof(tc)
			decomp, err := entry.Decompress(&testCaller, proof)
			if err != nil {
                t.Fatal("Decompression error:", err)
			}
			if proofsEqual(decomp, proof) {
				t.Fatalf("Decompression is a no-op")
			}
		})
	}
}
