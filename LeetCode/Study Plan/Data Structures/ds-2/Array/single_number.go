/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.
*/

package main

import "fmt"

func single_element(arr []int) int {
	val := arr[0]
	n := len(arr)
	for i := 1; i < n; i++ {
		val = val ^ arr[i]
	}
	return val
}

func main() {
	fmt.Println(single_element([]int{2, 2, 1}))
}
