import ripemd160 from "ripemd160";
import shajs from "sha.js";

import { proofs } from "./generated/codecimpl";
import { toHex } from "./helpers";

export function applyLeaf(
  leaf: proofs.ILeafOp,
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
    ...pvalue
  ]);
  return doHash(ensureHash(leaf.hash), data);
}

export function applyInner(
  inner: proofs.IInnerOp,
  child: Uint8Array
): Uint8Array {
  if (child.length === 0) {
    throw new Error("Inner op needs child value");
  }
  const preimage = new Uint8Array([
    ...ensureBytes(inner.prefix),
    ...child,
    ...ensureBytes(inner.suffix)
  ]);
  return doHash(ensureHash(inner.hash), preimage);
}

function ensure<T>(maybe: T | undefined | null, value: T): T {
  return maybe === undefined || maybe === null ? value : maybe;
}

const ensureHash = (h: proofs.HashOp | null | undefined) =>
  ensure(h, proofs.HashOp.NO_HASH);
const ensureLength = (l: proofs.LengthOp | null | undefined) =>
  ensure(l, proofs.LengthOp.NO_PREFIX);
const ensureBytes = (b: Uint8Array | null | undefined) =>
  ensure(b, new Uint8Array([]));

function prepareLeafData(
  hashOp: proofs.HashOp,
  lengthOp: proofs.LengthOp,
  data: Uint8Array
): Uint8Array {
  const h = doHashOrNoop(hashOp, data);
  return doLengthOp(lengthOp, h);
}

// doHashOrNoop will return the preimage untouched if hashOp == NONE,
// otherwise, perform doHash
function doHashOrNoop(hashOp: proofs.HashOp, preimage: Uint8Array): Uint8Array {
  if (hashOp === proofs.HashOp.NO_HASH) {
    return preimage;
  }
  return doHash(hashOp, preimage);
}

function rp160(preimage: Uint8Array): Uint8Array {
  // this is a bit tricky to work with besides buffer
  return new Uint8Array(
    new ripemd160().update(toHex(preimage), "hex" as any).digest()
  );
}

function s256(preimage: Uint8Array): Uint8Array {
  return new Uint8Array(
    shajs("sha256")
      .update(preimage)
      .digest()
  );
}

// doHash will preform the specified hash on the preimage.
// if hashOp == NONE, it will return an error (use doHashOrNoop if you want different behavior)
export function doHash(
  hashOp: proofs.HashOp,
  preimage: Uint8Array
): Uint8Array {
  switch (hashOp) {
    case proofs.HashOp.SHA256:
      return s256(preimage);
    case proofs.HashOp.SHA512:
      return new Uint8Array(
        shajs("sha512")
          .update(preimage)
          .digest()
      );
    case proofs.HashOp.RIPEMD160:
      // this requires string or Buffer....
      return rp160(preimage);
    case proofs.HashOp.BITCOIN:
      return rp160(s256(preimage));
  }
  throw new Error(`Unsupported hashop: ${hashOp}`);
}

// doLengthOp will calculate the proper prefix and return it prepended
//   doLengthOp(op, data) -> length(data) || data
function doLengthOp(lengthOp: proofs.LengthOp, data: Uint8Array): Uint8Array {
  switch (lengthOp) {
    case proofs.LengthOp.NO_PREFIX:
      return data;
    case proofs.LengthOp.VAR_PROTO:
      return new Uint8Array([...encodeVarintProto(data.length), ...data]);
    case proofs.LengthOp.REQUIRE_32_BYTES:
      if (data.length !== 32) {
        throw new Error(`Length is ${data.length}, not 32 bytes`);
      }
      return data;
    case proofs.LengthOp.REQUIRE_64_BYTES:
      if (data.length !== 64) {
        throw new Error(`Length is ${data.length}, not 64 bytes`);
      }
      return data;
    // TODO
    // case LengthOp_VAR_RLP:
    // case LengthOp_FIXED32_BIG:
    // case LengthOp_FIXED64_BIG:
    // case LengthOp_FIXED32_LITTLE:
    // case LengthOp_FIXED64_LITTLE:
  }
  throw new Error(`Unsupported lengthop: ${lengthOp}`);
}

function encodeVarintProto(l: number): Uint8Array {
  // TODO: handle numbers > 127
  return new Uint8Array([l]);

  // // avoid multiple allocs for normal case
  // res := make([]byte, 0, 8)
  // for l >= 1<<7 {
  // 	res = append(res, uint8(l&0x7f|0x80))
  // 	l >>= 7
  // }
  // res = append(res, uint8(l))
  // return res
}
