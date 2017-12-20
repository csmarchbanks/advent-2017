package main

import "testing"

func TestSpin(t *testing.T) {
	tcs := []struct {
		in          string
		instruction string
		out         string
	}{
		{"abcde", "1", "eabcd"},
		{"abcde", "3", "cdeab"},
	}
	for _, tc := range tcs {
		result := Spin(tc.in, tc.instruction)
		if result != tc.out {
			t.Errorf("Spin(%s, %s) => %s, expected %s\n", tc.in, tc.instruction, result, tc.out)
		}
	}
}

func TestExchange(t *testing.T) {
	tcs := []struct {
		in          string
		instruction string
		out         string
	}{
		{"abcde", "1/4", "aecdb"},
		{"eabcd", "3/4", "eabdc"},
	}
	for _, tc := range tcs {
		result := Exchange(tc.in, tc.instruction)
		if result != tc.out {
			t.Errorf("Spin(%s, %s) => %s, expected %s\n", tc.in, tc.instruction, result, tc.out)
		}
	}
}

func TestPartner(t *testing.T) {
	tcs := []struct {
		in          string
		instruction string
		out         string
	}{
		{"eabdc", "e/b", "baedc"},
	}
	for _, tc := range tcs {
		result := Partner(tc.in, tc.instruction)
		if result != tc.out {
			t.Errorf("Spin(%s, %s) => %s, expected %s\n", tc.in, tc.instruction, result, tc.out)
		}
	}
}
