package main

import "testing"

func TestMatchingPairs(t *testing.T) {
	a, b := NewGenerator(65, 16807), NewGenerator(8921, 48271)
	expectedResult := 588
	result := MatchingPairs(a, b, 40000000)
	if result != expectedResult {
		t.Errorf("MatchingPairs => %d, expected %d\n", result, expectedResult)
	}
}
