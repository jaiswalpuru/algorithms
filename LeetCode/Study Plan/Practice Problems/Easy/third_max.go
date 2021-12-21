/*
Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist,
return the maximum number.
*/

package main

import (
	"fmt"
	"math"
)

func third_max(arr []int) int {
	f, s, t := int(math.MinInt64), int(math.MinInt64), int(math.MinInt64)
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > f {
			f = arr[i]
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] > s && arr[i] != f {
			s = arr[i]
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] > t && arr[i] != f && arr[i] != s {
			t = arr[i]
		}
	}

	if t == int(math.MinInt64) {
		return f
	}
	return t

}

func main() {
	fmt.Println(third_max([]int{3, 2, 1}))
}
