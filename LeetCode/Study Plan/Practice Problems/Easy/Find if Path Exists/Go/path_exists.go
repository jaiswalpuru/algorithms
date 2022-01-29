package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}

	return graph
}

func path_exists(edges [][]int, n int, src, dst int) bool {
	graph := make_graph(edges)

	q := make([]int, 0)
	q = append(q, src)
	visited := make([]bool, n)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		visited[curr] = true
		if curr == dst {
			return true
		}
		for j := 0; j < len(graph[curr]); j++ {
			if !visited[graph[curr][j]] {
				q = append(q, graph[curr][j])
			}
		}
	}
	return false
}

func main() {
	fmt.Println(path_exists([][]int{
		{0, 1}, {1, 2}, {2, 0},
	}, 3, 0, 2))
}
