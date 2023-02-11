package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
	Prims algorithm to find MST
	O((V+E)logV) - using binary heaps
	works for both directed and undirected.
*/

type Pair struct{ v, wt int }
type M []Pair

func (m M) Len() int              { return len(m) }
func (m M) Less(i, j int) bool    { return m[i].wt < m[j].wt }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func make_graph(edges [][]int) map[int][]Pair {
	g := make(map[int][]Pair)
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], Pair{edges[i][1], edges[i][2]})
		g[edges[i][1]] = append(g[edges[i][1]], Pair{edges[i][0], edges[i][2]})
	}
	return g
}

func prims(edges [][]int, n int) int {
	g := make_graph(edges)
	visited := make([]bool, n)
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt64
	}
	parent := make([]int, n)
	dis[0] = 0
	parent[0] = -1
	mh := &M{}
	heap.Push(mh, Pair{0, 0})
	edge_cnt := 0
	min_cost := 0
	for mh.Len() > 0 && edge_cnt < n-1 {
		curr := heap.Pop(mh).(Pair)
		if visited[curr.v] {
			continue
		}
		visited[curr.v] = true
		min_cost += curr.wt
		for i := 0; i < len(g[curr.v]); i++ {
			node := g[curr.v][i]
			if !visited[node.v] {
				if dis[node.v] > node.wt {
					parent[node.v] = curr.v
					dis[node.v] = node.wt
					heap.Push(mh, Pair{node.v, node.wt})
				}
			}
		}
		edge_cnt++
	}
	fmt.Println("Parents : ", parent)
	fmt.Println("Distance : ", dis)
	return min_cost
}

func main() {
	g1 := [][]int{
		{0, 1, 5}, {1, 2, 4}, {2, 9, 2},
		{0, 4, 1}, {0, 3, 4}, {1, 3, 2},
		{2, 7, 4}, {2, 8, 1}, {9, 8, 0},
		{4, 5, 1}, {5, 6, 7}, {6, 8, 4},
		{4, 3, 2}, {5, 3, 5}, {3, 6, 11},
		{6, 7, 1}, {3, 7, 2}, {7, 8, 6},
	}
	fmt.Println(prims(g1, 10))
}
