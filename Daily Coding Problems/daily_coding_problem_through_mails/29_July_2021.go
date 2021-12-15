/*
Given an integer list where each number represents the number of hops you can make,
determine whether you can reach to the last index starting at index 0.

For example, [2, 0, 1, 0] returns True while [1, 1, 0, 1] returns False.
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func can_reach(arr []int) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		arr[i-1] = arr[i-1] - 1
		if arr[i-1] < 0 {
			return false
		}
		arr[i] = max(arr[i], arr[i-1])
	}
	if arr[n-1] >= 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(can_reach([]int{2, 3, 1, 1, 4}))
	fmt.Println(can_reach([]int{3, 2, 1, 0, 4}))
	fmt.Println(can_reach([]int{2, 5, 0, 0}))
	fmt.Println(can_reach([]int{0, 1}))
	fmt.Println(can_reach([]int{2, 0, 0}))
}
