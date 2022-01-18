package main

import (
	"fmt"
	"sort"
)

func find(arr []int) []int {
	n := len(arr)
	if n%2 != 0 {
		return []int{}
	}

	res := []int{}

	original_val := make(map[int]int)

	for i := 0; i < n; i++ {
		original_val[arr[i]]++
	}
	fmt.Println(original_val)

	sort.Ints(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			if original_val[arr[i]] >= 2 {
				original_val[arr[i]] -= 2
				res = append(res, arr[i])
				if original_val[arr[i]] == 0 {
					delete(original_val, arr[i])
				}
			} else {
				return []int{}
			}
		} else {
			v1, k1 := original_val[arr[i]]
			v2, k2 := original_val[arr[i]*2]
			if k1 && k2 {
				res = append(res, arr[i])
				original_val[arr[i]] = v1 - 1
				original_val[arr[i]*2] = v2 - 1
				if v1-1 == 0 {
					delete(original_val, arr[i])
				}
				if v2-1 == 0 {
					delete(original_val, arr[i]*2)
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(find([]int{6, 3, 0, 1}))
}
