package main

import "fmt"

//------------------------brute force try to generate all possibilities---------------------
func edit_distance(s1, s2 string) int { return _edit_distance(0, 0, s1, s2) }

func _edit_distance(i, j int, s1, s2 string) int {
	if i == len(s1) {
		return len(s2) - j
	}

	if j == len(s2) {
		return len(s1) - i
	}

	if s1[i] == s2[j] {
		return _edit_distance(i+1, j+1, s1, s2)
	} else {
		insert := 1 + _edit_distance(i, j+1, s1, s2)
		remove := 1 + _edit_distance(i+1, j, s1, s2)
		swap := 1 + _edit_distance(i+1, j+1, s1, s2)
		return min(insert, min(remove, swap))
	}
}

//------------------------brute force try to generate all possibilities---------------------

//--------------------------------memoization----------------------------
func edit_distance_memo(s1, s2 string) int {
	n, m := len(s1), len(s2)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}
	return _edit_distance_memo(0, 0, s1, s2, &dp)
}

func _edit_distance_memo(i, j int, s1, s2 string, dp *[][]int) int {
	if i == len(s1) {
		return len(s2) - j
	}
	if j == len(s2) {
		return len(s1) - i
	}

	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}

	if s1[i] == s2[j] {
		(*dp)[i][j] = _edit_distance_memo(i+1, j+1, s1, s2, dp)
	} else {
		insert := 1 + _edit_distance_memo(i, j+1, s1, s2, dp)
		remove := 1 + _edit_distance_memo(i+1, j, s1, s2, dp)
		swap := 1 + _edit_distance_memo(i+1, j+1, s1, s2, dp)
		(*dp)[i][j] = min(insert, min(remove, swap))
	}
	return (*dp)[i][j]
}

//--------------------------------memoization----------------------------

//---------------------------dp--------------------------------
func edit_distance_dp(s1, s2 string) int {
	n, m := len(s1)+1, len(s2)+1
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
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}
	return dp[n-1][m-1]
}

//---------------------------dp--------------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(edit_distance_dp("horse", "ros"))
}
