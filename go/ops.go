package proofs

import (
	"crypto"
	// adds sha256 capability to crypto.SHA256
	_ "crypto/sha256"
	// adds sha512 capability to crypto.SHA512
	_ "crypto/sha512"
	fmt "fmt"
)

func ApplyOps(ops []*ProofOp, args ...[]byte) ([]byte, error) {
	first, rem := ops[0], ops[1:]
	res, err := ApplyOp(first, args...)
	if err != nil {
		return nil, err
	}
	for _, step := range rem {
		res, err = ApplyOp(step, res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func ApplyOp(op *ProofOp, args ...[]byte) ([]byte, error) {
	o := op.Op
	switch o.(type) {
	case *ProofOp_Leaf:
		if len(args) != 2 {
			return nil, fmt.Errorf("Need key and value args, got %d", len(args))
		}
		return ApplyLeafOp(op.GetLeaf(), args[0], args[1])
	case *ProofOp_Inner:
		if len(args) != 1 {
			return nil, fmt.Errorf("Need one child hash, got %d", len(args))
		}
		return ApplyInnerOp(op.GetInner(), args[0])
	default:
		panic("Unknown proof op")
	}
}

func ApplyLeafOp(op *LeafOp, key []byte, value []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Leaf node needs key")
	}
	if len(value) == 0 {
		return nil, fmt.Errorf("Leaf node needs value")
	}
	// TODO: lengthop before or after hash ???
	pkey, err := prepareLeafData(op.PrehashKey, op.Length, key)
	if err != nil {
		return nil, err
	}
	pvalue, err := prepareLeafData(op.PrehashValue, op.Length, value)
	if err != nil {
		return nil, err
	}
	data := append(op.Prefix, pkey...)
	data = append(data, pvalue...)
	fmt.Printf("data: %X\n", data)
	return doHash(op.Hash, data)
}

func prepareLeafData(hashOp HashOp, lengthOp LengthOp, data []byte) ([]byte, error) {
	hdata, err := doHashOrNoop(hashOp, data)
	if err != nil {
		return nil, err
	}
	ldata, err := doLengthOp(lengthOp, hdata)
	return ldata, err
}

func ApplyInnerOp(op *InnerOp, child []byte) ([]byte, error) {
	preimage := append(op.Prefix, child...)
	preimage = append(preimage, op.Suffix...)
	return doHash(op.Hash, preimage)
}

// doHashOrNoop will return the preimage untouched if hashOp == NONE,
// otherwise, perform doHash
func doHashOrNoop(hashOp HashOp, preimage []byte) ([]byte, error) {
	if hashOp == HashOp_NO_HASH {
		return preimage, nil
	}
	return doHash(hashOp, preimage)
}

// doHash will preform the specified hash on the preimage.
// if hashOp == NONE, it will return an error (use doHashOrNoop if you want different behavior)
func doHash(hashOp HashOp, preimage []byte) ([]byte, error) {
	switch hashOp {
	case HashOp_SHA256:
		hash := crypto.SHA256.New()
		hash.Write(preimage)
		return hash.Sum(nil), nil
	case HashOp_SHA512:
		hash := crypto.SHA512.New()
		hash.Write(preimage)
		return hash.Sum(nil), nil
	}
	return nil, fmt.Errorf("Unsupported hashop: %d", hashOp)
}

// doLengthOp will calculate the proper prefix and return it prepended
//   doLengthOp(op, data) -> length(data) || data
func doLengthOp(lengthOp LengthOp, data []byte) ([]byte, error) {
	switch lengthOp {
	case LengthOp_NO_PREFIX:
		return data, nil
	case LengthOp_VAR_PROTO:
		res := append(encodeVarintProto(len(data)), data...)
		return res, nil
	case LengthOp_REQUIRE_32_BYTES:
		if len(data) != 32 {
			return nil, fmt.Errorf("Data was %d bytes, not 32", len(data))
		}
		return data, nil
	case LengthOp_REQUIRE_64_BYTES:
		if len(data) != 64 {
			return nil, fmt.Errorf("Data was %d bytes, not 64", len(data))
		}
		return data, nil
		// TODO
		// case LengthOp_VAR_RLP:
		// case LengthOp_FIXED32_BIG:
		// case LengthOp_FIXED64_BIG:
		// case LengthOp_FIXED32_LITTLE:
		// case LengthOp_FIXED64_LITTLE:
	}
	return nil, fmt.Errorf("Unsupported lengthop: %d", lengthOp)
}

func encodeVarintProto(l int) []byte {
	// avoid multiple allocs for normal case
	res := make([]byte, 0, 8)
	for l >= 1<<7 {
		res = append(res, uint8(l&0x7f|0x80))
		l >>= 7
	}
	res = append(res, uint8(l))
	return res
}
