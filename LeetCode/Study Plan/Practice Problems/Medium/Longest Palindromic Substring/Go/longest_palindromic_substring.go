package main

import "fmt"

func longest_palindromic_substring(s string) string {
	res := ""
	_recur(s, &res)
	return res
}

//-----------------brute force way--------------
func _recur(s string, res *string) {
	if s == "" {
		return
	}

	if is_palindrome(s) {
		if len(s) > len(*res) {
			*res = s
		}
		return
	}

	_recur(s[1:], res)
	_recur(s[:len(s)-1], res)
}

//-----------------brute force way--------------

//---------------efficient way-------------------
func lps_memo(s string) string {
	n := len(s)
	dp := make([][]int, n)
	lps := ""

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
		lps = string(s[i])
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 1
			lps = s[i : i+2]
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			if dp[i+1][j-1] == 1 && s[i] == s[j] {
				dp[i][j] = 1
				if len(lps) < j-i+1 {
					lps = s[i : j+1]
				}
			}
		}
	}
	return lps
}

//---------------efficient way-------------------

func is_palindrome(s string) bool {
	n := len(s)
	if n == 1 {
		return true
	}
	if n == 2 {
		return s[0] == s[1]
	}
	i, j := 0, n-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(lps_memo("babad"))
}
