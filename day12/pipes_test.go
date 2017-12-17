package main

import "testing"

var input = []string{
	"0 <-> 2",
	"1 <-> 1",
	"2 <-> 0, 3, 4",
	"3 <-> 2, 4",
	"4 <-> 2, 3, 6",
	"5 <-> 6",
	"6 <-> 4, 5",
}

func TestCountGroups(t *testing.T) {
	expectedResult := 6
	graph := NewGraph(input)
	result := graph.CountInGroup(0)
	if result != expectedResult {
		t.Errorf("CountInGroup => %d, expected %d\n", result, expectedResult)
	}
}

func TestNGroups(t *testing.T) {
	expectedResult := 2
	graph := NewGraph(input)
	result := graph.NGroups()
	if result != expectedResult {
		t.Errorf("NGroups => %d, expected %d\n", result, expectedResult)
	}
}
