/*
Given an array of integers, write a function to determine whether the array could become non-decreasing
by modifying at most 1 element.

For example, given the array [10, 5, 7], you should return true, since we can modify the 10 into a 1
to make the array non-decreasing.

Given the array [10, 5, 1], you should return false, since we can't modify any one element to get a
non-decreasing array.
*/

package main

import "fmt"

func is_possible(arr []int) bool {
	cnt := 0
	n := len(arr)
	if n == 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if arr[0] > arr[1] {
		arr[0] = arr[1] / 2
		cnt++
	}

	for i := 1; i < n-1; i++ {
		if (arr[i] > arr[i-1] && arr[i] > arr[i+1]) || (arr[i] < arr[i-1] && arr[i] < arr[i+1]) {

			arr[i] = (arr[i-1] + arr[i+1]) / 2

			if arr[i-1] == arr[i] || arr[i] == arr[i+1] {
				return false
			}
			cnt++
		}
	}
	if arr[n-1] < arr[n-2] {
		cnt++
	}
	if cnt > 1 {
		return false
	}
	return true
}

func main() {
	fmt.Println(is_possible([]int{10, 5, 7}))
	fmt.Println(is_possible([]int{10, 5, 1}))
}
