package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Graph struct {
	programMap map[int]*Program
}

type Program struct {
	ID          int
	connections []*Program
}

func addProgramToResults(program *Program, results map[int]bool) {
	results[program.ID] = true
	for _, connection := range program.connections {
		_, found := results[connection.ID]
		if !found {
			addProgramToResults(connection, results)
		}
	}
}

func (g *Graph) CountInGroup(id int) int {
	results := make(map[int]bool)
	addProgramToResults(g.programMap[id], results)
	return len(results)
}

func (g *Graph) NGroups() (count int) {
	results := make(map[int]bool)
	for _, program := range g.programMap {
		_, found := results[program.ID]
		if !found {
			count++
			addProgramToResults(program, results)
		}
	}
	return
}

var programParser = regexp.MustCompile("(\\d+) <-> (.*)")

func createAndAddConnections(connectionsString string, programMap map[int]*Program) []*Program {
	connectionStrIds := strings.Split(connectionsString, ", ")
	connections := []*Program{}
	for _, strId := range connectionStrIds {
		connId, err := strconv.Atoi(strId)
		if err != nil {
			panic(err)
		}
		connection, found := programMap[connId]
		if !found {
			connection = &Program{ID: connId}
			programMap[connId] = connection
		}
		connections = append(connections, connection)
	}
	return connections
}

func NewGraph(inputs []string) *Graph {
	programMap := make(map[int]*Program)
	for _, input := range inputs {
		parsed := programParser.FindStringSubmatch(input)
		id, err := strconv.Atoi(parsed[1])
		if err != nil {
			panic(err)
		}

		connections := createAndAddConnections(parsed[2], programMap)

		program, found := programMap[id]
		if !found {
			program = &Program{ID: id, connections: connections}
		} else {
			program.connections = connections
		}
		programMap[id] = program
	}
	return &Graph{programMap: programMap}
}

func NewGraphFromFile(filename string) *Graph {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return NewGraph(input)
}

func main() {
	graph := NewGraphFromFile("input")
	fmt.Println(graph.CountInGroup(0))
	fmt.Println(graph.NGroups())
}
