package main

import "fmt"

func number_of_closed_island(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	closed_island := 0
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if !visited[i][j] && grid[i][j] == 0 {
				if dfs(grid, &visited, i, j, n, m) {
					closed_island++
				}
			}
		}
	}
	return closed_island
}

var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func dfs(grid [][]int, visited *[][]bool, i, j, n, m int) bool {
	if i == 0 || j == 0 || j == m-1 || i == n-1 {
		return false
	}
	(*visited)[i][j] = true
	ans := true
	for k := 0; k < 4; k++ {
		x, y := i+directions[k][0], j+directions[k][1]
		if x >= 0 && y >= 0 && x < n && y < m && grid[x][y] == 0 && !(*visited)[x][y] {
			if !dfs(grid, visited, x, y, n, m) {
				ans = false
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(number_of_closed_island([][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}))
}
