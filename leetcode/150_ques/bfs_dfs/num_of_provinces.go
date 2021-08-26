/*
There are n cities. Some of them are connected, while some are not. If city a is connected directly
with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.

A province is a group of directly or indirectly connected cities and no other cities outside of the group.

You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are
directly connected, and isConnected[i][j] = 0 otherwise.

Return the total number of provinces.

Example 1:
Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
Output: 2

Example 2:
Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
Output: 3
*/

package main

import "fmt"

func num_of_province(arr [][]int) int {
	province := 0
	n := len(arr)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		if !visited[i] {
			province++
			dfs(arr, visited, i, n)
		}
	}
	return province
}

func dfs(arr [][]int, visited []bool, i, n int) {
	if visited[i] {
		return
	}
	visited[i] = true

	for j := 0; j < n; j++ {
		if arr[i][j] == 1 && i != j { //checking if two buildings are connected or not and the buildings are not the same
			dfs(arr, visited, j, n)
		}
	}
}

func main() {
	fmt.Println(num_of_province([][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}))
}
