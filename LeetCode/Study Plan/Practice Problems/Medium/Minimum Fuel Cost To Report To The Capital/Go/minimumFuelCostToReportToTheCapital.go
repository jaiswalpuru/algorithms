package main

import (
	"fmt"
	"math"
)

func minimumFuelCost(roads [][]int, seats int) int64 {
	g := makeGraph(roads)
	size := len(roads) + 1
	degree := make([]int, size)

	for i := 0; i < size-1; i++ {
		degree[roads[i][0]]++
		degree[roads[i][1]]++
	}

	q := []int{}
	subtree := make([]int, size)

	for i := 1; i < size; i++ {
		subtree[i] = 1
		if degree[i] == 1 {
			q = append(q, i)
		}
	}

	fuel := int64(0)
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		fuel += int64(math.Ceil(float64(subtree[curr]) / float64(seats)))
		for i := 0; i < len(g[curr]); i++ {
			v := g[curr][i]
			degree[v]--
			subtree[v] += subtree[curr]
			if degree[v] == 1 && v != 0 {
				q = append(q, v)
			}
		}
	}

	return fuel
}

func makeGraph(edges [][]int) map[int][]int {
	g := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], edges[i][1])
		g[edges[i][1]] = append(g[edges[i][1]], edges[i][0])
	}
	return g
}

func main() {
	fmt.Println(minimumFuelCost([][]int{
		{0, 1}, {0, 2}, {0, 3},
	}, 5))
}
