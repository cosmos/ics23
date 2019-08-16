package proofs

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestIavlVectors(t *testing.T) {
	type TestData struct {
		RootHash string `json:"root"`
		Proof    string `json:"proof"`
		Key    string `json:"key"`
		Value    string `json:"value"`
	}

	iavl := filepath.Join("..", "testdata", "iavl")
	tendermint := filepath.Join("..", "testdata", "tendermint")
	cases := []struct {
		dir      string
		filename string
		spec     *ProofSpec
	}{
		{dir: iavl, filename: "exist_left.json", spec: IavlSpec},
		{dir: iavl, filename: "exist_right.json", spec: IavlSpec},
		{dir: iavl, filename: "exist_middle.json", spec: IavlSpec},
		{dir: iavl, filename: "nonexist_left.json", spec: IavlSpec},
		{dir: iavl, filename: "nonexist_right.json", spec: IavlSpec},
		{dir: iavl, filename: "nonexist_middle.json", spec: IavlSpec},
		{dir: tendermint, filename: "exist_left.json", spec: TendermintSpec},
		{dir: tendermint, filename: "exist_right.json", spec: TendermintSpec},
		{dir: tendermint, filename: "exist_middle.json", spec: TendermintSpec},
		{dir: tendermint, filename: "nonexist_left.json", spec: TendermintSpec},
		{dir: tendermint, filename: "nonexist_right.json", spec: TendermintSpec},
		{dir: tendermint, filename: "nonexist_middle.json", spec: TendermintSpec},
	}

	for _, tc := range cases {
		name := fmt.Sprintf("%s/%s", tc.dir, tc.filename)
		t.Run(name, func(t *testing.T) {
			// load the file into a json struct
			name := filepath.Join(tc.dir, tc.filename)
			bz, err := ioutil.ReadFile(name)
			if err != nil {
				t.Fatalf("Read file: %+v", err)
			}
			var data TestData
			err = json.Unmarshal(bz, &data)
			if err != nil {
				t.Fatalf("Unmarshal json: %+v", err)
			}

			// parse the protobuf object
			var proof CommitmentProof
			err = proof.Unmarshal(mustHex(t, data.Proof))
			if err != nil {
				t.Fatalf("Unmarshal protobuf: %+v", err)
			}

			root := CommitmentRoot(mustHex(t, data.RootHash))
			key := mustHex(t, data.Key)

			if data.Value == "" {
				// non-existence
				valid := VerifyNonMembership(tc.spec, root, &proof, key)
				if !valid {
					t.Fatal("Invalid proof")
				}
			} else {
				value := mustHex(t, data.Value)
				valid := VerifyMembership(tc.spec, root, &proof, key, value)
				if !valid {
					t.Fatal("Invalid proof")
				}
			}
		})
	}
}

func mustHex(t *testing.T, data string) []byte {
	res, err := hex.DecodeString(data)
	if err != nil {
		t.Fatalf("decoding hex: %v", err)
	}
	return res
}
