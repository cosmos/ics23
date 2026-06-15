# ICS23 formal model (Lean 4)

A mechanized model of the ICS23 verifier with machine-checked soundness proofs,
per [`docs/rfc/001-formal-verification.md`](../docs/rfc/001-formal-verification.md).
Property catalogue: [`docs/verification/properties.md`](../docs/verification/properties.md).

## Build

Requires [`elan`](https://github.com/leanprover/elan) (the toolchain is pinned in
`lean-toolchain`; `lake` fetches it automatically).

```sh
cd lean
lake build
```

A clean build prints one expected warning — the *abstract-root*
`nonexistence_sound` (NonExistSound.lean) is the one deliberate `sorry`: an
opaque root carries no tree structure, so byte-order cannot be connected to
tree position (findings F3/F4). Both theorems are instead proved in the
honest-root tree models (below). Everything else is proof-complete.

## Layout

| File | Mirrors | Contents |
|------|---------|----------|
| `Ics23/Types.lean` | `proofs.proto`, `cosmos.ics23.v1.rs` | proto types |
| `Ics23/Ops.lean` | `rust/src/ops.rs` | `applyLeaf`, `applyInner`, `doHash` family, `doLength` |
| `Ics23/Verify.lean` | `rust/src/verify.rs` | existence verifier |
| `Ics23/NonExist.lean` | `rust/src/verify.rs` | non-existence verifier |
| `Ics23/Specs.lean` | `rust/src/api.rs` | IAVL / Tendermint / SMT specs |
| `Ics23/Varint.lean` | — | varint self-delimiting lemmas (A1) |
| `Ics23/Soundness.lean` | — | `WellFormed`, collisions, Theorem C |
| `Ics23/Existence.lean` | — | Theorem A, byte-level (`existence_binding`) |
| `Ics23/NonExistSound.lean` | — | byte order; abstract-root Theorem B (the `sorry`) |
| `Ics23/Order.lean` | — | order lemmas for the neighbor checks |
| `Ics23/LeafInj.lean` | — | joint leaf injectivity, proved for all three specs |
| `Ics23/Tree.lean` | — | honest-root `MTree` model; Theorem A (`membership_sound`) |
| `Ics23/TreeNonExist.lean` | — | Theorem B, total (`nonexistence_sound_tree_total`) |
| `Ics23/SmtTree.lean` | — | sparse (SMT/JMT) model; Theorem A (`membership_sound_smt`) |
| `Ics23/SmtNonExist.lean` | — | Theorem B for SMT (`nonexistence_sound_smt_total`) |
| `Ics23/IavlTree.lean` | — | IAVL model (`WFTreeI`); Theorem A (`membership_sound_iavl`) |
| `Ics23/IavlNonExist.lean` | — | Theorem B for IAVL (`nonexistence_sound_iavl_total`) |
| `Ics23/IavlPrefix.lean` | — | IAVL prefix structure; F3 witness |
| `Ics23/Sha256.lean` | — | pure-Lean SHA-256 (validated, incl. padding boundaries) |
| `Ics23/Executable.lean` | — | end-to-end executable runs, forgery refutations |
| `Ics23/Corpus.lean` | — | regression corpus (proven accept/reject facts) |
| `Ics23/TestVectors.lean` | `testdata/` | GENERATED differential vectors (`rust/examples/lean_testdata.rs`) |

The model is parameterized over an abstract hash family and makes no
collision-resistance assumption: soundness theorems conclude by exhibiting a
`HashCollision`. See the modeling-assumptions section of the property catalogue
for where (and why) the model intentionally differs from the Rust.

## Status

- **Model complete:** existence and non-existence verifiers, both executable
  over an abstract hash family.
- **Proved:**
  - Spec well-formedness certificates for all three shipped specs (Theorem C:
    `iavl_wellFormed`, `tendermint_wellFormed`, `smt_wellFormed`).
  - Regression corpus: spec-level invariant violations (domain separation,
    child-size, prefix-window, malformed specs, depth bounds) as machine-checked
    facts (`Ics23/Corpus.lean`).
  - Building blocks of Theorem A: `innerImage_inj` (preimage cancellation),
    `applyInner_inj` (one step injective up to a collision), and
    `applyPath_sameops_inj` (folding a shared op-list is injective up to a
    collision) — the inductive backbone.
  - The varint length prefix is self-delimiting (`varintEncode_append_inj`,
    `Varint.lean`), giving leaf-encoding injectivity (A1) for length-prefixed
    specs; `doLength_varProto_inj` / `doLength_noPrefix_inj`.
  - **Same-shape existence binding (Theorem A) for all three shipped leaf
    shapes**: `existence_binding_sameshape` (general, parameterized by length
    injectivity) with corollaries `existence_binding_sameshape_noPrefix` (SMT/
    JMT) and `existence_binding_sameshape_varProto` (IAVL / Tendermint). Two
    proofs sharing tree shape that bind one key to two values force a hash
    collision — the value-swap forgery, end to end.
  - **Equal-length binding** `existence_binding_eqlen` — strengthens the above
    to proofs of equal *depth* with arbitrary (differing) inner ops, concluding
    the honest disjunction `HashCollision ∨ PositionalAmbiguity` (the F3
    obstacle). Built on `applyPath_eqlen_merge`.
  - **Same-leaf binding, any depth** `existence_binding_sameleaf` — the strongest
    result: same leaf op, *arbitrary differing-length* paths ⇒ `HashCollision ∨
    PositionalAmbiguity`. Built on the root-side structural core `applyPath_merge`
    (+ `applyPath_snoc`) and leaf/inner domain separation
    (`leafHash_innerImage_collision`). Instantiated for all three shipped specs:
    `existence_binding_{iavl,tendermint,smt}` (side conditions closed by `decide`).
  - **Finding F3 is formalized and machine-checked** (`PositionalAmbiguity`,
    witnesses in `Executable.lean`/`IavlPrefix.lean`): the general
    `existence_binding` is correctly stated as a disjunction, since a
    collision-only conclusion is provably too strong byte-level.
  - Byte-ordering facts behind the neighbor checks: `bytesLt_irrefl`,
    `bytesLt_ne` (`NonExistSound.lean`).
  - **General existence binding — fully proved** (`existence_binding_shaped`):
    for the production-spec shape, two proofs binding one key to two values under
    one root — with *no* assumption on their leaf ops or paths — yield the honest
    three-way disjunction `HashCollision ∨ PositionalAmbiguity ∨ LeafAmbiguity`.
    The ambiguity arms are real machine-checkable obstructions (F3 + leaf-level
    analogue); collapsing them needs the symbolic-Merkle model. Built on
    `applyPath_merge`, `ensureLeaf_eq`, `leafHash_innerImage_collision`.
  - **Honest-root Theorem A** (`Tree.lean`): `membership_sound` — an existence
    proof verifying against `root = rootHash t` of a real tree implies genuine
    membership, no ambiguity arm (`split_pins` resolves F3 against a real
    node). Instantiated: `membership_sound_tendermint` with leaf injectivity
    *proved* (`LeafInj.lean`), leaving `FixedHash` as the only assumption.
  - **Honest-root Theorem B, total** (`TreeNonExist.lean`):
    `nonexistence_sound_tree_total` — for a key-sorted well-formed tree, any
    verifying non-existence proof plus a verifying existence proof for the
    same key yields a `HashCollision`, covering all verifier-accepted shapes
    (two-sided / left-only / right-only). Instantiated:
    `nonexistence_sound_tree_tendermint_total`.
  - **SMT/JMT instantiation** (`SmtTree.lean`, `SmtNonExist.lean`): both
    theorems for the sparse model, where empty subtrees hash to the
    `empty_child` placeholder and keys sort by their SHA-256 prehash
    (`keyForComparison`). One added assumption, `EmptyChildFree` (no exhibited
    preimage of the all-zero placeholder — the standard SMT assumption):
    `membership_sound_smt`, `nonexistence_sound_smt_total`.
  - **IAVL instantiation** (`IavlTree.lean`, `IavlNonExist.lean`): both
    theorems for the IAVL shape — variable 4–12 byte prefixes and 33-byte
    amino child slots — via `split_pins_var`, which pins a node split from
    the suffix length alone (no prefix pinning, so `min ≠ max` is fine):
    `membership_sound_iavl`, `nonexistence_sound_iavl_total`. **Honest-root
    Theorems A and B now cover all three shipped specs.**
- **Stated, deliberately unproved (the one whitelisted `sorry`):**
  - Abstract-root Theorem B (`nonexistence_sound`, NonExistSound.lean) — an
    opaque root has no tree structure to connect byte-order to position
    (F3/F4); adjacency is inherently a statement about a real tree, which is
    exactly what the honest-root theorems above capture.
- **Executable end to end:** a concrete SHA-256 (`Sha256.lean`, validated
  against the vectors in `rust/src/ops.rs` plus padding-boundary regressions)
  and `concreteHash` make the verifier runnable; `Executable.lean` computes
  real roots and refutes value-swap / wrong-shape forgeries by
  `native_decide`.
- **Phase 2a differential oracle:** `rust/examples/lean_testdata.rs` decodes
  the shared `testdata/` vectors and generates `TestVectors.lean`, so
  `lake build` re-verifies all 18 vectors against the model with real
  SHA-256; CI fails on regeneration drift. Its first run caught a real bug:
  `Sha256.pad`'s truncating `Nat` subtraction dropped the spill-over padding
  block for lengths `≡ 56..62 (mod 64)` — missed by every Tendermint/SMT
  preimage, exposed by IAVL's 57-byte leaf preimages. (Executable layer only;
  the soundness theorems are abstract over `HashFn`.)
- **CI:** `.github/workflows/lean.yml` builds all proofs and fails if any
  unexpected `sorry` appears (exactly one is whitelisted).
- **Next:** transcribe the Zellic findings into the corpus; optionally extend
  the differential oracle to the negative/malformed unit-test vectors
  (`testdata/TestCheckAgainstSpecData.json` etc.).
