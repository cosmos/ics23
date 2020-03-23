import { ics23 } from "./generated/codecimpl";

import { fromHex, toAscii } from "./helpers";
import { calculateExistenceRoot, ensureSpec, IavlSpec } from "./proofs";

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

describe("ensureSpec", () => {
  const validLeaf = IavlSpec.leafSpec;
  const invalidLeaf = {
    prefix: Uint8Array.from([0]),
    hash: ics23.HashOp.SHA512,
    prehashValue: ics23.HashOp.NO_HASH,
    prehashKey: ics23.HashOp.NO_HASH,
    length: ics23.LengthOp.VAR_PROTO
  };

  const validInner = {
    hash: ics23.HashOp.SHA256,
    prefix: fromHex("deadbeef00cafe00")
    // suffix: Uint8Array.from([]),
  };
  const invalidInner = {
    hash: ics23.HashOp.SHA256,
    prefix: fromHex("aa")
  };
  const invalidInnerHash = {
    hash: ics23.HashOp.SHA512,
    prefix: fromHex("deadbeef00cafe00")
  };

  const depthLimitedSpec = {
    leafSpec: IavlSpec.leafSpec,
    innerSpec: IavlSpec.innerSpec,
    minDepth: 2,
    maxDepth: 4
  };

  it("rejects empty proof", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar")
    };
    expect(() => ensureSpec(proof, IavlSpec)).toThrow();
  });

  it("accepts one valid leaf", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar"),
      leaf: validLeaf
    };
    // fail if this throws (invalid spec)
    ensureSpec(proof, IavlSpec);
  });

  it("rejects invalid leaf", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar"),
      leaf: invalidLeaf
    };
    expect(() => ensureSpec(proof, IavlSpec)).toThrow();
  });

  it("rejects inner without leaf", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar"),
      path: [validInner]
    };
    expect(() => ensureSpec(proof, IavlSpec)).toThrow();
  });

  it("accepts leaf with one inner", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar"),
      leaf: validLeaf,
      path: [validInner]
    };
    // fail if this throws (invalid spec)
    ensureSpec(proof, IavlSpec);
  });

  it("rejects with invalid inner (prefix)", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar"),
      leaf: validLeaf,
      path: [invalidInner, validInner]
    };
    expect(() => ensureSpec(proof, IavlSpec)).toThrow();
  });

  it("rejects with invalid inner (hash)", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar"),
      leaf: validLeaf,
      path: [validInner, invalidInnerHash]
    };
    expect(() => ensureSpec(proof, IavlSpec)).toThrow();
  });

  it("accepts depth limited with proper number of nodes", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar"),
      leaf: validLeaf,
      path: [validInner, validInner, validInner]
    };
    // fail if this throws (invalid spec)
    ensureSpec(proof, depthLimitedSpec);
  });

  it("rejects depth limited with too few nodes", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar"),
      leaf: validLeaf,
      path: [validInner]
    };
    expect(() => ensureSpec(proof, IavlSpec)).toThrow();
  });

  it("rejects depth limited with too many nodes", () => {
    const proof: ics23.IExistenceProof = {
      key: toAscii("foo"),
      value: toAscii("bar"),
      leaf: validLeaf,
      path: [validInner, validInner, validInner, validInner, validInner]
    };
    expect(() => ensureSpec(proof, IavlSpec)).toThrow();
  });
});
