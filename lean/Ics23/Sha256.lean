/-
A concrete SHA-256 in Lean, so the verifier model becomes executable end to end
(Phase 2a groundwork: a concrete `HashFn` to run the model and, eventually,
differentially test it against the Rust/Go implementations).

Validated below against the vectors embedded in `rust/src/ops.rs`.
-/
import Ics23.Ops

namespace Ics23.Sha256

/-- Right-rotate a 32-bit word. -/
def rotr (x : UInt32) (n : UInt32) : UInt32 := (x >>> n) ||| (x <<< (32 - n))

def K : List UInt32 :=
  [0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5,
   0xd807aa98, 0x12835b01, 0x243185be, 0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174,
   0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa, 0x5cb0a9dc, 0x76f988da,
   0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967,
   0x27b70a85, 0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85,
   0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3, 0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070,
   0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f, 0x682e6ff3,
   0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2]

def H0 : List UInt32 :=
  [0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19]

/-- Pad a message to a multiple of 64 bytes per the SHA-256 spec. -/
def pad (msg : List UInt8) : List UInt8 :=
  let len := msg.length
  let bitLen : UInt64 := (UInt64.ofNat len) * 8
  -- 0x80, then k zero bytes so total ≡ 56 (mod 64), then 8-byte big-endian bit length.
  -- `120 - r` keeps the subtraction non-truncating for r ∈ [0, 63]: Nat `56 - r`
  -- clamps to 0 for r > 56, which dropped the spill-over padding block for
  -- messages with `length % 64 ∈ [56, 62]` (caught by the IAVL test vectors —
  -- their 57-byte leaf preimages land exactly in that window).
  let withOne := msg ++ [0x80]
  let zeros := (120 - (withOne.length % 64)) % 64
  let lenBytes : List UInt8 :=
    (List.range 8).map (fun i => UInt8.ofNat ((bitLen >>> (UInt64.ofNat (8 * (7 - i)))).toNat % 256))
  withOne ++ List.replicate zeros 0 ++ lenBytes

/-- Pack 4 big-endian bytes into a word. -/
def beWord (a b c d : UInt8) : UInt32 :=
  (a.toUInt32 <<< 24) ||| (b.toUInt32 <<< 16) ||| (c.toUInt32 <<< 8) ||| d.toUInt32

/-- Split a 64-byte block into its first 16 message-schedule words. -/
def blockWords (block : List UInt8) : Array UInt32 := Id.run do
  let arr := block.toArray
  let mut ws : Array UInt32 := #[]
  for i in [0:16] do
    ws := ws.push (beWord arr[4*i]! arr[4*i+1]! arr[4*i+2]! arr[4*i+3]!)
  return ws

/-- Extend 16 schedule words to 64. -/
def schedule (w0 : Array UInt32) : Array UInt32 := Id.run do
  let mut w := w0
  for i in [16:64] do
    let s0 := rotr w[i-15]! 7 ^^^ rotr w[i-15]! 18 ^^^ (w[i-15]! >>> 3)
    let s1 := rotr w[i-2]! 17 ^^^ rotr w[i-2]! 19 ^^^ (w[i-2]! >>> 10)
    w := w.push (w[i-16]! + s0 + w[i-7]! + s1)
  return w

/-- Compress one block into the running hash state. -/
def compress (st : Array UInt32) (w : Array UInt32) : Array UInt32 := Id.run do
  let mut a := st[0]!; let mut b := st[1]!; let mut c := st[2]!; let mut d := st[3]!
  let mut e := st[4]!; let mut f := st[5]!; let mut g := st[6]!; let mut h := st[7]!
  let karr := K.toArray
  for i in [0:64] do
    let s1 := rotr e 6 ^^^ rotr e 11 ^^^ rotr e 25
    let ch := (e &&& f) ^^^ ((~~~e) &&& g)
    let t1 := h + s1 + ch + karr[i]! + w[i]!
    let s0 := rotr a 2 ^^^ rotr a 13 ^^^ rotr a 22
    let maj := (a &&& b) ^^^ (a &&& c) ^^^ (b &&& c)
    let t2 := s0 + maj
    h := g; g := f; f := e; e := d + t1; d := c; c := b; b := a; a := t1 + t2
  return #[st[0]! + a, st[1]! + b, st[2]! + c, st[3]! + d,
           st[4]! + e, st[5]! + f, st[6]! + g, st[7]! + h]

/-- Hash a message, returning the 32-byte digest. -/
def hash (msg : List UInt8) : List UInt8 := Id.run do
  let padded := (pad msg).toArray
  let mut st := H0.toArray
  let nBlocks := padded.size / 64
  for blk in [0:nBlocks] do
    let block := (padded.extract (blk*64) (blk*64+64)).toList
    st := compress st (schedule (blockWords block))
  -- serialize state big-endian
  let mut out : List UInt8 := []
  for word in st.toList do
    out := out ++ [UInt8.ofNat ((word >>> 24).toNat % 256), UInt8.ofNat ((word >>> 16).toNat % 256),
                   UInt8.ofNat ((word >>> 8).toNat % 256), UInt8.ofNat (word.toNat % 256)]
  return out

/-- Hex string of a digest, for validation. -/
def toHex (bs : List UInt8) : String :=
  let digit : Nat → Char := fun n =>
    if n < 10 then Char.ofNat (48 + n) else Char.ofNat (97 + n - 10)
  String.ofList (bs.flatMap (fun b => [digit (b.toNat / 16), digit (b.toNat % 16)]))

-- "food" → c1f026582fe6e8cb620d0c85a72fe421ddded756662a8ec00ed4c297ad10676b (rust/src/ops.rs)
example : toHex (hash [0x66, 0x6f, 0x6f, 0x64])
    = "c1f026582fe6e8cb620d0c85a72fe421ddded756662a8ec00ed4c297ad10676b" := by native_decide

-- "foobar" (apply_leaf of foo/bar, no prehash/length) →
-- c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2
example : toHex (hash [0x66, 0x6f, 0x6f, 0x62, 0x61, 0x72])
    = "c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2" := by native_decide

-- empty string → e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
example : toHex (hash [])
    = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" := by native_decide

/-! Padding-boundary regressions: lengths around `% 64 ∈ [55, 64]`, where the
8-byte length field no longer fits the final block and padding must spill into
an extra block. The original `pad` used truncating `Nat` subtraction and
silently dropped bytes exactly in the `[56, 62]` window — missed by the three
vectors above and by every Tendermint/SMT preimage, caught by the IAVL
differential vectors (TestVectors.lean). Digests cross-checked against
Python's `hashlib`. -/

example : toHex (hash (List.replicate 55 0x61))
    = "9f4390f8d30c2dd92ec9f095b65e2b9ae9b0a925a5258e241c9f1e910f734318" := by native_decide
example : toHex (hash (List.replicate 56 0x61))
    = "b35439a4ac6f0948b6d6f9e3c6af0f5f590ce20f1bde7090ef7970686ec6738a" := by native_decide
example : toHex (hash (List.replicate 57 0x61))
    = "f13b2d724659eb3bf47f2dd6af1accc87b81f09f59f2b75e5c0bed6589dfe8c6" := by native_decide
example : toHex (hash (List.replicate 62 0x61))
    = "f506898cc7c2e092f9eb9fadae7ba50383f5b46a2a4fe5597dbb553a78981268" := by native_decide
example : toHex (hash (List.replicate 63 0x61))
    = "7d3e74a05d7db15bce4ad9ec0658ea98e3f06eeecf16b4c6fff2da457ddc2f34" := by native_decide
example : toHex (hash (List.replicate 64 0x61))
    = "ffe054fe7ae0cb6dc65c3af9b61d5209f439851db43d0ba5997337df154668eb" := by native_decide

end Ics23.Sha256
