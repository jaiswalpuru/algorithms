package main

import "fmt"

var (
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

func bfs(edges [][]int, n int) {
	if n == 0 {
		return
	}
	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = WHITE
	}
	q := []int{0}
	color[0] = GRAY
	graph := make_graph(edges)
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for v := 0; v < len(graph[u]); v++ {
			if color[graph[u][v]] == WHITE {
				color[graph[u][v]] = GRAY
				q = append(q, graph[u][v])
			}
		}
		color[u] = BLACK
		fmt.Print(u, " ")
	}
	fmt.Println()
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
	bfs(edges, 5)
}
