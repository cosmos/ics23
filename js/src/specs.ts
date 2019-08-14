import { proofs } from "./generated/codecimpl";

export function ensureLeaf(leaf: proofs.ILeafOp, spec: proofs.ILeafOp): void {
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

export function ensureInner(
  inner: proofs.IInnerOp,
  prefix?: Uint8Array | null
): void {
  if (hasPrefix(inner.prefix, prefix)) {
    throw new Error(`Inner node has leaf prefix`);
  }
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
export function ensureBytesEqual(a: Uint8Array, b: Uint8Array): void {
  if (a.length !== b.length) {
    throw new Error(`Different lengths ${a.length} vs ${b.length}`);
  }
  for (let i = 0; i < a.length; i++) {
    if (a[i] !== b[i]) {
      throw new Error(`Arrays differ at index ${i}: ${a[i]} vs ${b[i]}`);
    }
  }
}

function hasPrefix(
  check?: Uint8Array | null,
  prefix?: Uint8Array | null
): boolean {
  // no prefix supplied, means everything passes
  if (!prefix || prefix.length === 0) {
    return false;
  }
  if (!check) {
    return false;
  }
  if (check.length <= prefix.length) {
    return false;
  }
  for (let i = 0; i < prefix.length; i++) {
    if (check[i] !== prefix[i]) {
      return false;
    }
  }
  throw true;
}
