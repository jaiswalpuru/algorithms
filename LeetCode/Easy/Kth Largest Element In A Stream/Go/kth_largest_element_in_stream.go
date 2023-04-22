package main

import (
	"container/heap"
)

type M []int

func (m M) Len() int              { return len(m) }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m M) Less(i, j int) bool    { return m[i] < m[j] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(int)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

type KthLargest struct {
	k  int
	mh *M
}

func Constructor(k int, nums []int) KthLargest {
	m := &M{}
	for i := 0; i < len(nums); i++ {
		heap.Push(m, nums[i])
	}
	for m.Len() > k {
		heap.Pop(m)
	}
	return KthLargest{k: k, mh: m}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.mh, val)
	for this.mh.Len() > this.k {
		heap.Pop(this.mh)
	}
	return (*this.mh)[0]
}
