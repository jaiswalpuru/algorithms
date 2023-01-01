package main

import "fmt"

func squarring_the_sorted_array(nums []int) []int {
	l, r := 0, len(nums)-1
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		val := 0
		if abs(nums[l]) < abs(nums[r]) {
			val = nums[r]
			r--
		} else {
			val = nums[l]
			l++
		}
		res[i] = val * val
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	fmt.Println(squarring_the_sorted_array([]int{-4, -1, 0, 3, 10}))
}
