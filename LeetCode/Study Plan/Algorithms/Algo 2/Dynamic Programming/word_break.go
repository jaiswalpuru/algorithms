/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated
sequence of one or more dictionary words.

Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/

package main

import "fmt"

func word_break(s string, dic []string) bool {
	n, m := len(s), len(dic)

	if n == 0 || m == 0 {
		return false
	}
	mp := make(map[string]struct{})
	for i := 0; i < m; i++ {
		mp[dic[i]] = struct{}{}
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				substr := s[j:i]
				if _, ok := mp[substr]; ok {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[n]
}

func word_break_bfs(s string, dic []string) bool {
	n, m := len(s), len(dic)
	if n == 0 || m == 0 {
		return false
	}

	mp := make(map[string]struct{})
	for i := 0; i < m; i++ {
		mp[dic[i]] = struct{}{}
	}

	visited := make([]bool, n)

	q := []int{0}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		if visited[curr] {
			continue
		}

		for j := curr + 1; j <= n; j++ {
			if _, ok := mp[string(s[curr:j])]; ok {
				q = append(q, j)
				if j == n {
					return true
				}
			}
		}
		visited[curr] = true
	}
	return false
}

func main() {
	fmt.Println(word_break_bfs("leetcode", []string{"leet", "code"}))
	fmt.Println(word_break_bfs("catsandog", []string{"cat", "dog", "sand", "and", "cat"}))
	fmt.Println(word_break("applepenapple", []string{"apple", "pen"}))
	fmt.Println(word_break("aaaaaaa", []string{"aaaa", "aa"}))
}
