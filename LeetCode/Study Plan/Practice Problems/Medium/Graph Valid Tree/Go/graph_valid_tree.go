package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

func graph_valid_tree(n int, edges [][]int) bool {

	graph := make_graph(edges)

	visited := make(map[int]int)

	q := []int{0}
	visited[0] = -1

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for i := 0; i < len(graph[curr]); i++ {
			if visited[curr] == graph[curr][i] {
				continue
			}
			if _, ok := visited[graph[curr][i]]; ok {
				return false
			}
			q = append(q, graph[curr][i])
			visited[graph[curr][i]] = curr
		}
	}
	return len(visited) == n
}

func main() {
	fmt.Println(graph_valid_tree(5, [][]int{
		{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4},
	}))
}
