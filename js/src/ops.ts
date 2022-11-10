import { ripemd160 } from "@noble/hashes/ripemd160";
import { sha256 } from "@noble/hashes/sha256";
import { sha512, sha512_256 } from "@noble/hashes/sha512";

import { ics23 } from "./generated/codecimpl";

export function applyLeaf(
  leaf: ics23.ILeafOp,
  key: Uint8Array,
  value: Uint8Array
): Uint8Array {
  if (key.length === 0) {
    throw new Error("Missing key");
  }
  if (value.length === 0) {
    throw new Error("Missing value");
  }
  const pkey = prepareLeafData(
    ensureHash(leaf.prehashKey),
    ensureLength(leaf.length),
    key
  );
  const pvalue = prepareLeafData(
    ensureHash(leaf.prehashValue),
    ensureLength(leaf.length),
    value
  );
  const data = new Uint8Array([
    ...ensureBytes(leaf.prefix),
    ...pkey,
    ...pvalue,
  ]);
  return doHash(ensureHash(leaf.hash), data);
}

export function applyInner(
  inner: ics23.IInnerOp,
  child: Uint8Array
): Uint8Array {
  if (child.length === 0) {
    throw new Error("Inner op needs child value");
  }
  const preimage = new Uint8Array([
    ...ensureBytes(inner.prefix),
    ...child,
    ...ensureBytes(inner.suffix),
  ]);
  return doHash(ensureHash(inner.hash), preimage);
}

function ensure<T>(maybe: T | undefined | null, value: T): T {
  return maybe === undefined || maybe === null ? value : maybe;
}

const ensureHash = (h: ics23.HashOp | null | undefined): ics23.HashOp =>
  ensure(h, ics23.HashOp.NO_HASH);
const ensureLength = (l: ics23.LengthOp | null | undefined): ics23.LengthOp =>
  ensure(l, ics23.LengthOp.NO_PREFIX);
const ensureBytes = (b: Uint8Array | null | undefined): Uint8Array =>
  ensure(b, new Uint8Array([]));

function prepareLeafData(
  hashOp: ics23.HashOp,
  lengthOp: ics23.LengthOp,
  data: Uint8Array
): Uint8Array {
  const h = doHashOrNoop(hashOp, data);
  return doLengthOp(lengthOp, h);
}

// doHashOrNoop will return the preimage untouched if hashOp == NONE,
// otherwise, perform doHash
function doHashOrNoop(hashOp: ics23.HashOp, preimage: Uint8Array): Uint8Array {
  if (hashOp === ics23.HashOp.NO_HASH) {
    return preimage;
  }
  return doHash(hashOp, preimage);
}

// doHash will preform the specified hash on the preimage.
// if hashOp == NONE, it will return an error (use doHashOrNoop if you want different behavior)
export function doHash(hashOp: ics23.HashOp, preimage: Uint8Array): Uint8Array {
  switch (hashOp) {
    case ics23.HashOp.SHA256:
      return sha256(preimage);
    case ics23.HashOp.SHA512:
      return sha512(preimage);
    case ics23.HashOp.RIPEMD160:
      return ripemd160(preimage);
    case ics23.HashOp.BITCOIN:
      return ripemd160(sha256(preimage));
    case ics23.HashOp.SHA512_256:
      return sha512_256(preimage);
  }
  throw new Error(`Unsupported hashop: ${hashOp}`);
}

// doLengthOp will calculate the proper prefix and return it prepended
//   doLengthOp(op, data) -> length(data) || data
function doLengthOp(lengthOp: ics23.LengthOp, data: Uint8Array): Uint8Array {
  switch (lengthOp) {
    case ics23.LengthOp.NO_PREFIX:
      return data;
    case ics23.LengthOp.VAR_PROTO:
      return new Uint8Array([...encodeVarintProto(data.length), ...data]);
    case ics23.LengthOp.REQUIRE_32_BYTES:
      if (data.length !== 32) {
        throw new Error(`Length is ${data.length}, not 32 bytes`);
      }
      return data;
    case ics23.LengthOp.REQUIRE_64_BYTES:
      if (data.length !== 64) {
        throw new Error(`Length is ${data.length}, not 64 bytes`);
      }
      return data;
    case ics23.LengthOp.FIXED32_LITTLE:
      return new Uint8Array([...encodeFixed32Le(data.length), ...data]);
    // TODO
    // case LengthOp_VAR_RLP:
    // case LengthOp_FIXED32_BIG:
    // case LengthOp_FIXED64_BIG:
    // case LengthOp_FIXED64_LITTLE:
  }
  throw new Error(`Unsupported lengthop: ${lengthOp}`);
}

function encodeVarintProto(n: number): Uint8Array {
  let enc: readonly number[] = [];
  let l = n;
  while (l >= 128) {
    const b = (l % 128) + 128;
    enc = [...enc, b];
    l = l / 128;
  }
  enc = [...enc, l];
  return new Uint8Array(enc);
}

function encodeFixed32Le(n: number): Uint8Array {
  const enc = new Uint8Array(4);
  let l = n;
  for (let i = enc.length; i > 0; i--) {
    enc[Math.abs(i - enc.length)] = l % 256;
    l = Math.floor(l / 256);
  }
  return enc;
}
