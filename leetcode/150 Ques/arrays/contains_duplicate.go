/*
Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
*/

package main

import "fmt"

func contains_duplicate(arr []int) bool {
	mp := make(map[int]bool)
	n := len(arr)

	for i := 0; i < n; i++ {
		if _, ok := mp[arr[i]]; ok {
			return true
		}
		mp[arr[i]] = true
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(contains_duplicate(arr))
	arr = []int{1, 2, 3, 4}
	fmt.Println(contains_duplicate(arr))
	arr = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(contains_duplicate(arr))
}
