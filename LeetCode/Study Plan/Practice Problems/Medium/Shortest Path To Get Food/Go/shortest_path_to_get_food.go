package main

import "fmt"

func is_safe(i, j, n, m int, visited [][]bool, grid [][]byte) bool {
	if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == 'X' || grid[i][j] == '*' || visited[i][j] {
		return false
	}
	return true
}

var direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func shortest_path_to_get_food(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	res := -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '*' {
				res = bfs(i, j, n, m, &visited, grid)
			}
		}
	}
	return res
}

func bfs(i, j, n, m int, visited *[][]bool, grid [][]byte) int {
	type Pair struct{ x, y, dis int }

	q := []Pair{{i, j, 0}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			x, y := curr.x+direction[i][0], curr.y+direction[i][1]
			if is_safe(x, y, n, m, *visited, grid) {
				if grid[x][y] == '#' {
					return curr.dis + 1
				}
				(*visited)[x][y] = true
				q = append(q, Pair{x, y, curr.dis + 1})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(shortest_path_to_get_food([][]byte{
		{'X', 'X', 'X', 'X', 'X', 'X'}, {'X', '*', 'O', 'O', 'O', 'X'}, {'X', 'O', 'O', '#', 'O', 'X'}, {'X', 'X', 'X', 'X', 'X', 'X'},
	}))
}
