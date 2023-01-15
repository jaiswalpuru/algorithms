package main

import (
	"container/heap"
	"fmt"
	"math"
)

//-----------------------------------Using dijkstra----------------------------------------
type P []Pair

func (m P) Len() int              { return len(m) }
func (m P) Less(i, j int) bool    { return m[i].w < m[j].w }
func (m P) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *P) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *P) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func network_delay_time_dijkstra(times [][]int, n, k int) int {
	g := make_graph(times)
	visited := make([]bool, n+1)
	mh := &P{}
	dis := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dis[i] = math.MaxInt64
	}
	dis[k] = 0
	heap.Push(mh, Pair{k, 0})
	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Pair)
		if visited[curr.v] {
			continue
		}
		visited[curr.v] = true
		for i := 0; i < len(g[curr.v]); i++ {
			node := g[curr.v][i]
			if !visited[node.v] {
				if dis[node.v] > curr.w+node.w {
					dis[node.v] = curr.w + node.w
					heap.Push(mh, Pair{node.v, dis[node.v]})
				}
			}
		}
	}
	min_cost := -1
	for i := 1; i < n+1; i++ {
		if dis[i] == math.MaxInt64 {
			return -1
		}
		min_cost = max(min_cost, dis[i])
	}
	return min_cost
}

//-----------------------------------Using dijkstra----------------------------------------

//-----------------------------------Using bfs----------------------------------------
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

type p struct {
	v int
	t int
}

func network_delay_time(edges [][]int, n, k int) int {
	graph := make_graph(edges)
	q := []p{{k, 0}}

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
				q = append(q, p{graph[curr.v][i].v, distance[graph[curr.v][i].v]})
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

//-----------------------------------Using bfs----------------------------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(network_delay_time_dijkstra([][]int{
		{2, 1, 1}, {2, 3, 1}, {3, 4, 1},
	}, 4, 2))
}
