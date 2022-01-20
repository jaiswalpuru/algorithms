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

func fruits(arr []int) int {
	n := len(arr)
	cnt := make(map[int]int, 3)
	res := -1
	for i, j := 0, 0; i < n; i++ {
		cnt[arr[i]]++
		for len(cnt) > 2 {
			cnt[arr[j]]--
			if cnt[arr[j]] == 0 {
				delete(cnt, arr[j])
			}
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}

func main() {
	fmt.Println(fruits([]int{0, 1, 6, 6, 4, 4, 6}))
}
