package main

import (
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

func main() {
	fmt.Println(kth_smallest_element_in_sorted_matrix([][]int{
		{1, 5, 9}, {10, 11, 13}, {12, 13, 15},
	}, 8))
}
