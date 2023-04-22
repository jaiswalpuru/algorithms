package main

import (
	"fmt"
)

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func loud_and_rich(richer [][]int, quiet []int) []int {
	graph := make_graph(richer)
	n := len(quiet)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = -1
	}

	for i := 0; i < n; i++ {
		dfs(graph, i, &res, quiet)
	}
	return res
}

func dfs(graph map[int][]int, node int, res *[]int, quiet []int) int {
	if (*res)[node] == -1 {
		(*res)[node] = node
		for i := 0; i < len(graph[node]); i++ {
			end := dfs(graph, graph[node][i], res, quiet)
			if quiet[end] < quiet[(*res)[node]] {
				(*res)[node] = end
			}
		}
	}
	return (*res)[node]
}

func main() {
	fmt.Println(loud_and_rich([][]int{
		{1, 0}, {2, 1}, {3, 1}, {3, 7}, {4, 3}, {5, 3}, {6, 3},
	}, []int{3, 2, 5, 4, 6, 1, 7, 0}))
	fmt.Println(loud_and_rich([][]int{
		{0, 1}, {1, 2},
	}, []int{1, 0, 2}))
}
