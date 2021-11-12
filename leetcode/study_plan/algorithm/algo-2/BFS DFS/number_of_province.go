/*
There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b,
and city b is connected directly with city c, then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected,
and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.
*/

package main

import "fmt"

func dfs(grid [][]int, visited []bool, i, n int) {
	if visited[i] {
		return
	}
	visited[i] = true
	for j := 0; j < n; j++ {
		if grid[i][j] == 1 && i != j {
			dfs(grid, visited, j, n)
		}
	}
}

func number_of_province(grid [][]int) int {
	n := len(grid)
	cnt := 0
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if !visited[i] {
			cnt++
			dfs(grid, visited, i, n)
		}
	}
	return cnt
}

func main() {
	fmt.Println(number_of_province([][]int{
		{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
	}))
	fmt.Println(number_of_province([][]int{
		{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
	}))
}
