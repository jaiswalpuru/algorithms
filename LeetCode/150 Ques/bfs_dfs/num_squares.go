/*
Given an integer n, return the least number of perfect square numbers that sum to n.

A perfect square is an integer that is the square of an integer; in other words, it is the
product of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.

Example 1:
Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.

Example 2:
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
*/

package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func num_squares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		min := i
		for j := 1; j*j <= i; j++ {
			if cur := 1 + dp[i-j*j]; min > cur {
				min = cur
			}
		}
		dp[i] = min
	}

	return dp[n]
}

func num_squares_bfs(n int) int {
	queue := make([]int, 0)

	queue = append(queue, n)

	lvl := 1
	for len(queue) > 0 {
		n := len(queue)

		for ; n > 0; n-- {
			p := queue[0]
			queue = queue[1:]
			for i := int(math.Sqrt(float64(p))); i > 0; i-- {
				v := p - i*i
				if v == 0 {
					return lvl
				}
				if v < 0 {
					continue
				}
				queue = append(queue, v)
			}
		}
		lvl++
	}
	return lvl
}

func main() {
	fmt.Println(num_squares(12))
	fmt.Println(num_squares_bfs(12))
	fmt.Println(num_squares(13))
	fmt.Println(num_squares_bfs(13))
}
