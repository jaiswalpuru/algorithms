/*
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.
*/

package main

import (
	"fmt"
	"sort"
)

func binary_search(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	low, high := 0, len(arr)
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func intersection_array(a, b []int) []int {
	res := []int{}
	n := len(a)
	sort.Ints(b)
	visited := make(map[int]bool)

	for i := 0; i < n; i++ {
		if binary_search(b, a[i]) != -1 {
			if _, ok := visited[a[i]]; !ok {
				res = append(res, a[i])
			}
			visited[a[i]] = true
		}
	}
	return res
}

func main() {
	fmt.Println(intersection_array([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection_array([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
