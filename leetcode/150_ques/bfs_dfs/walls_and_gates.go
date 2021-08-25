/*
You are given a m x n 2D grid initialized with these three possible values.
-1 - A wall or an obstacle.
0 - A gate.
INF - Infinity means an empty room. We use the value 2^31 - 1 = 2147483647 to represent INF as you may
assume that the distance to a gate is less than2147483647.
Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it
should be filled with INF.

For example, given the 2D grid:
INF  -1  0  INF
INF INF INF  -1
INF  -1 INF  -1
  0  -1 INF INF
After running your function, the 2D grid should be:
  3  -1   0   1
  2   2   1  -1
  1  -1   2  -1
  0  -1   3   4
*/

package main

import (
	"fmt"
	"math"
)

const INF = int(math.MaxInt32)

func dfs(arr [][]int, i, j, distance int, visited [][]bool, n, m int) {
	if i < 0 || j < 0 || i >= n || j >= m || arr[i][j] == -1 || visited[i][j] || distance > arr[i][j] {
		return
	}
	arr[i][j] = distance
	visited[i][j] = true

	dfs(arr, i-1, j, distance+1, visited, n, m)
	dfs(arr, i+1, j, distance+1, visited, n, m)
	dfs(arr, i, j-1, distance+1, visited, n, m)
	dfs(arr, i, j+1, distance+1, visited, n, m)
	visited[i][j] = false
}

func walls_and_gates(arr [][]int) [][]int {
	n, m := len(arr), len(arr[0])
	if n == 0 {
		return nil
	}

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				dfs(arr, i, j, 0, visited, n, m)
			}
		}
	}

	return arr
}

func main() {
	arr := [][]int{
		{INF, -1, 0, INF},
		{INF, INF, INF, -1},
		{INF, -1, INF, -1},
		{0, -1, INF, INF},
	}
	fmt.Println(walls_and_gates(arr))
}
