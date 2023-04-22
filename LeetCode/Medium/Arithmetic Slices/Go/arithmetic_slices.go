package main

import "fmt"

//This is somewhat brute force, this can also be solved using recursion
func arithmetic_slices(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}
	cnt := 0
	for i := 0; i < n-1; i++ {
		diff := arr[i+1] - arr[i]
		for j := i + 2; j < n; j++ {
			if arr[j]-arr[j-1] == diff {
				cnt++
			} else {
				break
			}
		}
	}
	return cnt
}

//-------------------------Using recursion---------------------------
func arithmetic_slices_recursion(arr []int) int {
	cnt := 0
	_arithmetic_slices(arr, len(arr)-1, &cnt)
	return cnt
}

func _arithmetic_slices(arr []int, ind int, cnt *int) int {
	if ind < 2 {
		return 0
	}
	temp := 0
	if arr[ind]-arr[ind-1] == arr[ind-1]-arr[ind-2] {
		temp = 1 + _arithmetic_slices(arr, ind-1, cnt)
		*cnt += temp
	} else {
		_arithmetic_slices(arr, ind-1, cnt)
	}
	return temp
}

//--------------------------------------------------------------------

//-------------------------------Using memoization-------------------------------------
func arithmetic_slices_dp(arr []int) int {
	n := len(arr)
	dp := make([]int, n)
	sum := 0
	for i := 2; i < n; i++ {
		if arr[i]-arr[i-1] == arr[i-1]-arr[i-2] {
			dp[i] = 1 + dp[i-1]
			sum += dp[i]
		}
	}
	return sum
}

//-------------------------------------------------------------------------------------

func main() {
	fmt.Println(arithmetic_slices_dp([]int{1, 2, 3, 4}))
}
