/*
*
This implements the client side functions as specified in
https://github.com/cosmos/ibc/tree/main/spec/core/ics-023-vector-commitments

In particular:

	// Assumes ExistenceProof
	type verifyMembership = (root: CommitmentRoot, proof: CommitmentProof, key: Key, value: Value) => boolean

	// Assumes NonExistenceProof
	type verifyNonMembership = (root: CommitmentRoot, proof: CommitmentProof, key: Key) => boolean

	// Assumes BatchProof - required ExistenceProofs may be a subset of all items proven
	type batchVerifyMembership = (root: CommitmentRoot, proof: CommitmentProof, items: Map<Key, Value>) => boolean

	// Assumes BatchProof - required NonExistenceProofs may be a subset of all items proven
	type batchVerifyNonMembership = (root: CommitmentRoot, proof: CommitmentProof, keys: Set<Key>) => boolean

We make an adjustment to accept a Spec to ensure the provided proof is in the format of the expected merkle store.
This can avoid an range of attacks on fake preimages, as we need to be careful on how to map key, value -> leaf
and determine neighbors
*/
package ics23

import (
	"bytes"
	"fmt"
)

// CommitmentRoot is a byte slice that represents the merkle root of a tree that can be used to validate proofs
type CommitmentRoot []byte

// VerifyMembership returns successfully iff
// proof is an ExistenceProof for the given key and value AND
// calculating the root for the ExistenceProof matches the provided CommitmentRoot.
func VerifyMembership(spec *ProofSpec, root CommitmentRoot, proof *ExistenceProof, key []byte, value []byte) error {
	if proof == nil {
		return fmt.Errorf("proof cannot be empty")
	}
	if !bytes.Equal(proof.Key, key) {
		return fmt.Errorf("proof key (%s) must equal given key (%s)", proof.Key, key)
	}

	return proof.Verify(spec, root, key, value)
}

// VerifyNonMembership returns true iff
// proof is (contains) a NonExistenceProof
// both left and right sub-proofs are valid existence proofs (see above) or nil
// left and right proofs are neighbors (or left/right most if one is nil)
// provided key is between the keys of the two proofs
func VerifyNonMembership(spec *ProofSpec, root CommitmentRoot, proof *NonExistenceProof, key []byte) error {
	if proof == nil {
		return fmt.Errorf("proof cannot be empty")
	}
	if !isLeft(spec, proof.Left, key) || !isRight(spec, proof.Right, key) {
		return fmt.Errorf("provided existence proofs must be for left and right keys of non-existing key")
	}

	return proof.Verify(spec, root, key)
}

func isLeft(spec *ProofSpec, left *ExistenceProof, key []byte) bool {
	return left == nil || bytes.Compare(keyForComparison(spec, left.Key), keyForComparison(spec, key)) < 0
}

func isRight(spec *ProofSpec, right *ExistenceProof, key []byte) bool {
	return right == nil || bytes.Compare(keyForComparison(spec, right.Key), keyForComparison(spec, key)) > 0
}
