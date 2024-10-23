package ics23

import (
	"bytes"
	"crypto"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"

	// adds sha256 capability to crypto.SHA256
	_ "crypto/sha256"
	// adds sha512 capability to crypto.SHA512
	_ "crypto/sha512"
	// adds blake2b capability to crypto.BLAKE2b_512
	_ "golang.org/x/crypto/blake2b"
	// adds blake2s capability to crypto.BLAKE2s_256
	_ "golang.org/x/crypto/blake2s"
	// adds ripemd160 capability to crypto.RIPEMD160
	_ "golang.org/x/crypto/ripemd160" //nolint:staticcheck
)

// validateIavlOps validates the prefix to ensure it begins with
// the height, size, and version of the IAVL tree. Each varint must be a bounded value.
// In addition, the remaining bytes are validated to ensure they correspond to the correct
// length. The layerNum is the inverse of the tree depth, i.e. depth=0 means leaf, depth>=1 means inner node
func validateIavlOps(op opType, layerNum int) error {
	r := bytes.NewReader(op.GetPrefix())

	height, err := binary.ReadVarint(r)
	if err != nil {
		return fmt.Errorf("failed to read IAVL height varint: %w", err)
	}
	if int(height) < 0 || int(height) < layerNum {
		return fmt.Errorf("IAVL height (%d) must be non-negative and greater than or equal to the layer number (%d)", height, layerNum)
	}

	size, err := binary.ReadVarint(r)
	if err != nil {
		return fmt.Errorf("failed to read IAVL size varint: %w", err)
	}

	if int(size) < 0 {
		return fmt.Errorf("IAVL size must be non-negative")
	}

	version, err := binary.ReadVarint(r)
	if err != nil {
		return fmt.Errorf("failed to read IAVL version varint: %w", err)
	}

	if int(version) < 0 {
		return fmt.Errorf("IAVL version must be non-negative")
	}

	// grab the length of the remainder of the prefix
	remLen := r.Len()
	if layerNum == 0 {
		// leaf node
		if remLen != 0 {
			return fmt.Errorf("expected remaining prefix length to be 0, got: %d", remLen)
		}
		if height != 0 {
			return fmt.Errorf("expected leaf node height to be 0, got: %d", remLen)
		}
		if size != 1 {
			return fmt.Errorf("expected leaf node size to be 1, got: %d", remLen)
		}
	} else {
		// inner node
		//
		// when the child comes from the left, the suffix if filled in
		// prefix: height | size | version | length byte (1 remainder)
		//
		// when the child comes from the right, the suffix is empty
		// prefix: height | size | version | length byte | 32 byte hash | next length byte (34 remainder)
		if remLen != 1 && remLen != 34 {
			return fmt.Errorf("remainder of prefix must be of length 1 or 34, got: %d", remLen)
		}
		if op.GetHash() != HashOp_SHA256 {
			return fmt.Errorf("IAVL hash op must be %v", HashOp_SHA256)
		}
	}
	return nil
}

// validateTendermintOps validates the prefix to ensure it begins with []byte{1}.
func validateTendermintOps(op *InnerOp) error {
	if len(op.Prefix) == 0 {
		return fmt.Errorf("inner op prefix must not be empty")
	}

	innerPrefix := []byte{1}
	if op.Suffix != nil {
		if !bytes.Equal(op.Prefix, innerPrefix) {
			return fmt.Errorf("expected inner op prefix: %v, got: %v", innerPrefix, op.Prefix)
		}
	}

	if !bytes.HasPrefix(op.Prefix, innerPrefix) {
		return fmt.Errorf("expected inner op prefix to begin with: %v, got: %v", innerPrefix, op.Prefix[:1])
	}

	return nil
}

// Apply will calculate the leaf hash given the key and value being proven
func (op *LeafOp) Apply(key []byte, value []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, errors.New("leaf op needs key")
	}
	if len(value) == 0 {
		return nil, errors.New("leaf op needs value")
	}
	pkey, err := prepareLeafData(op.PrehashKey, op.Length, key)
	if err != nil {
		return nil, fmt.Errorf("prehash key, %w", err)
	}
	pvalue, err := prepareLeafData(op.PrehashValue, op.Length, value)
	if err != nil {
		return nil, fmt.Errorf("prehash value, %w", err)
	}

	data := op.Prefix
	data = append(data, pkey...)
	data = append(data, pvalue...)

	return doHash(op.Hash, data)
}

// Apply will calculate the hash of the next step, given the hash of the previous step
func (op *InnerOp) Apply(child []byte) ([]byte, error) {
	if len(child) == 0 {
		return nil, errors.New("inner op needs child value")
	}
	preimage := op.Prefix
	preimage = append(preimage, child...)
	preimage = append(preimage, op.Suffix...)
	return doHash(op.Hash, preimage)
}

// CheckAgainstSpec will verify the LeafOp is in the format defined in spec
func (op *LeafOp) CheckAgainstSpec(spec *ProofSpec) error {
	if spec == nil {
		return errors.New("op and spec must be non-nil")
	}
	lspec := spec.LeafSpec
	if lspec == nil {
		return errors.New("spec.LeafSpec must be non-nil")
	}

	if spec.SpecEquals(IavlSpec) {
		err := validateIavlOps(op, 0)
		if err != nil {
			return err
		}
	}

	if op.Hash != lspec.Hash {
		return fmt.Errorf("unexpected HashOp: %d", op.Hash)
	}
	if op.PrehashKey != lspec.PrehashKey {
		return fmt.Errorf("unexpected PrehashKey: %d", op.PrehashKey)
	}
	if op.PrehashValue != lspec.PrehashValue {
		return fmt.Errorf("unexpected PrehashValue: %d", op.PrehashValue)
	}
	if op.Length != lspec.Length {
		return fmt.Errorf("unexpected LengthOp: %d", op.Length)
	}
	if !bytes.HasPrefix(op.Prefix, lspec.Prefix) {
		return fmt.Errorf("leaf Prefix doesn't start with %X", lspec.Prefix)
	}
	return nil
}

// CheckAgainstSpec will verify the InnerOp is in the format defined in spec
func (op *InnerOp) CheckAgainstSpec(spec *ProofSpec, b int) error {
	if spec == nil {
		return errors.New("op and spec must be both non-nil")
	}
	if spec.InnerSpec == nil {
		return errors.New("spec.InnerSpec must be non-nil")
	}
	if spec.LeafSpec == nil {
		return errors.New("spec.LeafSpec must be non-nil")
	}

	if op.Hash != spec.InnerSpec.Hash {
		return fmt.Errorf("unexpected HashOp: %d", op.Hash)
	}

	if spec.SpecEquals(IavlSpec) {
		err := validateIavlOps(op, b)
		if err != nil {
			return err
		}
	}

	if spec.SpecEquals(TendermintSpec) {
		err := validateTendermintOps(op)
		if err != nil {
			return err
		}
	}

	leafPrefix := spec.LeafSpec.Prefix
	if bytes.HasPrefix(op.Prefix, leafPrefix) {
		return fmt.Errorf("inner Prefix starts with %X", leafPrefix)
	}
	if len(op.Prefix) < int(spec.InnerSpec.MinPrefixLength) {
		return fmt.Errorf("innerOp prefix too short (%d)", len(op.Prefix))
	}
	maxLeftChildBytes := (len(spec.InnerSpec.ChildOrder) - 1) * int(spec.InnerSpec.ChildSize)
	if len(op.Prefix) > int(spec.InnerSpec.MaxPrefixLength)+maxLeftChildBytes {
		return fmt.Errorf("innerOp prefix too long (%d)", len(op.Prefix))
	}

	if spec.InnerSpec.ChildSize <= 0 {
		return errors.New("spec.InnerSpec.ChildSize must be >= 1")
	}

	if spec.InnerSpec.MaxPrefixLength >= spec.InnerSpec.MinPrefixLength+spec.InnerSpec.ChildSize {
		return errors.New("spec.InnerSpec.MaxPrefixLength must be < spec.InnerSpec.MinPrefixLength + spec.InnerSpec.ChildSize")
	}

	// ensures soundness, with suffix having to be of correct length
	if len(op.Suffix)%int(spec.InnerSpec.ChildSize) != 0 {
		return fmt.Errorf("InnerOp suffix malformed")
	}

	return nil
}

// doHash will preform the specified hash on the preimage.
// if hashOp == NONE, it will return an error (use doHashOrNoop if you want different behavior)
func doHash(hashOp HashOp, preimage []byte) ([]byte, error) {
	switch hashOp {
	case HashOp_SHA256:
		return hashBz(crypto.SHA256, preimage)
	case HashOp_SHA512:
		return hashBz(crypto.SHA512, preimage)
	case HashOp_RIPEMD160:
		return hashBz(crypto.RIPEMD160, preimage)
	case HashOp_BITCOIN:
		// ripemd160(sha256(x))
		tmp, err := hashBz(crypto.SHA256, preimage)
		if err != nil {
			return nil, err
		}
		return hashBz(crypto.RIPEMD160, tmp)
	case HashOp_SHA512_256:
		return hashBz(crypto.SHA512_256, preimage)
	case HashOp_BLAKE2B_512:
		return hashBz(crypto.BLAKE2b_512, preimage)
	case HashOp_BLAKE2S_256:
		return hashBz(crypto.BLAKE2s_256, preimage)
	}
	return nil, fmt.Errorf("unsupported hashop: %d", hashOp)
}

type hasher interface {
	New() hash.Hash
}

func hashBz(h hasher, preimage []byte) ([]byte, error) {
	hh := h.New()
	_, err := hh.Write(preimage)
	if err != nil {
		return nil, err
	}
	return hh.Sum(nil), nil
}

func prepareLeafData(hashOp HashOp, lengthOp LengthOp, data []byte) ([]byte, error) {
	// TODO: lengthop before or after hash ???
	hdata, err := doHashOrNoop(hashOp, data)
	if err != nil {
		return nil, err
	}

	return doLengthOp(lengthOp, hdata)
}

type opType interface {
	GetPrefix() []byte
	GetHash() HashOp
	Reset()
	String() string
}

// doLengthOp will calculate the proper prefix and return it prepended
//
//	doLengthOp(op, data) -> length(data) || data
func doLengthOp(lengthOp LengthOp, data []byte) ([]byte, error) {
	switch lengthOp {
	case LengthOp_NO_PREFIX:
		return data, nil
	case LengthOp_VAR_PROTO:
		res := append(encodeVarintProto(len(data)), data...)
		return res, nil
	case LengthOp_REQUIRE_32_BYTES:
		if len(data) != 32 {
			return nil, fmt.Errorf("data was %d bytes, not 32", len(data))
		}
		return data, nil
	case LengthOp_REQUIRE_64_BYTES:
		if len(data) != 64 {
			return nil, fmt.Errorf("data was %d bytes, not 64", len(data))
		}
		return data, nil
	case LengthOp_FIXED32_BIG:
		res := make([]byte, 4, 4+len(data))
		binary.BigEndian.PutUint32(res[:4], uint32(len(data)))
		res = append(res, data...)
		return res, nil
	case LengthOp_FIXED32_LITTLE:
		res := make([]byte, 4, 4+len(data))
		binary.LittleEndian.PutUint32(res[:4], uint32(len(data)))
		res = append(res, data...)
		return res, nil
	case LengthOp_FIXED64_BIG:
		res := make([]byte, 8, 8+len(data))
		binary.BigEndian.PutUint64(res[:8], uint64(len(data)))
		res = append(res, data...)
		return res, nil
	case LengthOp_FIXED64_LITTLE:
		res := make([]byte, 8, 8+len(data))
		binary.LittleEndian.PutUint64(res[:8], uint64(len(data)))
		res = append(res, data...)
		return res, nil
		// TODO
		// case LengthOp_VAR_RLP:
	}
	return nil, fmt.Errorf("unsupported lengthop: %d", lengthOp)
}

// doHashOrNoop will return the preimage untouched if hashOp == NONE,
// otherwise, perform doHash
func doHashOrNoop(hashOp HashOp, preimage []byte) ([]byte, error) {
	if hashOp == HashOp_NO_HASH {
		return preimage, nil
	}
	return doHash(hashOp, preimage)
}
