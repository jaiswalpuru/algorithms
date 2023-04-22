package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longest_substr_with_twodistinct(s string) int {
	n := len(s)
	l, r := 0, 0

	hash_map := make(map[byte]int)

	max_len := 1
	for r < n {
		hash_map[s[r]] = r
		r++
		min_val := r
		if len(hash_map) == 3 {
			for _, v := range hash_map {
				min_val = min(min_val, v)
			}
			delete(hash_map, s[min_val])
			l = min_val + 1
		}
		max_len = max(max_len, r-l)
	}
	return max_len
}

func main() {
	fmt.Println(longest_substr_with_twodistinct("eceab"))
}
