package main

import "fmt"

func find_all_numbers_disappeared_in_array(nums []int) []int {
	n := len(nums)
	arr := make([]int, n+1)
	res := []int{}
	for i := 0; i < n; i++ {
		arr[nums[i]] = nums[i]
	}
	for i := 1; i <= n; i++ {
		if arr[i] == 0 {
			res = append(res, i)
		}
	}
	return res

}

func main() {
	fmt.Println(find_all_numbers_disappeared_in_array([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
