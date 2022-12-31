package main

import "fmt"

func find_all_duplicate_numbers_in_an_array(nums []int) []int {
	i := 0
	res := []int{}
	for i < len(nums) {
		ind := nums[i] - 1
		if nums[i] != nums[ind] {
			nums[i], nums[ind] = nums[ind], nums[i]
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			res = append(res, nums[nums[i]-1])
		}
	}
	return res
}

func main() {
	fmt.Println(find_all_duplicate_numbers_in_an_array([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
