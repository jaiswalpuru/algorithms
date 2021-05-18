/*
	Given an array of numbers, find the maximum sum of any contiguous subarray of the array.
	[34,-50,42,14,-5,86]
	Max Sum := 137
	Elements := [42,14,-5,86]

	Brute force complexity : O(n^3)
	Kadane's algorithm : O(n)
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max_subarray_sum(arr []int) int {
	max_ending_here, max_so_far := 0, 0
	for i := range arr {
		max_ending_here = max(arr[i], max_ending_here+arr[i])
		max_so_far = max(max_so_far, max_ending_here)
	}
	return max_so_far
}

func min_subarray_sum(arr []int) int {
	min_ending_here, min_so_far := 0, 0
	for i := range arr {
		min_ending_here = min(arr[i], min_ending_here+arr[i])
		min_so_far = min(min_so_far, min_ending_here)
	}
	return min_so_far
}

func sum_array(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func max_circular_array(arr []int) int {
	max_circular_array_sum := sum_array(arr) - min_subarray_sum(arr)
	return max(max_circular_array_sum, max_subarray_sum(arr))
}

func main() {
	arr := []int{34, -50, 42, 14, -5, 86}
	fmt.Println("Maximum subarray sum : ", max_subarray_sum(arr))

	//get the largest wrap around sum.
	//for this we can form minimum subarray sum and subtract this from the total subarray
	fmt.Println("Minimum subarray sum : ", min_subarray_sum(arr))

	fmt.Println("Maximum circular subarray sum : ", max_circular_array(arr))
}
