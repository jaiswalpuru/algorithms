package main

import (
	"fmt"
	"math"
)

type Pair struct {
	v, wt int
}

func minScore(n int, roads [][]int) int {
	g := make(map[int][]Pair)
	for i := 0; i < len(roads); i++ {
		g[roads[i][0]] = append(g[roads[i][0]], Pair{roads[i][1], roads[i][2]})
		g[roads[i][1]] = append(g[roads[i][1]], Pair{roads[i][0], roads[i][2]})
	}

	minDis := math.MaxInt64
	visited := make(map[int]bool)
	q := []int{1}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for i := 0; i < len(g[u]); i++ {
			to := g[u][i]
			minDis = min(minDis, to.wt)
			if visited[to.v] {
				continue
			}
			visited[to.v] = true
			q = append(q, to.v)
		}
	}

	return minDis
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minScore(4, [][]int{
		{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7},
	}))
}
