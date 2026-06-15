/-
The verifier model, run end to end with a concrete SHA-256.

This grounds the abstract model: roots are computed with real hashing and
forgeries are rejected by computation (`native_decide`). It is the seed of the
Phase 2a differential oracle — the same `concreteHash` can drive the model
against the Rust/Go implementations over shared test vectors.
-/
import Ics23.Verify
import Ics23.NonExist
import Ics23.Specs
import Ics23.Sha256

namespace Ics23

/-- A concrete `HashFn` covering the ops the shipped specs use (`noHash`,
`sha256`). Other ops are placeholders — not exercised by IAVL/Tendermint/SMT. -/
def concreteHash : HashFn := fun op data =>
  match op with
  | .sha256 => Sha256.hash data
  | _ => data

/-! ## End-to-end: a single-leaf IAVL proof -/

def demoLeafProof : ExistenceProof :=
  { key := [0x01], value := [0x02], leaf := iavlSpec.leafSpec, path := [] }

/-- The root computed by the model with real SHA-256. -/
def demoLeafRoot : Bytes := (calculateExistenceRoot concreteHash iavlSpec demoLeafProof).getD []

/-- The honest proof verifies. -/
example : verifyExistence concreteHash demoLeafProof iavlSpec demoLeafRoot [0x01] [0x02] = true := by
  native_decide

/-- A different value under the same root is rejected (the hash chain no longer
matches) — a value-swap forgery, refuted by computation. -/
def demoLeafForgery : ExistenceProof := { demoLeafProof with value := [0x03] }

example : verifyExistence concreteHash demoLeafForgery iavlSpec demoLeafRoot [0x01] [0x03] = false := by
  native_decide

/-! ## End-to-end: an IAVL proof with one inner step -/

def demoInner : InnerOp := { hash := .sha256, prefixBytes := [0x01, 0x02, 0x03, 0x04], suffix := [] }

def demoPathProof : ExistenceProof := { demoLeafProof with path := [demoInner] }

def demoPathRoot : Bytes := (calculateExistenceRoot concreteHash iavlSpec demoPathProof).getD []

/-- The honest two-level proof verifies. -/
example : verifyExistence concreteHash demoPathProof iavlSpec demoPathRoot [0x01] [0x02] = true := by
  native_decide

/-- The single-leaf root does not validate the two-level proof (and vice versa):
distinct tree shapes give distinct roots. -/
example : verifyExistence concreteHash demoPathProof iavlSpec demoLeafRoot [0x01] [0x02] = false := by
  native_decide

/-! ## End-to-end: non-existence over a real 2-leaf Tendermint tree

Leaves at keys `01` and `03`; the tree's inner node is `sha256(0x01 ‖ left ‖ right)`.
We prove non-membership of `02` (which sorts strictly between) and show the
verifier *refuses* to prove non-membership of `01` (which exists). -/

def tmLeaf : LeafOp := tendermintSpec.leafSpec
def lhA : Bytes := (applyLeaf concreteHash tmLeaf [0x01] [0x0a]).getD []
def lhB : Bytes := (applyLeaf concreteHash tmLeaf [0x03] [0x0b]).getD []

/-- Left leaf's path: it is the left child, right sibling hash in the suffix. -/
def exA : ExistenceProof :=
  { key := [0x01], value := [0x0a], leaf := tmLeaf,
    path := [{ hash := .sha256, prefixBytes := [1], suffix := lhB }] }

/-- Right leaf's path: it is the right child, left sibling hash in the prefix. -/
def exB : ExistenceProof :=
  { key := [0x03], value := [0x0b], leaf := tmLeaf,
    path := [{ hash := .sha256, prefixBytes := [1] ++ lhA, suffix := [] }] }

def tmRoot : Bytes := (calculateExistenceRoot concreteHash tendermintSpec exA).getD []

/-- Both leaves hash up to the same root. -/
example : calculateExistenceRoot concreteHash tendermintSpec exB = some tmRoot := by native_decide

/-- Non-membership of `02` (strictly between the two leaves) verifies. -/
def nexProof : NonExistenceProof := { key := [0x02], left := some exA, right := some exB }

example : verifyNonExistence concreteHash nexProof tendermintSpec tmRoot [0x02] = true := by
  native_decide

/-- The verifier refuses to prove non-membership of `01`, which exists — the
Theorem B property, demonstrated computationally. -/
def nexForgery : NonExistenceProof := { key := [0x01], left := some exA, right := some exB }

example : verifyNonExistence concreteHash nexForgery tendermintSpec tmRoot [0x01] = false := by
  native_decide

/-! ## Finding F3, machine-checked: positional ambiguity of a node hash

For the Tendermint spec the 65-byte preimage `[1] ‖ A ‖ B` is accepted by
`ensure_inner` both as a **left**-child step (child `A`, sibling `B` in the
suffix) and as a **right**-child step (child `B`, sibling `A` in the prefix).
Both readings produce the *same* node hash from *different* children — so the
spec checks alone do not pin a node's recursive child. See
`docs/verification/properties.md` (F3): this is why general `existence_binding`
needs more than collision resistance. -/

def ambA : Bytes := List.replicate 32 1
def ambB : Bytes := List.replicate 32 2

/-- Left-child reading of the node. -/
def opLeft : InnerOp := { hash := .sha256, prefixBytes := [1], suffix := ambB }

/-- Right-child reading of the *same* node bytes. -/
def opRight : InnerOp := { hash := .sha256, prefixBytes := [1] ++ ambA, suffix := [] }

example : ensureInner opLeft tendermintSpec = true := by native_decide
example : ensureInner opRight tendermintSpec = true := by native_decide

/-- The two accepted readings hash to the same node from different children. -/
example : applyInner concreteHash opLeft ambA = applyInner concreteHash opRight ambB := by
  native_decide

example : ambA ≠ ambB := by native_decide

end Ics23
