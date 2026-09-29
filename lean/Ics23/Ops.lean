/-
Leaf and inner operations: the hashing core, mirroring `rust/src/ops.rs`.

The hash function is a parameter `HashFn`, never axiomatized. Soundness theorems
quantify over every `HashFn` and conclude by *exhibiting* a collision, so no
assumption about any concrete hash (SHA-256, etc.) is introduced. The only law
we ever require of a `HashFn` is that `noHash` is the identity, which is true of
the real `do_hash` and is passed explicitly where needed.
-/
import Ics23.Types

namespace Ics23

/-- An abstract hash family, keyed by `HashOp`. The real implementation is
`do_hash`; proofs treat it opaquely. -/
abbrev HashFn := HashOp → Bytes → Bytes

/-- Unsigned LEB128 varint encoding of a length, matching prost's
`encode_varint` as used by `LengthOp.varProto`. -/
def varintEncode (n : Nat) : List UInt8 :=
  if n < 0x80 then
    [UInt8.ofNat n]
  else
    UInt8.ofNat (n % 0x80 + 0x80) :: varintEncode (n / 0x80)
termination_by n
decreasing_by omega

/-- Big-endian 4-byte length. -/
def u32be (n : Nat) : List UInt8 :=
  [UInt8.ofNat (n >>> 24), UInt8.ofNat (n >>> 16), UInt8.ofNat (n >>> 8), UInt8.ofNat n]

/-- Big-endian 8-byte length. -/
def u64be (n : Nat) : List UInt8 :=
  [UInt8.ofNat (n >>> 56), UInt8.ofNat (n >>> 48), UInt8.ofNat (n >>> 40), UInt8.ofNat (n >>> 32),
   UInt8.ofNat (n >>> 24), UInt8.ofNat (n >>> 16), UInt8.ofNat (n >>> 8), UInt8.ofNat n]

/-- `do_length`: prepend (or require) a length encoding. Returns `none` for the
failing `Require*` cases, mirroring the Rust `ensure!`. -/
def doLength : LengthOp → Bytes → Option Bytes
  | .noPrefix, d => some d
  | .require32Bytes, d => if d.length = 32 then some d else none
  | .require64Bytes, d => if d.length = 64 then some d else none
  | .varProto, d => some (varintEncode d.length ++ d)
  | .fixed32Big, d => some (u32be d.length ++ d)
  | .fixed32Little, d => some ((u32be d.length).reverse ++ d)
  | .fixed64Big, d => some (u64be d.length ++ d)
  | .fixed64Little, d => some ((u64be d.length).reverse ++ d)

/-- `prepare_leaf_data`: prehash then length-encode. `none` on empty input,
matching the Rust `ensure!(!data.is_empty(), ...)`. -/
def prepareLeafData (H : HashFn) (prehash : HashOp) (length : LengthOp)
    (data : Bytes) : Option Bytes :=
  if data.isEmpty then none else doLength length (H prehash data)

/-- `apply_leaf`: `hash(prefix ++ encode(prehashKey key) ++ encode(prehashValue value))`. -/
def applyLeaf (H : HashFn) (leaf : LeafOp) (key value : Bytes) : Option Bytes :=
  match prepareLeafData H leaf.prehashKey leaf.length key,
        prepareLeafData H leaf.prehashValue leaf.length value with
  | some pk, some pv => some (H leaf.hash (leaf.prefixBytes ++ pk ++ pv))
  | _, _ => none

/-- `apply_inner`: `hash(prefix ++ child ++ suffix)`. `none` on empty child,
matching `ensure!(!child.is_empty(), ...)`. -/
def applyInner (H : HashFn) (inner : InnerOp) (child : Bytes) : Option Bytes :=
  if child.isEmpty then none
  else some (H inner.hash (inner.prefixBytes ++ child ++ inner.suffix))

end Ics23
