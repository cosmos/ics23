/-
Honest-root SMT/JMT tree model.

The `MTree` model (Tree.lean) cannot represent sparse-merkle trees: an SMT node
may have an *empty* subtree whose digest is not a hash image but the constant
`empty_child` placeholder (for `smt_spec`, 32 zero bytes — the same length as a
real digest, which is exactly why the Tendermint-shape bridge lemmas
`ensureRightMost_suffix_nil` / `ensureLeftMost_suffix_cs` exclude it). This file
models that: `SMTree` adds an `.empty` constructor, and `rootHashS` maps it to
the placeholder constant `ec`.

The price is one new, explicitly-stated assumption beyond collision resistance:
`EmptyChildFree` — the spec's hash op has no (exhibited) preimage of the
placeholder. For SHA-256 and the all-zero placeholder this is the standard
sparse-merkle-tree assumption; without it a prover who knows `x` with
`H(x) = 0^32` could graft a fake empty subtree over a real one.

With it, Theorem A (`membership_soundS`) transfers: an accepted existence proof
against an honest SMT root reaches a genuine leaf, because the verifier's split
of each node is pinned to a real child (`split_pins`, reused byte-level), and a
path can never terminate inside an empty subtree (its digest would be a hash
image equal to the placeholder).
-/
import Ics23.Tree
import Ics23.TreeNonExist

namespace Ics23

/-- A binary sparse-merkle tree: like `MTree` but subtrees may be empty, and
node hashing has no mid/suffix bytes (`node = H ih (pre ++ lh ++ rh)`, the
SMT/JMT shape). -/
inductive SMTree where
  | empty : SMTree
  | leaf : LeafOp → Bytes → Bytes → SMTree
  | node : HashOp → Bytes → SMTree → SMTree → SMTree
  deriving Inhabited

/-- Root hash with placeholder `ec`: an empty subtree's digest is the constant
`ec` (`empty_child`), not a hash image. -/
def rootHashS (H : HashFn) (ec : Bytes) : SMTree → Option Bytes
  | .empty => some ec
  | .leaf op k v => applyLeaf H op k v
  | .node ih pre l r =>
    match rootHashS H ec l, rootHashS H ec r with
    | some lh, some rh => some (H ih (pre ++ lh ++ rh))
    | _, _ => none

/-- `(key, value)` is a leaf of the tree. -/
def TreeMemberS (key val : Bytes) : SMTree → Prop
  | .empty => False
  | .leaf _ k v => k = key ∧ v = val
  | .node _ _ l r => TreeMemberS key val l ∨ TreeMemberS key val r

/-- SMT-shaped conformance to a spec: nodes use the inner-spec hash with a
fixed-length, non-leaf-prefixed prefix; leaves carry exactly the spec leaf op;
empty subtrees are unconstrained. -/
def WFTreeS (s : ProofSpec) (b : UInt8) : SMTree → Prop
  | .empty => True
  | .leaf op _ _ => op = s.leafSpec
  | .node ih pre l r =>
      ih = s.innerSpec.hash ∧
      (pre.length : Int) = s.innerSpec.minPrefixLength ∧
      pre ≠ [] ∧ pre.head? ≠ some b ∧
      WFTreeS s b l ∧ WFTreeS s b r

/-- The SMT placeholder assumption: the hash op `op` has no exhibited preimage
of the placeholder `ec`. For SHA-256 and `ec = 0^32` this is the standard
sparse-merkle assumption (finding such a preimage is a break of preimage
resistance at a fixed target). -/
def EmptyChildFree (H : HashFn) (op : HashOp) (ec : Bytes) : Prop :=
  ∀ x, H op x ≠ ec

/-- A successful `applyLeaf` output is an image of the leaf's hash op. -/
theorem applyLeaf_image (H : HashFn) (leaf : LeafOp) (k v lh : Bytes)
    (h : applyLeaf H leaf k v = some lh) : ∃ x, H leaf.hash x = lh := by
  unfold applyLeaf at h
  cases hpk : prepareLeafData H leaf.prehashKey leaf.length k with
  | none => simp [hpk] at h
  | some pk =>
  cases hpv : prepareLeafData H leaf.prehashValue leaf.length v with
  | none => simp [hpk, hpv] at h
  | some pv =>
  rw [hpk, hpv] at h; simp only [Option.some.injEq] at h
  exact ⟨_, h⟩

/-- Folding spec-conformant inner ops preserves being an image of the spec's
inner hash op. -/
theorem applyPath_image (H : HashFn) (s : ProofSpec) :
    ∀ (path : List InnerOp) (lh res : Bytes),
      (∃ x, H s.innerSpec.hash x = lh) →
      (∀ op ∈ path, ensureInner op s = true) →
      applyPath H s.innerSpec lh path = some res →
      ∃ x, H s.innerSpec.hash x = res := by
  intro path
  induction path with
  | nil =>
    intro lh res him _ hap
    simp only [applyPath, Option.some.injEq] at hap
    exact hap ▸ him
  | cons op rest ih =>
    intro lh res him hall hap
    simp only [applyPath] at hap
    cases hA : applyInner H op lh with
    | none => simp [hA] at hap
    | some h' =>
      simp only [hA] at hap
      by_cases g : ((h'.length : Int) > s.innerSpec.childSize ∧ s.innerSpec.childSize ≥ 32)
      · simp [g] at hap
      · rw [if_neg g] at hap
        refine ih h' res ?_ (fun o ho => hall o (List.mem_cons_of_mem op ho)) hap
        have himg := applyInner_image H op lh h' hA
        rw [ensureInner_hash op s (hall op (List.mem_cons_self ..))] at himg
        exact ⟨_, himg⟩

/-- A leaf-rooted, spec-conformant fold never produces the placeholder. This is
the structural payoff of `EmptyChildFree`: a verifying proof cannot terminate
(or pass through) an empty subtree's digest. -/
theorem applyPath_ne_emptyChild (H : HashFn) (s : ProofSpec) (ec : Bytes)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (key value lh res : Bytes) (path : List InnerOp)
    (hlh : applyLeaf H s.leafSpec key value = some lh)
    (hall : ∀ op ∈ path, ensureInner op s = true)
    (hap : applyPath H s.innerSpec lh path = some res) :
    res ≠ ec := by
  obtain ⟨x, hx⟩ := applyLeaf_image H s.leafSpec key value lh hlh
  rw [hlhash] at hx
  obtain ⟨y, hy⟩ := applyPath_image H s path lh res ⟨x, hx⟩ hall hap
  exact fun he => hef y (hy.trans he)

/-- An SMT root hash is a `cs`-length digest (the placeholder included, since
`ec.length = cs` for the SMT shape). -/
theorem rootHashS_len (H : HashFn) (ec : Bytes) (cs : Nat) (hH : FixedHash H cs)
    (hecl : ec.length = cs) :
    ∀ (t : SMTree) (r : Bytes), rootHashS H ec t = some r → r.length = cs := by
  intro t
  induction t with
  | empty =>
    intro r h
    simp only [rootHashS, Option.some.injEq] at h
    rw [← h]; exact hecl
  | leaf op k v =>
    intro r h; exact applyLeaf_len H cs hH op k v r h
  | node ih pre l r _ _ =>
    intro root h
    unfold rootHashS at h
    cases hl : rootHashS H ec l with
    | none => rw [hl] at h; simp at h
    | some lh =>
      cases hr : rootHashS H ec r with
      | none => rw [hl, hr] at h; simp at h
      | some rh =>
        rw [hl, hr] at h; simp only [Option.some.injEq] at h
        rw [← h]; exact hH _ _

/-- Under `EmptyChildFree`, only the empty tree hashes to the placeholder: a
well-formed leaf or node digest is a hash image. -/
theorem rootHashS_ec_empty (H : HashFn) (s : ProofSpec) (b : UInt8) (ec : Bytes)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (t : SMTree) (hwf : WFTreeS s b t) (h : rootHashS H ec t = some ec) :
    t = .empty := by
  cases t with
  | empty => rfl
  | leaf op k v =>
    exfalso
    simp only [WFTreeS] at hwf; subst hwf
    obtain ⟨x, hx⟩ := applyLeaf_image H s.leafSpec k v ec h
    rw [hlhash] at hx
    exact hef x hx
  | node ih pre l r =>
    exfalso
    obtain ⟨hih, _, _, _, _, _⟩ := hwf
    rw [rootHashS] at h
    cases hl : rootHashS H ec l with
    | none => rw [hl] at h; simp at h
    | some lh =>
    cases hr : rootHashS H ec r with
    | none => rw [hl, hr] at h; simp at h
    | some rh =>
      rw [hl, hr] at h
      simp only [Option.some.injEq] at h
      rw [hih] at h
      exact hef _ h

/-- Core induction (SMT analog of `reaches`): a proof whose leaf hash folds to a
real SMT's root reaches a genuine leaf. The new case versus `MTree`: the fold
can never land in an empty subtree, because its digest is the placeholder and
the fold's value is always a hash image (`applyPath_ne_emptyChild`). -/
theorem reachesS (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat) (ec : Bytes)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hbin : s.innerSpec.childOrder.length = 2)
    (hecl : ec.length = cs)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : SMTree) (key value lh : Bytes) (path : List InnerOp) (root : Bytes),
      WFTreeS s b t →
      applyLeaf H s.leafSpec key value = some lh →
      (∀ op ∈ path, ensureInner op s = true) →
      rootHashS H ec t = some root →
      applyPath H s.innerSpec lh path = some root →
      TreeMemberS key value t ∨ HashCollision H := by
  intro t
  induction t with
  | empty =>
    intro key value lh path root _ hlh hpath hrh hap
    exfalso
    simp only [rootHashS, Option.some.injEq] at hrh
    exact applyPath_ne_emptyChild H s ec hihash.symm hef key value lh root path
      hlh hpath hap hrh.symm
  | leaf top tk tv =>
    intro key value lh path root hwf hlh hpath hrh hap
    simp only [WFTreeS] at hwf; subst hwf
    rw [rootHashS] at hrh
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
  | node ih pre l r ihl ihr =>
    intro key value lh path root hwf hlh hpath hrh hap
    obtain ⟨hih, hppre, hprene, hprehead, hwfl, hwfr⟩ := hwf
    subst hih
    rw [rootHashS] at hrh
    cases hl : rootHashS H ec l with
    | none => rw [hl] at hrh; simp at hrh
    | some lhL =>
    cases hr : rootHashS H ec r with
    | none => rw [hl, hr] at hrh; simp at hrh
    | some rhR =>
    rw [hl, hr] at hrh
    simp only [Option.some.injEq] at hrh
    rcases List.eq_nil_or_concat path with hpnil | ⟨q, topOp, hpc⟩
    · -- empty path: the proof's leaf hash equals the node hash ⇒ domain collision
      subst hpnil
      simp only [applyPath, Option.some.injEq] at hap
      obtain ⟨tail, hlheq⟩ := applyLeaf_head H s.leafSpec key value lh b hlpre hlh
      rw [← hihash] at hlheq
      refine Or.inr (leaf_inner_domain_collision H s.innerSpec.hash (b :: tail)
        (pre ++ lhL ++ rhR) b (by simp) ?_ ?_)
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
      by_cases hpe : topOp.prefixBytes ++ m ++ topOp.suffix = pre ++ lhL ++ rhR
      · have hmlen : m.length = cs :=
          applyPath_len H s.innerSpec cs hH q lh m
            (applyLeaf_len H cs hH s.leafSpec key value lh hlh) hpm
        obtain ⟨hb1, hb2, hb7⟩ := split_bounds topOp s cs pre.length hcsspec hmm hbin hppre
          (hpath topOp (by simp))
        rcases split_pins pre lhL rhR topOp.prefixBytes m topOp.suffix cs pre.length
          hpe rfl (rootHashS_len H ec cs hH hecl l lhL hl)
          (rootHashS_len H ec cs hH hecl r rhR hr)
          hmlen hcs hb1 hb2 hb7 with hml | hmr
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

/-- **Honest-root Theorem A, SMT tree model.** An accepted existence proof
against an honest SMT root names a genuine `(key, value)` leaf — up to a hash
collision, given the placeholder assumption. -/
theorem membership_soundS (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat) (ec : Bytes)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hbin : s.innerSpec.childOrder.length = 2)
    (hecl : ec.length = cs)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : SMTree) (ep : ExistenceProof) (root key value : Bytes),
      WFTreeS s b t →
      ep.leaf = s.leafSpec →
      rootHashS H ec t = some root →
      verifyExistence H ep s root key value = true →
      TreeMemberS key value t ∨ HashCollision H := by
  intro t ep root key value hwf hep hrh hver
  have hkey : ep.key = key ∧ ep.value = value := by
    unfold verifyExistence at hver
    simp only [Bool.and_eq_true] at hver
    obtain ⟨⟨⟨_, hkk⟩, hvv⟩, _⟩ := hver
    exact ⟨by simpa using hkk, by simpa using hvv⟩
  rw [← hkey.1, ← hkey.2] at hver
  obtain ⟨lh, hlf, hap, hinn⟩ := verifyExistence_navigates H s ep root hep hver
  rw [hkey.1, hkey.2] at hlf
  exact reachesS H s b cs ec hH hcs hcsspec hlhash.symm hlpre hmin hmm hbin hecl hef hLInj
    t key value lh ep.path root hwf hlf hinn hrh hap

/-- **Theorem A for the SMT/JMT spec.** Structural side conditions discharged by
computation, joint leaf injectivity proved (`leafInj_smt`). The remaining
hypotheses are the two genuine cryptographic assumptions of the sparse-merkle
construction: fixed 32-byte digests (`FixedHash`) and no exhibited preimage of
the all-zero placeholder (`EmptyChildFree`). -/
theorem membership_sound_smt (H : HashFn)
    (hH : FixedHash H 32)
    (hef : EmptyChildFree H .sha256 (List.replicate 32 0)) :
    ∀ (t : SMTree) (ep : ExistenceProof) (root key value : Bytes),
      WFTreeS smtSpec 0 t →
      ep.leaf = smtSpec.leafSpec →
      rootHashS H (List.replicate 32 0) t = some root →
      verifyExistence H ep smtSpec root key value = true →
      TreeMemberS key value t ∨ HashCollision H :=
  membership_soundS H smtSpec 0 32 (List.replicate 32 0) hH (by decide) (by decide)
    (by decide) (by decide) (by decide) (by decide) (by decide) (by decide)
    hef (leafInj_smt H hH)

end Ics23
