package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Pair struct {
	en, proc, ind int
}

type M []Pair

func (m M) Len() int { return len(m) }
func (m M) Less(i, j int) bool {
	if m[i].proc == m[j].proc {
		return m[i].ind < m[j].ind
	}
	return m[i].proc < m[j].proc
}
func (m M) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *M) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *M) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func single_threaded_cpu(tasks [][]int) []int {
	t := make([]Pair, len(tasks))
	for i := 0; i < len(tasks); i++ {
		t[i].en = tasks[i][0]
		t[i].proc = tasks[i][1]
		t[i].ind = i
	}
	sort.Slice(t, func(i, j int) bool {
		if t[i].en == t[j].en {
			return t[i].proc < t[j].proc
		}
		return t[i].en < t[j].en
	})
	h := &M{}
	heap.Push(h, t[0])
	i := 1
	res := []int{}
	total := t[0].en
	for h.Len() > 0 {
		curr := heap.Pop(h).(Pair)
		res = append(res, curr.ind)
		total += curr.proc
		for i < len(t) && t[i].en <= total {
			heap.Push(h, t[i])
			i++
		}
		if h.Len() == 0 && i < len(t) {
			heap.Push(h, t[i])
			total += t[i].en
			i++
		}
	}
	return res
}

func main() {
	fmt.Println(single_threaded_cpu([][]int{
		{1, 2}, {2, 4}, {3, 2}, {4, 1},
	}))
}
