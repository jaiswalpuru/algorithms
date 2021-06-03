package main

import "fmt"

type GraphType string

const DIRECTED = "DIRECTED"
const UNDIRECTED = "UNDIRECTED"

func add_edge(u, v int, graph *[][]int, gtype GraphType) {
	(*graph)[u] = append((*graph)[u], v)
	if gtype == UNDIRECTED {
		(*graph)[v] = append((*graph)[v], u)
	}
}

func dfs(curr int, graph [][]int, visited *[]bool, ordering *[]int) {
	if !(*visited)[curr] {
		(*visited)[curr] = true
		*ordering = append(*ordering, curr)
		for j := 0; j < len(graph[curr]); j++ {
			if !(*visited)[graph[curr][j]] {
				dfs(graph[curr][j], graph, visited, ordering)
			}
		}
	}
}

func bfs(curr int, graph [][]int, visited *[]bool, ordering *[]int) {
	queue := []int{}
	queue = append(queue, curr)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if (*visited)[curr] {
			continue
		}
		(*ordering) = append((*ordering), curr)
		(*visited)[curr] = true
		for i := 0; i < len(graph[curr]); i++ {
			if (*visited)[graph[curr][i]] {
				continue
			}
			queue = append(queue, graph[curr][i])
		}
	}
}

func main() {
	graph := make([][]int, 5)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 0)
	}

	add_edge(0, 1, &graph, UNDIRECTED)
	add_edge(0, 2, &graph, UNDIRECTED)
	add_edge(0, 3, &graph, UNDIRECTED)
	add_edge(1, 2, &graph, UNDIRECTED)
	add_edge(2, 3, &graph, UNDIRECTED)
	add_edge(2, 4, &graph, UNDIRECTED)
	add_edge(3, 4, &graph, UNDIRECTED)

	visited := make([]bool, 5)
	ordering := make([]int, 0)
	dfs(0, graph, &visited, &ordering)
	fmt.Println("DFS : ", ordering)

	visited = make([]bool, 5)
	ordering = make([]int, 0)
	bfs(0, graph, &visited, &ordering)
	fmt.Println("BFS : ", ordering)
}
