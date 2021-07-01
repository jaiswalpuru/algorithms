/*
Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must appear as many times as it shows in both arrays and you may
return the result in any order.


Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.
*/

package main

import "fmt"

func intersection(arr_one, arr_two []int) []int {
	mp := make(map[int]bool)
	n, m := len(arr_one), len(arr_two)
	result := []int{}

	for i := 0; i < n; i++ {
		if _, ok := mp[arr_one[i]]; !ok {
			mp[arr_one[i]] = true
		}
	}

	for i := 0; i < m; i++ {
		if _, ok := mp[arr_two[i]]; ok {
			result = append(result, arr_two[i])
		}
	}

	return result
}

func main() {
	arr_one := []int{1, 2, 2, 1}
	arr_two := []int{2, 2}
	fmt.Println(intersection(arr_one, arr_two))
	arr_one = []int{9, 4, 9, 8, 4}
	arr_two = []int{4, 9, 5}
	fmt.Println(intersection(arr_one, arr_two))
}
