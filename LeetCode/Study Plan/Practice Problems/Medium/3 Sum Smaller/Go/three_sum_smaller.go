package main

import "fmt"
import "sort"

//--------------------brute force----------------------
func three_sum_smaller(arr []int, target int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i]+arr[j]+arr[k] < target {
					res++
				}
			}
		}
	}
	return res
}

//--------------------brute force----------------------

//--------------------efficient approach---------------
func three_sum_smaller_eff(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		res += two_sum(arr, i+1, target-arr[i])
	}
	return res
}

func two_sum(arr []int, ind, target int) int {
	cnt := 0
	for i := ind; i < len(arr)-1; i++ {
		cnt += b_search(arr, i, target-arr[i]) - i
	}
	return cnt
}

func b_search(arr []int, i, target int) int {
	l, r := i, len(arr)-1
	for l < r {
		mid := (l + r + 1) / 2
		if arr[mid] < target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

//--------------------efficient approach---------------

func main() {
	fmt.Println(three_sum_smaller_eff([]int{-2, 0, 1, 3}, 2))
}
