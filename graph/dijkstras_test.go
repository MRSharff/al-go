package graph

import (
	"testing"
)

type edge struct {
	v, u Node
}

type myGraph struct {
	adjacencyList map[Node][]Node
	edgeWeights   map[edge]int
}

func (m myGraph) Nodes() []Node {
	var nodes []Node
	for n := range m.adjacencyList {
		nodes = append(nodes, n)
	}
	return nodes
}

func (m myGraph) Neighbors(f Node) []Node {
	return m.adjacencyList[f]
}

func (m myGraph) Weight(f Node, neighbor Node) int {
	if f < neighbor {
		return m.edgeWeights[edge{f, neighbor}]
	}
	return m.edgeWeights[edge{neighbor, f}]
}

func TestDijkstras(t *testing.T) {
	g := myGraph{}
	g.adjacencyList = make(map[Node][]Node)
	g.edgeWeights = make(map[edge]int)
	g.adjacencyList[0] = append(g.adjacencyList[0], 1, 2, 3)
	g.adjacencyList[1] = append(g.adjacencyList[1], 0, 3)
	g.adjacencyList[2] = append(g.adjacencyList[2], 0)
	g.adjacencyList[3] = append(g.adjacencyList[3], 0, 1, 4)
	g.adjacencyList[4] = append(g.adjacencyList[4], 3)
	g.edgeWeights[edge{0, 1}] = 4
	g.edgeWeights[edge{0, 2}] = 4
	g.edgeWeights[edge{0, 3}] = 1
	g.edgeWeights[edge{3, 4}] = 5
	g.edgeWeights[edge{1, 3}] = 1

	shortestPath := Dijkstras(g, 0, 1)
	if shortestPath != 2 {
		t.Fail()
	}
}

func TestPriorityQueue(t *testing.T) {
	nodes := []Node{1, 2, 3, 4, 5}
	var pq PriorityQueue
	for i, n := range nodes {
		pq.Push(n, len(nodes)-i)
	}

	reversed := []Node{5, 4, 3, 2, 1}
	for i, n := 0, pq.Pop(); pq.Len() > 0; i, n = i+1, pq.Pop() {
		if reversed[i] != n {
			t.Errorf("Expected %d to be %d\n", n, reversed[i])
		}
	}
}
