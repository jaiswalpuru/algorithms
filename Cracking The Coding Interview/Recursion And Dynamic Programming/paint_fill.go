package main

import "fmt"

func paint_fill(arr [][]int, r, c int, n_color int) bool {
	if arr[r][c] == n_color {
		return false
	}
	return recur(arr, r, c, arr[r][c], n_color)
}

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func recur(arr [][]int, r, c, old, n_ew int) bool {
	if r < 0 || c < 0 || r >= len(arr) || c >= len(arr[0]) {
		return false
	}
	if arr[r][c] == old {
		arr[r][c] = n_ew
		for i := 0; i < 4; i++ {
			x, y := dir[i][0]+r, dir[i][1]+c
			recur(arr, x, y, old, n_ew)
		}
	}
	return true
}

func main() {
	fmt.Println(paint_fill([][]int{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
	}, 0, 1, 1))
}
