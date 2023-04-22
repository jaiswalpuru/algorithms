package main

import (
	"fmt"
)

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func num_nodes_subtree_with_same_label(n int, edges [][]int, labels string) []int {
	g := make_graph(edges)
	l := []byte(labels)
	res := []int{}
	visited := make(map[int]bool)
	cnt := make([][26]int, n)
	dfs(g, l, 0, &visited, &cnt)
	for i := 0; i < n; i++ {
		res = append(res, cnt[i][l[i]-'a'])
	}
	return res
}

func dfs(g map[int][]int, l []byte, u int, visited *map[int]bool, count *[][26]int) {
	(*count)[u][l[u]-'a'] = 1
	(*visited)[u] = true
	for i := 0; i < len(g[u]); i++ {
		v := g[u][i]
		if !(*visited)[v] {
			dfs(g, l, v, visited, count)
			for j := 0; j < 26; j++ {
				(*count)[u][j] += (*count)[v][j]
			}
		}
	}
}

func main() {
	fmt.Println(num_nodes_subtree_with_same_label(4, [][]int{
		{0, 1}, {1, 2}, {0, 3},
	}, "bbbb"))
}
