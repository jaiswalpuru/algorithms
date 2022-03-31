package main

import (
	"fmt"
	"math"
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

//--------------------------using recursion--------------------------
func split_array_largest_sum(arr []int, m int) int {
	n := len(arr)
	return _split_array_largest_sum(arr, 0, n, m)
}

func _split_array_largest_sum(arr []int, start, n, m int) int {
	//if all the subarrays used and reached the end
	if m == 0 && start == n {
		return 0
	}

	//either reached the end or all the subarrays have been used but not both
	if m == 0 || start == n {
		return math.MaxInt64
	}

	res := math.MaxInt64
	cur_sum := 0

	for i := start; i < n; i++ {
		cur_sum += arr[i]
		next_sub_arr_sum := _split_array_largest_sum(arr, i+1, n, m-1)
		res = min(res, max(cur_sum, next_sub_arr_sum))
	}

	return res
}
//--------------------------using recurion--------------------------

//------------------------using memoization------------------------
func split_larget_sum_using_memo(arr []int, m int) int {
	n := len(arr)
	memo := make([][]int, n)

	for i := 0; i < n; i++ {
		memo[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			memo[i][j] = -1
		}
	}

	return _split_array_largest_sum_memo(arr, &memo, 0, n, m)
}

func _split_array_largest_sum_memo(arr []int, memo *[][]int, start, n, m int) int {
	if m == 0 && start == n {
		return 0
	}

	if m == 0 || start == n {
		return math.MaxInt64
	}

	if (*memo)[start][m] != -1 {
		return (*memo)[start][m]
	}

	res := math.MaxInt64
	curr_sum := 0
	for i := start; i < n; i++ {
		curr_sum += arr[i]
		next_sub_arr_sum := _split_array_largest_sum_memo(arr, memo, i+1, n, m-1)
		res = min(res, max(curr_sum, next_sub_arr_sum))
	}

	(*memo)[start][m] = res
	return (*memo)[start][m]
}
//------------------------using memoization------------------------

//--------------------------using memoization but with prefx array sum (refer LC)--------------------------
func split_larget_sum_using_memo_prefix(arr []int, m int) int {
	n := len(arr)
	prefix := make([]int, n+1)
	prefix[0] = arr[0]

	for i:=0;i<n;i++{
		prefix[i+1] = prefix[i] + arr[i]
	}

	memo := make([][]int, n)
	for i:=0;i<n;i++{
		memo[i] = make([]int, m+1)
		for j:=0;j<=m;j++{
			memo[i][j] =-1
		}
	}

	return _split_array_largest_sum_memo_prefix(prefix, m, &memo,0)
}

func _split_array_largest_sum_memo_prefix(prefix []int, m int,memo *[][]int,ind int) int {
	n := len(prefix)-1

	if (*memo)[ind][m] != -1{
		return (*memo)[ind][m]
	}

	if m==1{
		(*memo)[ind][m] = prefix[n] - prefix[ind]
		return (*memo)[ind][m]
	}

	res := math.MaxInt64
	for i:=ind;i<=n-m;i++{
		curr_sum := prefix[i+1] - prefix[ind]

		next_sub_arr_sum := max(curr_sum, _split_array_largest_sum_memo_prefix(prefix, m-1,memo, i+1))
		res = min(res, next_sub_arr_sum)

		if curr_sum >= res {
			break
		}
	}

	(*memo)[ind][m] = res
	return (*memo)[ind][m]
}
//--------------------------using memoization but with prefx array sum (refer LC)--------------------------

func main() {
	fmt.Println(split_larget_sum_using_memo_prefix([]int{7, 2, 5, 10, 8}, 2))
}
