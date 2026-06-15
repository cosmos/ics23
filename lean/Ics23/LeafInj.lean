/-
Joint leaf injectivity (`hLInj`) discharged from the leaf encoding.

Theorems A and B both take a hypothesis `hLInj`: two leaves hashing to the same
value share their `(key, value)` — unless a hash collision is exhibited. The
honest-root proofs assumed it; here we *prove* it for the length-prefixed
(`varProto`) leaf shape used by the IAVL and Tendermint specs, with no axioms.

The reduction is exactly "varint self-delimiting + collision resistance":
  apply_leaf k v = H_hash (prefix ++ frame(H_phk k) ++ frame(H_phv v))
with `frame x = varintEncode |x| ++ x`. Equal images give either a collision on
the leaf `hash` op (differing preimages) or equal preimages; the latter cancels
the fixed `prefix`, then `varintEncode_append_inj` + `List.append_inj` peel the
two frames, leaving `H_phk k₁ = H_phk k₂` and `H_phv v₁ = H_phv v₂` — each of
which is either an equality of keys/values or a collision on the prehash op.
-/
import Ics23.Ops
import Ics23.Varint
import Ics23.Specs
import Ics23.Soundness

namespace Ics23

/-- The explicit preimage of a `varProto`-framed leaf hash. -/
theorem applyLeaf_varProto_form (H : HashFn) (ls : LeafOp) (hlen : ls.length = .varProto)
    (k v r : Bytes) (h : applyLeaf H ls k v = some r) :
    r = H ls.hash (ls.prefixBytes
        ++ (varintEncode (H ls.prehashKey k).length ++ H ls.prehashKey k)
        ++ (varintEncode (H ls.prehashValue v).length ++ H ls.prehashValue v)) := by
  unfold applyLeaf at h
  rw [hlen] at h
  cases hke : prepareLeafData H ls.prehashKey .varProto k with
  | none => rw [hke] at h; simp at h
  | some pk =>
  cases hve : prepareLeafData H ls.prehashValue .varProto v with
  | none => rw [hke, hve] at h; simp at h
  | some pv =>
    rw [hke, hve] at h
    simp only [Option.some.injEq] at h
    have hpk : pk = varintEncode (H ls.prehashKey k).length ++ H ls.prehashKey k := by
      unfold prepareLeafData at hke
      by_cases hkemp : k.isEmpty
      · rw [if_pos hkemp] at hke; exact absurd hke (by simp)
      · rw [if_neg hkemp] at hke; simp only [doLength, Option.some.injEq] at hke; exact hke.symm
    have hpv : pv = varintEncode (H ls.prehashValue v).length ++ H ls.prehashValue v := by
      unfold prepareLeafData at hve
      by_cases hvemp : v.isEmpty
      · rw [if_pos hvemp] at hve; exact absurd hve (by simp)
      · rw [if_neg hvemp] at hve; simp only [doLength, Option.some.injEq] at hve; exact hve.symm
    rw [← h, hpk, hpv]

/-- **Joint leaf injectivity for any `varProto`-framed leaf op.** Two leaves
hashing to the same value share their key and value, or a hash collision is
exhibited. No assumption on `H` beyond what a collision witnesses. -/
theorem leafInj_varProto (H : HashFn) (ls : LeafOp) (hlen : ls.length = .varProto) :
    ∀ k₁ v₁ k₂ v₂ r,
      applyLeaf H ls k₁ v₁ = some r →
      applyLeaf H ls k₂ v₂ = some r →
      (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H := by
  intro k₁ v₁ k₂ v₂ r h1 h2
  have e1 := applyLeaf_varProto_form H ls hlen k₁ v₁ r h1
  have e2 := applyLeaf_varProto_form H ls hlen k₂ v₂ r h2
  by_cases hP :
      (ls.prefixBytes ++ (varintEncode (H ls.prehashKey k₁).length ++ H ls.prehashKey k₁)
        ++ (varintEncode (H ls.prehashValue v₁).length ++ H ls.prehashValue v₁))
      = (ls.prefixBytes ++ (varintEncode (H ls.prehashKey k₂).length ++ H ls.prehashKey k₂)
        ++ (varintEncode (H ls.prehashValue v₂).length ++ H ls.prehashValue v₂))
  · -- equal preimages: cancel the prefix, then peel the two varint frames
    simp only [List.append_assoc] at hP
    have hP' := List.append_cancel_left hP
    obtain ⟨hklen, hkrest⟩ := varintEncode_append_inj (H ls.prehashKey k₁).length
      (H ls.prehashKey k₂).length _ _ hP'
    obtain ⟨hkeq, hveqframe⟩ := List.append_inj hkrest hklen
    obtain ⟨_, hveq⟩ := varintEncode_append_inj (H ls.prehashValue v₁).length
      (H ls.prehashValue v₂).length _ _ hveqframe
    by_cases hk : k₁ = k₂
    · by_cases hv : v₁ = v₂
      · exact Or.inl ⟨hk, hv⟩
      · exact Or.inr (hashCollision_of H ls.prehashValue v₁ v₂ hv hveq)
    · exact Or.inr (hashCollision_of H ls.prehashKey k₁ k₂ hk hkeq)
  · -- differing preimages with equal image: a collision on the leaf hash op
    exact Or.inr (hashCollision_of H ls.hash _ _ hP (e1.symm.trans e2))

/-- Joint leaf injectivity for the Tendermint spec (its leaf op is `varProto`). -/
theorem leafInj_tendermint (H : HashFn) :
    ∀ k₁ v₁ k₂ v₂ r,
      applyLeaf H tendermintSpec.leafSpec k₁ v₁ = some r →
      applyLeaf H tendermintSpec.leafSpec k₂ v₂ = some r →
      (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H :=
  leafInj_varProto H tendermintSpec.leafSpec (by rfl)

/-- Joint leaf injectivity for the IAVL spec (identical `varProto` leaf op). -/
theorem leafInj_iavl (H : HashFn) :
    ∀ k₁ v₁ k₂ v₂ r,
      applyLeaf H iavlSpec.leafSpec k₁ v₁ = some r →
      applyLeaf H iavlSpec.leafSpec k₂ v₂ = some r →
      (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H :=
  leafInj_varProto H iavlSpec.leafSpec (by rfl)

/-! ## The SMT / `noPrefix` shape: fixed-length split

The SMT spec uses `length := .noPrefix` (no varint frame) but prehashes both key
and value, so the leaf hash is `H_hash(prefix ++ H_phk k ++ H_phv v)` with the two
prehash digests each `cs` bytes wide (`FixedHash`). The fixed width is what makes
the concatenation uniquely splittable — the analogue of the varint frame above. -/

/-- The explicit preimage of a `noPrefix` leaf hash. -/
theorem applyLeaf_noPrefix_form (H : HashFn) (ls : LeafOp) (hlen : ls.length = .noPrefix)
    (k v r : Bytes) (h : applyLeaf H ls k v = some r) :
    r = H ls.hash (ls.prefixBytes ++ H ls.prehashKey k ++ H ls.prehashValue v) := by
  unfold applyLeaf at h
  rw [hlen] at h
  cases hke : prepareLeafData H ls.prehashKey .noPrefix k with
  | none => rw [hke] at h; simp at h
  | some pk =>
  cases hve : prepareLeafData H ls.prehashValue .noPrefix v with
  | none => rw [hke, hve] at h; simp at h
  | some pv =>
    rw [hke, hve] at h
    simp only [Option.some.injEq] at h
    have hpk : pk = H ls.prehashKey k := by
      unfold prepareLeafData at hke
      by_cases hkemp : k.isEmpty
      · rw [if_pos hkemp] at hke; exact absurd hke (by simp)
      · rw [if_neg hkemp] at hke; simp only [doLength, Option.some.injEq] at hke; exact hke.symm
    have hpv : pv = H ls.prehashValue v := by
      unfold prepareLeafData at hve
      by_cases hvemp : v.isEmpty
      · rw [if_pos hvemp] at hve; exact absurd hve (by simp)
      · rw [if_neg hvemp] at hve; simp only [doLength, Option.some.injEq] at hve; exact hve.symm
    rw [← h, hpk, hpv]

/-- **Joint leaf injectivity for any `noPrefix` leaf op over a fixed-length hash.**
Both prehash digests are `cs` bytes wide, so the concatenation splits at the fixed
boundary; the rest mirrors `leafInj_varProto`. -/
theorem leafInj_noPrefix (H : HashFn) (cs : Nat) (hH : FixedHash H cs)
    (ls : LeafOp) (hlen : ls.length = .noPrefix) :
    ∀ k₁ v₁ k₂ v₂ r,
      applyLeaf H ls k₁ v₁ = some r →
      applyLeaf H ls k₂ v₂ = some r →
      (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H := by
  intro k₁ v₁ k₂ v₂ r h1 h2
  have e1 := applyLeaf_noPrefix_form H ls hlen k₁ v₁ r h1
  have e2 := applyLeaf_noPrefix_form H ls hlen k₂ v₂ r h2
  by_cases hP :
      (ls.prefixBytes ++ H ls.prehashKey k₁ ++ H ls.prehashValue v₁)
      = (ls.prefixBytes ++ H ls.prehashKey k₂ ++ H ls.prehashValue v₂)
  · -- equal preimages: cancel the prefix, then split at the fixed digest width
    simp only [List.append_assoc] at hP
    have hP' := List.append_cancel_left hP
    have hlenK : (H ls.prehashKey k₁).length = (H ls.prehashKey k₂).length :=
      (hH ls.prehashKey k₁).trans (hH ls.prehashKey k₂).symm
    obtain ⟨hkeq, hveq⟩ := List.append_inj hP' hlenK
    by_cases hk : k₁ = k₂
    · by_cases hv : v₁ = v₂
      · exact Or.inl ⟨hk, hv⟩
      · exact Or.inr (hashCollision_of H ls.prehashValue v₁ v₂ hv hveq)
    · exact Or.inr (hashCollision_of H ls.prehashKey k₁ k₂ hk hkeq)
  · exact Or.inr (hashCollision_of H ls.hash _ _ hP (e1.symm.trans e2))

/-- Joint leaf injectivity for the SMT/JMT spec (`noPrefix`, prehashed key and
value), over a fixed 32-byte hash. -/
theorem leafInj_smt (H : HashFn) (hH : FixedHash H 32) :
    ∀ k₁ v₁ k₂ v₂ r,
      applyLeaf H smtSpec.leafSpec k₁ v₁ = some r →
      applyLeaf H smtSpec.leafSpec k₂ v₂ = some r →
      (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H :=
  leafInj_noPrefix H 32 hH smtSpec.leafSpec (by rfl)

end Ics23
