package main

import "fmt"

func steps_to_make_array_non_decreasing(nums []int) int {
	res := 0
	stack := [][]int{{nums[len(nums)-1], 0}}
	cnt := 0
	for i := len(nums) - 2; i >= 0; i-- {
		cnt = 0
		for len(stack) > 0 && nums[i] > stack[len(stack)-1][0] {
			cnt = max(cnt+1, stack[len(stack)-1][1])
			stack = stack[:len(stack)-1]
		}
		res = max(res, cnt)
		stack = append(stack, []int{nums[i], cnt})
	}
    return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(steps_to_make_array_non_decreasing([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}))
}
