/-
The IAVL-specific prefix checks, mirroring `ensure_iavl_prefix`,
`ensure_inner_prefix`, and `ensure_leaf_prefix` in `rust/src/api.rs`. These were
abstracted away in `Verify.lean` (a sound omission for soundness). Here we model
them faithfully to test whether re-including them resolves finding F3
(positional ambiguity) for the IAVL spec.

The IAVL prefix encodes three signed (zigzag) protobuf varints — height, size,
version — and `ensure_inner_prefix` then requires the *remaining* byte count to
be 1 (this node is the left child: just its length byte follows) or 34 (this
node is the right child: length byte + 32-byte left sibling + length byte).
-/
import Ics23.Varint

namespace Ics23

/-- Zigzag-decode an unsigned varint value to a signed integer (prost's
`read_varint`: `x = ux >> 1`, negated when the low bit is set). -/
def zigzagDecode (u : Nat) : Int :=
  if u % 2 = 0 then (u / 2 : Int) else -((u / 2 : Int)) - 1

/-- Read one signed (zigzag) varint, returning the value and the rest. -/
def readVarintSigned (data : Bytes) : Option (Int × Bytes) :=
  match varintDecode data with
  | some (u, rest) => some (zigzagDecode u, rest)
  | none => none

/-- `ensure_iavl_prefix`: decode height/size/version with their bounds, returning
the number of bytes remaining after them (or `none` on a malformed prefix). -/
def ensureIavlPrefix (data : Bytes) (minHeight : Int) : Option Nat :=
  match readVarintSigned data with
  | none => none
  | some (height, r1) =>
    if height < minHeight then none
    else match readVarintSigned r1 with
      | none => none
      | some (size, r2) =>
        if size < 0 then none
        else match readVarintSigned r2 with
          | none => none
          | some (version, r3) => if version < 0 then none else some r3.length

/-- `ensure_inner_prefix` for the IAVL spec: the prefix decodes as height/size/
version and the remaining bytes are 1 (left child) or 34 (right child), with a
SHA-256 hash op. -/
def ensureInnerPrefixIavl (pre : Bytes) (minHeight : Int) (hashIsSha256 : Bool) : Bool :=
  match ensureIavlPrefix pre minHeight with
  | none => false
  | some remaining => (remaining = 1 || remaining = 34) && hashIsSha256

/-- `ensure_leaf_prefix` for the IAVL spec: prefix is exactly height/size/version
(no remaining bytes). -/
def ensureLeafPrefixIavl (pre : Bytes) : Bool :=
  match ensureIavlPrefix pre 0 with
  | none => false
  | some remaining => remaining = 0

/-! ## Does the IAVL prefix check resolve F3?

Take the IAVL node preimage `P = hsv ‖ 0x20 ‖ A ‖ 0x20 ‖ B`, where `hsv` is
height=1, size=2, version=3 (zigzag varints `[2,4,6]`) and `0x20 = 32` is the
child length byte. Read it two ways:

* **left child**: prefix `= hsv ‖ 0x20` (`[2,4,6,32]`), child `= A`, suffix `= 0x20 ‖ B`.
* **right child**: prefix `= hsv ‖ 0x20 ‖ A ‖ 0x20`, child `= B`, suffix `= []`.

If `ensure_inner_prefix` accepted only one, IAVL would pin the position. It does
not: the remaining-byte count is 1 for the left reading and 34 for the right
reading, and the check admits both. So F3 persists for IAVL. -/

def f3_hsv : Bytes := [2, 4, 6]
def f3_A : Bytes := List.replicate 32 0xAA
def f3_B : Bytes := List.replicate 32 0xBB

/-- Left reading's prefix passes (`remaining = 1`). -/
example : ensureInnerPrefixIavl (f3_hsv ++ [32]) 0 true = true := by native_decide

/-- Right reading's prefix passes (`remaining = 34`) — same node bytes, other
position. So the IAVL prefix check does NOT disambiguate the child. -/
example : ensureInnerPrefixIavl (f3_hsv ++ [32] ++ f3_A ++ [32]) 0 true = true := by native_decide

/-- The two readings share the same node preimage but assign different children. -/
example :
    (f3_hsv ++ [32]) ++ f3_A ++ ([32] ++ f3_B)
      = (f3_hsv ++ [32] ++ f3_A ++ [32]) ++ f3_B ++ ([] : Bytes) := by native_decide

end Ics23
