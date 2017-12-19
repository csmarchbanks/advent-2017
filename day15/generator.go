package main

import (
	"fmt"
)

type Generator struct {
	value     uint64
	factor    uint64
	predicate func(uint64) bool
}

func NewGenerator(startValue, factor uint64) *Generator {
	return NewGeneratorWithPredicate(startValue, factor, func(a uint64) bool { return true })
}

func NewGeneratorWithPredicate(startValue, factor uint64, predicate func(uint64) bool) *Generator {
	return &Generator{value: startValue, factor: factor, predicate: predicate}
}

func (g *Generator) NextValue() uint64 {
	g.value *= g.factor
	g.value %= 2147483647
	if g.predicate(g.value) {
		return g.value
	}
	return g.NextValue()
}

func match(a, b uint64) bool {
	andBy := uint64(2)<<15 - 1
	a &= andBy
	b &= andBy
	return a == b
}

func MatchingPairs(a, b *Generator, trials int) (count int) {
	for i := 0; i < trials; i++ {
		if match(a.NextValue(), b.NextValue()) {
			count++
		}
	}
	return
}

func main() {
	a, b := NewGenerator(116, 16807), NewGenerator(299, 48271)
	fmt.Println(MatchingPairs(a, b, 40000000))

	aPredicate := func(a uint64) bool { return a%4 == 0 }
	bPredicate := func(a uint64) bool { return a%8 == 0 }
	a, b = NewGeneratorWithPredicate(116, 16807, aPredicate), NewGeneratorWithPredicate(299, 48271, bPredicate)
	fmt.Println(MatchingPairs(a, b, 5000000))
}
