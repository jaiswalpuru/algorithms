package main

import "fmt"

var dir = [][]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}}

//-------------------------brute force----------------------------
func longest_line_of_consecutive_one_in_matrix(arr [][]int) int {
	n, m := len(arr), len(arr[0])
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 1 {
				for k := 0; k < 4; k++ {
					cnt = max(cnt, recur(arr, i, j, dir[k]))
				}
			}
		}
	}
	return cnt
}

func is_safe(i, j, n, m int) bool {
	return !(i < 0 || j < 0 || i >= n || j >= m)
}

func recur(arr [][]int, i, j int, d []int) int {
	if is_safe(i, j, len(arr), len(arr[0])) && arr[i][j] == 1 {
		return 1 + recur(arr, i+d[0], j+d[1], d)
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-------------------------brute force----------------------------

//-------------------------memo----------------------------
func longest_line_of_consecutive_one_in_matrix_memo(arr [][]int) int {
	n, m := len(arr), len(arr[0])
	cnt := 0
	memo := make([][][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = make([]int, 4)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 1 {
				for k := 0; k < 4; k++ {
					cnt = max(cnt, _memo(arr, i, j, k, &memo))
				}
			}
		}
	}
	return cnt
}

func _memo(arr [][]int, i, j, k int, memo *[][][]int) int {
	if is_safe(i, j, len(arr), len(arr[0])) && arr[i][j] == 1 {
		if (*memo)[i][j][k] != 0 {
			return (*memo)[i][j][k]
		}
		val := 1 + _memo(arr, i+dir[k][0], j+dir[k][1], k, memo)
		(*memo)[i][j][k] = val
		return (*memo)[i][j][k]
	}
	return 0
}

//-------------------------memo----------------------------

func main() {
	fmt.Println(longest_line_of_consecutive_one_in_matrix_memo([][]int{
		{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 1},
	}))
}
