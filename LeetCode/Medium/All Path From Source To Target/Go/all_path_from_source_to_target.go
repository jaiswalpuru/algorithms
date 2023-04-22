package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)

	n := len(edges)
	for i := 0; i < n; i++ {
		graph[i] = append(graph[i], edges[i]...)
	}

	return graph
}

//----------------------------using bfs--------------------------------
func all_path_from_source_to_target(graph [][]int) [][]int {
	g := make_graph(graph)
	n := len(graph) - 1
	res := [][]int{}
	q := [][]int{}
	q = append(q, []int{0})
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		last := curr[len(curr)-1]
		for i := 0; i < len(g[last]); i++ {
			temp := curr
			temp = append(temp, g[last][i])
			if g[last][i] == n {
				res = append(res, append([]int{}, temp...))
			} else {
				q = append(q, append([]int{}, temp...))
			}
		}
	}
	return res
}

//-----------------------------using bfs---------------------------------

//--------------------------------using dfs-------------------------------
func dfs(src int, graph map[int][]int, res *[][]int, temp []int, n int) {
	if src == n {
		temp = append(temp, src)
		*res = append(*res, append([]int{}, temp...))
		return
	}

	for i := 0; i < len(graph[src]); i++ {
		temp = append(temp, src)
		dfs(graph[src][i], graph, res, temp, n)
		temp = temp[:len(temp)-1]
	}
}

func all_path_from_source_to_target_dfs(graph [][]int) [][]int {
	res := [][]int{}
	g := make_graph(graph)
	n := len(graph) - 1
	dfs(0, g, &res, []int{}, n)
	return res
}

//--------------------------------using dfs-------------------------------

func main() {
	fmt.Println(all_path_from_source_to_target_dfs([][]int{
		{1, 2}, {3}, {3}, {},
	}))
}
