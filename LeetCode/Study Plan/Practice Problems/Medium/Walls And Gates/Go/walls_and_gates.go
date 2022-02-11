package main

import (
	"fmt"
)

type Pair struct{ x, y, dis int }

func walls_and_gate(rooms [][]int) [][]int {
	queue := []Pair{}
	n, m := len(rooms), len(rooms[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, Pair{x: i, y: j, dis: 0})
			}
		}
	}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for len(queue) > 0 {
		i, j, dis := queue[0].x, queue[0].y, queue[0].dis
		visited[i][j] = true
		queue = queue[1:]
		if (i+1 < n && j < m) && rooms[i+1][j] != -1 {
			if rooms[i+1][j] == 0 && !visited[i+1][j] {
				queue = append(queue, Pair{i + 1, j, 0})
			} else {
				if rooms[i+1][j] > dis+1 {
					rooms[i+1][j] = dis + 1
					queue = append(queue, Pair{i + 1, j, dis + 1})
				}
			}
		}
		if (i-1 >= 0 && j < m) && rooms[i-1][j] != -1 {
			if rooms[i-1][j] == 0 && !visited[i-1][j] {
				queue = append(queue, Pair{i - 1, j, 0})
			} else {
				if rooms[i-1][j] > dis+1 {
					rooms[i-1][j] = dis + 1
					queue = append(queue, Pair{i - 1, j, dis + 1})
				}
			}
		}
		if (i < n && j+1 < m) && rooms[i][j+1] != -1 {
			if rooms[i][j+1] == 0 && !visited[i][j+1] {
				queue = append(queue, Pair{i, j + 1, 0})
			} else {
				if rooms[i][j+1] > dis+1 {
					rooms[i][j+1] = dis + 1
					queue = append(queue, Pair{i, j + 1, dis + 1})
				}
			}
		}
		if (i < n && j-1 >= 0) && rooms[i][j-1] != -1 {
			if rooms[i][j-1] == 0 && !visited[i][j-1] {
				queue = append(queue, Pair{i, j - 1, 0})
			} else {

				if rooms[i][j-1] > dis+1 {
					rooms[i][j-1] = dis + 1
					queue = append(queue, Pair{i, j - 1, dis + 1})
				}
			}
		}
	}

	return rooms
}

func main() {
	fmt.Println(walls_and_gate([][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}))
}
