package main

import (
	"fmt"
)

//1 -> GRAY (backward edge, which tells this is a loop)
//2 -> BLACK
var (
	WHITE = -1
	GRAY  = 1
	BLACK = 2
)

func all_path(n int, edges [][]int, src int, dest int) bool {
	g := make_graph(edges)
	color := make([]int, n)
	for i := 0; i < len(color); i++ {
		color[i] = WHITE
	}
	return dfs(src, dest, &color, g)
}

func dfs(src, dst int, color *[]int, g map[int][]int) bool {
	if (*color)[src] != WHITE {
		return (*color)[src] == BLACK
	}
	if len(g[src]) == 0 {
		return src == dst
	}
	(*color)[src] = GRAY
	for i := 0; i < len(g[src]); i++ {
		if !dfs(g[src][i], dst, color, g) {
			return false
		}
	}
	(*color)[src] = BLACK
	return true
}

func make_graph(edges [][]int) map[int][]int {
	vertices := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		vertices[edges[i][0]] = append(vertices[edges[i][0]], edges[i][1])
	}
	return vertices
}

func main() {
	fmt.Println(all_path(4,
		[][]int{
			{0, 1},
			{0, 2},
			{1, 3},
			{2, 3},
		}, 0, 3))
}
