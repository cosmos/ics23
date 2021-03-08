package ics23

import (
	"testing"

	ics23 "github.com/confio/ics23/go"
	"github.com/stretchr/testify/require"
)

func TestExistenceProof(t *testing.T) {
	session := Initialize(t)

	cases := ics23.ExistenceProofTestData()

	for name, tc := range cases {
		if tc.Proof.Leaf == nil {
			delete(cases, name)
		}
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := session.Calculate(ExistenceProofToABI(tc.Proof))
			if tc.IsErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, res, tc.Expected)
			}
		})
	}
}

func TestCheckAgainstSpec(t *testing.T) {
	session := Initialize(t)

	cases := ics23.CheckAgainstSpecTestData()

	for name, tc := range cases {
		if tc.Proof.Leaf == nil {
			delete(cases, name)
		}
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			ok, err := session.CheckAgainstSpec(ExistenceProofToABI(tc.Proof), LeafOpToABI(tc.Spec.LeafSpec))
			require.NoError(t, err)
			if tc.IsErr {
				require.False(t, ok)
			} else {
				require.True(t, ok)
			}
		})
	}
}
