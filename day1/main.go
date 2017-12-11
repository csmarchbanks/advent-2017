package main

import (
	"fmt"
	"io/ioutil"
)

func solveWithOffset(input []int, offset int) (sum int) {
	length := len(input)
	for i, v := range input {
		prevIndex := (i + length + offset) % length
		if v == input[prevIndex] {
			sum += v
		}
	}
	return
}

func SolveHalfwayAround(input []int) int {
	return solveWithOffset(input, len(input)/2)
}

func Solve(input []int) (sum int) {
	return solveWithOffset(input, -1)
}

func main() {
	input, err := ioutil.ReadFile("code.txt")
	if err != nil {
		panic(err)
	}
	ints := []int{}
	for _, v := range input {
		if v >= '0' && v <= '9' {
			ints = append(ints, int(v-'0'))
		} else {
			break
		}
	}
	fmt.Println(Solve(ints))
	fmt.Println(SolveHalfwayAround(ints))
}
