package main

import "fmt"

func number_of_zero_filled_subarrays(nums []int) int64 {
	num := 0
	res := int64(0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			num++
		} else {
			res += int64(num * (num + 1) / 2)
			num = 0
		}
	}
	if num > 0 {
		res += int64(num * (num + 1) / 2)
	}
	return res
}

func main() {
	fmt.Println(number_of_zero_filled_subarrays([]int{1, 3, 0, 0, 2, 0, 0, 4}))
}
