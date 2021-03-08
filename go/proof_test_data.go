package ics23

type ExistenceProofTestStruct struct {
	Proof    *ExistenceProof
	IsErr    bool
	Expected []byte
}

func ExistenceProofTestData() map[string]ExistenceProofTestStruct {
	return map[string]ExistenceProofTestStruct{
		"must have at least one step": {
			Proof: &ExistenceProof{
				Key:   []byte("foo"),
				Value: []byte("bar"),
			},
			IsErr: true,
		},
		// copied from ops_test / TestLeafOp
		"executes one leaf step": {
			Proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("some longer text"),
				Leaf: &LeafOp{
					Hash:   HashOp_SHA256,
					Length: LengthOp_VAR_PROTO,
				},
			},
			Expected: fromHex("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"),
		},
		// iavl Leaf: start with 0, length 3
		// inner prefix: !start with 0, length >= 4
		"demonstrate maliability of leaf if we change leaf algorithm": {
			Proof: &ExistenceProof{
				Key:   append([]byte{4}, []byte("food")...),
				Value: append([]byte{16}, []byte("some longer text")...),
				Leaf: &LeafOp{
					Hash: HashOp_SHA256,
				},
			},
			Expected: fromHex("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"),
		},
		"demonstrate maliability of leaf if we change leaf prefix": {
			Proof: &ExistenceProof{
				Key:   append([]byte("od"), byte(16)),
				Value: []byte("some longer text"),
				Leaf: &LeafOp{
					Prefix: []byte{4, 'f', 'o'},
					Hash:   HashOp_SHA256,
				},
			},
			Expected: fromHex("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"),
		},
		"cannot execute inner first": {
			Proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("some longer text"),
				Path: []*InnerOp{
					&InnerOp{
						Hash:   HashOp_SHA256,
						Prefix: fromHex("deadbeef00cafe00"),
					},
				},
			},
			IsErr: true,
		},
		"executes leaf then inner op": {
			Proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("some longer text"),
				Leaf: &LeafOp{
					Hash:   HashOp_SHA256,
					Length: LengthOp_VAR_PROTO,
				},
				// output: b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265
				Path: []*InnerOp{
					&InnerOp{
						Hash:   HashOp_SHA256,
						Prefix: fromHex("deadbeef00cafe00"),
					},
					// echo -n deadbeef00cafe00b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265 | xxd -r -p | sha256sum
				},
			},
			Expected: fromHex("836ea236a6902a665c2a004c920364f24cad52ded20b1e4f22c3179bfe25b2a9"),
		},
	}
}

type CheckLeafTestStruct struct {
	Leaf  *LeafOp
	Spec  *LeafOp
	IsErr bool
}

func CheckLeafTestData() map[string]CheckLeafTestStruct {
	return map[string]CheckLeafTestStruct{
		"empty spec, empty leaf": {
			Leaf: &LeafOp{},
			Spec: &LeafOp{},
		},
		"empty spec allows prefix": {
			Leaf: &LeafOp{Prefix: fromHex("aabb")},
			Spec: &LeafOp{},
		},
		"empty spec doesn't allow hashop": {
			Leaf:  &LeafOp{Hash: HashOp_SHA256},
			Spec:  &LeafOp{},
			IsErr: true,
		},
		"spec with different prefixes": {
			Leaf:  &LeafOp{Prefix: fromHex("aabb")},
			Spec:  &LeafOp{Prefix: fromHex("bb")},
			IsErr: true,
		},
		"leaf with empty prefix (but spec has one)": {
			Leaf:  &LeafOp{},
			Spec:  &LeafOp{Prefix: fromHex("bb")},
			IsErr: true,
		},
		"leaf and spec match, all fields full": {
			Leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			Spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
		},
		"leaf and spec differ on hash": {
			Leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA256,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			Spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			IsErr: true,
		},
		"leaf and spec differ on length": {
			Leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_NO_PREFIX,
			},
			Spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			IsErr: true,
		},
		"leaf and spec differ on prehash key": {
			Leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_SHA256,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			Spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			IsErr: true,
		},
		"leaf and spec differ on prehash value": {
			Leaf: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_NO_HASH,
				Length:       LengthOp_VAR_PROTO,
			},
			Spec: &LeafOp{
				Prefix:       fromHex("00"),
				Hash:         HashOp_SHA512,
				PrehashKey:   HashOp_NO_HASH,
				PrehashValue: HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
			},
			IsErr: true,
		},
	}
}

type CheckAgainstSpecTestStruct struct {
	Proof *ExistenceProof
	Spec  *ProofSpec
	IsErr bool
}

func CheckAgainstSpecTestData() map[string]CheckAgainstSpecTestStruct {
	validInner := &InnerOp{
		Prefix: fromHex("aa"),
	}
	invalidInner := &InnerOp{
		Prefix: fromHex("00aa"),
		Suffix: fromHex("bb"),
	}

	return map[string]CheckAgainstSpecTestStruct{
		"empty proof fails": {
			Proof: &ExistenceProof{
				Key:   []byte("foo"),
				Value: []byte("bar"),
			},
			Spec:  IavlSpec,
			IsErr: true,
		},
		"accepts one proper leaf": {
			Proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
			},
			Spec: IavlSpec,
		},
		"rejects invalid leaf": {
			Proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf: &LeafOp{
					Prefix: []byte{0},
					Hash:   HashOp_SHA256,
					Length: LengthOp_VAR_PROTO,
				},
			},
			Spec:  IavlSpec,
			IsErr: true,
		},
		"rejects only inner proof": {
			Proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Path: []*InnerOp{
					validInner,
				},
			},
			Spec:  IavlSpec,
			IsErr: true,
		},
		"accepts leaf with valid inner proofs": {
			Proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
				Path: []*InnerOp{
					validInner,
					validInner,
				},
			},
			Spec: IavlSpec,
		},
		"rejects leaf with invalid inner proofs": {
			Proof: &ExistenceProof{
				Key:   []byte("food"),
				Value: []byte("bar"),
				Leaf:  IavlSpec.LeafSpec,
				Path: []*InnerOp{
					validInner,
					invalidInner,
					validInner,
				},
			},
			Spec:  IavlSpec,
			IsErr: true,
		},
	}

}
