package main

import (
	"fmt"
)

func make_graph(n int, edges [][]int, labels string) map[int][]int {
	graph := make(map[int][]int)

	for i := 0; i < n-1; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func num_nodes_subtree_with_same_label(n int, edges [][]int, labels string) []int {
	graph := make_graph(n, edges, labels)
	cnt := make([][26]int, n)

	visited := make(map[int]bool)
	dfs(0, graph, visited, labels, cnt)
	res := []int{}
	for i := range cnt {
		res = append(res, cnt[i][labels[i]-'a'])
	}
	return res
}

func dfs(curr int, graph map[int][]int, visited map[int]bool, labels string, cnt [][26]int) {
	visited[curr] = true
	cnt[curr][labels[curr]-'a'] = 1
	for k := 0; k < len(graph[curr]); k++ {
		if !visited[graph[curr][k]] {
			dfs(graph[curr][k], graph, visited, labels, cnt)
			for i, val := range cnt[graph[curr][k]] {
				cnt[curr][i] += val
			}
		}
	}
}

func main() {
	fmt.Println(num_nodes_subtree_with_same_label(4, [][]int{
		{0, 1}, {1, 2}, {0, 3},
	}, "bbbb"))
}
