/*
Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

arr.length >= 3
There exists some i with 0 < i < arr.length - 1 such that:
arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
*/

package main

import "fmt"

func valid_mountain(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	i, j := 0, 0
	for i = 0; i < n-1; i++ {
		if arr[i] < arr[i+1] {
			continue
		} else {
			break
		}
	}
	if i == n-1 || i == 0 {
		return false
	}
	for j = i; j < n-1; j++ {
		if arr[j] <= arr[j+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(valid_mountain([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	// fmt.Println(valid_mountain([]int{2, 1}))
	// fmt.Println(valid_mountain([]int{3, 5, 5}))
	// fmt.Println(valid_mountain([]int{0, 3, 2, 1}))
	// fmt.Println(valid_mountain([]int{0, 2, 3, 3, 5, 2, 1, 0}))
	// fmt.Println(valid_mountain([]int{0, 2, 3, 4, 5, 2, 1, 0}))
	// fmt.Println(valid_mountain([]int{2, 0, 2}))
}
