package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
	}

	return graph
}

func check_pre_req(num_courses int, pre_req [][]int, queries [][]int) []bool {
	res := []bool{}
	is_reachable := make([][]bool, num_courses)
	for i := 0; i < num_courses; i++ {
		is_reachable[i] = make([]bool, num_courses)
	}
	graph := make_graph(pre_req)

	for i := 0; i < num_courses; i++ {
		q := []int{i}
		for len(q) > 0 {
			curr := q[0]
			q = q[1:]
			for j := 0; j < len(graph[curr]); j++ {
				if !is_reachable[i][graph[curr][j]] {
					is_reachable[i][graph[curr][j]] = true
					q = append(q, graph[curr][j])
				}
			}
		}
	}

	for i := 0; i < len(queries); i++ {
		res = append(res, is_reachable[queries[i][0]][queries[i][1]])
	}

	return res
}

func main() {
	fmt.Println(check_pre_req(2, [][]int{
		{1, 0},
	}, [][]int{
		{0, 1}, {1, 0},
	}))
}
