package main

import "fmt"

func minimum_insertions_to_make_string_palindrome(s string) int {
	memo := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		memo[i] = make([]int, len(s))
	}
	return len(s) - recur(s, 0, len(s)-1, &memo)
}

func recur(s string, l, r int, memo *[][]int) int {
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
		(*memo)[l][r] = 2 + recur(s, l+1, r-1, memo)
	} else {
		(*memo)[l][r] = max(recur(s, l+1, r, memo), recur(s, l, r-1, memo))
	}
	return (*memo)[l][r]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimum_insertions_to_make_string_palindrome("zzazz"))
}
