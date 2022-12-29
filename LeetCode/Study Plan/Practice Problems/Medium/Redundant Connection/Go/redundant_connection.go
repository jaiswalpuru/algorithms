package main

import "fmt"

//---------------------using dfs--------------------------------
func redundant_connection(edges [][]int) []int {
	graph := make(map[int][]int)
	n := len(edges)
	visited := make(map[int]bool)

	for i := 0; i < n; i++ {
		visited = make(map[int]bool)
		if dfs(graph, edges[i][0], edges[i][1], &visited) {
			return []int{edges[i][0], edges[i][1]}
		}
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return []int{}
}

func dfs(graph map[int][]int, src, target int, visited *map[int]bool) bool {
	if _, ok := (*visited)[src]; !ok {
		(*visited)[src] = true
		if src == target {
			return true
		}
		for i := 0; i < len(graph[src]); i++ {
			if dfs(graph, graph[src][i], target, visited) {
				return true
			}
		}
	}
	return false
}

//---------------------using dfs--------------------------------

//---------------------union find--------------------------------
func redundant_connection_union_find(edges [][]int) []int {
	n := len(edges)
	parent := make([]int, n+1)
	size := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		parent[i] = i
		size[i] = 1
	}
	for i := 0; i < len(edges); i++ {
		x, y := edges[i][0], edges[i][1]
		if find(x, &parent) == find(y, &parent) {
			return []int{x, y}
		}
		union(x, y, &parent, &size)
	}
	return nil
}

func find(x int, parent *[]int) int {
	if x == (*parent)[x] {
		return x
	}
	(*parent)[x] = find((*parent)[x], parent)
	return (*parent)[x]
}

func union(x, y int, parent *[]int, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] > (*size)[y] {
		(*parent)[y] = x
		(*size)[x] += (*size)[y]
	} else {
		(*parent)[x] = y
		(*size)[y] += (*size)[x]
	}
}

//---------------------union find--------------------------------

func main() {
	fmt.Println(redundant_connection_union_find([][]int{
		{1, 2}, {1, 3}, {2, 3},
	}))
}
