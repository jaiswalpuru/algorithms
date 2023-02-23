package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Pair struct{ profit, capital int }

type M []int

func (m M) Len() int      { return len(m) }
func (m M) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m M) Less(i, j int) bool {
	return m[i] > m[j]
}
func (m *M) Push(val interface{}) { *m = append(*m, val.(int)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	size := len(profits)
	project := make([]Pair, size)

	for i := 0; i < size; i++ {
		project[i].profit = profits[i]
		project[i].capital = capital[i]
	}

	sort.Slice(project, func(i, j int) bool {
		return project[i].capital < project[j].capital
	})

	mh := &M{}

	ptr := 0
	for i := 0; i < k; i++ {
		for ptr < size && project[ptr].capital <= w {
			heap.Push(mh, project[ptr].profit)
			ptr++
		}
		if mh.Len() == 0 {
			break
		}
		w += heap.Pop(mh).(int)
	}
	return w
}

func main() {
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
}
