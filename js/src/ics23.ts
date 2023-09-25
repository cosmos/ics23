import { decompress } from "./compress";
import { ics23 } from "./generated/codecimpl";
import { CommitmentRoot, verifyExistence, verifyNonExistence } from "./proofs";
import { keyForComparison } from "./proofs";
import { bytesBefore, bytesEqual } from "./specs";
/*
This implements the client side functions as specified in
https://github.com/cosmos/ics/tree/master/spec/ics-023-vector-commitments

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

/**
 * verifyMembership ensures proof is (contains) a valid existence proof for the given
 */
export function verifyMembership(
  proof: ics23.ICommitmentProof,
  spec: ics23.IProofSpec,
  root: CommitmentRoot,
  key: Uint8Array,
  value: Uint8Array
): boolean {
  const norm = decompress(proof);
  const exist = getExistForKey(norm, key);
  if (!exist) {
    return false;
  }
  try {
    verifyExistence(exist, spec, root, key, value);
    return true;
  } catch {
    return false;
  }
}

/**
 * verifyNonMembership ensures proof is (contains) a valid non-existence proof for the given key
 */
export function verifyNonMembership(
  proof: ics23.ICommitmentProof,
  spec: ics23.IProofSpec,
  root: CommitmentRoot,
  key: Uint8Array
): boolean {
  const norm = decompress(proof);
  const nonexist = getNonExistForKey(spec, norm, key);
  if (!nonexist) {
    return false;
  }
  try {
    verifyNonExistence(nonexist, spec, root, key);
    return true;
  } catch {
    return false;
  }
}

/**
 * batchVerifyMembership ensures proof is (contains) a valid existence proof for the given
 */
export function batchVerifyMembership(
  proof: ics23.ICommitmentProof,
  spec: ics23.IProofSpec,
  root: CommitmentRoot,
  items: Map<Uint8Array, Uint8Array>
): boolean {
  const norm = decompress(proof);
  for (const [key, value] of items.entries()) {
    if (!verifyMembership(norm, spec, root, key, value)) {
      return false;
    }
  }
  return true;
}

/**
 * batchVerifyNonMembership ensures proof is (contains) a valid existence proof for the given
 */
export function batchVerifyNonMembership(
  proof: ics23.ICommitmentProof,
  spec: ics23.IProofSpec,
  root: CommitmentRoot,
  keys: readonly Uint8Array[]
): boolean {
  const norm = decompress(proof);
  for (const key of keys) {
    if (!verifyNonMembership(norm, spec, root, key)) {
      return false;
    }
  }
  return true;
}

function getExistForKey(
  proof: ics23.ICommitmentProof,
  key: Uint8Array
): ics23.IExistenceProof | undefined | null {
  const match = (p: ics23.IExistenceProof | null | undefined): boolean =>
    !!p && bytesEqual(key, p.key!);
  if (match(proof.exist)) {
    return proof.exist!;
  } else if (proof.batch) {
    return proof.batch.entries!.map((x) => x.exist || null).find(match);
  }
  return undefined;
}

function getNonExistForKey(
  spec: ics23.IProofSpec,
  proof: ics23.ICommitmentProof,
  key: Uint8Array
): ics23.INonExistenceProof | undefined | null {
  const match = (p: ics23.INonExistenceProof | null | undefined): boolean => {
    return (
      !!p &&
      (!p.left ||
        bytesBefore(
          keyForComparison(spec, p.left.key!),
          keyForComparison(spec, key)
        )) &&
      (!p.right ||
        bytesBefore(
          keyForComparison(spec, key),
          keyForComparison(spec, p.right.key!)
        ))
    );
  };
  if (match(proof.nonexist)) {
    return proof.nonexist!;
  } else if (proof.batch) {
    return proof.batch.entries!.map((x) => x.nonexist || null).find(match);
  }
  return undefined;
}
