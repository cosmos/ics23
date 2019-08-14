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
		Proof    string `json:"existence"`
	}

	iavl := filepath.Join("..", "testdata", "iavl")
	tendermint := filepath.Join("..", "testdata", "tendermint")
	cases := []struct {
		dir      string
		filename string
		spec     *ProofSpec
	}{
		{dir: iavl, filename: "existence1.json", spec: IavlSpec},
		{dir: iavl, filename: "existence2.json", spec: IavlSpec},
		{dir: iavl, filename: "existence3.json", spec: IavlSpec},
		{dir: iavl, filename: "existence4.json", spec: IavlSpec},
		{dir: tendermint, filename: "existence1.json", spec: TendermintSpec},
		{dir: tendermint, filename: "existence2.json", spec: TendermintSpec},
		{dir: tendermint, filename: "existence3.json", spec: TendermintSpec},
		{dir: tendermint, filename: "existence4.json", spec: TendermintSpec},
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
			var proof ExistenceProof
			err = proof.Unmarshal(mustHex(t, data.Proof))
			if err != nil {
				t.Fatalf("Unmarshal protobuf: %+v", err)
			}

			root := CommitmentRoot(mustHex(t, data.RootHash))
			err = proof.Verify(tc.spec, root, proof.Key, proof.Value)
			if err != nil {
				t.Fatalf("Invalid proof: %+v", err)
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
