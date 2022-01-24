package main

import "fmt"

//----------------------------------------will give TLE with this approach----------------------------------------
func check_is_palindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func longest_palindromic_string(s string) string {
	res := ""
	_longest_palindromic_string(s, &res)
	return res
}

func _longest_palindromic_string(s string, res *string) {
	if s == "" {
		return
	}

	if check_is_palindrome(s) {
		if len(s) > len(*res) {
			*res = s
		}
		return
	}

	_longest_palindromic_string(s[1:], res)
	_longest_palindromic_string(s[:len(s)-1], res)
}

//------------------------------------------------------------------------------------------------------------------------

func longest_palindromic_string_dp(s string) string {
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

func main() {
	fmt.Println(longest_palindromic_string_dp("cbbd"))
}
