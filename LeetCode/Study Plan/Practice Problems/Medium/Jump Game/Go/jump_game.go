package main

import "fmt"

func jump_game(nums []int) bool {
	if nums[0] == 0 {
		return len(nums) == 1
	}
	jump := nums[0]
	for i := 1; i < len(nums); i++ {
		if jump == 0 {
			return false
		}
		jump = max(jump-1, nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(jump_game([]int{2, 3, 1, 1, 4}))
}
