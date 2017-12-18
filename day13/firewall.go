package main

import (
	"bufio"
	"fmt"
	"os"
)

type Firewall struct {
	layers []*Layer
	delay  int
}

func NewFirewall(layers ...*Layer) *Firewall {
	return &Firewall{layers: layers, delay: 0}
}

func NewFirewallFromFile(filename string) *Firewall {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	layers := []*Layer{}
	for scanner.Scan() {
		layers = append(layers, NewLayerFromString(scanner.Text()))
	}
	return NewFirewall(layers...)
}

func (firewall *Firewall) navigateFirewall() (severity int, found bool) {
	for _, layer := range firewall.layers {
		inColAt := firewall.delay + layer.Depth
		if layer.isFound(inColAt) {
			severity += layer.Severity()
			found = true
		}
	}
	return
}

func (firewall *Firewall) navigateFirewallWithDelay(delay int) (int, bool) {
	firewall.delay = delay
	return firewall.navigateFirewall()
}

func (firewall *Firewall) FindSeverity() (severity int) {
	severity, _ = firewall.navigateFirewall()
	return
}

func (firewall *Firewall) FindSmallestDelay() (delay int) {
	for {
		_, found := firewall.navigateFirewallWithDelay(delay)
		if !found {
			return
		}
		delay++
	}
}

func main() {
	firewall := NewFirewallFromFile("input")
	fmt.Println(firewall.FindSeverity())
	fmt.Println(firewall.FindSmallestDelay())
}
