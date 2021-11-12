/*
Given two strings s and p, return an array of all the start indices of p's anagrams in s.
You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
*/

package main

import "fmt"

func find_all_anagrams(s, p string) []int {
	ans := []int{}
	mp := make(map[byte]int)
	visited := make(map[byte]int)
	n, m := len(p), len(s)
	for i := 0; i < n; i++ {
		mp[p[i]]++
	}

	for i := 0; i < m; i++ {
		visited[s[i]]++
		if i >= n {
			if _, ok := visited[s[i-n]]; ok {
				if visited[s[i-n]] == 1 {
					delete(visited, s[i-n])
				} else {
					visited[s[i-n]]--
				}
			}
		}
		if len(mp) == len(visited) {
			flag := false
			for k := range visited {
				if mp[k] != visited[k] {
					flag = true
				}
			}
			if !flag {
				ans = append(ans, i-n+1)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(find_all_anagrams("cbaebabacd", "abc"))
	fmt.Println(find_all_anagrams("abab", "ab"))
	fmt.Println(find_all_anagrams("baa", "aa"))
	fmt.Println(find_all_anagrams("abaacbabc", "abc"))
}
