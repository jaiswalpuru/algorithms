package main

import (
	"fmt"
	"sort"
)

func rearrangeArray(nums []int) []int {
	sort.Ints(nums)
	size := len(nums)
	k, l := 0, 1
	rearrangeArray := make([]int, size)
	median := float64(0)
	if size%2 == 0 {
		median = float64(nums[size/2]+nums[size/2-1]) / 2
	} else {
		median = float64(nums[size/2])
	}

	for i := 0; i < size; i++ {
		if float64(nums[i]) >= median {
			rearrangeArray[k] = nums[i]
			k += 2
		} else {
			rearrangeArray[l] = nums[i]
			l += 2
		}
	}
	return rearrangeArray
}

func main() {
	fmt.Println(rearrangeArray(([]int{1, 2, 3, 4, 5})))
}
