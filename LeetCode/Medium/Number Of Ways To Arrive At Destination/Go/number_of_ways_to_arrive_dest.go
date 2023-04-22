package main

import (
	"container/heap"
	"fmt"
	"math"
)

const MOD = 1e9 + 7

type Pair struct {
	val    int
	weight int
}

type P []Pair

func (mh P) Len() int              { return len(mh) }
func (mh P) Less(i, j int) bool    { return mh[i].weight < mh[j].weight }
func (mh P) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *P) Push(val interface{}) { *mh = append(*mh, val.(Pair)) }
func (mh *P) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func make_graph(edges [][]int) map[int][]Pair {
	g := make(map[int][]Pair)
	n := len(edges)
	for i := 0; i < n; i++ {
		g[edges[i][0]] = append(g[edges[i][0]], Pair{val: edges[i][1], weight: edges[i][2]})
		g[edges[i][1]] = append(g[edges[i][1]], Pair{val: edges[i][0], weight: edges[i][2]})
	}
	return g
}

func number_of_ways(n int, edges [][]int) int {
	graph := make_graph(edges)
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt64
	}

	mh := &P{}
	dis[0] = 0
	heap.Push(mh, Pair{val: 0, weight: 0})

	count := make([]int, n)
	count[0] = 1
	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Pair)
		node := curr.val
		weight := curr.weight
		for i := 0; i < len(graph[node]); i++ {
			if dis[graph[node][i].val] > graph[node][i].weight+weight {
				dis[graph[node][i].val] = graph[node][i].weight + weight
				count[graph[node][i].val] = count[node]
				heap.Push(mh, Pair{val: graph[node][i].val, weight: dis[graph[node][i].val]})
			} else if dis[graph[node][i].val] == graph[node][i].weight+weight {
				count[graph[node][i].val] = count[graph[node][i].val] + count[node]
			}
		}
	}

	return count[n-1]
}

func main() {
	fmt.Println(number_of_ways(7, [][]int{
		{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3}, {3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2},
	}))
}
