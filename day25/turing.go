package main

import (
	"fmt"
)

type State struct {
	instructions []*Instruction
}

type Instruction struct {
	write     int
	move      int
	nextState *State
}

type Machine struct {
	tape   map[int]int
	cursor int
	state  *State
}

func NewMachine(initialState *State) *Machine {
	return &Machine{
		tape:   make(map[int]int),
		cursor: 0,
		state:  initialState,
	}
}

func (machine *Machine) DoInstruction() {
	instruction := machine.state.instructions[machine.tape[machine.cursor]]
	if instruction.write == 0 {
		delete(machine.tape, machine.cursor)
	} else {
		machine.tape[machine.cursor] = 1
	}
	machine.cursor += instruction.move
	machine.state = instruction.nextState
}

func main() {
	steps := 12667664
	stateA := &State{}
	stateB := &State{}
	stateC := &State{}
	stateD := &State{}
	stateE := &State{}
	stateF := &State{}
	stateA.instructions = []*Instruction{
		&Instruction{1, 1, stateB},
		&Instruction{0, -1, stateC},
	}
	stateB.instructions = []*Instruction{
		&Instruction{1, -1, stateA},
		&Instruction{1, 1, stateD},
	}
	stateC.instructions = []*Instruction{
		&Instruction{0, -1, stateB},
		&Instruction{0, -1, stateE},
	}
	stateD.instructions = []*Instruction{
		&Instruction{1, 1, stateA},
		&Instruction{0, 1, stateB},
	}
	stateE.instructions = []*Instruction{
		&Instruction{1, -1, stateF},
		&Instruction{1, -1, stateC},
	}
	stateF.instructions = []*Instruction{
		&Instruction{1, 1, stateD},
		&Instruction{1, 1, stateA},
	}

	machine := NewMachine(stateA)

	for i := 0; i < steps; i++ {
		machine.DoInstruction()
	}

	fmt.Println(len(machine.tape))
}
