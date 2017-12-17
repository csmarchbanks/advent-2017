package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Solve(jumps []int) (count int) {
	for i := 0; i >= 0 && i < len(jumps); {
		v := jumps[i]
		jumps[i]++
		i += v
		count++
	}
	return
}

func Solve2(jumps []int) (count int) {
	for i := 0; i >= 0 && i < len(jumps); {
		v := jumps[i]
		if v > 2 {
			jumps[i]--
		} else {
			jumps[i]++
		}
		i += v
		count++
	}
	return
}

func readJumps(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	jumps := []int{}
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		jumps = append(jumps, v)
	}
	return jumps
}

func main() {
	jumps := readJumps("input")
	jumps2 := make([]int, len(jumps), len(jumps))
	copy(jumps2, jumps)
	fmt.Println(Solve(jumps))
	fmt.Println(Solve2(jumps2))
}
