package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
	}
	return graph
}

func route_between_nodes(edges [][]int, src, dst int) bool {
	graph := make_graph(edges)
	visited := make(map[int]bool)
	q := []int{src}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr == dst {
			return true
		}
		if visited[curr] {
			continue
		}
		visited[curr] = true
		for i := 0; i < len(graph[curr]); i++ {
			if !visited[graph[curr][i]] {
				q = append(q, graph[curr][i])
			}
		}
	}
	return false
}

func main() {
	fmt.Println(route_between_nodes([][]int{
		{1, 2}, {2, 3},
	}, 1, 3))
	fmt.Println(route_between_nodes([][]int{
		{1, 2}, {2, 3},
	}, 1, 4))
}
