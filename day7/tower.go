package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	children []*Node
	parent   *Node
	name     string
	weight   int
}

var inputRegexp = regexp.MustCompile("(\\w+) \\((\\d+)\\)[ -> ]*(.*)")

func readFile(filename string) map[string]*Node {
	cache := make(map[string]*Node)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		groups := inputRegexp.FindStringSubmatch(input)
		name := groups[1]
		weight, err := strconv.Atoi(groups[2])
		if err != nil {
			panic(err)
		}
		childrenNames := []string{}
		if len(groups) > 3 {
			childrenNames = strings.Split(groups[3], ", ")
		}
		// find or create node
		node, found := cache[name]
		if !found {
			node = &Node{name: name}
		}
		node.weight = weight

		// deal with children
		children := []*Node{}
		for _, childName := range childrenNames {
			child, found := cache[childName]
			if !found {
				child = &Node{name: childName}
			}
			child.parent = node
			cache[childName] = child
			children = append(children, child)
		}
		node.children = children
		cache[name] = node
	}
	return cache
}

func FindRoot(cache map[string]*Node) *Node {
	for _, v := range cache {
		if v.parent == nil {
			return v
		}
	}
	return nil
}

func CalculateWeight(root *Node) (sum int) {
	sum = root.weight
	for _, child := range root.children {
		sum += CalculateWeight(child)
	}
	return
}

// DifferentIndex returns the index of the child that is different, and the
// correct value that the children should all be. If no children are different returns -1
func DifferentIndex(childrenWeights []int) (int, int) {
	if len(childrenWeights) == 0 {
		return -1, 0
	}
	sortedWeights := make([]int, len(childrenWeights))
	copy(sortedWeights, childrenWeights)
	sort.Ints(sortedWeights)
	correctValue := sortedWeights[len(sortedWeights)/2]
	for i, v := range childrenWeights {
		if v != correctValue {
			return i, correctValue
		}
	}
	return -1, correctValue
}

func FindNewWeightValue(root *Node, goalWeight int) int {
	weights := []int{}
	sumOfChildrenWeights := 0
	for _, child := range root.children {
		childWeight := CalculateWeight(child)
		sumOfChildrenWeights += childWeight
		weights = append(weights, childWeight)
	}

	i, correctWeight := DifferentIndex(weights)
	if i >= 0 {
		return FindNewWeightValue(root.children[i], correctWeight)
	}
	// if i < 0, then this is the bad node. Can find the weight by taking the goal weight minus the sum of the children weights
	return goalWeight - sumOfChildrenWeights
}

func main() {
	cache := readFile("input")
	root := FindRoot(cache)
	fmt.Printf("root=%s\n", root.name)
	fmt.Println(FindNewWeightValue(root, 0))
}
