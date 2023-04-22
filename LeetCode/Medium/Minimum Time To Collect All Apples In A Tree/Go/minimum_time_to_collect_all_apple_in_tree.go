package main

import "fmt"

//-------------------using dfs-----------------------------------
func minimum_time_to_collect_apples_in_tree(n int, edges [][]int, has_apple []bool) int {
	g := make_graph(edges)
	return dfs(g, 0, -1, has_apple) * 2
}

func dfs(g map[int][]int, src, parent int, has_apple []bool) int {
	sum := 0
	for i := 0; i < len(g[src]); i++ {
		if g[src][i] != parent {
			sum += dfs(g, g[src][i], src, has_apple)
		}
	}
	if src != 0 && (sum > 0 || has_apple[src]) {
		return sum + 1
	}
	return sum
}

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

//-------------------using dfs-----------------------------------

func main() {
	fmt.Println(minimum_time_to_collect_apples_in_tree(7, [][]int{
		{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6},
	}, []bool{false, false, true, false, true, true, false}))
}
