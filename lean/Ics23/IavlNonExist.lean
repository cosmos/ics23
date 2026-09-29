/-
Non-existence soundness (Theorem B) for the IAVL tree model.

Everything order-theoretic reuses TreeNonExist.lean verbatim: `maxKey`,
`minKey`, `SortedTree`, `IsSubtree`, the BST gap (`node_gap_no_member`), the
neighbor-walk decomposition (`ensureLeftNeighbor_spec`), and — because IAVL's
`emptyChild` is `[]` (length 0 ≠ 33 = `childSize`) — the Tendermint-style
padding bridges `ensureRightMost_suffix_nil` / `ensureLeftMost_suffix_cs`,
instantiated at `cs = 33`.

What changes is the navigation: with IAVL's variable 4–12 byte prefixes and
33-byte child slots (1-byte amino length prefix + 32-byte digest), the
`split_left` / `split_right` / `split_pins` family is replaced by
`split_var_left` / `split_var_right` / `split_pins_var` (IavlTree.lean), which
pin a node split from the suffix length alone. Key order is raw `bytesLt`
(`prehash_key_before_comparison = false` for IAVL).
-/
import Ics23.IavlTree

namespace Ics23

/-- An all-right path (`ensure_right_most`: each step has an empty suffix)
reaches the rightmost leaf of an IAVL-shaped tree — the proof's key is
`maxKey t`, up to a collision. -/
theorem reaches_maxI (H : HashFn) (s : ProofSpec) (b : UInt8) (d : Nat)
    (hH : FixedHash H d)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (key value lh : Bytes) (path : List InnerOp) (root : Bytes),
      WFTreeI s b t →
      applyLeaf H s.leafSpec key value = some lh →
      (∀ op ∈ path, ensureInner op s = true ∧ op.suffix = []) →
      rootHash H t = some root →
      applyPath H s.innerSpec lh path = some root →
      (key = maxKey t ∧ TreeMember key value t) ∨ HashCollision H := by
  intro t
  induction t with
  | leaf top tk tv =>
    intro key value lh path root hwf hlh hpath hrh hap
    simp only [WFTreeI] at hwf; subst hwf
    rw [rootHash] at hrh
    cases path with
    | nil =>
      simp only [applyPath, Option.some.injEq] at hap; subst hap
      rcases hLInj key value tk tv _ hlh hrh with ⟨hk, hv⟩ | hc
      · exact Or.inl ⟨hk, hk.symm, hv.symm⟩
      · exact Or.inr hc
    | cons op rest =>
      have hii : IsInnerImage H s root :=
        applyPath_result_isInnerImage H s (op :: rest) lh root (by simp)
          (fun o ho => (hpath o ho).1) hap
      exact Or.inr (leafHash_innerImage_collision H s s.leafSpec tk tv root b
        hihash.symm hlpre hmin (ensureLeaf_self s.leafSpec) hrh hii)
  | node ih pre mid suf l r _ ihr =>
    intro key value lh path root hwf hlh hpath hrh hap
    obtain ⟨hih, _, hsuf, _, _, hprene, hprehead, _, hwfr⟩ := hwf
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
    rcases List.eq_nil_or_concat path with hpnil | ⟨q, topOp, hpc⟩
    · subst hpnil
      simp only [applyPath, Option.some.injEq] at hap
      obtain ⟨tail, hlheq⟩ := applyLeaf_head H s.leafSpec key value lh b hlpre hlh
      rw [← hihash] at hlheq
      refine Or.inr (leaf_inner_domain_collision H s.innerSpec.hash (b :: tail)
        (pre ++ lhL ++ mid ++ rhR) b (by simp) ?_ ?_)
      · cases hpc2 : pre with
        | nil => exact absurd hpc2 hprene
        | cons x xs =>
          rw [hpc2] at hprehead; simp only [List.cons_append, List.head?_cons] at hprehead ⊢
          exact hprehead
      · rw [← hlheq, hap]; exact hrh.symm
    · rw [List.concat_eq_append] at hpc; subst hpc
      obtain ⟨m, hpm, htop, _⟩ := (applyPath_snoc H s.innerSpec q topOp lh root).mp hap
      have htopimg := applyInner_image H topOp m root htop
      rw [ensureInner_hash topOp s (hpath topOp (by simp)).1] at htopimg
      by_cases hpe : topOp.prefixBytes ++ m ++ topOp.suffix = pre ++ lhL ++ mid ++ rhR
      · have hmlen : m.length = d :=
          applyPath_len H s.innerSpec d hH q lh m
            (applyLeaf_len H d hH s.leafSpec key value lh hlh) hpm
        rw [(hpath topOp (by simp)).2] at hpe
        obtain ⟨_, hmr⟩ := split_var_right pre lhL mid rhR topOp.prefixBytes m d
          hpe (rootHash_len H d hH r rhR hr) hmlen
        rw [hmr] at hpm
        rcases ihr key value lh q rhR hwfr hlh
          (fun o ho => hpath o (List.mem_append.mpr (Or.inl ho))) hr hpm with ⟨hk, hmem⟩ | hcol
        · exact Or.inl ⟨by rw [hk]; rfl, Or.inr hmem⟩
        · exact Or.inr hcol
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpe (htopimg.trans hrh.symm))

/-- An all-left path (`ensure_left_most`: each step has a full `cs`-suffix)
reaches the leftmost leaf of an IAVL-shaped tree — the proof's key is
`minKey t`, up to a collision. -/
theorem reaches_minI (H : HashFn) (s : ProofSpec) (b : UInt8) (cs d : Nat)
    (hH : FixedHash H d)
    (hcsd : cs = d + 1)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (key value lh : Bytes) (path : List InnerOp) (root : Bytes),
      WFTreeI s b t →
      applyLeaf H s.leafSpec key value = some lh →
      (∀ op ∈ path, ensureInner op s = true ∧ op.suffix.length = cs) →
      rootHash H t = some root →
      applyPath H s.innerSpec lh path = some root →
      (key = minKey t ∧ TreeMember key value t) ∨ HashCollision H := by
  intro t
  induction t with
  | leaf top tk tv =>
    intro key value lh path root hwf hlh hpath hrh hap
    simp only [WFTreeI] at hwf; subst hwf
    rw [rootHash] at hrh
    cases path with
    | nil =>
      simp only [applyPath, Option.some.injEq] at hap; subst hap
      rcases hLInj key value tk tv _ hlh hrh with ⟨hk, hv⟩ | hc
      · exact Or.inl ⟨hk, hk.symm, hv.symm⟩
      · exact Or.inr hc
    | cons op rest =>
      have hii : IsInnerImage H s root :=
        applyPath_result_isInnerImage H s (op :: rest) lh root (by simp)
          (fun o ho => (hpath o ho).1) hap
      exact Or.inr (leafHash_innerImage_collision H s s.leafSpec tk tv root b
        hihash.symm hlpre hmin (ensureLeaf_self s.leafSpec) hrh hii)
  | node ih pre mid suf l r ihl _ =>
    intro key value lh path root hwf hlh hpath hrh hap
    obtain ⟨hih, hmid1, hsuf, _, _, hprene, hprehead, hwfl, _⟩ := hwf
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
    rcases List.eq_nil_or_concat path with hpnil | ⟨q, topOp, hpc⟩
    · subst hpnil
      simp only [applyPath, Option.some.injEq] at hap
      obtain ⟨tail, hlheq⟩ := applyLeaf_head H s.leafSpec key value lh b hlpre hlh
      rw [← hihash] at hlheq
      refine Or.inr (leaf_inner_domain_collision H s.innerSpec.hash (b :: tail)
        (pre ++ lhL ++ mid ++ rhR) b (by simp) ?_ ?_)
      · cases hpc2 : pre with
        | nil => exact absurd hpc2 hprene
        | cons x xs =>
          rw [hpc2] at hprehead; simp only [List.cons_append, List.head?_cons] at hprehead ⊢
          exact hprehead
      · rw [← hlheq, hap]; exact hrh.symm
    · rw [List.concat_eq_append] at hpc; subst hpc
      obtain ⟨m, hpm, htop, _⟩ := (applyPath_snoc H s.innerSpec q topOp lh root).mp hap
      have htopimg := applyInner_image H topOp m root htop
      rw [ensureInner_hash topOp s (hpath topOp (by simp)).1] at htopimg
      by_cases hpe : topOp.prefixBytes ++ m ++ topOp.suffix = pre ++ lhL ++ mid ++ rhR
      · have hmlen : m.length = d :=
          applyPath_len H s.innerSpec d hH q lh m
            (applyLeaf_len H d hH s.leafSpec key value lh hlh) hpm
        obtain ⟨_, hml, _⟩ := split_var_left pre lhL mid rhR topOp.prefixBytes m topOp.suffix cs d
          hpe (rootHash_len H d hH l lhL hl) (rootHash_len H d hH r rhR hr) hmlen
          (by omega) (hpath topOp (by simp)).2
        rw [hml] at hpm
        rcases ihl key value lh q lhL hwfl hlh
          (fun o ho => hpath o (List.mem_append.mpr (Or.inl ho))) hl hpm with ⟨hk, hmem⟩ | hcol
        · exact Or.inl ⟨by rw [hk]; rfl, Or.inl hmem⟩
        · exact Or.inr hcol
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpe (htopimg.trans hrh.symm))

/-- **Neighbor walk → divergence node (IAVL).** Two verifying existence proofs
whose reversed paths share a root-side prefix and then diverge as a left-step
(with right-most / left-most remainders) navigate to a common node `N` of the
honest IAVL tree: the left proof reaches `maxKey N.left`, the right proof
`minKey N.right` — up to a hash collision. -/
theorem neighbor_divergenceI (H : HashFn) (s : ProofSpec) (b : UInt8) (cs d : Nat)
    (hH : FixedHash H d)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hcsd : cs = d + 1)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmaxcs : s.innerSpec.maxPrefixLength < s.innerSpec.childSize)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (lhLeaf rhLeaf root : Bytes)
      (leftKey leftVal rightKey rightVal : Bytes)
      (pathL pathR : List InnerOp)
      (topLeft topRight : InnerOp) (restL restR : List InnerOp),
      WFTreeI s b t →
      applyLeaf H s.leafSpec leftKey leftVal = some lhLeaf →
      applyLeaf H s.leafSpec rightKey rightVal = some rhLeaf →
      (∀ op ∈ pathL, ensureInner op s = true) →
      (∀ op ∈ pathR, ensureInner op s = true) →
      rootHash H t = some root →
      applyPath H s.innerSpec lhLeaf pathL = some root →
      applyPath H s.innerSpec rhLeaf pathR = some root →
      dropCommonPrefix pathL.reverse pathR.reverse = (some topLeft, restL, some topRight, restR) →
      isLeftStep s.innerSpec topLeft topRight = true →
      ensureRightMost s.innerSpec restL.reverse = true →
      ensureLeftMost s.innerSpec restR.reverse = true →
      (∃ ihN preN midN lN rN, IsSubtree (.node ihN preN midN [] lN rN) t ∧
        (leftKey = maxKey lN ∨ HashCollision H) ∧
        (rightKey = minKey rN ∨ HashCollision H))
      ∨ HashCollision H := by
  intro t
  induction t with
  | leaf top tk tv =>
    intro lhLeaf rhLeaf root leftKey leftVal rightKey rightVal pathL pathR
      topLeft topRight restL restR
      hwf hlhL hlhR hpathL hpathR hrh hapL hapR hdcp hls hrm hlm
    cases pathL with
    | nil =>
      exfalso
      rw [List.reverse_nil] at hdcp
      rw [show dropCommonPrefix [] pathR.reverse
          = ((none : Option InnerOp), ([] : List InnerOp), (none : Option InnerOp),
             ([] : List InnerOp)) from rfl] at hdcp
      exact absurd (congrArg Prod.fst hdcp) (by simp)
    | cons op rest =>
      simp only [WFTreeI] at hwf; subst hwf
      rw [rootHash] at hrh
      have hii : IsInnerImage H s root :=
        applyPath_result_isInnerImage H s (op :: rest) lhLeaf root (by simp) hpathL hapL
      exact Or.inr (leafHash_innerImage_collision H s s.leafSpec tk tv root b
        hihash.symm hlpre hmin (ensureLeaf_self s.leafSpec) hrh hii)
  | node ih pre mid suf l r ihl ihr =>
    intro lhLeaf rhLeaf root leftKey leftVal rightKey rightVal pathL pathR
      topLeft topRight restL restR
      hwf hlhL hlhR hpathL hpathR hrh hapL hapR hdcp hls hrm hlm
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
    -- both paths must be nonempty (else a leaf hash equals the node hash)
    rcases List.eq_nil_or_concat pathL with hLnil | ⟨qL, topOpL, hLc⟩
    · subst hLnil
      simp only [applyPath, Option.some.injEq] at hapL
      obtain ⟨tail, hlheq⟩ := applyLeaf_head H s.leafSpec leftKey leftVal lhLeaf b hlpre hlhL
      rw [← hihash] at hlheq
      refine Or.inr (leaf_inner_domain_collision H s.innerSpec.hash (b :: tail)
        (pre ++ lhL ++ mid ++ rhR) b (by simp) ?_ ?_)
      · cases hpc2 : pre with
        | nil => exact absurd hpc2 hprene
        | cons x xs => rw [hpc2] at hprehead; simpa using hprehead
      · rw [← hlheq, hapL]; exact hrh.symm
    rcases List.eq_nil_or_concat pathR with hRnil | ⟨qR, topOpR, hRc⟩
    · subst hRnil
      simp only [applyPath, Option.some.injEq] at hapR
      obtain ⟨tail, hlheq⟩ := applyLeaf_head H s.leafSpec rightKey rightVal rhLeaf b hlpre hlhR
      rw [← hihash] at hlheq
      refine Or.inr (leaf_inner_domain_collision H s.innerSpec.hash (b :: tail)
        (pre ++ lhL ++ mid ++ rhR) b (by simp) ?_ ?_)
      · cases hpc2 : pre with
        | nil => exact absurd hpc2 hprene
        | cons x xs => rw [hpc2] at hprehead; simpa using hprehead
      · rw [← hlheq, hapR]; exact hrh.symm
    rw [List.concat_eq_append] at hLc hRc; subst hLc; subst hRc
    simp only [List.reverse_append, List.reverse_singleton, List.singleton_append] at hdcp
    obtain ⟨mL, hpmL, htopL, _⟩ := (applyPath_snoc H s.innerSpec qL topOpL lhLeaf root).mp hapL
    obtain ⟨mR, hpmR, htopR, _⟩ := (applyPath_snoc H s.innerSpec qR topOpR rhLeaf root).mp hapR
    have hmLlen : mL.length = d := applyPath_len H s.innerSpec d hH qL lhLeaf mL
      (applyLeaf_len H d hH s.leafSpec leftKey leftVal lhLeaf hlhL) hpmL
    have hmRlen : mR.length = d := applyPath_len H s.innerSpec d hH qR rhLeaf mR
      (applyLeaf_len H d hH s.leafSpec rightKey rightVal rhLeaf hlhR) hpmR
    have htopimgL := applyInner_image H topOpL mL root htopL
    rw [ensureInner_hash topOpL s (hpathL topOpL (by simp))] at htopimgL
    have htopimgR := applyInner_image H topOpR mR root htopR
    rw [ensureInner_hash topOpR s (hpathR topOpR (by simp))] at htopimgR
    have hlhLlen := rootHash_len H d hH l lhL hl
    have hrhRlen := rootHash_len H d hH r rhR hr
    unfold dropCommonPrefix at hdcp
    by_cases heq : eqPS topOpL topOpR = true
    · -- common root step: both proofs navigate into the SAME child; recurse
      rw [if_pos heq] at hdcp
      unfold eqPS at heq
      simp only [Bool.and_eq_true, beq_iff_eq] at heq
      obtain ⟨hpreEq, hsufEq⟩ := heq
      by_cases hpeL : topOpL.prefixBytes ++ mL ++ topOpL.suffix = pre ++ lhL ++ mid ++ rhR
      · by_cases hpeR : topOpR.prefixBytes ++ mR ++ topOpR.suffix = pre ++ lhL ++ mid ++ rhR
        · -- both top ops genuinely match the node split: mL = mR
          have hcat : topOpL.prefixBytes ++ mL = topOpL.prefixBytes ++ mR := by
            have hee : (topOpL.prefixBytes ++ mL) ++ topOpL.suffix
                = (topOpL.prefixBytes ++ mR) ++ topOpL.suffix := by
              rw [hpeL]; rw [hpreEq, hsufEq] at *; rw [hpeR]
            exact (List.append_inj hee (by simp [hmLlen, hmRlen])).1
          have hmEq : mL = mR := List.append_cancel_left hcat
          have hsufmod := ensureInner_suffix_mod topOpL s cs hcsspec
            (hpathL topOpL (by simp))
          rcases split_pins_var pre lhL mid rhR topOpL.prefixBytes mL topOpL.suffix cs d
            hpeL hlhLlen hrhRlen hmLlen (by omega) hprelt hsufmod with ⟨_, hmlL, _⟩ | ⟨_, hmrL, _⟩
          · rw [hmlL] at hpmL
            rw [← hmEq, hmlL] at hpmR
            rcases ihl lhLeaf rhLeaf lhL leftKey leftVal rightKey rightVal qL qR
              topLeft topRight restL restR hwfl hlhL hlhR
              (fun o ho => hpathL o (List.mem_append.mpr (Or.inl ho)))
              (fun o ho => hpathR o (List.mem_append.mpr (Or.inl ho)))
              hl hpmL hpmR hdcp hls hrm hlm with hex | hc
            · obtain ⟨ihN, preN, midN, lN, rN, hsub, h1, h2⟩ := hex
              exact Or.inl ⟨ihN, preN, midN, lN, rN,
                by simp only [IsSubtree]; exact Or.inr (Or.inl hsub), h1, h2⟩
            · exact Or.inr hc
          · rw [hmrL] at hpmL
            rw [← hmEq, hmrL] at hpmR
            rcases ihr lhLeaf rhLeaf rhR leftKey leftVal rightKey rightVal qL qR
              topLeft topRight restL restR hwfr hlhL hlhR
              (fun o ho => hpathL o (List.mem_append.mpr (Or.inl ho)))
              (fun o ho => hpathR o (List.mem_append.mpr (Or.inl ho)))
              hr hpmL hpmR hdcp hls hrm hlm with hex | hc
            · obtain ⟨ihN, preN, midN, lN, rN, hsub, h1, h2⟩ := hex
              exact Or.inl ⟨ihN, preN, midN, lN, rN,
                by simp only [IsSubtree]; exact Or.inr (Or.inr hsub), h1, h2⟩
            · exact Or.inr hc
        · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpeR (htopimgR.trans hrh.symm))
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpeL (htopimgL.trans hrh.symm))
    · -- divergence here: this node IS the neighbor node N
      rw [if_neg heq] at hdcp
      simp only [Prod.mk.injEq, Option.some.injEq] at hdcp
      obtain ⟨htL, hrestL, htR, hrestR⟩ := hdcp
      subst htL; subst htR
      rw [← hrestL, List.reverse_reverse] at hrm
      rw [← hrestR, List.reverse_reverse] at hlm
      obtain ⟨hsufL, hsufR⟩ := isLeftStep_binary s.innerSpec cs hco hcsspec topOpL topOpR hls
      by_cases hpeL : topOpL.prefixBytes ++ mL ++ topOpL.suffix = pre ++ lhL ++ mid ++ rhR
      · by_cases hpeR : topOpR.prefixBytes ++ mR ++ topOpR.suffix = pre ++ lhL ++ mid ++ rhR
        · obtain ⟨_, hmlL, _⟩ := split_var_left pre lhL mid rhR topOpL.prefixBytes mL
            topOpL.suffix cs d hpeL hlhLlen hrhRlen hmLlen (by omega) hsufL
          rw [List.length_eq_zero_iff.mp hsufR] at hpeR
          obtain ⟨_, hmrR⟩ := split_var_right pre lhL mid rhR topOpR.prefixBytes mR d
            hpeR hrhRlen hmRlen
          rw [hmlL] at hpmL
          rw [hmrR] at hpmR
          have hqLspec : ∀ op ∈ qL, ensureInner op s = true ∧ op.suffix = [] := fun op hop =>
            ⟨hpathL op (List.mem_append.mpr (Or.inl hop)),
             ensureRightMost_suffix_nil s.innerSpec cs hco hcsspec hec qL hrm op hop⟩
          have hqRspec : ∀ op ∈ qR, ensureInner op s = true ∧ op.suffix.length = cs := fun op hop =>
            ⟨hpathR op (List.mem_append.mpr (Or.inl hop)),
             ensureLeftMost_suffix_cs s.innerSpec cs hco hcsspec hec qR hlm op hop⟩
          refine Or.inl ⟨s.innerSpec.hash, pre, mid, l, r, Or.inl rfl, ?_, ?_⟩
          · rcases reaches_maxI H s b d hH hihash hlpre hmin hLInj
              l leftKey leftVal lhLeaf qL lhL hwfl hlhL hqLspec hl hpmL with ⟨hk, _⟩ | hc
            · exact Or.inl hk
            · exact Or.inr hc
          · rcases reaches_minI H s b cs d hH hcsd hihash hlpre hmin hLInj
              r rightKey rightVal rhLeaf qR rhR hwfr hlhR hqRspec hr hpmR with ⟨hk, _⟩ | hc
            · exact Or.inl hk
            · exact Or.inr hc
        · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpeR (htopimgR.trans hrh.symm))
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpeL (htopimgL.trans hrh.symm))

/-- **Theorem B (non-existence soundness), IAVL tree model — two-sided.** For a
key-sorted honest IAVL tree, a two-sided non-existence proof for `key` and an
existence proof for `key` cannot both verify without a hash collision. -/
theorem nonexistence_soundI (H : HashFn) (s : ProofSpec) (b : UInt8) (cs d : Nat)
    (hH : FixedHash H d) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hcsd : cs = d + 1)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmaxcs : s.innerSpec.maxPrefixLength < s.innerSpec.childSize)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hkfc : ∀ k, keyForComparison H s k = k)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : MTree) (hwf : WFTreeI s b t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep lp rp : ExistenceProof)
    (hnl : nep.left = some lp) (hnr : nep.right = some rp)
    (hepleaf : ep.leaf = s.leafSpec) (hlpleaf : lp.leaf = s.leafSpec) (hrpleaf : rp.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  -- the absent key is a genuine member of the honest tree (or a collision)
  rcases membership_soundI H s b cs d hH hcs hcsspec hcsd hlhash hlpre hmin hmaxcs hLInj
    t ep root key value hwf hepleaf hroot hex with hmem | hc
  · -- bracketing proofs verify and bracket the key
    obtain ⟨hlver, hllt⟩ := verifyNonExistence_left H s root key nep lp hnl hne
    obtain ⟨hrver, hrlt⟩ := verifyNonExistence_right H s root key nep rp hnr hne
    have hnb : ensureLeftNeighbor s.innerSpec lp.path rp.path = true := by
      unfold verifyNonExistence at hne
      simp only [hnl, hnr, Bool.and_eq_true] at hne
      exact hne.2
    obtain ⟨lhLeaf, hlapL, hlap, hlinn⟩ := verifyExistence_navigates H s lp root hlpleaf hlver
    obtain ⟨rhLeaf, hrapL, hrap, hrinn⟩ := verifyExistence_navigates H s rp root hrpleaf hrver
    obtain ⟨topLeft, restL, topRight, restR, hdcp, hls, hrm, hlm⟩ :=
      ensureLeftNeighbor_spec s.innerSpec lp.path rp.path hnb
    rcases neighbor_divergenceI H s b cs d hH hcsspec hcsd hlhash.symm hlpre hmin hmaxcs hco hec
      hLInj t lhLeaf rhLeaf root lp.key lp.value rp.key rp.value lp.path rp.path
      topLeft topRight restL restR
      hwf hlapL hrapL hlinn hrinn hroot hlap hrap hdcp hls hrm hlm with hdiv | hc
    · obtain ⟨ihN, preN, midN, lN, rN, hsub, hLeq, hReq⟩ := hdiv
      rcases hLeq with hLeq | hc
      · rcases hReq with hReq | hc
        · -- maxKey N.left < key < minKey N.right, but key is a member: contradiction
          rw [hkfc, hkfc] at hllt hrlt
          exact (node_gap_no_member t hsort ihN preN midN [] lN rN hsub key value hmem
            (by rw [← hLeq]; exact hllt) (by rw [← hReq]; exact hrlt)).elim
        · exact hc
      · exact hc
    · exact hc
  · exact hc

/-- **Theorem B, IAVL tree model — left-only.** The lone left neighbor's
right-most path makes it the rightmost leaf (`reaches_maxI`), so
`maxKey t < key` contradicts `member_le_maxKey` for a verifying existence
proof. -/
theorem nonexistence_soundI_leftOnly (H : HashFn) (s : ProofSpec) (b : UInt8) (cs d : Nat)
    (hH : FixedHash H d) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hcsd : cs = d + 1)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmaxcs : s.innerSpec.maxPrefixLength < s.innerSpec.childSize)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hkfc : ∀ k, keyForComparison H s k = k)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : MTree) (hwf : WFTreeI s b t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep lp : ExistenceProof)
    (hnl : nep.left = some lp) (hnr : nep.right = none)
    (hepleaf : ep.leaf = s.leafSpec) (hlpleaf : lp.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  rcases membership_soundI H s b cs d hH hcs hcsspec hcsd hlhash hlpre hmin hmaxcs hLInj
    t ep root key value hwf hepleaf hroot hex with hmem | hc
  · obtain ⟨hlver, hllt⟩ := verifyNonExistence_left H s root key nep lp hnl hne
    have hrm := verifyNonExistence_leftOnly H s root key nep lp hnl hnr hne
    obtain ⟨lhLeaf, hlapL, hlap, hlinn⟩ := verifyExistence_navigates H s lp root hlpleaf hlver
    rcases reaches_maxI H s b d hH hlhash.symm hlpre hmin hLInj
      t lp.key lp.value lhLeaf lp.path root hwf hlapL
      (fun op hop => ⟨hlinn op hop,
        ensureRightMost_suffix_nil s.innerSpec cs hco hcsspec hec lp.path hrm op hop⟩)
      hroot hlap with ⟨hk, _⟩ | hc
    · rw [hkfc, hkfc] at hllt
      exact (ble_not_gt key (maxKey t)
        (member_le_maxKey t key value hsort hmem) (hk ▸ hllt)).elim
    · exact hc
  · exact hc

/-- **Theorem B, IAVL tree model — right-only.** Mirror of
`nonexistence_soundI_leftOnly` via `reaches_minI` / `minKey_le_member`. -/
theorem nonexistence_soundI_rightOnly (H : HashFn) (s : ProofSpec) (b : UInt8) (cs d : Nat)
    (hH : FixedHash H d) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hcsd : cs = d + 1)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmaxcs : s.innerSpec.maxPrefixLength < s.innerSpec.childSize)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hkfc : ∀ k, keyForComparison H s k = k)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : MTree) (hwf : WFTreeI s b t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep rp : ExistenceProof)
    (hnl : nep.left = none) (hnr : nep.right = some rp)
    (hepleaf : ep.leaf = s.leafSpec) (hrpleaf : rp.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  rcases membership_soundI H s b cs d hH hcs hcsspec hcsd hlhash hlpre hmin hmaxcs hLInj
    t ep root key value hwf hepleaf hroot hex with hmem | hc
  · obtain ⟨hrver, hrlt⟩ := verifyNonExistence_right H s root key nep rp hnr hne
    have hlm := verifyNonExistence_rightOnly H s root key nep rp hnl hnr hne
    obtain ⟨rhLeaf, hrapL, hrap, hrinn⟩ := verifyExistence_navigates H s rp root hrpleaf hrver
    rcases reaches_minI H s b cs d hH hcsd hlhash.symm hlpre hmin hLInj
      t rp.key rp.value rhLeaf rp.path root hwf hrapL
      (fun op hop => ⟨hrinn op hop,
        ensureLeftMost_suffix_cs s.innerSpec cs hco hcsspec hec rp.path hlm op hop⟩)
      hroot hrap with ⟨hk, _⟩ | hc
    · rw [hkfc, hkfc] at hrlt
      exact (ble_not_gt (minKey t) key
        (minKey_le_member t key value hsort hmem) (hk ▸ hrlt)).elim
    · exact hc
  · exact hc

/-- **Theorem B, IAVL tree model — total.** Every verifier-accepted proof shape
is covered: two-sided, left-only, right-only (a no-neighbor proof never
verifies). -/
theorem nonexistence_soundI_total (H : HashFn) (s : ProofSpec) (b : UInt8) (cs d : Nat)
    (hH : FixedHash H d) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hcsd : cs = d + 1)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmaxcs : s.innerSpec.maxPrefixLength < s.innerSpec.childSize)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hkfc : ∀ k, keyForComparison H s k = k)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : MTree) (hwf : WFTreeI s b t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep : ExistenceProof)
    (hepleaf : ep.leaf = s.leafSpec)
    (hlleaf : ∀ p, nep.left = some p → p.leaf = s.leafSpec)
    (hrleaf : ∀ p, nep.right = some p → p.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  cases hnl : nep.left with
  | none =>
    cases hnr : nep.right with
    | none =>
      rw [verifyNonExistence_none H s root key nep hnl hnr] at hne
      exact Bool.noConfusion hne
    | some rp =>
      exact nonexistence_soundI_rightOnly H s b cs d hH hcs hcsspec hcsd hlhash hlpre hmin
        hmaxcs hco hec hkfc hLInj t hwf hsort root key value hroot nep ep rp hnl hnr hepleaf
        (hrleaf rp hnr) hne hex
  | some lp =>
    cases hnr : nep.right with
    | none =>
      exact nonexistence_soundI_leftOnly H s b cs d hH hcs hcsspec hcsd hlhash hlpre hmin
        hmaxcs hco hec hkfc hLInj t hwf hsort root key value hroot nep ep lp hnl hnr hepleaf
        (hlleaf lp hnl) hne hex
    | some rp =>
      exact nonexistence_soundI H s b cs d hH hcs hcsspec hcsd hlhash hlpre hmin hmaxcs hco
        hec hkfc hLInj t hwf hsort root key value hroot nep ep lp rp hnl hnr hepleaf
        (hlleaf lp hnl) (hrleaf rp hnr) hne hex

/-- **Theorem B for the IAVL spec — total.** All structural side conditions
discharged by computation; joint leaf injectivity proved (`leafInj_iavl`). The
only remaining hypothesis is collision resistance (`FixedHash`). With this,
honest-root Theorems A and B hold for **all three shipped specs**. -/
theorem nonexistence_sound_iavl_total (H : HashFn)
    (hH : FixedHash H 32)
    (t : MTree) (hwf : WFTreeI iavlSpec 0 t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep : ExistenceProof)
    (hepleaf : ep.leaf = iavlSpec.leafSpec)
    (hlleaf : ∀ p, nep.left = some p → p.leaf = iavlSpec.leafSpec)
    (hrleaf : ∀ p, nep.right = some p → p.leaf = iavlSpec.leafSpec)
    (hne : verifyNonExistence H nep iavlSpec root key = true)
    (hex : verifyExistence H ep iavlSpec root key value = true) :
    HashCollision H :=
  nonexistence_soundI_total H iavlSpec 0 33 32 hH (by decide) (by decide) (by decide)
    (by decide) (by decide) (by decide) (by decide) (by decide) (by decide)
    (fun k => by simp [keyForComparison, iavlSpec]) (leafInj_iavl H)
    t hwf hsort root key value hroot nep ep hepleaf hlleaf hrleaf hne hex

end Ics23
