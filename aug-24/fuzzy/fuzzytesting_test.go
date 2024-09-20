package main

import "testing"

func FuzzyReverse(t *testing.F) {
	tests := []string{"Hello", "World", ""}

	for _, test := range tests {
		t.Add(test)
	}

	t.Fuzz(func(t *testing.T, org string) {
		rev := Reverse(org)
		doubleRev := Reverse(rev)

		if org != doubleRev {
			t.Errorf("Before : %s , after %s", org, doubleRev)
		}
	})

}
