package main

import (
	"fmt"
)

func make_graph(edges [][]int) map[int][]int {
	n := len(edges)
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
	}
	return graph
}

func shortest_path_alternating(n int, red, blue [][]int) []int {
	r, b := make_graph(red), make_graph(blue)

	l := [2][]int{}
	l[0] = make([]int, n)
	l[1] = make([]int, n)

	for i := 1; i < n; i++ {
		l[0][i] = 2*n + 1
		l[1][i] = 2*n + 1
	}

	q := make([][2]int, 0)
	q = append(q, [2]int{0, 0})
	q = append(q, [2]int{0, 1})

	for len(q) > 0 {
		q_size := len(q)
		for i := 0; i < q_size; i++ {
			curr := q[0]
			q = q[1:]
			switch curr[1] {
			case 0: //color red
				for j := 0; j < len(b[curr[0]]); j++ {
					if l[1][b[curr[0]][j]] == 2*n+1 {
						l[1][b[curr[0]][j]] = 1 + l[0][curr[0]]
						q = append(q, [2]int{b[curr[0]][j], 1})
					}
				}
			case 1: // blue
				for j := 0; j < len(r[curr[0]]); j++ {
					if l[0][r[curr[0]][j]] == 2*n+1 {
						l[0][r[curr[0]][j]] = 1 + l[1][curr[0]]
						q = append(q, [2]int{r[curr[0]][j], 0})
					}
				}
			}
		}

	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		if l[0][i] < l[1][i] {
			res[i] = l[0][i]
		} else {
			res[i] = l[1][i]
		}

		if res[i] == 2*n+1 {
			res[i] = -1
		}
	}
	return res
}

func main() {
	fmt.Println(shortest_path_alternating(3,
		[][]int{
			{0, 1}, {0, 2},
		},
		[][]int{
			{1, 0},
		}))
}
