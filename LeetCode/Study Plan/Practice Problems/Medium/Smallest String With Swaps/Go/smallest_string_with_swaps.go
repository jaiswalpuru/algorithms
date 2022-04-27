package main

import "fmt"
import "sort"

func make_graph(edges [][]int) map[int][]int {
	n := len(edges)
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func dfs(s string, v int, chars *[]byte, indices *[]int, visited *map[int]bool, graph map[int][]int) {
	*chars = append(*chars, s[v])
	*indices = append(*indices, v)
	(*visited)[v] = true

	for i := 0; i < len(graph[v]); i++ {
		if !(*visited)[graph[v][i]] {
			dfs(s, graph[v][i], chars, indices, visited, graph)
		}
	}
}

func smallest_string_with_swaps(s string, pairs [][]int) string {
	graph := make_graph(pairs)

	n := len(s)
	res := make([]byte, n)
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		if !visited[i] {
			chars := []byte{}
			indices := []int{}
			dfs(s, i, &chars, &indices, &visited, graph)
			sort.Slice(chars, func(i, j int) bool {
				return chars[i] < chars[j]
			})
			sort.Ints(indices)

			for index := 0; index < len(chars); index++ {
				res[indices[index]] = chars[index]
			}
		}
	}

	return string(res)
}

func main() {
	fmt.Println(smallest_string_with_swaps("dcab", [][]int{
		{0, 3}, {1, 2},
	}))
}
