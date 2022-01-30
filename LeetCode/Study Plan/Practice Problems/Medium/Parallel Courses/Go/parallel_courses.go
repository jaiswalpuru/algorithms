package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
	}
	return g
}

func paralled_courses(n int, edges [][]int) int {
	indegree := make(map[int]int)
	m := len(edges)
	graph := make_graph(edges)

	for i := 0; i < m; i++ {
		indegree[edges[i][1]]++
	}

	nodes := []int{}

	for i := 1; i <= n; i++ {
		if indegree[i] == 0 {
			nodes = append(nodes, i)
		}
	}

	visited := make(map[int]bool)
	cnt := 0
	visit := 0
	for len(nodes) > 0 {
		n_size := len(nodes)
		cnt++
		for j := 0; j < n_size; j++ {
			visit++
			curr := nodes[0]
			nodes = nodes[1:]
			visited[curr] = true
			for i := 0; i < len(graph[curr]); i++ {
				indegree[graph[curr][i]]--
				if indegree[graph[curr][i]] == 0 {
					nodes = append(nodes, graph[curr][i])
				}
			}
		}
	}

	if visit == n {
		return cnt
	}
	return -1
}

func main() {
	fmt.Println(paralled_courses(3, [][]int{
		{1, 2}, {2, 3}, {3, 1},
	}))
}
