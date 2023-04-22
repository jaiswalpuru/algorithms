package main

import "fmt"

//---------------------------brute force---------------------------
func longest_palindromeic_subsequence(s string) int {
	return recur(0, len(s)-1, s)
}

func recur(l, r int, s string) int {
	if l == r {
		return 1
	}
	if l > r {
		return 0
	}
	if s[l] == s[r] {
		return 2 + recur(l+1, r-1, s)
	} else {
		return max(recur(l+1, r, s), recur(l, r-1, s))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//---------------------------brute force---------------------------

//---------------------------memo---------------------------
func longest_palindromeic_subsequence_memo(s string) int {
	memo := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		memo[i] = make([]int, len(s))
	}
	return recur(0, len(s)-1, s, &memo)
}

func recur(l, r int, s string, memo *[][]int) int {
	if l == r {
		return 1
	}
	if l > r {
		return 0
	}
	if (*memo)[l][r] != 0 {
		return (*memo)[l][r]
	}
	if s[l] == s[r] {
		(*memo)[l][r] = 2 + recur(l+1, r-1, s, memo)
	} else {
		(*memo)[l][r] = max(recur(l+1, r, s, memo), recur(l, r-1, s, memo))
	}
	return (*memo)[l][r]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//---------------------------memo---------------------------

func main() {
	fmt.Println(longest_palindromic_subsequence("bbbab"))
}
