/*
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s
such that every character in t (including duplicates) is included in the window. If there is no such
substring, return the empty string "".

The testcases will be generated such that the answer is unique.

A substring is a contiguous sequence of characters within the string.

Example 1:
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:
Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.

Example 3:
Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
*/

package main

import "fmt"

func get_cnt(str string) map[byte]int {
	mp := make(map[byte]int)
	n := len(str)
	for i := 0; i < n; i++ {
		mp[str[i]]++
	}
	return mp
}

func min_window_substring(s, t string) string {

	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	mp_t := get_cnt(t)
	mp_s := make(map[byte]int)

	l, r, n := 0, 0, len(s)
	formed, reqd := 0, len(mp_t)
	ans := []int{-1, 0, 0} //{length of window, left, right}

	for r < n {
		mp_s[s[r]]++

		if mp_s[s[r]] == mp_t[s[r]] {
			formed++
		}

		for l <= r && formed == reqd {
			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}
			mp_s[s[l]]--
			if mp_s[s[l]] < mp_t[s[l]] {
				formed--
			}
			l++
		}

		r++
	}
	if ans[0] == -1 {
		return ""
	}
	return s[ans[1] : ans[2]+1]
}

func main() {
	fmt.Println(min_window_substring("ADOBECODEBANC", "ABC"))
	fmt.Println(min_window_substring("a", "a"))
	fmt.Println(min_window_substring("a", "aa"))
}
