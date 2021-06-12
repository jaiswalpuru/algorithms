/*
	An sorted array of integers was rotated an unknown number of times.

	Given such an array, find the index of the element in the array in faster than linear time.
	If the element doesn't exist in the array, return null.

	For example, given the array [13, 18, 25, 2, 8, 10] and the element 8, return 4 (the index of 8 in the array).
*/

package main

import "fmt"

func find_index(arr []int, ele int) int {
	if len(arr) <= 0 {
		return -1
	}
	left, right := 0, len(arr)-1
	for left < right {
		mid_point := left + (right-left)/2
		if arr[mid_point] > arr[right] {
			left = mid_point + 1
		} else {
			right = mid_point
		}
	}

	start := left
	left, right = 0, len(arr)-1

	if ele >= arr[start] && ele <= arr[right] {
		left = start
	} else {
		right = start
	}

	for left <= right {
		mid_point := left + (right-left)/2
		if arr[mid_point] == ele {
			return mid_point
		} else if arr[mid_point] > ele {
			right = mid_point
		} else {
			left = mid_point
		}
	}

	return -1
}

func main() {
	arr := []int{13, 18, 25, 2, 8, 10}
	ele := 8 //return the index of this element

	fmt.Println(find_index(arr, ele))
}
