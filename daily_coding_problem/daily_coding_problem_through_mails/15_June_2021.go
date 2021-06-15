/*
	There is an N by M matrix of zeroes. Given N and M, write a function to count the number of ways of starting
	at the top-left corner and getting to the bottom-right corner. You can only move right or down.

	For example, given a 2 by 2 matrix, you should return 2, since there are two ways to get to the bottom-right:

	Right, then down
	Down, then right
	Given a 5 by 5 matrix, there are 70 ways to get to the bottom-right.
*/

package main

import "fmt"

//The complexity for this is exponential
func recursive_ways(m, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return recursive_ways(m-1, n) + recursive_ways(m, n-1)
}

//complexity O(n^2)
func num_path_dp(m, n int) int {
	count_ways := make([][]int, m)

	for i := 0; i < m; i++ {
		count_ways[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		count_ways[i][0] = 1
	}

	for j := 0; j < n; j++ {
		count_ways[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			count_ways[i][j] = count_ways[i][j-1] + count_ways[i-1][j]
		}
	}
	return count_ways[m-1][n-1]
}

func main() {
	fmt.Println(recursive_ways(5, 5))
	fmt.Println(num_path_dp(5, 5))
}
