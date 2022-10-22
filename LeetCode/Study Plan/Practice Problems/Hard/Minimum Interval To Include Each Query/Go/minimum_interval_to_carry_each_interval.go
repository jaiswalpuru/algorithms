package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

//---------------------------brute force--------------------------------
func minimum_interval_to_carry_each_interval(arr [][]int, q []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	n, m := len(q), len(arr)
	res := []int{}
	for i := 0; i < n; i++ {
		curr := q[i]
		min_val := math.MaxInt64
		for j := 0; j < m; j++ {
			if arr[j][0] > curr {
				break
			}
			if arr[j][0] <= curr && arr[j][1] >= curr {
				min_val = min(min_val, arr[j][1]-arr[j][0]+1)
			}
		}
		if min_val == math.MaxInt64 {
			res = append(res, -1)
		} else {
			res = append(res, min_val)
		}
	}
	return res
}

//---------------------------brute force--------------------------------

//---------------------------efficient--------------------------------
type Pair struct{ size, end int }
type P []Pair

func (a P) Len() int           { return len(a) }
func (a P) Less(i, j int) bool { return a[i].size < a[j].size }
func (a P) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a *P) Pop() interface{} {
	val := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return val
}
func (a *P) Push(val interface{}) { *a = append(*a, val.(Pair)) }

func minimum_interval_to_carry_each_interval_eff(arr [][]int, q []int) []int {
	n, m := len(arr), len(q)
	res := make([]int, m)
	type Q struct{ val, ind int }
	query_copy := make([]Q, m)
	for i := 0; i < m; i++ {
		query_copy[i] = Q{val: q[i], ind: i}
	}
	sort.Slice(query_copy, func(i, j int) bool {
		return query_copy[i].val < query_copy[j].val
	})
	sort.Slice(arr, func(i, j int) bool { return arr[i][0] < arr[j][0] })
	mh := &P{}
	heap.Init(mh)
	i := 0
	for j := 0; j < m; j++ {
		v, ind := query_copy[j].val, query_copy[j].ind
		for i < n && arr[i][0] <= v {
			heap.Push(mh, Pair{arr[i][1] - arr[i][0] + 1, arr[i][1]})
			i++
		}
		for mh.Len() > 0 && (*mh)[0].end < v {
			heap.Pop(mh)
		}
		if mh.Len() == 0 {
			res[ind] = -1
		} else {
			res[ind] = (*mh)[0].size
		}
	}
	return res
}

//---------------------------efficient--------------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimum_interval_to_carry_each_interval_eff([][]int{
		{1, 4}, {2, 4}, {3, 6}, {4, 4},
	}, []int{2, 3, 4, 5}))
}
