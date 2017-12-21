package main

import "testing"

func TestUsedSquares(t *testing.T) {
	tcs := []struct {
		in  string
		out int
	}{
		{"flqrgnkx", 8108},
		{"ffayrhll", 8190},
	}
	for _, tc := range tcs {
		result := UsedSquares(tc.in)
		if result != tc.out {
			t.Errorf("UsedSquares(%s) => %d, expected %d\n", tc.in, result, tc.out)
		}
	}
}

func TestCountRegions(t *testing.T) {
	tcs := []struct {
		in  string
		out int
	}{
		{"flqrgnkx", 1242},
		{"ffayrhll", 1134},
	}
	for _, tc := range tcs {
		result := CountRegions(tc.in)
		if result != tc.out {
			t.Errorf("UsedSquares(%s) => %d, expected %d\n", tc.in, result, tc.out)
		}
	}
}
