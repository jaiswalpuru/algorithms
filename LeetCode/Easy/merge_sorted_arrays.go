/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/

package main

import "fmt"

func merge_sorted_arrays(a, b []int, n, m int) []int {
	p, q := n-1, m-1
	for i := m + n - 1; i >= 0; i-- {
		if q < 0 {
			break
		}
		if p >= 0 && a[p] > b[q] {
			a[i] = a[p]
			p--
		} else {
			a[i] = b[q]
			q--
		}
	}
	return a
}

func main() {
	fmt.Println(merge_sorted_arrays([]int{4, 0, 0, 0, 0, 0}, []int{1, 2, 3, 5, 6}, 1, 5))
}
