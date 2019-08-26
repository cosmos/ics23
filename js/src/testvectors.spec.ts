import { readFileSync } from "fs";

import { ics23 } from "./generated/codecimpl";
import { fromHex } from "./helpers";
import { verifyMembership, verifyNonMembership } from "./ics23";
import { IavlSpec, TendermintSpec } from "./proofs";

describe("calculateExistenceRoot", () => {
  function validateTestVector(filepath: string, spec: ics23.IProofSpec): void {
    const content = readFileSync(filepath).toString();
    const { root, proof, key, value } = JSON.parse(content);
    expect(proof).toBeDefined();
    expect(root).toBeDefined();
    expect(key).toBeDefined();

    const rootHash = fromHex(root);
    const bkey = fromHex(key);
    const commit = ics23.CommitmentProof.decode(fromHex(proof));

    if (value) {
      const bvalue = fromHex(value);
      const valid = verifyMembership(commit, spec, rootHash, bkey, bvalue);
      expect(valid).toBe(true);
    } else {
      // verifyNonExistence(commit.nonexist!, spec, rootHash, bkey);
      const valid = verifyNonMembership(commit, spec, rootHash, bkey);
      expect(valid).toBe(true);
    }
  }

  it("should parse iavl left", () => {
    validateTestVector("../testdata/iavl/exist_left.json", IavlSpec);
  });
  it("should parse iavl right", () => {
    validateTestVector("../testdata/iavl/exist_right.json", IavlSpec);
  });
  it("should parse iavl middle", () => {
    validateTestVector("../testdata/iavl/exist_middle.json", IavlSpec);
  });
  it("should parse iavl left - nonexist", () => {
    validateTestVector("../testdata/iavl/nonexist_left.json", IavlSpec);
  });
  it("should parse iavl right - nonexist", () => {
    validateTestVector("../testdata/iavl/nonexist_right.json", IavlSpec);
  });
  it("should parse iavl middle - nonexist", () => {
    validateTestVector("../testdata/iavl/nonexist_middle.json", IavlSpec);
  });

  it("should parse tendermint left", () => {
    validateTestVector(
      "../testdata/tendermint/exist_left.json",
      TendermintSpec
    );
  });
  it("should parse tendermint right", () => {
    validateTestVector(
      "../testdata/tendermint/exist_right.json",
      TendermintSpec
    );
  });
  it("should parse tendermint middle", () => {
    validateTestVector(
      "../testdata/tendermint/exist_middle.json",
      TendermintSpec
    );
  });
  it("should parse tendermint left - nonexist", () => {
    validateTestVector(
      "../testdata/tendermint/nonexist_left.json",
      TendermintSpec
    );
  });
  it("should parse tendermint right - nonexist", () => {
    validateTestVector(
      "../testdata/tendermint/nonexist_right.json",
      TendermintSpec
    );
  });
  it("should parse tendermint middle - nonexist", () => {
    validateTestVector(
      "../testdata/tendermint/nonexist_middle.json",
      TendermintSpec
    );
  });
});
