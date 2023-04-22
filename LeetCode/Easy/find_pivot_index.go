/*
Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to
the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left.
This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.
*/

package main

import "fmt"

func pivot(arr []int) int {

	n := len(arr)
	prefix_sum := make([]int, n+1)
	prefix_sum[0] = 0

	for i := 1; i <= n; i++ {
		prefix_sum[i] = prefix_sum[i-1] + arr[i-1]
	}
	fmt.Println(prefix_sum)
	for i := 1; i <= n; i++ {
		left_sum := prefix_sum[i-1]
		right_sum := prefix_sum[len(prefix_sum)-1] - prefix_sum[i-1] - arr[i-1]
		if left_sum == right_sum {
			return i - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(pivot([]int{1, 7, 3, 6, 5, 6}))
}
