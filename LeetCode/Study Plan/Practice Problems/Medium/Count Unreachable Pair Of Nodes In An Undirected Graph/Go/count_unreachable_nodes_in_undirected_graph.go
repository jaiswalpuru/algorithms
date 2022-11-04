package main

import "fmt"

func count_unreachable_nodes_in_undirected_graph(n int, g [][]int) int {
	visited := make([]bool, n)
	graph := make_graph(g)
	total_conn := (n * (n - 1)) / 2
	for i := 0; i < n; i++ {
		if !visited[i] {
			cnt := 0
			dfs(i, graph, &cnt, &visited)
			total_conn -= (cnt * (cnt - 1)) / 2
		}
	}
	return total_conn
}

func dfs(curr int, graph map[int][]int, cnt *int, visited *[]bool) {
	(*visited)[curr] = true
	*cnt++
	for i := 0; i < len(graph[curr]); i++ {
		if !(*visited)[graph[curr][i]] {
			dfs(graph[curr][i], graph, cnt, visited)
		}
	}
}

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func main() {
	fmt.Println(count_unreachable_nodes_in_undirected_graph(3, [][]int{
		{0, 1}, {0, 2}, {1, 2},
	}))
}
