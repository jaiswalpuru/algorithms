/*
Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
The distance between two adjacent cells is 1.

Example 1:
Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]

Example 2:
Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]
*/
package main

import (
	"fmt"
	"math"
)

var INT_MAX = int(math.MaxInt32)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

//brute force approach
func update_matrix_bf(arr [][]int) [][]int {
	n, m := len(arr), len(arr[0])
	distance := make([][]int, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			distance[i][j] = INT_MAX
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				distance[i][j] = 0
			} else {
				for k := 0; k < n; k++ {
					for l := 0; l < m; l++ {
						if arr[k][l] == 0 {
							d := abs(k-i) + abs(l-j)
							distance[i][j] = min(distance[i][j], d)
						}
					}
				}
			}
		}
	}

	return distance
}

type Pair struct {
	i, j int
}

func update_matrix_bfs(arr [][]int) [][]int {
	n, m := len(arr), len(arr[0])

	dis := make([][]int, n)
	q := make([]Pair, 0)

	for i := 0; i < n; i++ {
		dis[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				dis[i][j] = 0
				q = append(q, Pair{i: i, j: j})
			} else {
				dis[i][j] = INT_MAX
			}
		}
	}

	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for i := 0; i < 4; i++ {
			r, c := curr.i+dir[i][0], curr.j+dir[i][1]
			if r >= 0 && c >= 0 && r < n && c < m {
				if dis[r][c] > dis[curr.i][curr.j] {
					dis[r][c] = dis[curr.i][curr.j] + 1
					q = append(q, Pair{i: r, j: c})
				}
			}
		}
	}

	return dis
}

func update_matrix_dp(arr [][]int) [][]int {
	n, m := len(arr), len(arr[0])
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dis[i][j] = INT_MAX
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				dis[i][j] = 0
			} else {
				if i > 0 {
					dis[i][j] = min(dis[i][j], dis[i-1][j]+1)
				}
				if j > 0 {
					dis[i][j] = min(dis[i][j], dis[i][j-1]+1)
				}
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i < n-1 {
				dis[i][j] = min(dis[i][j], dis[i+1][j]+1)
			}
			if j < m-1 {
				dis[i][j] = min(dis[i][j], dis[i][j+1]+1)
			}
		}
	}

	return dis
}

func main() {
	fmt.Println(update_matrix_bf([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))

	fmt.Println(update_matrix_bfs([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
	fmt.Println(update_matrix_dp([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}
