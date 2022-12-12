package main

import "fmt"

//-----------------------brute force------------------------
func maximum_alternating_subsequence_sum(arr []int) int64 {
	return recur(arr, 0, true)
}

func recur(arr []int, ind int, flag bool) int64 {
	if ind >= len(arr) {
		return int64(0)
	}
	if flag {
		val := recur(arr, ind+1, false) + int64(arr[ind])
		val = max(val, recur(arr, ind+1, true))
		return val
	} else {
		val := recur(arr, ind+1, true) - int64(arr[ind])
		val = max(val, recur(arr, ind+1, false))
		return val
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

//-----------------------brute force------------------------

//-------------------------memoization----------------------
func maximum_alternating_subsequence_sum_memo(arr []int) int64 {
	memo := make([][]int64, len(arr))
	for i := 0; i < len(arr); i++ {
		memo[i] = make([]int64, 2)
		memo[i][0] = -1
		memo[i][1] = -1
	}
	return _memo(arr, 0, 1, &memo)
}

func _memo(arr []int, ind int, flag int, memo *[][]int64) int64 {
	if ind >= len(arr) {
		return int64(0)
	}
	if (*memo)[ind][flag] != -1 {
		return (*memo)[ind][flag]
	}
	if flag == 1 {
		val := _memo(arr, ind+1, 0, memo) + int64(arr[ind])
		val = max(val, _memo(arr, ind+1, 1, memo))
		(*memo)[ind][flag] = val
		return (*memo)[ind][flag]
	} else {
		val := _memo(arr, ind+1, 1, memo) - int64(arr[ind])
		val = max(val, _memo(arr, ind+1, 0, memo))
		(*memo)[ind][flag] = val
		return (*memo)[ind][flag]
	}
}

//-------------------------memoization----------------------

func main() {
	fmt.Println(maximum_alternating_subsequence_sum_memo([]int{4, 2, 5, 3}))
}
