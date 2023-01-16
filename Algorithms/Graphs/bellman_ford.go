package main

import (
	"fmt"
	"math"
)

/*
	Bellman-Ford algorithm : computes the shortest
	path from source to all other nodes.
	In a graph with N vertices there are at most N-1
	edges between any two vertices.
	No negative wt cycles
*/

func bellman_ford(edges [][]int, v, start int) []float64 {
	dis := make([]float64, v)
	for i := 0; i < v; i++ {
		dis[i] = math.Inf(1)
	}
	dis[start] = 0
	//relax v-1 time all the edges
	for i := 0; i < v-1; i++ {
		for j := 0; j < len(edges); j++ {
			curr := edges[j]
			// dis[from] + cost < dis[to]
			if dis[curr[0]]+float64(curr[2]) < dis[curr[1]] {
				dis[curr[1]] = dis[curr[0]] + float64(curr[2])
			}
		}
	}
	//loop again to check for negative cycles
	for i := 0; i < v-1; i++ {
		for j := 0; j < len(edges); j++ {
			curr := edges[j]
			if dis[curr[0]]+float64(curr[2]) < dis[curr[1]] {
				dis[curr[1]] = math.Inf(-1)
			}
		}
	}
	return dis
}

func main() {
	//edges from, to, cost
	edges := [][]int{
		{0, 1, 1}, {1, 2, 1}, {2, 4, 1},
		{4, 3, -3}, {3, 2, 1}, {1, 5, 4},
		{1, 6, 4}, {5, 6, 5}, {6, 7, 4}, {5, 7, 3},
	}
	//apply relaxation for edges v-1 times i.e 9
	//v = 9
	fmt.Println(bellman_ford(edges, 9, 0))
}
