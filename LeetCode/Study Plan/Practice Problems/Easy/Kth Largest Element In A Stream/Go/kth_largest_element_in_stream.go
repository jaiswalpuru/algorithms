package main

import (
	"container/heap"
	"fmt"
)

type I []int

func (mh I) Len() int              { return len(mh) }
func (mh I) Less(i, j int) bool    { return mh[i] < mh[j] }
func (mh I) Swap(i, j int)         { mh[i], mh[j] = mh[j], mh[i] }
func (mh *I) Push(val interface{}) { *mh = append(*mh, val.(int)) }
func (mh *I) Pop() interface{} {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return val
}

type KthLargest struct {
	max_heap *I
	k        int
}

func Constructor(k int, arr []int) KthLargest {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	mh := &I{}
	n := len(arr)
	for i := 0; i < k && i < n; i++ {
		heap.Push(mh, arr[i])
	}
	return KthLargest{mh, k}
}

func (this *KthLargest) Add(val int) int {
	if this.max_heap.Len() < this.k {
		heap.Push(this.max_heap, val)
	} else if (*this.max_heap)[0] < val {
		heap.Pop(this.max_heap)
		heap.Push(this.max_heap, val)
	}
	return (*this.max_heap)[0]
}

func main() {

}
