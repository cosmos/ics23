/-
ICS23 proto types, modeled in Lean.

These mirror `proto/cosmos/ics23/v1/proofs.proto` and the Rust generated types
in `rust/src/cosmos.ics23.v1.rs`. Modeling decisions (kept deliberately faithful
to the verifier's soundness-relevant surface):

* Bytes are `List UInt8`.
* `HashOp`/`LengthOp` are inductives rather than the proto's `i32` enums. Only
  the operations the verifier actually supports are represented; an unsupported
  proto value corresponds to an input the verifier rejects.
* `ProofSpec.leafSpec`/`innerSpec` and `ExistenceProof.leaf` are non-optional
  here. In the proto/Rust they are `Option`, and a missing value causes an
  immediate `bail!`. Dropping the `None` cases only removes inputs the verifier
  rejects, so it is a sound abstraction for a soundness proof (it can only
  enlarge the accepted set, never shrink it).
* Spec integer fields (`childSize`, `minPrefixLength`, ...) are `Int`, matching
  the `i32` proto fields, so that malformed (e.g. negative) specs are
  representable and ruled out explicitly by `WellFormed` rather than by typing.
* `childOrder` entries are `Nat` (they are always non-negative indices).
-/

namespace Ics23

/-- A byte string. -/
abbrev Bytes := List UInt8

/-- Hash operations supported by the verifier (`do_hash` in `rust/src/ops.rs`). -/
inductive HashOp where
  | noHash
  | sha256
  | sha512
  | keccak256
  | ripemd160
  | bitcoin
  | sha512256
  | blake2b512
  | blake2s256
  | blake3
  deriving Repr, DecidableEq, BEq, Inhabited

instance : LawfulBEq HashOp where
  eq_of_beq {a b} h := by cases a <;> cases b <;> first | rfl | exact absurd h (by decide)
  rfl {a} := by cases a <;> rfl

/-- Length-prefix operations supported by the verifier (`do_length`). -/
inductive LengthOp where
  | noPrefix
  | varProto
  | require32Bytes
  | require64Bytes
  | fixed32Big
  | fixed32Little
  | fixed64Big
  | fixed64Little
  deriving Repr, DecidableEq, BEq, Inhabited

instance : LawfulBEq LengthOp where
  eq_of_beq {a b} h := by cases a <;> cases b <;> first | rfl | exact absurd h (by decide)
  rfl {a} := by cases a <;> rfl

/-- `LeafOp`: how a (key, value) pair is transformed into a leaf hash. -/
structure LeafOp where
  hash : HashOp
  prehashKey : HashOp
  prehashValue : HashOp
  length : LengthOp
  prefixBytes : Bytes  -- proto `prefix`; renamed: `prefix` is a Lean keyword
  deriving Repr, DecidableEq, BEq, Inhabited

/-- `InnerOp`: one step up the tree — `hash(prefix ++ child ++ suffix)`. -/
structure InnerOp where
  hash : HashOp
  prefixBytes : Bytes  -- proto `prefix`
  suffix : Bytes
  deriving Repr, DecidableEq, BEq, Inhabited

/-- `InnerSpec`: store-specific structure of inner nodes. -/
structure InnerSpec where
  childOrder : List Nat
  childSize : Int
  minPrefixLength : Int
  maxPrefixLength : Int
  emptyChild : Bytes
  hash : HashOp
  deriving Repr, DecidableEq, BEq, Inhabited

/-- `ProofSpec`: the full parameterization for a given merkle store. -/
structure ProofSpec where
  leafSpec : LeafOp
  innerSpec : InnerSpec
  minDepth : Int
  maxDepth : Int
  prehashKeyBeforeComparison : Bool
  deriving Repr, DecidableEq, BEq, Inhabited

/-- `ExistenceProof`: key/value plus the path of inner ops up to the root. -/
structure ExistenceProof where
  key : Bytes
  value : Bytes
  leaf : LeafOp
  path : List InnerOp
  deriving Repr, DecidableEq, BEq, Inhabited

/-- `NonExistenceProof`: the left and right neighbors bracketing an absent key. -/
structure NonExistenceProof where
  key : Bytes
  left : Option ExistenceProof
  right : Option ExistenceProof
  deriving Repr, DecidableEq, BEq, Inhabited

end Ics23
