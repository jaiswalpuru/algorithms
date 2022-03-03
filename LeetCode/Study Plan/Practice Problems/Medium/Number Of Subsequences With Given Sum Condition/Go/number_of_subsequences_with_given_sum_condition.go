package main

import (
	"fmt"
	"math"
	"sort"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-----------------------------dont use this method brute force -------------------------(TLE)
func number_of_subsequence_with_given_sum_k(arr []int, k int) int {
	res := [][]int{}
	cnt := 0
	_backtrack(arr, k, 0, 0, &res, []int{}, &cnt)
	return cnt
}

func _backtrack(arr []int, k, ind int, cur_sum int, res *[][]int, set []int, cnt *int) {

	if ind == len(arr) {
		min_val, max_val := math.MaxInt64, math.MinInt64
		for j := 0; j < len(set); j++ {
			min_val = min(min_val, set[j])
			max_val = max(max_val, set[j])
		}
		if min_val+max_val <= k && len(set) != 0 {
			(*cnt)++
		}
		return
	}

	temp := append(set, arr[ind])
	_backtrack(arr, k, ind+1, cur_sum+arr[ind], res, temp, cnt)
	temp = temp[:len(temp)-1]
	_backtrack(arr, k, ind+1, cur_sum, res, temp, cnt)
}

//---------------------------------------------------------------------------------

//--------------------Efficient soln--------------------------------------------
func number_of_subsequence_with_given_sum_k_effcient(arr []int, k int) int {
	var mod int = 1e9 + 7
	n := len(arr)
	sort.Ints(arr)

	power := make([]int, n)
	power[0] = 1
	for i := 1; i < n; i++ {
		power[i] = (power[i-1] << 1) % mod
	}

	res := 0
	j := n - 1
	for i := 0; i < n; i++ {
		for arr[i]+arr[j] > k && i < j {
			j--
		}
		if arr[i]+arr[j] <= k {
			res = (res%mod + power[j-i]%mod) % mod
		}
	}
	return res
}

func main() {
	fmt.Println(number_of_subsequence_with_given_sum_k_effcient([]int{2, 9, 4, 3, 6, 9, 1, 1}, 6))
}
