package main

import "fmt"

func parallel_courses(n int, relations [][]int, time []int) int {
	indegree := make([]int, n)
	g := make_graph(relations)
	for i := 0; i < len(relations); i++ {
		indegree[relations[i][1]-1]++
	}
	q := []int{}
	dis := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
			dis[i] = time[i]
		}
	}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for i := 0; i < len(g[curr]); i++ {
			indegree[g[curr][i]]--
			if indegree[g[curr][i]] == 0 {
				q = append(q, g[curr][i])
			}
			dis[g[curr][i]] = max(dis[g[curr][i]], time[g[curr][i]]+dis[curr])
		}
	}
	for i := 0; i < len(dis); i++ {
		res = max(res, dis[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func make_graph(edges [][]int) map[int][]int {
	n := len(edges)
	g := make(map[int][]int)
	for i := 0; i < n; i++ {
		g[edges[i][0]-1] = append(g[edges[i][0]-1], edges[i][1]-1)
	}
	return g
}

func main() {
	fmt.Println(parallel_courses(3, [][]int{
		{1, 3}, {2, 3},
	}, []int{3, 2, 5}))
}
