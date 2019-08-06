import { readFileSync } from "fs";

import { proofs } from "./generated/codecimpl";
import { fromHex } from "./helpers";
import {
  calculateExistenceRoot,
  ensureSpec,
  IavlSpec,
  TendermintSpec
} from "./proofs";

describe("calculateExistenceRoot", () => {
  function validateTestVector(filepath: string, spec: proofs.IProofSpec): void {
    const content = readFileSync(filepath).toString();
    const { root, existence } = JSON.parse(content);
    expect(existence).toBeDefined();
    expect(root).toBeDefined();

    const proof = proofs.ExistenceProof.decode(fromHex(existence));
    ensureSpec(proof, spec);
    const calc = calculateExistenceRoot(proof);

    const rootHash = fromHex(root);
    expect(calc).toEqual(rootHash);
  }

  it("should parse iavl 1", () => {
    validateTestVector("../testdata/iavl/existence1.json", IavlSpec);
  });
  it("should parse iavl 2", () => {
    validateTestVector("../testdata/iavl/existence2.json", IavlSpec);
  });
  it("should parse iavl 3", () => {
    validateTestVector("../testdata/iavl/existence3.json", IavlSpec);
  });
  it("should parse iavl 4", () => {
    validateTestVector("../testdata/iavl/existence4.json", IavlSpec);
  });

  it("should parse tendermint 1", () => {
    validateTestVector(
      "../testdata/tendermint/existence1.json",
      TendermintSpec
    );
  });
  it("should parse tendermint 2", () => {
    validateTestVector(
      "../testdata/tendermint/existence2.json",
      TendermintSpec
    );
  });
  it("should parse tendermint 3", () => {
    validateTestVector(
      "../testdata/tendermint/existence3.json",
      TendermintSpec
    );
  });
  it("should parse tendermint 4", () => {
    validateTestVector(
      "../testdata/tendermint/existence4.json",
      TendermintSpec
    );
  });
});
