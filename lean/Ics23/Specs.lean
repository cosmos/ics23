/-
The three shipped proof specs, mirroring `iavl_spec`, `tendermint_spec`, and
`smt_spec` in `rust/src/api.rs`. `Soundness.lean` certifies each is `WellFormed`.
-/
import Ics23.Types

namespace Ics23

/-- IAVL tree spec (`iavl_spec`). -/
def iavlSpec : ProofSpec where
  leafSpec :=
    { hash := .sha256, prehashKey := .noHash, prehashValue := .sha256,
      length := .varProto, prefixBytes := [0] }
  innerSpec :=
    { childOrder := [0, 1], childSize := 33, minPrefixLength := 4,
      maxPrefixLength := 12, emptyChild := [], hash := .sha256 }
  minDepth := 0
  maxDepth := 0
  prehashKeyBeforeComparison := false

/-- CometBFT/Tendermint simple-merkle spec (`tendermint_spec`). -/
def tendermintSpec : ProofSpec where
  leafSpec :=
    { hash := .sha256, prehashKey := .noHash, prehashValue := .sha256,
      length := .varProto, prefixBytes := [0] }
  innerSpec :=
    { childOrder := [0, 1], childSize := 32, minPrefixLength := 1,
      maxPrefixLength := 1, emptyChild := [], hash := .sha256 }
  minDepth := 0
  maxDepth := 0
  prehashKeyBeforeComparison := false

/-- Sparse-merkle-tree / JMT spec (`smt_spec`). Note `prehashKeyBeforeComparison`
and the 32-byte `emptyChild`. -/
def smtSpec : ProofSpec where
  leafSpec :=
    { hash := .sha256, prehashKey := .sha256, prehashValue := .sha256,
      length := .noPrefix, prefixBytes := [0] }
  innerSpec :=
    { childOrder := [0, 1], childSize := 32, minPrefixLength := 1,
      maxPrefixLength := 1,
      emptyChild := List.replicate 32 0, hash := .sha256 }
  minDepth := 0
  maxDepth := 0
  prehashKeyBeforeComparison := true

end Ics23
