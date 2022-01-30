package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	n := len(edges)

	for i := 0; i < n; i++ {
		g[i] = append(g[i], edges[i]...)
	}
	return g
}

func keys_and_rooms(edges [][]int) bool {
	graph := make_graph(edges)
	fmt.Println(graph)
	n := len(edges)
	visited := make(map[int]bool)
	q := []int{0}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		visited[curr] = true
		for i := 0; i < len(graph[curr]); i++ {
			if !visited[graph[curr][i]] {
				q = append(q, graph[curr][i])
			}
		}
	}

	for i := 0; i < n; i++ {
		if visited[i] != true {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(keys_and_rooms([][]int{
		{1, 3}, {3, 0, 1}, {2}, {0},
	}))
}
