package main

import (
	"regexp"
	"strconv"
)

type Layer struct {
	Depth int
	Range int
}

func NewLayer(depth, r int) *Layer {
	return &Layer{Depth: depth, Range: r}
}

var layerRegex = regexp.MustCompile("(\\d+): (\\d+)$")

func NewLayerFromString(input string) *Layer {
	parsed := layerRegex.FindStringSubmatch(input)
	depth, err := strconv.Atoi(parsed[1])
	if err != nil {
		panic(err)
	}
	r, err := strconv.Atoi(parsed[2])
	if err != nil {
		panic(err)
	}
	return NewLayer(depth, r)
}

func (layer *Layer) isFound(time int) bool {
	period := (layer.Range - 1) * 2
	return time%period == 0
}

func (layer *Layer) Severity() int {
	return layer.Range * layer.Depth
}
