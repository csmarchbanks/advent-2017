package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction uint8

const (
	up Direction = iota
	right
	down
	left
)

func (d Direction) Left() Direction {
	return (d - 1) % 4
}

func (d Direction) Right() Direction {
	return (d + 1) % 4
}

func (d Direction) Reverse() Direction {
	return (d + 2) % 4
}

type Virus interface {
	Burst()
	GetNodesInfected() int
}

type BaseVirus struct {
	direction     Direction
	i             int
	j             int
	cluster       *Cluster
	nodesInfected int
}

func NewBaseVirus(c *Cluster, i, j int) *BaseVirus {
	return &BaseVirus{
		direction:     up,
		i:             i,
		j:             j,
		cluster:       c,
		nodesInfected: 0,
	}
}

func (v BaseVirus) GetNodesInfected() int {
	return v.nodesInfected
}

func (v *BaseVirus) move() {
	switch v.direction {
	case up:
		v.i--
	case right:
		v.j++
	case down:
		v.i++
	case left:
		v.j--
	}
}

func (v *BaseVirus) Burst() {
	if v.cluster.GetNodeState(v.i, v.j) == Infected {
		v.direction = v.direction.Right()
		v.cluster.Clean(v.i, v.j)
	} else {
		v.direction = v.direction.Left()
		v.nodesInfected++
		v.cluster.Infect(v.i, v.j)
	}
	v.move()
}

type EvolvedVirus struct {
	*BaseVirus
}

func (v *EvolvedVirus) Burst() {
	switch v.cluster.GetNodeState(v.i, v.j) {
	case Clean:
		v.direction = v.direction.Left()
		v.cluster.Weaken(v.i, v.j)
	case Weakened:
		v.nodesInfected++
		v.cluster.Infect(v.i, v.j)
	case Infected:
		v.direction = v.direction.Right()
		v.cluster.Flag(v.i, v.j)
	case Flagged:
		v.direction = v.direction.Reverse()
		v.cluster.Clean(v.i, v.j)
	}
	v.move()
}

func readInputFromFile(filename string) [][]bool {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := [][]bool{}
	for scanner.Scan() {
		row := []bool{}
		for _, char := range scanner.Text() {
			if char == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		result = append(result, row)
	}
	return result
}

func simulate(virus Virus, times int) int {
	for i := 0; i < times; i++ {
		virus.Burst()
	}
	return virus.GetNodesInfected()
}

func simulateBase(input [][]bool) int {
	cluster := NewCluster(input)
	virus := NewBaseVirus(cluster, len(input)/2, len(input)/2)
	return simulate(virus, 10000)
}

func simulateEvolved(input [][]bool) int {
	cluster := NewCluster(input)
	virus := &EvolvedVirus{NewBaseVirus(cluster, len(input)/2, len(input)/2)}
	return simulate(virus, 10000000)
}

func main() {
	input := readInputFromFile("input")
	fmt.Println(simulateBase(input))
	fmt.Println(simulateEvolved(input))
}
