package main

import "fmt"

func check_if_array_is_consecutive(arr []int) bool {
	n := len(arr)
	min_val := arr[0]
	mp := make(map[int]bool)
	for i := 0; i < n; i++ {
		min_val = min(min_val, arr[i])
		mp[arr[i]] = true
	}

	for i := min_val; i <= min_val+n-1; i++ {
		if !mp[i] {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(check_if_array_is_consecutive([]int{1, 3, 4, 2}))
}
