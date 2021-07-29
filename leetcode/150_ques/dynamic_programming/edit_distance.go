/*
Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

You have the following three operations permitted on a word:

Insert a character
Delete a character
Replace a character

Example 1:
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func edit_distance_recursive(word1, word2 string) int {
	n, m := len(word1), len(word2)
	if n == 0 || m == 0 {
		return max(n, m)
	}
	d1 := edit_distance_recursive(word1[1:], word2) + 1
	d2 := edit_distance_recursive(word1, word2[1:]) + 1
	var d3 int
	if word1[0] == word2[0] {
		d3 = edit_distance_recursive(word1[1:], word2[1:])
	} else {
		d3 = edit_distance_recursive(word1[1:], word2[1:]) + 1
	}
	return min(d1, min(d2, d3))
}

func edit_distance_dp(word1, word2 string) int {
	n, m := len(word1)+1, len(word2)+1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		dp[i][0] = i
	}

	for j := 0; j < m; j++ {
		dp[0][j] = j
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}

	return dp[n-1][m-1]
}

func main() {
	fmt.Println(edit_distance_dp("horse", "ros"))
	fmt.Println(edit_distance_dp("kitten", "sitting"))
}
