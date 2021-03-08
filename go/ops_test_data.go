package ics23

import (
	"encoding/hex"
)

type LeafOpTestStruct struct {
	Op       *LeafOp
	Key      []byte
	Value    []byte
	IsErr    bool
	Expected []byte
}

func LeafOpTestData() map[string]LeafOpTestStruct {
	return map[string]LeafOpTestStruct{
		"hash foobar": {
			Op: &LeafOp{
				Hash: HashOp_SHA256,
				// no prehash, no length prefix
			},
			Key:   []byte("foo"),
			Value: []byte("bar"),
			// echo -n foobar | sha256sum
			Expected: fromHex("c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"),
		},
		"hash foobaz, sha-512": {
			Op: &LeafOp{
				Hash: HashOp_SHA512,
				// no prehash, no length prefix
			},
			Key:   []byte("foo"),
			Value: []byte("baz"),
			// echo -n foobaz | sha512sum
			Expected: fromHex("4f79f191298ec7461d60136c60f77c2ae8ddd85dbf6168bb925092d51bfb39b559219b39ae5385ba04946c87f64741385bef90578ea6fe6dac85dbf7ad3f79e1"),
		},
		"hash foobar (different break)": {
			Op: &LeafOp{
				Hash: HashOp_SHA256,
				// no prehash, no length prefix
			},
			Key:   []byte("f"),
			Value: []byte("oobar"),
			// echo -n foobar | sha256sum
			Expected: fromHex("c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"),
		},
		"requires key": {
			Op: &LeafOp{
				Hash: HashOp_SHA256,
				// no prehash, no length prefix
			},
			Key:   []byte("foo"),
			IsErr: true,
		},
		"requires value": {
			Op: &LeafOp{
				Hash: HashOp_SHA256,
				// no prehash, no length prefix
			},
			Value: []byte("bar"),
			IsErr: true,
		},
		"hash with length prefix": {
			Op: &LeafOp{
				Hash:   HashOp_SHA256,
				Length: LengthOp_VAR_PROTO,
				// no prehash
			},
			// echo -n food | xxs -ps
			// and manually compute length byte
			Key:   []byte("food"),             // 04666f6f64
			Value: []byte("some longer text"), // 10736f6d65206c6f6e6765722074657874
			// echo -n 04666f6f6410736f6d65206c6f6e6765722074657874 | xxd -r -p | sha256sum
			Expected: fromHex("b68f5d298e915ae1753dd333da1f9cf605411a5f2e12516be6758f365e6db265"),
		},
		"hash with prehash and length prefix": {
			Op: &LeafOp{
				Hash:         HashOp_SHA256,
				Length:       LengthOp_VAR_PROTO,
				PrehashValue: HashOp_SHA256,
			},
			Key: []byte("food"), // 04666f6f64
			// TODO: this is hash, then length....
			// echo -n yet another long string | sha256sum
			Value: []byte("yet another long string"), // 20a48c2d4f67b9f80374938535285ed285819d8a5a8fc1fccd1e3244e437cf290d
			// echo -n 04666f6f6420a48c2d4f67b9f80374938535285ed285819d8a5a8fc1fccd1e3244e437cf290d | xxd -r -p | sha256sum
			Expected: fromHex("87e0483e8fb624aef2e2f7b13f4166cda485baa8e39f437c83d74c94bedb148f"),
		},
	}
}

type InnerOpTestStruct struct {
	Op       *InnerOp
	Child    []byte
	IsErr    bool
	Expected []byte
}

func InnerOpTestData() map[string]InnerOpTestStruct {
	return map[string]InnerOpTestStruct{
		"requires child": {
			Op: &InnerOp{
				Hash:   HashOp_SHA256,
				Prefix: fromHex("0123456789"),
				Suffix: fromHex("deadbeef"),
			},
			IsErr: true,
		},
		"hash child with prefix and suffix": {
			Op: &InnerOp{
				Hash:   HashOp_SHA256,
				Prefix: fromHex("0123456789"),
				Suffix: fromHex("deadbeef"),
			},
			Child: fromHex("00cafe00"),
			// echo -n 012345678900cafe00deadbeef | xxd -r -p | sha256sum
			Expected: fromHex("0339f76086684506a6d42a60da4b5a719febd4d96d8b8d85ae92849e3a849a5e"),
		},
		"hash child with only prefix": {
			Op: &InnerOp{
				Hash:   HashOp_SHA256,
				Prefix: fromHex("00204080a0c0e0"),
			},
			Child: fromHex("ffccbb997755331100"),
			// echo -n 00204080a0c0e0ffccbb997755331100 | xxd -r -p | sha256sum
			Expected: fromHex("45bece1678cf2e9f4f2ae033e546fc35a2081b2415edcb13121a0e908dca1927"),
		},
		"hash child with only suffix": {
			Op: &InnerOp{
				Hash:   HashOp_SHA256,
				Suffix: []byte(" just kidding!"),
			},
			Child: []byte("this is a sha256 hash, really...."),
			// echo -n 'this is a sha256 hash, really.... just kidding!'  | sha256sum
			Expected: fromHex("79ef671d27e42a53fba2201c1bbc529a099af578ee8a38df140795db0ae2184b"),
		},
	}
}

type DoHashTestStruct struct {
	HashOp       HashOp
	Preimage     string
	ExpectedHash string
}

func DoHashTestData() map[string]DoHashTestStruct {
	return map[string]DoHashTestStruct{
		"sha256": {
			HashOp:   HashOp_SHA256,
			Preimage: "food",
			// echo -n food | sha256sum
			ExpectedHash: "c1f026582fe6e8cb620d0c85a72fe421ddded756662a8ec00ed4c297ad10676b",
		},
		"ripemd160": {
			HashOp:   HashOp_RIPEMD160,
			Preimage: "food",
			// echo -n food | openssl dgst -rmd160 -hex | cut -d' ' -f2
			ExpectedHash: "b1ab9988c7c7c5ec4b2b291adfeeee10e77cdd46",
		},
		"bitcoin": {
			HashOp:   HashOp_BITCOIN,
			Preimage: "food",
			// echo -n c1f026582fe6e8cb620d0c85a72fe421ddded756662a8ec00ed4c297ad10676b | xxd -r -p | openssl dgst -rmd160 -hex
			ExpectedHash: "0bcb587dfb4fc10b36d57f2bba1878f139b75d24",
		},
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
