package main

import "fmt"

/*
	Bridges : Any edge on whose removal
	disconnects the graph.
	This problem can be solved using Tarjan's
	algorithm.

	There is another approach to solve this problem
	using dfs + rank
*/

func bridges(g map[int][]int, n int) [][]int {
	bridge := [][]int{}
	low := make([]int, n)
	disc := make([]int, n)
	for i := 0; i < n; i++ {
		disc[i] = -1
	}
	time := 0
	dfs(g, 0, -1, &low, &disc, &time, &bridge)
	return bridge
}

func dfs(g map[int][]int, u, parent int, low, disc *[]int, time *int, bridge *[][]int) {
	*time++
	(*low)[u] = *time
	(*disc)[u] = *time
	for i := 0; i < len(g[u]); i++ {
		v := g[u][i]
		if v == parent {
			continue
		}
		if (*disc)[v] == -1 {
			dfs(g, v, u, low, disc, time, bridge)
			(*low)[u] = min((*low)[u], (*low)[v])
			if (*disc)[u] < (*low)[v] {
				*bridge = append(*bridge, []int{u, v})
			}
		}
		(*low)[u] = min((*low)[u], (*disc)[v])
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	g1 := map[int][]int{
		0: {1, 2},
		1: {0, 2, 3},
		2: {1, 0},
		3: {1},
	}
	fmt.Println(bridges(g1, 4))
}
