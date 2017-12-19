package utils

import (
	"reflect"
	"testing"
)

func TestReverseSubSlice(t *testing.T) {
	tcs := []struct {
		slice  []byte
		start  int
		length int
		result []byte
	}{
		{[]byte{0, 1, 2, 3, 4}, 0, 3, []byte{2, 1, 0, 3, 4}},
		{[]byte{2, 1, 0, 3, 4}, 3, 4, []byte{4, 3, 0, 1, 2}},
		{[]byte{4, 3, 0, 1, 2}, 3, 1, []byte{4, 3, 0, 1, 2}},
		{[]byte{4, 3, 0, 1, 2}, 1, 5, []byte{3, 4, 2, 1, 0}},
	}
	for _, tc := range tcs {
		reverseSubSlice(tc.slice, tc.start, tc.length)
		if !reflect.DeepEqual(tc.slice, tc.result) {
			t.Errorf("Result slice: %v, expected %v\n", tc.slice, tc.result)
		}
	}
}

func TestHashOnce(t *testing.T) {
	tcs := []struct {
		slice   []byte
		lengths []byte
		result  int
	}{
		{GenerateSlice(5), []byte{3, 4, 1, 5}, 12},
	}
	for _, tc := range tcs {
		result, _, _ := HashOnce(tc.slice, tc.lengths, 0, 0)
		if result != tc.result {
			t.Errorf("Hash(%v, %v) => %d, expected %d\n", tc.slice, tc.lengths, result, tc.result)
		}
	}
}

func TestHash(t *testing.T) {
	input := "120,93,0,90,5,80,129,74,1,165,204,255,254,2,50,113"
	expectedOutput := "d067d3f14d07e09c2e7308c3926605c4"
	result := Hash(input)
	if result != expectedOutput {
		t.Errorf("Hash(%s) => %s, expected %s\n", input, result, expectedOutput)
	}
}
