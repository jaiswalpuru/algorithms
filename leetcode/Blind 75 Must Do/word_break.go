/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into
a space-separated sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.

Example 1:
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
*/

package main

import "fmt"

func word_break(str string, dic []string) bool {
	n, m := len(dic), len(str)
	if n == 0 || m == 0 {
		return false
	}

	dp := make([]bool, m+1)
	dp[0] = true
	mp := make(map[string]struct{})
	for i := 0; i < n; i++ {
		mp[dic[i]] = struct{}{}
	}

	for i := 1; i <= m; i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				substr := str[j:i]
				if _, ok := mp[substr]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(dp)-1]
}

func main() {
	fmt.Println(word_break("leetcode", []string{"leet", "code"}))
}
