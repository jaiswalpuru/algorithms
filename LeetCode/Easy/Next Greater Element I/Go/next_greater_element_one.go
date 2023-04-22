package main

import "fmt"

func next_greater_element_one(nums1, nums2 []int) []int {
	ind := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		ind[nums2[i]] = i
	}
	n := len(nums2)
	cnt := make([]int, n)
	cnt[n-1] = -1
	stack := []int{}
	for k := n - 2; k >= 0; k-- {
		if nums2[k] < nums2[k+1] {
			stack = append(stack, nums2[k+1])
			cnt[k] = nums2[k+1]
		} else {
			if len(stack) == 0 {
				cnt[k] = -1
			} else {
				for len(stack) > 0 && stack[len(stack)-1] < nums2[k] {
					stack = stack[:len(stack)-1]
				}
				if len(stack) != 0 {
					cnt[k] = stack[len(stack)-1]
				} else {
					cnt[k] = -1
				}
			}
		}
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = cnt[ind[nums1[i]]]
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
	fmt.Println(next_greater_element_one([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
