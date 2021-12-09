/*
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/

package main

import "fmt"

//dp
func can_jump(arr []int) bool {
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i-1] > 0 {
			if arr[i-1] > arr[i] {
				arr[i] = arr[i-1] - 1
			}
		} else {
			return false
		}
	}
	return true
}

func _can_jump_recursive(arr []int, ind int) bool {
	if ind >= len(arr) {
		return true
	}
	if arr[ind-1] > 0 {
		if arr[ind-1] > arr[ind] {
			arr[ind] = arr[ind-1] - 1
		}
		return _can_jump_recursive(arr, ind+1)
	}
	return false
}

func can_jump_recursive(arr []int) bool {
	return _can_jump_recursive(arr, 1)
}

func main() {
	fmt.Println(can_jump_recursive([]int{2, 3, 1, 1, 4}))
	fmt.Println(can_jump_recursive([]int{3, 2, 1, 0, 4}))
	fmt.Println(can_jump_recursive([]int{2, 5, 0, 0}))
	fmt.Println(can_jump_recursive([]int{0, 1}))
	fmt.Println(can_jump_recursive([]int{2, 0, 0}))
}
