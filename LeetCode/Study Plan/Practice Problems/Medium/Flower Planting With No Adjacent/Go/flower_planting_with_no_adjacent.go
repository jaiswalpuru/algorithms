package main

import (
	"fmt"
)

func make_graph(edges [][]int) map[int][]int {
	n := len(edges)
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func garden_no_adj(n int, edges [][]int) []int {
	graph := make_graph(edges)
	color := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if color[i] != 0 {
			continue
		}
		q := []int{i}
		color[i] = 1
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]
			c := color[curr]
			for j := 0; j < len(graph[curr]); j++ {
				v := graph[curr][j]
				if color[v] == 0 {
					if c == 4 {
						color[v] = 1
					} else {
						color[v] = c + 1
					}
					q = append(q, v)
				} else if color[v] == c {
					if c == 4 {
						color[v] = 1
					} else {
						color[v] = c + 1
					}
				}
			}
		}
	}
	return color[1:]
}
func main() {
	fmt.Println(garden_no_adj(4, [][]int{
		{3, 4}, {4, 2}, {3, 2}, {1, 3},
	}))
}
