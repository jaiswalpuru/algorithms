package main

import (
	"fmt"
	"strconv"
)

type Pair struct{ x, y, step int }

func to_string(x, y int) string { return strconv.Itoa(x) + "->" + strconv.Itoa(y) }

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func nearest_exit_from_entrance_in_maze(maze [][]byte, ent []int) int {
	q := []Pair{{ent[0], ent[1], 0}}
	visited := make(map[string]bool)
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		if curr.x < 0 || curr.y < 0 || curr.x >= len(maze) || curr.y >= len(maze[0]) {
			continue
		}
		if maze[curr.x][curr.y] == '+' || visited[to_string(curr.x, curr.y)] {
			continue
		}
		if (curr.x == 0 || curr.x == len(maze)-1 || curr.y == 0 || curr.y == len(maze[0])-1) && !(curr.x == ent[0] && curr.y == ent[1]) {
			return curr.step
		}
		visited[to_string(curr.x, curr.y)] = true
		for i := 0; i < 4; i++ {
			x, y := curr.x+dir[i][0], curr.y+dir[i][1]
			if x == ent[0] && y == ent[1] {
				continue
			}
			q = append(q, Pair{x, y, curr.step + 1})
		}
	}
	return -1
}

func main() {
	fmt.Println(nearest_exit_from_entrance_in_maze([][]byte{
		{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'},
	}, []int{1, 2}))
}
