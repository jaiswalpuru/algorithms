/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example,
the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array
[a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.
*/

package main

import (
	"fmt"
)

func find_min(arr []int) int {
	res := arr[0]
	arr = arr[1:]
	l, h := 0, len(arr)-1
	for l <= h {
		mid := (l + h) / 2
		if res > arr[mid] {
			res = arr[mid]
			h = mid - 1
		}
		if res < arr[mid] {
			l = mid + 1
		}
	}
	return res
}

func main() {
	fmt.Println(find_min([]int{4, 5, 6, 7, 0, 1, 2}))
}
