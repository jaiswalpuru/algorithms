package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	n := len(edges)
	g := make(map[int][]int)
	for i := 0; i < n; i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

func make_connections(n int, connections [][]int) int {
	if n-1 > len(connections) {
		return -1
	}

	graph := make_graph(connections)
	visited := make([]bool, n)
	cnt := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			cnt++
			dfs(i, graph, &visited)
		}
	}
	return cnt - 1
}

func dfs(node int, graph map[int][]int, visited *[]bool) {
	(*visited)[node] = true
	for i := 0; i < len(graph[node]); i++ {
		if !(*visited)[graph[node][i]] {
			dfs(graph[node][i], graph, visited)
		}
	}
}

func main() {
	fmt.Println(make_connections(12, [][]int{
		{1, 5}, {1, 7}, {1, 2}, {1, 4}, {3, 7}, {4, 7}, {3, 5}, {0, 6}, {0, 1}, {0, 4}, {2, 6}, {0, 3}, {0, 2},
	}))
	fmt.Println(make_connections(6, [][]int{
		{0, 1}, {0, 2}, {0, 3}, {1, 2},
	}))
}
