package main

import "fmt"

func GenerateSpinlock(steps int) int {
	array := []int{0}
	currentPosition := 0
	for i := 1; i < 2018; i++ {
		// increment current position appropriately
		currentPosition += steps
		currentPosition %= len(array)

		// split array into before and after parts based on the current position
		part1 := array[:currentPosition+1]
		part2 := array[currentPosition+1:]

		// make a new array, copy part 1 into it
		newArray := make([]int, len(array)+1)
		copy(newArray, part1)

		// set the new value in the array
		newArray[currentPosition+1] = i

		// slice to copy part2 of the array into
		rest := newArray[currentPosition+2:]
		copy(rest, part2)

		array = newArray
		currentPosition++
	}
	fmt.Println(array[currentPosition-3 : currentPosition+3])
	return array[currentPosition+1]
}

func FindValueAfter0(steps int) (nAfter0 int) {
	currentPosition := 0
	length := 1
	for i := 1; i <= 50000000; i++ {
		currentPosition += steps
		currentPosition %= length
		currentPosition++
		if currentPosition == 1 {
			nAfter0 = i
		}
		length++
	}
	return
}

func main() {
	fmt.Println(GenerateSpinlock(301))
	fmt.Println(FindValueAfter0(301))
}
