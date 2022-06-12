package main

import "fmt"
import "math"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//-----------------------------brute force------------------------------
func minimum_path_cost_br(grid [][]int, move_cost [][]int) int {
	res := math.MaxInt64
	for col := 0; col < len(grid[0]); col++ {
		res = min(res, dfs_br(grid, move_cost, 0, col))
	}
	return res
}

func dfs_br(grid, move_cost [][]int, row, col int) int {
	if row >= len(grid) {
		return 0
	}

	res := math.MaxInt64
	for i := 0; i < len(move_cost[0]); i++ {
		t := 0
		if row != len(grid)-1 {
			t = move_cost[grid[row][col]][i]
		}
		res = min(res, dfs_br(grid, move_cost, row+1, i)+t)
	}
	return res + grid[row][col]
}

//-----------------------------brute force------------------------------

//-----------------------------memoization------------------------------
func minimum_path_cost(grid [][]int, move_cost [][]int) int {
	res := math.MaxInt64
	memo := make([][]int, len(grid))

	for i := 0; i < len(grid); i++ {
		memo[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			memo[i][j] = -1
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		res = min(res, dfs(grid, move_cost, 0, i, &memo))
	}
	return res
}

func dfs(grid, move_cost [][]int, r, c int, memo *[][]int) int {
	if r >= len(grid) {
		return 0
	}

	if (*memo)[r][c] != -1 {
		return (*memo)[r][c]
	}

	res := math.MaxInt64
	for i := 0; i < len(move_cost[0]); i++ {
		t := 0
		if r != len(grid)-1 {
			t = move_cost[grid[r][c]][i]
		}
		res = min(res, dfs(grid, move_cost, r+1, i, memo)+t)
	}
	(*memo)[r][c] = res + grid[r][c]
	return (*memo)[r][c]
}

//-----------------------------memoization------------------------------

func main() {
	fmt.Println(minimum_path_cost_br([][]int{
		{5, 3}, {4, 0}, {2, 1},
	}, [][]int{
		{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3},
	}))
}
