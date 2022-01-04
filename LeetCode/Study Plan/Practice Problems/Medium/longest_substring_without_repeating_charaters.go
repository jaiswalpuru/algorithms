/*
Given a string s, find the length of the longest substring without repeating characters.
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longest_substring(str string) int {
	n := len(str)
	i, j := 0, 0
	mp := make(map[string]int)

	res := 0
	for i < n {
		if _, ok := mp[string(str[i])]; ok {
			t := i - j
			res = max(res, t)
			if j <= mp[string(str[i])] {
				j = mp[string(str[i])] + 1
			}
		}
		mp[string(str[i])] = i
		i++
	}
	return max(res, i-j)
}

func main() {
	fmt.Println(longest_substring("bbbbb"))
}
