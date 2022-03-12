package main

import "fmt"

func make_graph(manager []int) map[int][]int {
	graph := make(map[int][]int)
	m := len(manager)
	for i := 0; i < m; i++ {
		graph[manager[i]] = append(graph[manager[i]], i)
	}
	return graph
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func time_to_inform_all_emp(n int, head_emp int, manager []int, inform_time []int) int {
	graph := make_graph(manager)

	type P struct{ val, dis int }

	q := []P{{head_emp, 0}}
	time := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		u := curr.val
		d := curr.dis
		time = max(time, d)
		for i := 0; i < len(graph[u]); i++ {
			q = append(q, P{graph[u][i], d + inform_time[u]})
		}
	}
	return time
}

func main() {
	fmt.Println(time_to_inform_all_emp(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}))
}
