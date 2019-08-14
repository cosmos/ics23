import { proofs } from "./generated/codecimpl";
import { applyInner, applyLeaf } from "./ops";
import {
  ensureBytesBefore,
  ensureBytesEqual,
  ensureInner,
  ensureLeaf
} from "./specs";

export const IavlSpec: proofs.IProofSpec = {
  leafSpec: {
    prefix: Uint8Array.from([0]),
    hash: proofs.HashOp.SHA256,
    prehashValue: proofs.HashOp.SHA256,
    prehashKey: proofs.HashOp.NO_HASH,
    length: proofs.LengthOp.VAR_PROTO
  }
};

export const TendermintSpec: proofs.IProofSpec = {
  leafSpec: {
    prefix: Uint8Array.from([0]),
    hash: proofs.HashOp.SHA256,
    prehashValue: proofs.HashOp.SHA256,
    prehashKey: proofs.HashOp.NO_HASH,
    length: proofs.LengthOp.VAR_PROTO
  }
};

export type CommitmentRoot = Uint8Array;

// verifyExistence will throw an error if the proof doesn't link key, value -> root
// or if it doesn't fulfill the spec
export function verifyExistence(
  proof: proofs.IExistenceProof,
  spec: proofs.IProofSpec,
  root: CommitmentRoot,
  key: Uint8Array,
  value: Uint8Array
): void {
  ensureSpec(proof, spec);
  const calc = calculateExistenceRoot(proof);
  ensureBytesEqual(calc, root);
  ensureBytesEqual(key, proof.key!);
  ensureBytesEqual(value, proof.value!);
}

// Verify does all checks to ensure the proof has valid non-existence proofs,
// and they ensure the given key is not in the CommitmentState,
// throwing an error if there is an issue
export function verifyNonExistence(
  proof: proofs.INonExistenceProof,
  spec: proofs.IProofSpec,
  root: CommitmentRoot,
  key: Uint8Array
): void {
  let leftKey: Uint8Array | undefined;
  let rightKey: Uint8Array | undefined;

  if (proof.left) {
    verifyExistence(proof.left, spec, root, proof.left.key!, proof.left.value!);
    leftKey = proof.left.key!;
  }
  if (proof.right) {
    verifyExistence(
      proof.right,
      spec,
      root,
      proof.right.key!,
      proof.right.value!
    );
    rightKey = proof.right.key!;
  }

  if (!leftKey && !rightKey) {
    throw new Error("neither left nor right proof defined");
  }

  if (!!leftKey) {
    ensureBytesBefore(leftKey, key);
  }
  if (!!rightKey) {
    ensureBytesBefore(key, rightKey);
  }

  if (!leftKey) {
    // TODO: ensure right proof is left most
    return;
  }

  if (!rightKey) {
    // TODO: ensure left proof is right most
    return;
  }

  // TODO: ensure left and right are neighbors
  return;
}

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
export function calculateExistenceRoot(
  proof: proofs.IExistenceProof
): CommitmentRoot {
  if (!proof.key || !proof.value) {
    throw new Error("Existence proof needs key and value set");
  }
  if (!proof.leaf) {
    throw new Error("Existence proof must start with a leaf operation");
  }
  const path = proof.path || [];

  let res = applyLeaf(proof.leaf, proof.key, proof.value);
  for (const inner of path) {
    res = applyInner(inner, res);
  }
  return res;
}

// ensureSpec throws an Error if proof doesn't fulfill spec
export function ensureSpec(
  proof: proofs.IExistenceProof,
  spec: proofs.IProofSpec
): void {
  if (!proof.leaf) {
    throw new Error("Existence proof must start with a leaf operation");
  }
  if (!spec.leafSpec) {
    throw new Error("Spec must include leafSpec");
  }
  ensureLeaf(proof.leaf, spec.leafSpec);
  const path = proof.path || [];
  for (const inner of path) {
    ensureInner(inner, spec.leafSpec.prefix);
  }
}
