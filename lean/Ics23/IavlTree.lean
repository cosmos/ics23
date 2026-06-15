/-
Honest-root IAVL tree model: Theorem A for the IAVL spec.

The Tendermint-shape development (Tree.lean) assumes `min_prefix_length =
max_prefix_length`, which pins an accepted inner op's prefix to the node's
prefix length (`split_bounds` → `split_pins`). IAVL violates it: the inner
prefix is 4–12 bytes of height/size/version varints (followed by the `0x20`
amino length byte for the left child), and the 33-byte child slots are the
`0x20` length byte plus the 32-byte digest. So an IAVL node preimage is

  `pre ++ lh ++ mid ++ rh`,  `|pre| ∈ [4,12]`, `|lh| = |rh| = 32`, `mid = [0x20]`

— exactly the `MTree` node shape with a 1-byte `mid` and empty `suf`.

The pinning argument changes: instead of prefix-length pinning, the verifier's
`suffix.length % child_size = 0` check together with the well-formedness fact
`max_prefix_length < child_size` (12 < 33) pins the suffix length to `0` or
`child_size`, and each choice resolves the whole split (`split_pins_var`). The
`max < min + child_size` family of checks is doing its job: no third,
straddling reading exists even with variable prefixes.
-/
import Ics23.Tree
import Ics23.TreeNonExist

namespace Ics23

/-- IAVL-shaped conformance of a tree to a spec: nodes use the inner-spec hash
with a 1-byte `mid` (the amino `0x20` length byte for the right child), empty
`suf`, a prefix whose length lies in the spec's `[min, max]` window (the
height/size/version varints plus the left child's length byte), nonempty and
not leaf-prefixed (domain separation); leaves carry exactly the spec leaf op. -/
def WFTreeI (s : ProofSpec) (b : UInt8) : MTree → Prop
  | .leaf op _ _ => op = s.leafSpec
  | .node ih pre mid suf l r =>
      ih = s.innerSpec.hash ∧ mid.length = 1 ∧ suf = [] ∧
      s.innerSpec.minPrefixLength ≤ (pre.length : Int) ∧
      (pre.length : Int) ≤ s.innerSpec.maxPrefixLength ∧
      pre ≠ [] ∧ pre.head? ≠ some b ∧
      WFTreeI s b l ∧ WFTreeI s b r

/-- `ensure_inner`'s suffix-multiple fact, extracted without the
`min = max` assumption `split_bounds` carries. -/
theorem ensureInner_suffix_mod (op : InnerOp) (s : ProofSpec) (cs : Nat)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hen : ensureInner op s = true) : op.suffix.length % cs = 0 := by
  unfold ensureInner at hen
  simp only [Bool.and_eq_true, decide_eq_true_eq] at hen
  have hc7 := hen.2
  have hcsInt : s.innerSpec.childSize = (cs : Int) := by omega
  rw [hcsInt] at hc7
  have : ((op.suffix.length % cs : Nat) : Int) = 0 := by
    rw [Int.natCast_emod]; exact hc7
  exact_mod_cast this

/-- A four-part node split with a full `cs`-length suffix resolves completely:
prefix = node prefix, child = left digest, suffix = `mid ++ rh`. -/
theorem split_var_left (pre lh mid rh topPre m topSuf : Bytes) (cs d : Nat)
    (hN : topPre ++ m ++ topSuf = pre ++ lh ++ mid ++ rh)
    (hlh : lh.length = d) (hrh : rh.length = d) (hm : m.length = d)
    (hmid : mid.length + d = cs)
    (hsufcs : topSuf.length = cs) :
    topPre = pre ∧ m = lh ∧ topSuf = mid ++ rh := by
  have hlen : topPre.length = pre.length := by
    have h := congrArg List.length hN
    simp only [List.length_append] at h; omega
  simp only [List.append_assoc] at hN
  have hA := List.append_inj hN hlen
  have hB := List.append_inj hA.2 (by rw [hm, hlh])
  exact ⟨hA.1, hB.1, hB.2⟩

/-- A four-part node split with an empty suffix resolves completely:
prefix = node prefix ++ left digest ++ `mid`, child = right digest. -/
theorem split_var_right (pre lh mid rh topPre m : Bytes) (d : Nat)
    (hN : topPre ++ m ++ [] = pre ++ lh ++ mid ++ rh)
    (hrh : rh.length = d) (hm : m.length = d) :
    topPre = pre ++ lh ++ mid ∧ m = rh := by
  rw [List.append_nil] at hN
  have hlen : topPre.length = ((pre ++ lh) ++ mid).length := by
    have h := congrArg List.length hN
    simp only [List.length_append] at h ⊢; omega
  exact List.append_inj hN hlen

/-- **F3 resolution for variable prefixes (the IAVL crux).** For a node preimage
`pre ++ lh ++ mid ++ rh` with `d`-byte digests, `cs = |mid| + d` byte child
slots, and `|pre| < cs`, the verifier's `suffix % cs = 0` check alone pins any
accepted `d`-length-child split to exactly the two genuine children — no prefix
pinning needed, so `min ≠ max` prefix windows are fine. -/
theorem split_pins_var (pre lh mid rh topPre m topSuf : Bytes) (cs d : Nat)
    (hN : topPre ++ m ++ topSuf = pre ++ lh ++ mid ++ rh)
    (hlh : lh.length = d) (hrh : rh.length = d) (hm : m.length = d)
    (hmid : mid.length + d = cs)
    (hprelt : pre.length < cs)
    (hsuf : topSuf.length % cs = 0) :
    (topPre = pre ∧ m = lh ∧ topSuf = mid ++ rh)
    ∨ (topPre = pre ++ lh ++ mid ∧ m = rh ∧ topSuf = []) := by
  have hlentot : topPre.length + m.length + topSuf.length
      = pre.length + lh.length + mid.length + rh.length := by
    have h := congrArg List.length hN
    simp only [List.length_append] at h
    omega
  have hdvd : cs ∣ topSuf.length := Nat.dvd_of_mod_eq_zero hsuf
  have hsuflen : topSuf.length = 0 ∨ topSuf.length = cs := by
    rcases Nat.eq_zero_or_pos topSuf.length with h0 | hpos
    · exact Or.inl h0
    · have hge : cs ≤ topSuf.length := Nat.le_of_dvd hpos hdvd
      rcases Nat.eq_or_lt_of_le hge with heq | hlt
      · exact Or.inr heq.symm
      · -- a ≥ 2·cs suffix would force a negative prefix: |topPre| + |topSuf|
        -- = |pre| + cs with |pre| < cs
        exfalso
        obtain ⟨k, hk⟩ := hdvd
        have hk2 : 2 ≤ k := by
          rcases Nat.lt_or_ge k 2 with hlt2 | hge2
          · exfalso
            have hk01 : k = 0 ∨ k = 1 := by omega
            rcases hk01 with rfl | rfl
            · rw [Nat.mul_zero] at hk; omega
            · rw [Nat.mul_one] at hk; omega
          · exact hge2
        have h2cs : 2 * cs ≤ topSuf.length := by
          rw [hk, Nat.mul_comm cs k]
          exact Nat.mul_le_mul_right cs hk2
        omega
  rcases hsuflen with hs0 | hscs
  · exact Or.inr (by
      have hsnil : topSuf = [] := List.length_eq_zero_iff.mp hs0
      rw [hsnil] at hN
      obtain ⟨h1, h2⟩ := split_var_right pre lh mid rh topPre m d hN hrh hm
      exact ⟨h1, h2, hsnil⟩)
  · exact Or.inl (split_var_left pre lh mid rh topPre m topSuf cs d hN hlh hrh hm hmid hscs)

/-- Core induction (IAVL analog of `reaches`): a proof whose leaf hash folds to
a real IAVL-shaped tree's root reaches a genuine leaf, with each step forced
into a real child by `split_pins_var`. -/
theorem reachesI (H : HashFn) (s : ProofSpec) (b : UInt8) (cs d : Nat)
    (hH : FixedHash H d) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hcsd : cs = d + 1)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmaxcs : s.innerSpec.maxPrefixLength < s.innerSpec.childSize)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (key value lh : Bytes) (path : List InnerOp) (root : Bytes),
      WFTreeI s b t →
      applyLeaf H s.leafSpec key value = some lh →
      (∀ op ∈ path, ensureInner op s = true) →
      rootHash H t = some root →
      applyPath H s.innerSpec lh path = some root →
      TreeMember key value t ∨ HashCollision H := by
  intro t
  induction t with
  | leaf top tk tv =>
    intro key value lh path root hwf hlh hpath hrh hap
    simp only [WFTreeI] at hwf; subst hwf
    rw [rootHash] at hrh
    cases path with
    | nil =>
      simp only [applyPath, Option.some.injEq] at hap
      subst hap
      rcases hLInj key value tk tv _ hlh hrh with ⟨hk, hv⟩ | hc
      · exact Or.inl ⟨hk.symm, hv.symm⟩
      · exact Or.inr hc
    | cons op rest =>
      have hii : IsInnerImage H s root :=
        applyPath_result_isInnerImage H s (op :: rest) lh root (by simp) hpath hap
      exact Or.inr (leafHash_innerImage_collision H s s.leafSpec tk tv root b
        hihash.symm hlpre hmin (ensureLeaf_self s.leafSpec) hrh hii)
  | node ih pre mid suf l r ihl ihr =>
    intro key value lh path root hwf hlh hpath hrh hap
    obtain ⟨hih, hmid1, hsuf, hpremin, hpremax, hprene, hprehead, hwfl, hwfr⟩ := hwf
    subst hsuf; subst hih
    rw [rootHash] at hrh
    cases hl : rootHash H l with
    | none => rw [hl] at hrh; simp at hrh
    | some lhL =>
    cases hr : rootHash H r with
    | none => rw [hl, hr] at hrh; simp at hrh
    | some rhR =>
    rw [hl, hr] at hrh
    simp only [List.append_nil, Option.some.injEq] at hrh
    have hcsInt : s.innerSpec.childSize = (cs : Int) := by omega
    have hprelt : pre.length < cs := by omega
    rcases List.eq_nil_or_concat path with hpnil | ⟨q, topOp, hpc⟩
    · -- empty path: the proof's leaf hash equals the node hash ⇒ domain collision
      subst hpnil
      simp only [applyPath, Option.some.injEq] at hap
      obtain ⟨tail, hlheq⟩ := applyLeaf_head H s.leafSpec key value lh b hlpre hlh
      rw [← hihash] at hlheq
      refine Or.inr (leaf_inner_domain_collision H s.innerSpec.hash (b :: tail)
        (pre ++ lhL ++ mid ++ rhR) b (by simp) ?_ ?_)
      · cases hpc2 : pre with
        | nil => exact absurd hpc2 hprene
        | cons x xs =>
          rw [hpc2] at hprehead
          simp only [List.cons_append, List.head?_cons] at hprehead ⊢
          exact hprehead
      · rw [← hlheq, hap]; exact hrh.symm
    · -- nonempty path: peel the root op and force it into a genuine child
      rw [List.concat_eq_append] at hpc; subst hpc
      obtain ⟨m, hpm, htop, _⟩ := (applyPath_snoc H s.innerSpec q topOp lh root).mp hap
      have htopimg := applyInner_image H topOp m root htop
      rw [ensureInner_hash topOp s (hpath topOp (by simp))] at htopimg
      by_cases hpe : topOp.prefixBytes ++ m ++ topOp.suffix = pre ++ lhL ++ mid ++ rhR
      · have hmlen : m.length = d :=
          applyPath_len H s.innerSpec d hH q lh m
            (applyLeaf_len H d hH s.leafSpec key value lh hlh) hpm
        have hsufmod := ensureInner_suffix_mod topOp s cs hcsspec
          (hpath topOp (by simp))
        rcases split_pins_var pre lhL mid rhR topOp.prefixBytes m topOp.suffix cs d
          hpe (rootHash_len H d hH l lhL hl) (rootHash_len H d hH r rhR hr)
          hmlen (by omega) hprelt hsufmod with ⟨_, hml, _⟩ | ⟨_, hmr, _⟩
        · rw [hml] at hpm
          rcases ihl key value lh q lhL hwfl hlh
            (fun o ho => hpath o (List.mem_append.mpr (Or.inl ho))) hl hpm with hmem | hcol
          · exact Or.inl (Or.inl hmem)
          · exact Or.inr hcol
        · rw [hmr] at hpm
          rcases ihr key value lh q rhR hwfr hlh
            (fun o ho => hpath o (List.mem_append.mpr (Or.inl ho))) hr hpm with hmem | hcol
          · exact Or.inl (Or.inr hmem)
          · exact Or.inr hcol
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpe (htopimg.trans hrh.symm))

/-- **Honest-root Theorem A, IAVL tree model.** An accepted existence proof
against an honest IAVL root names a genuine `(key, value)` leaf — up to a hash
collision. -/
theorem membership_soundI (H : HashFn) (s : ProofSpec) (b : UInt8) (cs d : Nat)
    (hH : FixedHash H d) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hcsd : cs = d + 1)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmaxcs : s.innerSpec.maxPrefixLength < s.innerSpec.childSize)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (ep : ExistenceProof) (root key value : Bytes),
      WFTreeI s b t →
      ep.leaf = s.leafSpec →
      rootHash H t = some root →
      verifyExistence H ep s root key value = true →
      TreeMember key value t ∨ HashCollision H := by
  intro t ep root key value hwf hep hrh hver
  have hkey : ep.key = key ∧ ep.value = value := by
    unfold verifyExistence at hver
    simp only [Bool.and_eq_true] at hver
    obtain ⟨⟨⟨_, hkk⟩, hvv⟩, _⟩ := hver
    exact ⟨by simpa using hkk, by simpa using hvv⟩
  rw [← hkey.1, ← hkey.2] at hver
  obtain ⟨lh, hlf, hap, hinn⟩ := verifyExistence_navigates H s ep root hep hver
  rw [hkey.1, hkey.2] at hlf
  exact reachesI H s b cs d hH hcs hcsspec hcsd hlhash.symm hlpre hmin hmaxcs hLInj
    t key value lh ep.path root hwf hlf hinn hrh hap

/-- **Theorem A for the IAVL spec.** Structural side conditions discharged by
computation; joint leaf injectivity proved (`leafInj_iavl`); the single
remaining hypothesis is the genuine cryptographic assumption (`FixedHash`). -/
theorem membership_sound_iavl (H : HashFn)
    (hH : FixedHash H 32) :
    ∀ (t : MTree) (ep : ExistenceProof) (root key value : Bytes),
      WFTreeI iavlSpec 0 t →
      ep.leaf = iavlSpec.leafSpec →
      rootHash H t = some root →
      verifyExistence H ep iavlSpec root key value = true →
      TreeMember key value t ∨ HashCollision H :=
  membership_soundI H iavlSpec 0 33 32 hH (by decide) (by decide) (by decide)
    (by decide) (by decide) (by decide) (by decide) (leafInj_iavl H)

end Ics23
