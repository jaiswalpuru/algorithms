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

func dfs_search(vertex int, graph [][]int, visited *[]bool, parent int) bool {
	(*visited)[vertex] = true

	for i := 0; i < len(graph[vertex]); i++ {
		if !(*visited)[graph[vertex][i]] {
			if dfs_search(graph[vertex][i], graph, visited, vertex) {
				return true
			}
		} else if parent != graph[vertex][i] {
			return true
		}
	}
	return false
}

func has_cycle(graph [][]int) bool {
	visited := make([]bool, len(graph))
	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			if dfs_search(i, graph, &visited, -1) {
				return true
			}
		}
	}
	return false
}

func main() {
	graph := make([][]int, 5)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 0)
	}

	add_edge(0, 1, &graph, UNDIRECTED)
	add_edge(0, 3, &graph, UNDIRECTED)
	add_edge(1, 2, &graph, UNDIRECTED)
	add_edge(2, 3, &graph, UNDIRECTED)
	add_edge(2, 4, &graph, UNDIRECTED)
	add_edge(3, 4, &graph, UNDIRECTED)

	fmt.Println(has_cycle(graph))
}
