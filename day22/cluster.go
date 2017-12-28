package main

type NodeState uint8

const (
	Clean NodeState = iota
	Weakened
	Infected
	Flagged
)

type Cluster struct {
	computerMap map[int64]NodeState
}

func NewCluster(input [][]bool) *Cluster {
	cluster := &Cluster{computerMap: make(map[int64]NodeState)}
	for i, row := range input {
		for j, v := range row {
			if v {
				cluster.Infect(i, j)
			}
		}
	}
	return cluster
}

func getKey(i, j int) int64 {
	return int64(i)<<32 + int64(j)
}

func (c *Cluster) GetNodeState(i, j int) NodeState {
	return c.computerMap[getKey(i, j)]
}

func (c *Cluster) Infect(i, j int) {
	c.computerMap[getKey(i, j)] = Infected
}

func (c *Cluster) Weaken(i, j int) {
	c.computerMap[getKey(i, j)] = Weakened
}

func (c *Cluster) Flag(i, j int) {
	c.computerMap[getKey(i, j)] = Flagged
}

func (c *Cluster) Clean(i, j int) {
	delete(c.computerMap, getKey(i, j))
}
