package main

import "fmt"

//---------------brute force----------------------
func longest_common_subsequence(s1, s2 string) int {
	n, m := len(s1)-1, len(s2)-1
	return recur(s1, s2, n, m)
}

func recur(s1, s2 string, n, m int) int {
	if n < 0 || m < 0 {
		return 0
	}
	if s1[n] == s2[m] {
		return 1 + recur(s1, s2, n-1, m-1)
	}
	return max(recur(s1, s2, n, m-1), recur(s1, s2, n-1, m))
}

//---------------brute force----------------------

//-----------------memoization---------------------
func longest_common_subsequence_eff(s1, s2 string) int {
	n, m := len(s1), len(s2)
	memo := make([][]int, n+1)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, m)
		for j := 0; j < m; j++ {
			memo[i][j] = -1
		}
	}
	return _memo(s1, s2, n-1, m-1, &memo)
}

func _memo(s1, s2 string, n, m int, memo *[][]int) int {
	if n < 0 || m < 0 {
		return 0
	}
	if (*memo)[n][m] != -1 {
		return (*memo)[n][m]
	}
	take, dnt_take := 0, 0
	if s1[n] == s2[m] {
		take = 1 + _memo(s1, s2, n-1, m-1, memo)
	}
	dnt_take = max(_memo(s1, s2, n-1, m, memo), _memo(s1, s2, n, m-1, memo))
	(*memo)[n][m] = max(take, dnt_take)
	return (*memo)[n][m]
}

//-----------------memoization---------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longest_common_subsequence_eff("abcde", "ace"))
}
