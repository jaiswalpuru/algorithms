package main

import (
	"fmt"
	"sort"
)

//------------------union find-----------------
func smallest_string_with_swaps_union_find(s string, pairs [][]int) string {
	parent := make([]int, len(s))
	size := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		parent[i] = i
		size[i] = 1
	}
	for i := 0; i < len(pairs); i++ {
		curr := pairs[i]
		union(curr[0], curr[1], &parent, &size)
	}
	group := make(map[int][]int)
	for i := 0; i < len(parent); i++ {
		g := find(i, &parent)
		group[g] = append(group[g], i)
	}
	res := make([]byte, len(s))
	for _, v := range group {
		st := []byte{}
		for i := 0; i < len(v); i++ {
			st = append(st, s[v[i]])
		}
		sort.Slice(st, func(i, j int) bool { return st[i] < st[j] })
		for i := 0; i < len(v); i++ {
			res[v[i]] = st[i]
		}
	}
	return string(res)
}

func find(x int, parent *[]int) int {
	if x != (*parent)[x] {
		(*parent)[x] = find((*parent)[x], parent)
	}
	return (*parent)[x]
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

//------------------union find-----------------

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
