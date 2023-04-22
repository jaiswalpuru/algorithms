package main

import "fmt"

func sort_colors(nums []int) {
	hash_map := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		hash_map[nums[i]]++
	}
	k := 0
	for i := range []int{0, 1, 2} {
		for hash_map[i] > 0 {
			nums[k] = i
			hash_map[i]--
			k++
		}
	}
	fmt.Println(nums)
}

func sort_colors_one_pass(nums []int) {
	curr := 0
	p0, p2 := 0, len(nums)-1
	for curr <= p2 {
		if nums[curr] == 0 {
			nums[curr], nums[p0] = nums[p0], nums[curr]
			curr++
			p0++
		} else if nums[curr] == 2 {
			nums[curr], nums[p2] = nums[p2], nums[curr]
			p2--
		} else {
			curr++
		}
	}
	fmt.Println(nums)
}

func main() {
	sort_colors([]int{2, 0, 2, 1, 1, 0})
	sort_colors_one_pass([]int{2, 0, 2, 1, 1, 0})
}
