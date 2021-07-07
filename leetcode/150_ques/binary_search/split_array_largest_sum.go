/*
Given an array nums which consists of non-negative integers and an integer m,
you can split the array into m non-empty continuous subarrays.

Write an algorithm to minimize the largest sum among these m subarrays.

Example 1:
Input: nums = [7,2,5,10,8], m = 2
Output: 18
Explanation:
There are four ways to split nums into two subarrays.
The best way is to split it into [7,2,5] and [10,8],
where the largest sum among the two subarrays is only 18.

Example 2:
Input: nums = [1,2,3,4,5], m = 2
Output: 9

Example 3:
Input: nums = [1,4,4], m = 3
Output: 4
*/

package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func check(arr []int, mid, n, k int) bool {
	cnt, sum := 0, 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > mid {
			return false
		}
		sum += arr[i]
		if sum > mid {
			cnt++
			sum = arr[i]
		}
	}
	cnt++
	if cnt <= k {
		return true
	}
	fmt.Println(mid, cnt)
	return false
}

func split_array_largest_sum(arr []int, m int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	k, sum, res := 0, 0, -1
	for i := 0; i < n; i++ {
		k = max(k, arr[i])
		sum += arr[i]
	}
	low, high := k, sum
	for low <= high {
		mid := low + (high-low)/2
		if check(arr, mid, n, m) {
			res = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}

func main() {
	arr := []int{7, 2, 5, 10, 8}
	fmt.Println(split_array_largest_sum(arr, 2))
	fmt.Println(split_array_largest_sum([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(split_array_largest_sum([]int{1, 4, 4}, 3))
}
