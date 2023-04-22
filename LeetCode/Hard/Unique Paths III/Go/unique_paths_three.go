package main

import "fmt"

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func unique_paths_three(grid [][]int) int {
	r, c := 0, 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				r, c = i, j
				break
			}
		}
	}
	obstacles := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == -1 {
				obstacles++
			}
		}
	}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	res := 0
	backtrack(grid, r, c, &visited, &res, n, m, obstacles, 1)
	return res
}

func backtrack(grid [][]int, r, c int, visited *[][]bool, res *int, n, m, obs int, t int) {
	if r < 0 || c < 0 || r >= n || c >= m || (*visited)[r][c] {
		return
	}
	if grid[r][c] == -1 {
		t++
		return
	}
	if grid[r][c] == 2 {
		if t == (n*m - obs) {
			*res++
		}
		return
	}
	(*visited)[r][c] = true
	for i := 0; i < 4; i++ {
		x, y := r+dir[i][0], c+dir[i][1]
		backtrack(grid, x, y, visited, res, n, m, obs, t+1)
	}
	(*visited)[r][c] = false
}

func main() {
	fmt.Println(unique_paths_three([][]int{
		{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1},
	}))
}
