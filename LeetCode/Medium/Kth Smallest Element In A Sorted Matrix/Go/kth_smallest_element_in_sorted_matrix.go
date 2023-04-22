package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//-----------------------------brute force------------------------------
func kth_smallest_element_in_sorted_matrix(arr [][]int, k int) int {
	n, m := len(arr), len(arr[0])
	l := 0
	res := make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[l] = arr[i][j]
			l++
		}
	}
	sort.Ints(res)
	return res[k-1]
}

//-----------------------------brute force------------------------------

//-----------------------------using min heap---------------------------
type Trip struct{ val, r, c int }
type MH []Trip

func (m MH) Less(i, j int) bool    { return m[i].val < m[j].val }
func (m MH) Swap(i, j int)         { m[i], m[j] = m[j], m[i] }
func (m MH) Len() int              { return len(m) }
func (m *MH) Push(val interface{}) { *m = append(*m, val.(Trip)) }
func (m *MH) Pop() interface{} {
	val := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return val
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func kth_smallest_element_in_sorted_matrix_min_heap(arr [][]int, k int) int {
	n := len(arr)
	mh := make(MH, 0)
	heap.Init(&mh)
	for i := 0; i < min(n, k); i++ {
		heap.Push(&mh, Trip{arr[i][0], i, 0})
	}
	ele := mh[0]
	for k > 0 {
		ele = heap.Pop(&mh).(Trip)
		r, c := ele.r, ele.c
		if c < n-1 {
			heap.Push(&mh, Trip{arr[r][c+1], r, c + 1})
		}
		k--
	}
	return ele.val
}

//-----------------------------using min heap---------------------------

func main() {
	fmt.Println(kth_smallest_element_in_sorted_matrix_min_heap([][]int{
		{1, 5, 9}, {10, 11, 13}, {12, 13, 15},
	}, 8))
}
