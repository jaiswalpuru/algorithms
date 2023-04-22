package main

import (
	"fmt"
	"sort"
)

type P struct {
	ind      int
	soldiers int
}

type I []P

func (p I) Len() int { return len(p) }
func (p I) Less(i, j int) bool {
	if p[i].soldiers == p[j].soldiers {
		return p[i].ind < p[j].ind
	}
	return p[i].soldiers < p[j].soldiers
}
func (p I) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func kth_weakest_row_in_matrix(arr [][]int, k int) []int {

	n, m := len(arr), len(arr[0])
	a := I{}

	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < m; j++ {
			if arr[i][j] == 1 {
				cnt++
			}
		}
		a = append(a, P{i, cnt})
	}
	fmt.Println(a)

	sort.Sort(a)
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, a[i].ind)
	}
	return res
}

func main() {
	fmt.Println(kth_weakest_row_in_matrix([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
		{0, 0, 0},
		{1, 1, 1},
		{1, 0, 0},
	}, 6))
}
