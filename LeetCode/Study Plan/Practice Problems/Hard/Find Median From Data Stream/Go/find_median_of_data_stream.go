package main

import (
	"container/heap"
	"sort"
)

//----------------------sorting when finding the median---------------
type MedianFinder struct {
	arr []int
}

func Constructor() MedianFinder {
	return MedianFinder{arr: []int{}}
}

func (this *MedianFinder) AddNum(num int) {
	this.arr = append(this.arr, num)
}

func (this *MedianFinder) FindMedian() float64 {
	sort.Ints(this.arr)
	if len(this.arr)%2 == 0 {
		return float64(this.arr[len(this.arr)/2]+this.arr[len(this.arr)/2-1]) / 2
	} else {
		return float64(this.arr[len(this.arr)/2])
	}
}

//----------------------sorting when finding the median---------------

//---------------------using two heaps----------------------------
type Min []int

func (m Min) Len() int              { return len(m) }
func (m Min) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m Min) Less(i, j int) bool    { return m[i] < m[j] }
func (m *Min) Push(val interface{}) { *m = append(*m, val.(int)) }
func (m *Min) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

type Max []int

func (m Max) Len() int              { return len(m) }
func (m Max) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m Max) Less(i, j int) bool    { return m[i] > m[j] }
func (m *Max) Push(val interface{}) { *m = append(*m, val.(int)) }
func (m *Max) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

type MedianFinder struct {
	min_heap *Min
	max_heap *Max
}

func Constructor() MedianFinder {
	max_heap := &Max{}
	min_heap := &Min{}
	return MedianFinder{max_heap: max_heap, min_heap: min_heap}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.min_heap, num)
	val := heap.Pop(this.min_heap).(int)
	heap.Push(this.max_heap, val)
	if this.min_heap.Len() < this.max_heap.Len() {
		val = heap.Pop(this.max_heap).(int)
		heap.Push(this.min_heap, val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.min_heap.Len() > this.max_heap.Len() {
		val := heap.Pop(this.min_heap).(int)
		heap.Push(this.min_heap, val)
		return float64(val)
	}
	v1, v2 := heap.Pop(this.min_heap).(int), heap.Pop(this.max_heap).(int)
	heap.Push(this.min_heap, v1)
	heap.Push(this.max_heap, v2)
	return float64(v1+v2) / 2
}

//---------------------using two heaps----------------------------
