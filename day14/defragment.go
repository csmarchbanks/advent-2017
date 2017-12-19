package main

import (
	"fmt"
	"math/bits"
	"strconv"

	"github.com/csmarchbanks/advent-2017/utils"
)

func OnesCount(hashPart string) int {
	val, err := strconv.ParseUint(hashPart, 16, 64)
	if err != nil {
		panic(err)
	}
	return bits.OnesCount64(val)
}

func UsedSquares(key string) (sum int) {
	for i := 0; i < 128; i++ {
		fullKey := fmt.Sprintf("%s-%d", key, i)
		hash := utils.Hash(fullKey)
		part1 := hash[0:16]
		part2 := hash[16:32]
		sum += OnesCount(part1) + OnesCount(part2)
	}
	return
}

func main() {
	fmt.Println(UsedSquares("ffayrhll"))
}
