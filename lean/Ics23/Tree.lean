/-
Honest-root Merkle tree model (Option 1).

The byte-level theorems (`existence_binding_shaped`, etc.) conclude a disjunction
with ambiguity arms (findings F3/F5) because, against an *arbitrary* root, the
verifier accepts inner ops whose prefix/child/suffix are adversarial bytes that
merely hash correctly. In deployment the root is the hash of a tree some honest
full node actually built (the adversary controls the proof, not the root's
provenance). This file models that: an inductive ordered Merkle tree with a
`rootHash`, against which accepted proofs follow *genuine* structure.

The payoff: both F3 "readings" of a node correspond to real left/right children,
so `membership_sound` (accepted existence proof ⇒ genuine membership, up to a
collision) holds with no ambiguity arm — which strengthens Theorem A and is the
engine for Theorem B.

Binary trees (all shipped specs are binary). A node's hash mirrors how the
verifier's inner ops combine: `H ih (pre ++ leftHash ++ mid ++ rightHash ++ suf)`,
so a left-child op `{ih, pre, mid++rh++suf}` and a right-child op
`{ih, pre++lh++mid, suf}` both reconstruct it.
-/
import Ics23.Verify
import Ics23.Soundness
import Ics23.Existence
import Ics23.LeafInj

namespace Ics23

/-- A binary Merkle tree with explicit structural bytes at each node. -/
inductive MTree where
  | leaf : LeafOp → Bytes → Bytes → MTree
  | node : HashOp → Bytes → Bytes → Bytes → MTree → MTree → MTree
  deriving Inhabited

/-- The root hash of a tree, mirroring the verifier's leaf/inner hashing. -/
def rootHash (H : HashFn) : MTree → Option Bytes
  | .leaf op k v => applyLeaf H op k v
  | .node ih pre mid suf l r =>
    match rootHash H l, rootHash H r with
    | some lh, some rh => some (H ih (pre ++ lh ++ mid ++ rh ++ suf))
    | _, _ => none

/-- `(key, value)` is a leaf of the tree. -/
def TreeMember (key val : Bytes) : MTree → Prop
  | .leaf _ k v => k = key ∧ v = val
  | .node _ _ _ _ l r => TreeMember key val l ∨ TreeMember key val r

/-- The left-child inner op for a node: child is the left subtree's hash. -/
def leftChildOp (ih : HashOp) (pre mid suf rh : Bytes) : InnerOp :=
  { hash := ih, prefixBytes := pre, suffix := mid ++ rh ++ suf }

/-- The right-child inner op for a node: the left sibling sits in the prefix. -/
def rightChildOp (ih : HashOp) (pre mid suf lh : Bytes) : InnerOp :=
  { hash := ih, prefixBytes := pre ++ lh ++ mid, suffix := suf }

/-- A left-child op applied to the (nonempty) left subtree hash reproduces the
node hash. -/
theorem leftChildOp_apply (H : HashFn) (ih : HashOp) (pre mid suf lh rh : Bytes)
    (hlh : lh.isEmpty = false) :
    applyInner H (leftChildOp ih pre mid suf rh) lh
      = some (H ih (pre ++ lh ++ mid ++ rh ++ suf)) := by
  unfold applyInner leftChildOp
  rw [if_neg (by rw [hlh]; simp)]
  simp [List.append_assoc]

/-- A right-child op applied to the (nonempty) right subtree hash reproduces the
node hash. -/
theorem rightChildOp_apply (H : HashFn) (ih : HashOp) (pre mid suf lh rh : Bytes)
    (hrh : rh.isEmpty = false) :
    applyInner H (rightChildOp ih pre mid suf lh) rh
      = some (H ih (pre ++ lh ++ mid ++ rh ++ suf)) := by
  unfold applyInner rightChildOp
  rw [if_neg (by rw [hrh]; simp)]

/-- **F3 resolution (the crux).** For a binary node `pre ++ lh ++ rh` with
`|lh| = |rh| = cs`, the verifier's checks — prefix length in `[p, p+cs]` and
`suffix.length % cs = 0` — pin any accepted split of the node into a `cs`-length
child to *exactly* the two genuine children: the child is the left or right
subtree hash. There is no straddling third reading. -/
theorem split_pins (pre lh rh topPre m topSuf : Bytes) (cs p : Nat)
    (hN : topPre ++ m ++ topSuf = pre ++ lh ++ rh)
    (hpre : pre.length = p) (hlh : lh.length = cs) (hrh : rh.length = cs)
    (hm : m.length = cs) (_hcs : 0 < cs)
    (hpb1 : p ≤ topPre.length) (_hpb2 : topPre.length ≤ p + cs)
    (hsuf : topSuf.length % cs = 0) :
    m = lh ∨ m = rh := by
  have hlentot : topPre.length + m.length + topSuf.length
      = pre.length + lh.length + rh.length := by
    have h := congrArg List.length hN
    simp only [List.length_append] at h
    omega
  rw [hm, hpre, hlh, hrh] at hlentot
  have hsuflen : topSuf.length = 0 ∨ topSuf.length = cs := by
    rcases Nat.eq_zero_or_pos topSuf.length with h0 | hpos
    · exact Or.inl h0
    · refine Or.inr ?_
      have hdvd : cs ∣ topSuf.length := Nat.dvd_of_mod_eq_zero hsuf
      have hge : cs ≤ topSuf.length := Nat.le_of_dvd hpos hdvd
      omega
  rcases hsuflen with hs0 | hscs
  · -- empty suffix: the child is the right subtree
    refine Or.inr ?_
    have hsnil : topSuf = [] := List.length_eq_zero_iff.mp hs0
    rw [hsnil, List.append_nil] at hN
    have hlen' : topPre.length = (pre ++ lh).length := by
      simp only [List.length_append]; omega
    exact (List.append_inj hN hlen').2
  · -- full-cs suffix: the child is the left subtree
    refine Or.inl ?_
    simp only [List.append_assoc] at hN
    have htplen : topPre.length = pre.length := by omega
    have hA := List.append_inj hN htplen
    exact (List.append_inj hA.2 (by rw [hm, hlh])).1

/-- Tendermint-shaped conformance of a tree to a spec: nodes use the inner-spec
hash with empty `mid`/`suf` (`node = H ih (pre ++ lh ++ rh)`), a node prefix that
is nonempty and not leaf-prefixed (domain separation), and leaves whose op is
exactly the spec leaf op. Models the CometBFT simple-merkle shape. -/
def WFTree (s : ProofSpec) (b : UInt8) : MTree → Prop
  | .leaf op _ _ => op = s.leafSpec
  | .node ih pre mid suf l r =>
      ih = s.innerSpec.hash ∧ mid = [] ∧ suf = [] ∧
      (pre.length : Int) = s.innerSpec.minPrefixLength ∧
      pre ≠ [] ∧ pre.head? ≠ some b ∧
      WFTree s b l ∧ WFTree s b r

/-- `applyLeaf` / `applyInner` outputs are `cs`-length digests. -/
theorem applyLeaf_len (H : HashFn) (cs : Nat) (hH : FixedHash H cs)
    (leaf : LeafOp) (k v r : Bytes) (h : applyLeaf H leaf k v = some r) : r.length = cs := by
  unfold applyLeaf at h
  cases hpk : prepareLeafData H leaf.prehashKey leaf.length k with
  | none => simp [hpk] at h
  | some pk =>
  cases hpv : prepareLeafData H leaf.prehashValue leaf.length v with
  | none => simp [hpk, hpv] at h
  | some pv =>
  rw [hpk, hpv] at h; simp only [Option.some.injEq] at h
  rw [← h]; exact hH _ _

theorem applyInner_len (H : HashFn) (cs : Nat) (hH : FixedHash H cs)
    (op : InnerOp) (c r : Bytes) (h : applyInner H op c = some r) : r.length = cs := by
  have := applyInner_image H op c r h
  rw [← this]; exact hH _ _

/-- A tree's root hash is a `cs`-length digest. -/
theorem rootHash_len (H : HashFn) (cs : Nat) (hH : FixedHash H cs) :
    ∀ (t : MTree) (r : Bytes), rootHash H t = some r → r.length = cs := by
  intro t
  induction t with
  | leaf op k v =>
    intro r hrh; unfold rootHash at hrh; exact applyLeaf_len H cs hH op k v r hrh
  | node ih pre mid suf l r _ _ =>
    intro root hrh
    unfold rootHash at hrh
    cases hl : rootHash H l with
    | none => rw [hl] at hrh; simp at hrh
    | some lh =>
      cases hr : rootHash H r with
      | none => rw [hl, hr] at hrh; simp at hrh
      | some rh =>
        rw [hl, hr] at hrh; simp only [Option.some.injEq] at hrh
        rw [← hrh]; exact hH _ _

/-- A path fold preserves the `cs`-length digest. -/
theorem applyPath_len (H : HashFn) (isp : InnerSpec) (cs : Nat) (hH : FixedHash H cs) :
    ∀ (p : List InnerOp) (h r : Bytes), h.length = cs →
      applyPath H isp h p = some r → r.length = cs := by
  intro p
  induction p with
  | nil =>
    intro h r hh hap; simp only [applyPath, Option.some.injEq] at hap; rw [← hap]; exact hh
  | cons op rest ih =>
    intro h r hh hap
    simp only [applyPath] at hap
    cases ha : applyInner H op h with
    | none => simp [ha] at hap
    | some h' =>
      simp only [ha] at hap
      by_cases g : (h'.length : Int) > isp.childSize ∧ isp.childSize ≥ 32
      · simp [g] at hap
      · rw [if_neg g] at hap
        exact ih h' r (applyInner_len H cs hH op h h' ha) hap

/-- Any leaf op conforms to itself. -/
theorem ensureLeaf_self (l : LeafOp) : ensureLeaf l l = true := by
  unfold ensureLeaf; simp [hasPrefix_refl]

/-- The leaf preimage starts with the leaf prefix byte `b`. -/
theorem applyLeaf_head (H : HashFn) (leaf : LeafOp) (k v r : Bytes) (b : UInt8)
    (hpre : leaf.prefixBytes = [b]) (h : applyLeaf H leaf k v = some r) :
    ∃ tail, r = H leaf.hash (b :: tail) := by
  unfold applyLeaf at h
  cases hpk : prepareLeafData H leaf.prehashKey leaf.length k with
  | none => simp [hpk] at h
  | some pk =>
  cases hpv : prepareLeafData H leaf.prehashValue leaf.length v with
  | none => simp [hpk, hpv] at h
  | some pv =>
  rw [hpk, hpv] at h; simp only [Option.some.injEq] at h
  exact ⟨pk ++ pv, by rw [← h, hpre]; rfl⟩

/-- From `ensure_inner`, for a binary spec with `min = max` prefix length, the
prefix length lies in `[p, p+cs]` and the suffix length is a multiple of `cs`
(`p = min` = the node prefix length). These are exactly `split_pins`' hypotheses. -/
theorem split_bounds (op : InnerOp) (s : ProofSpec) (cs p : Nat)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hbin : s.innerSpec.childOrder.length = 2)
    (hpp : (p : Int) = s.innerSpec.minPrefixLength)
    (hen : ensureInner op s = true) :
    p ≤ op.prefixBytes.length ∧ op.prefixBytes.length ≤ p + cs
      ∧ op.suffix.length % cs = 0 := by
  unfold ensureInner at hen
  simp only [Bool.and_eq_true, decide_eq_true_eq] at hen
  obtain ⟨⟨⟨⟨⟨⟨_, _⟩, hc3⟩, hc4⟩, hc5⟩, _⟩, hc7⟩ := hen
  have hcsInt : s.innerSpec.childSize = (cs : Int) := by omega
  rw [hbin] at hc4
  refine ⟨?_, ?_, ?_⟩
  · omega
  · -- |prefix| ≤ max + (2-1)*cs = min + cs = p + cs
    have : ((op.prefixBytes.length : Int)) ≤ s.innerSpec.minPrefixLength + (cs : Int) := by
      rw [hmm]; rw [hcsInt] at hc4; push_cast at hc4 ⊢; omega
    omega
  · -- (|suffix| : Int) % childSize = 0 → |suffix| % cs = 0
    rw [hcsInt] at hc7
    have : ((op.suffix.length % cs : Nat) : Int) = 0 := by
      rw [Int.natCast_emod]; exact hc7
    exact_mod_cast this

/-- A node split with empty suffix is the *right* child (`ensure_right_most`). -/
theorem split_right (pre lh rh topPre m topSuf : Bytes) (cs : Nat)
    (hN : topPre ++ m ++ topSuf = pre ++ lh ++ rh)
    (hlh : lh.length = cs) (hrh : rh.length = cs) (hm : m.length = cs)
    (hsuf0 : topSuf.length = 0) : m = rh := by
  have hsnil : topSuf = [] := List.length_eq_zero_iff.mp hsuf0
  rw [hsnil, List.append_nil] at hN
  have hlen : topPre.length = (pre ++ lh).length := by
    have h := congrArg List.length hN
    simp only [List.length_append] at h ⊢; omega
  exact (List.append_inj hN hlen).2

/-- A node split with a full `cs`-length suffix is the *left* child
(`ensure_left_most`). -/
theorem split_left (pre lh rh topPre m topSuf : Bytes) (cs : Nat)
    (hN : topPre ++ m ++ topSuf = pre ++ lh ++ rh)
    (hpre : pre.length ≤ topPre.length) (hlh : lh.length = cs) (hrh : rh.length = cs)
    (hm : m.length = cs) (hsuf : topSuf.length = cs) : m = lh := by
  have hlen : topPre.length = pre.length := by
    have h := congrArg List.length hN
    simp only [List.length_append] at h; omega
  simp only [List.append_assoc] at hN
  exact (List.append_inj (List.append_inj hN hlen).2 (by rw [hm, hlh])).1

/-- Core induction: a proof whose leaf hash folds to a real tree's root reaches a
genuine leaf — using `split_pins` to force each step into a real child. -/
theorem reaches (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hbin : s.innerSpec.childOrder.length = 2)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (key value lh : Bytes) (path : List InnerOp) (root : Bytes),
      WFTree s b t →
      applyLeaf H s.leafSpec key value = some lh →
      (∀ op ∈ path, ensureInner op s = true) →
      rootHash H t = some root →
      applyPath H s.innerSpec lh path = some root →
      TreeMember key value t ∨ HashCollision H := by
  intro t
  induction t with
  | leaf top tk tv =>
    intro key value lh path root hwf hlh hpath hrh hap
    -- WFTree leaf ⇒ top = s.leafSpec
    simp only [WFTree] at hwf; subst hwf
    rw [rootHash] at hrh
    cases path with
    | nil =>
      simp only [applyPath, Option.some.injEq] at hap
      -- lh = root, and lh = applyLeaf .. (key,value), root = applyLeaf .. (tk,tv)
      subst hap
      rcases hLInj key value tk tv _ hlh hrh with ⟨hk, hv⟩ | hc
      · exact Or.inl ⟨hk.symm, hv.symm⟩
      · exact Or.inr hc
    | cons op rest =>
      -- nonempty path ⇒ root is an inner image; but root is a leaf hash ⇒ collision
      have hii : IsInnerImage H s root :=
        applyPath_result_isInnerImage H s (op :: rest) lh root (by simp) hpath hap
      exact Or.inr (leafHash_innerImage_collision H s s.leafSpec tk tv root b
        hihash.symm hlpre hmin (ensureLeaf_self s.leafSpec) hrh hii)
  | node ih pre mid suf l r ihl ihr =>
    intro key value lh path root hwf hlh hpath hrh hap
    obtain ⟨hih, hmid, hsuf, hppre, hprene, hprehead, hwfl, hwfr⟩ := hwf
    subst hmid; subst hsuf; subst hih
    rw [rootHash] at hrh
    cases hl : rootHash H l with
    | none => rw [hl] at hrh; simp at hrh
    | some lhL =>
    cases hr : rootHash H r with
    | none => rw [hl, hr] at hrh; simp at hrh
    | some rhR =>
    rw [hl, hr] at hrh
    simp only [List.append_nil, Option.some.injEq] at hrh
    -- hrh : H s.innerSpec.hash (pre ++ lhL ++ rhR) = root
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
          hpe rfl (rootHash_len H cs hH l lhL hl) (rootHash_len H cs hH r rhR hr)
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

theorem membership_sound (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hbin : s.innerSpec.childOrder.length = 2)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (ep : ExistenceProof) (root key value : Bytes),
      WFTree s b t →
      ep.leaf = s.leafSpec →
      rootHash H t = some root →
      verifyExistence H ep s root key value = true →
      TreeMember key value t ∨ HashCollision H := by
  intro t ep root key value hwf hep hrh hver
  have hr := verifyExistence_root H ep s root key value hver
  have hkey : ep.key = key ∧ ep.value = value := by
    unfold verifyExistence at hver
    simp only [Bool.and_eq_true] at hver
    obtain ⟨⟨⟨_, hkk⟩, hvv⟩, _⟩ := hver
    exact ⟨by simpa using hkk, by simpa using hvv⟩
  obtain ⟨hk, hv⟩ := hkey
  cases hke : ep.key.isEmpty with
  | true => simp [calculateExistenceRoot, hke] at hr
  | false =>
  cases hve : ep.value.isEmpty with
  | true => simp [calculateExistenceRoot, hke, hve] at hr
  | false =>
  cases hlf : applyLeaf H ep.leaf ep.key ep.value with
  | none => simp [calculateExistenceRoot, hke, hve, hlf] at hr
  | some lh =>
    have hap : applyPath H s.innerSpec lh ep.path = some root := by
      rw [calculateExistenceRoot_eq H s ep lh hke hve hlf] at hr; exact hr
    rw [hep, hk, hv] at hlf
    exact reaches H s b cs hH hcs hcsspec hlhash.symm hlpre hmin hmm hbin hLInj
      t key value lh ep.path root hwf hlf
      (verifyExistence_inners H ep s root key value hver) hrh hap

/-- **Honest-root Theorem A for the Tendermint spec.** All structural side
conditions are discharged by computation, and joint leaf injectivity is now
*proved* (`leafInj_tendermint`, from varint self-delimiting + the collision
escape). The single remaining hypothesis is the genuine cryptographic assumption:
a fixed 32-byte digest (`FixedHash`). -/
theorem membership_sound_tendermint (H : HashFn)
    (hH : FixedHash H 32) :
    ∀ (t : MTree) (ep : ExistenceProof) (root key value : Bytes),
      WFTree tendermintSpec 0 t →
      ep.leaf = tendermintSpec.leafSpec →
      rootHash H t = some root →
      verifyExistence H ep tendermintSpec root key value = true →
      TreeMember key value t ∨ HashCollision H :=
  membership_sound H tendermintSpec 0 32 hH (by decide) (by decide) (by decide)
    (by decide) (by decide) (by decide) (by decide) (leafInj_tendermint H)

end Ics23
