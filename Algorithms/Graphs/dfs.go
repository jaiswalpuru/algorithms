package main

import "fmt"

/*
	Here in the make graph function i am considering an unidirected
	graph, the reason i am appending for both the u, v of every edge
	if it had been an directed graph just appending the u -> v will
	work.
	In this solution I will follow dfs with coloring as this will help
	in solving processing of nodes problem later.
	Assuming n == len(edges) and the graph is connected
*/

const (
	WHITE = -1
	GRAY  = 1
	BLACK = 2
)

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func dfs(edges [][]int, n int) {
	graph := make_graph(edges)
	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = WHITE
	}
	for i := 0; i < n; i++ {
		if color[i] == WHITE {
			_dfs(i, graph, &color)
		}
	}
	fmt.Println()
}

func _dfs(u int, g map[int][]int, color *[]int) {
	(*color)[u] = GRAY
	fmt.Print(u, " ")
	for v := 0; v < len(g[u]); v++ {
		if (*color)[g[u][v]] == WHITE {
			_dfs(g[u][v], g, color)
		}
	}
	(*color)[u] = BLACK
}

/*
	0 - 1 - 2
	|	  /
	4 - 3
*/
func main() {
	edges := [][]int{
		{0, 1}, {0, 4}, {1, 2}, {4, 3}, {3, 2},
	}
	dfs(edges, 5)
}
