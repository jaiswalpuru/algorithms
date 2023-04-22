package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-------------------------brute force-------------------------
func wiggle_subsequence(arr []int) int {
	return max(_recur(arr, 0, false), _recur(arr, 0, true))
}

func _recur(arr []int, ind int, is_up bool) int {
	if ind == len(arr) {
		return 0
	}

	not_include := _recur(arr, ind+1, is_up)
	include := 0
	if ind > 0 {
		if (is_up && arr[ind] > arr[ind-1]) || (!is_up && arr[ind] < arr[ind-1]) {
			include = 1 + _recur(arr, ind+1, !is_up)
		}
	} else {
		include = 1 + _recur(arr, ind+1, is_up)
	}

	return max(include, not_include)

}

//-------------------------brute force-------------------------

//-------------------------memoization -------------------------
func wiggle_subsequence_eff(arr []int) int {
	n := len(arr)
	memo := make([][2]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i][0] = -1
		memo[i][1] = -1
	}
	return max(_memo(arr, 0, 0, &memo), _memo(arr, 0, 1, &memo))
}

func _memo(arr []int, ind int, prev int, memo *[][2]int) int {
	if (*memo)[ind][prev] != -1 {
		return (*memo)[ind][prev]
	}

	if ind == len(arr) {
		return 0
	}

	include := 0
	next := 0
	if prev == 0 {
		next = 1
	}
	if ind > 0 {
		if (prev == 1 && arr[ind] > arr[ind-1]) || (prev == 0 && arr[ind] < arr[ind-1]) {
			include = 1 + _memo(arr, ind+1, next, memo)
		}
	} else {
		include = 1 + _memo(arr, ind+1, next, memo)
	}
	not_include := _memo(arr, ind+1, prev, memo)
	(*memo)[ind][prev] = max(include, not_include)
	return (*memo)[ind][prev]
}

//-------------------------memoization -------------------------

func main() {
	fmt.Println(wiggle_subsequence_eff([]int{1, 7, 4, 9, 2, 5}))
}
