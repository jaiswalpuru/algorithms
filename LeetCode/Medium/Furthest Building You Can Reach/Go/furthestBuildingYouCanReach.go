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

func furthestBuilding(heights []int, bricks int, ladders int) int {
	m := &M{}
	for i := 0; i < len(heights)-1; i++ {
		jump := heights[i+1] - heights[i]
		if jump <= 0 {
			continue
		}
		heap.Push(m, jump)
		if m.Len() <= ladders {
			continue
		}
		bricks -= heap.Pop(m).(int)
		if bricks < 0 {
			return i
		}
	}
	return len(heights) - 1
}

func main() {
	fmt.Println(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))
}
