package main

import (
	"fmt"
)

type P struct {
	coord []int
	dis   int
}

var direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func shortest_bridge(arr [][]int) int {
	n, m := len(arr), len(arr[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	i, j := 0, 0
	for i = 0; i < n; i++ {
		f := true
		for j = 0; j < m; j++ {
			if arr[i][j] == 1 {
				f = false
				break
			}
		}
		if !f {
			break
		}
	}

	q := []P{}
	dfs(arr, i, j, n, m, &visited, &q)
	for len(q) > 0 {
		x, y := q[0].coord[0], q[0].coord[1]
		dis := q[0].dis
		q = q[1:]
		for i := 0; i < 4; i++ {
			c_x, c_y := x+direction[i][0], y+direction[i][1]
			if is_safe(c_x, c_y, n, m) && !visited[c_x][c_y] {
				if arr[c_x][c_y] == 1 {
					return dis
				} else {
					q = append(q, P{[]int{c_x, c_y}, dis + 1})
				}
				visited[c_x][c_y] = true
			}
		}
	}

	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func is_safe(i, j, n, m int) bool {
	if i < 0 || j < 0 || i >= n || j >= m {
		return false
	}
	return true
}

func dfs(arr [][]int, i, j, n, m int, visited *[][]bool, q *[]P) {
	(*visited)[i][j] = true
	*q = append(*q, P{[]int{i, j}, 0})

	for k := 0; k < 4; k++ {
		x, y := i+direction[k][0], j+direction[k][1]
		if is_safe(x, y, n, m) && arr[x][y] == 1 && !(*visited)[x][y] {
			dfs(arr, x, y, n, m, visited, q)
		}
	}
}

func main() {
	fmt.Println(shortest_bridge([][]int{
		{0, 1}, {1, 0},
	}))
}
