package main

import (
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func New(val int) *Node { return &Node{Val: val} }

func clone_graph(g *Node) *Node {
	if g == nil {
		return nil
	}
	visited, new_node := make(map[int]*Node), &Node{Val: g.Val, Neighbors: nil}
	visited[g.Val] = new_node
	q := []*Node{g}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		n := len(curr.Neighbors)
		for i := 0; i < n; i++ {
			if _, ok := visited[curr.Neighbors[i].Val]; !ok {
				visited[curr.Neighbors[i].Val] = &Node{Val: curr.Neighbors[i].Val, Neighbors: nil}
				q = append(q, curr.Neighbors[i])
			}
			visited[curr.Val].Neighbors = append(visited[curr.Val].Neighbors, visited[curr.Neighbors[i].Val])
		}
	}
	return new_node
}

func main() {
	node_1 := New(1)
	node_2 := New(2)
	node_3 := New(3)
	node_4 := New(4)
	node_1.Neighbors = append(node_1.Neighbors, []*Node{node_2, node_4}...)
	node_2.Neighbors = append(node_2.Neighbors, []*Node{node_1, node_3}...)
	node_3.Neighbors = append(node_3.Neighbors, []*Node{node_2, node_4}...)
	node_4.Neighbors = append(node_4.Neighbors, []*Node{node_1, node_3}...)
	fmt.Println(clone_graph(node_1))
}
