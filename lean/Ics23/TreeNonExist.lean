/-
Non-existence soundness (Theorem B) via the honest-root tree model.

`membership_sound` (Tree.lean) already gives that a non-existence proof's
bracketing neighbors — and any claimed key — are *genuine* members of the real
tree. What remains is the ordered-tree (BST) reasoning: in a key-sorted tree, no
member sits strictly between two position-adjacent leaves. This file builds that.

Key order is the raw byte order `bytesLt` (the Tendermint / `prehash_key_before_
comparison = false` case; the SMT case composes with the prehash and is analogous).
-/
import Ics23.Tree
import Ics23.NonExist
import Ics23.NonExistSound

namespace Ics23

/-- The rightmost leaf's key. -/
def maxKey : MTree → Bytes
  | .leaf _ k _ => k
  | .node _ _ _ _ _ r => maxKey r

/-- The leftmost leaf's key. -/
def minKey : MTree → Bytes
  | .leaf _ k _ => k
  | .node _ _ _ _ l _ => minKey l

/-- A key-sorted (BST) tree: at each node, the left subtree's max key is strictly
below the right subtree's min key, recursively. -/
def SortedTree : MTree → Prop
  | .leaf _ _ _ => True
  | .node _ _ _ _ l r => SortedTree l ∧ SortedTree r ∧ bytesLt (maxKey l) (minKey r) = true

/-- `min ≤ max` for any tree (with `≤` meaning `= ∨ bytesLt`). -/
theorem minKey_le_maxKey (t : MTree) (hs : SortedTree t) :
    minKey t = maxKey t ∨ bytesLt (minKey t) (maxKey t) = true := by
  induction t with
  | leaf _ k _ => exact Or.inl rfl
  | node _ _ _ _ l r ihl ihr =>
    obtain ⟨hsl, hsr, hlr⟩ := hs
    -- minKey node = minKey l, maxKey node = maxKey r
    have hl := ihl hsl   -- minKey l ≤ maxKey l
    have hr := ihr hsr   -- minKey r ≤ maxKey r
    refine Or.inr ?_
    -- minKey l ≤ maxKey l < minKey r ≤ maxKey r
    show bytesLt (minKey l) (maxKey r) = true
    have step1 : bytesLt (minKey l) (minKey r) = true := by
      rcases hl with h | h
      · rw [h]; exact hlr
      · exact bytesLt_trans _ _ _ h hlr
    rcases hr with h | h
    · rw [← h]; exact step1
    · exact bytesLt_trans _ _ _ step1 h

/-- A member's key is `≤` the tree's max key. -/
theorem member_le_maxKey (t : MTree) (key value : Bytes)
    (hs : SortedTree t) (hm : TreeMember key value t) :
    key = maxKey t ∨ bytesLt key (maxKey t) = true := by
  induction t with
  | leaf _ k _ =>
    simp only [TreeMember] at hm
    exact Or.inl hm.1.symm
  | node _ _ _ _ l r ihl ihr =>
    obtain ⟨hsl, hsr, hlr⟩ := hs
    simp only [TreeMember] at hm
    rcases hm with hml | hmr
    · -- in left: key ≤ maxKey l < minKey r ≤ maxKey r
      refine Or.inr ?_
      have hkl := ihl hsl hml
      have hrr := minKey_le_maxKey r hsr
      have hkr : bytesLt key (minKey r) = true := by
        rcases hkl with h | h
        · rw [h]; exact hlr
        · exact bytesLt_trans _ _ _ h hlr
      show bytesLt key (maxKey r) = true
      rcases hrr with h | h
      · rw [← h]; exact hkr
      · exact bytesLt_trans _ _ _ hkr h
    · exact ihr hsr hmr

/-- A member's key is `≥` the tree's min key. -/
theorem minKey_le_member (t : MTree) (key value : Bytes)
    (hs : SortedTree t) (hm : TreeMember key value t) :
    minKey t = key ∨ bytesLt (minKey t) key = true := by
  induction t with
  | leaf _ k _ =>
    simp only [TreeMember] at hm
    exact Or.inl hm.1
  | node _ _ _ _ l r ihl ihr =>
    obtain ⟨hsl, hsr, hlr⟩ := hs
    simp only [TreeMember] at hm
    rcases hm with hml | hmr
    · exact ihl hsl hml
    · -- in right: minKey l ≤ maxKey l < minKey r ≤ key
      refine Or.inr ?_
      have hkr := ihr hsr hmr
      have hll := minKey_le_maxKey l hsl
      have hlk : bytesLt (maxKey l) key = true := by
        rcases hkr with h | h
        · rw [← h]; exact hlr
        · exact bytesLt_trans _ _ _ hlr h
      show bytesLt (minKey l) key = true
      rcases hll with h | h
      · rw [h]; exact hlk
      · exact bytesLt_trans _ _ _ h hlk

/-- `≤` (`= ∨ bytesLt`) is transitive. -/
theorem ble_trans (a b c : Bytes) (h1 : a = b ∨ bytesLt a b = true)
    (h2 : b = c ∨ bytesLt b c = true) : a = c ∨ bytesLt a c = true := by
  rcases h1 with h1 | h1
  · rcases h2 with h2 | h2
    · exact Or.inl (h1.trans h2)
    · exact Or.inr (h1 ▸ h2)
  · rcases h2 with h2 | h2
    · exact Or.inr (h2 ▸ h1)
    · exact Or.inr (bytesLt_trans _ _ _ h1 h2)

/-- `a ≤ b` and `b < a` are contradictory. -/
theorem ble_not_gt (a b : Bytes) (h1 : a = b ∨ bytesLt a b = true)
    (h2 : bytesLt b a = true) : False := by
  rcases h1 with h | h
  · rw [h, bytesLt_irrefl] at h2; exact Bool.noConfusion h2
  · have := bytesLt_trans _ _ _ h h2; rw [bytesLt_irrefl] at this; exact Bool.noConfusion this

/-- `N` is a subtree of `t`. -/
def IsSubtree : MTree → MTree → Prop
  | N, .leaf ih k v => N = .leaf ih k v
  | N, .node ih pre mid suf l r =>
      N = .node ih pre mid suf l r ∨ IsSubtree N l ∨ IsSubtree N r

/-- A subtree's min key is `≥` the whole tree's min key. -/
theorem subtree_minKey_le (t : MTree) (hs : SortedTree t) (N : MTree) (hsub : IsSubtree N t) :
    minKey t = minKey N ∨ bytesLt (minKey t) (minKey N) = true := by
  induction t with
  | leaf ih k v => simp only [IsSubtree] at hsub; rw [hsub]; exact Or.inl rfl
  | node ih pre mid suf l r ihl ihr =>
    obtain ⟨hsl, hsr, hlr⟩ := hs
    simp only [IsSubtree] at hsub
    rcases hsub with hN | hsubl | hsubr
    · rw [hN]; exact Or.inl rfl
    · exact ihl hsl hsubl  -- minKey node = minKey l
    · -- minKey node = minKey l ≤ maxKey l < minKey r ≤ minKey N
      refine Or.inr ?_
      show bytesLt (minKey l) (minKey N) = true
      have h1 : bytesLt (minKey l) (minKey r) = true := by
        rcases minKey_le_maxKey l hsl with h | h
        · rw [h]; exact hlr
        · exact bytesLt_trans _ _ _ h hlr
      rcases ihr hsr hsubr with h | h
      · rw [← h]; exact h1
      · exact bytesLt_trans _ _ _ h1 h

/-- A subtree's max key is `≤` the whole tree's max key. -/
theorem subtree_maxKey_ge (t : MTree) (hs : SortedTree t) (N : MTree) (hsub : IsSubtree N t) :
    maxKey N = maxKey t ∨ bytesLt (maxKey N) (maxKey t) = true := by
  induction t with
  | leaf ih k v => simp only [IsSubtree] at hsub; rw [hsub]; exact Or.inl rfl
  | node ih pre mid suf l r ihl ihr =>
    obtain ⟨hsl, hsr, hlr⟩ := hs
    simp only [IsSubtree] at hsub
    rcases hsub with hN | hsubl | hsubr
    · rw [hN]; exact Or.inl rfl
    · -- maxKey N ≤ maxKey l < minKey r ≤ maxKey r = maxKey node
      refine Or.inr ?_
      show bytesLt (maxKey N) (maxKey r) = true
      have h1 : bytesLt (maxKey N) (minKey r) = true := by
        rcases ihl hsl hsubl with h | h
        · rw [h]; exact hlr
        · exact bytesLt_trans _ _ _ h hlr
      rcases minKey_le_maxKey r hsr with h | h
      · rw [← h]; exact h1
      · exact bytesLt_trans _ _ _ h1 h
    · exact ihr hsr hsubr  -- maxKey node = maxKey r

/-- **BST gap (root case).** In a sorted node, no member sits strictly between
the left subtree's max key and the right subtree's min key. This is the core
ordered-tree fact behind non-existence: adjacent leaves have no member between. -/
theorem root_gap_no_member (ih : HashOp) (pre mid suf : Bytes) (l r : MTree)
    (hs : SortedTree (.node ih pre mid suf l r))
    (key' value' : Bytes) (hm : TreeMember key' value' (.node ih pre mid suf l r))
    (h1 : bytesLt (maxKey l) key' = true) (h2 : bytesLt key' (minKey r) = true) : False := by
  obtain ⟨hsl, hsr, _⟩ := hs
  simp only [TreeMember] at hm
  rcases hm with hml | hmr
  · rcases member_le_maxKey l key' value' hsl hml with h | h
    · rw [h] at h1; simp [bytesLt_irrefl] at h1
    · have := bytesLt_trans _ _ _ h1 h; simp [bytesLt_irrefl] at this
  · rcases minKey_le_member r key' value' hsr hmr with h | h
    · rw [← h] at h2; simp [bytesLt_irrefl] at h2
    · have := bytesLt_trans _ _ _ h2 h; simp [bytesLt_irrefl] at this

/-- **Subtree contiguity.** A member of a sorted tree whose key lies within a
subtree `N`'s key range is a member of `N`. (Sorted trees occupy contiguous key
ranges.) -/
theorem member_in_subtree_range (t : MTree) (hs : SortedTree t)
    (N : MTree) (hsub : IsSubtree N t) (key value : Bytes) (hm : TreeMember key value t)
    (hlo : minKey N = key ∨ bytesLt (minKey N) key = true)
    (hhi : key = maxKey N ∨ bytesLt key (maxKey N) = true) :
    TreeMember key value N := by
  induction t with
  | leaf ih k v => simp only [IsSubtree] at hsub; subst hsub; exact hm
  | node ih pre mid suf l r ihl ihr =>
    obtain ⟨hsl, hsr, hlr⟩ := hs
    simp only [IsSubtree] at hsub
    simp only [TreeMember] at hm
    rcases hsub with hN | hsubl | hsubr
    · subst hN; simp only [TreeMember]; exact hm
    · -- N ⊆ l : key ≤ maxKey N ≤ maxKey l < minKey r, so the member is in l
      have hkml : key = maxKey l ∨ bytesLt key (maxKey l) = true :=
        ble_trans _ _ _ hhi (subtree_maxKey_ge l hsl N hsubl)
      have hklt : bytesLt key (minKey r) = true := by
        rcases hkml with h | h
        · rw [h]; exact hlr
        · exact bytesLt_trans _ _ _ h hlr
      rcases hm with hml | hmr
      · exact ihl hsl hsubl hml
      · exact absurd (minKey_le_member r key value hsr hmr) (fun h => ble_not_gt _ _ h hklt)
    · -- N ⊆ r : minKey l < minKey r ≤ minKey N ≤ key, so the member is in r
      have hkmr : minKey r = key ∨ bytesLt (minKey r) key = true :=
        ble_trans _ _ _ (subtree_minKey_le r hsr N hsubr) hlo
      have hrlt : bytesLt (maxKey l) key = true := by
        rcases hkmr with h | h
        · rw [← h]; exact hlr
        · exact bytesLt_trans _ _ _ hlr h
      rcases hm with hml | hmr
      · exact absurd (member_le_maxKey l key value hsl hml) (fun h => ble_not_gt _ _ h hrlt)
      · exact ihr hsr hsubr hmr

/-- A subtree of a sorted tree is sorted. -/
theorem subtree_sorted (t : MTree) (hs : SortedTree t) (N : MTree) (hsub : IsSubtree N t) :
    SortedTree N := by
  induction t with
  | leaf _ _ _ => simp only [IsSubtree] at hsub; rw [hsub]; exact hs
  | node _ _ _ _ l r ihl ihr =>
    obtain ⟨hsl, hsr, hlr⟩ := hs
    simp only [IsSubtree] at hsub
    rcases hsub with hN | hl | hr
    · rw [hN]; exact ⟨hsl, hsr, hlr⟩
    · exact ihl hsl hl
    · exact ihr hsr hr

/-- **BST gap (general subtree).** For any subtree `N` of a sorted tree `t`, no
member of `t` sits strictly between `N`'s left-max and right-min keys. -/
theorem node_gap_no_member (t : MTree) (hs : SortedTree t)
    (ih : HashOp) (pre mid suf : Bytes) (l r : MTree)
    (hsub : IsSubtree (.node ih pre mid suf l r) t)
    (key value : Bytes) (hm : TreeMember key value t)
    (h1 : bytesLt (maxKey l) key = true) (h2 : bytesLt key (minKey r) = true) : False := by
  have hsN : SortedTree (.node ih pre mid suf l r) :=
    subtree_sorted t hs _ hsub
  have hsl := hsN.1
  have hsr := hsN.2.1
  have hlo : minKey (.node ih pre mid suf l r) = key ∨
      bytesLt (minKey (.node ih pre mid suf l r)) key = true :=
    ble_trans (minKey l) (maxKey l) key (minKey_le_maxKey l hsl) (Or.inr h1)
  have hhi : key = maxKey (.node ih pre mid suf l r) ∨
      bytesLt key (maxKey (.node ih pre mid suf l r)) = true :=
    ble_trans key (minKey r) (maxKey r) (Or.inr h2) (minKey_le_maxKey r hsr)
  have hmN := member_in_subtree_range t hs _ hsub key value hm hlo hhi
  exact root_gap_no_member ih pre mid suf l r hsN key value hmN h1 h2

/-- `ensure_inner`'s lower prefix bound. -/
theorem ensureInner_minle (op : InnerOp) (s : ProofSpec) (h : ensureInner op s = true) :
    s.innerSpec.minPrefixLength ≤ (op.prefixBytes.length : Int) := by
  unfold ensureInner at h
  simp only [Bool.and_eq_true, decide_eq_true_eq] at h
  exact h.1.1.1.1.2

/-- An all-right path (`ensure_right_most`: each step has empty suffix) reaches the
rightmost leaf — the proof's key is `maxKey t`, up to a collision. -/
theorem reaches_max (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat)
    (hH : FixedHash H cs)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (key value lh : Bytes) (path : List InnerOp) (root : Bytes),
      WFTree s b t →
      applyLeaf H s.leafSpec key value = some lh →
      (∀ op ∈ path, ensureInner op s = true ∧ op.suffix = []) →
      rootHash H t = some root →
      applyPath H s.innerSpec lh path = some root →
      (key = maxKey t ∧ TreeMember key value t) ∨ HashCollision H := by
  intro t
  induction t with
  | leaf top tk tv =>
    intro key value lh path root hwf hlh hpath hrh hap
    simp only [WFTree] at hwf; subst hwf
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
    obtain ⟨hih, hmid, hsuf, _, hprene, hprehead, _, hwfr⟩ := hwf
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
    rcases List.eq_nil_or_concat path with hpnil | ⟨q, topOp, hpc⟩
    · subst hpnil
      simp only [applyPath, Option.some.injEq] at hap
      obtain ⟨tail, hlheq⟩ := applyLeaf_head H s.leafSpec key value lh b hlpre hlh
      rw [← hihash] at hlheq
      refine Or.inr (leaf_inner_domain_collision H s.innerSpec.hash (b :: tail)
        (pre ++ lhL ++ rhR) b (by simp) ?_ ?_)
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
      by_cases hpe : topOp.prefixBytes ++ m ++ topOp.suffix = pre ++ lhL ++ rhR
      · have hmlen : m.length = cs :=
          applyPath_len H s.innerSpec cs hH q lh m
            (applyLeaf_len H cs hH s.leafSpec key value lh hlh) hpm
        have hmr : m = rhR := split_right pre lhL rhR topOp.prefixBytes m topOp.suffix cs
          hpe (rootHash_len H cs hH l lhL hl) (rootHash_len H cs hH r rhR hr) hmlen
          (by rw [(hpath topOp (by simp)).2]; rfl)
        rw [hmr] at hpm
        rcases ihr key value lh q rhR hwfr hlh
          (fun o ho => hpath o (List.mem_append.mpr (Or.inl ho))) hr hpm with ⟨hk, hmem⟩ | hcol
        · exact Or.inl ⟨by rw [hk]; rfl, Or.inr hmem⟩
        · exact Or.inr hcol
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpe (htopimg.trans hrh.symm))

/-- An all-left path (`ensure_left_most`: each step has a full `cs`-suffix) reaches
the leftmost leaf — the proof's key is `minKey t`, up to a collision. -/
theorem reaches_min (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat)
    (hH : FixedHash H cs)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (key value lh : Bytes) (path : List InnerOp) (root : Bytes),
      WFTree s b t →
      applyLeaf H s.leafSpec key value = some lh →
      (∀ op ∈ path, ensureInner op s = true ∧ op.suffix.length = cs) →
      rootHash H t = some root →
      applyPath H s.innerSpec lh path = some root →
      (key = minKey t ∧ TreeMember key value t) ∨ HashCollision H := by
  intro t
  induction t with
  | leaf top tk tv =>
    intro key value lh path root hwf hlh hpath hrh hap
    simp only [WFTree] at hwf; subst hwf
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
    obtain ⟨hih, hmid, hsuf, hppre, hprene, hprehead, hwfl, _⟩ := hwf
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
    rcases List.eq_nil_or_concat path with hpnil | ⟨q, topOp, hpc⟩
    · subst hpnil
      simp only [applyPath, Option.some.injEq] at hap
      obtain ⟨tail, hlheq⟩ := applyLeaf_head H s.leafSpec key value lh b hlpre hlh
      rw [← hihash] at hlheq
      refine Or.inr (leaf_inner_domain_collision H s.innerSpec.hash (b :: tail)
        (pre ++ lhL ++ rhR) b (by simp) ?_ ?_)
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
      by_cases hpe : topOp.prefixBytes ++ m ++ topOp.suffix = pre ++ lhL ++ rhR
      · have hmlen : m.length = cs :=
          applyPath_len H s.innerSpec cs hH q lh m
            (applyLeaf_len H cs hH s.leafSpec key value lh hlh) hpm
        have hlow : pre.length ≤ topOp.prefixBytes.length := by
          have := ensureInner_minle topOp s (hpath topOp (by simp)).1
          omega
        have hml : m = lhL := split_left pre lhL rhR topOp.prefixBytes m topOp.suffix cs
          hpe hlow (rootHash_len H cs hH l lhL hl) (rootHash_len H cs hH r rhR hr) hmlen
          (hpath topOp (by simp)).2
        rw [hml] at hpm
        rcases ihl key value lh q lhL hwfl hlh
          (fun o ho => hpath o (List.mem_append.mpr (Or.inl ho))) hl hpm with ⟨hk, hmem⟩ | hcol
        · exact Or.inl ⟨by rw [hk]; rfl, Or.inl hmem⟩
        · exact Or.inr hcol
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpe (htopimg.trans hrh.symm))

/-- Structure extracted from `ensure_left_neighbor`: after stripping the common
root-side path, the first divergent ops form a left-step, the left remainder is
right-most, and the right remainder is left-most. -/
theorem ensureLeftNeighbor_spec (isp : InnerSpec) (left right : List InnerOp)
    (h : ensureLeftNeighbor isp left right = true) :
    ∃ topLeft restL topRight restR,
      dropCommonPrefix left.reverse right.reverse
        = (some topLeft, restL, some topRight, restR) ∧
      isLeftStep isp topLeft topRight = true ∧
      ensureRightMost isp restL.reverse = true ∧
      ensureLeftMost isp restR.reverse = true := by
  unfold ensureLeftNeighbor at h
  split at h
  · next topLeft restL topRight restR heq =>
    simp only [Bool.and_eq_true] at h
    exact ⟨topLeft, restL, topRight, restR, heq, h.1.1, h.1.2, h.2⟩
  · exact absurd h (by simp)

/-- Bridge (the padding-vs-navigation subtlety, finding F4/F5): for a binary spec
whose `emptyChild` sentinel is not a full `cs`-byte digest (the Tendermint shape,
`emptyChild = []`), every step of an `ensure_right_most` path has an *empty*
suffix — i.e. it is a genuine right-pad, not an empty-branch placeholder. The
placeholder branch would require `emptyChild.length = cs`, excluded by `hec`. -/
theorem ensureRightMost_suffix_nil (isp : InnerSpec) (cs : Nat)
    (hco : isp.childOrder = [0, 1])
    (hcsspec : isp.childSize.toNat = cs)
    (hec : isp.emptyChild.length ≠ cs)
    (path : List InnerOp) (h : ensureRightMost isp path = true) :
    ∀ op, op ∈ path → op.suffix = [] := by
  intro op hop
  have hpadeq : getPadding isp (isp.childOrder.length - 1)
      = some { minPrefix := (1 : Int) * isp.childSize + isp.minPrefixLength,
               maxPrefix := (1 : Int) * isp.childSize + isp.maxPrefixLength,
               suffix := isp.childSize * ((2 : Int) - 1 - 1) } := by
    unfold getPadding; rw [hco]; rfl
  unfold ensureRightMost at h
  rw [hpadeq] at h
  rw [List.all_eq_true] at h
  have hstep := h op hop
  rw [Bool.or_eq_true] at hstep
  rcases hstep with hp | hr
  · -- genuine right-pad: suffix length equals pad.suffix = childSize * 0 = 0
    unfold hasPadding at hp
    simp only [Bool.and_eq_true, decide_eq_true_eq] at hp
    have hz : (op.suffix.length : Int) = 0 := by have := hp.2; simpa using this
    have : op.suffix.length = 0 := by exact_mod_cast hz
    exact List.length_eq_zero_iff.mp this
  · -- placeholder: forces emptyChild.length = cs, contradicting hec
    exfalso
    unfold rightBranchesAreEmpty at hr
    cases hofp : orderFromPadding isp op with
    | none => rw [hofp] at hr; simp at hr
    | some idx =>
      rw [hofp, hco] at hr
      simp only [List.length_cons, List.length_nil] at hr
      by_cases hrb : (2 - 1 - idx) = 0
      · rw [hrb] at hr; simp at hr
      · rw [if_neg hrb] at hr
        by_cases hsl : op.suffix.length = isp.childSize.toNat
        · rw [if_neg (by omega)] at hr
          have h0 : byteRange op.suffix (0 * isp.childSize.toNat) isp.childSize.toNat
              == some isp.emptyChild := by
            have := List.all_eq_true.mp hr 0 (by simp only [List.mem_range]; omega)
            simpa using this
          simp only [Nat.zero_mul, byteRange] at h0
          rw [if_pos (by omega)] at h0
          simp only [List.drop_zero, beq_iff_eq, Option.some.injEq] at h0
          rw [List.take_of_length_le (by omega)] at h0
          exact hec (by rw [← h0]; exact hsl.trans hcsspec)
        · rw [if_pos hsl] at hr; simp at hr

/-- `order_from_padding` returns an in-range branch whose padding the op matches. -/
theorem orderFromPadding_mem (isp : InnerSpec) (op : InnerOp) (idx : Nat)
    (h : orderFromPadding isp op = some idx) :
    idx < isp.childOrder.length ∧ ∃ pad, getPadding isp idx = some pad ∧ hasPadding op pad = true := by
  unfold orderFromPadding at h
  have hmem := List.mem_of_find?_eq_some h
  have hpred := List.find?_some h
  rw [List.mem_range] at hmem
  refine ⟨hmem, ?_⟩
  cases hgp : getPadding isp idx with
  | none => rw [hgp] at hpred; simp at hpred
  | some pad => rw [hgp] at hpred; exact ⟨pad, rfl, hpred⟩

/-- For a binary spec, an `is_left_step` pair has a full-`cs` left suffix and an
empty right suffix (orders 0 and 1). -/
theorem isLeftStep_binary (isp : InnerSpec) (cs : Nat)
    (hco : isp.childOrder = [0, 1]) (hcsspec : isp.childSize.toNat = cs)
    (l r : InnerOp) (h : isLeftStep isp l r = true) :
    l.suffix.length = cs ∧ r.suffix.length = 0 := by
  unfold isLeftStep at h
  cases hol : orderFromPadding isp l with
  | none => rw [hol] at h; simp at h
  | some li =>
  cases hor : orderFromPadding isp r with
  | none => rw [hol, hor] at h; simp at h
  | some ri =>
    rw [hol, hor] at h
    simp only [decide_eq_true_eq] at h
    obtain ⟨hlilt, padl, hpadl, hhpl⟩ := orderFromPadding_mem isp l li hol
    obtain ⟨hrilt, padr, hpadr, hhpr⟩ := orderFromPadding_mem isp r ri hor
    rw [hco] at hlilt hrilt
    simp only [List.length_cons, List.length_nil] at hlilt hrilt
    have hli0 : li = 0 := by omega
    have hri1 : ri = 1 := by omega
    subst hli0; subst hri1
    have hpl0 : getPadding isp 0
        = some { minPrefix := (0 : Int) * isp.childSize + isp.minPrefixLength,
                 maxPrefix := (0 : Int) * isp.childSize + isp.maxPrefixLength,
                 suffix := isp.childSize * ((2 : Int) - 1 - 0) } := by
      unfold getPadding; rw [hco]; rfl
    have hpr1 : getPadding isp 1
        = some { minPrefix := (1 : Int) * isp.childSize + isp.minPrefixLength,
                 maxPrefix := (1 : Int) * isp.childSize + isp.maxPrefixLength,
                 suffix := isp.childSize * ((2 : Int) - 1 - 1) } := by
      unfold getPadding; rw [hco]; rfl
    rw [hpl0, Option.some.injEq] at hpadl; subst hpadl
    rw [hpr1, Option.some.injEq] at hpadr; subst hpadr
    unfold hasPadding at hhpl hhpr
    simp only [Bool.and_eq_true, decide_eq_true_eq] at hhpl hhpr
    constructor
    · have hz : (l.suffix.length : Int) = isp.childSize * ((2 : Int) - 1 - 0) := hhpl.2
      have he1 : ((2 : Int) - 1 - 0) = 1 := by decide
      rw [he1, Int.mul_one] at hz
      have : l.suffix.length = isp.childSize.toNat := by omega
      rw [this, hcsspec]
    · have hz : (r.suffix.length : Int) = isp.childSize * ((2 : Int) - 1 - 1) := hhpr.2
      have he0 : ((2 : Int) - 1 - 1) = 0 := by decide
      rw [he0, Int.mul_zero] at hz
      exact_mod_cast hz

/-- Bridge (left-most, mirror of `ensureRightMost_suffix_nil`): every step of an
`ensure_left_most` path has a full `cs`-byte suffix — a genuine left-pad, not a
left-empty-branch placeholder (which would force `emptyChild.length = cs`). -/
theorem ensureLeftMost_suffix_cs (isp : InnerSpec) (cs : Nat)
    (hco : isp.childOrder = [0, 1])
    (hcsspec : isp.childSize.toNat = cs)
    (hec : isp.emptyChild.length ≠ cs)
    (path : List InnerOp) (h : ensureLeftMost isp path = true) :
    ∀ op, op ∈ path → op.suffix.length = cs := by
  intro op hop
  have hpadeq : getPadding isp 0
      = some { minPrefix := (0 : Int) * isp.childSize + isp.minPrefixLength,
               maxPrefix := (0 : Int) * isp.childSize + isp.maxPrefixLength,
               suffix := isp.childSize * ((2 : Int) - 1 - 0) } := by
    unfold getPadding; rw [hco]; rfl
  unfold ensureLeftMost at h
  rw [hpadeq] at h
  rw [List.all_eq_true] at h
  have hstep := h op hop
  rw [Bool.or_eq_true] at hstep
  rcases hstep with hp | hl
  · -- genuine left-pad: suffix length equals pad.suffix = childSize * 2 - childSize = cs
    unfold hasPadding at hp
    simp only [Bool.and_eq_true, decide_eq_true_eq] at hp
    have hz : (op.suffix.length : Int) = isp.childSize * ((2 : Int) - 1 - 0) := hp.2
    have he1 : ((2 : Int) - 1 - 0) = 1 := by decide
    rw [he1, Int.mul_one] at hz
    have : op.suffix.length = isp.childSize.toNat := by omega
    rw [this, hcsspec]
  · -- placeholder: forces emptyChild.length = cs, contradicting hec
    exfalso
    unfold leftBranchesAreEmpty at hl
    cases hofp : orderFromPadding isp op with
    | none => rw [hofp] at hl; simp at hl
    | some idx =>
      simp only [hofp] at hl
      by_cases hlb : idx = 0
      · rw [hlb] at hl; simp at hl
      · rw [if_neg hlb] at hl
        by_cases hpl : op.prefixBytes.length < idx * isp.childSize.toNat
        · rw [if_pos hpl] at hl; simp at hl
        · rw [if_neg hpl] at hl
          have hmul : isp.childSize.toNat ≤ idx * isp.childSize.toNat :=
            Nat.le_mul_of_pos_left _ (Nat.one_le_iff_ne_zero.mpr hlb)
          have h0 : byteRange op.prefixBytes
              ((op.prefixBytes.length - idx * isp.childSize.toNat) + 0 * isp.childSize.toNat)
              isp.childSize.toNat == some isp.emptyChild := by
            have := List.all_eq_true.mp hl 0 (by simp only [List.mem_range]; omega)
            simpa using this
          simp only [Nat.zero_mul, Nat.add_zero, byteRange] at h0
          rw [if_pos (by omega)] at h0
          simp only [beq_iff_eq, Option.some.injEq] at h0
          have hlen : ((op.prefixBytes.drop (op.prefixBytes.length - idx * isp.childSize.toNat)).take isp.childSize.toNat).length = isp.emptyChild.length := by rw [h0]
          rw [List.length_take, List.length_drop] at hlen
          exact hec (by omega)

/-- **Neighbor walk → divergence node.** Two verifying existence proofs whose
reversed paths share a root-side prefix and then diverge as a left-step (with
right-most / left-most remainders) navigate to a common node `N` (a subtree of the
honest tree): the left proof reaches `maxKey N.left`, the right proof reaches
`minKey N.right` — up to a hash collision. This is the byte→tree bridge: it turns
`ensure_left_neighbor`'s shape into a structural fact about where the proofs land. -/
theorem neighbor_divergence (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : MTree) (lhLeaf rhLeaf root : Bytes)
      (leftKey leftVal rightKey rightVal : Bytes)
      (pathL pathR : List InnerOp)
      (topLeft topRight : InnerOp) (restL restR : List InnerOp),
      WFTree s b t →
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
      (∃ ihN preN lN rN, IsSubtree (.node ihN preN [] [] lN rN) t ∧
        (leftKey = maxKey lN ∨ HashCollision H) ∧
        (rightKey = minKey rN ∨ HashCollision H))
      ∨ HashCollision H := by
  intro t
  induction t with
  | leaf top tk tv =>
    intro lhLeaf rhLeaf root leftKey leftVal rightKey rightVal pathL pathR topLeft topRight restL restR
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
      simp only [WFTree] at hwf; subst hwf
      rw [rootHash] at hrh
      have hii : IsInnerImage H s root :=
        applyPath_result_isInnerImage H s (op :: rest) lhLeaf root (by simp) hpathL hapL
      exact Or.inr (leafHash_innerImage_collision H s s.leafSpec tk tv root b
        hihash.symm hlpre hmin (ensureLeaf_self s.leafSpec) hrh hii)
  | node ih pre mid suf l r ihl ihr =>
    intro lhLeaf rhLeaf root leftKey leftVal rightKey rightVal pathL pathR topLeft topRight restL restR
      hwf hlhL hlhR hpathL hpathR hrh hapL hapR hdcp hls hrm hlm
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
    -- both paths must be nonempty (else a leaf hash equals the node hash)
    rcases List.eq_nil_or_concat pathL with hLnil | ⟨qL, topOpL, hLc⟩
    · subst hLnil
      simp only [applyPath, Option.some.injEq] at hapL
      obtain ⟨tail, hlheq⟩ := applyLeaf_head H s.leafSpec leftKey leftVal lhLeaf b hlpre hlhL
      rw [← hihash] at hlheq
      refine Or.inr (leaf_inner_domain_collision H s.innerSpec.hash (b :: tail)
        (pre ++ lhL ++ rhR) b (by simp) ?_ ?_)
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
        (pre ++ lhL ++ rhR) b (by simp) ?_ ?_)
      · cases hpc2 : pre with
        | nil => exact absurd hpc2 hprene
        | cons x xs => rw [hpc2] at hprehead; simpa using hprehead
      · rw [← hlheq, hapR]; exact hrh.symm
    rw [List.concat_eq_append] at hLc hRc; subst hLc; subst hRc
    simp only [List.reverse_append, List.reverse_singleton, List.singleton_append] at hdcp
    obtain ⟨mL, hpmL, htopL, _⟩ := (applyPath_snoc H s.innerSpec qL topOpL lhLeaf root).mp hapL
    obtain ⟨mR, hpmR, htopR, _⟩ := (applyPath_snoc H s.innerSpec qR topOpR rhLeaf root).mp hapR
    have hmLlen : mL.length = cs := applyPath_len H s.innerSpec cs hH qL lhLeaf mL
      (applyLeaf_len H cs hH s.leafSpec leftKey leftVal lhLeaf hlhL) hpmL
    have hmRlen : mR.length = cs := applyPath_len H s.innerSpec cs hH qR rhLeaf mR
      (applyLeaf_len H cs hH s.leafSpec rightKey rightVal rhLeaf hlhR) hpmR
    have htopimgL := applyInner_image H topOpL mL root htopL
    rw [ensureInner_hash topOpL s (hpathL topOpL (by simp))] at htopimgL
    have htopimgR := applyInner_image H topOpR mR root htopR
    rw [ensureInner_hash topOpR s (hpathR topOpR (by simp))] at htopimgR
    have hlhLlen := rootHash_len H cs hH l lhL hl
    have hrhRlen := rootHash_len H cs hH r rhR hr
    have hlowL : pre.length ≤ topOpL.prefixBytes.length := by
      have := ensureInner_minle topOpL s (hpathL topOpL (by simp)); omega
    have hlowR : pre.length ≤ topOpR.prefixBytes.length := by
      have := ensureInner_minle topOpR s (hpathR topOpR (by simp)); omega
    unfold dropCommonPrefix at hdcp
    by_cases heq : eqPS topOpL topOpR = true
    · -- common root step: both proofs navigate into the SAME child; recurse
      rw [if_pos heq] at hdcp
      unfold eqPS at heq
      simp only [Bool.and_eq_true, beq_iff_eq] at heq
      obtain ⟨hpreEq, hsufEq⟩ := heq
      by_cases hpeL : topOpL.prefixBytes ++ mL ++ topOpL.suffix = pre ++ lhL ++ rhR
      · by_cases hpeR : topOpR.prefixBytes ++ mR ++ topOpR.suffix = pre ++ lhL ++ rhR
        · -- both top ops genuinely match the node split: mL = mR
          have hcat : topOpL.prefixBytes ++ mL = topOpL.prefixBytes ++ mR := by
            have hee : (topOpL.prefixBytes ++ mL) ++ topOpL.suffix
                = (topOpL.prefixBytes ++ mR) ++ topOpL.suffix := by
              rw [hpeL]; rw [hpreEq, hsufEq] at *; rw [hpeR]
            exact (List.append_inj hee (by simp [hmLlen, hmRlen])).1
          have hmEq : mL = mR := List.append_cancel_left hcat
          obtain ⟨hb1, hb2, hb7⟩ := split_bounds topOpL s cs pre.length hcsspec hmm
            (by rw [hco]; rfl) hppre (hpathL topOpL (by simp))
          rcases split_pins pre lhL rhR topOpL.prefixBytes mL topOpL.suffix cs pre.length
            hpeL rfl hlhLlen hrhRlen hmLlen hcs hb1 hb2 hb7 with hmlL | hmrL
          · rw [hmlL] at hpmL
            rw [← hmEq, hmlL] at hpmR
            rcases ihl lhLeaf rhLeaf lhL leftKey leftVal rightKey rightVal qL qR
              topLeft topRight restL restR hwfl hlhL hlhR
              (fun o ho => hpathL o (List.mem_append.mpr (Or.inl ho)))
              (fun o ho => hpathR o (List.mem_append.mpr (Or.inl ho)))
              hl hpmL hpmR hdcp hls hrm hlm with hex | hc
            · obtain ⟨ihN, preN, lN, rN, hsub, h1, h2⟩ := hex
              exact Or.inl ⟨ihN, preN, lN, rN, by simp only [IsSubtree]; exact Or.inr (Or.inl hsub), h1, h2⟩
            · exact Or.inr hc
          · rw [hmrL] at hpmL
            rw [← hmEq, hmrL] at hpmR
            rcases ihr lhLeaf rhLeaf rhR leftKey leftVal rightKey rightVal qL qR
              topLeft topRight restL restR hwfr hlhL hlhR
              (fun o ho => hpathL o (List.mem_append.mpr (Or.inl ho)))
              (fun o ho => hpathR o (List.mem_append.mpr (Or.inl ho)))
              hr hpmL hpmR hdcp hls hrm hlm with hex | hc
            · obtain ⟨ihN, preN, lN, rN, hsub, h1, h2⟩ := hex
              exact Or.inl ⟨ihN, preN, lN, rN, by simp only [IsSubtree]; exact Or.inr (Or.inr hsub), h1, h2⟩
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
      by_cases hpeL : topOpL.prefixBytes ++ mL ++ topOpL.suffix = pre ++ lhL ++ rhR
      · by_cases hpeR : topOpR.prefixBytes ++ mR ++ topOpR.suffix = pre ++ lhL ++ rhR
        · have hmlL : mL = lhL := split_left pre lhL rhR topOpL.prefixBytes mL topOpL.suffix cs
            hpeL hlowL hlhLlen hrhRlen hmLlen hsufL
          have hmrR : mR = rhR := split_right pre lhL rhR topOpR.prefixBytes mR topOpR.suffix cs
            hpeR hlhLlen hrhRlen hmRlen hsufR
          rw [hmlL] at hpmL
          rw [hmrR] at hpmR
          have hqLspec : ∀ op ∈ qL, ensureInner op s = true ∧ op.suffix = [] := fun op hop =>
            ⟨hpathL op (List.mem_append.mpr (Or.inl hop)),
             ensureRightMost_suffix_nil s.innerSpec cs hco hcsspec hec qL hrm op hop⟩
          have hqRspec : ∀ op ∈ qR, ensureInner op s = true ∧ op.suffix.length = cs := fun op hop =>
            ⟨hpathR op (List.mem_append.mpr (Or.inl hop)),
             ensureLeftMost_suffix_cs s.innerSpec cs hco hcsspec hec qR hlm op hop⟩
          refine Or.inl ⟨s.innerSpec.hash, pre, l, r, Or.inl rfl, ?_, ?_⟩
          · rcases reaches_max H s b cs hH hihash hlpre hmin hLInj l leftKey leftVal lhLeaf qL lhL
              hwfl hlhL hqLspec hl hpmL with ⟨hk, _⟩ | hc
            · exact Or.inl hk
            · exact Or.inr hc
          · rcases reaches_min H s b cs hH hihash hlpre hmin hLInj r rightKey rightVal rhLeaf qR rhR
              hwfr hlhR hqRspec hr hpmR with ⟨hk, _⟩ | hc
            · exact Or.inl hk
            · exact Or.inr hc
        · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpeR (htopimgR.trans hrh.symm))
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpeL (htopimgL.trans hrh.symm))

/-- A verifying existence proof navigates: its leaf hash folds along its path to
the claimed root, and every path op is spec-conformant. (Extracted from the body
of `membership_sound`.) -/
theorem verifyExistence_navigates (H : HashFn) (s : ProofSpec) (p : ExistenceProof) (root : Bytes)
    (hep : p.leaf = s.leafSpec)
    (hver : verifyExistence H p s root p.key p.value = true) :
    ∃ lh, applyLeaf H s.leafSpec p.key p.value = some lh ∧
          applyPath H s.innerSpec lh p.path = some root ∧
          (∀ op ∈ p.path, ensureInner op s = true) := by
  have hr := verifyExistence_root H p s root p.key p.value hver
  cases hke : p.key.isEmpty with
  | true => simp [calculateExistenceRoot, hke] at hr
  | false =>
  cases hve : p.value.isEmpty with
  | true => simp [calculateExistenceRoot, hke, hve] at hr
  | false =>
  cases hlf : applyLeaf H p.leaf p.key p.value with
  | none => simp [calculateExistenceRoot, hke, hve, hlf] at hr
  | some lh =>
    have hap : applyPath H s.innerSpec lh p.path = some root := by
      rw [calculateExistenceRoot_eq H s p lh hke hve hlf] at hr; exact hr
    refine ⟨lh, ?_, hap, verifyExistence_inners H p s root p.key p.value hver⟩
    rw [← hep]; exact hlf

/-- **Theorem B (non-existence soundness), honest-root tree model — two-sided.**
The honest-root analog of `membership_sound`: for a key-sorted honest tree, a
two-sided non-existence proof for `key` and an existence proof for `key` cannot
both verify without a hash collision. The neighbor walk pins the bracketing
proofs to a divergence node `N`; the absent key, were it a genuine member, would
sit in the forbidden gap between `maxKey N.left` and `minKey N.right`, which
`node_gap_no_member` rules out in a sorted tree. -/
theorem nonexistence_sound_tree (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hkfc : ∀ k, keyForComparison H s k = k)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : MTree) (hwf : WFTree s b t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep lp rp : ExistenceProof)
    (hnl : nep.left = some lp) (hnr : nep.right = some rp)
    (hepleaf : ep.leaf = s.leafSpec) (hlpleaf : lp.leaf = s.leafSpec) (hrpleaf : rp.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  have hbin : s.innerSpec.childOrder.length = 2 := by rw [hco]; rfl
  -- the absent key is a genuine member of the honest tree (or a collision)
  rcases membership_sound H s b cs hH hcs hcsspec hlhash hlpre hmin hmm hbin hLInj
    t ep root key value hwf hepleaf hroot hex with hmem | hc
  · -- bracketing proofs verify and bracket the key
    obtain ⟨hlver, hllt⟩ := verifyNonExistence_left H s root key nep lp hnl hne
    obtain ⟨hrver, hrlt⟩ := verifyNonExistence_right H s root key nep rp hnr hne
    have hnb : ensureLeftNeighbor s.innerSpec lp.path rp.path = true := by
      unfold verifyNonExistence at hne
      simp only [hnl, hnr, Bool.and_eq_true] at hne
      exact hne.2
    -- both bracketing proofs navigate
    obtain ⟨lhLeaf, hlapL, hlap, hlinn⟩ := verifyExistence_navigates H s lp root hlpleaf hlver
    obtain ⟨rhLeaf, hrapL, hrap, hrinn⟩ := verifyExistence_navigates H s rp root hrpleaf hrver
    -- the neighbor walk decomposes
    obtain ⟨topLeft, restL, topRight, restR, hdcp, hls, hrm, hlm⟩ :=
      ensureLeftNeighbor_spec s.innerSpec lp.path rp.path hnb
    -- ... and pins both proofs to a divergence node N
    rcases neighbor_divergence H s b cs hH hcs hcsspec hlhash.symm hlpre hmin hmm hco hec hLInj
      t lhLeaf rhLeaf root lp.key lp.value rp.key rp.value lp.path rp.path topLeft topRight restL restR
      hwf hlapL hrapL hlinn hrinn hroot hlap hrap hdcp hls hrm hlm with hdiv | hc
    · obtain ⟨ihN, preN, lN, rN, hsub, hLeq, hReq⟩ := hdiv
      rcases hLeq with hLeq | hc
      · rcases hReq with hReq | hc
        · -- maxKey N.left < key < minKey N.right, but key is a member: contradiction
          rw [hkfc, hkfc] at hllt hrlt
          exact (node_gap_no_member t hsort ihN preN [] [] lN rN hsub key value hmem
            (by rw [← hLeq]; exact hllt) (by rw [← hReq]; exact hrlt)).elim
        · exact hc
      · exact hc
    · exact hc
  · exact hc

/-- **Theorem B for the Tendermint spec (two-sided).** Structural side conditions
discharged by computation, joint leaf injectivity now *proved*
(`leafInj_tendermint`); the single remaining hypothesis is the genuine
cryptographic assumption (`FixedHash`). Mirrors `membership_sound_tendermint`. -/
theorem nonexistence_sound_tree_tendermint (H : HashFn)
    (hH : FixedHash H 32)
    (t : MTree) (hwf : WFTree tendermintSpec 0 t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep lp rp : ExistenceProof)
    (hnl : nep.left = some lp) (hnr : nep.right = some rp)
    (hepleaf : ep.leaf = tendermintSpec.leafSpec)
    (hlpleaf : lp.leaf = tendermintSpec.leafSpec) (hrpleaf : rp.leaf = tendermintSpec.leafSpec)
    (hne : verifyNonExistence H nep tendermintSpec root key = true)
    (hex : verifyExistence H ep tendermintSpec root key value = true) :
    HashCollision H :=
  nonexistence_sound_tree H tendermintSpec 0 32 hH (by decide) (by decide) (by decide)
    (by decide) (by decide) (by decide) (by decide) (by decide)
    (fun k => by simp [keyForComparison, tendermintSpec]) (leafInj_tendermint H)
    t hwf hsort root key value hroot nep ep lp rp hnl hnr hepleaf hlpleaf hrpleaf hne hex

/-- **Theorem B, honest-root tree model — left-only.** A non-existence proof with
only a left neighbor claims `key` lies beyond the rightmost leaf: the verifier
requires the neighbor's path to be `ensure_right_most`, so the neighbor *is* the
rightmost leaf (`reaches_max`) and `maxKey t < key`. A verifying existence proof
for `key` would make it a genuine member with `key ≤ maxKey t`
(`member_le_maxKey`) — contradiction, so a hash collision occurred. -/
theorem nonexistence_sound_tree_leftOnly (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hkfc : ∀ k, keyForComparison H s k = k)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : MTree) (hwf : WFTree s b t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep lp : ExistenceProof)
    (hnl : nep.left = some lp) (hnr : nep.right = none)
    (hepleaf : ep.leaf = s.leafSpec) (hlpleaf : lp.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  have hbin : s.innerSpec.childOrder.length = 2 := by rw [hco]; rfl
  rcases membership_sound H s b cs hH hcs hcsspec hlhash hlpre hmin hmm hbin hLInj
    t ep root key value hwf hepleaf hroot hex with hmem | hc
  · obtain ⟨hlver, hllt⟩ := verifyNonExistence_left H s root key nep lp hnl hne
    have hrm := verifyNonExistence_leftOnly H s root key nep lp hnl hnr hne
    obtain ⟨lhLeaf, hlapL, hlap, hlinn⟩ := verifyExistence_navigates H s lp root hlpleaf hlver
    rcases reaches_max H s b cs hH hlhash.symm hlpre hmin hLInj
      t lp.key lp.value lhLeaf lp.path root hwf hlapL
      (fun op hop => ⟨hlinn op hop,
        ensureRightMost_suffix_nil s.innerSpec cs hco hcsspec hec lp.path hrm op hop⟩)
      hroot hlap with ⟨hk, _⟩ | hc
    · rw [hkfc, hkfc] at hllt
      exact (ble_not_gt key (maxKey t)
        (member_le_maxKey t key value hsort hmem) (hk ▸ hllt)).elim
    · exact hc
  · exact hc

/-- **Theorem B, honest-root tree model — right-only.** Mirror of
`nonexistence_sound_tree_leftOnly`: the right neighbor's `ensure_left_most` path
makes it the leftmost leaf (`reaches_min`) with `key < minKey t`, while a
verifying existence proof would force `minKey t ≤ key` (`minKey_le_member`). -/
theorem nonexistence_sound_tree_rightOnly (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hkfc : ∀ k, keyForComparison H s k = k)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : MTree) (hwf : WFTree s b t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep rp : ExistenceProof)
    (hnl : nep.left = none) (hnr : nep.right = some rp)
    (hepleaf : ep.leaf = s.leafSpec) (hrpleaf : rp.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  have hbin : s.innerSpec.childOrder.length = 2 := by rw [hco]; rfl
  rcases membership_sound H s b cs hH hcs hcsspec hlhash hlpre hmin hmm hbin hLInj
    t ep root key value hwf hepleaf hroot hex with hmem | hc
  · obtain ⟨hrver, hrlt⟩ := verifyNonExistence_right H s root key nep rp hnr hne
    have hlm := verifyNonExistence_rightOnly H s root key nep rp hnl hnr hne
    obtain ⟨rhLeaf, hrapL, hrap, hrinn⟩ := verifyExistence_navigates H s rp root hrpleaf hrver
    rcases reaches_min H s b cs hH hlhash.symm hlpre hmin hLInj
      t rp.key rp.value rhLeaf rp.path root hwf hrapL
      (fun op hop => ⟨hrinn op hop,
        ensureLeftMost_suffix_cs s.innerSpec cs hco hcsspec hec rp.path hlm op hop⟩)
      hroot hrap with ⟨hk, _⟩ | hc
    · rw [hkfc, hkfc] at hrlt
      exact (ble_not_gt (minKey t) key
        (minKey_le_member t key value hsort hmem) (hk ▸ hrlt)).elim
    · exact hc
  · exact hc

/-- **Theorem B, honest-root tree model — total.** Every proof shape the verifier
accepts is covered: two-sided, left-only, right-only (a no-neighbor proof never
verifies, `verifyNonExistence_none`). For a key-sorted honest tree, *any*
verifying non-existence proof for `key` together with a verifying existence
proof for `key` yields a hash collision. -/
theorem nonexistence_sound_tree_total (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hec : s.innerSpec.emptyChild.length ≠ cs)
    (hkfc : ∀ k, keyForComparison H s k = k)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : MTree) (hwf : WFTree s b t) (hsort : SortedTree t)
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
      exact nonexistence_sound_tree_rightOnly H s b cs hH hcs hcsspec hlhash hlpre hmin hmm
        hco hec hkfc hLInj t hwf hsort root key value hroot nep ep rp hnl hnr hepleaf
        (hrleaf rp hnr) hne hex
  | some lp =>
    cases hnr : nep.right with
    | none =>
      exact nonexistence_sound_tree_leftOnly H s b cs hH hcs hcsspec hlhash hlpre hmin hmm
        hco hec hkfc hLInj t hwf hsort root key value hroot nep ep lp hnl hnr hepleaf
        (hlleaf lp hnl) hne hex
    | some rp =>
      exact nonexistence_sound_tree H s b cs hH hcs hcsspec hlhash hlpre hmin hmm hco hec
        hkfc hLInj t hwf hsort root key value hroot nep ep lp rp hnl hnr hepleaf
        (hlleaf lp hnl) (hrleaf rp hnr) hne hex

/-- **Theorem B for the Tendermint spec — total.** All structural side conditions
discharged by computation; the only remaining hypothesis is collision resistance
(`FixedHash`). Subsumes the two-sided and one-sided cases. -/
theorem nonexistence_sound_tree_tendermint_total (H : HashFn)
    (hH : FixedHash H 32)
    (t : MTree) (hwf : WFTree tendermintSpec 0 t) (hsort : SortedTree t)
    (root key value : Bytes) (hroot : rootHash H t = some root)
    (nep : NonExistenceProof) (ep : ExistenceProof)
    (hepleaf : ep.leaf = tendermintSpec.leafSpec)
    (hlleaf : ∀ p, nep.left = some p → p.leaf = tendermintSpec.leafSpec)
    (hrleaf : ∀ p, nep.right = some p → p.leaf = tendermintSpec.leafSpec)
    (hne : verifyNonExistence H nep tendermintSpec root key = true)
    (hex : verifyExistence H ep tendermintSpec root key value = true) :
    HashCollision H :=
  nonexistence_sound_tree_total H tendermintSpec 0 32 hH (by decide) (by decide) (by decide)
    (by decide) (by decide) (by decide) (by decide) (by decide)
    (fun k => by simp [keyForComparison, tendermintSpec]) (leafInj_tendermint H)
    t hwf hsort root key value hroot nep ep hepleaf hlleaf hrleaf hne hex

end Ics23
