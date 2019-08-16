import { readFileSync } from "fs";

import { proofs } from "./generated/codecimpl";
import { fromHex } from "./helpers";
import { verifyMembership, verifyNonMembership } from "./ics23";
import { IavlSpec, TendermintSpec } from "./proofs";

describe("calculateExistenceRoot", () => {
  function validateTestVector(filepath: string, spec: proofs.IProofSpec): void {
    const content = readFileSync(filepath).toString();
    const { root, proof, key, value } = JSON.parse(content);
    expect(proof).toBeDefined();
    expect(root).toBeDefined();
    expect(key).toBeDefined();

    const rootHash = fromHex(root);
    const bkey = fromHex(key);
    const commit = proofs.CommitmentProof.decode(fromHex(proof));

    if (value) {
      const bvalue = fromHex(value);
      verifyMembership(commit, spec, rootHash, bkey, bvalue);
    } else {
      verifyNonMembership(commit, spec, rootHash, bkey);
    }
  }

  it("should parse iavl 1", () => {
    validateTestVector("../testdata/iavl/exist_left.json", IavlSpec);
  });
  it("should parse iavl 2", () => {
    validateTestVector("../testdata/iavl/exist_right.json", IavlSpec);
  });
  it("should parse iavl 3", () => {
    validateTestVector("../testdata/iavl/exist_middle.json", IavlSpec);
  });
  it("should parse iavl 1 - nonexist", () => {
    validateTestVector("../testdata/iavl/nonexist_left.json", IavlSpec);
  });
  it("should parse iavl 2 - nonexist", () => {
    validateTestVector("../testdata/iavl/nonexist_right.json", IavlSpec);
  });
  it("should parse iavl 3 - nonexist", () => {
    validateTestVector("../testdata/iavl/nonexist_middle.json", IavlSpec);
  });

  it("should parse tendermint 1", () => {
    validateTestVector(
      "../testdata/tendermint/exist_left.json",
      TendermintSpec
    );
  });
  it("should parse tendermint 2", () => {
    validateTestVector(
      "../testdata/tendermint/exist_right.json",
      TendermintSpec
    );
  });
  it("should parse tendermint 3", () => {
    validateTestVector(
      "../testdata/tendermint/exist_middle.json",
      TendermintSpec
    );
  });
  it("should parse tendermint 1 - nonexist", () => {
    validateTestVector(
      "../testdata/tendermint/nonexist_left.json",
      TendermintSpec
    );
  });
  it("should parse tendermint 2 - nonexist", () => {
    validateTestVector(
      "../testdata/tendermint/nonexist_right.json",
      TendermintSpec
    );
  });
  it("should parse tendermint 3 - nonexist", () => {
    validateTestVector(
      "../testdata/tendermint/nonexist_middle.json",
      TendermintSpec
    );
  });
});
