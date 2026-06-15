/-
Regression corpus: concrete inputs encoded as machine-checked acceptance/
rejection facts. These pin the spec-level invariants from
`docs/verification/properties.md` and double as the model's own test suite.

Everything here is decided by computation (no hash function required), so each
`example` is a proof, not a test that might silently stop running.
-/
import Ics23.Verify
import Ics23.NonExist
import Ics23.Specs
import Ics23.Soundness
import Ics23.Tree
import Ics23.TreeNonExist

namespace Ics23

/-! ## Positive controls -/

/-- A leaf-only proof whose leaf matches the IAVL leaf spec passes spec checks. -/
def iavlLeafProof : ExistenceProof :=
  { key := [1], value := [2], leaf := iavlSpec.leafSpec, path := [] }

example : checkExistenceSpec iavlLeafProof iavlSpec = true := by decide

/-- A well-formed Tendermint inner op is accepted by `ensureInner`. -/
def tmValidInner : InnerOp := { hash := .sha256, prefixBytes := [1], suffix := [] }

example : ensureInner tmValidInner tendermintSpec = true := by decide

/-! ## A2 — leaf/inner domain separation

An inner op whose prefix begins with the leaf prefix must be rejected, else a
leaf hash could be reinterpreted as an inner hash (depth confusion). -/
def tmLeafPrefixInner : InnerOp := { hash := .sha256, prefixBytes := [0, 9], suffix := [] }

example : ensureInner tmLeafPrefixInner tendermintSpec = false := by decide

/-! ## A3 — positional unambiguity -/

/-- Suffix length not a multiple of `child_size` is rejected. -/
def tmBadSuffixInner : InnerOp := { hash := .sha256, prefixBytes := [1], suffix := [0] }

example : ensureInner tmBadSuffixInner tendermintSpec = false := by decide

/-- Inner prefix shorter than `min_prefix_length` is rejected. -/
def tmShortPrefixInner : InnerOp := { hash := .sha256, prefixBytes := [], suffix := [] }

example : ensureInner tmShortPrefixInner tendermintSpec = false := by decide

/-! ## C — malformed specs are not well-formed -/

/-- Non-positive `child_size`. -/
def negChildSizeSpec : ProofSpec :=
  { tendermintSpec with
    innerSpec := { tendermintSpec.innerSpec with childSize := 0 } }

example : wellFormedB negChildSizeSpec = false := by decide

/-- A1-split shape: `NoPrefix` length with variable-length prehash on both
fields — the `key ++ value` boundary is ambiguous, so the spec is unsafe. -/
def splitSpec : ProofSpec :=
  { tendermintSpec with
    leafSpec := { tendermintSpec.leafSpec with
                  length := .noPrefix, prehashValue := .noHash } }

example : wellFormedB splitSpec = false := by decide

/-! ## Depth bounds -/

/-- With `min_depth ≠ 0`, a path shorter than `min_depth` is rejected. -/
def depthBoundedSpec : ProofSpec := { iavlSpec with minDepth := 2, maxDepth := 4 }

def shallowProof : ExistenceProof :=
  { key := [1], value := [2], leaf := iavlSpec.leafSpec, path := [] }

example : checkExistenceSpec shallowProof depthBoundedSpec = false := by decide

/-! ## Non-existence: padding, ordering, and empty-branch placeholders

These exercise the SMT inner spec (binary, `child_size = 32`, 32-byte zero
`empty_child`) — the logic where Dragonberry-class bugs live. Mirrors the Rust
`check_empty_branch` test. `native_decide` is used for the 32-byte cases. -/

/-- Left child position: prefix length 1, 32-byte right sibling in the suffix. -/
def smtLeftStep : InnerOp :=
  { hash := .sha256, prefixBytes := [7], suffix := List.replicate 32 0 }

/-- Right child position: 32-byte left sibling in the prefix, empty suffix. -/
def smtRightStep : InnerOp :=
  { hash := .sha256, prefixBytes := List.replicate 32 0 ++ [7], suffix := [] }

example : orderFromPadding smtSpec.innerSpec smtLeftStep = some 0 := by native_decide
example : orderFromPadding smtSpec.innerSpec smtRightStep = some 1 := by native_decide

/-- The left child is immediately left of the right child. -/
example : isLeftStep smtSpec.innerSpec smtLeftStep smtRightStep = true := by native_decide

/-- A genuine left child is left-most. -/
example : ensureLeftMost smtSpec.innerSpec [smtLeftStep] = true := by native_decide

/-- A genuine right child is right-most. -/
example : ensureRightMost smtSpec.innerSpec [smtRightStep] = true := by native_decide

/-- A placeholder at branch 1 whose left sibling is the empty child is a valid
left-most step (the left subtree is empty). -/
def smtLeftPlaceholder : InnerOp :=
  { hash := .sha256, prefixBytes := [9] ++ List.replicate 32 0, suffix := [] }

example : leftBranchesAreEmpty smtSpec.innerSpec smtLeftPlaceholder = true := by native_decide
example : ensureLeftMost smtSpec.innerSpec [smtLeftPlaceholder] = true := by native_decide

/-- A branch-1 step whose left sibling is *not* the empty child is not a valid
left-most placeholder. -/
def smtNotLeftmost : InnerOp :=
  { hash := .sha256, prefixBytes := [9] ++ List.replicate 31 0 ++ [5], suffix := [] }

example : leftBranchesAreEmpty smtSpec.innerSpec smtNotLeftmost = false := by native_decide
example : ensureLeftMost smtSpec.innerSpec [smtNotLeftmost] = false := by native_decide

/-! ## Zellic audit findings (July 2024)

`docs/audits/ICS-23 - Zellic Audit Report.pdf`. Each finding is recorded here as
either a machine-checked fact about the model (where the model has a decidable
witness) or, for the implementation- or liveness-only findings, a note pointing
to where the property lives. The two critical *soundness* findings (3.1, 3.2) map
directly onto hypotheses/checks the formal soundness theorems already rely on. -/

/-! ### 3.1 — Nonexistence soundness depends on sorted keys (Critical)

The headline finding: a nonexistence proof for `b` and an existence proof for
`b` can both verify when the committed tree is *not* key-sorted. The PoC builds
this tree (Figure 3.1) — `b` sits left of `a` under `Y`, so the leaves are out
of order even though `a < b < c` lexicographically:

```
        X
       / \
      Y  "c"
     / \
   "b" "a"
```

In the formal model this is exactly the `SortedTree` hypothesis of
`nonexistence_sound_tree` (and `SortedTreeS` for SMT, `SortedTree` for IAVL):
non-existence soundness is *conditional* on a store invariant the verifier never
checks (modeling note F4). The witness below reconstructs the Zellic tree and
machine-checks that (a) `b` is a genuine member and (b) the tree is provably
**not** `SortedTree` — so the theorem's precondition is precisely what rules the
attack out. -/

/-- The Zellic Figure 3.1 tree as an `MTree` (hash/leaf ops are irrelevant to
ordering, so any concrete ops do). Keys: `a = 0x61`, `b = 0x62`, `c = 0x63`. -/
def zellicUnsortedTree : MTree :=
  let lop := tendermintSpec.leafSpec
  .node .sha256 [1] [] []
    (.node .sha256 [1] [] [] (.leaf lop [0x62] [0x62]) (.leaf lop [0x61] [0x61]))
    (.leaf lop [0x63] [0x63])

/-- The absent-claimed key `b` is in fact a genuine member of the tree. -/
example : TreeMember [0x62] [0x62] zellicUnsortedTree := Or.inl (Or.inl ⟨rfl, rfl⟩)

/-- The bracketing keys order as `a < b < c`, so the forged nonexistence proof's
own ordering checks (`left.key < key < right.key`) pass — order alone does not
imply absence. -/
example : bytesLt [0x61] [0x62] = true := by decide
example : bytesLt [0x62] [0x63] = true := by decide

/-- …yet the tree is **not** `SortedTree`: under `Y`, `maxKey(left) = b` is not
`< minKey(right) = a`. This is the precondition `nonexistence_sound_tree`
requires; without it the gap argument (`node_gap_no_member`) does not hold. -/
example : ¬ SortedTree zellicUnsortedTree := fun h => absurd h.1.2.2 (by decide)

/-! ### 3.2 — Forging nonexistence proofs with incorrect `MaxPrefixLength` (Critical)

If `max_prefix_length ≥ min_prefix_length + child_size`, the same bytes can be
read as both prefix and child in different parts of verification, breaking
nonexistence soundness and admitting nodes with more children than `child_order`
specifies. The PoC sets `MaxPrefixLength = 33` on `tendermint_spec`
(`min = 1`, `child_size = 32`, so `33 = min + child_size`).

The model rules this out structurally: `innerWFB` requires
`max_prefix_length < min_prefix_length + child_size`, and the honest-root
soundness theorems pin a node split via exactly this bound (`split_pins` /
`split_pins_var` need `max_prefix_length < child_size + min`). -/

/-- The Zellic 3.2 malformed spec: `max = min + child_size` (the boundary). -/
def zellicMaxPrefixSpec : ProofSpec :=
  { tendermintSpec with
    innerSpec := { tendermintSpec.innerSpec with maxPrefixLength := 33 } }

example : wellFormedB zellicMaxPrefixSpec = false := by decide

/-- An inner op whose prefix lands in the now-overlapping window is still
rejected by `ensure_inner` against that spec — the `max < min + child_size`
check is duplicated at the op level. -/
def zellicOverlapInner : InnerOp :=
  { hash := .sha256, prefixBytes := List.replicate 33 1, suffix := [] }

example : ensureInner zellicOverlapInner zellicMaxPrefixSpec = false := by decide

/-! ### 3.3 — TypeScript does not validate IAVL prefixes (Critical)

TypeScript-implementation-specific (the Dragonberry class, where the attacker
needs only a single key's suffix, not control of tree shape). The TS library
was *removed* upstream rather than fixed, and the Lean model models the Rust/Go
verifier, so there is no model artifact. The model deliberately *omits* the IAVL
prefix checks (modeling assumption 2 in `properties.md`) as a sound
over-approximation — it accepts at least every proof the stricter code accepts —
so a soundness theorem here transfers to the implementations that *do* check.
Related: 3.8. -/

/-! ### 3.4 — Zero division in `ensure_inner` (High)

`child_size = 0` triggers a divide-by-zero panic in the Rust `ensure_inner`. The
model is panic-free here (Lean `Int` `%` is total), and the soundness-relevant
gate is well-formedness: a zero (or negative) `child_size` spec is not
`WellFormed`. See the positive witness `negChildSizeSpec` above
(`wellFormedB … = false`); the op-level check is below. The Rust panic-freedom
itself is a liveness property covered by the Kani harnesses (Phase 3). -/

/-- `ensure_inner` rejects any op against a zero-`child_size` spec (the
`child_size > 0` conjunct), so the model never reaches the division. -/
example : ensureInner tmValidInner negChildSizeSpec = false := by decide

/-! ### 3.5 — Unrestricted proof size (Medium)

No `MaxDepth` ⇒ arbitrarily large, expensive-to-verify proofs (a liveness/DoS
concern, outside soundness scope). The model *has* the depth gate
(`checkExistenceSpec`'s `min_depth`/`max_depth` bounds); the shipped specs set
`min_depth = max_depth = 0`, which disables it — faithfully reflecting the
finding (the fix is to set a bound). The `depthBoundedSpec` witnesses above show
the gate working; the over-deep rejection is below. -/

/-- A valid IAVL inner op (prefix length 4 ∈ [4, 45], empty suffix). -/
def iavlDepthInner : InnerOp :=
  { hash := .sha256, prefixBytes := [1, 2, 3, 4], suffix := [] }

/-- A path deeper than `max_depth = 4` is rejected once the depth gate is armed
(`min_depth ≠ 0`). -/
def zellicDeepProof : ExistenceProof :=
  { key := [1], value := [2], leaf := iavlSpec.leafSpec,
    path := List.replicate 5 iavlDepthInner }

example : checkExistenceSpec zellicDeepProof depthBoundedSpec = false := by decide

/-! ### 3.6 — Panic with out-of-bounds `InnerSpec.ChildOrder` (Low)

`getPosition` panics for a branch absent from a non-contiguous `child_order`
(e.g. `[0, 2]`). The model's analog (`getPadding`) is *total* — it returns
`none` for an out-of-range branch instead of panicking (faithfulness note in
`NonExist.lean`: where the Rust would panic, the model returns `none`/`false`,
a sound divergence). So the 3.6 panic class cannot occur in the model. The
honest-root theorems are additionally stated only for the binary
`child_order = [0, 1]` shipped shape. -/

/-- The Zellic 3.6 non-contiguous child order. -/
def zellicGapInnerSpec : InnerSpec :=
  { smtSpec.innerSpec with childOrder := [0, 2] }

/-- Total, not panicking: branch `1` is absent from `[0, 2]` ⇒ `none`. -/
example : getPadding zellicGapInnerSpec 1 = none := by decide

/-- A present branch still resolves. -/
example : (getPadding zellicGapInnerSpec 0).isSome = true := by decide

/-! ### 3.7 — Panic in `decompressExist` (Low)

Out-of-bounds index when decompressing a compressed batch proof. Compressed and
batched proofs were *removed* upstream (PRs 390/391/392) and are **out of scope**
for the model, which covers single existence/non-existence proofs only. No model
artifact. -/

/-! ### 3.8 — Checks for `IavlSpec`/`TendermintSpec` could be made stricter (Informational)

Recommends tightening the verifier to reject more malformed-but-currently-accepted
IAVL/Tendermint proofs (IAVL leaf prefix's second varint `= 1`; Tendermint inner
prefix `= [1]`). The model's verifier (`ensureInner`/`ensureLeaf`) is the looser
upstream surface, but the honest-root *tree* models bake the stricter shape into
their well-formedness predicates: `WFTree tendermintSpec` requires each node
prefix to have length `= min_prefix_length = 1` and not begin with the leaf byte,
and `WFTreeI`/`WFTreeS` similarly constrain the IAVL/SMT node shape. Related: 3.3. -/

end Ics23
