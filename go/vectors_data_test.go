package ics23

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
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

type TestVectorsStruct struct {
	Dir      string
	Filename string
	Spec     *ProofSpec
}

func VectorsTestData() []TestVectorsStruct {
	iavl := filepath.Join("..", "testdata", "iavl")
	tendermint := filepath.Join("..", "testdata", "tendermint")
	smt := filepath.Join("..", "testdata", "smt")
	cases := []TestVectorsStruct{
		{Dir: iavl, Filename: "exist_left.json", Spec: IavlSpec},
		{Dir: iavl, Filename: "exist_right.json", Spec: IavlSpec},
		{Dir: iavl, Filename: "exist_middle.json", Spec: IavlSpec},
		{Dir: iavl, Filename: "nonexist_left.json", Spec: IavlSpec},
		{Dir: iavl, Filename: "nonexist_right.json", Spec: IavlSpec},
		{Dir: iavl, Filename: "nonexist_middle.json", Spec: IavlSpec},
		{Dir: tendermint, Filename: "exist_left.json", Spec: TendermintSpec},
		{Dir: tendermint, Filename: "exist_right.json", Spec: TendermintSpec},
		{Dir: tendermint, Filename: "exist_middle.json", Spec: TendermintSpec},
		{Dir: tendermint, Filename: "nonexist_left.json", Spec: TendermintSpec},
		{Dir: tendermint, Filename: "nonexist_right.json", Spec: TendermintSpec},
		{Dir: tendermint, Filename: "nonexist_middle.json", Spec: TendermintSpec},
		{Dir: smt, Filename: "exist_left.json", Spec: SmtSpec},
		{Dir: smt, Filename: "exist_right.json", Spec: SmtSpec},
		{Dir: smt, Filename: "exist_middle.json", Spec: SmtSpec},
		{Dir: smt, Filename: "nonexist_left.json", Spec: SmtSpec},
		{Dir: smt, Filename: "nonexist_right.json", Spec: SmtSpec},
		{Dir: smt, Filename: "nonexist_middle.json", Spec: SmtSpec},
	}
	return cases
}

// BatchVector is what is stored in the file
type BatchVector struct {
	RootHash string `json:"root"`
	Proof    string `json:"proof"`
	Items    []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
}

type BatchVectorData struct {
	Spec    *ProofSpec
	Proof   *CommitmentProof
	Ref     *RefData
	Invalid bool // default is valid
}

func LoadFile(tb testing.TB, dir string, filename string) (*CommitmentProof, *RefData) {
	tb.Helper()
	// load the file into a json struct
	name := filepath.Join(dir, filename)
	bz, err := os.ReadFile(name)
	if err != nil {
		tb.Fatalf("Read file: %+v", err)
	}
	var data TestVector
	err = json.Unmarshal(bz, &data)
	if err != nil {
		tb.Fatalf("Unmarshal json: %+v", err)
	}
	// parse the protobuf object
	var proof CommitmentProof
	err = proof.Unmarshal(mustHex(tb, data.Proof))
	if err != nil {
		tb.Fatalf("Unmarshal protobuf: %+v", err)
	}
	var ref RefData
	ref.RootHash = CommitmentRoot(mustHex(tb, data.RootHash))
	ref.Key = mustHex(tb, data.Key)
	if data.Value != "" {
		ref.Value = mustHex(tb, data.Value)
	}
	return &proof, &ref
}

func mustHex(tb testing.TB, data string) []byte {
	tb.Helper()
	if data == "" {
		return nil
	}
	res, err := hex.DecodeString(data)
	if err != nil {
		tb.Fatalf("decoding hex: %v", err)
	}
	return res
}
