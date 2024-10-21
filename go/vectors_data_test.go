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

func DecompressBatchVectorsTestData(t *testing.T) map[string]*CommitmentProof {
	t.Helper()
	iavl := filepath.Join("..", "testdata", "iavl")
	tendermint := filepath.Join("..", "testdata", "tendermint")
	smt := filepath.Join("..", "testdata", "smt")
	// note that these batches are already compressed
	batchIAVL, _ := loadBatch(t, iavl, "batch_exist.json")
	batchTM, _ := loadBatch(t, tendermint, "batch_nonexist.json")
	batchSMT, _ := loadBatch(t, smt, "batch_nonexist.json")
	return map[string]*CommitmentProof{
		"iavl":       batchIAVL,
		"tendermint": batchTM,
		"smt":        batchSMT,
	}
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

func buildBatch(t *testing.T, dir string, filenames []string) (*CommitmentProof, []*RefData) {
	t.Helper()
	refs := make([]*RefData, len(filenames))
	proofs := make([]*CommitmentProof, len(filenames))
	for i, fn := range filenames {
		proofs[i], refs[i] = LoadFile(t, dir, fn)
	}
	batch, err := CombineProofs(proofs)
	if err != nil {
		t.Fatalf("Generating batch: %v", err)
	}
	return batch, refs
}

func loadBatch(t *testing.T, dir string, filename string) (*CommitmentProof, []*RefData) {
	t.Helper()
	// load the file into a json struct
	name := filepath.Join(dir, filename)
	bz, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("Read file: %+v", err)
	}
	var data BatchVector
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
	root := mustHex(t, data.RootHash)
	refs := make([]*RefData, len(data.Items))
	for i, item := range data.Items {
		refs[i] = &RefData{
			RootHash: root,
			Key:      mustHex(t, item.Key),
			Value:    mustHex(t, item.Value),
		}
	}
	return &proof, refs
}
