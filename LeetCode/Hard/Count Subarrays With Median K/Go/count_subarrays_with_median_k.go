package main

import "fmt"

func count_subarrays_with_median_k(nums []int, k int) int {
	ind := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > k {
			nums[i] = 1
		} else if nums[i] < k {
			nums[i] = -1
		} else {
			nums[i] = 0
			ind = i
		}
	}
	sum := 0
	mp := make(map[int]int)
	for i := ind; i >= 0; i-- {
		sum += nums[i]
		mp[sum]++
	}
	sum = 0
	res := 0
	for i := ind; i < len(nums); i++ {
		sum += nums[i]
		if mp[-1*sum] != 0 {
			res += mp[-1*sum]
		}
		if mp[-1*sum+1] != 0 {
			res += mp[-1*sum+1]
		}
	}
	return res
}

func main() {
	fmt.Println(count_subarrays_with_median_k([]int{3, 2, 1, 4, 5}, 1))
}
