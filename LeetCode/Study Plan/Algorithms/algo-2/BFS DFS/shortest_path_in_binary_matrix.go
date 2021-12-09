/*
Given an n x n binary matrix grid, return the length of the shortest clear path in the matrix.
If there is no clear path, return -1.

A clear path in a binary matrix is a path from the top-left cell (i.e., (0, 0)) to the bottom-right cell
(i.e., (n - 1, n - 1)) such that:

All the visited cells of the path are 0.
All the adjacent cells of the path are 8-directionally connected
(i.e., they are different and they share an edge or a corner).
The length of a clear path is the number of visited cells of this path.
*/

package main

import "fmt"

func shortest_path(arr [][]int) int {
	n := len(arr)
	if arr[0][0] != 0 || arr[n-1][n-1] != 0 {
		return -1
	}
	directions := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	arr[0][0] = 1
	q := [][]int{{0, 0}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		dis := arr[curr[0]][curr[1]]
		if curr[0] == n-1 && curr[1] == n-1 {
			return dis
		}

		//store all the neighbours
		m := len(directions)
		neighbours := make([][]int, 0)
		for i := 0; i < m; i++ {
			row, col := curr[0]+directions[i][0], curr[1]+directions[i][1]
			if row < 0 || col < 0 || row >= n || col >= n || arr[row][col] != 0 {
				continue
			}
			neighbours = append(neighbours, []int{row, col})
		}

		m = len(neighbours)
		for i := 0; i < m; i++ {
			row, col := neighbours[i][0], neighbours[i][1]
			q = append(q, []int{row, col})
			arr[row][col] = dis + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(shortest_path([][]int{
		{0, 1}, {1, 0},
	}))
	fmt.Println(shortest_path([][]int{
		{0, 0, 0}, {1, 1, 0}, {1, 1, 0},
	}))
	fmt.Println(shortest_path([][]int{
		{0, 1, 1, 0, 0, 0}, {0, 1, 0, 1, 1, 0}, {0, 1, 1, 0, 1, 0}, {0, 0, 0, 1, 1, 0}, {1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 0},
	}))
}
