import shajs from "sha.js";

import { proofs } from "./generated/codecimpl";

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
  const pkey = prepareLeafData(leaf.prehashKey, leaf.length, key);
  const pvalue = prepareLeafData(leaf.prehashValue, leaf.length, value);
  const prefix = leaf.prefix || new Uint8Array([]);
  const data = new Uint8Array([...prefix, ...pkey, ...pvalue]);
  return doHash(leaf.hash, data);
}

function prepareLeafData(
  hashOp: proofs.HashOp | undefined | null,
  lengthOp: proofs.LengthOp | undefined | null,
  data: Uint8Array
): Uint8Array {
  const h = doHashOrNoop(hashOp, data);
  return doLengthOp(lengthOp, h);
}

// doHashOrNoop will return the preimage untouched if hashOp == NONE,
// otherwise, perform doHash
function doHashOrNoop(
  hashOp: proofs.HashOp | undefined | null,
  preimage: Uint8Array
): Uint8Array {
  if (
    hashOp === undefined ||
    hashOp === null ||
    hashOp === proofs.HashOp.NO_HASH
  ) {
    return preimage;
  }
  return doHash(hashOp, preimage);
}

// doHash will preform the specified hash on the preimage.
// if hashOp == NONE, it will return an error (use doHashOrNoop if you want different behavior)
function doHash(
  hashOp: proofs.HashOp | undefined | null,
  preimage: Uint8Array
): Uint8Array {
  switch (hashOp) {
    case proofs.HashOp.SHA256:
      return new Uint8Array(
        shajs("sha256")
          .update(preimage)
          .digest()
      );
    case proofs.HashOp.SHA512:
      return new Uint8Array(
        shajs("sha512")
          .update(preimage)
          .digest()
      );
  }
  throw new Error(`Unsupported hashop: ${hashOp}`);
}

// doLengthOp will calculate the proper prefix and return it prepended
//   doLengthOp(op, data) -> length(data) || data
function doLengthOp(
  lengthOp: proofs.LengthOp | undefined | null,
  data: Uint8Array
): Uint8Array {
  switch (lengthOp) {
    case null:
    case undefined:
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

// TODO
function encodeVarintProto(l: number): Uint8Array {
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
