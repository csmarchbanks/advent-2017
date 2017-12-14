package main

import (
	"testing"
)

func TestDoInstructions(t *testing.T) {
	instructions := []string{
		"b inc 5 if a > 1",
		"a inc 1 if b < 5",
		"c dec -10 if a >= 1",
		"c inc -20 if c == 10",
	}
	expectedMax, expectedMaxEver := 1, 10
	registers, maxEver := DoInstructions(instructions)
	max := FindMax(registers)
	if max != expectedMax {
		t.Errorf("FindMax => %d, expected %d\n", max, expectedMax)
	}
	if maxEver != expectedMaxEver {
		t.Errorf("maxEver => %d, expected %d\n", maxEver, expectedMaxEver)
	}
}
