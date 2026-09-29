/-
Soundness scaffolding for the existence verifier.

This file contains:
  * `HashCollision` — the constructive target of every soundness theorem.
  * `WellFormed` — the decidable side condition on a `ProofSpec` under which
    the verifier is sound (Theorem C certifies the shipped specs satisfy it).
  * supporting lemmas that are fully proved.
  * Theorem A (`existence_binding`) — stated precisely, proof in progress.

Theorem A is the headline soundness property. Its proof is staged: the
structural lemmas it rests on are being landed first; the top-level statement is
recorded now (with `sorry`) so the target is fixed and reviewable. No `sorry`
is relied upon by `WellFormed`, the spec certificates, or the proved lemmas.
-/
import Ics23.Verify
import Ics23.Specs

namespace Ics23

/-! ## The collision target -/

/-- A concrete collision in the hash family `H`: two distinct preimages under
the same `HashOp` with the same image. Every soundness theorem concludes by
exhibiting one of these, so the results are unconditional (no hash assumptions). -/
def HashCollision (H : HashFn) : Prop :=
  ∃ (op : HashOp) (a b : Bytes), a ≠ b ∧ H op a = H op b

/-- A fixed-length hash family (`cs`-byte digests), with `cs = child_size`. The
honest setting: SHA-256 outputs 32 bytes and the binary specs set
`child_size = 32`. -/
def FixedHash (H : HashFn) (cs : Nat) : Prop := ∀ (op : HashOp) (d : Bytes), (H op d).length = cs

/-! ## Well-formedness of a `ProofSpec`

These are the side conditions Theorems A and B assume. They are phrased as a
decidable `Bool` predicate so the shipped specs can be certified by computation. -/

/-- A length op whose output length is determined by the encoding itself
(length-prefixed) or fixed by construction (`Require*`). `noPrefix` is the only
non-determining op. -/
def lengthDetermining : LengthOp → Bool
  | .noPrefix => false
  | _ => true

/-- The fixed output length of a hash op, or `none` for `noHash` (identity,
hence variable length). -/
def fixedOutputLen : HashOp → Option Nat
  | .noHash => none
  | .sha256 => some 32
  | .sha512 => some 64
  | .keccak256 => some 32
  | .ripemd160 => some 20
  | .bitcoin => some 20
  | .sha512256 => some 32
  | .blake2b512 => some 64
  | .blake2s256 => some 32
  | .blake3 => some 32

/-- The leaf encoding `prefix ++ enc(prehashKey key) ++ enc(prehashValue value)`
parses unambiguously into its two fields: either the `LengthOp` delimits them, or
both prehash images have fixed length. This is the heart of leaf injectivity. -/
def leafDelimitingB (leaf : LeafOp) : Bool :=
  lengthDetermining leaf.length
  || ((fixedOutputLen leaf.prehashKey).isSome && (fixedOutputLen leaf.prehashValue).isSome)

/-- `l` is a permutation of `[0, …, l.length - 1]`. Decidable, and sufficient
because every index in range appearing in a list of that length forces a perm. -/
def isPermRange (l : List Nat) : Bool :=
  (List.range l.length).all (fun i => l.contains i)

/-- Inner-spec well-formedness: positive child size, sensible prefix window,
the `max < min + childSize` invariant (the same one `ensure_inner` enforces per
op), at least a binary node, a valid child ordering, and an `emptyChild` that is
either absent or exactly `childSize` bytes. -/
def innerWFB (isp : InnerSpec) : Bool :=
  (isp.childSize > 0)
  && (isp.minPrefixLength ≥ 0)
  && (isp.maxPrefixLength ≥ isp.minPrefixLength)
  && (isp.maxPrefixLength < isp.minPrefixLength + isp.childSize)
  && (isp.childOrder.length ≥ 2)
  && isPermRange isp.childOrder
  && (isp.emptyChild.isEmpty || ((isp.emptyChild.length : Int) = isp.childSize))

/-- The full decidable well-formedness predicate. The nonempty leaf prefix is
what makes leaf/inner domain separation (the `!has_prefix(leafPrefix, ...)`
check in `ensure_inner`) meaningful. -/
def wellFormedB (s : ProofSpec) : Bool :=
  leafDelimitingB s.leafSpec
  && (s.leafSpec.prefixBytes.length > 0)
  && innerWFB s.innerSpec

/-- A spec is well-formed when it passes the decidable check. -/
def WellFormed (s : ProofSpec) : Prop := wellFormedB s = true

instance (s : ProofSpec) : Decidable (WellFormed s) := by
  unfold WellFormed; infer_instance

/-! ## Theorem C: the shipped specs are well-formed

These hold by computation. They are the standalone, immediately-useful artifact:
a machine-checked answer to "is this `ProofSpec` safe to accept?". -/

theorem iavl_wellFormed : WellFormed iavlSpec := by decide

theorem tendermint_wellFormed : WellFormed tendermintSpec := by decide

theorem smt_wellFormed : WellFormed smtSpec := by decide

/-! ## Supporting lemmas (fully proved) -/

/-- An inner op's preimage is injective in its child: same op, same surrounding
prefix/suffix, so equal images force equal children. This is the cancellation
step used when walking two proofs up to a shared node. -/
theorem innerImage_inj (op : InnerOp) (c₁ c₂ : Bytes) :
    op.prefixBytes ++ c₁ ++ op.suffix = op.prefixBytes ++ c₂ ++ op.suffix ↔ c₁ = c₂ := by
  constructor
  · intro h
    have h2 : op.prefixBytes ++ c₁ = op.prefixBytes ++ c₂ := List.append_cancel_right h
    exact List.append_cancel_left h2
  · intro h; rw [h]

/-- `hasPrefix` is reflexive: every byte string is a prefix of itself. -/
theorem hasPrefix_refl (b : Bytes) : hasPrefix b b = true := by
  unfold hasPrefix
  simp

/-- If `[b]` is a prefix of `data`, then `data` starts with `b`. -/
theorem hasPrefix_single_head (b : UInt8) (data : Bytes)
    (h : hasPrefix [b] data = true) : data.head? = some b := by
  cases data with
  | nil => simp [hasPrefix] at h
  | cons x xs =>
    unfold hasPrefix at h
    simp only [List.length_cons, List.length_nil, List.take_succ_cons, List.take_zero,
      Bool.and_eq_true] at h
    have hx : b = x := by have := eq_of_beq h.2; simpa using this
    simp [hx]

/-- If `[b]` is *not* a prefix of a non-empty `data`, then `data` does not start
with `b`. -/
theorem not_hasPrefix_single_head (b x : UInt8) (xs : Bytes)
    (h : hasPrefix [b] (x :: xs) = false) : (x :: xs).head? ≠ some b := by
  unfold hasPrefix at h
  simp only [List.length_cons, List.length_nil, List.take_succ_cons, List.take_zero,
    Bool.and_eq_false_iff, decide_eq_false_iff_not, Nat.not_le] at h
  simp only [List.head?_cons, ne_eq, Option.some.injEq]
  rcases h with hc | hbeq
  · omega
  · intro hxb; rw [hxb] at hbeq; simp at hbeq

/-- Packaging a witnessed clash as a `HashCollision`. -/
theorem hashCollision_of (H : HashFn) (op : HashOp) (a b : Bytes)
    (hne : a ≠ b) (heq : H op a = H op b) : HashCollision H :=
  ⟨op, a, b, hne, heq⟩

/-- Leaf/inner domain separation as a collision: if a leaf preimage (starting
with the leaf prefix byte `b`) and an inner preimage (not starting with `b`) hash
to the same value under the same op, that is a collision. This discharges the
length-mismatch case of binding, where a short proof's leaf hash meets a long
proof's intermediate node hash. -/
theorem leaf_inner_domain_collision (H : HashFn) (op : HashOp)
    (lpre ipre : Bytes) (b : UInt8)
    (hl0 : lpre.head? = some b) (hi0 : ipre.head? ≠ some b)
    (heq : H op lpre = H op ipre) : HashCollision H := by
  refine hashCollision_of H op lpre ipre ?_ heq
  intro hpe; rw [hpe] at hl0; exact hi0 hl0

/-- The core inductive step of existence binding: one inner step is injective in
its child *up to a collision*. If two children hash to the same node under the
same inner op, then either the children are equal or the differing preimages
are an explicit collision. -/
theorem applyInner_inj (H : HashFn) (op : InnerOp) (c₁ c₂ r : Bytes)
    (h1 : applyInner H op c₁ = some r) (h2 : applyInner H op c₂ = some r) :
    c₁ = c₂ ∨ HashCollision H := by
  by_cases e1 : c₁.isEmpty = true
  · rw [applyInner, if_pos e1] at h1; simp at h1
  · by_cases e2 : c₂.isEmpty = true
    · rw [applyInner, if_pos e2] at h2; simp at h2
    · rw [applyInner, if_neg e1] at h1
      rw [applyInner, if_neg e2] at h2
      have hi1 := Option.some.inj h1
      have hi2 := Option.some.inj h2
      by_cases himg :
          op.prefixBytes ++ c₁ ++ op.suffix = op.prefixBytes ++ c₂ ++ op.suffix
      · exact Or.inl ((innerImage_inj op c₁ c₂).mp himg)
      · exact Or.inr (hashCollision_of H op.hash _ _ himg (hi1.trans hi2.symm))

/-- Folding the *same* op-list over two starting hashes to the same root forces
the starting hashes equal, up to a collision. This is the inductive backbone of
existence binding for the shared portion of two proof paths. -/
theorem applyPath_sameops_inj (H : HashFn) (isp : InnerSpec) :
    ∀ (path : List InnerOp) (h₁ h₂ r : Bytes),
      applyPath H isp h₁ path = some r →
      applyPath H isp h₂ path = some r →
      h₁ = h₂ ∨ HashCollision H := by
  intro path
  induction path with
  | nil =>
    intro h₁ h₂ r e1 e2
    simp only [applyPath, Option.some.injEq] at e1 e2
    exact Or.inl (e1.trans e2.symm)
  | cons step rest ih =>
    intro h₁ h₂ r e1 e2
    simp only [applyPath] at e1 e2
    cases hA1 : applyInner H step h₁ with
    | none => simp [hA1] at e1
    | some h₁' =>
      cases hA2 : applyInner H step h₂ with
      | none => simp [hA2] at e2
      | some h₂' =>
        simp only [hA1] at e1
        simp only [hA2] at e2
        by_cases g1 : (h₁'.length : Int) > isp.childSize ∧ isp.childSize ≥ 32
        · simp [g1] at e1
        · by_cases g2 : (h₂'.length : Int) > isp.childSize ∧ isp.childSize ≥ 32
          · simp [g2] at e2
          · rw [if_neg g1] at e1
            rw [if_neg g2] at e2
            rcases ih h₁' h₂' r e1 e2 with hh | hc
            · subst hh
              exact applyInner_inj H step h₁ h₂ h₁' hA1 hA2
            · exact Or.inr hc

/-- The positional ambiguity of finding F3, formalized: two distinct inner ops
that both pass `ensureInner` for `s` and decompose the *same* node preimage with
*different* children. This is the obstacle to a pure collision reduction for
general binding (see `docs/verification/properties.md`, `IavlPrefix.lean`). -/
def PositionalAmbiguity (s : ProofSpec) : Prop :=
  ∃ (op₁ op₂ : InnerOp) (c₁ c₂ : Bytes),
    ensureInner op₁ s = true ∧ ensureInner op₂ s = true ∧ op₁ ≠ op₂ ∧
    op₁.prefixBytes ++ c₁ ++ op₁.suffix = op₂.prefixBytes ++ c₂ ++ op₂.suffix

/-- The leaf-level analogue of `PositionalAmbiguity` (F3 at the leaf): two
spec-conformant leaf ops with *different* prefixes producing the same leaf hash.
Because `ensure_leaf` only requires the spec prefix to be *a* prefix (not exactly
equal), an adversary can pad the leaf prefix; resolving this needs the symbolic
hash model, exactly as for `PositionalAmbiguity`. -/
def LeafAmbiguity (H : HashFn) (s : ProofSpec) : Prop :=
  ∃ (leaf₁ leaf₂ : LeafOp) (k₁ v₁ k₂ v₂ L : Bytes),
    ensureLeaf leaf₁ s.leafSpec = true ∧ ensureLeaf leaf₂ s.leafSpec = true ∧
    leaf₁.prefixBytes ≠ leaf₂.prefixBytes ∧
    applyLeaf H leaf₁ k₁ v₁ = some L ∧ applyLeaf H leaf₂ k₂ v₂ = some L

/-- `h` is the image of some spec-conformant inner op — i.e. a non-leaf node
hash. Used to discharge the length-mismatch case of binding via leaf/inner
domain separation. -/
def IsInnerImage (H : HashFn) (s : ProofSpec) (h : Bytes) : Prop :=
  ∃ (op : InnerOp) (c : Bytes), ensureInner op s = true ∧ applyInner H op c = some h

/-- The result of folding a non-empty, spec-conformant path is an inner-node
image. (Component of the differing-length case of binding.) -/
theorem applyPath_result_isInnerImage (H : HashFn) (s : ProofSpec) :
    ∀ (p : List InnerOp) (h r : Bytes), p ≠ [] →
      (∀ op ∈ p, ensureInner op s = true) →
      applyPath H s.innerSpec h p = some r → IsInnerImage H s r := by
  intro p
  induction p with
  | nil => intro h r hne _ _; exact absurd rfl hne
  | cons op rest ih =>
    intro h r _ hall happ
    simp only [applyPath] at happ
    cases hA : applyInner H op h with
    | none => simp [hA] at happ
    | some h' =>
      simp only [hA] at happ
      by_cases g : (h'.length : Int) > s.innerSpec.childSize ∧ s.innerSpec.childSize ≥ 32
      · simp [g] at happ
      · rw [if_neg g] at happ
        cases rest with
        | nil =>
          simp only [applyPath, Option.some.injEq] at happ
          refine ⟨op, h, hall op (List.mem_cons_self ..), ?_⟩
          rw [hA]; exact congrArg some happ
        | cons op2 rest2 =>
          exact ih h' r (by simp) (fun o ho => hall o (List.mem_cons_of_mem op ho)) happ

/-- The image equation extracted from a successful `applyInner`. -/
theorem applyInner_image (H : HashFn) (op : InnerOp) (c r : Bytes)
    (happ : applyInner H op c = some r) :
    H op.hash (op.prefixBytes ++ c ++ op.suffix) = r := by
  unfold applyInner at happ
  by_cases e : c.isEmpty = true
  · rw [if_pos e] at happ; exact absurd happ (by simp)
  · rw [if_neg e] at happ; exact Option.some.inj happ

/-- A spec-conformant inner op uses the inner spec's hash. -/
theorem ensureInner_hash (op : InnerOp) (s : ProofSpec)
    (h : ensureInner op s = true) : op.hash = s.innerSpec.hash := by
  unfold ensureInner at h
  simp only [Bool.and_eq_true] at h
  obtain ⟨⟨⟨⟨⟨⟨hh, _⟩, _⟩, _⟩, _⟩, _⟩, _⟩ := h
  exact eq_of_beq hh

/-- **Binding for equal-length paths.** Two spec-conformant paths of the *same
length* folding two inputs to the same root force the inputs equal — or exhibit a
collision or the F3 positional ambiguity. This strengthens the same-shape result
(which required identical ops) to arbitrary differing ops at equal depth; the
equal length is what sidesteps the leaf/inner length-mismatch case. -/
theorem applyPath_eqlen_merge (H : HashFn) (s : ProofSpec) :
    ∀ (p₁ p₂ : List InnerOp), p₁.length = p₂.length →
      ∀ (h₁ h₂ r : Bytes),
      (∀ op ∈ p₁, ensureInner op s = true) → (∀ op ∈ p₂, ensureInner op s = true) →
      applyPath H s.innerSpec h₁ p₁ = some r → applyPath H s.innerSpec h₂ p₂ = some r →
      HashCollision H ∨ PositionalAmbiguity s ∨ h₁ = h₂ := by
  intro p₁
  induction p₁ with
  | nil =>
    intro p₂ hlen h₁ h₂ r _ _ e1 e2
    cases p₂ with
    | nil =>
      simp only [applyPath, Option.some.injEq] at e1 e2
      exact Or.inr (Or.inr (e1.trans e2.symm))
    | cons => simp at hlen
  | cons op1 rest1 ih =>
    intro p₂ hlen h₁ h₂ r hall1 hall2 e1 e2
    cases p₂ with
    | nil => simp at hlen
    | cons op2 rest2 =>
      have hlen' : rest1.length = rest2.length := by simpa using hlen
      simp only [applyPath] at e1 e2
      cases hA1 : applyInner H op1 h₁ with
      | none => simp [hA1] at e1
      | some h₁' =>
        cases hA2 : applyInner H op2 h₂ with
        | none => simp [hA2] at e2
        | some h₂' =>
          simp only [hA1] at e1
          simp only [hA2] at e2
          by_cases g1 : (h₁'.length : Int) > s.innerSpec.childSize ∧ s.innerSpec.childSize ≥ 32
          · simp [g1] at e1
          · by_cases g2 : (h₂'.length : Int) > s.innerSpec.childSize ∧ s.innerSpec.childSize ≥ 32
            · simp [g2] at e2
            · rw [if_neg g1] at e1
              rw [if_neg g2] at e2
              have hin1 : ensureInner op1 s = true := hall1 op1 (List.mem_cons_self ..)
              have hin2 : ensureInner op2 s = true := hall2 op2 (List.mem_cons_self ..)
              rcases ih rest2 hlen' h₁' h₂' r
                  (fun o ho => hall1 o (List.mem_cons_of_mem op1 ho))
                  (fun o ho => hall2 o (List.mem_cons_of_mem op2 ho)) e1 e2 with hc | ha | heq
              · exact Or.inl hc
              · exact Or.inr (Or.inl ha)
              · subst heq
                have hP1 := applyInner_image H op1 h₁ h₁' hA1
                have hP2 := applyInner_image H op2 h₂ h₁' hA2
                have hhash : op1.hash = op2.hash :=
                  (ensureInner_hash op1 s hin1).trans (ensureInner_hash op2 s hin2).symm
                rw [hhash] at hP1
                -- hP1 : H op2.hash P1 = h₁',  hP2 : H op2.hash P2 = h₁'
                by_cases hPeq :
                    op1.prefixBytes ++ h₁ ++ op1.suffix = op2.prefixBytes ++ h₂ ++ op2.suffix
                · by_cases hop : op1 = op2
                  · subst hop
                    exact Or.inr (Or.inr ((innerImage_inj op1 h₁ h₂).mp hPeq))
                  · exact Or.inr (Or.inl ⟨op1, op2, h₁, h₂, hin1, hin2, hop, hPeq⟩)
                · exact Or.inl (hashCollision_of H op2.hash _ _ hPeq (hP1.trans hP2.symm))

/-- Root-side characterization of a path fold: peeling the last (root-most) op.
Enables aligning two differing-length folds from the root. -/
theorem applyPath_snoc (H : HashFn) (isp : InnerSpec) :
    ∀ (p : List InnerOp) (op : InnerOp) (h r : Bytes),
      applyPath H isp h (p ++ [op]) = some r ↔
      ∃ m, applyPath H isp h p = some m ∧ applyInner H op m = some r ∧
           ¬((r.length : Int) > isp.childSize ∧ isp.childSize ≥ 32) := by
  intro p
  induction p with
  | nil =>
    intro op h r
    simp only [List.nil_append]
    constructor
    · intro hh
      simp only [applyPath] at hh
      cases hA : applyInner H op h with
      | none => simp only [hA] at hh; simp at hh
      | some h' =>
        simp only [hA] at hh
        by_cases g : (h'.length : Int) > isp.childSize ∧ isp.childSize ≥ 32
        · rw [if_pos g] at hh; simp at hh
        · rw [if_neg g] at hh
          simp only [Option.some.injEq] at hh
          subst hh
          exact ⟨h, rfl, hA, g⟩
    · rintro ⟨m, hm, hop, hg⟩
      simp only [applyPath, Option.some.injEq] at hm
      subst hm
      simp only [applyPath, hop]
      rw [if_neg hg]
  | cons a rest ih =>
    intro op h r
    simp only [List.cons_append]
    constructor
    · intro hh
      simp only [applyPath] at hh
      cases hA : applyInner H a h with
      | none => simp only [hA] at hh; simp at hh
      | some h' =>
        simp only [hA] at hh
        by_cases g : (h'.length : Int) > isp.childSize ∧ isp.childSize ≥ 32
        · rw [if_pos g] at hh; simp at hh
        · rw [if_neg g] at hh
          obtain ⟨m, hm, hop, hgr⟩ := (ih op h' r).mp hh
          refine ⟨m, ?_, hop, hgr⟩
          simp only [applyPath, hA]; rw [if_neg g]; exact hm
    · rintro ⟨m, hm, hop, hgr⟩
      simp only [applyPath] at hm
      cases hA : applyInner H a h with
      | none => simp only [hA] at hm; simp at hm
      | some h' =>
        simp only [hA] at hm
        by_cases g : (h'.length : Int) > isp.childSize ∧ isp.childSize ≥ 32
        · rw [if_pos g] at hm; simp at hm
        · rw [if_neg g] at hm
          simp only [applyPath, hA]; rw [if_neg g]
          exact (ih op h' r).mpr ⟨m, hm, hop, hgr⟩

/-- **Root-side merge.** Two spec-conformant paths folding `h₁`, `h₂` to the same
root yield: a collision, the F3 ambiguity, equal inputs, or one input being an
inner-node image (the length-mismatch case, discharged at the leaf level by
domain separation). Root-side recursion keeps `h₁, h₂` fixed, so the conclusion is
about the original inputs. This is the structural core of general binding. -/
theorem applyPath_merge (H : HashFn) (s : ProofSpec) :
    ∀ (n : Nat) (p₁ p₂ : List InnerOp) (h₁ h₂ r : Bytes),
      p₁.length + p₂.length = n →
      (∀ op ∈ p₁, ensureInner op s = true) → (∀ op ∈ p₂, ensureInner op s = true) →
      applyPath H s.innerSpec h₁ p₁ = some r → applyPath H s.innerSpec h₂ p₂ = some r →
      HashCollision H ∨ PositionalAmbiguity s ∨ h₁ = h₂
        ∨ IsInnerImage H s h₁ ∨ IsInnerImage H s h₂ := by
  intro n
  induction n using Nat.strongRecOn with
  | _ n ih =>
    intro p₁ p₂ h₁ h₂ r hn hin1 hin2 e1 e2
    rcases List.eq_nil_or_concat p₁ with hp1 | ⟨q₁, op1, hp1⟩
    · subst hp1
      simp only [applyPath, Option.some.injEq] at e1
      cases hp2c : p₂ with
      | nil =>
        rw [hp2c] at e2
        simp only [applyPath, Option.some.injEq] at e2
        exact Or.inr (Or.inr (Or.inl (e1.trans e2.symm)))
      | cons a2 rest2 =>
        have himg : IsInnerImage H s r :=
          applyPath_result_isInnerImage H s p₂ h₂ r (by rw [hp2c]; simp)
            hin2 e2
        rw [← e1] at himg
        exact Or.inr (Or.inr (Or.inr (Or.inl himg)))
    · rw [List.concat_eq_append] at hp1
      subst hp1
      rcases List.eq_nil_or_concat p₂ with hp2 | ⟨q₂, op2, hp2⟩
      · subst hp2
        simp only [applyPath, Option.some.injEq] at e2
        have himg : IsInnerImage H s r :=
          applyPath_result_isInnerImage H s (q₁ ++ [op1]) h₁ r (by simp) hin1 e1
        rw [← e2] at himg
        exact Or.inr (Or.inr (Or.inr (Or.inr himg)))
      · rw [List.concat_eq_append] at hp2
        subst hp2
        obtain ⟨m₁, hm1, hop1, _⟩ := (applyPath_snoc H s.innerSpec q₁ op1 h₁ r).mp e1
        obtain ⟨m₂, hm2, hop2, _⟩ := (applyPath_snoc H s.innerSpec q₂ op2 h₂ r).mp e2
        have hP1 := applyInner_image H op1 m₁ r hop1
        have hP2 := applyInner_image H op2 m₂ r hop2
        have hin1' : ensureInner op1 s = true := hin1 op1 (by simp)
        have hin2' : ensureInner op2 s = true := hin2 op2 (by simp)
        have hhash : op1.hash = op2.hash :=
          (ensureInner_hash op1 s hin1').trans (ensureInner_hash op2 s hin2').symm
        rw [hhash] at hP1
        by_cases hPeq :
            op1.prefixBytes ++ m₁ ++ op1.suffix = op2.prefixBytes ++ m₂ ++ op2.suffix
        · by_cases hop : op1 = op2
          · subst hop
            have hmeq : m₁ = m₂ := (innerImage_inj op1 m₁ m₂).mp hPeq
            subst hmeq
            have hlt : q₁.length + q₂.length < n := by
              simp only [List.length_append, List.length_cons, List.length_nil] at hn; omega
            exact ih (q₁.length + q₂.length) hlt q₁ q₂ h₁ h₂ m₁ rfl
              (fun o ho => hin1 o (List.mem_append.mpr (Or.inl ho)))
              (fun o ho => hin2 o (List.mem_append.mpr (Or.inl ho))) hm1 hm2
          · exact Or.inr (Or.inl ⟨op1, op2, m₁, m₂, hin1', hin2', hop, hPeq⟩)
        · exact Or.inl (hashCollision_of H op2.hash _ _ hPeq (hP1.trans hP2.symm))

/-! ## Theorem A: existence binding (soundness)

The general existence-binding theorem is `Ics23.existence_binding_shaped` (in
`Existence.lean`), proved with **no `sorry`** for the production-spec shape
(single-byte leaf prefix, shared leaf/inner hash op, `min_prefix_length ≥ 1`),
with corollaries `existence_binding_{iavl,tendermint,smt}`. Its conclusion is the
honest three-way disjunction

  `HashCollision H ∨ PositionalAmbiguity s ∨ LeafAmbiguity H s`,

which is exactly what the abstract-hash model can establish: a forger who binds a
key to two values either found a hash collision, or exploited the inner
positional ambiguity F3, or the leaf-level analogue. A collision-only conclusion
is provably too strong (machine-checked in `IavlPrefix.lean`); collapsing the two
ambiguity arms requires the symbolic "Merkle" hash model. The same-leaf case
(`existence_binding_sameleaf`) — the honest case, since a key's leaf op is
determined in a real tree — yields just `HashCollision ∨ PositionalAmbiguity`. -/

end Ics23
