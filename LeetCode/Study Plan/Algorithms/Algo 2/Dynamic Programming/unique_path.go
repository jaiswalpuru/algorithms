/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right
corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?
*/

package main

import "fmt"

func unique_path(m, n int) int {
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		arr[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		arr[i][0] = 1
	}

	for i := 0; i < n; i++ {
		arr[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arr[i][j] = arr[i][j-1] + arr[i-1][j]
		}
	}

	return arr[m-1][n-1]
}

func main() {
	fmt.Println(unique_path(3, 2))
	fmt.Println(unique_path(7, 3))
	fmt.Println(unique_path(3, 3))
}
