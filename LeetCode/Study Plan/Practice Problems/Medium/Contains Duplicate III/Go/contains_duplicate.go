package main

import (
	"fmt"
)

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

//-------------------------brute force--------------------------
func contains_duplicate(arr []int, k, t int) bool {
	n := len(arr)

	for i := 0; i < n; i++ {
		val := arr[i]
		for j := 0; j < n; j++ {
			if i != j {
				if abs(val-arr[j]) <= t && abs(i-j) <= k {
					return true
				}
			}
		}
	}
	return false
}

//-------------------------brute force--------------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-------------------------efficient--------------------------
func contains_duplicate_eff(arr []int, k, t int) bool { //reference LC
	//using idea of bucket sort
	mp := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		key := arr[i] / max(1, t)
		if arr[i] < 0 {
			key--
		}

		if _, ok := mp[key]; ok {
			return true
		}

		if v, ok := mp[key-1]; ok && abs(arr[i]-v) <= t {
			return true
		}

		if v, ok := mp[key+1]; ok && abs(arr[i]-v) <= t {
			return true
		}

		mp[key] = arr[i]
		if i >= k {
			delete(mp, arr[i-k]/max(t, 1))
		}
	}
	return false
}

//-------------------------efficient--------------------------

func main() {
	fmt.Println(contains_duplicate_eff([]int{1, 2, 3, 1}, 3, 0))
}
