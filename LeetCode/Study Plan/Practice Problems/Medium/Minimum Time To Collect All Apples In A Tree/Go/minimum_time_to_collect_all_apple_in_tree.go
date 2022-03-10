package main

import "fmt"

func minimum_time_to_collect_apples_in_tree(n int, edges [][]int, has_apple []bool) int {
	graph := make(map[int][]int)
	m := len(edges)
	for i := 0; i < m; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}

	return dfs(graph, 0, -1, has_apple) * 2 //basically post order
}

func dfs(graph map[int][]int, curr, parent int, has_apple []bool) int {
	sum := 0
	for i := 0; i < len(graph[curr]); i++ {
		if graph[curr][i] == parent {
			continue
		}
		sum += dfs(graph, graph[curr][i], curr, has_apple)
	}

	if curr != 0 && (sum > 0 || has_apple[curr]) {
		return sum + 1
	}
	return sum
}

func main() {
	fmt.Println(minimum_time_to_collect_apples_in_tree(7, [][]int{
		{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6},
	}, []bool{false, false, true, false, true, true, false}))
}
