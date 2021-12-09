/*
Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
The relative order of the elements may be changed.

Since it is impossible to change the length of the array in some languages, you must instead have the
result be placed in the first part of the array nums. More formally, if there are k elements after removing
the duplicates, then the first k elements of nums should hold the final result. It does not matter what you
leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.
*/

package main

import "fmt"

func remove_elements(arr []int, val int) []int {
	n := len(arr)

	if n == 0 {
		return arr
	}

	for i := 0; i < n; i++ {
		if arr[i] == val {
			arr[i] = -1
		}
	}

	j := 0
	for i := 0; i < n; i++ {
		if arr[i] != -1 {
			arr[j] = arr[i]
			j++
		}
	}
	return arr[:j]
}

func main() {
	fmt.Println(remove_elements([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
