/*
Given an array of integers nums and an integer target, return indices of the two numbers such that
they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].


Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]


Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

package main

import "fmt"

func main() {
	fmt.Println(two_sum([]int{2, 7, 11, 15}, 9))
	fmt.Println(two_sum([]int{3, 2, 4}, 6))
}

func two_sum(nums []int, target int) []int {

	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}
	ans := []int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		temp := target - nums[i]
		if _, ok := mp[temp]; ok && mp[temp] != i {
			ans = append(ans, mp[temp])
			ans = append(ans, i)
			break
		}
	}
	return ans
}
