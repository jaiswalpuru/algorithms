package main

import (
	"container/heap"
	"fmt"
)

type M []string

func (m M) Len() int      { return len(m) }
func (m M) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m M) Less(i, j int) bool {
	if len(m[i]) == len(m[j]) {
		return m[i] > m[j]
	}
	return len(m[i]) > len(m[j])
}
func (m *M) Push(val interface{}) { *m = append(*m, val.(string)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func kthLargestNumber(nums []string, k int) string {
	maxHeap := &M{}
	res := ""
	for i := 0; i < len(nums); i++ {
		heap.Push(maxHeap, nums[i])
	}
	for k > 0 {
		res = heap.Pop(maxHeap).(string)
		k--
	}
	return res
}

func main() {
	fmt.Println(kthLargestNumber([]string{"3", "6", "7", "10"}, 4))
}
