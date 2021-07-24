/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach
the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

Example 1:
Input: m = 3, n = 7
Output: 28

Example 2:
Input: m = 3, n = 2
Output: 3
Explanation:
From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down

Example 3:
Input: m = 7, n = 3
Output: 28

Example 4:
Input: m = 3, n = 3
Output: 6
*/

package main

import "fmt"

func unique_path(m, n int) int {

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		res[i][0] = 1
	}

	for i := 0; i < n; i++ {
		res[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i][j-1] + res[i-1][j]
		}
	}
	return res[m-1][n-1]

}

func main() {
	fmt.Println(unique_path(3, 7))
	fmt.Println(unique_path(3, 2))
	fmt.Println(unique_path(7, 3))
	fmt.Println(unique_path(3, 3))
}
