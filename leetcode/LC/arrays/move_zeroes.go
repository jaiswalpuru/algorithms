/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of
the non-zero elements.

Note that you must do this in-place without making a copy of the array.

Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
Example 2:

Input: nums = [0]
Output: [0]
*/

package main

import "fmt"

func move_zeroes(arr []int) []int {
	n := len(arr)

	k, cnt := 0, 0
	for i := 0; i < n; i++ {
		if arr[i] != 0 {
			arr[k] = arr[i]
			k++
		} else {
			cnt++
		}
	}

	for i := k; i < n; i++ {
		arr[i] = 0
	}

	return arr
}

func main() {
	arr := []int{0, 1, 0, 3, 12}
	fmt.Println(move_zeroes(arr))
}
