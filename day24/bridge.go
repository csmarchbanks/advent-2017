package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Component struct {
	used  bool
	lPort int
	rPort int
}

func (c *Component) String() string {
	return fmt.Sprintf("%d/%d-%t", c.lPort, c.rPort, c.used)
}

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func ReadComponentCache(filename string) map[int][]*Component {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	cache := make(map[int][]*Component)
	for scanner.Scan() {
		ports := strings.Split(scanner.Text(), "/")
		component := &Component{
			used:  false,
			lPort: Atoi(ports[0]),
			rPort: Atoi(ports[1]),
		}
		cache[component.lPort] = append(cache[component.lPort], component)
		cache[component.rPort] = append(cache[component.rPort], component)
	}
	return cache
}

func calculateMaxLengthAndStrength(component *Component, availablePort int, currentLength int, components map[int][]*Component) (int, int) {
	maxLength, maxStrength := 0, 0
	component.used = true
	defer func() {
		component.used = false
	}()
	for _, c := range components[availablePort] {
		if !c.used {
			aPort := c.lPort
			if aPort == availablePort {
				aPort = c.rPort
			}
			length, strength := calculateMaxLengthAndStrength(c, aPort, 1, components)

			if length > maxLength || (length == maxLength && strength > maxStrength) {
				maxLength = length
				maxStrength = strength
			}
		}
	}
	return currentLength + maxLength, maxStrength + component.lPort + component.rPort
}

func BuildBridge(components map[int][]*Component) (maxLength, maxStrength int) {
	for _, component := range components[0] {
		availablePort := component.lPort
		if availablePort == 0 {
			availablePort = component.rPort
		}
		length, strength := calculateMaxLengthAndStrength(component, availablePort, 1, components)
		if length > maxLength || (length == maxLength && strength > maxStrength) {
			maxLength = length
			maxStrength = strength
		}
	}
	return
}

func main() {
	start := time.Now()
	cache := ReadComponentCache("input")
	fmt.Println(BuildBridge(cache))
	fmt.Printf("Took: %s\n", time.Since(start))
}
