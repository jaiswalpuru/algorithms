package main

import (
	"fmt"
	"math"
)

//-----------------------brute force-----------------------
func arithmetic_slices_two(arr []int) int {
	res := 0
	recur(arr, &res, 0, []int{})
	return res
}
func recur(arr []int, res *int, ind int, set []int) {
	if ind == len(arr) {
		if len(set) <= 2 {
			return
		}
		f := true
		for j := len(set) - 1; j >= 2; j-- {
			if set[j]-set[j-1] != set[j-1]-set[j-2] {
				f = false
				break
			}
		}
		if f {
			*res++
		}
		return
	}
	recur(arr, res, ind+1, append(set, arr[ind]))
	recur(arr, res, ind+1, set)
}

//-----------------------brute force-----------------------

//-----------------------efficient way-----------------------
func arithmetic_slices_two_eff(arr []int) int {
	n := len(arr)
	mp := make([]map[int]int, n)
	res := 0
	for i := 0; i < n; i++ {
		mp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := arr[i] - arr[j]
			if diff < math.MinInt64 || diff > math.MaxInt64 {
				continue
			}
			sum := mp[j][diff]
			origin := mp[i][diff]
			mp[i][diff] += origin + sum + 1
			res += sum
		}
	}
	return res
}

//-----------------------efficient way-----------------------

func main() {
	fmt.Println(arithmetic_slices_two_eff([]int{2, 4, 6, 8, 10}))
}
