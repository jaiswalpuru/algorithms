package main

import "fmt"

func course_schedule_two(n int, pre_req [][]int) []int {
	g := make_graph(pre_req)
	indegree := make([]int, n)
	for i := 0; i < len(pre_req); i++ {
		indegree[pre_req[i][0]]++
	}
	q := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	res := []int{}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		res = append(res, curr)
		for i := 0; i < len(g[curr]); i++ {
			indegree[g[curr][i]]--
			if indegree[g[curr][i]] == 0 {
				q = append(q, g[curr][i])
			}
		}
	}
	for i := 0; i < n; i++ {
		if indegree[i] != 0 {
			return nil
		}
	}
	return res
}

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

func main() {
	fmt.Println(course_schedule_two(2, [][]int{{1, 0}}))
}
