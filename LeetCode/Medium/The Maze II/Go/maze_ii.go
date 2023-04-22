package main

import (
	"fmt"
	"math"
)

//----------------------------using bfs----------------------------
func shortest_distance(maze [][]int, start []int, dst []int) int {
	n, m := len(maze), len(maze[0])

	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, m)
		for j := 0; j < m; j++ {
			distance[i][j] = math.MaxInt64
		}
	}

	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	distance[start[0]][start[1]] = 0
	q := [][]int{start}
	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		for i := 0; i < 4; i++ {
			x := s[0] + dirs[i][0]
			y := s[1] + dirs[i][1]
			cnt := 0
			for x >= 0 && y >= 0 && x < n && y < m && maze[x][y] == 0 {
				x += dirs[i][0]
				y += dirs[i][1]
				cnt++
			}
			if distance[s[0]][s[1]]+cnt < distance[x-dirs[i][0]][y-dirs[i][1]] {
				distance[x-dirs[i][0]][y-dirs[i][1]] = distance[s[0]][s[1]] + cnt
				q = append(q, []int{x - dirs[i][0], y - dirs[i][1]})
			}
		}
	}
	if distance[dst[0]][dst[1]] == math.MaxInt64 {
		return -1
	}
	return distance[dst[0]][dst[1]]
}

//------------------------------------------------------------------------------------

//---------------------------------using dfs------------------------------------
func shortest_distance_dfs(maze [][]int, start []int, dst []int) int {
	n, m := len(maze), len(maze[0])

	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, m)
		for j := 0; j < m; j++ {
			distance[i][j] = math.MaxInt64
		}
	}

	distance[start[0]][start[1]] = 0
	dfs(maze, start, dst, distance)
	if distance[dst[0]][dst[1]] == math.MaxInt64 {
		return -1
	}
	return distance[dst[0]][dst[1]]
}

func dfs(maze [][]int, start []int, dst []int, distance [][]int) {
	dirs := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for i := 0; i < 4; i++ {
		x := start[0] + dirs[i][0]
		y := start[1] + dirs[i][1]
		cnt := 0
		for x >= 0 && y >= 0 && x < len(maze) && y < len(maze[0]) && maze[x][y] == 0 {
			x += dirs[i][0]
			y += dirs[i][1]
			cnt++
		}
		if distance[start[0]][start[1]]+cnt < distance[x-dirs[i][0]][y-dirs[i][1]] {
			distance[x-dirs[i][0]][y-dirs[i][1]] = distance[start[0]][start[1]] + cnt
			dfs(maze, []int{x - dirs[i][0], y - dirs[i][1]}, dst, distance)
		}

	}
}

//--------------------------using dfs-----------------------------

func main() {
	fmt.Println(shortest_distance_dfs([][]int{
		{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0},
	}, []int{0, 4}, []int{4, 4}))
}
