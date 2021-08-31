/*
Given an undirected graph, how to check if there is a cycle in the graph?
*/

package main

import "fmt"

func _is_cyclic(v int, visited []bool, parent int, g [][]int) bool {
	visited[v] = true

	n := len(g[v])

	for i := 0; i < n; i++ {
		if !visited[g[v][i]] {
			if _is_cyclic(g[v][i], visited, v, g) {
				return true
			}
		} else if g[v][i] != parent {
			return true
		}
	}
	return false
}

func is_cyclic(g [][]int) bool {
	nodes := len(g)
	visited := make([]bool, nodes)

	for i := 0; i < nodes; i++ {
		if !visited[i] {
			if _is_cyclic(i, visited, -1, g) {
				return true
			}
		}
	}
	return false
}

func main() {
	g1 := make([][]int, 5) //Graph with 5 nodes (0-4)

	for i := 0; i < 5; i++ {
		g1[i] = make([]int, 0)
	}

	g1[1] = append(g1[1], 0)
	g1[0] = append(g1[0], 2)
	g1[2] = append(g1[2], 1)
	g1[0] = append(g1[0], 3)
	g1[3] = append(g1[3], 4)

	fmt.Println(is_cyclic(g1))

	g2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		g2[i] = make([]int, 0)
	}
	g2[0] = append(g2[0], 1)
	g2[1] = append(g2[1], 2)
	fmt.Println(is_cyclic(g2))

}
