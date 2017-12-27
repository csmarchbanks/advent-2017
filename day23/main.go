package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Processor struct {
	vals     map[string]int
	mulCount int
}

func NewProcessor() *Processor {
	vals := make(map[string]int)
	return &Processor{vals: vals}
}

func (p *Processor) getIntValue(register string) int {
	v, err := strconv.Atoi(register)
	if err != nil {
		v = p.vals[register]
	}
	return v
}

func (p *Processor) DoInstruction(instruction string) int {
	input := strings.Split(instruction, " ")

	switch input[0] {
	case "set":
		p.vals[input[1]] = p.getIntValue(input[2])
	case "sub":
		p.vals[input[1]] -= p.getIntValue(input[2])
	case "mul":
		p.mulCount++
		p.vals[input[1]] *= p.getIntValue(input[2])
	case "jnz":
		if p.getIntValue(input[1]) != 0 {
			return p.getIntValue(input[2])
		}
	default:
		panic(fmt.Errorf("Instruction (%s) not found", input[0]))
	}
	return 1
}

func (p *Processor) DoInstructions(instructions []string) {
	for i := 0; i < len(instructions); {
		i += p.DoInstruction(instructions[i])
	}
}

func readInstructionsFromFile(filename string) []string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	str := string(bytes)
	str = strings.Trim(str, "\n")
	return strings.Split(str, "\n")
}

func IsPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	processor := NewProcessor()
	instructions := readInstructionsFromFile("input")
	processor.DoInstructions(instructions)
	fmt.Println(processor.mulCount)

	// Part two checks for all not prime numbers between starting b and b + 17000 incrementing by 17
	b, c, h := 0, 0, 0
	b = 84*100 + 100000
	c = b + 17000
	for ; b <= c; b += 17 {
		if !IsPrime(b) {
			h++
		}
	}
	fmt.Println(h)
}
