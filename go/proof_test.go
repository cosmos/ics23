package ics23

import (
	"bytes"
	"fmt"
	"testing"
)

func TestExistenceProof(t *testing.T) {
	cases := ExistenceProofTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := tc.Proof.Calculate()
			// short-circuit with error case
			if tc.IsErr {
				if err == nil {
					t.Fatal("Expected error, but got none")
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
			}
			if !bytes.Equal(res, tc.Expected) {
				t.Errorf("Bad result: %s vs %s", toHex(res), toHex(tc.Expected))
			}
		})
	}
}

func TestCheckLeaf(t *testing.T) {
	cases := CheckLeafTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.Leaf.CheckAgainstSpec(&ProofSpec{LeafSpec: tc.Spec})
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got nil")
			} else if !tc.IsErr && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}

func TestCheckAgainstSpec(t *testing.T) {
	cases := CheckAgainstSpecTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.Proof.CheckAgainstSpec(tc.Spec)
			if tc.Err == "" && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			} else if tc.Err != "" && tc.Err != err.Error() {
				t.Fatalf("Expected error: %s, got %s", tc.Err, err.Error())
			}
		})
	}
}

func TestEmptyBranch(t *testing.T) {
	cases := EmptyBranchTestData(t)

	for _, tc := range cases {
		t.Run("case", func(t *testing.T) {
			if err := tc.Op.CheckAgainstSpec(tc.Spec, 1); err != nil {
				t.Errorf("Invalid InnerOp: %v", err)
			}
			leftBranchesAreEmpty, err := leftBranchesAreEmpty(tc.Spec.InnerSpec, tc.Op)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if leftBranchesAreEmpty != tc.IsLeft {
				t.Errorf("Expected leftBranchesAreEmpty to be %t but it wasn't", tc.IsLeft)
			}
			rightBranchesAreEmpty, err := rightBranchesAreEmpty(tc.Spec.InnerSpec, tc.Op)
			if err != nil {
				t.Errorf("Error: %v", err)
			}
			if rightBranchesAreEmpty != tc.IsRight {
				t.Errorf("Expected rightBranchesAreEmpty to be %t but it wasn't", tc.IsRight)
			}
		})
	}
}

func TestGetPosition(t *testing.T) {
	cases := []struct {
		name     string
		order    []int32
		branch   int32
		expPos   int
		expError error
	}{
		{
			name:     "success: first branch",
			order:    IavlSpec.InnerSpec.ChildOrder,
			branch:   0,
			expPos:   0,
			expError: nil,
		},
		{
			name:     "success: second branch",
			order:    IavlSpec.InnerSpec.ChildOrder,
			branch:   1,
			expPos:   1,
			expError: nil,
		},
		{
			name:     "failure: negative branch value",
			order:    IavlSpec.InnerSpec.ChildOrder,
			branch:   -1,
			expPos:   -1,
			expError: fmt.Errorf("invalid branch: %d", int32(-1)),
		},
		{
			name:     "failure: branch is greater than length of child order",
			order:    IavlSpec.InnerSpec.ChildOrder,
			branch:   3,
			expPos:   -1,
			expError: fmt.Errorf("invalid branch: %d", int32(3)),
		},
		{
			name:     "failure: branch not found in child order",
			order:    []int32{0, 2},
			branch:   1,
			expPos:   -1,
			expError: fmt.Errorf("branch %d not found in order %v", 1, []int32{0, 2}),
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			pos, err := getPosition(tc.order, tc.branch)
			if tc.expError == nil && err != nil {
				t.Fatal(err)
			}

			if tc.expError != nil && (err.Error() != tc.expError.Error() || pos != -1) {
				t.Fatalf("expected: %v, got: %v", tc.expError, err)
			}

			if tc.expPos != pos {
				t.Fatalf("expected position: %d, got: %d", tc.expPos, pos)
			}
		})

	}
}
