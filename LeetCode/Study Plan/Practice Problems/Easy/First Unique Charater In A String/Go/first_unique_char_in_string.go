package main

import "fmt"

func first_unique_character(s string) int {
	mp := make(map[byte]int)
	n := len(s)
	for i := 0; i < n; i++ {
		mp[s[i]]++
	}

	for i := 0; i < n; i++ {
		if mp[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(first_unique_character("leetcode"))
}
