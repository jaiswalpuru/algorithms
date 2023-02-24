package main

import (
	"container/heap"
	"fmt"
	"math"
)

type M []int

func (m M) Len() int              { return len(m) }
func (m M) Less(i, j int) bool    { return m[i] > m[j] }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(int)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func minimumDeviation(nums []int) int {
	size := len(nums)
	maxHeap := &M{}
	minEvenElement := math.MaxInt64
	minDeviation := math.MaxInt64

	for i := 0; i < size; i++ {
		if nums[i]%2 == 0 {
			minEvenElement = min(minEvenElement, nums[i])
			heap.Push(maxHeap, nums[i])
		} else {
			minEvenElement = min(minEvenElement, nums[i]*2)
			heap.Push(maxHeap, nums[i]*2)
		}
	}

	for maxHeap.Len() > 0 {
		topElement := heap.Pop(maxHeap).(int)
		minDeviation = min(minDeviation, topElement-minEvenElement)
		if topElement%2 == 0 {
			heap.Push(maxHeap, topElement/2)
			minEvenElement = min(minEvenElement, topElement/2)
		} else {
			break
		}
	}

	return minDeviation
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimumDeviation([]int{1, 2, 3, 4}))
}
