package main

import "fmt"

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func island_perimeter(arr [][]int) int {
	n, m := len(arr), len(arr[0])
	r, c := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 1 {
				r, c = i, j
				break
			}
		}
	}
	res := 0
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	recur(r, c, arr, &res, &visited)
	return res
}

func recur(r, c int, arr [][]int, res *int, visited *[][]bool) {
	(*visited)[r][c] = true
	for i := 0; i < 4; i++ {
		x, y := dir[i][0]+r, dir[i][1]+c
		if !is_safe(x, y, len(arr), len(arr[0]), arr) {
			*res++
			continue
		}
		if !(*visited)[x][y] {
			recur(x, y, arr, res, visited)
		}
	}
}

func is_safe(i, j, n, m int, arr [][]int) bool {
	return !(i < 0 || j < 0 || i >= n || j >= m || arr[i][j] == 0)
}

func main() {
	fmt.Println(island_perimeter([][]int{
		{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0},
	}))
}
