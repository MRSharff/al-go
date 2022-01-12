package graph

import (
	"container/heap"
)

type Node int

type Graph interface {
	Neighbors(f Node) []Node
	Weight(f Node, neighbor Node) int
}

type Element struct {
	value           Node
	priority, index int
}

type nodeHeap []*Element

func (n nodeHeap) Len() int {
	return len(n)
}

func (n nodeHeap) Less(i, j int) bool {
	return n[i].priority < n[j].priority
}

func (n nodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
	n[i].index = i
	n[j].index = j
}

func (n *nodeHeap) Push(x interface{}) {
	l := len(*n)
	e := x.(*Element)
	e.index = l
	*n = append(*n, e)
}

func (n *nodeHeap) Pop() interface{} {
	old := *n
	l := len(old)
	element := old[l-1]
	old[l-1] = nil     // avoid memory leak
	element.index = -1 // for safety
	*n = old[0 : l-1]
	return element
}

type PriorityQueue struct {
	nodes nodeHeap
}

func (pq *PriorityQueue) Push(n Node, priority int) {
	if pq.nodes == nil {
		nh := make(nodeHeap, 0)
		pq.nodes = nh
	}
	element := &Element{n, priority, -1}
	heap.Push(&(pq.nodes), element)
}

func (pq *PriorityQueue) Pop() Node {
	n := heap.Pop(&(pq.nodes)).(*Element).value
	return n
}

func (pq *PriorityQueue) Update(n Node, priority int) {
	e := pq.nodes[n]
	e.priority = priority
	e.value = n
	heap.Fix(&(pq.nodes), e.index)
}

func (pq *PriorityQueue) Len() int {
	return len(pq.nodes)
}

func Dijkstras(g Graph, start Node, end Node) int {
	// This algorithm is a breadth first search, which means using a queue for our frontier.
	// A priority queue will help us optimize finding the Node with minimum distance in the frontier.
	// Even better, we can use a monotonic priority queue.
	//
	// https://en.wikipedia.org/wiki/Priority_queue
	// https://en.wikipedia.org/wiki/Monotone_priority_queue
	// First, we will implement a heap based priority queue

	var distance = make(map[Node]int)

	// For each node in this set, the shortest distance from the source node is known
	var settled = make(map[Node]bool)

	// Each node in this set has been visited at least once.
	// The shortest distance is known when using edges we have traversed, which means
	// we may find a shorter distance still.
	var frontier PriorityQueue

	frontier.Push(start, 0)
	distance[start] = 0

	for frontier.Len() > 0 {
		f := frontier.Pop()
		settled[f] = true
		if f == end {
			return distance[f]
		}
		for _, neighbor := range g.Neighbors(f) {
			if settled[neighbor] {
				continue
			}
			d := distance[f] + g.Weight(f, neighbor)
			oldD, seen := distance[neighbor]
			if !seen {
				distance[neighbor] = d
				frontier.Push(neighbor, d)
			} else if d < oldD {
				frontier.Update(neighbor, d)
				distance[neighbor] = d
			}
		}
	}

	return -1
}
