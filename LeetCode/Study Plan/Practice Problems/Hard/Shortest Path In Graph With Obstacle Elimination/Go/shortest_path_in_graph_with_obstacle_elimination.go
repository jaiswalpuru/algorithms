package main

import (
	"fmt"
	"strconv"
)

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type S struct{ x, y, k, step int }

func shortest_path_in_graph_with_obstacle_elimination(grid [][]int, k int) int {
	q := []S{{0, 0, k, 0}}
	visited := make(map[string]bool)
	visited["0"+"0"+strconv.Itoa(k)] = true
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.x == len(grid)-1 && curr.y == len(grid[0])-1 {
			return curr.step
		}
		for i := 0; i < 4; i++ {
			x, y := curr.x+dir[i][0], curr.y+dir[i][1]
			if x < 0 || y < 0 || x == len(grid) || y == len(grid[0]) {
				continue
			}
			k_update := curr.k - grid[x][y]
			new_path := strconv.Itoa(x) + " " + strconv.Itoa(y) + " " + strconv.Itoa(k_update)
			if k_update >= 0 && !visited[new_path] {
				visited[new_path] = true
				q = append(q, S{x, y, k_update, curr.step + 1})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(shortest_path_in_graph_with_obstacle_elimination([][]int{
		{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0},
	}, 1))
}
