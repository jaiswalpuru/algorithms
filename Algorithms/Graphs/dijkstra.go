package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
	Dijkstra algorithm:
	Single source shortest path algorithm.
	O(ElogV)
*/

type Pair struct{ v, wt int }
type P []Pair

func (m P) Len() int              { return len(m) }
func (m P) Less(i, j int) bool    { return m[i].wt < m[j].wt }
func (m P) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *P) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *P) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

//dis[0] = will be MaxInt as it is not considered in this graph
func dijkstra(g map[int][]Pair, src, n int) []int {
	dis := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dis[i] = math.MaxInt64
	}
	dis[src] = 0
	dis[0] = -1
	parent := make([]int, n+1)
	parent[src] = -1
	parent[0] = -1
	visited := make([]bool, n+1)
	mh := &P{}
	heap.Push(mh, Pair{src, 0})
	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Pair)
		if visited[curr.v] {
			continue
		}
		visited[curr.v] = true
		for i := 0; i < len(g[curr.v]); i++ {
			node := g[curr.v][i]
			if !visited[node.v] && dis[node.v] > node.wt+curr.wt {
				parent[node.v] = curr.v
				dis[node.v] = node.wt + curr.wt
				heap.Push(mh, Pair{node.v, dis[node.v]})
			}
		}
	}
	fmt.Println("Parents for each node : ", parent)
	return dis
}

func main() {
	g1 := map[int][]Pair{
		2: {{1, 1}, {3, 1}},
		3: {{4, 1}},
	}
	fmt.Println("Shortest Path to all other vertices from node 2 : ", dijkstra(g1, 2, 4))
}
