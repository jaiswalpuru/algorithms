package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func longest_sub_k_dis(str string, k int) int {
	n := len(str)
	if n*k == 0 {
		return 0
	}

	l, r := 0, 0
	hash_map := make(map[byte]int)
	max_len := 1

	for r < n {
		hash_map[str[r]] = r
		r++
		min_val := r
		if len(hash_map) == k+1 {
			for _, v := range hash_map {
				min_val = min(min_val, v)
			}
			l = min_val + 1
			delete(hash_map, str[min_val])
		}
		max_len = max(max_len, r-l)
	}
	return max_len
}

func main() {
	fmt.Println(longest_sub_k_dis("a", 0))
}
