package main

import "fmt"

func circular_array_loop(nums []int) bool {
	color := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if color[i] == 0 && dfs(nums, &color, i) {
			return true
		}
	}
	return false
}

func dfs(nums []int, color *[]int, start int) bool {
	if (*color)[start] == 2 {
		return false
	}
	(*color)[start] = 1
	next_node := ((nums[start]+start)%len(nums) + len(nums)) % len(nums)
	if next_node == start || nums[start]*nums[next_node] < 0 {
		(*color)[start] = 2
		return false
	}
	if (*color)[next_node] == 1 {
		(*color)[start] = 2
		return true
	}
	if dfs(nums, color, next_node) {
		return true
	}
	(*color)[start] = 2
	return false
}

func main() {
	fmt.Println(circular_array_loop([]int{-2, 1, 1, 2, 2}))
}
