package main

import "fmt"

//----------------------dfs-----------------------------
func make_graph(edges [][]int) map[int][]int {
	n := len(edges)
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
		graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	return graph
}

func dfs(start int, visited *[]bool, graph map[int][]int) {
	(*visited)[start] = true
	for i := 0; i < len(graph[start]); i++ {
		if !(*visited)[graph[start][i]] {
			dfs(graph[start][i], visited, graph)
		}
	}
}

func num_connected_components(n int, edges [][]int) int {
	components := 0

	graph := make_graph(edges)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			components++
			dfs(i, &visited, graph)
		}
	}

	return components
}

//----------------------dfs-----------------------------

//----------------------using union find-----------------------------
func num_connected_components_union_find(n int, edges [][]int) int {
	size := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = i
		parent[i] = i
	}
	for i := 0; i < len(edges); i++ {
		union(edges[i][0], edges[i][1], &parent, &size)
	}
	num := 0
	for i := 0; i < n; i++ {
		if parent[i] == i {
			num++
		}
	}
	return num
}

func union(u, v int, parent *[]int, size *[]int) {
	u = find(u, parent)
	v = find(v, parent)
	if u == v {
		return
	}
	if (*size)[u] > (*size)[v] {
		(*parent)[v] = u
		(*size)[u] += (*size)[v]
	} else {
		(*parent)[u] = v
		(*size)[v] += (*size)[u]
	}
}

func find(u int, parent *[]int) int {
	if u == (*parent)[u] {
		return u
	}
	(*parent)[u] = find((*parent)[u], parent)
	return (*parent)[u]
}

//----------------------using union find-----------------------------

func main() {
	fmt.Println(num_connected_components_union_find(2, [][]int{
		{1, 0},
	}))
}
