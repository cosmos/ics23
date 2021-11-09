package ics23

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
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

func TestVectorsData() []TestVectorsStruct {
	iavl := filepath.Join("..", "testdata", "iavl")
	tendermint := filepath.Join("..", "testdata", "tendermint")
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

func TestBatchVectorsData(t *testing.T) map[string]BatchVectorData {
	iavl := filepath.Join("..", "testdata", "iavl")
	tendermint := filepath.Join("..", "testdata", "tendermint")
	// Note that each item has a different commitment root,
	// so maybe not ideal (cannot check multiple entries)
	batch_iavl, refs_iavl := buildBatch(t, iavl, []string{
		"exist_left.json",
		"exist_right.json",
		"exist_middle.json",
		"nonexist_left.json",
		"nonexist_right.json",
		"nonexist_middle.json",
	})
	batch_tm, refs_tm := buildBatch(t, tendermint, []string{
		"exist_left.json",
		"exist_right.json",
		"exist_middle.json",
		"nonexist_left.json",
		"nonexist_right.json",
		"nonexist_middle.json",
	})
	batch_tm_exist, refs_tm_exist := loadBatch(t, tendermint, "batch_exist.json")
	batch_tm_nonexist, refs_tm_nonexist := loadBatch(t, tendermint, "batch_nonexist.json")
	batch_iavl_exist, refs_iavl_exist := loadBatch(t, iavl, "batch_exist.json")
	batch_iavl_nonexist, refs_iavl_nonexist := loadBatch(t, iavl, "batch_nonexist.json")
	return map[string]BatchVectorData{
		"iavl 0": {Spec: IavlSpec, Proof: batch_iavl, Ref: refs_iavl[0]},
		"iavl 1": {Spec: IavlSpec, Proof: batch_iavl, Ref: refs_iavl[1]},
		"iavl 2": {Spec: IavlSpec, Proof: batch_iavl, Ref: refs_iavl[2]},
		"iavl 3": {Spec: IavlSpec, Proof: batch_iavl, Ref: refs_iavl[3]},
		"iavl 4": {Spec: IavlSpec, Proof: batch_iavl, Ref: refs_iavl[4]},
		"iavl 5": {Spec: IavlSpec, Proof: batch_iavl, Ref: refs_iavl[5]},
		// Note this spec only differs for non-existence proofs
		"iavl invalid 1":      {Spec: TendermintSpec, Proof: batch_iavl, Ref: refs_iavl[4], Invalid: true},
		"iavl invalid 2":      {Spec: IavlSpec, Proof: batch_iavl, Ref: refs_tm[0], Invalid: true},
		"iavl batch exist":    {Spec: IavlSpec, Proof: batch_iavl_exist, Ref: refs_iavl_exist[17]},
		"iavl batch nonexist": {Spec: IavlSpec, Proof: batch_iavl_nonexist, Ref: refs_iavl_nonexist[7]},
		"tm 0":                {Spec: TendermintSpec, Proof: batch_tm, Ref: refs_tm[0]},
		"tm 1":                {Spec: TendermintSpec, Proof: batch_tm, Ref: refs_tm[1]},
		"tm 2":                {Spec: TendermintSpec, Proof: batch_tm, Ref: refs_tm[2]},
		"tm 3":                {Spec: TendermintSpec, Proof: batch_tm, Ref: refs_tm[3]},
		"tm 4":                {Spec: TendermintSpec, Proof: batch_tm, Ref: refs_tm[4]},
		"tm 5":                {Spec: TendermintSpec, Proof: batch_tm, Ref: refs_tm[5]},
		// Note this spec only differs for non-existence proofs
		"tm invalid 1":      {Spec: IavlSpec, Proof: batch_tm, Ref: refs_tm[4], Invalid: true},
		"tm invalid 2":      {Spec: TendermintSpec, Proof: batch_tm, Ref: refs_iavl[0], Invalid: true},
		"tm batch exist":    {Spec: TendermintSpec, Proof: batch_tm_exist, Ref: refs_tm_exist[10]},
		"tm batch nonexist": {Spec: TendermintSpec, Proof: batch_tm_nonexist, Ref: refs_tm_nonexist[3]},
	}
}

func TestDecompressBatchVectorsData(t *testing.T) map[string]*CommitmentProof {
	iavl := filepath.Join("..", "testdata", "iavl")
	tendermint := filepath.Join("..", "testdata", "tendermint")
	// note that these batches are already compressed
	batch_iavl, _ := loadBatch(t, iavl, "batch_exist.json")
	batch_tm, _ := loadBatch(t, tendermint, "batch_nonexist.json")
	return map[string]*CommitmentProof{
		"iavl":       batch_iavl,
		"tendermint": batch_tm,
	}
}

func LoadFile(t *testing.T, dir string, filename string) (*CommitmentProof, *RefData) {
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
	var proof CommitmentProof
	err = proof.Unmarshal(mustHex(t, data.Proof))
	if err != nil {
		t.Fatalf("Unmarshal protobuf: %+v", err)
	}
	var ref RefData
	ref.RootHash = CommitmentRoot(mustHex(t, data.RootHash))
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

func buildBatch(t *testing.T, dir string, filenames []string) (*CommitmentProof, []*RefData) {
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
	// load the file into a json struct
	name := filepath.Join(dir, filename)
	bz, err := ioutil.ReadFile(name)
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
	var refs = make([]*RefData, len(data.Items))
	for i, item := range data.Items {
		refs[i] = &RefData{
			RootHash: root,
			Key:      mustHex(t, item.Key),
			Value:    mustHex(t, item.Value),
		}
	}
	return &proof, refs
}
