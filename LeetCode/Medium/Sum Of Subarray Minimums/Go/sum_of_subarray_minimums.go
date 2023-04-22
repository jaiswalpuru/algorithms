package main

import "fmt"

var mod = int(1e9) + 7

//----------------------brute force-------------------
func sum_of_subarray_minimums(arr []int) int {
	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	for i := 0; i < n; i++ {
		min_val := arr[i]
		for j := i + 1; j < n; j++ {
			min_val = min(min_val, arr[j])
			sum = (sum + min_val) % mod
		}
	}
	return sum
}

//----------------------brute force-------------------

//----------------------monotonic stack + contribution of each element (LC refered)-------------------
func sum_of_subarray_minimums_eff(arr []int) int {
	n := len(arr)
	st := []int{}
	sum := 0
	for i := 0; i <= n; i++ {
		for len(st) > 0 && (i == n || arr[st[len(st)-1]] >= arr[i]) {
			mid := st[len(st)-1]
			st = st[:len(st)-1]
			left := 0
			if len(st) == 0 {
				left = -1
			} else {
				left = st[len(st)-1]
			}
			right := i
			cnt := (mid - left) * (right - mid) % mod
			sum += (arr[mid] * cnt) % mod
			sum %= mod
		}
		st = append(st, i)
	}
	return sum
}

//----------------------monotonic stack + contribution of each element (LC refered) -------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(sum_of_subarray_minimums_eff([]int{11, 81, 94, 43, 3}))
}
