package main

import "fmt"

func number_of_enclaves(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	cnt := 0
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				t := 0
				if dfs(grid, &visited, i, j, n, m, &t) {
					cnt += t
				}
			}
		}
	}
	return cnt
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dfs(grid [][]int, visited *[][]bool, i, j, n, m int, t *int) bool {
	if i == 0 || j == 0 || i == n-1 || j == m-1 {
		return false
	}
	(*visited)[i][j] = true
	ans := true
	*t += 1
	for k := 0; k < 4; k++ {
		x, y := i+directions[k][0], j+directions[k][1]
		if x >= 0 && y >= 0 && x < n && y < m && grid[x][y] == 1 && !(*visited)[x][y] {
			if !dfs(grid, visited, x, y, n, m, t) {
				ans = false
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(number_of_enclaves([][]int{
		{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0},
	}))
}
