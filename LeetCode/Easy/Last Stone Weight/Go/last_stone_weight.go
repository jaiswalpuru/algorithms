package main

import (
	"container/heap"
	"fmt"
)

type I []int

func (mh I) Len() int              { return len(mh) }
func (mh I) Less(i, j int) bool    { return mh[i] > mh[j] }
func (mh I) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *I) Push(val interface{}) { *mh = append(*mh, val.(int)) }
func (mh *I) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

func last_stone_weight(arr []int) int {
	var max_heap I
	max_heap = arr
	heap.Init(&max_heap)
	for max_heap.Len() > 1 {
		v1 := heap.Pop(&max_heap).(int)
		v2 := heap.Pop(&max_heap).(int)
		if v2 != v1 {
			heap.Push(&max_heap, abs(v2-v1))
		}
	}

	if max_heap.Len() == 0 {
		return 0
	}

	return heap.Pop(&max_heap).(int)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	fmt.Println(last_stone_weight([]int{2, 7, 4, 1, 8, 1}))
}
