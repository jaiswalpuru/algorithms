package main

import "fmt"

/*
	This topological sort is using dfs.
	Any node can be the start node.
*/

func dfs(u int, visited *[]bool, order *[]int, g map[int][]int) {
	(*visited)[u] = true
	for i := 0; i < len(g[u]); i++ {
		if !(*visited)[g[u][i]] {
			dfs(g[u][i], visited, order, g)
		}
	}
	*order = append(*order, u)
}

func topological_sort(g map[int][]int, n int) []int {
	ordering := make([]int, n)
	visited := make([]bool, n)
	j := n - 1
	for i := 0; i < n; i++ {
		if !visited[i] {
			order := []int{}
			dfs(i, &visited, &order, g)
			for k := 0; k < len(order); k++ {
				ordering[j] = order[k]
				j--
			}
		}
	}
	return ordering
}

func main() {
	graph := map[int][]int{
		5: {2, 0},
		4: {0, 1},
		2: {3},
		3: {1},
	}
	fmt.Println(topological_sort(graph, 6))
}
