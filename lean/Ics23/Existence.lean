/-
Existence binding (Theorem A), developed incrementally.

This file proves binding for the *same-shape* case: two existence proofs that
share the same leaf op and path ops, binding one key to two different values
under one root, yield a hash collision. It is parameterized by injectivity of
the leaf's length encoding (`hLeafInj`) and instantiated for both shipped leaf
shapes:

* `existence_binding_sameshape_noPrefix` — SMT/JMT (`NoPrefix` length).
* `existence_binding_sameshape_varProto` — IAVL / Tendermint (`VarProto` length).

The full chain is exercised: root extraction → path-fold injectivity
(`applyPath_sameops_inj`) → leaf-image cancellation → leaf-encoding injectivity
(`doLength_*_inj`). No `sorry` is used. The remaining gap to the *general*
`existence_binding` (in `Soundness.lean`) is the differing-path-structure case.
-/
import Ics23.Soundness
import Ics23.Varint

namespace Ics23

/-- A successful `verifyExistence` pins the computed root. -/
theorem verifyExistence_root (H : HashFn) (p : ExistenceProof) (s : ProofSpec)
    (root key value : Bytes) (h : verifyExistence H p s root key value = true) :
    calculateExistenceRoot H s p = some root := by
  unfold verifyExistence at h
  simp only [Bool.and_eq_true] at h
  obtain ⟨⟨⟨_, _⟩, _⟩, hr⟩ := h
  cases hc : calculateExistenceRoot H s p with
  | none => rw [hc] at hr; simp at hr
  | some r =>
    rw [hc] at hr
    simp only [Option.some.injEq]
    exact (eq_of_beq hr)

/-- `calculateExistenceRoot` factored into its leaf and path stages. -/
theorem calculateExistenceRoot_eq (H : HashFn) (s : ProofSpec) (p : ExistenceProof)
    (lh : Bytes) (hk : p.key.isEmpty = false) (hv : p.value.isEmpty = false)
    (hl : applyLeaf H p.leaf p.key p.value = some lh) :
    calculateExistenceRoot H s p = applyPath H s.innerSpec lh p.path := by
  simp [calculateExistenceRoot, hk, hv, hl]

/-- `doLength .noPrefix` is injective. -/
theorem doLength_noPrefix_inj (a b : Bytes)
    (h : doLength .noPrefix a = doLength .noPrefix b) : a = b := by
  simp only [doLength, Option.some.injEq] at h; exact h

/-- Leaf-value injectivity up to a collision: if two values encode to the same
leaf-value field under an injective length op, they are equal or their prehash
collides. -/
theorem prepareLeafData_inj (H : HashFn) (pre : HashOp) (L : LengthOp) (v₁ v₂ : Bytes)
    (hinj : ∀ a b, doLength L a = doLength L b → a = b)
    (h : prepareLeafData H pre L v₁ = prepareLeafData H pre L v₂)
    (hne1 : v₁.isEmpty = false) (hne2 : v₂.isEmpty = false) :
    v₁ = v₂ ∨ HashCollision H := by
  unfold prepareLeafData at h
  rw [if_neg (by simp [hne1]), if_neg (by simp [hne2])] at h
  have h' := hinj _ _ h
  by_cases hv : v₁ = v₂
  · exact Or.inl hv
  · exact Or.inr (hashCollision_of H pre v₁ v₂ hv h')

/-- Two different values binding to the same leaf hash under one leaf op force a
collision (the base case of binding). -/
theorem leaf_value_collision (H : HashFn) (leaf : LeafOp) (key v₁ v₂ L : Bytes)
    (hLeafInj : ∀ a b, doLength leaf.length a = doLength leaf.length b → a = b)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false) (hv : v₁ ≠ v₂)
    (hl1 : applyLeaf H leaf key v₁ = some L) (hl2 : applyLeaf H leaf key v₂ = some L) :
    HashCollision H := by
  unfold applyLeaf at hl1 hl2
  cases hpk : prepareLeafData H leaf.prehashKey leaf.length key with
  | none => simp [hpk] at hl1
  | some pk =>
  cases hpv1 : prepareLeafData H leaf.prehashValue leaf.length v₁ with
  | none => simp [hpk, hpv1] at hl1
  | some pv1 =>
  cases hpv2 : prepareLeafData H leaf.prehashValue leaf.length v₂ with
  | none => simp [hpk, hpv2] at hl2
  | some pv2 =>
  simp only [hpk, hpv1, Option.some.injEq] at hl1
  simp only [hpk, hpv2, Option.some.injEq] at hl2
  by_cases himg : leaf.prefixBytes ++ pk ++ pv1 = leaf.prefixBytes ++ pk ++ pv2
  · have hpveq : pv1 = pv2 := List.append_cancel_left himg
    have hpv : prepareLeafData H leaf.prehashValue leaf.length v₁
             = prepareLeafData H leaf.prehashValue leaf.length v₂ := by rw [hpv1, hpv2, hpveq]
    rcases prepareLeafData_inj H leaf.prehashValue leaf.length v₁ v₂ hLeafInj hpv hv1ne hv2ne
      with hveq | hcol
    · exact absurd hveq hv
    · exact hcol
  · exact hashCollision_of H leaf.hash _ _ himg (by rw [hl1, hl2])

/-- A leaf hash cannot also be an inner-node image: if it is, the leaf preimage
(starting with the leaf prefix byte) and the inner preimage (not) collide. The
length-mismatch case of binding. Requires the (shipped-spec) shape: a single-byte
leaf prefix, a shared hash op, and `min_prefix_length ≥ 1`. -/
theorem leafHash_innerImage_collision (H : HashFn) (s : ProofSpec)
    (leaf : LeafOp) (key val L : Bytes) (b : UInt8)
    (hsh : s.leafSpec.hash = s.innerSpec.hash)
    (hpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hleaf : ensureLeaf leaf s.leafSpec = true)
    (hl : applyLeaf H leaf key val = some L)
    (hii : IsInnerImage H s L) : HashCollision H := by
  obtain ⟨op, c, hin, hic⟩ := hii
  have hIimg := applyInner_image H op c L hic
  unfold applyLeaf at hl
  cases hpk : prepareLeafData H leaf.prehashKey leaf.length key with
  | none => simp [hpk] at hl
  | some pk =>
  cases hpv : prepareLeafData H leaf.prehashValue leaf.length val with
  | none => simp [hpk, hpv] at hl
  | some pv =>
  simp only [hpk, hpv, Option.some.injEq] at hl
  unfold ensureLeaf at hleaf
  simp only [Bool.and_eq_true] at hleaf
  obtain ⟨⟨⟨⟨hLhash, _⟩, _⟩, _⟩, hLpre⟩ := hleaf
  have hleafhash : leaf.hash = s.leafSpec.hash := (eq_of_beq hLhash).symm
  have hophash : op.hash = s.innerSpec.hash := ensureInner_hash op s hin
  have hLph : leaf.prefixBytes.head? = some b := by
    apply hasPrefix_single_head; rw [← hpre]; exact hLpre
  have hinx := hin
  unfold ensureInner at hinx
  simp only [Bool.and_eq_true] at hinx
  obtain ⟨⟨⟨⟨⟨⟨_, hnp⟩, hmp⟩, _⟩, _⟩, _⟩, _⟩ := hinx
  have hnp' : hasPrefix s.leafSpec.prefixBytes op.prefixBytes = false := by simpa using hnp
  have hopne : op.prefixBytes ≠ [] := by
    intro he
    rw [he] at hmp
    simp only [List.length_nil, decide_eq_true_eq] at hmp
    omega
  have hleafPre : (leaf.prefixBytes ++ pk ++ pv).head? = some b := by
    cases hlp : leaf.prefixBytes with
    | nil => rw [hlp] at hLph; simp at hLph
    | cons y ys =>
      rw [hlp] at hLph; simp only [List.head?_cons, Option.some.injEq] at hLph
      simp [hLph]
  have hinnerPre : (op.prefixBytes ++ c ++ op.suffix).head? ≠ some b := by
    cases hop : op.prefixBytes with
    | nil => exact absurd hop hopne
    | cons y ys =>
      have hh : hasPrefix [b] (y :: ys) = false := by rw [← hpre, ← hop]; exact hnp'
      have hne := not_hasPrefix_single_head b y ys hh
      simp only [List.cons_append, List.head?_cons] at hne ⊢
      exact hne
  refine leaf_inner_domain_collision H leaf.hash _ _ b hleafPre hinnerPre ?_
  rw [hl, hleafhash, hsh, ← hophash, hIimg]

/-- Every inner op of a verifying existence proof is spec-conformant. -/
theorem verifyExistence_inners (H : HashFn) (p : ExistenceProof) (s : ProofSpec)
    (root key value : Bytes) (h : verifyExistence H p s root key value = true) :
    ∀ op ∈ p.path, ensureInner op s = true := by
  unfold verifyExistence at h
  simp only [Bool.and_eq_true] at h
  obtain ⟨⟨⟨hces, _⟩, _⟩, _⟩ := h
  unfold checkExistenceSpec at hces
  simp only [Bool.and_eq_true] at hces
  obtain ⟨⟨_, _⟩, hpath⟩ := hces
  intro op hop
  exact List.all_eq_true.mp hpath op hop

/-- A verifying existence proof's leaf op is spec-conformant. -/
theorem verifyExistence_leaf (H : HashFn) (p : ExistenceProof) (s : ProofSpec)
    (root key value : Bytes) (h : verifyExistence H p s root key value = true) :
    ensureLeaf p.leaf s.leafSpec = true := by
  unfold verifyExistence at h
  simp only [Bool.and_eq_true] at h
  obtain ⟨⟨⟨hces, _⟩, _⟩, _⟩ := h
  unfold checkExistenceSpec at hces
  simp only [Bool.and_eq_true] at hces
  exact hces.1.1

/-- Two leaf ops conforming to the same leaf spec with equal prefixes are equal
(the spec pins every other field). -/
theorem ensureLeaf_eq (l1 l2 spec : LeafOp)
    (h1 : ensureLeaf l1 spec = true) (h2 : ensureLeaf l2 spec = true)
    (hp : l1.prefixBytes = l2.prefixBytes) : l1 = l2 := by
  unfold ensureLeaf at h1 h2
  simp only [Bool.and_eq_true, beq_iff_eq] at h1 h2
  obtain ⟨⟨⟨⟨e1h, e1pk⟩, e1pv⟩, e1l⟩, _⟩ := h1
  obtain ⟨⟨⟨⟨e2h, e2pk⟩, e2pv⟩, e2l⟩, _⟩ := h2
  cases l1; cases l2; simp_all

/-- **Theorem A, equal-length case.** Two existence proofs sharing the same leaf
op and the same path *length* (but possibly different inner ops), binding one key
to two different values under one root, yield a hash collision OR exhibit the F3
positional ambiguity. Strengthens the same-shape theorem (identical ops) to equal
depth with arbitrary ops; closes the honest disjunction for equal-depth proofs. -/
theorem existence_binding_eqlen
    (H : HashFn) (s : ProofSpec) (root key v₁ v₂ : Bytes)
    (p₁ p₂ : ExistenceProof)
    (hleafEq : p₁.leaf = p₂.leaf)
    (hlen : p₁.path.length = p₂.path.length)
    (hLeafInj : ∀ a b, doLength p₁.leaf.length a = doLength p₁.leaf.length b → a = b)
    (hk1 : p₁.key = key) (hk2 : p₂.key = key)
    (hkne : key.isEmpty = false)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false)
    (hvv1 : p₁.value = v₁) (hvv2 : p₂.value = v₂)
    (hv : v₁ ≠ v₂)
    (h₁ : verifyExistence H p₁ s root key v₁ = true)
    (h₂ : verifyExistence H p₂ s root key v₂ = true) :
    HashCollision H ∨ PositionalAmbiguity s := by
  have r1 := verifyExistence_root H p₁ s root key v₁ h₁
  have r2 := verifyExistence_root H p₂ s root key v₂ h₂
  have hk1e : p₁.key.isEmpty = false := by rw [hk1]; exact hkne
  have hk2e : p₂.key.isEmpty = false := by rw [hk2]; exact hkne
  have hv1e : p₁.value.isEmpty = false := by rw [hvv1]; exact hv1ne
  have hv2e : p₂.value.isEmpty = false := by rw [hvv2]; exact hv2ne
  cases hl1 : applyLeaf H p₁.leaf p₁.key p₁.value with
  | none => simp [calculateExistenceRoot, hk1e, hv1e, hl1] at r1
  | some lh₁ =>
  cases hl2 : applyLeaf H p₂.leaf p₂.key p₂.value with
  | none => simp [calculateExistenceRoot, hk2e, hv2e, hl2] at r2
  | some lh₂ =>
  have e1 : applyPath H s.innerSpec lh₁ p₁.path = some root := by
    rw [← calculateExistenceRoot_eq H s p₁ lh₁ hk1e hv1e hl1]; exact r1
  have e2 : applyPath H s.innerSpec lh₂ p₂.path = some root := by
    rw [← calculateExistenceRoot_eq H s p₂ lh₂ hk2e hv2e hl2]; exact r2
  rcases applyPath_eqlen_merge H s p₁.path p₂.path hlen lh₁ lh₂ root
      (verifyExistence_inners H p₁ s root key v₁ h₁)
      (verifyExistence_inners H p₂ s root key v₂ h₂) e1 e2 with hc | ha | hlh
  · exact Or.inl hc
  · exact Or.inr ha
  · -- lh₁ = lh₂: same leaf step as same-shape ⇒ collision or v₁ = v₂ (contra)
    refine Or.inl ?_
    rw [hk1, hvv1] at hl1
    rw [hk2, hvv2, ← hleafEq] at hl2
    unfold applyLeaf at hl1 hl2
    cases hpk : prepareLeafData H p₁.leaf.prehashKey p₁.leaf.length key with
    | none => simp [hpk] at hl1
    | some pk =>
    cases hpv1 : prepareLeafData H p₁.leaf.prehashValue p₁.leaf.length v₁ with
    | none => simp [hpk, hpv1] at hl1
    | some pv1 =>
    cases hpv2 : prepareLeafData H p₁.leaf.prehashValue p₁.leaf.length v₂ with
    | none => simp [hpk, hpv2] at hl2
    | some pv2 =>
    simp only [hpk, hpv1, Option.some.injEq] at hl1
    simp only [hpk, hpv2, Option.some.injEq] at hl2
    by_cases himg :
        p₁.leaf.prefixBytes ++ pk ++ pv1 = p₁.leaf.prefixBytes ++ pk ++ pv2
    · have hpveq : pv1 = pv2 := List.append_cancel_left himg
      have hpv : prepareLeafData H p₁.leaf.prehashValue p₁.leaf.length v₁
               = prepareLeafData H p₁.leaf.prehashValue p₁.leaf.length v₂ := by
        rw [hpv1, hpv2, hpveq]
      rcases prepareLeafData_inj H p₁.leaf.prehashValue p₁.leaf.length v₁ v₂
          hLeafInj hpv hv1ne hv2ne with hveq | hcol
      · exact absurd hveq hv
      · exact hcol
    · exact hashCollision_of H p₁.leaf.hash _ _ himg (by rw [hl1, hl2]; exact hlh)

/-- **Theorem A, same-shape case (general length op).** Two existence proofs
sharing the same leaf op and path ops, binding the same (nonempty) key to two
different (nonempty) values under one root, yield a hash collision — provided the
leaf's length encoding is injective. -/
theorem existence_binding_sameshape
    (H : HashFn) (s : ProofSpec) (root key v₁ v₂ : Bytes)
    (p₁ p₂ : ExistenceProof)
    (hleafEq : p₁.leaf = p₂.leaf)
    (hpathEq : p₁.path = p₂.path)
    (hLeafInj : ∀ a b, doLength p₁.leaf.length a = doLength p₁.leaf.length b → a = b)
    (hk1 : p₁.key = key) (hk2 : p₂.key = key)
    (hkne : key.isEmpty = false)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false)
    (hvv1 : p₁.value = v₁) (hvv2 : p₂.value = v₂)
    (hv : v₁ ≠ v₂)
    (h₁ : verifyExistence H p₁ s root key v₁ = true)
    (h₂ : verifyExistence H p₂ s root key v₂ = true) :
    HashCollision H := by
  have r1 := verifyExistence_root H p₁ s root key v₁ h₁
  have r2 := verifyExistence_root H p₂ s root key v₂ h₂
  have hk1e : p₁.key.isEmpty = false := by rw [hk1]; exact hkne
  have hk2e : p₂.key.isEmpty = false := by rw [hk2]; exact hkne
  have hv1e : p₁.value.isEmpty = false := by rw [hvv1]; exact hv1ne
  have hv2e : p₂.value.isEmpty = false := by rw [hvv2]; exact hv2ne
  cases hl1 : applyLeaf H p₁.leaf p₁.key p₁.value with
  | none => simp [calculateExistenceRoot, hk1e, hv1e, hl1] at r1
  | some lh₁ =>
  cases hl2 : applyLeaf H p₂.leaf p₂.key p₂.value with
  | none => simp [calculateExistenceRoot, hk2e, hv2e, hl2] at r2
  | some lh₂ =>
  have e1 : applyPath H s.innerSpec lh₁ p₁.path = some root := by
    rw [← calculateExistenceRoot_eq H s p₁ lh₁ hk1e hv1e hl1]; exact r1
  have e2 : applyPath H s.innerSpec lh₂ p₂.path = some root := by
    rw [← calculateExistenceRoot_eq H s p₂ lh₂ hk2e hv2e hl2]; exact r2
  rw [hpathEq] at e1
  rcases applyPath_sameops_inj H s.innerSpec p₂.path lh₁ lh₂ root e1 e2 with hlh | hc
  · rw [hk1, hvv1] at hl1
    rw [hk2, hvv2, ← hleafEq] at hl2
    unfold applyLeaf at hl1 hl2
    cases hpk : prepareLeafData H p₁.leaf.prehashKey p₁.leaf.length key with
    | none => simp [hpk] at hl1
    | some pk =>
    cases hpv1 : prepareLeafData H p₁.leaf.prehashValue p₁.leaf.length v₁ with
    | none => simp [hpk, hpv1] at hl1
    | some pv1 =>
    cases hpv2 : prepareLeafData H p₁.leaf.prehashValue p₁.leaf.length v₂ with
    | none => simp [hpk, hpv2] at hl2
    | some pv2 =>
    simp only [hpk, hpv1, Option.some.injEq] at hl1
    simp only [hpk, hpv2, Option.some.injEq] at hl2
    by_cases himg :
        p₁.leaf.prefixBytes ++ pk ++ pv1 = p₁.leaf.prefixBytes ++ pk ++ pv2
    · have hpveq : pv1 = pv2 := List.append_cancel_left himg
      have hpv : prepareLeafData H p₁.leaf.prehashValue p₁.leaf.length v₁
               = prepareLeafData H p₁.leaf.prehashValue p₁.leaf.length v₂ := by
        rw [hpv1, hpv2, hpveq]
      rcases prepareLeafData_inj H p₁.leaf.prehashValue p₁.leaf.length v₁ v₂
          hLeafInj hpv hv1ne hv2ne with hveq | hcol
      · exact absurd hveq hv
      · exact hcol
    · exact hashCollision_of H p₁.leaf.hash _ _ himg (by rw [hl1, hl2]; exact hlh)
  · exact hc

/-- **Theorem A, SMT/JMT same-shape** (`NoPrefix` length). -/
theorem existence_binding_sameshape_noPrefix
    (H : HashFn) (s : ProofSpec) (root key v₁ v₂ : Bytes)
    (p₁ p₂ : ExistenceProof)
    (hleafEq : p₁.leaf = p₂.leaf) (hpathEq : p₁.path = p₂.path)
    (hLen : p₁.leaf.length = .noPrefix)
    (hk1 : p₁.key = key) (hk2 : p₂.key = key)
    (hkne : key.isEmpty = false)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false)
    (hvv1 : p₁.value = v₁) (hvv2 : p₂.value = v₂) (hv : v₁ ≠ v₂)
    (h₁ : verifyExistence H p₁ s root key v₁ = true)
    (h₂ : verifyExistence H p₂ s root key v₂ = true) :
    HashCollision H :=
  existence_binding_sameshape H s root key v₁ v₂ p₁ p₂ hleafEq hpathEq
    (fun a b h => by rw [hLen] at h; exact doLength_noPrefix_inj a b h)
    hk1 hk2 hkne hv1ne hv2ne hvv1 hvv2 hv h₁ h₂

/-- **Theorem A, IAVL / Tendermint same-shape** (`VarProto` length). -/
theorem existence_binding_sameshape_varProto
    (H : HashFn) (s : ProofSpec) (root key v₁ v₂ : Bytes)
    (p₁ p₂ : ExistenceProof)
    (hleafEq : p₁.leaf = p₂.leaf) (hpathEq : p₁.path = p₂.path)
    (hLen : p₁.leaf.length = .varProto)
    (hk1 : p₁.key = key) (hk2 : p₂.key = key)
    (hkne : key.isEmpty = false)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false)
    (hvv1 : p₁.value = v₁) (hvv2 : p₂.value = v₂) (hv : v₁ ≠ v₂)
    (h₁ : verifyExistence H p₁ s root key v₁ = true)
    (h₂ : verifyExistence H p₂ s root key v₂ = true) :
    HashCollision H :=
  existence_binding_sameshape H s root key v₁ v₂ p₁ p₂ hleafEq hpathEq
    (fun a b h => by rw [hLen] at h; exact doLength_varProto_inj a b h)
    hk1 hk2 hkne hv1ne hv2ne hvv1 hvv2 hv h₁ h₂

/-- **Theorem A, same-leaf case (arbitrary path length).** For a spec with a
single-byte leaf prefix, a shared leaf/inner hash op, and `min_prefix_length ≥ 1`
(all true for IAVL/Tendermint/SMT), two existence proofs sharing the same leaf op
— but with *arbitrary, differing-length* paths — that bind one key to two values
under one root yield a hash collision OR the F3 positional ambiguity. This is the
strongest binding result: it removes the equal-length restriction by handling the
length-mismatch via leaf/inner domain separation. -/
theorem existence_binding_sameleaf
    (H : HashFn) (s : ProofSpec) (root key v₁ v₂ : Bytes)
    (p₁ p₂ : ExistenceProof) (b : UInt8)
    (hsh : s.leafSpec.hash = s.innerSpec.hash)
    (hpreShape : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hleafEq : p₁.leaf = p₂.leaf)
    (hLeafInj : ∀ a b, doLength p₁.leaf.length a = doLength p₁.leaf.length b → a = b)
    (hk1 : p₁.key = key) (hk2 : p₂.key = key)
    (hkne : key.isEmpty = false)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false)
    (hvv1 : p₁.value = v₁) (hvv2 : p₂.value = v₂)
    (hv : v₁ ≠ v₂)
    (h₁ : verifyExistence H p₁ s root key v₁ = true)
    (h₂ : verifyExistence H p₂ s root key v₂ = true) :
    HashCollision H ∨ PositionalAmbiguity s := by
  have r1 := verifyExistence_root H p₁ s root key v₁ h₁
  have r2 := verifyExistence_root H p₂ s root key v₂ h₂
  have hk1e : p₁.key.isEmpty = false := by rw [hk1]; exact hkne
  have hk2e : p₂.key.isEmpty = false := by rw [hk2]; exact hkne
  have hv1e : p₁.value.isEmpty = false := by rw [hvv1]; exact hv1ne
  have hv2e : p₂.value.isEmpty = false := by rw [hvv2]; exact hv2ne
  cases hl1 : applyLeaf H p₁.leaf p₁.key p₁.value with
  | none => simp [calculateExistenceRoot, hk1e, hv1e, hl1] at r1
  | some lh₁ =>
  cases hl2 : applyLeaf H p₂.leaf p₂.key p₂.value with
  | none => simp [calculateExistenceRoot, hk2e, hv2e, hl2] at r2
  | some lh₂ =>
  have e1 : applyPath H s.innerSpec lh₁ p₁.path = some root := by
    rw [← calculateExistenceRoot_eq H s p₁ lh₁ hk1e hv1e hl1]; exact r1
  have e2 : applyPath H s.innerSpec lh₂ p₂.path = some root := by
    rw [← calculateExistenceRoot_eq H s p₂ lh₂ hk2e hv2e hl2]; exact r2
  -- normalize the leaf-hash facts to (key, vᵢ)
  rw [hk1, hvv1] at hl1
  rw [hk2, hvv2] at hl2
  rcases applyPath_merge H s (p₁.path.length + p₂.path.length) p₁.path p₂.path lh₁ lh₂ root rfl
      (verifyExistence_inners H p₁ s root key v₁ h₁)
      (verifyExistence_inners H p₂ s root key v₂ h₂) e1 e2 with hc | ha | hlheq | hii1 | hii2
  · exact Or.inl hc
  · exact Or.inr ha
  · -- lh₁ = lh₂ : leaf injectivity
    refine Or.inl ?_
    rw [hlheq] at hl1
    rw [← hleafEq] at hl2
    exact leaf_value_collision H p₁.leaf key v₁ v₂ lh₂ hLeafInj hv1ne hv2ne hv hl1 hl2
  · -- lh₁ is an inner image: leaf/inner domain separation
    exact Or.inl (leafHash_innerImage_collision H s p₁.leaf key v₁ lh₁ b hsh hpreShape hmin
      (verifyExistence_leaf H p₁ s root key v₁ h₁) hl1 hii1)
  · -- lh₂ is an inner image
    exact Or.inl (leafHash_innerImage_collision H s p₂.leaf key v₂ lh₂ b hsh hpreShape hmin
      (verifyExistence_leaf H p₂ s root key v₂ h₂) hl2 hii2)

/-- Same-leaf binding instantiated for the **IAVL** spec. -/
theorem existence_binding_iavl
    (H : HashFn) (root key v₁ v₂ : Bytes) (p₁ p₂ : ExistenceProof)
    (hleafEq : p₁.leaf = p₂.leaf) (hLen : p₁.leaf.length = .varProto)
    (hk1 : p₁.key = key) (hk2 : p₂.key = key) (hkne : key.isEmpty = false)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false)
    (hvv1 : p₁.value = v₁) (hvv2 : p₂.value = v₂) (hv : v₁ ≠ v₂)
    (h₁ : verifyExistence H p₁ iavlSpec root key v₁ = true)
    (h₂ : verifyExistence H p₂ iavlSpec root key v₂ = true) :
    HashCollision H ∨ PositionalAmbiguity iavlSpec :=
  existence_binding_sameleaf H iavlSpec root key v₁ v₂ p₁ p₂ 0
    (by decide) (by decide) (by decide) hleafEq
    (fun a b h => by rw [hLen] at h; exact doLength_varProto_inj a b h)
    hk1 hk2 hkne hv1ne hv2ne hvv1 hvv2 hv h₁ h₂

/-- Same-leaf binding instantiated for the **Tendermint** spec. -/
theorem existence_binding_tendermint
    (H : HashFn) (root key v₁ v₂ : Bytes) (p₁ p₂ : ExistenceProof)
    (hleafEq : p₁.leaf = p₂.leaf) (hLen : p₁.leaf.length = .varProto)
    (hk1 : p₁.key = key) (hk2 : p₂.key = key) (hkne : key.isEmpty = false)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false)
    (hvv1 : p₁.value = v₁) (hvv2 : p₂.value = v₂) (hv : v₁ ≠ v₂)
    (h₁ : verifyExistence H p₁ tendermintSpec root key v₁ = true)
    (h₂ : verifyExistence H p₂ tendermintSpec root key v₂ = true) :
    HashCollision H ∨ PositionalAmbiguity tendermintSpec :=
  existence_binding_sameleaf H tendermintSpec root key v₁ v₂ p₁ p₂ 0
    (by decide) (by decide) (by decide) hleafEq
    (fun a b h => by rw [hLen] at h; exact doLength_varProto_inj a b h)
    hk1 hk2 hkne hv1ne hv2ne hvv1 hvv2 hv h₁ h₂

/-- Same-leaf binding instantiated for the **SMT** spec (`NoPrefix` length). -/
theorem existence_binding_smt
    (H : HashFn) (root key v₁ v₂ : Bytes) (p₁ p₂ : ExistenceProof)
    (hleafEq : p₁.leaf = p₂.leaf) (hLen : p₁.leaf.length = .noPrefix)
    (hk1 : p₁.key = key) (hk2 : p₂.key = key) (hkne : key.isEmpty = false)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false)
    (hvv1 : p₁.value = v₁) (hvv2 : p₂.value = v₂) (hv : v₁ ≠ v₂)
    (h₁ : verifyExistence H p₁ smtSpec root key v₁ = true)
    (h₂ : verifyExistence H p₂ smtSpec root key v₂ = true) :
    HashCollision H ∨ PositionalAmbiguity smtSpec :=
  existence_binding_sameleaf H smtSpec root key v₁ v₂ p₁ p₂ 0
    (by decide) (by decide) (by decide) hleafEq
    (fun a b h => by rw [hLen] at h; exact doLength_noPrefix_inj a b h)
    hk1 hk2 hkne hv1ne hv2ne hvv1 hvv2 hv h₁ h₂

/-- **Theorem A, general (production-spec shape).** Two existence proofs binding
one key to two different values under one root — with NO assumption relating their
leaf ops or paths — yield the honest three-way disjunction: a hash collision, the
inner positional ambiguity (F3), or its leaf-level analogue. Proved with no
`sorry` for the shape shared by IAVL/Tendermint/SMT (single-byte leaf prefix,
shared leaf/inner hash op, `min_prefix_length ≥ 1`). The two ambiguity arms are
real, machine-checkable obstructions that the abstract-hash model cannot rule out;
collapsing them to a bare collision requires the symbolic-Merkle model. -/
theorem existence_binding_shaped
    (H : HashFn) (s : ProofSpec) (root key v₁ v₂ : Bytes)
    (p₁ p₂ : ExistenceProof) (b : UInt8)
    (hsh : s.leafSpec.hash = s.innerSpec.hash)
    (hpreShape : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hLeafInj : ∀ a b, doLength p₁.leaf.length a = doLength p₁.leaf.length b → a = b)
    (hk1 : p₁.key = key) (hk2 : p₂.key = key)
    (hkne : key.isEmpty = false)
    (hv1ne : v₁.isEmpty = false) (hv2ne : v₂.isEmpty = false)
    (hvv1 : p₁.value = v₁) (hvv2 : p₂.value = v₂)
    (hv : v₁ ≠ v₂)
    (h₁ : verifyExistence H p₁ s root key v₁ = true)
    (h₂ : verifyExistence H p₂ s root key v₂ = true) :
    HashCollision H ∨ PositionalAmbiguity s ∨ LeafAmbiguity H s := by
  have r1 := verifyExistence_root H p₁ s root key v₁ h₁
  have r2 := verifyExistence_root H p₂ s root key v₂ h₂
  have hk1e : p₁.key.isEmpty = false := by rw [hk1]; exact hkne
  have hk2e : p₂.key.isEmpty = false := by rw [hk2]; exact hkne
  have hv1e : p₁.value.isEmpty = false := by rw [hvv1]; exact hv1ne
  have hv2e : p₂.value.isEmpty = false := by rw [hvv2]; exact hv2ne
  cases hl1 : applyLeaf H p₁.leaf p₁.key p₁.value with
  | none => simp [calculateExistenceRoot, hk1e, hv1e, hl1] at r1
  | some lh₁ =>
  cases hl2 : applyLeaf H p₂.leaf p₂.key p₂.value with
  | none => simp [calculateExistenceRoot, hk2e, hv2e, hl2] at r2
  | some lh₂ =>
  have e1 : applyPath H s.innerSpec lh₁ p₁.path = some root := by
    rw [← calculateExistenceRoot_eq H s p₁ lh₁ hk1e hv1e hl1]; exact r1
  have e2 : applyPath H s.innerSpec lh₂ p₂.path = some root := by
    rw [← calculateExistenceRoot_eq H s p₂ lh₂ hk2e hv2e hl2]; exact r2
  rw [hk1, hvv1] at hl1
  rw [hk2, hvv2] at hl2
  rcases applyPath_merge H s (p₁.path.length + p₂.path.length) p₁.path p₂.path lh₁ lh₂ root rfl
      (verifyExistence_inners H p₁ s root key v₁ h₁)
      (verifyExistence_inners H p₂ s root key v₂ h₂) e1 e2 with hc | ha | hlheq | hii1 | hii2
  · exact Or.inl hc
  · exact Or.inr (Or.inl ha)
  · by_cases hpeq : p₁.leaf.prefixBytes = p₂.leaf.prefixBytes
    · have hleq : p₁.leaf = p₂.leaf :=
        ensureLeaf_eq p₁.leaf p₂.leaf s.leafSpec
          (verifyExistence_leaf H p₁ s root key v₁ h₁)
          (verifyExistence_leaf H p₂ s root key v₂ h₂) hpeq
      refine Or.inl ?_
      rw [hlheq] at hl1
      rw [← hleq] at hl2
      exact leaf_value_collision H p₁.leaf key v₁ v₂ lh₂ hLeafInj hv1ne hv2ne hv hl1 hl2
    · refine Or.inr (Or.inr ⟨p₁.leaf, p₂.leaf, key, v₁, key, v₂, lh₂,
        verifyExistence_leaf H p₁ s root key v₁ h₁,
        verifyExistence_leaf H p₂ s root key v₂ h₂, hpeq, ?_, hl2⟩)
      rw [hlheq] at hl1; exact hl1
  · exact Or.inl (leafHash_innerImage_collision H s p₁.leaf key v₁ lh₁ b hsh hpreShape hmin
      (verifyExistence_leaf H p₁ s root key v₁ h₁) hl1 hii1)
  · exact Or.inl (leafHash_innerImage_collision H s p₂.leaf key v₂ lh₂ b hsh hpreShape hmin
      (verifyExistence_leaf H p₂ s root key v₂ h₂) hl2 hii2)

end Ics23
