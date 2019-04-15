package proofs

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestLeafOp(t *testing.T) {
	cases := map[string]struct {
		op       *LeafOp
		key      []byte
		value    []byte
		isErr    bool
		expected []byte
	}{
		"hash foobar": {
			op: &LeafOp{
				Hash: HashOp_SHA256,
				// no prehash, no length prefix
			},
			key:   []byte("foo"),
			value: []byte("bar"),
			// echo -n foobar | sha256sum
			expected: fromHex("c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"),
		},
		"requires key": {
			op: &LeafOp{
				Hash: HashOp_SHA256,
				// no prehash, no length prefix
			},
			key:   []byte("foo"),
			isErr: true,
		},
		"requires value": {
			op: &LeafOp{
				Hash: HashOp_SHA256,
				// no prehash, no length prefix
			},
			value: []byte("bar"),
			isErr: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := ApplyLeafOp(tc.op, tc.key, tc.value)
			// short-circuit with error case
			if tc.isErr {
				if err == nil {
					t.Fatal("Expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(res, tc.expected) {
				t.Errorf("Bad result: %s vs %s", toHex(res), toHex(tc.expected))
			}
		})
	}
}

func fromHex(data string) []byte {
	res, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return res
}

func toHex(data []byte) string {
	return hex.EncodeToString(data)
}
