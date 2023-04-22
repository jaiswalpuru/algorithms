package main

import (
	"container/heap"
	"fmt"
)

type M []Node

func (m M) Len() int              { return len(m) }
func (m M) Less(i, j int) bool    { return m[i].wt < m[j].wt }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(Node)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func optimize_water_distribution_in_village(n int, wells []int, pipes [][]int) int {
	//add virtual node 0 which connects with all nodes in wells
	graph := make_graph(pipes, wells)
	mh := &M{}
	for i := 0; i < len(graph[0]); i++ {
		heap.Push(mh, graph[0][i])
	}
	set := make(map[int]bool)
	total_cost := 0
	set[0] = true
	for len(set) < n+1 {
		edge := heap.Pop(mh).(Node)
		next_house := edge.u
		cost := edge.wt
		if set[next_house] {
			continue
		}
		set[next_house] = true
		total_cost += cost
		for i := 0; i < len(graph[next_house]); i++ {
			if !set[graph[next_house][i].u] {
				heap.Push(mh, graph[next_house][i])
			}
		}
	}
	return total_cost
}

type Node struct {
	u  int
	wt int
}

func make_graph(edges [][]int, wells []int) map[int][]Node {
	g := make(map[int][]Node)
	for i := 0; i < len(wells); i++ {
		g[0] = append(g[0], Node{i + 1, wells[i]})
	}
	for i := 0; i < len(edges); i++ {
		g[edges[i][0]] = append(g[edges[i][0]], Node{edges[i][1], edges[i][2]})
		g[edges[i][1]] = append(g[edges[i][1]], Node{edges[i][0], edges[i][2]})
	}
	return g
}

func main() {
	fmt.Println(optimize_water_distribution_in_village(3, []int{1, 2, 2}, [][]int{
		{1, 2, 1}, {2, 3, 1},
	}))
}
