/*
Given an array of integers nums and an integer k, return the number of contiguous subarrays
where the product of all the elements in the subarray is strictly less than k.
*/

package main

import "fmt"

func subarray_product(arr []int, k int) int {
	if k < 1 {
		return 0
	}
	cnt := 0
	i, j := 0, 0
	n := len(arr)
	val := 1
	for i < n && j < n {
		val *= arr[j]
		for val >= k && i < n {
			val /= arr[i]
			i++
			if i > n {
				break
			}
		}
		if i < n {
			cnt += j - i + 1
		}
		j++
	}
	return cnt
}

func main() {
	fmt.Println(subarray_product([]int{10, 5, 2, 6}, 100))
	fmt.Println(subarray_product([]int{1, 2, 3}, 0))
	fmt.Println(subarray_product([]int{1, 1, 1}, 1))
}
