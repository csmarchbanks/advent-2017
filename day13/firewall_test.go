package main

import "testing"

func TestFindSeverity(t *testing.T) {
	firewall := NewFirewall(NewLayer(0, 3), NewLayer(1, 2), NewLayer(4, 4), NewLayer(6, 4))
	expectedSeverity := 24
	result := firewall.FindSeverity()
	if result != expectedSeverity {
		t.Errorf("Expected severity = %d, got %d\n", expectedSeverity, result)
	}
}

func TestFindSmallestDelay(t *testing.T) {
	firewall := NewFirewall(NewLayer(0, 3), NewLayer(1, 2), NewLayer(4, 4), NewLayer(6, 4))
	expectedDelay := 10
	result := firewall.FindSmallestDelay()
	if result != expectedDelay {
		t.Errorf("Expected delay = %d, got %d\n", expectedDelay, result)
	}
}
