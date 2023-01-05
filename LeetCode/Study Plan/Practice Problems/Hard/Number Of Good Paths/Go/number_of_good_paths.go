package main

import (
	"fmt"
	"sort"
)

//----------------------union find-------------------------
func number_of_good_paths_union(vals []int, edges [][]int) int {
	graph := make_graph(edges)
	parent := make([]int, len(vals))
	size := make([]int, len(vals))
	for i := 0; i < len(vals); i++ {
		parent[i] = i
		size[i] = 1
	}
	val := make(map[int][]int)
	for i := 0; i < len(vals); i++ {
		val[vals[i]] = append(val[vals[i]], i)
	}
	sorted_val := []int{}
	for k, _ := range val {
		sorted_val = append(sorted_val, k)
	}
	res := 0
	sort.Ints(sorted_val)
	for i := 0; i < len(sorted_val); i++ {
		nodes := val[sorted_val[i]]
		for j := 0; j < len(nodes); j++ {
			curr := nodes[j]
			for k := 0; k < len(graph[curr]); k++ {
				if vals[curr] >= vals[graph[curr][k]] {
					union(curr, graph[curr][k], &parent, &size)
				}
			}
		}
		groups := make(map[int]int)
		for j := 0; j < len(nodes); j++ {
			groups[find(nodes[j], &parent)]++
		}
		for _, v := range groups {
			res += v * (v + 1) / 2
		}
	}
	return res
}

func union(x, y int, parent *[]int, size *[]int) {
	x = find(x, parent)
	y = find(y, parent)
	if x == y {
		return
	}
	if (*size)[x] >= (*size)[y] {
		(*size)[x] += (*size)[y]
		(*parent)[y] = x
	} else {
		(*size)[y] += (*size)[x]
		(*parent)[x] = y
	}
}

func find(x int, parent *[]int) int {
	if x != (*parent)[x] {
		(*parent)[x] = find((*parent)[x], parent)
	}
	return (*parent)[x]
}

//----------------------union find-------------------------

//-----------------------DFS (TLE)-----------------------------
func number_of_good_paths(vals []int, edges [][]int) int {
	graph := make_graph(edges)
	res := len(vals)
	for i := 0; i < len(vals); i++ {
		visited := make(map[int]bool)
		dfs(i, graph, vals, &res, &visited, vals[i])
	}
	return res / 2
}

func dfs(curr_node int, g map[int][]int, val []int, path *int, visited *map[int]bool, parent_val int) {
	(*visited)[curr_node] = true
	if parent_val == val[curr_node] {
		*path++
	}
	for i := 0; i < len(g[curr_node]); i++ {
		if !(*visited)[g[curr_node][i]] && val[g[curr_node][i]] <= parent_val {
			dfs(g[curr_node][i], g, val, path, visited, parent_val)
		}
	}
}

func make_graph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

//-----------------------DFS (TLE)-----------------------------

func main() {
	fmt.Println(number_of_good_paths_union([]int{1, 3, 2, 1, 3}, [][]int{
		{0, 1}, {0, 2}, {2, 3}, {2, 4},
	}))
}
