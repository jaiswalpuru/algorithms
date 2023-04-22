package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var dir = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func max_area_island(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				cnt := 0
				dfs(grid, i, j, &visited, &cnt, n, m)
				res = max(res, cnt)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int, visited *[][]bool, cnt *int, n, m int) {
	if i >= n || j >= m || i < 0 || j < 0 || (*visited)[i][j] || grid[i][j] == 0 {
		return
	}
	(*visited)[i][j] = true
	*cnt++
	for k := 0; k < len(dir); k++ {
		x, y := i+dir[k][0], j+dir[k][1]
		dfs(grid, x, y, visited, cnt, n, m)
	}
}

func main() {
	fmt.Println(max_area_island([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}))
}
