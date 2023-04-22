package main

import "fmt"

func detect_cycle(arr [][]byte) bool {
	n, m := len(arr), len(arr[0])

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !visited[i][j] {
				if dfs(arr, &visited, i, j, n, m, -1, -1, arr[i][j]) {
					return true
				}
			}
		}
	}

	return false
}

var dir = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func is_valid(i, j, n, m int) bool { return !(i < 0 || j < 0 || i >= n || j >= m) }

func dfs(arr [][]byte, visited *[][]bool, i, j, n, m, p_i, p_j int, val byte) bool {
	(*visited)[i][j] = true

	for k := 0; k < 4; k++ {
		x, y := i+dir[k][0], j+dir[k][1]
		if is_valid(x, y, n, m) && arr[x][y] == val && (p_i != x || p_j != y) {
			if (*visited)[x][y] {
				return true
			} else {
				if dfs(arr, visited, x, y, n, m, i, j, val) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(detect_cycle([][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'a', 'a', 'a'},
	}))
}
