package main

import "fmt"

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

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

//---------------------------brute force---------------------------------
func maximum_absolute_sum_of_any_subarray(arr []int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		temp := arr[i]
		sum_temp := abs(temp)
		for j := i + 1; j < n; j++ {
			temp += arr[j]
			sum_temp = max(sum_temp, abs(temp))
		}
		res = max(res, sum_temp)
	}
	return res
}

//---------------------------brute force---------------------------------

//---------------------------efficient approach--------------------------
func max_abs_sum_of_any_subarray(arr []int) int {
	n := len(arr)
	max_val := arr[0]
	min_val := arr[0]
	temp := arr[0]
	for i := 1; i < n; i++ {
		temp = max(arr[i], arr[i]+temp)
		max_val = max(max_val, temp)
	}
	temp = arr[0]
	for i := 1; i < n; i++ {
		temp = min(arr[i], temp+arr[i])
		min_val = min(min_val, temp)
	}
	return max(abs(max_val), abs(min_val))
}

//---------------------------efficient approach--------------------------

func main() {
	fmt.Println(max_abs_sum_of_any_subarray([]int{1, -3, 2, 3, -4}))
}
