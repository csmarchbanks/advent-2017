package main

import "testing"

func TestSolve(t *testing.T) {
	tcs := []struct {
		input  []int
		result int
	}{
		{[]int{1, 1, 2, 2}, 3},
		{[]int{1, 1, 1, 1}, 4},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{9, 1, 2, 1, 2, 1, 2, 9}, 9},
	}
	for _, tc := range tcs {
		result := Solve(tc.input)
		if result != tc.result {
			t.Errorf("Solve(%v) => %d, expected %d\n", tc.input, result, tc.result)
		}
	}
}

func TestSolveWithOffset(t *testing.T) {
	tcs := []struct {
		input  []int
		result int
	}{
		{[]int{1, 2, 1, 2}, 6},
		{[]int{1, 2, 2, 1}, 0},
		{[]int{1, 2, 3, 4, 2, 5}, 4},
		{[]int{1, 2, 3, 1, 2, 3}, 12},
		{[]int{1, 2, 1, 3, 1, 4, 1, 5}, 4},
	}
	for _, tc := range tcs {
		result := SolveHalfwayAround(tc.input)
		if result != tc.result {
			t.Errorf("Solve(%v) => %d, expected %d\n", tc.input, result, tc.result)
		}
	}
}
