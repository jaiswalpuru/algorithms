package main

import "fmt"

func make_graph(edges [][]int) map[int]map[int]int {
	graph := make(map[int]map[int]int)
	n := len(edges)

	for i := 0; i < n; i++ {
		if _, ok := graph[edges[i][0]]; !ok {
			graph[edges[i][0]] = make(map[int]int)
		}
		if _, ok := graph[edges[i][1]]; !ok {
			graph[edges[i][1]] = make(map[int]int)
		}
		graph[edges[i][0]][edges[i][1]] = edges[i][0]
		graph[edges[i][1]][edges[i][0]] = edges[i][1]
	}
	return graph
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximal_network_rank(n int, edges [][]int) int {
	graph := make_graph(edges)

	max_rank := 0
	for i := 0; i < n; i++ {
		rank := len(graph[i])
		sum := 0
		node := -1
		for j := 0; j < n; j++ {
			if i != j {
				if sum < len(graph[j]) {
					sum = len(graph[j])
					node = j
					if _, ok := graph[i][node]; ok {
						sum -= 1
					}
				}
			}
		}
		max_rank = max(max_rank, rank+sum)
	}
	return max_rank
}

func maximal_network_rank_using_degrees(n int, edges [][]int) int {
	graph := make([][]bool, n)
	degrees := make([]int, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}

	for i := 0; i < len(edges); i++ {
		graph[edges[i][0]][edges[i][1]] = true
		graph[edges[i][1]][edges[i][0]] = true
		degrees[edges[i][0]]++
		degrees[edges[i][1]]++
	}

	max_connection := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] {
				max_connection = max(max_connection, degrees[i]+degrees[j]-1)
			} else {
				max_connection = max(max_connection, degrees[i]+degrees[j])
			}
		}
	}
	return max_connection
}

func main() {
	fmt.Println(maximal_network_rank(5, [][]int{
		{2, 3}, {0, 3}, {0, 4}, {4, 1},
	}))
}
