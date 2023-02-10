package main

import "fmt"

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func maxDistance(grid [][]int) int {
	size := len(grid)
	distance := -1
	q := [][2]int{}
	visited := make([][]int, size)
	for i := 0; i < size; i++ {
		visited[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			visited[i][j] = grid[i][j]
			if grid[i][j] == 1 {
				q = append(q, [2]int{i, j})
			}
		}
	}

	for len(q) > 0 {
		qLen := len(q)
		for k := 0; k < qLen; k++ {
			curr := q[0]
			q = q[1:]
			for i := 0; i < 4; i++ {
				x, y := curr[0]+dir[i][0], curr[1]+dir[i][1]
				if x >= 0 && y >= 0 && x < size && y < size && visited[x][y] == 0 {
					visited[x][y] = 1
					q = append(q, [2]int{x, y})
				}
			}
		}
		distance++
	}

	if distance == 0 {
		return -1
	}
	return distance
}

func main() {
	fmt.Println(maxDistance([][]int{
		{1, 0, 1}, {0, 0, 0}, {1, 0, 1},
	}))
}
