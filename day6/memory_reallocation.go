package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Max returns the index and value of the maximum entry in the array
func Max(array []int) (maxIndex, maxValue int) {
	for i, v := range array {
		if v > maxValue {
			maxIndex = i
			maxValue = v
		}
	}
	return
}

func redistribute(blocks []int) {
	offset, v := Max(blocks)
	blocks[offset] = 0
	for i := 0; i < v; i++ {
		blocks[(i+offset+1)%len(blocks)]++
	}
}

func RedistributeUntilRepeat(blocks []int) (count, diff int) {
	cache := make(map[string]int)
	for {
		key := fmt.Sprint(blocks)
		if firstSeen, found := cache[key]; found {
			diff = count - firstSeen
			return
		}
		cache[key] = count
		redistribute(blocks)
		count++
	}
}

func readFile(filename string) []int {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	str := string(bytes)
	str = strings.TrimSpace(str)
	strArray := strings.Split(str, "\t")
	blocks := []int{}
	for _, strValue := range strArray {
		v, err := strconv.Atoi(strValue)
		if err != nil {
			panic(err)
		}
		blocks = append(blocks, v)
	}
	return blocks
}

func main() {
	blocks := readFile("input")
	fmt.Println(RedistributeUntilRepeat(blocks))
}
