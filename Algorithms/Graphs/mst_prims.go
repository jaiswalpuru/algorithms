package main

import (
	"fmt"
	"log"
)

type minHeapNode interface {
	lessThan(minHeapNode) bool
}

type MinHeap struct {
	nodes []minHeapNode
}

//called after insertion of an element
func (mh *MinHeap) percolateUp(index int) {
	if index == 0 || index >= len(mh.nodes) {
		return
	}

	parentIndex := (index - 1) / 2
	if mh.nodes[index].lessThan(mh.nodes[parentIndex]) {
		temp := mh.nodes[parentIndex]
		mh.nodes[parentIndex] = mh.nodes[index]
		mh.nodes[index] = temp
		mh.percolateUp(parentIndex)
	}
}

//this is called when we delete an element
func (mh *MinHeap) percolateDown(index int) {
	if index*2+1 >= len(mh.nodes) { //not a single node stop sifting up
		return
	}
	childIndex := index*2 + 1
	childNode := mh.nodes[childIndex]

	//check if second child node is less than the first, if so then swap it
	if index*2+2 < len(mh.nodes) && mh.nodes[index*2+2].lessThan(childNode) {
		childIndex = index*2 + 2
		childNode = mh.nodes[childIndex]
	}

	//if childnode is less than the parent then swap it
	if childNode.lessThan(mh.nodes[index]) {
		temp := mh.nodes[index]
		mh.nodes[index] = mh.nodes[childIndex]
		mh.nodes[childIndex] = temp
		mh.percolateDown(childIndex)
	}
}

//insert an element into the heap
func (mh *MinHeap) Insert(node minHeapNode) {
	mh.nodes = append(mh.nodes, node)
	mh.percolateUp(len(mh.nodes) - 1)
}

//deletes the min value element from the heap
func (mh *MinHeap) Pop() (minHeapNode, error) { //deletes the node with the minimum value, it is the root node

	if len(mh.nodes) == 0 {
		return nil, fmt.Errorf("MinHeap is empty")
	}

	node := mh.nodes[0]
	mh.nodes[0] = mh.nodes[len(mh.nodes)-1] //swap the root node with the last node present in the heap
	mh.nodes = mh.nodes[:len(mh.nodes)-1]
	mh.percolateDown(0)
	return node, nil

}

//returns the minimum element or the top element from the min heap
func (mh MinHeap) Fetch() (minHeapNode, error) {

	if len(mh.nodes) == 0 {
		return nil, fmt.Errorf("MinHeap is empty")
	}

	return mh.nodes[0], nil
}

//reutrn the number of nodes in min heap
func (mh MinHeap) Size() int { return len(mh.nodes) }

//returns new min heap
func NewMinHeap(nodes []minHeapNode) *MinHeap {

	minHeap := &MinHeap{nodes: nodes}

	if len(minHeap.nodes) <= 1 {
		return minHeap
	}

	for i := len(minHeap.nodes) / 2; i >= 0; i-- {
		minHeap.percolateUp(i)
	}

	return minHeap
}

type Node interface{}

type edge struct {
	nodes  [2]Node
	weight int
}

func (e edge) String() string { return fmt.Sprintf("%s<-%d->%s", e.nodes[0], e.weight, e.nodes[1]) }

func (e edge) lessThan(mh minHeapNode) bool {
	edge, ok := mh.(*edge)
	if !ok {
		log.Fatal("Programming error invalid type assertion")
	}

	return e.weight < edge.weight
}

type Graph struct {
	edges map[Node]*MinHeap
}

func NewGraph() *Graph { return &Graph{edges: make(map[Node]*MinHeap)} }

func (g *Graph) Insert(nodeA, nodeB Node, weight int) {
	insert := func(node Node, edge *edge) {
		nodeEdges, ok := g.edges[node]
		if !ok {
			g.edges[node] = NewMinHeap([]minHeapNode{})
			nodeEdges = g.edges[node]
		}
		nodeEdges.Insert(edge)
	}

	edge := &edge{nodes: [2]Node{nodeA, nodeB},
		weight: weight,
	}

	insert(nodeA, edge)
	insert(nodeB, edge)
}

func (g Graph) Get(node Node) (*MinHeap, error) {
	nodeEdges, ok := g.edges[node]
	if !ok {
		return nil, fmt.Errorf("Node not found")
	}
	return nodeEdges, nil
}

func (g *Graph) IterNode(cb func(Node) bool) {
	for node := range g.edges {
		if shouldContinue := cb(node); !shouldContinue {
			break
		}
	}
}

func main() {
	fmt.Println()
}
