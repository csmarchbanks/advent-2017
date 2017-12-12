package main

import "testing"

func TestSolve(t *testing.T) {
	tcs := []struct {
		n      int
		result int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}
	for _, tc := range tcs {
		result := Solve(tc.n)
		if result != tc.result {
			t.Errorf("Solve(%d) => %d, expected %d\n", tc.n, result, tc.result)
		}
	}
}
