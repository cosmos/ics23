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
  },
  innerSpec: {
    childOrder: [0, 1],
    minPrefixLength: 4,
    maxPrefixLength: 12,
    childSize: 33
  }
};

export const TendermintSpec: proofs.IProofSpec = {
  leafSpec: {
    prefix: Uint8Array.from([0]),
    hash: proofs.HashOp.SHA256,
    prehashValue: proofs.HashOp.SHA256,
    prehashKey: proofs.HashOp.NO_HASH,
    length: proofs.LengthOp.VAR_PROTO
  },
  innerSpec: {
    childOrder: [0, 1],
    minPrefixLength: 1,
    maxPrefixLength: 1,
    childSize: 32
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

  if (!spec.innerSpec) {
    throw new Error("");
  }
  if (!leftKey) {
    ensureLeftMost(spec.innerSpec, proof.right!.path!);
  } else if (!rightKey) {
    ensureRightMost(spec.innerSpec, proof.left!.path!);
  } else {
    throw new Error("Unimplemented");
    // TODO: ensure left and right are neighbors
  }
  return;
}

export function ensureLeftMost(
  spec: proofs.IInnerSpec,
  path: ReadonlyArray<proofs.IInnerOp>
): void {
  const { minPrefix, maxPrefix, suffix } = getPadding(spec, 0);

  // ensure every step has a prefix and suffix defined to be leftmost
  for (const step of path) {
    if (!hasPadding(step, minPrefix, maxPrefix, suffix)) {
      throw new Error("Step not leftmost");
    }
  }
}

export function ensureRightMost(
  spec: proofs.IInnerSpec,
  path: ReadonlyArray<proofs.IInnerOp>
): void {
  const len = spec.childOrder!.length - 1;
  const { minPrefix, maxPrefix, suffix } = getPadding(spec, len);

  // ensure every step has a prefix and suffix defined to be leftmost
  for (const step of path) {
    if (!hasPadding(step, minPrefix, maxPrefix, suffix)) {
      throw new Error("Step not leftmost");
    }
  }
}

function hasPadding(
  op: proofs.IInnerOp,
  minPrefix: number,
  maxPrefix: number,
  suffix: number
): boolean {
  if ((op.prefix || []).length < minPrefix) {
    return false;
  }
  if ((op.prefix || []).length > maxPrefix) {
    return false;
  }
  return (op.suffix || []).length === suffix;
}

interface PaddingResult {
  readonly minPrefix: number;
  readonly maxPrefix: number;
  readonly suffix: number;
}
function getPadding(spec: proofs.IInnerSpec, branch: number): PaddingResult {
  const idx = getPosition(spec.childOrder!, branch);

  // count how many children are in the prefix
  const prefix = idx * spec.childSize!;
  const minPrefix = prefix + spec.minPrefixLength!;
  const maxPrefix = prefix + spec.maxPrefixLength!;

  // count how many children are in the suffix
  const suffix = (spec.childOrder!.length - 1 - idx) * spec.childSize!;
  return { minPrefix, maxPrefix, suffix };
}

function getPosition(order: ReadonlyArray<number>, branch: number): number {
  if (branch < 0 || branch >= order.length) {
    throw new Error(`Invalid branch: ${branch}`);
  }
  return order.findIndex(val => val === branch);
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
