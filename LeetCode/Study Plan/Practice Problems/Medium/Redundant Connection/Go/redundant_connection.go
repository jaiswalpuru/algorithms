package main

import "fmt"

func redundant_connection(edges [][]int) []int {
	graph := make(map[int][]int)
	n := len(edges)
	visited := make(map[int]bool)

	for i := 0; i < n; i++ {
		visited = make(map[int]bool)
		if !(len(graph[edges[i][0]]) == 0) && !(len(graph[edges[i][1]]) == 0) && dfs(graph, edges[i][0], edges[i][1], &visited) {
			return []int{edges[i][0], edges[i][1]}
		}
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return []int{}
}

func dfs(graph map[int][]int, src, target int, visited *map[int]bool) bool {
	if _, ok := (*visited)[src]; !ok {
		(*visited)[src] = true
		if src == target {
			return true
		}
		for i := 0; i < len(graph[src]); i++ {
			if dfs(graph, graph[src][i], target, visited) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(redundant_connection([][]int{
		{1, 2}, {1, 3}, {2, 3},
	}))
}
