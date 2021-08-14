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

import "fmt"

func intersection_two_arrays(arr_one, arr_two []int) []int {
	n, m := len(arr_one), len(arr_two)
	visited := make(map[int]bool)
	temp := make(map[int]bool)
	for i := 0; i < n; i++ {
		visited[arr_one[i]] = true
	}

	res := []int{}
	for i := 0; i < m; i++ {
		if _, ok := visited[arr_two[i]]; ok {
			if _, ok := temp[arr_two[i]]; !ok {
				res = append(res, arr_two[i])
				temp[arr_two[i]] = true
			}
		}
	}

	return res
}

func main() {
	fmt.Println(intersection_two_arrays([]int{1, 2, 2, 2}, []int{2, 2}))
}
