package main

import "fmt"

//---------------union find----------------------
func graph_valid_tree_union_find(n int, edges [][]int) bool {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	for i := 0; i < len(edges); i++ {
		if !union(edges[i][0], edges[i][1], &parent, &size) {
			return false
		}
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			cnt++
		}
	}
	return !(cnt > 1)
}

func union(x, y int, parent *[]int, size *[]int) bool {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return false
	}
	if (*size)[x] >= (*size)[y] {
		(*size)[x] += (*size)[y]
		(*parent)[y] = x
	} else {
		(*size)[y] += (*size)[x]
		(*parent)[x] = y
	}
	return true
}

func find(x int, parent *[]int) int {
	if x != (*parent)[x] {
		(*parent)[x] = find((*parent)[x], parent)
	}
	return (*parent)[x]
}

//---------------union find----------------------

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	n := len(edges)
	for i := 0; i < n; i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

func graph_valid_tree(n int, edges [][]int) bool {

	graph := make_graph(edges)

	visited := make(map[int]int)

	q := []int{0}
	visited[0] = -1

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for i := 0; i < len(graph[curr]); i++ {
			if visited[curr] == graph[curr][i] {
				continue
			}
			if _, ok := visited[graph[curr][i]]; ok {
				return false
			}
			q = append(q, graph[curr][i])
			visited[graph[curr][i]] = curr
		}
	}
	return len(visited) == n
}

func main() {
	fmt.Println(graph_valid_tree(5, [][]int{
		{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4},
	}))
}
