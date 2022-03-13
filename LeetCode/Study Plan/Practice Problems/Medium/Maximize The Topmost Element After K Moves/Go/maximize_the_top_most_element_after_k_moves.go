package main

import "fmt"

func maximize(arr []int, k int) int {
	n := len(arr)
	if n == 1 {
		if k%2 == 1 {
			return -1
		} else {
			return arr[0]
		}
	}

	max_val := 0
	for i := 0; i < n; i++ {
		if k == i || k >= i+2 {
			max_val = max(max_val, arr[i])
		}
	}
	return max_val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximize([]int{5, 2, 2, 4, 0, 6}, 4))
}
