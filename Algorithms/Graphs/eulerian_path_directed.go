package main

import "fmt"

//Refer to the videos by William Fiset.

/*
	Rules :
		Directed graph :
		For eulerian circuit :
			- Every vertex has equal in and out
			degree.
		For eulerian path in directed graph :
			- Atmost one vertex has out-in=1
			and atmost one vertex has in-out=1
			and all other vertex have equal in and
			out degrees.
*/

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
	}
	return graph
}

func in_out_degree(g map[int][]int, n int) (in []int, out []int) {
	in = make([]int, n)
	out = make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = len(g[i])
	}
	for _, v := range g {
		for i := 0; i < len(v); i++ {
			in[v[i]]++
		}
	}
	return
}

func eulerian_path_exists(in, out []int) bool {
	start, end := 0, 0
	for i := 0; i < len(in); i++ {
		if in[i]-out[i] > 1 || out[i]-in[i] > 1 {
			return false
		}
		if out[i]-in[i] == 1 {
			start++
		}
		if in[i]-out[i] == 1 {
			end++
		}
	}
	return (start == 0 && end == 0) || (start == 1 && end == 1)
}

func find_start_node(in, out []int) int {
	start := 0
	for i := 0; i < len(in); i++ {
		if out[i]-in[i] == 1 {
			return i
		}
		if out[i] > 0 {
			start = i
		}
	}
	return start
}

func eulerian_path(edges [][]int, n int) []int {
	graph := make_graph(edges)
	in, out := in_out_degree(graph, n)
	if !eulerian_path_exists(in, out) {
		return nil
	}
	start := find_start_node(in, out)
	path := []int{}
	dfs(start, out, graph, &path)
	for i := 0; i < len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
	}
	return path
}

func dfs(start int, out []int, graph map[int][]int, path *[]int) {
	for out[start] != 0 {
		out[start]--
		v := graph[start][out[start]]
		dfs(v, out, graph, path)
	}
	*path = append(*path, start)
}

func main() {
	edges := [][]int{
		{2, 2}, {1, 2}, {2, 4},
		{2, 4}, {4, 6}, {5, 6},
		{3, 5}, {6, 3}, {4, 3},
		{3, 2}, {1, 3}, {3, 1},
	}
	fmt.Println(eulerian_path(edges, 7))
}
