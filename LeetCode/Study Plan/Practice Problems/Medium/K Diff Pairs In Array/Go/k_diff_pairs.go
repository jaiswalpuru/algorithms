package main

import (
	"fmt"
	"sort"
	"strconv"
)

var itoa = strconv.Itoa

func to_string(a, b int) string {
	v1, v2 := "", ""
	if a < 0 {
		v1 += "-" + itoa(a)
	} else {
		v1 += itoa(a)
	}
	if b < 0 {
		v2 += "-" + itoa(b)
	}
	return v1 + v2
}
func k_diff(arr []int, k int) int {
	n := len(arr)
	sort.Ints(arr)
	fmt.Println(arr)
	mp := make(map[string]struct{})

	res := 0
	for i := 0; i < n; i++ {
		val := arr[i] - k
		l, r := 0, n-1
		mid := -1
		f := false
		for l < r {
			mid = (l + r) / 2
			if arr[mid] == val {
				f = true
				break
			}
			if arr[mid] > val {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if f && mid != i {
			if _, ok := mp[to_string(arr[i], arr[mid])]; !ok {
				res++
				mp[to_string(arr[i], arr[mid])] = struct{}{}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(k_diff([]int{1, 3, 1, 4, 5}, 0))
}
