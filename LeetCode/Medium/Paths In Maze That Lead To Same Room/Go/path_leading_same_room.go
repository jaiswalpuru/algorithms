package main

import "fmt"

func make_graph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func path_lead_to_same_room(n int, corridors [][]int) int {
	graph := make_graph(corridors)
	visited := make(map[int]bool)
	total := 0
	q := []int{}
	ongoing := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 0; j < len(graph[i]); j++ {
			if !visited[graph[i][j]] {
				q = append(q, graph[i][j])
				ongoing[graph[i][j]] = 1
			}
		}

		for len(q) > 0 {
			curr_node := q[0]
			q = q[1:]
			for j := 0; j < len(graph[curr_node]); j++ {
				if ongoing[graph[curr_node][j]] == 1 {
					total++
				}
			}
			ongoing[curr_node] = 0
		}
		visited[i] = true
	}
	return total
}

func main() {
	fmt.Println(path_lead_to_same_room(5, [][]int{
		{4, 1}, {4, 2}, {3, 4}, {5, 3}, {5, 2}, {1, 3}, {3, 2}, {2, 1}, {5, 4}, {5, 1},
	}))
}
