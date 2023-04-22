package main

import (
	"fmt"
	"math"
)

//--------------------brute force------------------------
func increasing_triplet_subsequence(arr []int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i] < arr[j] && arr[j] < arr[k] {
					return true
				}
			}
		}
	}
	return false
}

//--------------------brute force------------------------

//--------------------efficient approach------------------------
func increasing_triplet_subsequence_eff(arr []int) bool {
	first, second := math.MaxInt64, math.MaxInt64
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] <= first {
			first = arr[i]
		} else if arr[i] <= second {
			second = arr[i]
		} else {
			return true
		}
	}
	return false
}

//--------------------effcient approach------------------------

func main() {
	fmt.Println(increasing_triplet_subsequence_eff([]int{1, 2, 3, 4, 5}))
}
