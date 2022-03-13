package main

import (
	"fmt"
)

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//---------------------------brute force (TLE)---------------------------
func path_with_max_min_val(arr [][]int) int {
	n, m := len(arr), len(arr[0])

	score := min(arr[0][0], arr[n-1][m-1]) //always

	for score >= 0 {
		if path_exists(arr, score) {
			return score
		} else {
			score -= 1
		}
	}
	return -1
}

func path_exists(arr [][]int, score int) bool {
	n, m := len(arr), len(arr[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	q := [][]int{{0, 0}}
	visited[0][0] = true
	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]
		if r == n-1 && c == m-1 {
			return true
		}
		for i := 0; i < 4; i++ {
			nr, nc := r+dir[i][0], c+dir[i][1]
			if nr >= 0 && nc >= 0 && nr < n && nc < m && !visited[nr][nc] && arr[nr][nc] >= score {
				visited[nr][nc] = true
				q = append(q, []int{nr, nc})
			}
		}
	}
	return false
}

//---------------------------------------------------------------------------------

//---------------------------Effcient approach (BFS+binary search)---------------------------

func path_with_max_min_val_eff(arr [][]int) int {
	n, m := len(arr), len(arr[0])

	l, r := 0, min(arr[0][0], arr[n-1][m-1]) //always

	for l < r {
		mid := (l + r + 1) / 2
		if path_exists(arr, mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

//---------------------------------------------------------------------------------------------

func main() {
	fmt.Println(path_with_max_min_val([][]int{
		{3, 4, 6, 3, 4}, {0, 2, 1, 1, 7}, {8, 8, 3, 2, 7}, {3, 2, 4, 9, 8}, {4, 1, 2, 0, 0}, {4, 6, 5, 4, 3},
	}))
}
