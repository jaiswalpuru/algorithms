package main

import (
	"container/heap"
	"fmt"
	"math"
)

//-----------------------using bellman ford---------------------
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	currDis := make([]int, n)
	prevDis := make([]int, n)
	for i := 0; i < n; i++ {
		currDis[i] = int(1e7)
		prevDis[i] = int(1e7)
	}
	prevDis[src] = 0
	// as we need to find the cheapest flights within k stops
	// so we try to relax every edges atmost k times
	for i := 0; i <= k; i++ {
		currDis[src] = 0
		for j := 0; j < len(flights); j++ {
			from := flights[j][0]
			to := flights[j][1]
			cost := flights[j][2]
			if prevDis[from] < int(1e7) {
				currDis[to] = min(currDis[to], prevDis[from]+cost)
			}
		}
		prevDis = make([]int, n)
		copy(prevDis, currDis)
	}
	if prevDis[dst] == int(1e7) {
		return -1
	}
	return prevDis[dst]
}

//-----------------------using bellman ford---------------------

type Pair struct {
	next,
	weight,
	stops int
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

func makeGraph(flights [][]int) map[int][]Pair {
	n := len(flights)
	graph := make(map[int][]Pair)
	for i := 0; i < n; i++ {
		graph[flights[i][0]] = append(graph[flights[i][0]], Pair{next: flights[i][1], weight: flights[i][2]})
	}
	return graph
}

func findCheapestPriceDijkstra(n int, flights [][]int, src, dst, k int) int {
	graph := makeGraph(flights)

	mh := &P{}
	heap.Push(mh, Pair{next: src, weight: 0, stops: 0})

	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt64
	}
	count := make([]int, n)
	dis[src] = 0
	count[src] = 0
	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Pair)
		u := curr.next
		w := curr.weight
		stop := curr.stops
		if u == dst {
			return w
		}

		if stop == k+1 {
			continue
		}
		count[u] = stop
		for v := 0; v < len(graph[u]); v++ {
			if dis[graph[u][v].next] > graph[u][v].weight+w {
				dis[graph[u][v].next] = graph[u][v].weight + w
				heap.Push(mh, Pair{next: graph[u][v].next, weight: dis[graph[u][v].next], stops: stop + 1})
			} else if stop < count[dis[graph[u][v].next]] {
				heap.Push(mh, Pair{next: graph[u][v].next, weight: graph[u][v].weight + w, stops: stop + 1})
			}
		}
	}
	if dis[dst] == math.MaxInt64 {
		return -1
	}
	return dis[dst]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(findCheapestPrice(4, [][]int{
		{0, 1, 1},
		{0, 2, 5},
		{1, 2, 1},
		{2, 3, 1},
	}, 0, 3, 1))
}
