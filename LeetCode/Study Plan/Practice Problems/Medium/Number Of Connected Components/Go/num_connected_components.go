package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	n := len(edges)
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func dfs(start int, visited *[]bool, graph map[int][]int) {
	(*visited)[start] = true
	for i := 0; i < len(graph[start]); i++ {
		if !(*visited)[graph[start][i]] {
			dfs(graph[start][i], visited, graph)
		}
	}
}

func num_connected_components(n int, edges [][]int) int {
	components := 0

	graph := make_graph(edges)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			components++
			dfs(i, &visited, graph)
		}
	}

	return components
}

func main() {
	fmt.Println(num_connected_components(2, [][]int{
		{1, 0},
	}))
}
