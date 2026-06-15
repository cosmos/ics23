/-
Non-existence soundness (Theorem B) for the SMT/JMT tree model.

The Tendermint-shape development (TreeNonExist.lean) excludes the SMT spec at
two points, both rooted in `emptyChild`:

* `ensureRightMost_suffix_nil` / `ensureLeftMost_suffix_cs` prove that the
  verifier's `*_branches_are_empty` placeholder arm is *unreachable* when
  `emptyChild.length ≠ childSize`. For `smt_spec` the placeholder is a full
  32-byte digest, so the arm is genuinely reachable: a right-most path may step
  *left* under a node whose right subtree is empty, and symmetrically.
* `MTree` cannot represent those empty subtrees at all (fixed in SmtTree.lean).

Here the placeholder arm is embraced instead of excluded: an `ensure_right_most`
step either has an empty suffix (a genuine right step) or a suffix equal to
`emptyChild` — which, under `EmptyChildFree`, pins the right sibling to a
genuinely *empty* subtree, so the path still reaches the rightmost (nonempty)
leaf. Key order throughout is on *comparison keys* `kf key` (for SMT,
`keyForComparison` prehashes with SHA-256), so sortedness is sortedness of the
hashed key space — exactly how a real SMT/JMT arranges its leaves.
-/
import Ics23.SmtTree

namespace Ics23

/-! ## Order infrastructure: Option-valued extreme keys, comparison-key sortedness -/

/-- The rightmost (nonempty) leaf's key, if the tree has any leaf. -/
def maxKeyS : SMTree → Option Bytes
  | .empty => none
  | .leaf _ k _ => some k
  | .node _ _ l r =>
    match maxKeyS r with
    | some k => some k
    | none => maxKeyS l

/-- The leftmost (nonempty) leaf's key, if the tree has any leaf. -/
def minKeyS : SMTree → Option Bytes
  | .empty => none
  | .leaf _ k _ => some k
  | .node _ _ l r =>
    match minKeyS l with
    | some k => some k
    | none => minKeyS r

/-- Sorted w.r.t. comparison keys `kf` (for SMT, the prehashed key — see
`keyForComparison`): every member of a left subtree is strictly below every
member of the sibling right subtree. -/
def SortedTreeS (kf : Bytes → Bytes) : SMTree → Prop
  | .empty => True
  | .leaf _ _ _ => True
  | .node _ _ l r => SortedTreeS kf l ∧ SortedTreeS kf r ∧
      ∀ k₁ v₁ k₂ v₂, TreeMemberS k₁ v₁ l → TreeMemberS k₂ v₂ r →
        bytesLt (kf k₁) (kf k₂) = true

/-- `N` is a subtree of `t`. -/
def IsSubtreeS : SMTree → SMTree → Prop
  | N, .empty => N = .empty
  | N, .leaf op k v => N = .leaf op k v
  | N, .node ih pre l r => N = .node ih pre l r ∨ IsSubtreeS N l ∨ IsSubtreeS N r

/-- The max key, when present, is itself a member. -/
theorem maxKeyS_mem : ∀ (t : SMTree) (mk : Bytes), maxKeyS t = some mk →
    ∃ v, TreeMemberS mk v t := by
  intro t
  induction t with
  | empty => intro mk h; simp [maxKeyS] at h
  | leaf op k v =>
    intro mk h
    simp only [maxKeyS, Option.some.injEq] at h
    exact ⟨v, h, rfl⟩
  | node ih pre l r ihl ihr =>
    intro mk h
    unfold maxKeyS at h
    cases hr : maxKeyS r with
    | some k =>
      rw [hr] at h; simp only [Option.some.injEq] at h; subst h
      obtain ⟨v, hv⟩ := ihr k hr
      exact ⟨v, Or.inr hv⟩
    | none =>
      rw [hr] at h
      obtain ⟨v, hv⟩ := ihl mk h
      exact ⟨v, Or.inl hv⟩

/-- The min key, when present, is itself a member. -/
theorem minKeyS_mem : ∀ (t : SMTree) (mk : Bytes), minKeyS t = some mk →
    ∃ v, TreeMemberS mk v t := by
  intro t
  induction t with
  | empty => intro mk h; simp [minKeyS] at h
  | leaf op k v =>
    intro mk h
    simp only [minKeyS, Option.some.injEq] at h
    exact ⟨v, h, rfl⟩
  | node ih pre l r ihl ihr =>
    intro mk h
    unfold minKeyS at h
    cases hl : minKeyS l with
    | some k =>
      rw [hl] at h; simp only [Option.some.injEq] at h; subst h
      obtain ⟨v, hv⟩ := ihl k hl
      exact ⟨v, Or.inl hv⟩
    | none =>
      rw [hl] at h
      obtain ⟨v, hv⟩ := ihr mk h
      exact ⟨v, Or.inr hv⟩

/-- A tree with no max key has no members. -/
theorem maxKeyS_none_no_member : ∀ (t : SMTree), maxKeyS t = none →
    ∀ k v, ¬ TreeMemberS k v t := by
  intro t
  induction t with
  | empty => intro _ k v hm; exact hm
  | leaf op k' v' => intro h; simp [maxKeyS] at h
  | node ih pre l r ihl ihr =>
    intro h k v hm
    unfold maxKeyS at h
    cases hr : maxKeyS r with
    | some k' => rw [hr] at h; simp at h
    | none =>
      rw [hr] at h
      rcases hm with hml | hmr
      · exact ihl h k v hml
      · exact ihr hr k v hmr

/-- A tree with no min key has no members. -/
theorem minKeyS_none_no_member : ∀ (t : SMTree), minKeyS t = none →
    ∀ k v, ¬ TreeMemberS k v t := by
  intro t
  induction t with
  | empty => intro _ k v hm; exact hm
  | leaf op k' v' => intro h; simp [minKeyS] at h
  | node ih pre l r ihl ihr =>
    intro h k v hm
    unfold minKeyS at h
    cases hl : minKeyS l with
    | some k' => rw [hl] at h; simp at h
    | none =>
      rw [hl] at h
      rcases hm with hml | hmr
      · exact ihl hl k v hml
      · exact ihr h k v hmr

/-- A member's comparison key is `≤` the tree's max comparison key. -/
theorem member_le_maxKeyS (kf : Bytes → Bytes) : ∀ (t : SMTree),
    SortedTreeS kf t → ∀ key value, TreeMemberS key value t →
    ∃ mk, maxKeyS t = some mk ∧ (kf key = kf mk ∨ bytesLt (kf key) (kf mk) = true) := by
  intro t
  induction t with
  | empty => intro _ key value hm; exact absurd hm (by simp [TreeMemberS])
  | leaf op k v =>
    intro _ key value hm
    obtain ⟨hk, _⟩ := hm
    exact ⟨k, rfl, Or.inl (by rw [hk])⟩
  | node ih pre l r ihl ihr =>
    intro hs key value hm
    obtain ⟨hsl, hsr, hlr⟩ := hs
    rcases hm with hml | hmr
    · obtain ⟨mkl, hmkl, hle⟩ := ihl hsl key value hml
      cases hr : maxKeyS r with
      | none =>
        refine ⟨mkl, ?_, hle⟩
        unfold maxKeyS; rw [hr]; exact hmkl
      | some mkr =>
        refine ⟨mkr, by unfold maxKeyS; rw [hr], ?_⟩
        obtain ⟨vl, hvl⟩ := maxKeyS_mem l mkl hmkl
        obtain ⟨vr, hvr⟩ := maxKeyS_mem r mkr hr
        have hlt : bytesLt (kf mkl) (kf mkr) = true := hlr mkl vl mkr vr hvl hvr
        rcases hle with h | h
        · exact Or.inr (h ▸ hlt)
        · exact Or.inr (bytesLt_trans _ _ _ h hlt)
    · obtain ⟨mkr, hmkr, hle⟩ := ihr hsr key value hmr
      exact ⟨mkr, by unfold maxKeyS; rw [hmkr], hle⟩

/-- The tree's min comparison key is `≤` a member's comparison key. -/
theorem minKeyS_le_member (kf : Bytes → Bytes) : ∀ (t : SMTree),
    SortedTreeS kf t → ∀ key value, TreeMemberS key value t →
    ∃ mk, minKeyS t = some mk ∧ (kf mk = kf key ∨ bytesLt (kf mk) (kf key) = true) := by
  intro t
  induction t with
  | empty => intro _ key value hm; exact absurd hm (by simp [TreeMemberS])
  | leaf op k v =>
    intro _ key value hm
    obtain ⟨hk, _⟩ := hm
    exact ⟨k, rfl, Or.inl (by rw [hk])⟩
  | node ih pre l r ihl ihr =>
    intro hs key value hm
    obtain ⟨hsl, hsr, hlr⟩ := hs
    rcases hm with hml | hmr
    · obtain ⟨mkl, hmkl, hle⟩ := ihl hsl key value hml
      exact ⟨mkl, by unfold minKeyS; rw [hmkl], hle⟩
    · obtain ⟨mkr, hmkr, hle⟩ := ihr hsr key value hmr
      cases hl : minKeyS l with
      | none =>
        refine ⟨mkr, ?_, hle⟩
        unfold minKeyS; rw [hl]; exact hmkr
      | some mkl =>
        refine ⟨mkl, by unfold minKeyS; rw [hl], ?_⟩
        obtain ⟨vl, hvl⟩ := minKeyS_mem l mkl hl
        obtain ⟨vr, hvr⟩ := minKeyS_mem r mkr hmkr
        have hlt : bytesLt (kf mkl) (kf mkr) = true := hlr mkl vl mkr vr hvl hvr
        rcases hle with h | h
        · exact Or.inr (h ▸ hlt)
        · exact Or.inr (bytesLt_trans _ _ _ hlt h)

/-- A member of a subtree is a member of the whole tree. -/
theorem subtree_memberS : ∀ (t N : SMTree), IsSubtreeS N t →
    ∀ k v, TreeMemberS k v N → TreeMemberS k v t := by
  intro t
  induction t with
  | empty => intro N hsub k v hm; simp only [IsSubtreeS] at hsub; subst hsub; exact hm
  | leaf op k' v' => intro N hsub k v hm; simp only [IsSubtreeS] at hsub; subst hsub; exact hm
  | node ih pre l r ihl ihr =>
    intro N hsub k v hm
    simp only [IsSubtreeS] at hsub
    rcases hsub with hN | hl | hr
    · subst hN; exact hm
    · exact Or.inl (ihl N hl k v hm)
    · exact Or.inr (ihr N hr k v hm)

/-- A subtree of a sorted tree is sorted. -/
theorem subtree_sortedS (kf : Bytes → Bytes) : ∀ (t : SMTree), SortedTreeS kf t →
    ∀ N, IsSubtreeS N t → SortedTreeS kf N := by
  intro t
  induction t with
  | empty => intro hs N hsub; simp only [IsSubtreeS] at hsub; subst hsub; exact hs
  | leaf op k v => intro hs N hsub; simp only [IsSubtreeS] at hsub; subst hsub; exact hs
  | node ih pre l r ihl ihr =>
    intro hs N hsub
    obtain ⟨hsl, hsr, hlr⟩ := hs
    simp only [IsSubtreeS] at hsub
    rcases hsub with hN | hl | hr
    · subst hN; exact ⟨hsl, hsr, hlr⟩
    · exact ihl hsl N hl
    · exact ihr hsr N hr

/-- **Subtree contiguity (members form).** A member of a sorted tree whose
comparison key is bracketed by two members of a subtree `N` is itself a member
of `N`. Sorted trees occupy contiguous comparison-key ranges. -/
theorem member_bracketed_in_subtreeS (kf : Bytes → Bytes) : ∀ (t : SMTree),
    SortedTreeS kf t → ∀ N, IsSubtreeS N t →
    ∀ key value a va b vb,
      TreeMemberS key value t →
      TreeMemberS a va N → TreeMemberS b vb N →
      (kf a = kf key ∨ bytesLt (kf a) (kf key) = true) →
      (kf key = kf b ∨ bytesLt (kf key) (kf b) = true) →
      TreeMemberS key value N := by
  intro t
  induction t with
  | empty => intro _ N _ key value a va b vb hm; exact absurd hm (by simp [TreeMemberS])
  | leaf op k v =>
    intro _ N hsub key value a va b vb hm _ _ _ _
    simp only [IsSubtreeS] at hsub; subst hsub; exact hm
  | node ih pre l r ihl ihr =>
    intro hs N hsub key value a va b vb hm ha hb hlo hhi
    obtain ⟨hsl, hsr, hlr⟩ := hs
    simp only [IsSubtreeS] at hsub
    rcases hsub with hN | hsubl | hsubr
    · subst hN; exact hm
    · -- N ⊆ l: a member of t in r would violate key ≤ kf b with b ∈ l
      have hbl : TreeMemberS b vb l := subtree_memberS l N hsubl b vb hb
      rcases hm with hml | hmr
      · exact ihl hsl N hsubl key value a va b vb hml ha hb hlo hhi
      · exact absurd (hlr b vb key value hbl hmr) (fun h => ble_not_gt _ _ hhi h)
    · -- N ⊆ r: a member of t in l would violate kf a ≤ key with a ∈ r
      have har : TreeMemberS a va r := subtree_memberS r N hsubr a va ha
      rcases hm with hml | hmr
      · exact absurd (hlr key value a va hml har) (fun h => ble_not_gt _ _ hlo h)
      · exact ihr hsr N hsubr key value a va b vb hmr ha hb hlo hhi

/-- **BST gap (SMT).** For a node-subtree `N` of a sorted tree, no member of the
tree has a comparison key strictly between `N`'s left-max and right-min. -/
theorem node_gap_no_memberS (kf : Bytes → Bytes) (t : SMTree) (hs : SortedTreeS kf t)
    (ih : HashOp) (pre : Bytes) (l r : SMTree)
    (hsub : IsSubtreeS (.node ih pre l r) t)
    (key value : Bytes) (hm : TreeMemberS key value t)
    (kl kr : Bytes) (hkl : maxKeyS l = some kl) (hkr : minKeyS r = some kr)
    (h1 : bytesLt (kf kl) (kf key) = true) (h2 : bytesLt (kf key) (kf kr) = true) : False := by
  obtain ⟨vl, hvl⟩ := maxKeyS_mem l kl hkl
  obtain ⟨vr, hvr⟩ := minKeyS_mem r kr hkr
  have hmN : TreeMemberS key value (.node ih pre l r) :=
    member_bracketed_in_subtreeS kf t hs _ hsub key value kl vl kr vr hm
      (Or.inl hvl) (Or.inr hvr) (Or.inr h1) (Or.inr h2)
  have hsN := subtree_sortedS kf t hs _ hsub
  obtain ⟨hsl, hsr, _⟩ := hsN
  rcases hmN with hml | hmr
  · -- key ∈ l: key ≤ maxKeyS l = kl, but kl < key
    obtain ⟨mk, hmk, hle⟩ := member_le_maxKeyS kf l hsl key value hml
    rw [hkl] at hmk
    simp only [Option.some.injEq] at hmk; subst hmk
    exact ble_not_gt _ _ hle h1
  · -- key ∈ r: kr = minKeyS r ≤ key, but key < kr
    obtain ⟨mk, hmk, hle⟩ := minKeyS_le_member kf r hsr key value hmr
    rw [hkr] at hmk
    simp only [Option.some.injEq] at hmk; subst hmk
    exact ble_not_gt _ _ hle h2

/-! ## SMT padding bridges

The placeholder (`*_branches_are_empty`) arm is reachable for the SMT shape
(`emptyChild.length = childSize`), so instead of excluding it
(`ensureRightMost_suffix_nil`) we extract what it pins down: the sibling slot's
bytes are exactly `emptyChild`. -/

/-- An `ensure_right_most` step (binary spec) either has an empty suffix — a
genuine right step — or its suffix is exactly `emptyChild` (a left step whose
right sibling is the empty placeholder). -/
theorem ensureRightMost_step_smt (isp : InnerSpec) (cs : Nat)
    (hco : isp.childOrder = [0, 1])
    (hcsspec : isp.childSize.toNat = cs)
    (path : List InnerOp) (h : ensureRightMost isp path = true) :
    ∀ op, op ∈ path → op.suffix = [] ∨ op.suffix = isp.emptyChild := by
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
    refine Or.inl ?_
    unfold hasPadding at hp
    simp only [Bool.and_eq_true, decide_eq_true_eq] at hp
    have hz : (op.suffix.length : Int) = 0 := by have := hp.2; simpa using this
    have : op.suffix.length = 0 := by exact_mod_cast hz
    exact List.length_eq_zero_iff.mp this
  · -- placeholder: the (single) right-sibling block is exactly emptyChild
    refine Or.inr ?_
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
          exact h0
        · rw [if_pos hsl] at hr; simp at hr

/-- An `ensure_left_most` step (binary spec) either has a full-`cs` suffix — a
genuine left step — or it is a right step whose left sibling is the empty
placeholder: empty suffix, and the prefix's trailing `cs` bytes are exactly
`emptyChild`. -/
theorem ensureLeftMost_step_smt (isp : InnerSpec) (cs : Nat)
    (hco : isp.childOrder = [0, 1])
    (hcsspec : isp.childSize.toNat = cs)
    (path : List InnerOp) (h : ensureLeftMost isp path = true) :
    ∀ op, op ∈ path → op.suffix.length = cs ∨
      (op.suffix = [] ∧ cs ≤ op.prefixBytes.length ∧
       op.prefixBytes.drop (op.prefixBytes.length - cs) = isp.emptyChild) := by
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
  · -- genuine left-pad: suffix length equals pad.suffix = childSize
    refine Or.inl ?_
    unfold hasPadding at hp
    simp only [Bool.and_eq_true, decide_eq_true_eq] at hp
    have hz : (op.suffix.length : Int) = isp.childSize * ((2 : Int) - 1 - 0) := hp.2
    have he1 : ((2 : Int) - 1 - 0) = 1 := by decide
    rw [he1, Int.mul_one] at hz
    have : op.suffix.length = isp.childSize.toNat := by omega
    rw [this, hcsspec]
  · -- placeholder: a right step over an empty left sibling
    refine Or.inr ?_
    unfold leftBranchesAreEmpty at hl
    cases hofp : orderFromPadding isp op with
    | none => rw [hofp] at hl; simp at hl
    | some idx =>
      simp only [hofp] at hl
      by_cases hlb : idx = 0
      · rw [hlb] at hl; simp at hl
      · rw [if_neg hlb] at hl
        -- binary: idx = 1, and its padding forces an empty suffix
        obtain ⟨hilt, pad, hpad, hhp⟩ := orderFromPadding_mem isp op idx hofp
        rw [hco] at hilt
        simp only [List.length_cons, List.length_nil] at hilt
        have hi1 : idx = 1 := by omega
        subst hi1
        have hpr1 : getPadding isp 1
            = some { minPrefix := (1 : Int) * isp.childSize + isp.minPrefixLength,
                     maxPrefix := (1 : Int) * isp.childSize + isp.maxPrefixLength,
                     suffix := isp.childSize * ((2 : Int) - 1 - 1) } := by
          unfold getPadding; rw [hco]; rfl
        rw [hpr1, Option.some.injEq] at hpad; subst hpad
        unfold hasPadding at hhp
        simp only [Bool.and_eq_true, decide_eq_true_eq] at hhp
        have hsuf0 : op.suffix = [] := by
          have hz : (op.suffix.length : Int) = isp.childSize * ((2 : Int) - 1 - 1) := hhp.2
          have he0 : ((2 : Int) - 1 - 1) = 0 := by decide
          rw [he0, Int.mul_zero] at hz
          exact List.length_eq_zero_iff.mp (by exact_mod_cast hz)
        by_cases hpl : op.prefixBytes.length < 1 * isp.childSize.toNat
        · rw [if_pos hpl] at hl; simp at hl
        · rw [if_neg hpl] at hl
          rw [Nat.one_mul] at hpl
          have h0 : byteRange op.prefixBytes
              ((op.prefixBytes.length - 1 * isp.childSize.toNat) + 0 * isp.childSize.toNat)
              isp.childSize.toNat == some isp.emptyChild := by
            have := List.all_eq_true.mp hl 0 (by simp only [List.mem_range]; omega)
            simpa using this
          simp only [Nat.zero_mul, Nat.one_mul, Nat.add_zero, byteRange] at h0
          rw [if_pos (by omega)] at h0
          simp only [beq_iff_eq, Option.some.injEq] at h0
          rw [List.take_of_length_le (by rw [List.length_drop]; omega)] at h0
          rw [← hcsspec]
          exact ⟨hsuf0, by omega, h0⟩

/-! ## Navigation: extreme-key paths in the SMT model -/

/-- A node split with a full `cs`-length suffix resolves completely: the prefix
is the node prefix, the child is the left subtree hash, and the suffix is the
right subtree hash. (Strengthens `split_left` to also name the suffix, which for
SMT pins the right sibling.) -/
theorem split_all (pre lh rh topPre m topSuf : Bytes) (cs : Nat)
    (hN : topPre ++ m ++ topSuf = pre ++ lh ++ rh)
    (hlh : lh.length = cs) (hrh : rh.length = cs)
    (hm : m.length = cs) (hsuf : topSuf.length = cs) :
    topPre = pre ∧ m = lh ∧ topSuf = rh := by
  have hlen : topPre.length = pre.length := by
    have h := congrArg List.length hN
    simp only [List.length_append] at h; omega
  simp only [List.append_assoc] at hN
  have hA := List.append_inj hN hlen
  have hB := List.append_inj hA.2 (by rw [hm, hlh])
  exact ⟨hA.1, hB.1, hB.2⟩

/-- A node split with an empty suffix resolves completely: the prefix is the
node prefix followed by the left subtree hash, and the child is the right
subtree hash. (Strengthens `split_right` to also name the prefix, which for SMT
pins the left sibling.) -/
theorem split_all_right (pre lh rh topPre m : Bytes) (cs : Nat)
    (hN : topPre ++ m ++ [] = pre ++ lh ++ rh)
    (hlh : lh.length = cs) (hrh : rh.length = cs) (hm : m.length = cs) :
    topPre = pre ++ lh ∧ m = rh := by
  rw [List.append_nil] at hN
  have hlen : topPre.length = (pre ++ lh).length := by
    have h := congrArg List.length hN
    simp only [List.length_append] at h ⊢; omega
  exact List.append_inj hN hlen

/-- An `ensure_right_most`-shaped path (each step: empty suffix, or suffix
`= emptyChild`) reaches the rightmost *nonempty* leaf: the proof's key is
`maxKeyS t`. A placeholder step pins the skipped right sibling to the empty
subtree (`rootHashS_ec_empty`), so the maximum is preserved. -/
theorem reaches_maxS (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat) (ec : Bytes)
    (hH : FixedHash H cs)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hecl : ec.length = cs)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : SMTree) (key value lh : Bytes) (path : List InnerOp) (root : Bytes),
      WFTreeS s b t →
      applyLeaf H s.leafSpec key value = some lh →
      (∀ op ∈ path, ensureInner op s = true ∧ (op.suffix = [] ∨ op.suffix = ec)) →
      rootHashS H ec t = some root →
      applyPath H s.innerSpec lh path = some root →
      (maxKeyS t = some key ∧ TreeMemberS key value t) ∨ HashCollision H := by
  intro t
  induction t with
  | empty =>
    intro key value lh path root _ hlh hpath hrh hap
    exfalso
    simp only [rootHashS, Option.some.injEq] at hrh
    exact applyPath_ne_emptyChild H s ec hihash.symm hef key value lh root path
      hlh (fun o ho => (hpath o ho).1) hap hrh.symm
  | leaf top tk tv =>
    intro key value lh path root hwf hlh hpath hrh hap
    simp only [WFTreeS] at hwf; subst hwf
    rw [rootHashS] at hrh
    cases path with
    | nil =>
      simp only [applyPath, Option.some.injEq] at hap; subst hap
      rcases hLInj key value tk tv _ hlh hrh with ⟨hk, hv⟩ | hc
      · exact Or.inl ⟨by simp only [maxKeyS]; rw [hk], hk.symm, hv.symm⟩
      · exact Or.inr hc
    | cons op rest =>
      have hii : IsInnerImage H s root :=
        applyPath_result_isInnerImage H s (op :: rest) lh root (by simp)
          (fun o ho => (hpath o ho).1) hap
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
        have hlhLlen := rootHashS_len H ec cs hH hecl l lhL hl
        have hrhRlen := rootHashS_len H ec cs hH hecl r rhR hr
        have hq : ∀ op ∈ q, ensureInner op s = true ∧ (op.suffix = [] ∨ op.suffix = ec) :=
          fun o ho => hpath o (List.mem_append.mpr (Or.inl ho))
        rcases (hpath topOp (by simp)).2 with hsnil | hsec
        · -- right step: recurse into the right subtree
          rw [hsnil] at hpe
          obtain ⟨_, hmr⟩ := split_all_right pre lhL rhR topOp.prefixBytes m cs
            hpe hlhLlen hrhRlen hmlen
          rw [hmr] at hpm
          rcases ihr key value lh q rhR hwfr hlh hq hr hpm with ⟨hk, hmem⟩ | hcol
          · refine Or.inl ⟨?_, Or.inr hmem⟩
            show (match maxKeyS r with | some k => some k | none => maxKeyS l) = some key
            rw [hk]
          · exact Or.inr hcol
        · -- placeholder step: the right sibling is the empty subtree
          obtain ⟨_, hml, hsr⟩ := split_all pre lhL rhR topOp.prefixBytes m topOp.suffix cs
            hpe hlhLlen hrhRlen hmlen (by rw [hsec]; exact hecl)
          have hrEC : rootHashS H ec r = some ec := by rw [hr, ← hsr, hsec]
          have hrempty : r = .empty :=
            rootHashS_ec_empty H s b ec hihash.symm hef r hwfr hrEC
          rw [hml] at hpm
          rcases ihl key value lh q lhL hwfl hlh hq hl hpm with ⟨hk, hmem⟩ | hcol
          · refine Or.inl ⟨?_, Or.inl hmem⟩
            show (match maxKeyS r with | some k => some k | none => maxKeyS l) = some key
            rw [hrempty]
            show (match (none : Option Bytes) with
              | some k => some k | none => maxKeyS l) = some key
            exact hk
          · exact Or.inr hcol
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpe (htopimg.trans hrh.symm))

/-- An `ensure_left_most`-shaped path (each step: full-`cs` suffix, or empty
suffix with trailing-`emptyChild` prefix) reaches the leftmost *nonempty* leaf:
the proof's key is `minKeyS t`. Mirror of `reaches_maxS`. -/
theorem reaches_minS (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat) (ec : Bytes)
    (hH : FixedHash H cs)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hecl : ec.length = cs)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : SMTree) (key value lh : Bytes) (path : List InnerOp) (root : Bytes),
      WFTreeS s b t →
      applyLeaf H s.leafSpec key value = some lh →
      (∀ op ∈ path, ensureInner op s = true ∧
        (op.suffix.length = cs ∨
         (op.suffix = [] ∧ cs ≤ op.prefixBytes.length ∧
          op.prefixBytes.drop (op.prefixBytes.length - cs) = ec))) →
      rootHashS H ec t = some root →
      applyPath H s.innerSpec lh path = some root →
      (minKeyS t = some key ∧ TreeMemberS key value t) ∨ HashCollision H := by
  intro t
  induction t with
  | empty =>
    intro key value lh path root _ hlh hpath hrh hap
    exfalso
    simp only [rootHashS, Option.some.injEq] at hrh
    exact applyPath_ne_emptyChild H s ec hihash.symm hef key value lh root path
      hlh (fun o ho => (hpath o ho).1) hap hrh.symm
  | leaf top tk tv =>
    intro key value lh path root hwf hlh hpath hrh hap
    simp only [WFTreeS] at hwf; subst hwf
    rw [rootHashS] at hrh
    cases path with
    | nil =>
      simp only [applyPath, Option.some.injEq] at hap; subst hap
      rcases hLInj key value tk tv _ hlh hrh with ⟨hk, hv⟩ | hc
      · exact Or.inl ⟨by simp only [minKeyS]; rw [hk], hk.symm, hv.symm⟩
      · exact Or.inr hc
    | cons op rest =>
      have hii : IsInnerImage H s root :=
        applyPath_result_isInnerImage H s (op :: rest) lh root (by simp)
          (fun o ho => (hpath o ho).1) hap
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
        have hlhLlen := rootHashS_len H ec cs hH hecl l lhL hl
        have hrhRlen := rootHashS_len H ec cs hH hecl r rhR hr
        have hq : ∀ op ∈ q, ensureInner op s = true ∧
            (op.suffix.length = cs ∨
             (op.suffix = [] ∧ cs ≤ op.prefixBytes.length ∧
              op.prefixBytes.drop (op.prefixBytes.length - cs) = ec)) :=
          fun o ho => hpath o (List.mem_append.mpr (Or.inl ho))
        rcases (hpath topOp (by simp)).2 with hscs | ⟨hsnil, hple, hpec⟩
        · -- left step: recurse into the left subtree
          obtain ⟨_, hml, _⟩ := split_all pre lhL rhR topOp.prefixBytes m topOp.suffix cs
            hpe hlhLlen hrhRlen hmlen hscs
          rw [hml] at hpm
          rcases ihl key value lh q lhL hwfl hlh hq hl hpm with ⟨hk, hmem⟩ | hcol
          · refine Or.inl ⟨?_, Or.inl hmem⟩
            show (match minKeyS l with | some k => some k | none => minKeyS r) = some key
            rw [hk]
          · exact Or.inr hcol
        · -- placeholder step: the left sibling is the empty subtree
          rw [hsnil] at hpe
          obtain ⟨hpfx, hmr⟩ := split_all_right pre lhL rhR topOp.prefixBytes m cs
            hpe hlhLlen hrhRlen hmlen
          have hlEC : lhL = ec := by
            rw [hpfx] at hpec
            have hplen : (pre ++ lhL).length - cs = pre.length := by
              simp only [List.length_append]; omega
            rw [hplen, List.drop_left] at hpec
            exact hpec
          have hrempty : l = .empty :=
            rootHashS_ec_empty H s b ec hihash.symm hef l hwfl (by rw [hl, hlEC])
          rw [hmr] at hpm
          rcases ihr key value lh q rhR hwfr hlh hq hr hpm with ⟨hk, hmem⟩ | hcol
          · refine Or.inl ⟨?_, Or.inr hmem⟩
            show (match minKeyS l with | some k => some k | none => minKeyS r) = some key
            rw [hrempty]
            show (match (none : Option Bytes) with
              | some k => some k | none => minKeyS r) = some key
            exact hk
          · exact Or.inr hcol
      · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpe (htopimg.trans hrh.symm))

/-! ## Neighbor divergence and Theorem B for the SMT model -/

/-- **Neighbor walk → divergence node (SMT).** Two verifying existence proofs
whose reversed paths share a root-side prefix and then diverge as a left-step
(with SMT right-most / left-most remainders) navigate to a common node `N` of
the honest SMT: the left proof reaches the max key of `N`'s left subtree, the
right proof the min key of `N`'s right subtree — up to a hash collision. -/
theorem neighbor_divergence_smt (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat) (ec : Bytes)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hihash : s.innerSpec.hash = s.leafSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hecl : ec.length = cs)
    (hecspec : s.innerSpec.emptyChild = ec)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H) :
    ∀ (t : SMTree) (lhLeaf rhLeaf root : Bytes)
      (leftKey leftVal rightKey rightVal : Bytes)
      (pathL pathR : List InnerOp)
      (topLeft topRight : InnerOp) (restL restR : List InnerOp),
      WFTreeS s b t →
      applyLeaf H s.leafSpec leftKey leftVal = some lhLeaf →
      applyLeaf H s.leafSpec rightKey rightVal = some rhLeaf →
      (∀ op ∈ pathL, ensureInner op s = true) →
      (∀ op ∈ pathR, ensureInner op s = true) →
      rootHashS H ec t = some root →
      applyPath H s.innerSpec lhLeaf pathL = some root →
      applyPath H s.innerSpec rhLeaf pathR = some root →
      dropCommonPrefix pathL.reverse pathR.reverse = (some topLeft, restL, some topRight, restR) →
      isLeftStep s.innerSpec topLeft topRight = true →
      ensureRightMost s.innerSpec restL.reverse = true →
      ensureLeftMost s.innerSpec restR.reverse = true →
      (∃ ihN preN lN rN, IsSubtreeS (.node ihN preN lN rN) t ∧
        (maxKeyS lN = some leftKey ∨ HashCollision H) ∧
        (minKeyS rN = some rightKey ∨ HashCollision H))
      ∨ HashCollision H := by
    intro t
    induction t with
    | empty =>
      intro lhLeaf rhLeaf root leftKey leftVal rightKey rightVal pathL pathR
        topLeft topRight restL restR
        _ hlhL _ hpathL _ hrh hapL _ _ _ _ _
      exfalso
      simp only [rootHashS, Option.some.injEq] at hrh
      exact applyPath_ne_emptyChild H s ec hihash.symm hef leftKey leftVal lhLeaf root pathL
        hlhL hpathL hapL hrh.symm
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
        simp only [WFTreeS] at hwf; subst hwf
        rw [rootHashS] at hrh
        have hii : IsInnerImage H s root :=
          applyPath_result_isInnerImage H s (op :: rest) lhLeaf root (by simp) hpathL hapL
        exact Or.inr (leafHash_innerImage_collision H s s.leafSpec tk tv root b
          hihash.symm hlpre hmin (ensureLeaf_self s.leafSpec) hrh hii)
    | node ih pre l r ihl ihr =>
      intro lhLeaf rhLeaf root leftKey leftVal rightKey rightVal pathL pathR
        topLeft topRight restL restR
        hwf hlhL hlhR hpathL hpathR hrh hapL hapR hdcp hls hrm hlm
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
      have hlhLlen := rootHashS_len H ec cs hH hecl l lhL hl
      have hrhRlen := rootHashS_len H ec cs hH hecl r rhR hr
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
                exact Or.inl ⟨ihN, preN, lN, rN,
                  by simp only [IsSubtreeS]; exact Or.inr (Or.inl hsub), h1, h2⟩
              · exact Or.inr hc
            · rw [hmrL] at hpmL
              rw [← hmEq, hmrL] at hpmR
              rcases ihr lhLeaf rhLeaf rhR leftKey leftVal rightKey rightVal qL qR
                topLeft topRight restL restR hwfr hlhL hlhR
                (fun o ho => hpathL o (List.mem_append.mpr (Or.inl ho)))
                (fun o ho => hpathR o (List.mem_append.mpr (Or.inl ho)))
                hr hpmL hpmR hdcp hls hrm hlm with hex | hc
              · obtain ⟨ihN, preN, lN, rN, hsub, h1, h2⟩ := hex
                exact Or.inl ⟨ihN, preN, lN, rN,
                  by simp only [IsSubtreeS]; exact Or.inr (Or.inr hsub), h1, h2⟩
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
            have hqLspec : ∀ op ∈ qL, ensureInner op s = true ∧
                (op.suffix = [] ∨ op.suffix = ec) := fun op hop =>
              ⟨hpathL op (List.mem_append.mpr (Or.inl hop)),
               hecspec ▸ ensureRightMost_step_smt s.innerSpec cs hco hcsspec qL hrm op hop⟩
            have hqRspec : ∀ op ∈ qR, ensureInner op s = true ∧
                (op.suffix.length = cs ∨
                 (op.suffix = [] ∧ cs ≤ op.prefixBytes.length ∧
                  op.prefixBytes.drop (op.prefixBytes.length - cs) = ec)) := fun op hop =>
              ⟨hpathR op (List.mem_append.mpr (Or.inl hop)),
               hecspec ▸ ensureLeftMost_step_smt s.innerSpec cs hco hcsspec qR hlm op hop⟩
            refine Or.inl ⟨s.innerSpec.hash, pre, l, r, Or.inl rfl, ?_, ?_⟩
            · rcases reaches_maxS H s b cs ec hH hihash hlpre hmin hecl hef hLInj
                l leftKey leftVal lhLeaf qL lhL hwfl hlhL hqLspec hl hpmL with ⟨hk, _⟩ | hc
              · exact Or.inl hk
              · exact Or.inr hc
            · rcases reaches_minS H s b cs ec hH hihash hlpre hmin hecl hef hLInj
                r rightKey rightVal rhLeaf qR rhR hwfr hlhR hqRspec hr hpmR with ⟨hk, _⟩ | hc
              · exact Or.inl hk
              · exact Or.inr hc
          · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpeR (htopimgR.trans hrh.symm))
        · exact Or.inr (hashCollision_of H s.innerSpec.hash _ _ hpeL (htopimgL.trans hrh.symm))

/-- **Theorem B (non-existence soundness), SMT tree model — two-sided.** For an
SMT sorted on comparison keys, a two-sided non-existence proof for `key` and an
existence proof for `key` cannot both verify without a hash collision. -/
theorem nonexistence_soundS (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat) (ec : Bytes)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hecl : ec.length = cs)
    (hecspec : s.innerSpec.emptyChild = ec)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : SMTree) (hwf : WFTreeS s b t) (hsort : SortedTreeS (keyForComparison H s) t)
    (root key value : Bytes) (hroot : rootHashS H ec t = some root)
    (nep : NonExistenceProof) (ep lp rp : ExistenceProof)
    (hnl : nep.left = some lp) (hnr : nep.right = some rp)
    (hepleaf : ep.leaf = s.leafSpec) (hlpleaf : lp.leaf = s.leafSpec) (hrpleaf : rp.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  have hbin : s.innerSpec.childOrder.length = 2 := by rw [hco]; rfl
  -- the absent key is a genuine member of the honest tree (or a collision)
  rcases membership_soundS H s b cs ec hH hcs hcsspec hlhash hlpre hmin hmm hbin hecl hef hLInj
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
    rcases neighbor_divergence_smt H s b cs ec hH hcs hcsspec hlhash.symm hlpre hmin hmm hco
      hecl hecspec hef hLInj
      t lhLeaf rhLeaf root lp.key lp.value rp.key rp.value lp.path rp.path
      topLeft topRight restL restR
      hwf hlapL hrapL hlinn hrinn hroot hlap hrap hdcp hls hrm hlm with hdiv | hc
    · obtain ⟨ihN, preN, lN, rN, hsub, hLeq, hReq⟩ := hdiv
      rcases hLeq with hLeq | hc
      · rcases hReq with hReq | hc
        · -- the absent key sits in the forbidden gap of the divergence node
          exact (node_gap_no_memberS (keyForComparison H s) t hsort ihN preN lN rN hsub
            key value hmem lp.key rp.key hLeq hReq hllt hrlt).elim
        · exact hc
      · exact hc
    · exact hc
  · exact hc

/-- **Theorem B, SMT tree model — left-only.** The lone left neighbor's
right-most path makes it the rightmost nonempty leaf (`reaches_maxS`), so
`key`'s comparison key exceeds the tree's maximum while a verifying existence
proof would bound it by that maximum. -/
theorem nonexistence_soundS_leftOnly (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat) (ec : Bytes)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hecl : ec.length = cs)
    (hecspec : s.innerSpec.emptyChild = ec)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : SMTree) (hwf : WFTreeS s b t) (hsort : SortedTreeS (keyForComparison H s) t)
    (root key value : Bytes) (hroot : rootHashS H ec t = some root)
    (nep : NonExistenceProof) (ep lp : ExistenceProof)
    (hnl : nep.left = some lp) (hnr : nep.right = none)
    (hepleaf : ep.leaf = s.leafSpec) (hlpleaf : lp.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  have hbin : s.innerSpec.childOrder.length = 2 := by rw [hco]; rfl
  rcases membership_soundS H s b cs ec hH hcs hcsspec hlhash hlpre hmin hmm hbin hecl hef hLInj
    t ep root key value hwf hepleaf hroot hex with hmem | hc
  · obtain ⟨hlver, hllt⟩ := verifyNonExistence_left H s root key nep lp hnl hne
    have hrm := verifyNonExistence_leftOnly H s root key nep lp hnl hnr hne
    obtain ⟨lhLeaf, hlapL, hlap, hlinn⟩ := verifyExistence_navigates H s lp root hlpleaf hlver
    rcases reaches_maxS H s b cs ec hH hlhash.symm hlpre hmin hecl hef hLInj
      t lp.key lp.value lhLeaf lp.path root hwf hlapL
      (fun op hop => ⟨hlinn op hop,
        hecspec ▸ ensureRightMost_step_smt s.innerSpec cs hco hcsspec lp.path hrm op hop⟩)
      hroot hlap with ⟨hk, _⟩ | hc
    · obtain ⟨mk, hmk, hle⟩ := member_le_maxKeyS (keyForComparison H s) t hsort key value hmem
      rw [hk] at hmk
      simp only [Option.some.injEq] at hmk
      subst hmk
      exact (ble_not_gt _ _ hle hllt).elim
    · exact hc
  · exact hc

/-- **Theorem B, SMT tree model — right-only.** Mirror of
`nonexistence_soundS_leftOnly` via `reaches_minS` / `minKeyS_le_member`. -/
theorem nonexistence_soundS_rightOnly (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat) (ec : Bytes)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hecl : ec.length = cs)
    (hecspec : s.innerSpec.emptyChild = ec)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : SMTree) (hwf : WFTreeS s b t) (hsort : SortedTreeS (keyForComparison H s) t)
    (root key value : Bytes) (hroot : rootHashS H ec t = some root)
    (nep : NonExistenceProof) (ep rp : ExistenceProof)
    (hnl : nep.left = none) (hnr : nep.right = some rp)
    (hepleaf : ep.leaf = s.leafSpec) (hrpleaf : rp.leaf = s.leafSpec)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  have hbin : s.innerSpec.childOrder.length = 2 := by rw [hco]; rfl
  rcases membership_soundS H s b cs ec hH hcs hcsspec hlhash hlpre hmin hmm hbin hecl hef hLInj
    t ep root key value hwf hepleaf hroot hex with hmem | hc
  · obtain ⟨hrver, hrlt⟩ := verifyNonExistence_right H s root key nep rp hnr hne
    have hlm := verifyNonExistence_rightOnly H s root key nep rp hnl hnr hne
    obtain ⟨rhLeaf, hrapL, hrap, hrinn⟩ := verifyExistence_navigates H s rp root hrpleaf hrver
    rcases reaches_minS H s b cs ec hH hlhash.symm hlpre hmin hecl hef hLInj
      t rp.key rp.value rhLeaf rp.path root hwf hrapL
      (fun op hop => ⟨hrinn op hop,
        hecspec ▸ ensureLeftMost_step_smt s.innerSpec cs hco hcsspec rp.path hlm op hop⟩)
      hroot hrap with ⟨hk, _⟩ | hc
    · obtain ⟨mk, hmk, hle⟩ := minKeyS_le_member (keyForComparison H s) t hsort key value hmem
      rw [hk] at hmk
      simp only [Option.some.injEq] at hmk
      subst hmk
      exact (ble_not_gt _ _ hle hrlt).elim
    · exact hc
  · exact hc

/-- **Theorem B, SMT tree model — total.** Every verifier-accepted proof shape
is covered: two-sided, left-only, right-only (a no-neighbor proof never
verifies). -/
theorem nonexistence_soundS_total (H : HashFn) (s : ProofSpec) (b : UInt8) (cs : Nat) (ec : Bytes)
    (hH : FixedHash H cs) (hcs : 0 < cs)
    (hcsspec : s.innerSpec.childSize.toNat = cs)
    (hlhash : s.leafSpec.hash = s.innerSpec.hash)
    (hlpre : s.leafSpec.prefixBytes = [b])
    (hmin : 1 ≤ s.innerSpec.minPrefixLength)
    (hmm : s.innerSpec.minPrefixLength = s.innerSpec.maxPrefixLength)
    (hco : s.innerSpec.childOrder = [0, 1])
    (hecl : ec.length = cs)
    (hecspec : s.innerSpec.emptyChild = ec)
    (hef : EmptyChildFree H s.innerSpec.hash ec)
    (hLInj : ∀ k₁ v₁ k₂ v₂ r, applyLeaf H s.leafSpec k₁ v₁ = some r →
      applyLeaf H s.leafSpec k₂ v₂ = some r → (k₁ = k₂ ∧ v₁ = v₂) ∨ HashCollision H)
    (t : SMTree) (hwf : WFTreeS s b t) (hsort : SortedTreeS (keyForComparison H s) t)
    (root key value : Bytes) (hroot : rootHashS H ec t = some root)
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
      exact nonexistence_soundS_rightOnly H s b cs ec hH hcs hcsspec hlhash hlpre hmin hmm
        hco hecl hecspec hef hLInj t hwf hsort root key value hroot nep ep rp hnl hnr hepleaf
        (hrleaf rp hnr) hne hex
  | some lp =>
    cases hnr : nep.right with
    | none =>
      exact nonexistence_soundS_leftOnly H s b cs ec hH hcs hcsspec hlhash hlpre hmin hmm
        hco hecl hecspec hef hLInj t hwf hsort root key value hroot nep ep lp hnl hnr hepleaf
        (hlleaf lp hnl) hne hex
    | some rp =>
      exact nonexistence_soundS H s b cs ec hH hcs hcsspec hlhash hlpre hmin hmm hco hecl
        hecspec hef hLInj t hwf hsort root key value hroot nep ep lp rp hnl hnr hepleaf
        (hlleaf lp hnl) (hrleaf rp hnr) hne hex

/-- **Theorem B for the SMT/JMT spec — total.** All structural side conditions
discharged by computation; joint leaf injectivity proved (`leafInj_smt`). The
remaining hypotheses are the sparse-merkle construction's two cryptographic
assumptions — fixed 32-byte digests (`FixedHash`) and no exhibited preimage of
the all-zero placeholder (`EmptyChildFree`) — plus the store invariant that the
honest tree is sorted on *prehashed* keys (`keyForComparison`, which for
`smt_spec` is SHA-256 of the raw key), exactly how a real SMT/JMT arranges its
leaves. -/
theorem nonexistence_sound_smt_total (H : HashFn)
    (hH : FixedHash H 32)
    (hef : EmptyChildFree H .sha256 (List.replicate 32 0))
    (t : SMTree) (hwf : WFTreeS smtSpec 0 t)
    (hsort : SortedTreeS (keyForComparison H smtSpec) t)
    (root key value : Bytes) (hroot : rootHashS H (List.replicate 32 0) t = some root)
    (nep : NonExistenceProof) (ep : ExistenceProof)
    (hepleaf : ep.leaf = smtSpec.leafSpec)
    (hlleaf : ∀ p, nep.left = some p → p.leaf = smtSpec.leafSpec)
    (hrleaf : ∀ p, nep.right = some p → p.leaf = smtSpec.leafSpec)
    (hne : verifyNonExistence H nep smtSpec root key = true)
    (hex : verifyExistence H ep smtSpec root key value = true) :
    HashCollision H :=
  nonexistence_soundS_total H smtSpec 0 32 (List.replicate 32 0) hH (by decide) (by decide)
    (by decide) (by decide) (by decide) (by decide) (by decide) (by decide) (by decide)
    hef (leafInj_smt H hH) t hwf hsort root key value hroot nep ep hepleaf hlleaf hrleaf hne hex

end Ics23
