package ics23

import (
	"bytes"
	"fmt"
	"testing"
)

func TestVectors(t *testing.T) {
	cases := VectorsTestData()
	for _, tc := range cases {
		tc := tc
		name := fmt.Sprintf("%s/%s", tc.Dir, tc.Filename)
		t.Run(name, func(t *testing.T) {
			proof, ref := LoadFile(t, tc.Dir, tc.Filename)
			// Test Calculate method
			calculatedRoot, err := proof.Calculate()
			if err != nil {
				t.Fatal("proof.Calculate() returned error")
			}
			if !bytes.Equal(ref.RootHash, calculatedRoot) {
				t.Fatalf("calculated root: %X did not match expected root: %X", calculatedRoot, ref.RootHash)
			}
			// Test Verify method
			if ref.Value == nil {
				// non-existence
				valid := VerifyNonMembership(tc.Spec, ref.RootHash, proof, ref.Key)
				if !valid {
					t.Fatal("Invalid proof")
				}
			} else {
				valid := VerifyMembership(tc.Spec, ref.RootHash, proof, ref.Key, ref.Value)
				if !valid {
					t.Fatal("Invalid proof")
				}
			}
		})
	}
}

func TestDecompressBatchVectors(t *testing.T) {
	cases := DecompressBatchVectorsTestData(t)
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			small, err := tc.Marshal()
			if err != nil {
				t.Fatalf("Marshal batch %v", err)
			}

			decomp := Decompress(tc)
			if decomp == tc {
				t.Fatalf("Decompression is a no-op")
			}
			big, err := decomp.Marshal()
			if err != nil {
				t.Fatalf("Marshal batch %v", err)
			}
			if len(small) >= len(big) {
				t.Fatalf("Compression doesn't reduce size")
			}

			restore := Compress(tc)
			resmall, err := restore.Marshal()
			if err != nil {
				t.Fatalf("Marshal batch %v", err)
			}
			if len(resmall) != len(small) {
				t.Fatalf("Decompressed len %d, original len %d", len(resmall), len(small))
			}
			if !bytes.Equal(resmall, small) {
				t.Fatal("Decompressed batch proof differs from original")
			}
		})
	}
}
