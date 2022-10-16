package main

import (
	"fmt"
	"math"
)

func larget_positive_integer_that_exists_with_its_negative(arr []int) int {
	mp := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			mp[arr[i]] = true
		}
	}
	res := math.MinInt64
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 && mp[-1*arr[i]] {
			if res < arr[i] {
				res = arr[i]
			}
		}
	}
	if res == math.MinInt64 {
		return -1
	}
	return res
}

func main() {
	fmt.Println(larget_positive_integer_that_exists_with_its_negative([]int{-1, 2, -3, 3}))
}
