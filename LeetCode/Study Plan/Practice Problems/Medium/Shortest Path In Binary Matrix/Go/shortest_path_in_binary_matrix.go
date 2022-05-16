package main

import "fmt"

func is_safe(i, j, n int, grid [][]int) bool {
	return i >= 0 && j >= 0 && i < n && j < n && grid[i][j] == 0
}

var dir = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func shortest_path_binary_matrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	q := [][]int{{0, 0}}
	grid[0][0] = 1
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		i, j := curr[0], curr[1]
		dis := grid[i][j]
		if i == n-1 && j == n-1 {
			return dis
		}

		for k := 0; k < 8; k++ {
			x, y := i+dir[k][0], j+dir[k][1]
			if is_safe(x, y, n, grid) {
				q = append(q, []int{x, y})
				grid[x][y] = dis + 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(shortest_path_binary_matrix([][]int{{0, 1}, {1, 0}}))
}
