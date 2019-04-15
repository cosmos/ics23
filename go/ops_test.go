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
		"hash foobaz, sha-512": {
			op: &LeafOp{
				Hash: HashOp_SHA512,
				// no prehash, no length prefix
			},
			key:   []byte("foo"),
			value: []byte("baz"),
			// echo -n foobaz | sha512sum
			expected: fromHex("4f79f191298ec7461d60136c60f77c2ae8ddd85dbf6168bb925092d51bfb39b559219b39ae5385ba04946c87f64741385bef90578ea6fe6dac85dbf7ad3f79e1"),
		},
		"hash foobar (different break)": {
			op: &LeafOp{
				Hash: HashOp_SHA256,
				// no prehash, no length prefix
			},
			key:   []byte("f"),
			value: []byte("oobar"),
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
		"hash with length prefix": {
			op: &LeafOp{
				Hash:   HashOp_SHA256,
				Length: LengthOp_VAR_PROTO,
				// no prehash, no length prefix
			},
			// echo -n food | xxs -ps
			// and manually compute length byte
			key:   []byte("food"),             // 04666f6f64
			value: []byte("some longer text"), // 10736f6d65206c6f6e6765722074657874
			// echo -n 04666f6f6410736f6d65206c6f6e6765722074657874 | xxd -r -p | sha256sum
			expected: fromHex("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"),
		},
		"hash with prehash and length prefix": {
			op: &LeafOp{
				Hash:         HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
				PrehashValue: HashOp_SHA256,
				// no prehash, no length prefix
			},
			key: []byte("food"), // 04666f6f64
			// TODO: this is hash, then length....
			// echo -n yet another long string | sha256sum
			value: []byte("yet another long string"), // 20a48c2d4f67b9f80374938535285ed285819d8a5a8fc1fccd1e3244e437cf290d
			// echo -n 04666f6f6420a48c2d4f67b9f80374938535285ed285819d8a5a8fc1fccd1e3244e437cf290d | xxd -r -p | sha256sum
			expected: fromHex("87e0483e8fb624aef2e2f7b13f4166cda485baa8e39f437c83d74c94bedb148f"),
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
