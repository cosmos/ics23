/-
The varint length prefix is self-delimiting: `varintEncode n ++ rest` parses
back to `(n, rest)`. This is the A1 ingredient for length-prefixed specs
(IAVL/Tendermint), the analogue of the fixed-length argument used for the SMT
shape. From it, `doLength .varProto` is injective in its data argument.
-/
import Ics23.Ops

namespace Ics23

/-- Decode a leading unsigned LEB128 varint, returning the value and the
remaining bytes. Inverse of `varintEncode` on the prefix. -/
def varintDecode : List UInt8 → Option (Nat × List UInt8)
  | [] => none
  | b :: rest =>
    if b.toNat < 0x80 then
      some (b.toNat, rest)
    else
      match varintDecode rest with
      | some (n, rest') => some (b.toNat % 0x80 + 0x80 * n, rest')
      | none => none

/-- Decoding the encoding of `n` followed by anything recovers `n` and the
remainder exactly. -/
theorem varintDecode_roundtrip (n : Nat) (rest : Bytes) :
    varintDecode (varintEncode n ++ rest) = some (n, rest) := by
  induction n using Nat.strongRecOn with
  | _ n ih =>
    rw [varintEncode]
    by_cases hn : n < 0x80
    · have ht : (UInt8.ofNat n).toNat = n := by simp; omega
      simp only [if_pos hn, List.cons_append, List.nil_append, varintDecode, ht]
    · have hb : (UInt8.ofNat (n % 0x80 + 0x80)).toNat = n % 0x80 + 0x80 := by
        simp; omega
      have hge : ¬ (UInt8.ofNat (n % 0x80 + 0x80)).toNat < 0x80 := by rw [hb]; omega
      have hlt : n / 0x80 < n := Nat.div_lt_self (by omega) (by omega)
      rw [if_neg hn]
      simp only [List.cons_append, varintDecode, if_neg hge, ih (n / 0x80) hlt,
        Option.some.injEq, Prod.mk.injEq, and_true]
      rw [hb]; omega

/-- `varintEncode` is self-delimiting under concatenation. -/
theorem varintEncode_append_inj (n m : Nat) (x y : Bytes)
    (h : varintEncode n ++ x = varintEncode m ++ y) : n = m ∧ x = y := by
  have h1 := varintDecode_roundtrip n x
  have h2 := varintDecode_roundtrip m y
  rw [h, h2] at h1
  simp only [Option.some.injEq, Prod.mk.injEq] at h1
  exact ⟨h1.1.symm, h1.2.symm⟩

/-- `doLength .varProto` is injective in its data. -/
theorem doLength_varProto_inj (a b : Bytes)
    (h : doLength .varProto a = doLength .varProto b) : a = b := by
  simp only [doLength, Option.some.injEq] at h
  exact (varintEncode_append_inj a.length b.length a b h).2

end Ics23
