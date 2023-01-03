package main

import "fmt"

func check_if_there_is_a_path_with_equal_number_of_zeroes_and_ones(grid [][]int) bool {
	memo := [100][100][101]bool{}
	return (len(grid)+len(grid[0]))%2 == 1 && dfs(0, 0, len(grid), len(grid[0]), grid, 0, &memo)
}

func dfs(i, j, n, m int, grid [][]int, cnt int, memo *[100][100][101]bool) bool {
	if i == n || j == m || cnt > (n+m-1)/2 || (*memo)[i][j][cnt] {
		return false
	}
	(*memo)[i][j][cnt] = true
	cnt += grid[i][j]
	if i == n-1 && j == m-1 {
		return cnt == (n+m-1)/2
	}
	return dfs(i+1, j, n, m, grid, cnt, memo) || dfs(i, j+1, n, m, grid, cnt, memo)
}

func main() {
	fmt.Println(check_if_there_is_a_path_with_equal_number_of_zeroes_and_ones([][]int{
		{0, 1, 0, 0}, {0, 1, 0, 0}, {1, 0, 1, 0},
	}))
}
