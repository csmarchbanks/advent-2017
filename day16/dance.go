package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Spin(dancers string, instruction string) string {
	moveBy, err := strconv.Atoi(instruction)
	if err != nil {
		panic(err)
	}
	splitPoint := len(dancers) - moveBy
	start := dancers[splitPoint:]
	end := dancers[:splitPoint]
	return start + end
}

func replaceAtIndex(str string, replacement byte, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func swap(i, j int, dancers string) string {
	str := replaceAtIndex(dancers, dancers[i], j)
	return replaceAtIndex(str, dancers[j], i)
}

func Exchange(dancers string, instruction string) string {
	values := strings.Split(instruction, "/")
	i, err := strconv.Atoi(values[0])
	if err != nil {
		panic(err)
	}
	j, err := strconv.Atoi(values[1])
	if err != nil {
		panic(err)
	}
	return swap(i, j, dancers)
}

func Partner(dancers string, instruction string) string {
	values := strings.Split(instruction, "/")
	i := strings.LastIndex(dancers, values[0])
	j := strings.LastIndex(dancers, values[1])
	return swap(i, j, dancers)
}

func DanceMove(dancers string, instruction string) string {
	switch instruction[0] {
	case 's':
		return Spin(dancers, instruction[1:])
	case 'x':
		return Exchange(dancers, instruction[1:])
	case 'p':
		return Partner(dancers, instruction[1:])
	}
	panic("Unrecognized move")
}

func Dance(dancers string, instructions []string) string {
	for _, instruction := range instructions {
		dancers = DanceMove(dancers, instruction)
	}
	return dancers
}

func readInstructionsFromFile(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	strInstructions := string(bytes)
	trimmed := strings.Trim(strInstructions, "\n")
	return strings.Split(trimmed, ",")
}

func findAfterBillionDances(input string, instructions []string) string {
	cache := make(map[string]int)
	// i = 0 corresponds to the original input
	for i := 0; i < 100000; i++ {
		cacheResult, found := cache[input]
		if found {
			period := i - cacheResult
			remaining := 1000000000 - i
			nDancesRemaining := remaining % period
			for j := 0; j < nDancesRemaining; j++ {
				input = Dance(input, instructions)
			}
			return input
		}
		cache[input] = i
		input = Dance(input, instructions)
	}

	panic("No repeat found")
}

func main() {
	input := "abcdefghijklmnop"
	instructions := readInstructionsFromFile("input")
	fmt.Println(Dance(input, instructions))
	fmt.Println(findAfterBillionDances(input, instructions))
}
