package main

import "testing"

func TestMax(t *testing.T) {
	tcs := []struct {
		in       []int
		maxIndex int
		maxValue int
	}{
		{[]int{1, 2, 3, 4}, 3, 4},
		{[]int{4, 7, 1, 2}, 1, 7},
	}
	for _, tc := range tcs {
		i, v := Max(tc.in)
		if i != tc.maxIndex || v != tc.maxValue {
			t.Errorf("Max(%v) => %d, %d, expected %d, %d\n", tc.in, i, v, tc.maxIndex, tc.maxValue)
		}
	}
}

func TestRedistributeUntilRepeat(t *testing.T) {
	tcs := []struct {
		in   []int
		out  int
		diff int
	}{
		{[]int{0, 2, 7, 0}, 5, 4},
	}
	for _, tc := range tcs {
		result, diff := RedistributeUntilRepeat(tc.in)
		if result != tc.out {
			t.Errorf("RedistributeUntilRepeat(%v) => %d, %d, expected %d, %d\n", tc.in, result, diff, tc.out, tc.diff)
		}
	}
}
