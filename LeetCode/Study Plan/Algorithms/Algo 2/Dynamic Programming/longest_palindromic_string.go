/*
Given a string s, return the longest palindromic substring in s.
*/

package main

import "fmt"

func expand_from_middle(s string, left, right int) int {
	if s == "" || left > right {
		return 0
	}
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1 //-1 because for left >=0 so in the end it will have -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lps(s string) string {
	if len(s) == 0 {
		return ""
	}
	start, end := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		l1 := expand_from_middle(s, i, i)
		l2 := expand_from_middle(s, i, i+1)
		l := max(l1, l2)
		if l > end-start {
			start = i - (l-1)/2
			end = i + (l / 2)
		}
	}

	return s[start : end+1]
}

func lps_dp(s string) string {
	n := len(s)
	lps := ""
	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		arr[i][i] = 1
		lps = string(s[i])
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			arr[i][i+1] = 1
			lps = s[i : i+2]
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			if arr[i+1][j-1] == 1 && s[i] == s[j] {
				arr[i][j] = 1
				if len(lps) < (j - i + 1) {
					lps = s[i : j+1]
				}
			}
		}
	}
	return lps
}

func main() {
	fmt.Println(lps_dp("babad"))
	fmt.Println(lps_dp("cbbd"))
	fmt.Println(lps_dp("aaaaa"))
	fmt.Println(lps_dp("ac"))
}
