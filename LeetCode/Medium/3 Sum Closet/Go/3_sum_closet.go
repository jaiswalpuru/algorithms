package main

import (
	"fmt"
	"math"
	"sort"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func three_sum_closet(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	min_val := math.MaxInt32
	for i := 0; i < n; i++ {
		j, k := i+1, n-1
		for j < k {
			sum := arr[i] + arr[j] + arr[k]
			if abs(sum-target) <= abs(min_val) {
				min_val = target - sum
			}
			if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return target - min_val
}

func main() {
	fmt.Println(three_sum_closet([]int{-1, 2, 1, -4}, 1))
}
