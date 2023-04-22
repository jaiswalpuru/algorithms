package main

import "fmt"

//Note the ball is rolling
func has_path(maze [][]int, start []int, dst []int) bool {
	n, m := len(maze), len(maze[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	return _has_path(maze, start[0], start[1], dst[0], dst[1], n, m, &visited)
}

func _has_path(maze [][]int, i, j, x, y, n, m int, visited *[][]bool) bool {
	if (*visited)[i][j] {
		return false
	}
	if i == x && j == y {
		return true
	}
	(*visited)[i][j] = true
	r, l, u, d := j+1, j-1, i-1, i+1
	for r < m && maze[i][r] == 0 {
		r++
	}
	if _has_path(maze, i, r-1, x, y, n, m, visited) {
		return true
	}

	for l >= 0 && maze[i][l] == 0 {
		l--
	}
	if _has_path(maze, i, l+1, x, y, n, m, visited) {
		return true
	}

	for u >= 0 && maze[u][j] == 0 {
		u--
	}
	if _has_path(maze, u+1, j, x, y, n, m, visited) {
		return true
	}

	for d < n && maze[d][j] == 0 {
		d++
	}
	if _has_path(maze, d-1, j, x, y, n, m, visited) {
		return true
	}
	return false
}

func main() {
	fmt.Println(has_path([][]int{
		{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0},
	}, []int{0, 4}, []int{4, 4}))
}
