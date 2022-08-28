package main

import (
	"container/heap"
	"fmt"
)

type I []int

func (a I) Len() int              { return len(a) }
func (a I) Less(i, j int) bool    { return a[i] < a[j] }
func (a I) Swap(i, j int)         { a[i], a[j] = a[j], a[i] }
func (a *I) Push(val interface{}) { *a = append(*a, val.(int)) }
func (a *I) Pop() interface{} {
	val := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return val
}

func sort_the_matrix_diagonally(arr [][]int) [][]int {
	mp := make(map[int]I)
	n, m := len(arr), len(arr[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if _, ok := mp[i-j]; !ok {
				mh := &I{}
				heap.Init(mh)
				mp[i-j] = *mh
			}
			mh := mp[i-j]
			heap.Push(&mh, arr[i][j])
			mp[i-j] = mh
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mh := mp[i-j]
			arr[i][j] = heap.Pop(&mh).(int)
			mp[i-j] = mh
		}
	}
	return arr
}

func main() {
	fmt.Println(sort_the_matrix_diagonally([][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}))
}
