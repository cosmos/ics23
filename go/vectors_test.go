package proofs

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

// TestVector is what is stored in the file
type TestVector struct {
	RootHash string `json:"root"`
	Proof    string `json:"proof"`
	Key    string `json:"key"`
	Value    string `json:"value"`
}

// RefData is parsed version of everything except the CommitmentProof itself
type RefData struct{
	RootHash []byte
	Key []byte
	Value []byte
}

func TestVectors(t *testing.T) {

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
			proof, ref := loadFile(t, tc.dir, tc.filename)
			if ref.Value == nil {
				// non-existence
				valid := VerifyNonMembership(tc.spec, ref.RootHash, proof, ref.Key)
				if !valid {
					t.Fatal("Invalid proof")
				}
			} else {
				valid := VerifyMembership(tc.spec, ref.RootHash, proof, ref.Key, ref.Value)
				if !valid {
					t.Fatal("Invalid proof")
				}
			}
		})
	}
}

func loadFile(t *testing.T, dir string, filename string) (*CommitmentProof, *RefData) {
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

func loadBatch(t *testing.T, dir string, filenames []string) (*CommitmentProof, []*RefData) {
	var proof *CommitmentProof
	refs := make([]*RefData, len(filenames))
	entries := make([]*BatchEntry, len(filenames))

	for i, fn := range filenames {
		proof, refs[i] = loadFile(t, dir, fn)
		if ex := proof.GetExist(); ex != nil {
			entries[i] = &BatchEntry{
				Proof: &BatchEntry_Exist{
					Exist: ex,
				},
			}
		} else if non := proof.GetNonexist(); non != nil {
			entries[i] = &BatchEntry{
				Proof: &BatchEntry_Nonexist{
					Nonexist: non,
				},
			}
		} else {
			t.Fatalf("Loaded proof neither exist or nonexist: %s\n%#v", fn, proof.GetProof())
		}
	}

	result := &CommitmentProof{
		Proof: &CommitmentProof_Batch{
			Batch: &BatchProof{
				Entries: entries,
			},
		},
	}

	return result, refs
}


func TestBatchVectors(t *testing.T) {
	iavl := filepath.Join("..", "testdata", "iavl")
	tendermint := filepath.Join("..", "testdata", "tendermint")

	// Note that each item has a different commitment root,
	// so maybe not ideal (cannot check multiple entries)
	batch_iavl, refs_iavl := loadBatch(t, iavl, []string{
		"exist_left.json",
		"exist_right.json",
		"exist_middle.json",
		"nonexist_left.json",
		"nonexist_right.json",
		"nonexist_middle.json",
	})

	batch_tm, refs_tm := loadBatch(t, tendermint, []string{
		"exist_left.json",
		"exist_right.json",
		"exist_middle.json",
		"nonexist_left.json",
		"nonexist_right.json",
		"nonexist_middle.json",
	})


	cases := map[string]struct {
		spec     *ProofSpec
		proof 	 *CommitmentProof
		ref 	*RefData
		invalid bool // default is valid
	}{
		"iavl 0": {spec: IavlSpec, proof: batch_iavl, ref: refs_iavl[0]},
		"iavl 1": {spec: IavlSpec, proof: batch_iavl, ref: refs_iavl[1]},
		"iavl 2": {spec: IavlSpec, proof: batch_iavl, ref: refs_iavl[2]},
		"iavl 3": {spec: IavlSpec, proof: batch_iavl, ref: refs_iavl[3]},
		"iavl 4": {spec: IavlSpec, proof: batch_iavl, ref: refs_iavl[4]},
		"iavl 5": {spec: IavlSpec, proof: batch_iavl, ref: refs_iavl[5]},
		// Note this spec only differs for non-existence proofs
		"iavl invalid 1": {spec: TendermintSpec, proof: batch_iavl, ref: refs_iavl[4], invalid: true},
		"iavl invalid 2": {spec: IavlSpec, proof: batch_iavl, ref: refs_tm[0], invalid: true},
		"tm 0": {spec: TendermintSpec, proof: batch_tm, ref: refs_tm[0]},
		"tm 1": {spec: TendermintSpec, proof: batch_tm, ref: refs_tm[1]},
		"tm 2": {spec: TendermintSpec, proof: batch_tm, ref: refs_tm[2]},
		"tm 3": {spec: TendermintSpec, proof: batch_tm, ref: refs_tm[3]},
		"tm 4": {spec: TendermintSpec, proof: batch_tm, ref: refs_tm[4]},
		"tm 5": {spec: TendermintSpec, proof: batch_tm, ref: refs_tm[5]},
		// Note this spec only differs for non-existence proofs
		"tm invalid 1": {spec: IavlSpec, proof: batch_tm, ref: refs_tm[4], invalid: true},
		"tm invalid 2": {spec: TendermintSpec, proof: batch_tm, ref: refs_iavl[0], invalid: true},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// try one proof
				if tc.ref.Value == nil {
					// non-existence
					valid := VerifyNonMembership(tc.spec, tc.ref.RootHash, tc.proof, tc.ref.Key)
					if valid == tc.invalid {
						t.Fatalf("Expected proof validity: %t", !tc.invalid)
					}
					keys := [][]byte{tc.ref.Key}
					valid = BatchVerifyNonMembership(tc.spec, tc.ref.RootHash, tc.proof, keys)
					if valid == tc.invalid {
						t.Fatalf("Expected batch proof validity: %t", !tc.invalid)
					}
				} else {
					valid := VerifyMembership(tc.spec, tc.ref.RootHash, tc.proof, tc.ref.Key, tc.ref.Value)
					if valid == tc.invalid {
						t.Fatalf("Expected proof validity: %t", !tc.invalid)
					}
					items := make(map[string][]byte)
					items[string(tc.ref.Key)] = tc.ref.Value
					valid = BatchVerifyMembership(tc.spec, tc.ref.RootHash, tc.proof, items)
					if valid == tc.invalid {
						t.Fatalf("Expected batch proof validity: %t", !tc.invalid)
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
