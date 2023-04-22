package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Pair struct {
	city int
	cost int
}

func make_graph(edges [][]int) map[int][]Pair {
	graph := make(map[int][]Pair)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], Pair{edges[i][1], edges[i][2]})
		graph[edges[i][1]] = append(graph[edges[i][1]], Pair{edges[i][0], edges[i][2]})
	}
	return graph
}

type P []Pair

func (mh P) Len() int              { return len(mh) }
func (mh P) Less(i, j int) bool    { return mh[i].cost < mh[j].cost }
func (mh P) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *P) Push(val interface{}) { *mh = append(*mh, val.(Pair)) }
func (mh *P) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func connecting_cities_with_min_cost(n int, edges [][]int) int {
	graph := make_graph(edges)
	mh := &P{}
	visited := make(map[int]bool)
	min_cost := make([]int, n+1)
	for i := 0; i <= n; i++ {
		min_cost[i] = math.MaxInt64
	}
	heap.Push(mh, Pair{1, 0}) //pushing the first node with weight zero
	visit, res := 0, 0
	for visit < n && mh.Len() > 0 {
		curr := heap.Pop(mh).(Pair)
		city := curr.city
		cost := curr.cost
		if visited[city] {
			continue
		}
		visit++
		visited[city] = true
		res += cost
		for j := 0; j < len(graph[city]); j++ {
			if !visited[graph[city][j].city] && min_cost[graph[city][j].city] > graph[city][j].cost {
				min_cost[graph[city][j].city] = graph[city][j].cost
				heap.Push(mh, Pair{graph[city][j].city, graph[city][j].cost})
			}
		}
		fmt.Println(mh)
	}
	if visit == n {
		return res
	}
	return -1
}

func main() {
	fmt.Println(connecting_cities_with_min_cost(3, [][]int{
		{1, 2, 7}, {1, 3, 5}, {2, 3, 1},
	}))
}
