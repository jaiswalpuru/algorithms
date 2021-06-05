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

func traverse(graph [][]int, curr int, result *map[int]int) int {
	descendants := 0

	for child := range graph[curr] {
		num_nodes := traverse(graph, graph[curr][child], result)
		(*result)[graph[curr][child]] += num_nodes - 1
		descendants += num_nodes
	}
	return descendants + 1
}

func max_edges(graph [][]int) int {

	start := 1

	vertices := make(map[int]int)
	traverse(graph, start, &vertices)
	fmt.Println(vertices)
	ans := []int{}

	for _, v := range vertices {
		if v%2 == 1 {
			ans = append(ans, v)
		}
	}
	return len(ans)
}

func main() {

	graph := make([][]int, 9)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, 0)
	}

	add_edge(1, 2, &graph, DIRECTED)
	add_edge(1, 3, &graph, DIRECTED)
	add_edge(3, 4, &graph, DIRECTED)
	add_edge(3, 5, &graph, DIRECTED)
	add_edge(4, 6, &graph, DIRECTED)
	add_edge(4, 7, &graph, DIRECTED)
	add_edge(4, 8, &graph, DIRECTED)

	fmt.Println(max_edges(graph))
}
