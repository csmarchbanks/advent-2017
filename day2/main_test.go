package main

import "testing"

func TestChecksum(t *testing.T) {
	input := [][]int{
		[]int{5, 1, 9, 5},
		[]int{7, 5, 3},
		[]int{2, 4, 6, 8},
	}
	expectedResult := 18
	result := Checksum(input)
	if result != expectedResult {
		t.Errorf("Expected: %d, got %d\n", expectedResult, result)
	}
}

func TestEvenDivisionSum(t *testing.T) {
	input := [][]int{
		[]int{5, 9, 2, 8},
		[]int{9, 4, 7, 3},
		[]int{3, 8, 6, 5},
	}
	expectedResult := 9
	result := EvenDivisionSum(input)
	if result != expectedResult {
		t.Errorf("Expected: %d, got %d\n", expectedResult, result)
	}
}
