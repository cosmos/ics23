import { proofs } from "./generated/codecimpl";

import { fromHex, toAscii } from "./helpers";
import { applyLeaf } from "./ops";

// "hash foobar": {
// 	op: &LeafOp{
// 		Hash: HashOp_SHA256,
// 		// no prehash, no length prefix
// 	},
// 	key:   []byte("foo"),
// 	value: []byte("bar"),
// 	// echo -n foobar | sha256sum
// 	expected: fromHex("c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"),
// },

describe("applyLeaf", () => {
  it("hashes foobar", () => {
    const op: proofs.ILeafOp = { hash: proofs.HashOp.SHA256 };
    const key = toAscii("foo");
    const value = toAscii("bar");
    const expected = fromHex(
      "c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"
    );
    expect(applyLeaf(op, key, value)).toEqual(expected);
  });

  it("hashes foobaz with sha-512", () => {
    const op: proofs.ILeafOp = { hash: proofs.HashOp.SHA512 };
    const key = toAscii("foo");
    const value = toAscii("baz");
    const expected = fromHex(
      "4f79f191298ec7461d60136c60f77c2ae8ddd85dbf6168bb925092d51bfb39b559219b39ae5385ba04946c87f64741385bef90578ea6fe6dac85dbf7ad3f79e1"
    );
    expect(applyLeaf(op, key, value)).toEqual(expected);
  });

  it("hashes foobar (different breakpoint)", () => {
    const op: proofs.ILeafOp = { hash: proofs.HashOp.SHA256 };
    const key = toAscii("f");
    const value = toAscii("oobar");
    const expected = fromHex(
      "c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"
    );
    expect(applyLeaf(op, key, value)).toEqual(expected);
  });

  it("hashes with length prefix", () => {
    const op: proofs.ILeafOp = {
      hash: proofs.HashOp.SHA256,
      length: proofs.LengthOp.VAR_PROTO
    };
    const key = toAscii("food");
    const value = toAscii("some longer text");
    const expected = fromHex(
      "b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"
    );
    expect(applyLeaf(op, key, value)).toEqual(expected);
  });
});
