/*
Given an array of numbers, find the length of the longest increasing subsequence in the array.
The subsequence does not necessarily have to be contiguous.

For example, given the array [0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15],
the longest increasing subsequence has length 6: it is 0, 2, 6, 9, 11, 15.
*/

package main

import "fmt"

func longest_inc_subseq_recursive(arr []int, n int, max *int) int {
	if n == 1 {
		return 1
	}
	res := 0
	max_end_here := 1

	for i := 1; i < n; i++ {
		res = longest_inc_subseq_recursive(arr, i, max)
		if arr[i-1] < arr[n-1] && res+1 > max_end_here {
			max_end_here = res + 1
		}
	}
	if *max < max_end_here {
		*max = max_end_here
	}
	return max_end_here
}

func longest_inc_subseq_recursive_wrapper(arr []int) int {
	n := len(arr)
	max := 1
	longest_inc_subseq_recursive(arr, n, &max)
	return max
}

func longest_inc_subseq(arr []int) {
	//length := 0

}

func main() {
	arr := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	fmt.Println(longest_inc_subseq_recursive_wrapper(arr))
	//longest_inc_subseq(arr)
}
