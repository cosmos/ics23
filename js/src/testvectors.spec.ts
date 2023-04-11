import { readFileSync } from "fs";

import { compress } from "./compress";
import { ics23 } from "./generated/codecimpl";
import {
  batchVerifyMembership,
  batchVerifyNonMembership,
  verifyMembership,
  verifyNonMembership,
} from "./ics23";
import { iavlSpec, smtSpec, tendermintSpec } from "./proofs";
import { fromHex } from "./testhelpers.spec";

describe("calculateExistenceRoot", () => {
  interface RefData {
    readonly root: Uint8Array;
    readonly key: Uint8Array;
    readonly value?: Uint8Array;
  }

  interface LoadResult {
    readonly proof: ics23.ICommitmentProof;
    readonly data: RefData;
  }

  function loadFile(filepath: string): LoadResult {
    const content = readFileSync(filepath).toString();
    const { root, proof, key, value } = JSON.parse(content);
    expect(proof).toBeDefined();
    expect(root).toBeDefined();
    expect(key).toBeDefined();

    const commit = ics23.CommitmentProof.decode(fromHex(proof));

    const data = {
      root: fromHex(root),
      key: fromHex(key),
      value: value ? fromHex(value) : undefined,
    };

    return { proof: commit, data };
  }

  interface BatchResult {
    readonly proof: ics23.ICommitmentProof;
    readonly data: readonly RefData[];
  }

  function validateTestVector(filepath: string, spec: ics23.IProofSpec): void {
    const {
      proof,
      data: { root, key, value },
    } = loadFile(filepath);
    if (value) {
      const valid = verifyMembership(proof, spec, root, key, value);
      expect(valid).toBe(true);
    } else {
      const valid = verifyNonMembership(proof, spec, root, key);
      expect(valid).toBe(true);
    }
  }

  it("should parse iavl left", () => {
    validateTestVector("../testdata/iavl/exist_left.json", iavlSpec);
  });
  it("should parse iavl right", () => {
    validateTestVector("../testdata/iavl/exist_right.json", iavlSpec);
  });
  it("should parse iavl middle", () => {
    validateTestVector("../testdata/iavl/exist_middle.json", iavlSpec);
  });
  it("should parse iavl left - nonexist", () => {
    validateTestVector("../testdata/iavl/nonexist_left.json", iavlSpec);
  });
  it("should parse iavl right - nonexist", () => {
    validateTestVector("../testdata/iavl/nonexist_right.json", iavlSpec);
  });
  it("should parse iavl middle - nonexist", () => {
    validateTestVector("../testdata/iavl/nonexist_middle.json", iavlSpec);
  });

  it("should parse tendermint left", () => {
    validateTestVector(
      "../testdata/tendermint/exist_left.json",
      tendermintSpec
    );
  });
  it("should parse tendermint right", () => {
    validateTestVector(
      "../testdata/tendermint/exist_right.json",
      tendermintSpec
    );
  });
  it("should parse tendermint middle", () => {
    validateTestVector(
      "../testdata/tendermint/exist_middle.json",
      tendermintSpec
    );
  });
  it("should parse tendermint left - nonexist", () => {
    validateTestVector(
      "../testdata/tendermint/nonexist_left.json",
      tendermintSpec
    );
  });
  it("should parse tendermint right - nonexist", () => {
    validateTestVector(
      "../testdata/tendermint/nonexist_right.json",
      tendermintSpec
    );
  });
  it("should parse tendermint middle - nonexist", () => {
    validateTestVector(
      "../testdata/tendermint/nonexist_middle.json",
      tendermintSpec
    );
  });

  function loadBatch(files: readonly string[]): BatchResult {
    let refs: readonly RefData[] = [];
    let entries: readonly ics23.IBatchEntry[] = [];

    for (const file of files) {
      const { proof, data } = loadFile(file);
      refs = [...refs, data];
      if (proof.exist) {
        entries = [...entries, { exist: proof.exist }];
      } else if (proof.nonexist) {
        entries = [...entries, { nonexist: proof.nonexist }];
      }
    }
    const commit: ics23.ICommitmentProof = {
      batch: {
        entries: entries as ics23.IBatchEntry[],
      },
    };

    return {
      proof: commit,
      data: refs,
    };
  }

  function validateBatch(
    proof: ics23.ICommitmentProof,
    spec: ics23.IProofSpec,
    data: RefData
  ): void {
    const { root, key, value } = data;
    if (value) {
      let valid = verifyMembership(proof, spec, root, key, value);
      expect(valid).toBe(true);
      const items = new Map([[key, value]]);
      valid = batchVerifyMembership(proof, spec, root, items);
      expect(valid).toBe(true);
    } else {
      let valid = verifyNonMembership(proof, spec, root, key);
      expect(valid).toBe(true);
      const keys: readonly Uint8Array[] = [key];
      valid = batchVerifyNonMembership(proof, spec, root, keys);
      expect(valid).toBe(true);
    }
  }

  it("should validate iavl batch exist", () => {
    const { proof, data } = loadBatch([
      "../testdata/iavl/exist_left.json",
      "../testdata/iavl/exist_right.json",
      "../testdata/iavl/exist_middle.json",
      "../testdata/iavl/nonexist_left.json",
      "../testdata/iavl/nonexist_right.json",
      "../testdata/iavl/nonexist_middle.json",
    ]);
    validateBatch(proof, iavlSpec, data[0]);
  });

  it("should validate iavl batch nonexist", () => {
    const { proof, data } = loadBatch([
      "../testdata/iavl/exist_left.json",
      "../testdata/iavl/exist_right.json",
      "../testdata/iavl/exist_middle.json",
      "../testdata/iavl/nonexist_left.json",
      "../testdata/iavl/nonexist_right.json",
      "../testdata/iavl/nonexist_middle.json",
    ]);
    validateBatch(proof, iavlSpec, data[5]);
  });

  it("should validate compressed iavl batch exist", () => {
    const { proof, data } = loadBatch([
      "../testdata/iavl/exist_left.json",
      "../testdata/iavl/exist_right.json",
      "../testdata/iavl/exist_middle.json",
      "../testdata/iavl/nonexist_left.json",
      "../testdata/iavl/nonexist_right.json",
      "../testdata/iavl/nonexist_middle.json",
    ]);
    const small = compress(proof);

    // ensure this is acutally a different format
    const origBin = ics23.CommitmentProof.encode(proof).finish();
    const origBin2 = ics23.CommitmentProof.encode(proof).finish();
    const smallBin = ics23.CommitmentProof.encode(small).finish();
    expect(origBin).toEqual(origBin2);
    expect(origBin).not.toEqual(smallBin);

    validateBatch(small, iavlSpec, data[0]);
  });

  it("should validate compressed iavl batch nonexist", () => {
    const { proof, data } = loadBatch([
      "../testdata/iavl/exist_left.json",
      "../testdata/iavl/exist_right.json",
      "../testdata/iavl/exist_middle.json",
      "../testdata/iavl/nonexist_left.json",
      "../testdata/iavl/nonexist_right.json",
      "../testdata/iavl/nonexist_middle.json",
    ]);
    const small = compress(proof);

    // ensure this is acutally a different format
    const origBin = ics23.CommitmentProof.encode(proof).finish();
    const origBin2 = ics23.CommitmentProof.encode(proof).finish();
    const smallBin = ics23.CommitmentProof.encode(small).finish();
    expect(origBin).toEqual(origBin2);
    expect(origBin).not.toEqual(smallBin);

    validateBatch(small, iavlSpec, data[5]);
  });

  it("should validate tendermint batch exist", () => {
    const { proof, data } = loadBatch([
      "../testdata/tendermint/exist_left.json",
      "../testdata/tendermint/exist_right.json",
      "../testdata/tendermint/exist_middle.json",
      "../testdata/tendermint/nonexist_left.json",
      "../testdata/tendermint/nonexist_right.json",
      "../testdata/tendermint/nonexist_middle.json",
    ]);
    validateBatch(proof, tendermintSpec, data[2]);
  });

  it("should validate tendermint batch nonexist", () => {
    const { proof, data } = loadBatch([
      "../testdata/tendermint/exist_left.json",
      "../testdata/tendermint/exist_right.json",
      "../testdata/tendermint/exist_middle.json",
      "../testdata/tendermint/nonexist_left.json",
      "../testdata/tendermint/nonexist_right.json",
      "../testdata/tendermint/nonexist_middle.json",
    ]);
    validateBatch(proof, tendermintSpec, data[3]);
  });

  it("should validate smt batch exist", () => {
    const { proof, data } = loadBatch([
      "../testdata/smt/exist_left.json",
      "../testdata/smt/exist_right.json",
      "../testdata/smt/exist_middle.json",
      "../testdata/smt/nonexist_left.json",
      "../testdata/smt/nonexist_right.json",
      "../testdata/smt/nonexist_middle.json",
    ]);
    validateBatch(proof, smtSpec, data[2]);
  });

  it("should validate smt batch nonexist", () => {
    const { proof, data } = loadBatch([
      "../testdata/smt/exist_left.json",
      "../testdata/smt/exist_right.json",
      "../testdata/smt/exist_middle.json",
      "../testdata/smt/nonexist_left.json",
      "../testdata/smt/nonexist_right.json",
      "../testdata/smt/nonexist_middle.json",
    ]);
    validateBatch(proof, smtSpec, data[3]);
  });
});
