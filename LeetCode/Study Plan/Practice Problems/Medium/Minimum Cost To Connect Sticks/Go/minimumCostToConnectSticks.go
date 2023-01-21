package main

import (
	"container/heap"
	"fmt"
)

type M []int

func (m M) Len() int              { return len(m) }
func (m M) Less(i, j int) bool    { return m[i] < m[j] }
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(int)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func connectSticks(sticks []int) int {
	m := &M{}
	for i := 0; i < len(sticks); i++ {
		heap.Push(m, sticks[i])
	}
	cost := 0
	for m.Len() > 1 {
		v1, v2 := heap.Pop(m).(int), heap.Pop(m).(int)
		cost += v1 + v2
		heap.Push(m, v1+v2)
	}
	return cost
}

func main() {
	fmt.Println(connectSticks([]int{2, 4, 3}))
}
