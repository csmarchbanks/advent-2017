package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func readFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	matrix := [][]int{}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		row := []int{}
		for _, v := range line {
			intValue, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			row = append(row, intValue)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func rowDifference(row []int) int {
	min, max := 9999999, 0
	for _, v := range row {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func rowDivision(row []int) int {
	sort.Ints(row)
	for i, v := range row {
		for j := i + 1; j < len(row); j++ {
			if row[j]%v == 0 {
				return row[j] / v
			}
		}
	}
	return 0
}

func Checksum(data [][]int) (sum int) {
	for _, row := range data {
		sum += rowDifference(row)
	}
	return
}

func EvenDivisionSum(data [][]int) (sum int) {
	for _, row := range data {
		sum += rowDivision(row)
	}
	return
}

func main() {
	data := readFile("input")
	fmt.Println(Checksum(data))
	fmt.Println(EvenDivisionSum(data))
}
