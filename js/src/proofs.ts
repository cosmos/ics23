import { proofs } from "./generated/codecimpl";
import { applyInner, applyLeaf } from "./ops";

// Calculate determines the root hash that matches the given proof.
// You must validate the result is what you have in a header.
// Returns error if the calculations cannot be performed.
export function calculateExistenceRoot(
  proof: proofs.IExistenceProof
): Uint8Array {
  if (!proof.key || !proof.value) {
    throw new Error("Existence proof needs key and value set");
  }
  if (!proof.steps || proof.steps.length === 0) {
    throw new Error("Existence Proof needs at least one step");
  }

  const [{ leaf }, ...rem] = proof.steps;
  if (!leaf) {
    throw new Error("Existence proof must start with a leaf operation");
  }
  let res = applyLeaf(leaf, proof.key, proof.value);

  for (const { inner } of rem) {
    if (!inner) {
      throw new Error("All subsequent steps must be inner ops");
    }
    res = applyInner(inner, res);
  }

  return res;
}
