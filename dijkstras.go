package main

import (
	"sort"
)

type Node int

type Graph interface {
	Nodes() []Node
	Neighbors(f Node) []Node
	Weight(f Node, neighbor Node) int
}

func Dijkstras(g Graph, start Node, end Node) int {
	// This algorithm is a breadth first search, which means using a queue for our frontier.

	var distance = make(map[Node]int)

	// For each node in this set, the shortest distance from the source node is known
	var settled = make(map[Node]bool)

	// Each node in this set has been visited at least once.
	// The shortest distance is known when using edges we have traversed, which means
	// we may find a shorter distance still.
	var frontier = make(map[Node]bool)

	frontier[start] = true
	distance[start] = 0

	for len(frontier) > 0 {
		f := min(frontier, distance)
		delete(frontier, f)
		settled[f] = true
		if f == end {
			return distance[f]
		}
		for _, neighbor := range g.Neighbors(f) {
			d := distance[f] + g.Weight(f, neighbor)
			isFarOff := !settled[neighbor] && !frontier[neighbor]
			if isFarOff {
				distance[neighbor] = d
				frontier[neighbor] = true
			} else if d < distance[neighbor] {
				distance[neighbor] = d
			}

		}
	}

	return -1
}

func min(frontier map[Node]bool, distance map[Node]int) Node {
	var nodes []Node
	for n := range frontier {
		nodes = append(nodes, n)
	}
	sort.Slice(nodes, func(i, j int) bool {
		return distance[nodes[i]] < distance[nodes[j]]
	})
	return nodes[0]
}
