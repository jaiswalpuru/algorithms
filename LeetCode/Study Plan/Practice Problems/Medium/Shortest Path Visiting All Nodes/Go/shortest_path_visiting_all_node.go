package main

import (
	"fmt"
)

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[i] = append(graph[i], edges[i]...)
	}
	return graph
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func shortest_path_visiting_all_nodes(graph [][]int) int {
	if len(graph) == 1 {
		return 0
	}
	g := make_graph(graph)
	n := len(graph)
	end_mask := (1 << n) - 1

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, end_mask)
	}

	q := [][]int{}

	for i := 0; i < n; i++ {
		q = append(q, []int{i, 1 << i})
		visited[i][1<<i] = true
	}

	steps := 0
	for len(q) > 0 {
		m := len(q)
		for i := 0; i < m; i++ {
			curr := q[0]
			q = q[1:]
			node := curr[0]
			mask := curr[1]
			for v := 0; v < len(g[node]); v++ {
				next_mask := mask | (1 << g[node][v])
				if next_mask == end_mask {
					return steps + 1
				}
				if !visited[g[node][v]][next_mask] {
					visited[g[node][v]][next_mask] = true
					q = append(q, []int{g[node][v], next_mask})
				}
			}
		}
		steps++
	}
	return -1
}

func main() {
	fmt.Println(shortest_path_visiting_all_nodes([][]int{
		{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2},
	}))
}

// func shortest_path_visiting_all_nodes_using_dfs(edges [][]int) int {
// 	graph := make_graph(edges)
// 	n := len(graph)

// 	ending_mask := (1 << n) - 1
// 	cache := make([][]int, n+1)
// 	for i := 0; i <= n; i++ {
// 		cache[i] = make([]int, ending_mask+1)
// 	}

// 	optimal_path := math.MaxInt64
// 	for node := 0; node < n; node++ {
// 		optimal_path = min(optimal_path, dfs(graph, &cache, node, ending_mask))
// 	}

// 	return optimal_path
// }

// func dfs(graph map[int][]int, cache *[][]int, node, ending_mask int) int { // refer LC solutions
// 	if (*cache)[node][ending_mask] != 0 {
// 		return (*cache)[node][ending_mask]
// 	}

// 	if ending_mask&(ending_mask-1) == 0 {
// 		// when only has a single mask
// 		return 0
// 	}

// 	(*cache)[node][ending_mask] = math.MaxInt64 - 1
// 	for i := 0; i < len(graph[node]); i++ {
// 		if ending_mask&(1<<graph[node][i]) != 0 {
// 			already_visited := dfs(graph, cache, graph[node][i], ending_mask)
// 			not_visited := 1 + dfs(graph, cache, graph[node][i], ending_mask^(1<<node))
// 			best_option := min(already_visited, not_visited)
// 			(*cache)[node][ending_mask] = min(1+best_option, (*cache)[node][ending_mask])
// 		}
// 	}
// 	return (*cache)[node][ending_mask]
// }
