/*
Given an integer array nums which is sorted in ascending order and all of its elements are unique and given also an integer k,
return the kth missing number starting from the leftmost number of the array.
*/

package main

import "fmt"

//missing will return number of missing number from the current index element
func missing(idx int, arr []int) int { return arr[idx] - arr[0] - idx }

func missing_element_in_sorted_array(arr []int, k int) int {
	n := len(arr)
	if k > missing(n-1, arr) {
		return arr[n-1] + k - missing(n-1, arr)
	}
	l, h := 0, n-1
	for l < h {
		mid := (l + h) / 2
		if k > missing(mid, arr) {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return arr[l-1] + k - missing(l-1, arr) //returning the element from the left of the array
}

func main() {
	fmt.Println(missing_element_in_sorted_array([]int{4, 7, 9, 10}, 1))
	fmt.Println(missing_element_in_sorted_array([]int{1, 2, 4}, 3))
}
