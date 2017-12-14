package main

import (
	"fmt"
	"io/ioutil"
)

func Score(c chan byte) (count, garbageCount int) {
	nLeft := 0
	inGarbage := false
	ignoreNext := false
	for char := range c {
		if ignoreNext {
			ignoreNext = false
		} else if inGarbage {
			switch char {
			case '!':
				ignoreNext = true
			case '>':
				inGarbage = false
			default:
				garbageCount++
			}
		} else {
			switch char {
			case '{':
				nLeft++
			case '}':
				count += nLeft
				nLeft--
			case '<':
				inGarbage = true
			case '!':
				ignoreNext = true
			}
		}
	}
	return
}

func addBytesToChan(bytes []byte, c chan byte) {
	for _, char := range bytes {
		c <- char
	}
	close(c)
}

func ScoreFromString(input string) (int, int) {
	c := make(chan byte)
	bytes := []byte(input)
	go addBytesToChan(bytes, c)
	return Score(c)
}

func ScoreFromFile(filename string) (int, int) {
	c := make(chan byte)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	go addBytesToChan(bytes, c)
	return Score(c)
}

func main() {
	fmt.Println(ScoreFromFile("input"))
}
