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

type Node struct {
	v           int
	disc_time   int
	finish_time int
}

func make_graph(edges [][]int) map[int][]Node {
	graph := make(map[int][]Node)
	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], Node{v: edges[i][1]})
		graph[edges[i][1]] = append(graph[edges[i][1]], Node{v: edges[i][0]})
	}
	return graph
}

func dfs(edges [][]int, n int) {
	graph := make_graph(edges)
	color := make([]int, n)
	for i := 0; i < n; i++ {
		color[i] = WHITE
	}
	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i].v = i
	}
	time := 0
	for i := 0; i < n; i++ {
		if color[i] == WHITE {
			_dfs(nodes[i], &nodes, graph, &color, &time)
		}
	}
	fmt.Printf("%+v\n", nodes)
}

func _dfs(u Node, nodes *[]Node, g map[int][]Node, color *[]int, time *int) {
	*time += 1
	(*color)[u.v] = GRAY
	u.disc_time = *time
	for v := 0; v < len(g[u.v]); v++ {
		if (*color)[g[u.v][v].v] == WHITE {
			_dfs(g[u.v][v], nodes, g, color, time)
		}
	}
	*time += 1
	u.finish_time = *time
	(*nodes)[u.v] = u
	(*color)[u.v] = BLACK
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
