/*
Given a string s, return the longest palindromic substring in s.

Example 1:
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:
Input: s = "cbbd"
Output: "bb"

Example 3:
Input: s = "a"
Output: "a"

Example 4:
Input: s = "ac"
Output: "a"
*/

package main

import "fmt"

func reverse(str string) string {
	reverse := []rune(str)
	n := len(reverse)
	for i := 0; i < n/2; i++ {
		reverse[i], reverse[n-i-1] = reverse[n-1-i], reverse[i]
	}
	return string(reverse)
}

func check_palin(temp string) bool {
	if temp == reverse(temp) {
		return true
	}
	return false
}

func lps(str string) string {
	n := len(str)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	max_len := 0
	left := -1
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if str[i] == str[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if j-i+1 > max_len {
					max_len = j - i + 1
					left = i
				}
			}
		}
	}
	return str[left : left+max_len]
}

func main() {
	fmt.Println(lps("cbbd"))
	fmt.Println(lps("babad"))
}
