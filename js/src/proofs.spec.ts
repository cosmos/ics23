import { ics23 } from "./generated/codecimpl";

import { fromHex, toAscii } from "./helpers";
import { calculateExistenceRoot } from "./proofs";

describe("calculateExistenceRoot", () => {
  it("must have at least one step", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar")
    };
    expect(() => calculateExistenceRoot(proof)).toThrow();
  });
  it("executes one leaf step", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("food"),
      value: toAscii("some longer text"),
      leaf: {
        hash: ics23.HashOp.SHA256,
        length: ics23.LengthOp.VAR_PROTO
      }
    };
    const expected = fromHex(
      "b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"
    );
    expect(calculateExistenceRoot(proof)).toEqual(expected);
  });
  it("cannot execute inner first", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("food"),
      value: toAscii("some longer text"),
      path: [
        {
          hash: ics23.HashOp.SHA256,
          prefix: fromHex("deadbeef00cafe00")
        }
      ]
    };
    expect(() => calculateExistenceRoot(proof)).toThrow();
  });
  it("can execute leaf then inner", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("food"),
      value: toAscii("some longer text"),
      leaf: {
        hash: ics23.HashOp.SHA256,
        length: ics23.LengthOp.VAR_PROTO
      },
      // output: b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265
      path: [
        {
          hash: ics23.HashOp.SHA256,
          prefix: fromHex("deadbeef00cafe00")
        }
        // echo -n deadbeef00cafe00b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265 | xxd -r -p | sha256sum
      ]
    };
    const expected = fromHex(
      "836ea236a6902a665c2a004c920364f24cad52ded20b1e4f22c3179bfe25b2a9"
    );
    expect(calculateExistenceRoot(proof)).toEqual(expected);
  });
});
