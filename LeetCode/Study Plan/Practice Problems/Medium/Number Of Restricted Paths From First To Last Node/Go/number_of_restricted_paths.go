package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

type Pair struct{ next, weight int }

type P []Pair

func (p P) Len() int              { return len(p) }
func (p P) Less(i, j int) bool    { return p[i].weight < p[j].weight }
func (p P) Swap(i, j int)         { p[i], p[j] = p[j], p[i] }
func (p *P) Push(val interface{}) { *p = append(*p, val.(Pair)) }
func (p *P) Pop() interface{} {
	old := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return old
}

func make_graph(edges [][]int) map[int][]Pair {
	graph := make(map[int][]Pair)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], Pair{next: edges[i][1], weight: edges[i][2]})
		graph[edges[i][1]] = append(graph[edges[i][1]], Pair{next: edges[i][0], weight: edges[i][2]})
	}

	return graph
}

func restricted_paths(n int, edges [][]int) int {
	graph := make_graph(edges)
	mh := &P{}
	dis_from_end := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dis_from_end[i] = math.MaxInt64
	}
	dis_from_end[n] = 0
	heap.Push(mh, Pair{next: n, weight: 0})
	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Pair)
		node := curr.next
		weight := curr.weight

		for i := 0; i < len(graph[node]); i++ {
			if dis_from_end[graph[node][i].next] > graph[node][i].weight+weight {
				dis_from_end[graph[node][i].next] = graph[node][i].weight + weight
				heap.Push(mh, Pair{next: graph[node][i].next, weight: dis_from_end[graph[node][i].next]})
			}
		}
	}

	restricted_path := make([]int, len(dis_from_end))
	restricted_path[n] = 1
	list := [][2]int{}
	for i := 1; i < len(dis_from_end); i++ {
		list = append(list, [2]int{dis_from_end[i], i})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i][0] < list[j][0]
	})
	for _, v := range list {
		for _, e := range graph[v[1]] {
			if dis_from_end[e.next] > dis_from_end[v[1]] {
				restricted_path[e.next] = (restricted_path[e.next] + restricted_path[v[1]]) % 1000000007
			}
		}
	}
	return restricted_path[1]
}

func main() {
	fmt.Println(restricted_paths(5, [][]int{
		{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10},
	}))
}
