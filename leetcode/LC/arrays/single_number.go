/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

Example 1:

Input: nums = [2,2,1]
Output: 1
Example 2:

Input: nums = [4,1,2,1,2]
Output: 4
Example 3:

Input: nums = [1]
Output: 1
*/

package main

import "fmt"

func single_number(arr []int) int {
	n := len(arr)
	ans := arr[0]

	for i := 1; i < n; i++ {
		ans ^= arr[i]
	}
	return ans
}

func main() {
	arr := []int{2, 2, 1}
	fmt.Println(single_number(arr))
	arr = []int{4, 1, 2, 1, 2}
	fmt.Println(single_number(arr))
}
