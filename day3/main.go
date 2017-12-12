package main

import (
	"fmt"
	"math"
)

type stateMachine struct {
	i              int
	j              int
	direction      int
	moves          int
	movesUntilTurn int
}

const (
	right = iota
	up
	left
	down
)

func newStateMachine(i, j int) *stateMachine {
	return &stateMachine{i, j, right, 0, 1}
}

func (sm *stateMachine) turn() {
	sm.direction++
	sm.direction = sm.direction % 4
}

func (sm *stateMachine) move() (int, int) {
	if sm.moves == sm.movesUntilTurn {
		sm.turn()
		sm.moves = 0
		if sm.direction == left || sm.direction == right {
			sm.movesUntilTurn++
		}
	}
	sm.moves++
	if sm.direction == right {
		sm.j++
	} else if sm.direction == down {
		sm.i++
	} else if sm.direction == left {
		sm.j--
	} else if sm.direction == up {
		sm.i--
	}
	return sm.i, sm.j
}

func Solve(n int) int {
	sm := newStateMachine(0, 0)
	for i := 1; i < n; i++ {
		sm.move()
	}
	row, col := sm.i, sm.j
	if row < 0 {
		row *= -1
	}
	if col < 0 {
		col *= -1
	}
	return row + col
}

func sumOfNeighbors(matrix [][]int, row int, col int) (sum int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				sum += matrix[row+i][col+j]
			}
		}
	}
	return
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}
}

func Solve2(n int) int {
	matrixSize := int(math.Sqrt(float64(n))) + 4 // give some extra padding for checking zeroed neighbors
	matrix := make([][]int, matrixSize, matrixSize)
	for i := range matrix {
		matrix[i] = make([]int, matrixSize)
	}
	startI, startJ := matrixSize/2, matrixSize/2
	sm := newStateMachine(startI, startJ)
	matrix[startI][startJ] = 1
	for {
		i, j := sm.move()
		value := sumOfNeighbors(matrix, i, j)
		matrix[i][j] = value
		if value > n {
			fmt.Println(value)
			break
		}
	}
	//printMatrix(matrix)
	return 0
}

func main() {
	value := 312051
	fmt.Println(Solve(value))
	Solve2(value)
}
