package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Duet struct {
	registers      map[string]int
	sendCount      int
	sendChannel    chan int
	receiveChannel chan int
	index          int
}

func (duet *Duet) getIntValue(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		intValue = duet.registers[value]
	}
	return intValue
}

func (duet *Duet) Snd(value string) {
	duet.sendCount++
	//fmt.Printf("%d has sent %d\n", duet.index, duet.getIntValue(value))
	duet.sendChannel <- duet.getIntValue(value)
}

func (duet *Duet) Set(x, y string) {
	duet.registers[x] = duet.getIntValue(y)
}

func (duet *Duet) Add(x, y string) {
	duet.registers[x] += duet.getIntValue(y)
}

func (duet *Duet) Mul(x, y string) {
	duet.registers[x] *= duet.getIntValue(y)
}

func (duet *Duet) Mod(x, y string) {
	duet.registers[x] %= duet.getIntValue(y)
}

func (duet *Duet) Rcv(x string) error {
	timeout := time.After(1 * time.Second)
	select {
	case duet.registers[x] = <-duet.receiveChannel:
		return nil
	case <-timeout:
		return errors.New("Deadlock detected")
	}
}

func (duet *Duet) Jgz(x, y string) int {
	if duet.getIntValue(x) > 0 {
		return duet.getIntValue(y)
	}
	return 1
}

func (duet *Duet) DoInstructions(instructions []string, resultCh chan int) {
	i := 0
	for {
		instruction := instructions[i%len(instructions)]
		values := strings.Split(instruction, " ")
		operation := values[0]
		switch operation {
		case "snd":
			duet.Snd(values[1])
		case "set":
			duet.Set(values[1], values[2])
		case "add":
			duet.Add(values[1], values[2])
		case "mul":
			duet.Mul(values[1], values[2])
		case "mod":
			duet.Mod(values[1], values[2])
		case "rcv":
			error := duet.Rcv(values[1])
			if error != nil {
				resultCh <- duet.sendCount
			}
		case "jgz":
			i += duet.Jgz(values[1], values[2])
			continue
		default:
			panic("Unrecognized instruction" + instruction)
		}
		i++
	}
}

func NewDuet(sendChannel, receiveChannel chan int, index int) *Duet {
	registers := make(map[string]int)
	registers["p"] = index
	return &Duet{
		registers:      registers,
		sendChannel:    sendChannel,
		receiveChannel: receiveChannel,
		index:          index,
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

func main() {
	instructions := readInstructionsFromFile("input")
	/*
		instructions = []string{
			"snd 1",
			"snd 2",
			"snd p",
			"rcv a",
			"rcv b",
			"rcv c",
			"rcv d",
		}
	*/
	ch1, ch2 := make(chan int, 500), make(chan int, 500)
	resultCh1, resultCh2 := make(chan int), make(chan int)
	duet0, duet1 := NewDuet(ch1, ch2, 0), NewDuet(ch2, ch1, 1)
	go duet0.DoInstructions(instructions, resultCh1)
	go duet1.DoInstructions(instructions, resultCh2)
	fmt.Println(<-resultCh2)
}
