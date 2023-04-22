package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Pair struct{ v, wt int }

func make_graph(edges [][]int) map[int][]Pair {
	graph := make(map[int][]Pair)
	n := len(edges)
	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], Pair{v: edges[i][1], wt: edges[i][2]})
		graph[edges[i][1]] = append(graph[edges[i][1]], Pair{v: edges[i][0], wt: edges[i][2]})
	}
	return graph
}

type P []Pair

func (mh P) Len() int              { return len(mh) }
func (mh P) Less(i, j int) bool    { return mh[i].wt < mh[j].wt }
func (mh P) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *P) Push(val interface{}) { *mh = append(*mh, val.(Pair)) }
func (mh *P) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func run_dijkstra(node int, graph map[int][]Pair, distance_threshold, n int) int {
	distance := make([]int, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt64
	}

	mh := &P{}
	distance[node] = 0
	heap.Push(mh, Pair{v: node, wt: 0})

	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Pair)
		u := curr.v
		cost := curr.wt

		for i := 0; i < len(graph[u]); i++ {
			v := graph[u][i].v
			cost_v := graph[u][i].wt
			if distance[v] >= cost+cost_v {
				distance[v] = cost + cost_v
				heap.Push(mh, Pair{v: v, wt: distance[v]})
			}
		}
	}
	cities := 0
	for i := 0; i < n; i++ {
		if i != node {
			if distance[i] <= distance_threshold {
				cities++
			}
		}
	}
	return cities
}

func find_the_city(n int, edges [][]int, distance_threshold int) int {
	graph := make_graph(edges)
	mp := make(map[int]int)
	min_city := -1
	mp[-1] = math.MaxInt64
	for i := 0; i < n; i++ {
		mp[i] = run_dijkstra(i, graph, distance_threshold, n)
		if mp[min_city] > mp[i] {
			min_city = i
		} else if mp[min_city] == mp[i] {
			if i > min_city {
				min_city = i
			}
		}
	}
	return min_city
}

func main() {
	fmt.Println(find_the_city(4, [][]int{
		{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1},
	}, 4))
}
