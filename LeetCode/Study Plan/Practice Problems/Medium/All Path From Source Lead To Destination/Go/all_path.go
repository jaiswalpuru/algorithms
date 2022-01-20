package main

import (
	"fmt"
)

func make_graph(edges [][]int) map[int][]int {
	vertices := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		vertices[edges[i][0]] = append(vertices[edges[i][0]], edges[i][1])
	}
	return vertices
}

func dfs(src, dst int, visited *[]bool, edge_map map[int][]int) bool {
	if len(edge_map[src]) == 0 {
		if src == dst {
			return true
		}
		return false
	}

	(*visited)[src] = true
	for i := 0; i < len(edge_map[src]); i++ {
		if (*visited)[edge_map[src][i]] || !dfs(edge_map[src][i], dst, visited, edge_map) {
			return false
		}
		(*visited)[edge_map[src][i]] = false
	}
	return true
}

func all_path(n int, edges [][]int, src int, dest int) bool {
	edge_map := make_graph(edges)

	visited := make([]bool, n)

	return dfs(src, dest, &visited, edge_map)
}

func main() {
	fmt.Println(all_path(4,
		[][]int{
			{0, 1},
			{0, 2},
			{1, 3},
			{2, 3},
		}, 0, 3))
}
