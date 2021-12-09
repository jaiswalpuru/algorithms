/*
Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

class Node {
    public int val;
    public List<Node> neighbors;
}

Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed).
For example, the first node with val == 1, the second node with val == 2, and so on.
The graph is represented in the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph.
Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the
given node as a reference to the cloned graph.

Example 1:
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

Example 2:
Input: adjList = [[]]
Output: [[]]
Explanation: Note that the input contains one empty list. The graph consists of only one node with val = 1 and it does not have any neighbors.

Example 3:
Input: adjList = []
Output: []
Explanation: This an empty graph, it does not have any nodes.

Example 4:
Input: adjList = [[2],[1]]
Output: [[2],[1]]
*/

package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func New(val int) *Node { return &Node{Val: val, Neighbors: nil} }

func create_graph(arr [][]int) *Node {
	node := New(1)
	node.Neighbors = append(node.Neighbors, New(2))
	node.Neighbors = append(node.Neighbors, New(4))
	node.Neighbors[0].Neighbors = append(node.Neighbors[0].Neighbors, node)
	node.Neighbors[0].Neighbors = append(node.Neighbors[0].Neighbors, New(3))
	node.Neighbors[0].Neighbors[1].Neighbors = append(node.Neighbors[0].Neighbors[1].Neighbors, node.Neighbors...)
	node.Neighbors[1].Neighbors = append(node.Neighbors[1].Neighbors, node)
	node.Neighbors[1].Neighbors = append(node.Neighbors[1].Neighbors, node.Neighbors[0].Neighbors[1])
	return node
}

func clone_graph(graph *Node) *Node {
	if graph == nil {
		return nil
	}

	q := []*Node{graph}
	visited, new_node := make(map[int]*Node), &Node{Val: graph.Val, Neighbors: nil}
	visited[graph.Val] = new_node
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, v := range curr.Neighbors {
			if _, ok := visited[v.Val]; !ok {
				visited[v.Val] = &Node{v.Val, nil}
				q = append(q, v)
			}
			visited[curr.Val].Neighbors = append(visited[curr.Val].Neighbors, visited[v.Val])
		}
	}
	return new_node
}

func main() {
	graph := create_graph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	fmt.Println(graph)
	fmt.Println(graph.Neighbors[0])
	fmt.Println(graph.Neighbors[1])
	fmt.Println(graph.Neighbors[0].Neighbors[0])
	fmt.Println(graph.Neighbors[0].Neighbors[1])
	fmt.Println(graph.Neighbors[0].Neighbors[1].Neighbors[0])
	fmt.Println(graph.Neighbors[0].Neighbors[1].Neighbors[1])
	graph_clone := clone_graph(graph)
	fmt.Println(graph_clone)
	fmt.Println(graph_clone.Neighbors[0])
	fmt.Println(graph_clone.Neighbors[1])
	fmt.Println(graph_clone.Neighbors[0].Neighbors[0])
	fmt.Println(graph_clone.Neighbors[0].Neighbors[1])
	fmt.Println(graph_clone.Neighbors[0].Neighbors[1].Neighbors[0])
	fmt.Println(graph_clone.Neighbors[0].Neighbors[1].Neighbors[1])
}
