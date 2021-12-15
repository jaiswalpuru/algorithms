/*
	Compute running median of sequence of numbers. That is, given a stream of numbers ,
	print out the median of the list so far after each new element
	IP : [2,1,5,7,2,0,5]
	OP :
		2
		1.5
		2
		3.5
		2
		2
		2
*/

package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Len() int           { return len(h) }

func (h *MaxHeap) Push(val interface{}) { *h = append(*h, val.(int)) }

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

type MinHeap []int

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], (h)[j] = h[j], h[i] }
func (h MinHeap) Len() int           { return len(h) }

func (h *MinHeap) Push(val interface{}) { *h = append(*h, val.(int)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func get_median(min_heap MinHeap, max_heap MaxHeap) float64 {
	if len(min_heap) > len(max_heap) {
		return float64(min_heap[0])
	} else if len(min_heap) < len(max_heap) {
		return float64(max_heap[0])
	} else {
		return float64(min_heap[0]+max_heap[0]) / 2.0
	}
}

func add(num int, min_heap *MinHeap, max_heap *MaxHeap) {
	if (*min_heap).Len()+(*max_heap).Len() < 1 {
		heap.Push(min_heap, num)
		return
	}

	median := get_median(*min_heap, *max_heap)

	if float64(num) > median {
		heap.Push(min_heap, num)
	} else {
		heap.Push(max_heap, num)
	}
}

func rebalance(min_heap *MinHeap, max_heap *MaxHeap) {
	if len(*min_heap) > len(*max_heap)+1 {
		root := heap.Pop(min_heap).(int)
		heap.Push(max_heap, root)
	} else if len(*max_heap) > len(*min_heap)+1 {
		root := heap.Pop(max_heap).(int)
		heap.Push(min_heap, root)
	}
}

func print_median(min_heap MinHeap, max_heap MaxHeap) {
	fmt.Println(get_median(min_heap, max_heap))
}

func main() {
	arr := []int{2, 1, 5, 7, 2, 0, 5}

	max_heap := make(MaxHeap, 0)
	min_heap := make(MinHeap, 0)

	heap.Init(&min_heap)
	heap.Init(&max_heap)

	for i := 0; i < len(arr); i++ {
		add(arr[i], &min_heap, &max_heap)
		rebalance(&min_heap, &max_heap)
		print_median(min_heap, max_heap)
	}
}
