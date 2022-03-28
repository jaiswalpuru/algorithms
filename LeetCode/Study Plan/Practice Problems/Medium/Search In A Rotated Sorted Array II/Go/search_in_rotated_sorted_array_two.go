package main

import (
	"fmt"
)

//------------------------brute  force------------------------------
func search_in_rotated_sorted_array_two(arr []int, target int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}

//------------------------brute  force------------------------------

//------------------------using binary search--------------------------
func search_in_rotated_sorted_array_two_bs(arr []int, target int) bool {
	if len(arr) == 0 {
		return false
	}

	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return true
		}

		if !is_binary_search_helpful(arr, arr[mid], l) {
			l++
			continue
		}

		//which array does middle element belongs to
		mid_first := exists_in_first(arr, l, arr[mid])
		//which element does target belongs to
		exists_first := exists_in_first(arr, l, target)

		//mid_first, if arr[mid] exists in first array
		if mid_first != exists_first {
			if mid_first {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if arr[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}

func exists_in_first(arr []int, start, target int) bool { return arr[start] <= target }

func is_binary_search_helpful(arr []int, target int, start int) bool { return arr[start] != target }

//------------------------using binary search--------------------------

func main() {
	fmt.Println(search_in_rotated_sorted_array_two([]int{2, 5, 6, 0, 0, 1, 2}, 0))
}
