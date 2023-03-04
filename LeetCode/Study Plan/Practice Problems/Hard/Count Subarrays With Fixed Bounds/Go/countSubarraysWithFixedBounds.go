package main

import (
	"fmt"
	"math"
)

//using two pointers approach
func countSubarrays(nums []int, minK int, maxK int) int64 {
	size := len(nums)
	subarray := int64(0)
	minPos := -1
	maxPos := -1
	left := -1

	for i := 0; i < size; i++ {
		if nums[i] < minK || nums[i] > maxK {
			left = i
		}
		if nums[i] == minK {
			minPos = i
		}
		if nums[i] == maxK {
			maxPos = i
		}
		subarray += int64(max(0, min(maxPos, minPos)-left))
	}
	return subarray
}

func countSubarraysBruteForce(nums []int, minK int, maxK int) int64 {
	size := len(nums)

	subarray := int64(0)
	for i := 0; i < size; i++ {
		minVal := math.MaxInt64
		maxVal := math.MinInt64
		for j := i; j < size; j++ {
			minVal = min(minVal, nums[j])
			maxVal = max(maxVal, nums[j])
			if minVal == minK && maxVal == maxK {
				subarray++
			}
		}
	}

	return subarray
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
}
