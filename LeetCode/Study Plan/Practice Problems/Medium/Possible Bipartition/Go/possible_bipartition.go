package main

import (
	"fmt"
)

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

func possible_bipartition(n int, dislikes [][]int) bool {
	graph := make_graph(dislikes)
	visited := make(map[int]int)

	for i := 1; i <= n; i++ {
		if visited[i] != 0 {
			continue
		}
		visited[i] = 1
		q := []int{i}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]
			for j := 0; j < len(graph[curr]); j++ {
				if visited[graph[curr][j]] == visited[curr] {
					return false
				}
				if visited[graph[curr][j]] == 0 {
					q = append(q, graph[curr][j])
					if visited[curr] == 1 {
						visited[graph[curr][j]] = 2
					} else {
						visited[graph[curr][j]] = 1
					}
				}
			}
		}
	}
	return true
}

func main() {
	fmt.Println(possible_bipartition(3, [][]int{
		{1, 2}, {1, 3}, {2, 3},
	}))
}
