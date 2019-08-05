import { proofs } from "./generated/codecimpl";
import { applyInner, applyLeaf } from "./ops";

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
export function calculateExistenceRoot(
  proof: proofs.IExistenceProof
): Uint8Array {
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

function ensureLeaf(leaf: proofs.ILeafOp, spec: proofs.ILeafOp): void {
  if (leaf.hash !== spec.hash) {
    throw new Error(`Unexpected hashOp: ${leaf.hash}`);
  }
  if (leaf.prehashKey !== spec.prehashKey) {
    throw new Error(`Unexpected prehashKey: ${leaf.prehashKey}`);
  }
  if (leaf.prehashValue !== spec.prehashValue) {
    throw new Error(`Unexpected prehashValue: ${leaf.prehashValue}`);
  }
  if (leaf.length !== spec.length) {
    throw new Error(`Unexpected length op: ${leaf.length}`);
  }
  ensurePrefix(leaf.prefix, spec.prefix);
}

function ensureInner(inner: proofs.IInnerOp, prefix?: Uint8Array | null): void {
  ensurePrefix(inner.prefix, prefix);
}

function ensurePrefix(
  check?: Uint8Array | null,
  prefix?: Uint8Array | null
): void {
  // no prefix supplied, means everything passes
  if (!prefix || prefix.length === 0) {
    return;
  }
  if (!check) {
    throw new Error(`Target bytes missing`);
  }
  ensureBytesEqual(prefix, check.slice(0, prefix.length));
}

// ensureBytesEqual throws an error if the arrays are different
function ensureBytesEqual(a: Uint8Array, b: Uint8Array): void {
  if (a.length !== b.length) {
    throw new Error(`Different lengths ${a.length} vs ${b.length}`);
  }
  for (let i = 0; i < a.length; i++) {
    if (a[i] !== b[i]) {
      throw new Error(`Arrays differ at index ${i}: ${a[i]} vs ${b[i]}`);
    }
  }
}
