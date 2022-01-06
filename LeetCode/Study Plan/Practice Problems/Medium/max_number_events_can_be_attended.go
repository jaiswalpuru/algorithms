/*
You are given an array of events where events[i] = [startDayi, endDayi]. Every event i starts at startDayi and ends at endDayi.

You can attend an event i at any day d where startTimei <= d <= endTimei. You can only attend one event at any time d.

Return the maximum number of events you can attend.
*/

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Item [][]int

func (a Item) Len() int      { return len(a) }
func (a Item) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Item) Less(i, j int) bool {
	if a[i][0] == a[j][0] {
		return a[i][1] < a[j][1]
	}
	return a[i][0] < a[j][0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h *MinHeap) Swap(i, j int)     { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func events(arr [][]int) int {
	n := len(arr)
	item := make(Item, n)
	for i := 0; i < n; i++ {
		item[i] = make([]int, 2)
		item[i][0] = arr[i][0]
		item[i][1] = arr[i][1]
	}
	sort.Sort(item)

	pq := &MinHeap{}
	heap.Init(pq)
	cnt := 0
	j := 0
	for i := 1; i <= 100000; i++ {
		for ; j < n && item[j][0] == i; j++ {
			heap.Push(pq, item[j][1])
		}

		for pq.Len() > 0 {
			item := heap.Pop(pq)
			if i > item.(int) {
				continue
			}
			cnt += 1
			break
		}
	}

	return cnt
}

func main() {
	fmt.Println(events([][]int{
		{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1},
	}))
}
