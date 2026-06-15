/-
The non-existence verifier, mirroring `verify_non_existence` and its helpers
(`ensure_left_most`, `ensure_right_most`, `ensure_left_neighbor`, `get_padding`,
`has_padding`, `order_from_padding`, `is_left_step`, `left_branches_are_empty`,
`right_branches_are_empty`) in `rust/src/verify.rs`.

Faithfulness notes:
* Lexicographic `Vec<u8>` ordering is `bytesLt` (unsigned, shorter-is-less).
* Where the Rust would panic on an out-of-range slice or a `pop().unwrap()` of an
  empty path, the model returns `false`. Panics are a liveness/DoS concern, not a
  soundness one, and rejecting can only shrink the accepted set — a sound
  divergence. Panic-freedom of the Rust is covered separately (Phase 3, Kani).
* `get_padding`'s `child_order.find(|x| x == branch)` binds `idx = branch`, so
  `child_order` acts only as the index set `{0,…,n-1}` here; its permutation
  content matters in the `*_branches_are_empty` checks.
-/
import Ics23.Verify

namespace Ics23

/-- Lexicographic unsigned byte-string `<`, matching `Vec<u8>` `Ord`. -/
def bytesLt : Bytes → Bytes → Bool
  | [], [] => false
  | [], _ :: _ => true
  | _ :: _, [] => false
  | a :: as, b :: bs =>
    if a < b then true
    else if a == b then bytesLt as bs
    else false

/-- The key used for ordering comparisons: hashed first iff the spec sets
`prehash_key_before_comparison` (SMT/JMT). -/
def keyForComparison (H : HashFn) (s : ProofSpec) (key : Bytes) : Bytes :=
  if s.prehashKeyBeforeComparison then H s.leafSpec.prehashKey key else key

/-- Expected prefix/suffix sizes for an inner op sitting at a given branch. -/
structure Padding where
  minPrefix : Int
  maxPrefix : Int
  suffix : Int

/-- `get_padding`: with `idx = branch`, `branch * childSize` bytes of left
siblings in the prefix and `(n-1-branch) * childSize` bytes of right siblings in
the suffix. `none` if `branch` is out of range. -/
def getPadding (isp : InnerSpec) (branch : Nat) : Option Padding :=
  if isp.childOrder.contains branch then
    let cs := isp.childSize
    some { minPrefix := (branch : Int) * cs + isp.minPrefixLength,
           maxPrefix := (branch : Int) * cs + isp.maxPrefixLength,
           suffix := cs * ((isp.childOrder.length : Int) - 1 - (branch : Int)) }
  else none

/-- `has_padding`: the op's prefix length is within `[min,max]` and its suffix
length is exactly the expected number of right-sibling bytes. -/
def hasPadding (op : InnerOp) (pad : Padding) : Bool :=
  (pad.minPrefix ≤ (op.prefixBytes.length : Int))
  && ((op.prefixBytes.length : Int) ≤ pad.maxPrefix)
  && ((op.suffix.length : Int) = pad.suffix)

/-- `order_from_padding`: the unique branch whose padding the op matches. -/
def orderFromPadding (isp : InnerSpec) (op : InnerOp) : Option Nat :=
  (List.range isp.childOrder.length).find? (fun branch =>
    match getPadding isp branch with
    | some pad => hasPadding op pad
    | none => false)

/-- `data[from .. from+len]` if in range, else `none` (models the Rust slice
without panicking). -/
def byteRange (data : Bytes) (start len : Nat) : Option Bytes :=
  if start + len ≤ data.length then some ((data.drop start).take len) else none

/-- `left_branches_are_empty`: the prefix padding bytes are all `empty_child`,
i.e. this is a valid placeholder on a left-most path. -/
def leftBranchesAreEmpty (isp : InnerSpec) (op : InnerOp) : Bool :=
  match orderFromPadding isp op with
  | none => false
  | some idx =>
    let leftBranches := idx
    if leftBranches = 0 then false
    else
      let cs := isp.childSize.toNat
      if op.prefixBytes.length < leftBranches * cs then false
      else
        let actualPrefix := op.prefixBytes.length - leftBranches * cs
        (List.range leftBranches).all (fun i =>
          byteRange op.prefixBytes (actualPrefix + i * cs) cs == some isp.emptyChild)

/-- `right_branches_are_empty`: the suffix padding bytes are all `empty_child`. -/
def rightBranchesAreEmpty (isp : InnerSpec) (op : InnerOp) : Bool :=
  match orderFromPadding isp op with
  | none => false
  | some idx =>
    let rightBranches := isp.childOrder.length - 1 - idx
    let cs := isp.childSize.toNat
    if rightBranches = 0 then false
    else if op.suffix.length ≠ cs then false
    else
      (List.range rightBranches).all (fun i =>
        byteRange op.suffix (i * cs) cs == some isp.emptyChild)

/-- `ensure_left_most`: every step is the left padding or a valid left placeholder. -/
def ensureLeftMost (isp : InnerSpec) (path : List InnerOp) : Bool :=
  match getPadding isp 0 with
  | none => false
  | some pad => path.all (fun step => hasPadding step pad || leftBranchesAreEmpty isp step)

/-- `ensure_right_most`: every step is the right padding or a valid right placeholder. -/
def ensureRightMost (isp : InnerSpec) (path : List InnerOp) : Bool :=
  match getPadding isp (isp.childOrder.length - 1) with
  | none => false
  | some pad => path.all (fun step => hasPadding step pad || rightBranchesAreEmpty isp step)

/-- `is_left_step`: `left` sits immediately left of `right` among the branches. -/
def isLeftStep (isp : InnerSpec) (l r : InnerOp) : Bool :=
  match orderFromPadding isp l, orderFromPadding isp r with
  | some li, some ri => li + 1 = ri
  | _, _ => false

/-- Two inner ops match for the purposes of the neighbor walk (prefix & suffix). -/
def eqPS (a b : InnerOp) : Bool := a.prefixBytes == b.prefixBytes && a.suffix == b.suffix

/-- Drop the longest common (root-side) prefix of two reversed paths, returning
the first divergent op of each plus the remaining (still root→leaf) tails.
`none` if either runs out first (the Rust `pop().unwrap()` would panic). -/
def dropCommonPrefix : List InnerOp → List InnerOp →
    Option InnerOp × List InnerOp × Option InnerOp × List InnerOp
  | a :: as, b :: bs =>
    if eqPS a b then dropCommonPrefix as bs else (some a, as, some b, bs)
  | _, _ => (none, [], none, [])

/-- `ensure_left_neighbor`: strip the shared root path, require the first
divergence to be a left-step, and the inner remainders to be right-most (left
proof) and left-most (right proof). -/
def ensureLeftNeighbor (isp : InnerSpec) (left right : List InnerOp) : Bool :=
  match dropCommonPrefix left.reverse right.reverse with
  | (some topLeft, restL, some topRight, restR) =>
    isLeftStep isp topLeft topRight
    && ensureRightMost isp restL.reverse
    && ensureLeftMost isp restR.reverse
  | _ => false

/-- `verify_non_existence`: the bracketing existence proofs verify, the key lies
strictly between them, and the paths are genuine tree neighbors. -/
def verifyNonExistence (H : HashFn) (p : NonExistenceProof) (s : ProofSpec)
    (root key : Bytes) : Bool :=
  let kfc := keyForComparison H s
  (match p.left with
   | some l => verifyExistence H l s root l.key l.value && bytesLt (kfc l.key) (kfc key)
   | none => true)
  && (match p.right with
   | some r => verifyExistence H r s root r.key r.value && bytesLt (kfc key) (kfc r.key)
   | none => true)
  && (match p.left, p.right with
   | some l, none => ensureRightMost s.innerSpec l.path
   | none, some r => ensureLeftMost s.innerSpec r.path
   | some l, some r => ensureLeftNeighbor s.innerSpec l.path r.path
   | none, none => false)

end Ics23
