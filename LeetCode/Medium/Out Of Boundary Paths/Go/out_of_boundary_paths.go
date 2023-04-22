package main

import "fmt"

var mod = int(1e9 + 7)

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

//---------------------brute force---------------------------
func out_of_boundary_path(m, n, max_move, r, c int) int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}
	return _recur(arr, m, n, max_move, r, c)
}

func _recur(arr [][]int, m, n, max_move, r, c int) int {

	if max_move < 0 {
		return 0
	}

	if r >= m || c >= n || r < 0 || c < 0 {
		return 1
	}

	res := 0
	for i := 0; i < 4; i++ {
		x, y := dir[i][0]+r, dir[i][1]+c
		res = (res + _recur(arr, m, n, max_move-1, x, y)) % mod
	}
	return res % mod
}

//---------------------brute force---------------------------

//---------------------efficient approach---------------------------
func out_of_boundary_path_eff(m, n, max_move, r, c int) int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}

	memo := make([][][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int, max_move+1)
			for k := 0; k < max_move+1; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	return _memo(arr, &memo, m, n, max_move, r, c)
}

func _memo(arr [][]int, memo *[][][]int, m, n, max_move, r, c int) int {
	if max_move < 0 {
		return 0
	}

	if r >= m || c >= n || r < 0 || c < 0 {
		return 1
	}

	if (*memo)[r][c][max_move] != -1 {
		return (*memo)[r][c][max_move]
	}

	res := 0
	for i := 0; i < 4; i++ {
		x, y := r+dir[i][0], c+dir[i][1]
		res = (res + _memo(arr, memo, m, n, max_move-1, x, y)) % mod
	}
	(*memo)[r][c][max_move] = res % mod
	return (*memo)[r][c][max_move]
}

//---------------------efficient approach---------------------------

func main() {
	fmt.Println(out_of_boundary_path_eff(2, 2, 2, 0, 0))
}
