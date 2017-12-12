package main

import "testing"

func TestSolve(t *testing.T) {
	tcs := []struct {
		jumps  []int
		result int
	}{
		{[]int{0, 3, 0, 1, -3}, 5},
	}
	for _, tc := range tcs {
		result := Solve(tc.jumps)
		if result != tc.result {
			t.Errorf("Solve(%v) => %d, expected %d\n", tc.jumps, result, tc.result)
		}
	}
}

func TestSolve2(t *testing.T) {
	tcs := []struct {
		jumps  []int
		result int
	}{
		{[]int{0, 3, 0, 1, -3}, 10},
	}
	for _, tc := range tcs {
		result := Solve2(tc.jumps)
		if result != tc.result {
			t.Errorf("Solve2(%v) => %d, expected %d\n", tc.jumps, result, tc.result)
		}
	}
}
