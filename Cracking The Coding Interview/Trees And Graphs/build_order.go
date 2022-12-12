package main

import "fmt"

func build_order(vertices []int, edges [][]int) []int {
	n := len(vertices)
	in_degree := make([]int, n+1)
	graph := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		in_degree[edges[i][1]]++
	}
	q := []int{}
	for i := 1; i < len(in_degree); i++ {
		if in_degree[i] == 0 {
			q = append(q, i)
		}
	}
	res := []int{}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		res = append(res, curr)
		for i := 0; i < len(graph[curr]); i++ {
			in_degree[graph[curr][i]]--
			if in_degree[graph[curr][i]] == 0 {
				q = append(q, graph[curr][i])
			}
		}
	}
	return res
}

func main() {
	fmt.Println(build_order([]int{1, 2, 3, 4, 5, 6}, [][]int{{1, 4}, {6, 4}, {2, 4}, {6, 1}, {4, 3}}))
}
