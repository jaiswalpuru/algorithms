package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	n := len(edges)

	//let say graph is like [[1,0],[2,0]]
	// negative means there is a way from 1 to 0
	//positivie means there is wasy from 0 to 1

	for i := 0; i < n; i++ {
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
		graph[edges[i][0]] = append(graph[edges[i][0]], -edges[i][1])
	}
	return graph
}

func reorder_routes(n int, edges [][]int) int {
	graph := make_graph(edges)
	min_edges := 0
	visited := make(map[int]bool)
	q := []int{0}
	visited[0] = true
	for len(q) > 0 {
		m := len(q)
		for i := 0; i < m; i++ {
			curr := q[0]
			q = q[1:]
			for _, v := range graph[curr] {
				if !visited[abs(v)] {
					if v < 0 {
						min_edges++
					}
					q = append(q, abs(v))
					visited[abs(v)] = true
				}
			}
		}
	}

	return min_edges
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	fmt.Println(reorder_routes(3, [][]int{
		{1, 2}, {2, 0},
	}))
}
