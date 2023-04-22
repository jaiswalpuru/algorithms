package main

import (
	"fmt"
	"math"
)

func contains_duplicate(arr []int, k int) bool {
	n := len(arr)
	if n == 1 {
		return false
	}

	hash_map := make(map[int]int)
	for i := 0; i < n; i++ {
		if v, ok := hash_map[arr[i]]; ok {
			if int(math.Abs(float64(i-v))) <= k {
				return true
			} else {
				hash_map[arr[i]] = i
			}
		} else {
			hash_map[arr[i]] = i
		}
	}
	return false
}

func main() {
	fmt.Println(contains_duplicate([]int{1, 2, 3, 1}, 3))
}
