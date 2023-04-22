package main

import (
	"container/heap"
	"fmt"
	"sort"
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

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	m := &M{}
	heap.Push(m, intervals[0][1])
	minRooms := 1
	for i := 1; i < len(intervals); i++ {
		top := heap.Pop(m).(int)
		if top > intervals[i][0] {
			heap.Push(m, top)
			minRooms++
		}
		heap.Push(m, intervals[i][1])
	}
	return minRooms
}

func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}
