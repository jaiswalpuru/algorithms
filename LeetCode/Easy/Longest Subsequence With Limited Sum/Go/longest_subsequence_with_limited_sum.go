package main

import (
	"fmt"
	"sort"
)

func longest_subsequence_with_limited_sum(nums []int, queries []int) []int {
	sort.Ints(nums)
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}
	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = bsearch(prefix, queries[i])
	}
	return res
}

func bsearch(arr []int, val int) int {
	l, r := 0, len(arr)-1
	res := 0
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] <= val {
			res = max(res, mid+1)
			l = mid + 1
		} else {
			r = mid - 1
		}
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
	fmt.Println(longest_subsequence_with_limited_sum([]int{4, 5, 2, 1}, []int{3, 10, 21}))
}
