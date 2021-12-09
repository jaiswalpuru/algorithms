/*
You are given an array of binary strings strs and two integers m and n.

Return the size of the largest subset of strs such that there are at most m 0's and n 1's in the subset.

A set x is a subset of a set y if all elements of x are also elements of y.

Example 1:
Input: strs = ["10","0001","111001","1","0"], m = 5, n = 3
Output: 4
Explanation: The largest subset with at most 5 0's and 3 1's is {"10", "0001", "1", "0"}, so the answer is 4.
Other valid but smaller subsets include {"0001", "1"} and {"10", "1", "0"}.
{"111001"} is an invalid subset because it contains 4 1's, greater than the maximum of 3.

Example 2:
Input: strs = ["10","0","1"], m = 1, n = 1
Output: 2
Explanation: The largest subset is {"0", "1"}, so the answer is 2.
*/

package main

import "fmt"

type Pair struct {
	zero, one int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func get_count(str []string) map[string]Pair {
	data := make(map[string]Pair, 0)
	n := len(str)
	for i := 0; i < n; i++ {
		m := len(str[i])
		var p Pair
		for j := 0; j < m; j++ {
			if str[i][j] == '0' {
				p.zero++
			} else {
				p.one++
			}
		}
		data[str[i]] = p
	}
	return data
}

func ones_zeroes(str []string, m, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	dic := get_count(str)
	for _, s := range str {
		zero, one := dic[s].zero, dic[s].one
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-zero][j-one])
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(ones_zeroes([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(ones_zeroes([]string{"10", "0", "1"}, 1, 1))
	fmt.Println(ones_zeroes([]string{"111", "1000", "1000", "1000"}, 9, 3))
}
