package main

import "fmt"

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func search(grid [][]byte, word string) bool {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if word[0] == grid[i][j] {
				if dfs(i, j, grid, &visited, word, 0) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(i, j int, grid [][]byte, visited *[][]bool, word string, k int) bool {
	if k == len(word) {
		return true
	}
	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || (*visited)[i][j] || k > len(word) {
		return false
	}
	if word[k] != grid[i][j] {
		return false
	}
	(*visited)[i][j] = true
	res := false
	for p := 0; p < 4; p++ {
		res = res || dfs(i+dir[p][0], j+dir[p][1], grid, visited, word, k+1)
	}
	(*visited)[i][j] = false
	return res
}

func main() {
	fmt.Println(search([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE"))
}
