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
	minHeap *Min
	maxHeap *Max
}

func Constructor() MedianFinder {
	return MedianFinder{&Min{}, &Max{}}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.minHeap, num)
	heap.Push(this.maxHeap, heap.Pop(this.minHeap).(int))
	if this.minHeap.Len() < this.maxHeap.Len() {
		heap.Push(this.minHeap, heap.Pop(this.maxHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	var median float64
	if this.minHeap.Len() > this.maxHeap.Len() {
		val := heap.Pop(this.minHeap).(int)
		heap.Push(this.minHeap, val)
		median = float64(val)
	} else {
		v1, v2 := heap.Pop(this.minHeap).(int), heap.Pop(this.maxHeap).(int)
		median = float64(v1+v2) / 2
		heap.Push(this.minHeap, v1)
		heap.Push(this.maxHeap, v2)
	}
	return median
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
//---------------------using two heaps----------------------------
