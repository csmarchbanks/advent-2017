package main

import (
	"fmt"
	"math/big"

	"github.com/csmarchbanks/advent-2017/utils"
)

type Grid [128][128]bool

func OnesCount(hash string) (sum int) {
	val := big.NewInt(0)
	_, success := val.SetString(hash, 16)
	if !success {
		panic("Failed to set string")
	}
	for i := 0; i < val.BitLen(); i++ {
		sum += int(val.Bit(i))
	}
	return
}

func UsedSquares(key string) (sum int) {
	for i := 0; i < 128; i++ {
		fullKey := fmt.Sprintf("%s-%d", key, i)
		hash := utils.Hash(fullKey)
		sum += OnesCount(hash)
	}
	return
}

func CreateGrid(key string) Grid {
	grid := Grid{}
	for i := 0; i < 128; i++ {
		fullKey := fmt.Sprintf("%s-%d", key, i)
		hash := utils.Hash(fullKey)
		val := big.NewInt(0)
		val.SetString(hash, 16)
		for j := 0; j < 128; j++ {
			grid[i][127-j] = val.Bit(j) == 1
		}
	}
	return grid
}

func getKey(i, j int) int {
	return (i << 16) + j
}

func MapRegion(grid Grid, cache map[int]int, i, j, regionValue int) {
	if i >= 0 && i < 128 && j >= 0 && j < 128 && grid[i][j] {
		key := getKey(i, j)
		if _, found := cache[key]; found {
			return
		} else {
			cache[key] = regionValue
		}
		MapRegion(grid, cache, i+1, j, regionValue)
		MapRegion(grid, cache, i-1, j, regionValue)
		MapRegion(grid, cache, i, j+1, regionValue)
		MapRegion(grid, cache, i, j-1, regionValue)
	}
}

func CountRegions(key string) int {
	regions := []int{}
	cache := make(map[int]int)
	grid := CreateGrid(key)
	for i, row := range grid {
		for j, v := range row {
			key := getKey(i, j)
			if _, found := cache[key]; v && !found {
				regionValue := len(regions) + 1
				regions = append(regions, regionValue)
				MapRegion(grid, cache, i, j, regionValue)
			}
		}
	}
	return len(regions)
}

func main() {
	key := "ffayrhll"
	fmt.Println(UsedSquares(key))
	fmt.Println(CountRegions(key))
}
