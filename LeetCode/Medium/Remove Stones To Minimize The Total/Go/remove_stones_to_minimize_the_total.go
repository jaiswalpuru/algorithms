package main

import (
	"container/heap"
	"fmt"
	"math"
)

type M []int

func (m M) Len() int              { return len(m) }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m M) Less(i, j int) bool    { return m[i] > m[j] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(int)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func remove_stones_to_minimize_the_total(piles []int, k int) int {
	m := &M{}
	for i := 0; i < len(piles); i++ {
		heap.Push(m, piles[i])
	}
	for k > 0 {
		curr := heap.Pop(m).(int)
		curr = curr - int(math.Floor(float64(curr)/2))
		heap.Push(m, curr)
		k--
	}
	res := 0
	for m.Len() > 0 {
		curr := heap.Pop(m).(int)
		res += curr
	}
	return res
}

func main() {
	fmt.Println(remove_stones_to_minimize_the_total([]int{5, 9, 2}, 2))
}
