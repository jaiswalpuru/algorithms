package main

import (
	"fmt"
	"math"
)

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	g := makeGraph(edges)
	n := len(edges)
	dis1 := bfs(node1, n, g)
	dis2 := bfs(node2, n, g)
	minDis := math.MaxInt64
	ind := -1
	for i := 0; i < n; i++ {
		if dis1[i] == int(1e7) || dis2[i] == int(1e7) {
			continue
		}
		if minDis > max(dis1[i], dis2[i]) {
			ind = i
			minDis = max(dis1[i], dis2[i])
		}
	}
	return ind
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bfs(node, n int, g map[int][]int) []int {
	q := []Pair{{node, 0}}
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = int(1e7)
	}
	dis[node] = 0
	visited := make(map[int]bool)
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		dis[curr.v] = curr.dis
		if visited[curr.v] {
			continue
		}
		visited[curr.v] = true
		for i := 0; i < len(g[curr.v]); i++ {
			if !visited[g[curr.v][i]] {
				q = append(q, Pair{g[curr.v][i], curr.dis + 1})
			}
		}
	}
	return dis
}

func makeGraph(edges []int) map[int][]int {
	n := len(edges)
	g := make(map[int][]int)
	for i := 0; i < n; i++ {
		if edges[i] == -1 {
			continue
		}
		g[i] = append(g[i], edges[i])
	}
	return g
}

type Pair struct {
	v, dis int
}

func main() {
	fmt.Println(closestMeetingNode([]int{1, 2, -1}, 0, 2))
}
