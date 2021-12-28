/*
Given a string s, return the longest palindromic substring in s.
*/

package main

import "fmt"

//first applying brute force
func reverse(s string) string {
	str := []rune(s)
	n := len(s)
	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}
	return string(str)
}

func bf_longest_palindromic_string(s string) string {
	n := len(s)
	max_len := 0
	str := ""
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if s[i:j] == reverse(s[i:j]) {
				if max_len < (j - i + 1) {
					max_len = j - i + 1
					str = s[i:j]
				}
			}
		}
	}
	return str
}

func dp_longest_palindromic_string(s string) string {
	n := len(s)
	arr := make([][]int, n)
	lps := ""
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
				if len(lps) < j-i+1 {
					lps = s[i : j+1]
				}
			}
		}
	}

	return lps
}

func main() {
	fmt.Println(dp_longest_palindromic_string("babad"))
	fmt.Println(dp_longest_palindromic_string("cbbd"))
}
