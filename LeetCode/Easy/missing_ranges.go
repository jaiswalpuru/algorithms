/*
You are given an inclusive range [lower, upper] and a sorted unique integer array nums, where all elements are in the
inclusive range.

A number x is considered missing if x is in the range [lower, upper] and x is not in nums.

Return the smallest sorted list of ranges that cover every missing number exactly. That is, no element of nums is
in any of the ranges, and each missing number is in one of the ranges.

Each range [a,b] in the list should be output as:

"a->b" if a != b
"a" if a == b
*/

package main

import (
	"fmt"
	"strconv"
)

func missing_ranges(nums []int, l, r int) []string {
	stack := []string{}
	n := len(nums)

	if n == 0 {
		temp := ""
		if l == r {
			temp = strconv.Itoa(l)
		} else {
			temp = strconv.Itoa(l) + "->" + strconv.Itoa(r)
		}
		return []string{temp}
	}

	if nums[0] != l {
		temp := ""
		if l == nums[0]-1 {
			temp = strconv.Itoa(l)
		} else {
			temp = strconv.Itoa(l) + "->" + strconv.Itoa(nums[0]-1)
		}
		stack = append(stack, temp)
	}
	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] != 1 {
			temp := ""
			if nums[i-1]+1 == nums[i]-1 {
				temp = strconv.Itoa(nums[i-1] + 1)
			} else {
				temp = strconv.Itoa(nums[i-1]+1) + "->" + strconv.Itoa(nums[i]-1)
			}
			stack = append(stack, temp)
		}
	}
	if nums[n-1] < r {
		temp := ""
		if nums[n-1]+1 == r {
			temp = strconv.Itoa(r)
		} else {
			temp = strconv.Itoa(nums[n-1]+1) + "->" + strconv.Itoa(r)
		}

		stack = append(stack, temp)
	}
	return stack
}

func main() {
	fmt.Println(missing_ranges([]int{-1}, -1, 0))
	fmt.Println(missing_ranges([]int{0, 1, 3, 50, 75}, 0, 99))
	fmt.Println(missing_ranges([]int{-1}, -1, -1))
}
