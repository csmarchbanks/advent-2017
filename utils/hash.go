package utils

import (
	"encoding/hex"
)

func reverseSubSlice(slice []byte, start, length int) {
	for i, end := start, start+length-1; i < end; i, end = i+1, end-1 {
		startIndex := i % len(slice)
		endIndex := end % len(slice)
		tmp := slice[startIndex]
		slice[startIndex] = slice[endIndex]
		slice[endIndex] = tmp
	}
}

func HashOnce(list, lengths []byte, i, skipSize int) (int, int, int) {
	for _, length := range lengths {
		reverseSubSlice(list, i, int(length))
		i += int(length) + skipSize
		skipSize++
	}
	return int(list[0]) * int(list[1]), i, skipSize
}

func GenerateSlice(n int) []byte {
	result := []byte{}
	for i := 0; i < n; i++ {
		result = append(result, byte(i))
	}
	return result
}

func sparseHash(input string) []byte {
	list := GenerateSlice(256)
	bytes := []byte(input)
	bytes = append(bytes, 17, 31, 73, 47, 23)
	start, skipSize := 0, 0
	for i := 0; i < 64; i++ {
		_, start, skipSize = HashOnce(list, bytes, start, skipSize)
	}
	return list
}

func denseHash(bytes []byte) []byte {
	result := []byte{}
	for i := 0; i < 16; i++ {
		n := byte(0)
		for j := i * 16; j < (i+1)*16; j++ {
			n ^= bytes[j]
		}
		result = append(result, n)
	}
	return result
}

func Hash(input string) string {
	sparse := sparseHash(input)
	dense := denseHash(sparse)
	return hex.EncodeToString(dense)
}
