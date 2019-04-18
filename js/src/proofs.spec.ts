import { proofs } from "./generated/codecimpl";

import { fromHex, toAscii } from "./helpers";
import { calculateExistenceRoot } from "./proofs";

describe("calculateExistenceRoot", () => {
  it("must have at least one step", () => {
    const proof: proofs.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar")
    };
    expect(() => calculateExistenceRoot(proof)).toThrow();
  });
  it("executes one leaf step", () => {
    const proof: proofs.IExistenceProof = {
      key: toAscii("food"),
      value: toAscii("some longer text"),
      steps: [
        {
          leaf: {
            hash: proofs.HashOp.SHA256,
            length: proofs.LengthOp.VAR_PROTO
          }
        }
      ]
    };
    const expected = fromHex(
      "b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"
    );
    expect(calculateExistenceRoot(proof)).toEqual(expected);
  });
  it("cannot execute two leaves", () => {
    const proof: proofs.IExistenceProof = {
      key: toAscii("food"),
      value: toAscii("some longer text"),
      steps: [
        {
          leaf: {
            hash: proofs.HashOp.SHA256,
            length: proofs.LengthOp.VAR_PROTO
          }
        },
        {
          leaf: {
            hash: proofs.HashOp.SHA256,
            length: proofs.LengthOp.VAR_PROTO
          }
        }
      ]
    };
    expect(() => calculateExistenceRoot(proof)).toThrow();
  });
  it("cannot execute inner first", () => {
    const proof: proofs.IExistenceProof = {
      key: toAscii("food"),
      value: toAscii("some longer text"),
      steps: [
        {
          inner: {
            hash: proofs.HashOp.SHA256,
            prefix: fromHex("deadbeef00cafe00")
          }
        }
      ]
    };
    expect(() => calculateExistenceRoot(proof)).toThrow();
  });
  it("can execute leaf then inner", () => {
    const proof: proofs.IExistenceProof = {
      key: toAscii("food"),
      value: toAscii("some longer text"),
      steps: [
        {
          leaf: {
            hash: proofs.HashOp.SHA256,
            length: proofs.LengthOp.VAR_PROTO
          }
        },
        // output: b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265
        {
          inner: {
            hash: proofs.HashOp.SHA256,
            prefix: fromHex("deadbeef00cafe00")
          }
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
