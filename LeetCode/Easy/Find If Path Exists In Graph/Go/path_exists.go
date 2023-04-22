package main

import "fmt"

//----------------------bfs----------------------
func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}

	return graph
}

func path_exists(edges [][]int, n int, src, dst int) bool {
	graph := make_graph(edges)

	q := make([]int, 0)
	q = append(q, src)
	visited := make([]bool, n)

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		visited[curr] = true
		if curr == dst {
			return true
		}
		for j := 0; j < len(graph[curr]); j++ {
			if !visited[graph[curr][j]] {
				q = append(q, graph[curr][j])
			}
		}
	}
	return false
}

//----------------------bfs----------------------

//----------------------dfs----------------------
func valid_path_dfs(n int, edges [][]int, src int, dst int) bool {
	visited := make(map[int]bool)
	graph := make_graph(edges)
	return dfs(src, dst, graph, &visited)
}

func dfs(src, target int, graph map[int][]int, visited *map[int]bool) bool {
	(*visited)[src] = true
	if src == target {
		return true
	}
	res := false
	for i := 0; i < len(graph[src]); i++ {
		if !(*visited)[graph[src][i]] {
			res = res || dfs(graph[src][i], target, graph, visited)
		}
	}
	return res
}

//----------------------dfs----------------------

func main() {
	fmt.Println(path_exists([][]int{
		{0, 1}, {1, 2}, {2, 0},
	}, 3, 0, 2))
}
