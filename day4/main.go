package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func IsValid(passphrase string) bool {
	wordMap := make(map[string]bool)
	words := strings.Split(passphrase, " ")
	for _, word := range words {
		word = sortWord(word)
		_, found := wordMap[word]
		if found {
			return false
		}
		wordMap[word] = true
	}
	return true
}

func sortWord(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		if IsValid(scanner.Text()) {
			count++
		}
	}
	fmt.Println(count)
}
