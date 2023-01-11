package main

import "fmt"

/*
	Kosaraju algorithm is for check whether a directed graph is a strongly
	connected component or not.
	A strongly connected graph is one in which starting from any
	vertex, it can visit any other vertex.
*/

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
	}
	return g
}

func graph_transpose(g map[int][]int) map[int][]int {
	transpose := make(map[int][]int)
	for k, v := range g {
		for i := 0; i < len(v); i++ {
			transpose[v[i]] = append(transpose[v[i]], k)
		}
	}
	return transpose
}

func dfs(g map[int][]int, u int, visited *[]bool) {
	(*visited)[u] = true
	for i := 0; i < len(g[u]); i++ {
		if !(*visited)[g[u][i]] {
			dfs(g, g[u][i], visited)
		}
	}
}

func is_strong_connected_component(edges [][]int, n int) bool {
	g := make_graph(edges)
	visited := make([]bool, n)
	//start from any vertex
	dfs(g, 0, &visited)
	for i := 0; i < n; i++ {
		if !visited[i] {
			return false
		}
	}
	transpose := graph_transpose(g)
	visited = make([]bool, n)
	//do the dfs from the same vertex as used above
	dfs(transpose, 0, &visited)
	for i := 0; i < n; i++ {
		if !visited[i] {
			return false
		}
	}
	return true
}

/*
	g1 :
		0 -> 1
		1 -> 2
		2 -> 1, 3, 4
		3 -> 0
		4 -> 2
	g2 :
		0 -> 1
		1 -> 2
		2 -> 3
*/
func main() {
	g1 := [][]int{
		{0, 1}, {1, 2}, {2, 1}, {2, 3}, {2, 4}, {3, 0}, {4, 2},
	}
	g2 := [][]int{
		{0, 1}, {1, 2}, {2, 3},
	}
	fmt.Println(is_strong_connected_component(g1, 5))
	fmt.Println(is_strong_connected_component(g2, 4))
}
