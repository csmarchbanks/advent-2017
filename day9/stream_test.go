package main

import "testing"

func TestScoreFromString(t *testing.T) {
	tcs := []struct {
		in           string
		out          int
		garbageCount int
	}{
		{"{}", 1, 0},
		{"{{{}}}", 6, 0},
		{"{{},{}}", 5, 0},
		{"{{{},{},{{}}}}", 16, 0},
		{"{<a>,<a>,<a>,<a>}", 1, 4},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17},
	}
	for _, tc := range tcs {
		result, garbageCount := ScoreFromString(tc.in)
		if result != tc.out || garbageCount != tc.garbageCount {
			t.Errorf("CountGroupsFromString(%s) => %d, %d, expected %d, %d\n", tc.in, result, garbageCount, tc.out, tc.garbageCount)
		}
	}
}
