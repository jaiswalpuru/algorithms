package main

import "fmt"

/*
	Articulation points of a graph G
	is vertex whose removal disconnects G.
	This problem can be solved using tarjan's
	algorithm.
	There is another approach to solve this problem
	for reference look at critical connection in a
	graph.
*/

func articulation_points(g map[int][]int, n int) []int {
	articulation_point := []int{}
	low := make([]int, n)
	disc := make([]int, n)
	for i := 0; i < n; i++ {
		disc[i] = -1
	}
	time := 0
	dfs(g, 0, -1, &disc, &low, &time, &articulation_point)
	return articulation_point
}

func dfs(g map[int][]int, u, parent int, disc, low *[]int, time *int, articulation_points *[]int) {
	*time++
	(*disc)[u] = *time
	(*low)[u] = *time
	for i := 0; i < len(g[u]); i++ {
		v := g[u][i]
		if v == parent {
			continue
		}
		//if not discovered yet
		if (*disc)[v] == -1 {
			dfs(g, v, u, disc, low, time, articulation_points)
			(*low)[u] = min((*low)[u], (*low)[v])
			//we can say that u and v it is a articulation point if discovery of u is
			//less than or equal to low link value for its neighbors as this is not
			//present in the cycle which the other component has formed.
			if (*disc)[u] < (*low)[v] {
				*articulation_points = append(*articulation_points, u)
				*articulation_points = append(*articulation_points, v)
			}
		}
		(*low)[u] = min((*low)[u], (*disc)[v])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	g1 := map[int][]int{
		0: {1, 2},
		1: {0, 2, 3},
		2: {1, 0},
		3: {1},
	}
	fmt.Println(articulation_points(g1, 4))
}
