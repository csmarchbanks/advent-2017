package main

import "testing"

func TestIsValid(t *testing.T) {
	tcs := []struct {
		passphrase string
		valid      bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}
	for _, tc := range tcs {
		result := IsValid(tc.passphrase)
		if result != tc.valid {
			t.Errorf("IsValid(%s) => %t, expected %t\n", tc.passphrase, result, tc.valid)
		}
	}
}
