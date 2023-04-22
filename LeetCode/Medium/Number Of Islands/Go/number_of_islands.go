package main

import "fmt"

func number_of_islands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	islands := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				islands++
				dfs(grid, i, j, &visited, m, n)
			}
		}
	}
	return islands
}

func dfs(grid [][]byte, i, j int, visited *[][]bool, m, n int) {
	if i >= m || j >= n || i < 0 || j < 0 || (*visited)[i][j] || grid[i][j] == '0' {
		return
	}
	(*visited)[i][j] = true
	dfs(grid, i+1, j, visited, m, n)
	dfs(grid, i, j+1, visited, m, n)
	dfs(grid, i-1, j, visited, m, n)
	dfs(grid, i, j-1, visited, m, n)
}

func main() {
	fmt.Println(number_of_islands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}
