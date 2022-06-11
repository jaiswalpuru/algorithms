package main

import "fmt"

func minimum_operations_to_reduce_x_to_zero(arr []int, x int) int {
	n := len(arr)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	max_val := -1
	l, curr_sum := 0, 0
	for i := 0; i < n; i++ {
		curr_sum += arr[i]
		for curr_sum > sum-x && l <= i {
			curr_sum -= arr[l]
			l++
		}

		if curr_sum == sum-x {
			max_val = max(max_val, i-l+1)
		}
	}
	if max_val != -1 {
		return n - max_val
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimum_operations_to_reduce_x_to_zero([]int{1, 1, 4, 2, 3}, 5))
}
