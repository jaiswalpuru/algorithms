package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max_consecutive_one(arr []int) int {
	n := len(arr)
	ans := 0
	res := 0
	for i := 0; i < n; i++ {
		if arr[i] == 1 {
			res++
		} else {
			ans = max(ans, res)
			res = 0
		}
	}
	ans = max(ans, res)
	return ans
}

func main() {
	fmt.Println(max_consecutive_one([]int{1, 1, 0, 1, 1, 1}))
}
