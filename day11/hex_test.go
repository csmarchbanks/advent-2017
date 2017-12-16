package main

import "testing"

func TestFindShortestPath(t *testing.T) {
	tcs := []struct {
		in  []string
		out int
	}{
		{[]string{ne, ne, ne}, 3},
		{[]string{ne, ne, sw, sw}, 0},
		{[]string{ne, ne, s, s}, 2},
		{[]string{se, sw, se, sw, sw}, 3},
	}
	for _, tc := range tcs {
		result := FindShortestPath(tc.in)
		if result != tc.out {
			t.Errorf("FindShortestPath(%v) => %d, expected %d\n", tc.in, result, tc.out)
		}
	}
}
