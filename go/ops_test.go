package ics23

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"
)

func TestValidateIavlOps(t *testing.T) {
	var (
		op       opType
		layerNum int
	)
	cases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success",
			func() {},
			nil,
		},
		{
			"failure: reading varint",
			func() {
				op.(*InnerOp).Prefix = []byte{}
			},
			errors.New("failed to read IAVL height varint: EOF"),
		},
		{
			"failure: invalid height value",
			func() {
				op.(*InnerOp).Prefix = []byte{1}
			},
			errors.New("IAVL height (-1) must be non-negative and greater than or equal to the layer number (1)"),
		},
		{
			"failure: invalid size varint",
			func() {
				var varintBuf [binary.MaxVarintLen64]byte
				prefix := convertVarIntToBytes(int64(5), varintBuf)
				op.(*InnerOp).Prefix = prefix
			},
			errors.New("failed to read IAVL size varint: EOF"),
		},
		{
			"failure: invalid size value",
			func() {
				var varintBuf [binary.MaxVarintLen64]byte
				prefix := convertVarIntToBytes(int64(5), varintBuf)
				prefix = append(prefix, convertVarIntToBytes(int64(-1), varintBuf)...) // size
				op.(*InnerOp).Prefix = prefix
			},
			errors.New("IAVL size must be non-negative"),
		},
		{
			"failure: invalid version varint",
			func() {
				var varintBuf [binary.MaxVarintLen64]byte
				prefix := convertVarIntToBytes(int64(5), varintBuf)
				prefix = append(prefix, convertVarIntToBytes(int64(10), varintBuf)...)
				op.(*InnerOp).Prefix = prefix
			},
			errors.New("failed to read IAVL version varint: EOF"),
		},
		{
			"failure: invalid version value",
			func() {
				var varintBuf [binary.MaxVarintLen64]byte
				prefix := convertVarIntToBytes(int64(5), varintBuf)
				prefix = append(prefix, convertVarIntToBytes(int64(10), varintBuf)...)
				prefix = append(prefix, convertVarIntToBytes(int64(-1), varintBuf)...) // version
				op.(*InnerOp).Prefix = prefix
			},
			errors.New("IAVL version must be non-negative"),
		},
		{
			"failure: invalid remaining length with layer number is 0",
			func() {
				layerNum = 0
			},
			fmt.Errorf("expected remaining prefix length to be 0, got: 1"),
		},
		{
			"failure: invalid remaining length with non-zero layer number",
			func() {
				layerNum = 1
				op.(*InnerOp).Prefix = append(op.(*InnerOp).Prefix, []byte{1}...)
			},
			fmt.Errorf("remainder of prefix must be of length 1 or 34, got: 2"),
		},
		{
			"failure: invalid hash",
			func() {
				op.(*InnerOp).Hash = HashOp_NO_HASH
			},
			fmt.Errorf("IAVL hash op must be %v", HashOp_SHA256),
		},
	}
	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			op = &InnerOp{
				Hash:   HashOp_SHA256,
				Prefix: generateInnerOpPrefix(),
				Suffix: []byte(""),
			}
			layerNum = 1

			tc.malleate()

			err := validateIavlOps(op, layerNum)
			if tc.expError == nil && err != nil {
				t.Fatal(err)
			}

			if tc.expError != nil && err.Error() != tc.expError.Error() {
				t.Fatalf("expected: %v, got: %v", tc.expError, err)
			}
		})

	}
}

func TestValidateTendermintOps(t *testing.T) {
	var op *InnerOp
	cases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success: valid prefix when suffix populated",
			func() {},
			nil,
		},
		{
			"success: valid prefix when suffix empty",
			func() {
				op.Prefix = []byte{1, 2}
				op.Suffix = nil
			},
			nil,
		},
		{
			"failure: empty prefix and suffix",
			func() {
				op.Prefix = nil
				op.Suffix = nil
			},
			errors.New("inner op prefix must not be empty"),
		},
		{
			"failure: invalid prefix when suffix populated",
			func() {
				op.Prefix = []byte{0}
				op.Suffix = []byte{1}
			},
			fmt.Errorf("expected inner op prefix: %v, got: %v", []byte{1}, []byte{0}),
		},
		{
			"failure: invalid prefix when suffix empty",
			func() {
				op.Prefix = []byte{2, 1}
				op.Suffix = nil
			},
			fmt.Errorf("expected inner op prefix to begin with: %v, got: %v", []byte{1}, []byte{2}),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			op = &InnerOp{
				Hash:   HashOp_SHA256,
				Prefix: []byte{1},
				Suffix: []byte{1},
			}

			tc.malleate()

			err := validateTendermintOps(op)
			if tc.expError == nil && err != nil {
				t.Fatal(err)
			}

			if tc.expError != nil && err.Error() != tc.expError.Error() {
				t.Fatalf("expected: %v, got: %v", tc.expError, err)
			}
		})

	}
}

func TestLeafOp(t *testing.T) {
	cases := LeafOpTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := tc.Op.Apply(tc.Key, tc.Value)
			// short-circuit with error case
			if tc.IsErr && err == nil {
				t.Fatal("expected error, but got none")
			}
			if !tc.IsErr && err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(res, tc.Expected) {
				t.Errorf("bad result: %s vs %s", toHex(res), toHex(tc.Expected))
			}
		})
	}
}

func TestInnerOp(t *testing.T) {
	cases := InnerOpTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := tc.Op.Apply(tc.Child)
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got none")
			}
			if !tc.IsErr && err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(res, tc.Expected) {
				t.Errorf("Bad result: %s vs %s", toHex(res), toHex(tc.Expected))
			}
		})
	}
}

func TestDoHash(t *testing.T) {
	cases := DoHashTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := doHash(tc.HashOp, []byte(tc.Preimage))
			if err != nil {
				t.Fatal(err)
			}
			hexRes := hex.EncodeToString(res)
			if hexRes != tc.ExpectedHash {
				t.Fatalf("Expected %s got %s", tc.ExpectedHash, hexRes)
			}
		})
	}
}

func TestInnerOpCheckAgainstSpec(t *testing.T) {
	var (
		spec    *ProofSpec
		innerOp *InnerOp
	)

	cases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success",
			func() {},
			nil,
		},
		{
			"failure: empty spec",
			func() {
				spec = nil
			},
			errors.New("op and spec must be both non-nil"),
		},
		{
			"failure: empty inner spec",
			func() {
				spec.InnerSpec = nil
			},
			errors.New("spec.InnerSpec must be non-nil"),
		},
		{
			"failure: empty leaf spec",
			func() {
				spec.LeafSpec = nil
			},
			errors.New("spec.LeafSpec must be non-nil"),
		},
		{
			"failure: incorrect hash function inner op",
			func() {
				innerOp.Hash = HashOp_BITCOIN
			},
			fmt.Errorf("unexpected HashOp: %d", HashOp_BITCOIN),
		},
		{
			"failure: spec fails validation",
			func() {
				innerOp.Prefix = []byte{0x01}
			},
			fmt.Errorf("IAVL height (-1) must be non-negative and greater than or equal to the layer number (1)"),
		},
		{
			"failure: inner prefix starts with leaf prefix",
			func() {
				// change spec to be non-iavl spec to skip strict iavl validation
				spec.LeafSpec.PrehashKey = HashOp_BITCOIN
				innerOp.Prefix = append(spec.LeafSpec.Prefix, innerOp.Prefix...)
			},
			fmt.Errorf("inner Prefix starts with %X", []byte{0}),
		},
		{
			"failure: inner prefix too short",
			func() {
				// change spec to be non-iavl spec to skip strict iavl validation
				spec.LeafSpec.PrehashKey = HashOp_BITCOIN
				innerOp.Prefix = []byte{0x01}
			},
			errors.New("innerOp prefix too short (1)"),
		},
		{
			"failure: inner prefix too long",
			func() {
				// change spec to be non-iavl spec to skip strict iavl validation
				spec.LeafSpec.PrehashKey = HashOp_BITCOIN
				innerOp.Prefix = []byte("AgQIIGe3bHuC1g6+5/Qd0RoCU0waFu+nDCFzEDViMN/VrQwgIA==")
			},
			errors.New("innerOp prefix too long (52)"),
		},
		{
			"failure: child size must be greater than zero",
			func() {
				spec.InnerSpec.ChildSize = 0
			},
			errors.New("spec.InnerSpec.ChildSize must be >= 1"),
		},
		{
			"failure: MaxPrefixLength >= MinPrefixLength + ChildSize",
			func() {
				spec.InnerSpec.MaxPrefixLength = spec.InnerSpec.MinPrefixLength + spec.InnerSpec.ChildSize
			},
			errors.New("spec.InnerSpec.MaxPrefixLength must be < spec.InnerSpec.MinPrefixLength + spec.InnerSpec.ChildSize"),
		},
		{
			"failure: inner op suffix malformed",
			func() {
				innerOp.Suffix = []byte{0x01}
			},
			fmt.Errorf("InnerOp suffix malformed"),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			spec = &ProofSpec{ // IavlSpec
				LeafSpec: &LeafOp{
					Prefix:       []byte{0},
					PrehashKey:   HashOp_NO_HASH,
					Hash:         HashOp_SHA256,
					PrehashValue: HashOp_SHA256,
					Length:       LengthOp_VAR_PROTO,
				},
				InnerSpec: &InnerSpec{
					ChildOrder:      []int32{0, 1},
					MinPrefixLength: 4,
					MaxPrefixLength: 12,
					ChildSize:       33, // (with length byte)
					EmptyChild:      nil,
					Hash:            HashOp_SHA256,
				},
			}

			innerOp = &InnerOp{
				Hash:   HashOp_SHA256,
				Prefix: generateInnerOpPrefix(),
				Suffix: []byte(""),
			}
			tc.malleate()

			err := innerOp.CheckAgainstSpec(spec, 1) // use a layer number of 1
			if tc.expError == nil && err != nil {
				t.Fatal(err)
			}

			if tc.expError != nil && err.Error() != tc.expError.Error() {
				t.Fatalf("expected: %v, got: %v", tc.expError, err)
			}
		})

	}
}

func TestForgeNonExistenceProofWithIncorrectMaxPrefixLength(t *testing.T) {
	spec := &ProofSpec{ // TendermintSpec
		LeafSpec: &LeafOp{
			Prefix:       []byte{0},
			PrehashKey:   HashOp_NO_HASH,
			Hash:         HashOp_SHA256,
			PrehashValue: HashOp_SHA256,
			Length:       LengthOp_VAR_PROTO,
		},
		InnerSpec: &InnerSpec{
			ChildOrder:      []int32{0, 1},
			MinPrefixLength: 1,
			MaxPrefixLength: 1,
			ChildSize:       32, // (no length byte)
			Hash:            HashOp_SHA256,
		},
	}

	spec.InnerSpec.MaxPrefixLength = 33
	leafOp := spec.LeafSpec
	aLeaf, _ := leafOp.Apply([]byte("a"), []byte("a"))
	bLeaf, _ := leafOp.Apply([]byte("b"), []byte("b"))
	b2Leaf, _ := leafOp.Apply([]byte("b2"), []byte("b2"))

	cLeaf, _ := leafOp.Apply([]byte("c"), []byte("c"))
	aExist := ExistenceProof{
		Key:   []byte("a"),
		Value: []byte("a"),
		Leaf:  leafOp,
		Path: []*InnerOp{
			{
				Hash:   spec.InnerSpec.Hash,
				Prefix: []byte{1},
				Suffix: append(bLeaf, b2Leaf...),
			},
			{
				Hash:   spec.InnerSpec.Hash,
				Prefix: []byte{1},
				Suffix: cLeaf,
			},
		},
	}
	bExist := ExistenceProof{
		Key:   []byte("b"),
		Value: []byte("b"),
		Leaf:  leafOp,
		Path: []*InnerOp{
			{
				Hash:   spec.InnerSpec.Hash,
				Prefix: append([]byte{1}, aLeaf...),
				Suffix: b2Leaf,
			},
			{
				Hash:   spec.InnerSpec.Hash,
				Prefix: []byte{1},
				Suffix: cLeaf,
			},
		},
	}
	b2Exist := ExistenceProof{
		Key:   []byte("b2"),
		Value: []byte("b2"),
		Leaf:  leafOp,
		Path: []*InnerOp{
			{
				Hash:   spec.InnerSpec.Hash,
				Prefix: append(append([]byte{1}, aLeaf...), bLeaf...),
				Suffix: []byte{},
			},
			{
				Hash:   spec.InnerSpec.Hash,
				Prefix: []byte{1},
				Suffix: cLeaf,
			},
		},
	}
	yHash, _ := aExist.Path[0].Apply(aLeaf)
	cExist := ExistenceProof{
		Key:   []byte("c"),
		Value: []byte("c"),
		Leaf:  leafOp,
		Path: []*InnerOp{
			{
				Hash:   spec.InnerSpec.Hash,
				Prefix: append([]byte{1}, yHash...),
				Suffix: []byte{},
			},
		},
	}
	aNotExist := NonExistenceProof{
		Key:   []byte("a"),
		Left:  nil,
		Right: &bExist,
	}
	root, err := aExist.Calculate()
	if err != nil {
		t.Fatal("failed to calculate existence proof of leaf a")
	}

	expError := fmt.Errorf("inner, %w", errors.New("spec.InnerSpec.MaxPrefixLength must be < spec.InnerSpec.MinPrefixLength + spec.InnerSpec.ChildSize"))
	err = aExist.Verify(spec, root, []byte("a"), []byte("a"))
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove existence of leaf a returned incorrect error")
	}

	err = bExist.Verify(spec, root, []byte("b"), []byte("b"))
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove existence of leaf b returned incorrect error")
	}

	err = b2Exist.Verify(spec, root, []byte("b2"), []byte("b2"))
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove existence of third leaf returned incorrect error")
	}

	err = cExist.Verify(spec, root, []byte("c"), []byte("c"))
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove existence of leaf c returned incorrect error")
	}

	err = aNotExist.Verify(spec, root, []byte("a"))
	expError = fmt.Errorf("right proof, %w", expError)
	if err.Error() != expError.Error() {
		t.Fatal("attempting to prove non-existence of leaf a returned incorrect error")
	}
}

// generatePrefix generates a valid iavl prefix for an inner op.
func generateInnerOpPrefix() []byte {
	var (
		varintBuf  [binary.MaxVarintLen64]byte
		lengthByte byte = 0x20
	)
	height, size, version := 5, 10, 20
	prefix := convertVarIntToBytes(int64(height), varintBuf)
	prefix = append(prefix, convertVarIntToBytes(int64(size), varintBuf)...)
	prefix = append(prefix, convertVarIntToBytes(int64(version), varintBuf)...)
	prefix = append(prefix, lengthByte)
	return prefix
}

func convertVarIntToBytes(orig int64, buf [binary.MaxVarintLen64]byte) []byte {
	n := binary.PutVarint(buf[:], orig)
	return buf[:n]
}
