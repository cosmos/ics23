package proofs

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestIavlVectors(t *testing.T) {
	type TestData struct {
		RootHash string `json:"root"`
		Proof    string `json:"existence"`
	}

	cases := []struct {
		filename string
	}{
		// TODO: re-generate with new format
		// {filename: "existence1.json"},
		// {filename: "existence2.json"},
		// {filename: "existence3.json"},
		// {filename: "existence4.json"},
	}
	dir := filepath.Join("..", "testdata", "iavl")

	for _, tc := range cases {
		t.Run(tc.filename, func(t *testing.T) {
			// load the file into a json struct
			name := filepath.Join(dir, tc.filename)
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
			err = proof.CheckAgainstSpec(IavlSpec)
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
