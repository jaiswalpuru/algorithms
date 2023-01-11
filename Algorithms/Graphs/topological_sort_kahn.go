package main

import "fmt"

/*
	Topological sort using kahn algorithm
	O(V+E)
*/

func topological_sort(g map[int][]int, n int) []int {
	indegree := make([]int, n)
	for _, v := range g {
		for i := 0; i < len(v); i++ {
			indegree[v[i]]++
		}
	}
	q := []int{}
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	ordering := []int{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		ordering = append(ordering, u)
		for i := 0; i < len(g[u]); i++ {
			indegree[g[u][i]]--
			if indegree[g[u][i]] == 0 {
				q = append(q, g[u][i])
			}
		}
	}
	if len(ordering) != n {
		err := fmt.Errorf("%q", "Graph not a dag")
		fmt.Println(err)
		return nil
	}
	return ordering
}

func main() {
	g1 := map[int][]int{
		5: {2, 0},
		4: {0, 1},
		2: {3},
		3: {1},
	}
	g2 := map[int][]int{
		0: {1},
		1: {2},
		2: {1},
	}
	fmt.Println(topological_sort(g1, 6))
	fmt.Println(topological_sort(g2, 3))
}
