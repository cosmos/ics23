/-
Non-existence soundness (Theorem B), staged.

States Theorem B precisely and proves supporting facts about the lexicographic
byte ordering `bytesLt` that the neighbor checks rely on. The main proof needs
the ordered-tree semantics an `InnerSpec` describes and is landed separately;
the statement is recorded here so the target is fixed (mirroring how Theorem A
was staged). Only `nonexistence_sound` uses `sorry`.
-/
import Ics23.NonExist
import Ics23.Soundness
import Ics23.Order

namespace Ics23

/-- `bytesLt` is irreflexive: no byte string is strictly less than itself. -/
theorem bytesLt_irrefl (a : Bytes) : bytesLt a a = false := by
  induction a with
  | nil => rfl
  | cons x xs ih =>
    show (if x < x then true else if x == x then bytesLt xs xs else false) = false
    rw [if_neg (UInt8.lt_irrefl x), if_pos (by simp)]
    exact ih

/-- A strictly-ordered key cannot also be equal: `bytesLt` excludes equality. -/
theorem bytesLt_ne (a b : Bytes) (h : bytesLt a b = true) : a ≠ b := by
  intro hab
  rw [hab, bytesLt_irrefl] at h
  exact Bool.noConfusion h

/-- `bytesLt` is transitive. -/
theorem bytesLt_trans : ∀ (a b c : Bytes),
    bytesLt a b = true → bytesLt b c = true → bytesLt a c = true
  | [], [], _, h, _ => by simp [bytesLt] at h
  | _ :: _, [], _, h, _ => by simp [bytesLt] at h
  | [], _ :: _, [], _, h2 => by simp [bytesLt] at h2
  | [], _ :: _, _ :: _, _, _ => rfl
  | x :: xs, y :: ys, [], _, h2 => by simp [bytesLt] at h2
  | x :: xs, y :: ys, z :: zs, h1, h2 => by
    simp only [bytesLt] at h1 h2 ⊢
    by_cases hxy : x < y
    · -- x < y
      by_cases hyz : y < z
      · simp [UInt8.lt_trans hxy hyz]
      · simp only [hyz, if_false] at h2
        by_cases hyz' : (y == z) = true
        · have hyzeq : y = z := eq_of_beq hyz'
          subst hyzeq; simp [hxy]
        · simp [hyz'] at h2
    · -- ¬ x < y
      simp only [hxy, if_false] at h1
      by_cases hxy' : (x == y) = true
      · have hxyeq : x = y := eq_of_beq hxy'
        subst hxyeq
        simp only [beq_self_eq_true, if_true] at h1
        by_cases hxz : x < z
        · simp [hxz]
        · simp only [hxz, if_false] at h2 ⊢
          by_cases hxz' : (x == z) = true
          · have hxzeq : x = z := eq_of_beq hxz'
            subst hxzeq
            simp only [beq_self_eq_true, if_true] at h2 ⊢
            exact bytesLt_trans xs ys zs h1 h2
          · simp [hxz'] at h2
      · simp [hxy'] at h1

/-- Left-neighbor extraction: the left existence proof verifies and its key
sorts strictly before `key`. -/
theorem verifyNonExistence_left (H : HashFn) (s : ProofSpec) (root key : Bytes)
    (nep : NonExistenceProof) (l : ExistenceProof) (hl : nep.left = some l)
    (h : verifyNonExistence H nep s root key = true) :
    verifyExistence H l s root l.key l.value = true ∧
    bytesLt (keyForComparison H s l.key) (keyForComparison H s key) = true := by
  unfold verifyNonExistence at h
  simp only [hl, Bool.and_eq_true] at h
  exact h.1.1

/-- Right-neighbor extraction: the right existence proof verifies and its key
sorts strictly after `key`. -/
theorem verifyNonExistence_right (H : HashFn) (s : ProofSpec) (root key : Bytes)
    (nep : NonExistenceProof) (r : ExistenceProof) (hr : nep.right = some r)
    (h : verifyNonExistence H nep s root key = true) :
    verifyExistence H r s root r.key r.value = true ∧
    bytesLt (keyForComparison H s key) (keyForComparison H s r.key) = true := by
  unfold verifyNonExistence at h
  simp only [hr, Bool.and_eq_true] at h
  exact h.1.2

/-- One-sided extraction (left-only): with no right neighbor, the verifier
requires the left neighbor's path to be right-most — the claim is that `key`
lies beyond the last leaf of the tree. -/
theorem verifyNonExistence_leftOnly (H : HashFn) (s : ProofSpec) (root key : Bytes)
    (nep : NonExistenceProof) (l : ExistenceProof)
    (hl : nep.left = some l) (hr : nep.right = none)
    (h : verifyNonExistence H nep s root key = true) :
    ensureRightMost s.innerSpec l.path = true := by
  unfold verifyNonExistence at h
  simp only [hl, hr, Bool.and_eq_true] at h
  exact h.2

/-- One-sided extraction (right-only): with no left neighbor, the verifier
requires the right neighbor's path to be left-most — the claim is that `key`
lies before the first leaf of the tree. -/
theorem verifyNonExistence_rightOnly (H : HashFn) (s : ProofSpec) (root key : Bytes)
    (nep : NonExistenceProof) (r : ExistenceProof)
    (hl : nep.left = none) (hr : nep.right = some r)
    (h : verifyNonExistence H nep s root key = true) :
    ensureLeftMost s.innerSpec r.path = true := by
  unfold verifyNonExistence at h
  simp only [hl, hr, Bool.and_eq_true] at h
  exact h.2

/-- A no-neighbor non-existence proof never verifies. -/
theorem verifyNonExistence_none (H : HashFn) (s : ProofSpec) (root key : Bytes)
    (nep : NonExistenceProof) (hl : nep.left = none) (hr : nep.right = none) :
    verifyNonExistence H nep s root key = false := by
  unfold verifyNonExistence
  simp [hl, hr]

/-- A two-sided non-existence proof brackets `key` strictly: its left neighbor
sorts before its right neighbor. Composes the extraction lemmas with
transitivity; a fully-proved consistency property of the verifier. -/
theorem verifyNonExistence_neighbors_ordered
    (H : HashFn) (s : ProofSpec) (root key : Bytes)
    (nep : NonExistenceProof) (l r : ExistenceProof)
    (hl : nep.left = some l) (hr : nep.right = some r)
    (h : verifyNonExistence H nep s root key = true) :
    bytesLt (keyForComparison H s l.key) (keyForComparison H s r.key) = true :=
  bytesLt_trans _ _ _
    (verifyNonExistence_left H s root key nep l hl h).2
    (verifyNonExistence_right H s root key nep r hr h).2

/-- **Theorem B (non-existence soundness), abstract-root statement.** A
non-existence proof for `key` and an existence proof for `key` cannot both verify
under the same spec and root without a hash collision.

This abstract-root form is **deliberately left as the one `sorry`**: as written it
is not provable, for exactly the reasons findings F3/F4 record — an opaque `root`
carries no tree structure, so nothing connects byte-order (which the neighbor
checks constrain) to *position* in the tree, and nothing forbids `key`'s leaf
sitting at an unrelated position. This is the non-existence analog of why
Theorem A's real result is the honest-root `membership_sound` rather than an
abstract-root binding.

The genuine result is proved in the honest-root tree model:
`Ics23.nonexistence_sound_tree` (and `nonexistence_sound_tree_tendermint`) in
`Tree`/`TreeNonExist`. There, `root = rootHash t` for a key-sorted honest tree
`t`, and the full chain is discharged with no axioms beyond `FixedHash` + joint
leaf injectivity: `membership_sound` makes the absent key a genuine member,
`ensureLeftNeighbor_spec` + `neighbor_divergence` pin the bracketing proofs to a
divergence node `N`, and `node_gap_no_member` shows no member can sit between
`maxKey N.left` and `minKey N.right`. The padding-vs-navigation subtlety (F4/F5)
is resolved by the `ensureRightMost_suffix_nil` / `ensureLeftMost_suffix_cs`
bridges. Respects `prehash_key_before_comparison` via the `keyForComparison`
hypothesis (identity for Tendermint; the SMT/JMT prehash composes analogously). -/
theorem nonexistence_sound
    (H : HashFn) (hNoHash : ∀ b, H .noHash b = b)
    (s : ProofSpec) (hwf : WellFormed s)
    (root key value : Bytes)
    (hsorted : KeySorted H s root)
    (nep : NonExistenceProof) (ep : ExistenceProof)
    (hkey : ep.key = key)
    (hne : verifyNonExistence H nep s root key = true)
    (hex : verifyExistence H ep s root key value = true) :
    HashCollision H := by
  sorry

end Ics23
