package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPredicateSatisfied(registers map[string]int, register string, operator string, compareValue int) bool {
	value := registers[register]
	switch operator {
	case "==":
		return value == compareValue
	case "!=":
		return value != compareValue
	case "<":
		return value < compareValue
	case "<=":
		return value <= compareValue
	case ">":
		return value > compareValue
	case ">=":
		return value >= compareValue
	}
	panic("Unrecognized operator")
}

func parseInstrucation(registers map[string]int, instruction string) int {
	instructions := strings.Split(instruction, " ")
	register := instructions[0]
	operation := instructions[1]
	value, err := strconv.Atoi(instructions[2])
	if err != nil {
		panic(err)
	}
	variable := instructions[4]
	operator := instructions[5]
	compareValue, err := strconv.Atoi(instructions[6])
	if err != nil {
		panic(err)
	}
	if isPredicateSatisfied(registers, variable, operator, compareValue) {
		switch operation {
		case "inc":
			registers[register] += value
		case "dec":
			registers[register] -= value
		}
	}
	return registers[register]
}

func DoInstructions(instructions []string) (map[string]int, int) {
	registers := make(map[string]int)
	max := 0
	for _, instruction := range instructions {
		v := parseInstrucation(registers, instruction)
		if v > max {
			max = v
		}
	}
	return registers, max
}

func FindMax(registers map[string]int) (max int) {
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return max
}

func ParseInstructionsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	return instructions
}

func main() {
	instructions := ParseInstructionsFromFile("input")
	registers, maxEverValue := DoInstructions(instructions)
	fmt.Println(FindMax(registers))
	fmt.Println(maxEverValue)
}
