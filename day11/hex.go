package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	n  = "n"
	ne = "ne"
	se = "se"
	s  = "s"
	sw = "sw"
	nw = "nw"
)

var possibleDirections = []string{n, ne, se, s, sw, nw}
var directionIndices = func() map[string]int {
	result := make(map[string]int)
	for i, direction := range possibleDirections {
		result[direction] = i
	}
	return result
}()

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func combinationDirections(direction string) (string, string) {
	nDirections := len(possibleDirections)
	index := directionIndices[direction] + nDirections
	return possibleDirections[(index-1)%nDirections], possibleDirections[(index+1)%nDirections]
}

func oppositeDirection(direction string) string {
	nDirections := len(possibleDirections)
	index := directionIndices[direction] + nDirections/2
	return possibleDirections[index%nDirections]
}

func orderByDirection(directions []string) map[string]int {
	result := make(map[string]int)
	for _, direction := range possibleDirections {
		result[direction] = 0
	}
	for _, direction := range directions {
		result[direction]++
	}
	return result
}

func reduceCombinations(byDirection map[string]int) (updated bool) {
	for direction := range byDirection {
		left, right := combinationDirections(direction)
		leftCount := byDirection[left]
		rightCount := byDirection[right]
		count := min(leftCount, rightCount)
		if count > 0 {
			byDirection[direction] += count
			byDirection[left] -= count
			byDirection[right] -= count
			updated = true
		}
	}
	return
}

func reduceOpposites(byDirection map[string]int) (updated bool) {
	for direction, count := range byDirection {
		oppDirection := oppositeDirection(direction)
		oppositeCount := byDirection[oppDirection]

		if oppositeCount < count {
			count = oppositeCount
		}
		if count > 0 {
			byDirection[direction] -= count
			byDirection[oppDirection] -= count
			updated = true
		}
	}
	return
}

func sumCounts(byDirection map[string]int) (sum int) {
	for _, count := range byDirection {
		sum += count
	}
	return
}

func findShortestPathFromMap(byDirection map[string]int) int {
	for reduceOpposites(byDirection) || reduceCombinations(byDirection) {
	}

	return sumCounts(byDirection)
}

// FindShortestPath calculates the shortest path to the
// ending location
func FindShortestPath(directions []string) int {
	byDirection := orderByDirection(directions)
	return findShortestPathFromMap(byDirection)
}

// FurthestDistance calculates the furthest distance ever gone
// it does this by continuously recalculating the shortest path
// after adding one direction at a time to the byDirection map
func FurthestDistance(directions []string) (max int) {
	byDirection := orderByDirection([]string{})
	for _, direction := range directions {
		byDirection[direction]++
		distance := findShortestPathFromMap(byDirection)
		if distance > max {
			max = distance
		}
	}
	return
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ",")...)
	}
	return result
}

func main() {
	input := readFile("input")
	fmt.Println(FindShortestPath(input))
	fmt.Println(FurthestDistance(input))
}
