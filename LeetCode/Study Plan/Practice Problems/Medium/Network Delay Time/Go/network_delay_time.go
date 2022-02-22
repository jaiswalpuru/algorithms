package main

import (
	"fmt"
	"math"
)

type Pair struct {
	v int
	w int
}

func make_graph(edges [][]int) map[int][]Pair {
	graph := make(map[int][]Pair)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], Pair{edges[i][1], edges[i][2]})
	}
	return graph
}

type P struct {
	v int
	t int
}

func network_delay_time(edges [][]int, n, k int) int {
	graph := make_graph(edges)
	q := []P{{k, 0}}

	time := 0
	distance := make([]int, n+1)
	for i := 1; i <= n; i++ {
		distance[i] = math.MaxInt64
	}
	distance[k] = 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for i := 0; i < len(graph[curr.v]); i++ {
			if distance[graph[curr.v][i].v] > curr.t+graph[curr.v][i].w {
				distance[graph[curr.v][i].v] = curr.t + graph[curr.v][i].w
				q = append(q, P{graph[curr.v][i].v, distance[graph[curr.v][i].v]})
			}
		}
	}
	for i := 1; i <= n; i++ {
		time = max(time, distance[i])
	}
	if time == math.MaxInt64 {
		return -1
	}
	return time
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(network_delay_time([][]int{
		{2, 1, 1}, {2, 3, 1}, {3, 4, 1},
	}, 4, 2))
}
