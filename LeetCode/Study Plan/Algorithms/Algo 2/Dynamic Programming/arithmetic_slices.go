/*
An integer array is called arithmetic if it consists of at least three elements and if the difference between any two
consecutive elements is the same.

For example, [1,3,5,7,9], [7,7,7,7], and [3,-1,-5,-9] are arithmetic sequences.
Given an integer array nums, return the number of arithmetic subarrays of nums.

A subarray is a contiguous subsequence of the array.
*/

package main

import "fmt"

//brute force approach to solve this problem
func arithmetic_slices_bf(arr []int) int {
	n := len(arr)
	cnt := 0
	for i := 0; i < n-2; i++ {
		d := arr[i+1] - arr[i]
		for j := i + 2; j < n; j++ {
			k := 0
			for k = i + 1; k <= j; k++ {
				if arr[k]-arr[k-1] != d {
					break
				}
			}
			if k > j {
				cnt++
			}
		}
	}
	return cnt
}

//better brute force
func arithmetic_slices_bbf(arr []int) int {
	n := len(arr)
	cnt := 0
	for i := 0; i < n-2; i++ {
		d := arr[i+1] - arr[i]
		for j := i + 2; j < n; j++ {
			if arr[j]-arr[j-1] != d {
				break
			} else {
				cnt++
			}
		}
	}
	return cnt
}

//recursion
func arithmetic_slices_recursion(arr []int) int {
	sum := 0
	slices(arr, len(arr)-1, &sum)
	return sum
}

func slices(arr []int, n int, sum *int) int {
	if n < 2 {
		return 0
	}
	temp := 0
	if arr[n]-arr[n-1] == arr[n-1]-arr[n-2] {
		temp = 1 + slices(arr, n-1, sum)
		*sum += temp
	} else {
		slices(arr, n-1, sum)
	}
	return temp
}

//dp
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

func main() {
	fmt.Println(arithmetic_slices_recursion([]int{1, 2, 3, 4}))
	fmt.Println(arithmetic_slices_recursion([]int{1, 2}))
	fmt.Println(arithmetic_slices_dp([]int{1, 2, 3, 4}))
	fmt.Println(arithmetic_slices_dp([]int{1, 2}))
}
