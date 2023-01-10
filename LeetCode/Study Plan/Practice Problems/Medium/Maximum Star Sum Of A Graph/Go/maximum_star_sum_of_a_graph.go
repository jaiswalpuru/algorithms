package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Pair struct{ v, wt int }
type P []Pair

func (m P) Len() int              { return len(m) }
func (m P) Less(i, j int) bool    { return m[i].wt > m[j].wt }
func (m P) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m *P) Push(val interface{}) { *m = append(*m, val.(Pair)) }
func (m *P) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func maximum_star_sum_of_a_graph(vals []int, edges [][]int, k int) int {
	g := make(map[int]*P)
	for i := 0; i < len(edges); i++ {
		if _, ok := g[edges[i][0]]; !ok {
			g[edges[i][0]] = &P{}
		}
		if _, ok := g[edges[i][1]]; !ok {
			g[edges[i][1]] = &P{}
		}
		heap.Push(g[edges[i][0]], Pair{edges[i][1], vals[edges[i][1]]})
		heap.Push(g[edges[i][1]], Pair{edges[i][0], vals[edges[i][0]]})
	}
	res := math.MinInt64
	for i := 0; i < len(vals); i++ {
		res = max(res, vals[i])
	}
	for key, val := range g {
		curr := val
		k_temp := k
		sum := vals[key]
		for curr.Len() > 0 && k_temp > 0 {
			sum += heap.Pop(curr).(Pair).wt
			k_temp--
			res = max(res, sum)
		}
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximum_star_sum_of_a_graph([]int{1, 2, 3, 4, 10, -10, -20}, [][]int{
		{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}, {3, 6},
	}, 2))
}
