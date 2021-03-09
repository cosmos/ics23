package ics23

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	ics23 "github.com/confio/ics23/go"
	"github.com/stretchr/testify/require"
)

// TestVector is what is stored in the file
type TestVector struct {
	RootHash string `json:"root"`
	Proof    string `json:"proof"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}

// RefData is parsed version of everything except the CommitmentProof itself
type RefData struct {
	RootHash []byte
	Key      []byte
	Value    []byte
}

func TestVectors(t *testing.T) {

	iavl := filepath.Join("..", "testdata", "iavl")
	tendermint := filepath.Join("..", "testdata", "tendermint")
	cases := []struct {
		dir      string
		filename string
		spec     *ics23.ProofSpec
	}{
//		{dir: iavl, filename: "exist_left.json", spec: ics23.IavlSpec},
//		{dir: iavl, filename: "exist_right.json", spec: ics23.IavlSpec},
//		{dir: iavl, filename: "exist_middle.json", spec: ics23.IavlSpec},
		{dir: iavl, filename: "nonexist_left.json", spec: ics23.IavlSpec},
		{dir: iavl, filename: "nonexist_right.json", spec: ics23.IavlSpec},
		{dir: iavl, filename: "nonexist_middle.json", spec: ics23.IavlSpec},
		{dir: tendermint, filename: "exist_left.json", spec: ics23.TendermintSpec},
		{dir: tendermint, filename: "exist_right.json", spec: ics23.TendermintSpec},
		{dir: tendermint, filename: "exist_middle.json", spec: ics23.TendermintSpec},
		{dir: tendermint, filename: "nonexist_left.json", spec: ics23.TendermintSpec},
		{dir: tendermint, filename: "nonexist_right.json", spec: ics23.TendermintSpec},
		{dir: tendermint, filename: "nonexist_middle.json", spec: ics23.TendermintSpec},
	}

	session := Initialize(t)

	for _, tc := range cases {
		tc := tc
		name := fmt.Sprintf("%s/%s", tc.dir, tc.filename)
		t.Run(name, func(t *testing.T) {
			proof, ref := loadFile(t, tc.dir, tc.filename)
			// Test Calculate method
			calculatedRoot, err := proof.Calculate()
			if err != nil {
				t.Fatal("proof.Calculate() returned error")
			}
			if !bytes.Equal(ref.RootHash, calculatedRoot) {
				t.Fatalf("calculated root: %X did not match expected root: %X", calculatedRoot, ref.RootHash)
			}

			// Test Verify method
			if len(ref.Value) == 0 {
				// non-existence
				valid, err := session.VerifyNonMembership(ProofSpecToABI(tc.spec), ref.RootHash, NonExistenceProofToABI(proof.GetNonexist()), ref.Key)
				require.NoError(t, err)
				if !valid {
					t.Fatal("Invalid proof")
				}
			} else {
				valid, err := session.VerifyMembership(LeafOpToABI(tc.spec.LeafSpec), ref.RootHash, ExistenceProofToABI(proof.GetExist()), ref.Key, ref.Value)
				require.NoError(t, err)

				if !valid {
					t.Fatal("Invalid proof")
				}
			}
		})
	}
}

func loadFile(t *testing.T, dir string, filename string) (*ics23.CommitmentProof, *RefData) {
	// load the file into a json struct
	name := filepath.Join(dir, filename)
	bz, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf("Read file: %+v", err)
	}
	var data TestVector
	err = json.Unmarshal(bz, &data)
	if err != nil {
		t.Fatalf("Unmarshal json: %+v", err)
	}

	// parse the protobuf object
	var proof ics23.CommitmentProof
	err = proof.Unmarshal(mustHex(t, data.Proof))
	if err != nil {
		t.Fatalf("Unmarshal protobuf: %+v", err)
	}

	var ref RefData
	ref.RootHash = ics23.CommitmentRoot(mustHex(t, data.RootHash))
	ref.Key = mustHex(t, data.Key)
	if data.Value != "" {
		ref.Value = mustHex(t, data.Value)
	}

	return &proof, &ref
}

func mustHex(t *testing.T, data string) []byte {
	if data == "" {
		return nil
	}
	res, err := hex.DecodeString(data)
	if err != nil {
		t.Fatalf("decoding hex: %v", err)
	}
	return res
}
