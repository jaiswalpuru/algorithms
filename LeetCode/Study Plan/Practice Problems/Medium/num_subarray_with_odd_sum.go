/*
Given an array of integers arr, return the number of subarrays with an odd sum.

Since the answer can be very large, return it modulo 109 + 7.
*/

package main

import "fmt"

const (
	mod = 1000000007
)

//brute force TLE
func odd_sum_sub(arr []int) int {
	n := len(arr)

	cnt := 0
	for i := 0; i < n; i++ {
		sum := arr[i]
		if sum%2 != 0 {
			cnt++
		}
		for j := i + 1; j < n; j++ {
			sum += arr[j]
			if sum%2 != 0 {
				cnt++
			}
		}
	}

	return cnt
}

//odd even combination concept
func odd_sum_sub_eff(arr []int) int {
	n := len(arr)
	odd, even := 0, 0

	sum, ans := 0, 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if sum%2 == 0 {
			ans += odd
			even++
		} else {
			ans += 1 + even
			odd++
		}
	}
	return ans
}

func main() {
	fmt.Println(odd_sum_sub_eff([]int{1, 3, 5}))
}
