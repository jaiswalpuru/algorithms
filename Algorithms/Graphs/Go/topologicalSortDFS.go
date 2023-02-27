package main

import "fmt"

/*
	This topological sort is using dfs.
	Any node can be the start node.
*/

const (
	WHITE = 1
	GRAY  = 2
	BLACK = 3
)

func dfs(u int, visited *[]int, order *[]int, g map[int][]int, cycle *bool) {
	(*visited)[u] = GRAY
	for i := 0; i < len(g[u]); i++ {
		if (*visited)[g[u][i]] == WHITE {
			dfs(g[u][i], visited, order, g, cycle)
		} else if (*visited)[g[u][i]] == GRAY {
			*cycle = true
		}
	}
	*order = append(*order, u)
	(*visited)[u] = BLACK
}

func topological_sort(g map[int][]int, n int) []int {
	ordering := make([]int, n)
	visited := make([]int, n)
	for i := 0; i < n; i++ {
		visited[i] = WHITE
	}

	j := n - 1
	cycle := false
	for i := 0; i < n; i++ {
		if visited[i] == WHITE {
			order := []int{}
			dfs(i, &visited, &order, g, &cycle)
			if cycle {
				fmt.Println("Not a DAG")
				return nil
			}
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
