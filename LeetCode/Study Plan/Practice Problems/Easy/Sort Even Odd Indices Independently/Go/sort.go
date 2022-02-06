package main

import (
	"fmt"
	"sort"
)

func sort_even_odd(nums []int) []int {
	even := []int{}
	odd := []int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	sort.Ints(even)
	sort.Ints(odd)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			nums[i] = even[0]
			even = even[1:]
		} else {
			nums[i] = odd[len(odd)-1]
			odd = odd[:len(odd)-1]
		}
	}
	return nums
}

func main() {
	fmt.Println(sort_even_odd([]int{4, 1, 2, 3}))
}
