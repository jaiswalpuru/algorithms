/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.
*/

package main

import "fmt"

func dfs(grid [][]int, visited [][]bool, r, c, i, j int) {
	if i < 0 || j < 0 || i >= r || j >= c || visited[i][j] || grid[i][j] == 0 {
		return
	}
	visited[i][j] = true
	dfs(grid, visited, r, c, i+1, j)
	dfs(grid, visited, r, c, i-1, j)
	dfs(grid, visited, r, c, i, j+1)
	dfs(grid, visited, r, c, i, j-1)
}

func number_of_islands(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				cnt++
				dfs(grid, visited, n, m, i, j)
			}
		}
	}
	return cnt
}

func main() {
	fmt.Println(number_of_islands([][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))
	fmt.Println(number_of_islands([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}))
}
