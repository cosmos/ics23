package proofs

import (
	"bytes"
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
				t.Fatalf("Read file: %s", err)
			}
			var data TestData
			err = json.Unmarshal(bz, &data)
			if err != nil {
				t.Fatalf("Unmarshal json: %s", err)
			}

			// parse the protobuf object
			var proof ExistenceProof
			err = proof.Unmarshal(mustHex(t, data.Proof))
			if err != nil {
				t.Fatalf("Unmarshal protobuf: %s", err)
			}

			// ensure the proof is valid
			err = proof.CheckAgainstSpec(tc.spec)
			if err != nil {
				t.Fatalf("Failed Iavl check spec: %s", err)
			}

			calc, err := proof.Calculate()
			if err != nil {
				t.Fatalf("Calculating root hash: %s", err)
			}

			root := mustHex(t, data.RootHash)

			if !bytes.Equal(calc, root) {
				t.Errorf("Expected root %X calculated %X", root, calc)
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
