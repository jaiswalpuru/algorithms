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

var (
	ROW, COL int
)

func dfs(grid [][]int, row, col int, visited [][]bool) {
	if row < 0 || row >= ROW || col < 0 || col >= COL || (grid[row][col] == 0) || visited[row][col] {
		return
	}
	visited[row][col] = true
	dfs(grid, row+1, col, visited)
	dfs(grid, row-1, col, visited)
	dfs(grid, row, col+1, visited)
	dfs(grid, row, col-1, visited)
}

func num_of_island(grid [][]int) int {
	ROW = len(grid)
	COL = len(grid[0])
	visited := make([][]bool, ROW)
	for i := 0; i < ROW; i++ {
		visited[i] = make([]bool, COL)
	}

	cnt := 0

	for i := 0; i < ROW; i++ {
		for j := 0; j < COL; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				cnt++
				dfs(grid, i, j, visited)
			}
		}
	}

	return cnt
}

func main() {
	fmt.Println(num_of_island([][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}))

	fmt.Println(num_of_island([][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}))
}
