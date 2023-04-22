package main

import "fmt"

func first_missing_positive(nums []int) int {
	one := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			one++
			break
		}
	}
	if one == 0 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}
	for i := 0; i < len(nums); i++ {
		val := abs(nums[i])
		if val == len(nums) {
			nums[0] = -abs(nums[0])
		} else {
			nums[val] = -abs(nums[val])
		}
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}
	if nums[0] > 0 {
		return len(nums)
	}
	return len(nums) + 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	fmt.Println(first_missing_positive([]int{1, 2, 0}))
}
