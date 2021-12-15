/*
The edit distance between two strings refers to the minimum number of character insertions,
deletions, and substitutions required to change one string to the other. For example, the edit
distance between “kitten” and “sitting” is three: substitute the “k” for “s”, substitute the “e” for “i”,
and append a “g”.

Given two strings, compute the edit distance between them.
*/

package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func computeDistance(word1, word2 string) int {
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
	s1, s2 := "kitten", "sitting"
	fmt.Println(computeDistance(s1, s2))
}
