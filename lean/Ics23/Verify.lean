/-
The existence verifier, mirroring `verify_existence` / `check_existence_spec` /
`calculate_existence_root_for_spec` in `rust/src/verify.rs`.

Abstraction note: the IAVL-specific prefix sanity checks (`ensure_leaf_prefix`,
`ensure_inner_prefix` in `rust/src/api.rs`) are *omitted* here. They fire only
when the spec equals the IAVL spec and can only reject additional proofs. Leaving
them out enlarges the accepted set, so a soundness theorem proved against this
(more permissive) model implies the same property for the stricter Rust code.
-/
import Ics23.Ops

namespace Ics23

/-- `has_prefix`: is `pre` a prefix of `data`? -/
def hasPrefix (pre data : Bytes) : Bool :=
  pre.length ≤ data.length && pre == data.take pre.length

/-- `ensure_leaf`: the proof's leaf op must match the spec's leaf op. -/
def ensureLeaf (leaf leafSpec : LeafOp) : Bool :=
  leafSpec.hash == leaf.hash
  && leafSpec.prehashKey == leaf.prehashKey
  && leafSpec.prehashValue == leaf.prehashValue
  && leafSpec.length == leaf.length
  && hasPrefix leafSpec.prefixBytes leaf.prefixBytes

/-- `ensure_inner`: an inner op must conform to the inner spec, and in
particular must not carry the leaf prefix (domain separation). -/
def ensureInner (inner : InnerOp) (s : ProofSpec) : Bool :=
  let isp := s.innerSpec
  let maxLeftChildBytes : Int := ((isp.childOrder.length : Int) - 1) * isp.childSize
  inner.hash == isp.hash
  && (!hasPrefix s.leafSpec.prefixBytes inner.prefixBytes)
  && (isp.minPrefixLength ≤ (inner.prefixBytes.length : Int))
  && ((inner.prefixBytes.length : Int) ≤ isp.maxPrefixLength + maxLeftChildBytes)
  && (isp.childSize > 0)
  && (isp.maxPrefixLength < isp.minPrefixLength + isp.childSize)
  && ((inner.suffix.length : Int) % isp.childSize = 0)

/-- `check_existence_spec`: the leaf and every inner op conform to the spec, and
the depth bounds are respected when `minDepth ≠ 0` (matching the Rust gate). -/
def checkExistenceSpec (p : ExistenceProof) (s : ProofSpec) : Bool :=
  ensureLeaf p.leaf s.leafSpec
  && (if s.minDepth ≠ 0 then
        (s.minDepth ≤ (p.path.length : Int)) && ((p.path.length : Int) ≤ s.maxDepth)
      else true)
  && p.path.all (fun step => ensureInner step s)

/-- Fold the inner ops over the leaf hash, enforcing the spec's `child_size`
guard at each step (the `hash.len() > child_size && child_size >= 32` bail). -/
def applyPath (H : HashFn) (isp : InnerSpec) : Bytes → List InnerOp → Option Bytes
  | h, [] => some h
  | h, step :: rest =>
    match applyInner H step h with
    | none => none
    | some h' =>
      if (h'.length : Int) > isp.childSize ∧ isp.childSize ≥ 32 then none
      else applyPath H isp h' rest

/-- `calculate_existence_root_for_spec`: compute the root the proof claims. -/
def calculateExistenceRoot (H : HashFn) (s : ProofSpec)
    (p : ExistenceProof) : Option Bytes :=
  if p.key.isEmpty then none
  else if p.value.isEmpty then none
  else match applyLeaf H p.leaf p.key p.value with
    | none => none
    | some h => applyPath H s.innerSpec h p.path

/-- `verify_existence`: spec conformance, key/value binding, and root match. -/
def verifyExistence (H : HashFn) (p : ExistenceProof) (s : ProofSpec)
    (root key value : Bytes) : Bool :=
  checkExistenceSpec p s
  && p.key == key
  && p.value == value
  && (match calculateExistenceRoot H s p with
      | some r => r == root
      | none => false)

end Ics23
