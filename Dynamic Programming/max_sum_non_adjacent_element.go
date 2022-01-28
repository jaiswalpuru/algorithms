/*
You are given an array/list of N integers. You are supposed to return the maximum sum of the subsequence with the
constraint that no two elements are adjacent in the given array/list.
*/

package main

import "fmt"

//------------------recursion----------------------
func max_sum(arr []int) int {
	return _max_sum(arr, len(arr)-1)
}

func _max_sum(arr []int, start int) int {
	if start < 0 {
		return 0
	}
	if start == 0 {
		return arr[start]
	}

	pick := arr[start] + _max_sum(arr, start-2)
	not_pick := _max_sum(arr, start-1)
	return max(pick, not_pick)
}

//------------------------------------------------

//for memo and dp solutions see Leetcode/Study Plans/Practice Problems/Medium

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max_sum([]int{1, 2, 3, 1}))
}
