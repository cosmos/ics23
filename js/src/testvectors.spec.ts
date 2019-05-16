import { readFileSync } from "fs";

import { proofs } from "./generated/codecimpl";
import { fromHex } from "./helpers";
import { calculateExistenceRoot } from "./proofs";

describe("calculateExistenceRoot", () => {
  function validateTestVector(filepath: string): void {
    const content = readFileSync(filepath).toString();
    const { root, existence } = JSON.parse(content);
    expect(existence).toBeDefined();
    expect(root).toBeDefined();

    const rootHash = fromHex(root);
    const proof = proofs.ExistenceProof.decode(fromHex(existence));
    const calc = calculateExistenceRoot(proof);

    expect(calc).toEqual(rootHash);
  }

  it("should parse 1", () => {
    validateTestVector("../testdata/iavl/existence1.json");
  });
  it("should parse 2", () => {
    validateTestVector("../testdata/iavl/existence2.json");
  });
  it("should parse 3", () => {
    validateTestVector("../testdata/iavl/existence3.json");
  });
  it("should parse 4", () => {
    validateTestVector("../testdata/iavl/existence4.json");
  });
});
