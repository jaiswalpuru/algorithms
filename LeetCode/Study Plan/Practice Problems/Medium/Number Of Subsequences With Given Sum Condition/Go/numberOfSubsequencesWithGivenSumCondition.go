package main

import (
	"fmt"
	"sort"
)

var mod = int(1e9 + 7)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	size := len(nums)

	powers2 := make([]int, size)
	powers2[0] = 1
	for i := 1; i < size; i++ {
		powers2[i] = powers2[i-1] * 2 % mod
	}

	l, r := 0, size-1
	res := 0
	for l <= r {
		if nums[l]+nums[r] <= target {
			res = (res + powers2[r-l]) % mod
			l++
		} else {
			r--
		}
	}
	return res
}

func main() {
	fmt.Println(numSubseq([]int{2, 9, 4, 3, 6, 9, 1, 1}, 6))
}
