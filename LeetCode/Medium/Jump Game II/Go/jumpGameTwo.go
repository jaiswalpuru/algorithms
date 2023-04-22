package main

import "fmt"

func jump(nums []int) int {
	size := len(nums)
	minJumps := 0
	curr, farthest := 0, 0

	for i := 0; i < size-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == curr {
			minJumps += 1
			curr = farthest
		}
	}

	return minJumps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}
