/*
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water),
return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
You may assume all four edges of the grid are all surrounded by water.

Example 1:
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1

Example 2:
Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3
*/

package main

import "fmt"

func num_of_islands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	islands := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				islands++
				dfs(grid, i, j, visited, n, m)
			}
		}
	}
	return islands
}

func dfs(grid [][]byte, i, j int, visited [][]bool, row, col int) {
	if i < 0 || j < 0 || i >= row || j >= col || grid[i][j] == '0' || visited[i][j] {
		return
	}
	visited[i][j] = true
	dfs(grid, i+1, j, visited, row, col)
	dfs(grid, i-1, j, visited, row, col)
	dfs(grid, i, j+1, visited, row, col)
	dfs(grid, i, j-1, visited, row, col)
}

func main() {
	fmt.Println(num_of_islands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}
