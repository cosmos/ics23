import { ics23 } from "./generated/codecimpl";
import { applyInner, applyLeaf } from "./ops";
import { doHash } from "./ops";
import {
  bytesEqual,
  ensureBytesBefore,
  ensureBytesEqual,
  ensureInner,
  ensureLeaf,
} from "./specs";

export const iavlSpec: ics23.IProofSpec = {
  leafSpec: {
    prefix: Uint8Array.from([0]),
    hash: ics23.HashOp.SHA256,
    prehashValue: ics23.HashOp.SHA256,
    prehashKey: ics23.HashOp.NO_HASH,
    length: ics23.LengthOp.VAR_PROTO,
  },
  innerSpec: {
    childOrder: [0, 1],
    minPrefixLength: 4,
    maxPrefixLength: 12,
    childSize: 33,
    hash: ics23.HashOp.SHA256,
  },
};

export const tendermintSpec: ics23.IProofSpec = {
  leafSpec: {
    prefix: Uint8Array.from([0]),
    hash: ics23.HashOp.SHA256,
    prehashValue: ics23.HashOp.SHA256,
    prehashKey: ics23.HashOp.NO_HASH,
    length: ics23.LengthOp.VAR_PROTO,
  },
  innerSpec: {
    childOrder: [0, 1],
    minPrefixLength: 1,
    maxPrefixLength: 1,
    childSize: 32,
    hash: ics23.HashOp.SHA256,
  },
};

export const smtSpec: ics23.IProofSpec = {
  leafSpec: {
    hash: ics23.HashOp.SHA256,
    prehashKey: ics23.HashOp.SHA256,
    prehashValue: ics23.HashOp.SHA256,
    length: ics23.LengthOp.NO_PREFIX,
    prefix: Uint8Array.from([0]),
  },
  innerSpec: {
    childOrder: [0, 1],
    childSize: 32,
    minPrefixLength: 1,
    maxPrefixLength: 1,
    emptyChild: new Uint8Array(32),
    hash: ics23.HashOp.SHA256,
  },
  maxDepth: 256,
  prehashKeyBeforeComparison: true,
};

export type CommitmentRoot = Uint8Array;

export function keyForComparison(
  spec: ics23.IProofSpec,
  key: Uint8Array
): Uint8Array {
  if (!spec.prehashKeyBeforeComparison) {
    return key;
  }

  return doHash(spec.leafSpec!.prehashKey!, key);
}

// verifyExistence will throw an error if the proof doesn't link key, value -> root
// or if it doesn't fulfill the spec
export function verifyExistence(
  proof: ics23.IExistenceProof,
  spec: ics23.IProofSpec,
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
  proof: ics23.INonExistenceProof,
  spec: ics23.IProofSpec,
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

  if (leftKey) {
    ensureBytesBefore(
      keyForComparison(spec, leftKey),
      keyForComparison(spec, key)
    );
  }
  if (rightKey) {
    ensureBytesBefore(
      keyForComparison(spec, key),
      keyForComparison(spec, rightKey)
    );
  }

  if (!spec.innerSpec) {
    throw new Error("no inner spec");
  }
  if (!leftKey) {
    ensureLeftMost(spec.innerSpec, proof.right!.path!);
  } else if (!rightKey) {
    ensureRightMost(spec.innerSpec, proof.left!.path!);
  } else {
    ensureLeftNeighbor(spec.innerSpec, proof.left!.path!, proof.right!.path!);
  }
  return;
}

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
export function calculateExistenceRoot(
  proof: ics23.IExistenceProof
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
  proof: ics23.IExistenceProof,
  spec: ics23.IProofSpec
): void {
  if (!proof.leaf) {
    throw new Error("Existence proof must start with a leaf operation");
  }
  if (!spec.leafSpec) {
    throw new Error("Spec must include leafSpec");
  }
  if (!spec.innerSpec) {
    throw new Error("Spec must include innerSpec");
  }
  ensureLeaf(proof.leaf, spec.leafSpec);

  const path = proof.path || [];
  if (spec.minDepth && path.length < spec.minDepth) {
    throw new Error(`Too few inner nodes ${path.length}`);
  }
  if (spec.maxDepth && path.length > spec.maxDepth) {
    throw new Error(`Too many inner nodes ${path.length}`);
  }
  for (const inner of path) {
    ensureInner(inner, spec.leafSpec.prefix, spec.innerSpec);
  }
}

function ensureLeftMost(
  spec: ics23.IInnerSpec,
  path: readonly ics23.IInnerOp[]
): void {
  const { minPrefix, maxPrefix, suffix } = getPadding(spec, 0);

  // ensure every step has a prefix and suffix defined to be leftmost
  for (const step of path) {
    if (!hasPadding(step, minPrefix, maxPrefix, suffix)) {
      throw new Error("Step not leftmost");
    }
  }
}

function ensureRightMost(
  spec: ics23.IInnerSpec,
  path: readonly ics23.IInnerOp[]
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

export function ensureLeftNeighbor(
  spec: ics23.IInnerSpec,
  left: readonly ics23.IInnerOp[],
  right: readonly ics23.IInnerOp[]
): void {
  const mutleft: ics23.IInnerOp[] = [...left];
  const mutright: ics23.IInnerOp[] = [...right];

  let topleft = mutleft.pop()!;
  let topright = mutright.pop()!;
  while (
    bytesEqual(topleft.prefix!, topright.prefix!) &&
    bytesEqual(topleft.suffix!, topright.suffix!)
  ) {
    topleft = mutleft.pop()!;
    topright = mutright.pop()!;
  }

  // now topleft and topright are the first divergent nodes
  // make sure they are left and right of each other
  if (!isLeftStep(spec, topleft, topright)) {
    throw new Error(`Not left neightbor at first divergent step`);
  }

  // make sure the paths are left and right most possibilities respectively
  ensureRightMost(spec, mutleft);
  ensureLeftMost(spec, mutright);
}

// isLeftStep assumes left and right have common parents
// checks if left is exactly one slot to the left of right
function isLeftStep(
  spec: ics23.IInnerSpec,
  left: ics23.IInnerOp,
  right: ics23.IInnerOp
): boolean {
  const leftidx = orderFromPadding(spec, left);
  const rightidx = orderFromPadding(spec, right);
  return rightidx === leftidx + 1;
}

function orderFromPadding(
  spec: ics23.IInnerSpec,
  inner: ics23.IInnerOp
): number {
  for (let branch = 0; branch < spec.childOrder!.length; branch++) {
    const { minPrefix, maxPrefix, suffix } = getPadding(spec, branch);
    if (hasPadding(inner, minPrefix, maxPrefix, suffix)) {
      return branch;
    }
  }
  throw new Error(`Cannot find any valid spacing for this node`);
}

function hasPadding(
  op: ics23.IInnerOp,
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
function getPadding(spec: ics23.IInnerSpec, branch: number): PaddingResult {
  const idx = getPosition(spec.childOrder!, branch);

  // count how many children are in the prefix
  const prefix = idx * spec.childSize!;
  const minPrefix = prefix + spec.minPrefixLength!;
  const maxPrefix = prefix + spec.maxPrefixLength!;

  // count how many children are in the suffix
  const suffix = (spec.childOrder!.length - 1 - idx) * spec.childSize!;
  return { minPrefix, maxPrefix, suffix };
}

function getPosition(order: readonly number[], branch: number): number {
  if (branch < 0 || branch >= order.length) {
    throw new Error(`Invalid branch: ${branch}`);
  }
  return order.findIndex((val) => val === branch);
}
