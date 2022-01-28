package main

import "fmt"

func maze_runner(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	return _maze_runner(grid, m-1, n-1)
}

func _maze_runner(grid [][]int, m, n int) int {
	if m == 0 && n == 0 {
		if grid[m][n] == -1 {
			return 0
		}
		return 1
	}

	if m < 0 || n < 0 {
		return 0
	}

	if grid[m][n] == -1 {
		return 0
	}

	left := _maze_runner(grid, m, n-1)
	up := _maze_runner(grid, m-1, n)
	return left + up
}

func maze_runner_memo(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1
		}
	}

	_memo(grid, &dp, m-1, n-1)
	return dp[m-1][n-1]
}

func _memo(grid [][]int, dp *[][]int, m, n int) int {
	if m < 0 || n < 0 {
		return 0
	}

	if m == 0 && n == 0 {
		if grid[m][n] == -1 {
			return 0
		}
		return 1
	}

	if grid[m][n] == -1 {
		return 0
	}

	if (*dp)[m][n] != -1 {
		return (*dp)[m][n]
	}

	left := _memo(grid, dp, m, n-1)
	up := _memo(grid, dp, m-1, n)
	(*dp)[m][n] = left + up
	return (*dp)[m][n]
}

func main() {
	fmt.Println(maze_runner_memo([][]int{
		{0, 0, 0},
		{0, -1, 0},
		{0, 0, 0},
	}))
}
