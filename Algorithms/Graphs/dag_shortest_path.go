package main

import (
	"fmt"
	"math"
)

// source : CLRS

//dag shortest path. the algorithm starts by topological
//sorting the dag to impose a linear ordering. if a dag
//contains a path from u to v then in topological ordering
//u precedes v
func dag_shortest_path(g map[int][]Pair, n int, source int) []int {
	dis := make([]int, n)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt64
	}
	dis[source] = 0
	indegree := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(g[i]); j++ {
			indegree[g[i][j].v]++
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
			indegree[g[u][i].v]--
			if indegree[g[u][i].v] == 0 {
				q = append(q, g[u][i].v)
			}
		}
	}
	for i := 0; i < len(ordering); i++ {
		u := ordering[i]
		for j := 0; j < len(g[u]); j++ {
			v, wt := g[u][j].v, g[u][j].wt
			if dis[u] == math.MaxInt64 && wt > 0 {
				continue
			}
			if dis[v] > wt+dis[u] {
				parent[v] = u
				dis[v] = wt + dis[u]
			}
		}
	}
	fmt.Println("Parent for each node : ", parent)
	return dis
}

type Pair struct{ v, wt int }

func main() {
	g := map[int][]Pair{
		1: {{0, 1}, {2, 1}},
		2: {{3, 1}},
	}
	fmt.Println(dag_shortest_path(g, 4, 2))
}
