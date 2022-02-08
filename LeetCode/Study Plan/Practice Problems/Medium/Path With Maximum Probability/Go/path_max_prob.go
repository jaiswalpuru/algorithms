package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	val  int
	prob float64
}

type MaxHeap []Pair

func (mh MaxHeap) Len() int              { return len(mh) }
func (mh MaxHeap) Less(i, j int) bool    { return mh[i].prob > mh[j].prob }
func (mh MaxHeap) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MaxHeap) Push(val interface{}) { *mh = append(*mh, val.(Pair)) }
func (mh *MaxHeap) Pop() interface{} {
	old := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return old
}

func make_graph(edges [][]int, prob []float64) map[int][]Pair {
	graph := make(map[int][]Pair)

	n := len(edges)

	for i := 0; i < n; i++ {
		graph[edges[i][0]] = append(graph[edges[i][0]], Pair{val: edges[i][1], prob: prob[i]})
		graph[edges[i][1]] = append(graph[edges[i][1]], Pair{val: edges[i][0], prob: prob[i]})
	}
	return graph
}

func path_max_prob(n int, edges [][]int, prob []float64, start, end int) float64 {
	graph := make_graph(edges, prob)

	visited := make(map[int]bool)

	mh := &MaxHeap{}
	heap.Push(mh, Pair{val: start, prob: 1.0})
	for mh.Len() > 0 {
		curr := heap.Pop(mh).(Pair)
		visited[curr.val] = true
		if curr.val == end {
			return curr.prob
		}
		for i := 0; i < len(graph[curr.val]); i++ {
			if !visited[graph[curr.val][i].val] {
				heap.Push(mh, Pair{val: graph[curr.val][i].val, prob: curr.prob * graph[curr.val][i].prob})
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(path_max_prob(3, [][]int{
		{0, 1}, {1, 2}, {0, 2},
	}, []float64{0.5, 0.5, 0.2}, 0, 2))
}
