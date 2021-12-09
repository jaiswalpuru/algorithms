/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that
each unique element appears only once. The relative order of the elements should be kept the same.

Since it is impossible to change the length of the array in some languages,
you must instead have the result be placed in the first part of the array nums. More formally,
if there are k elements after removing the duplicates, then the first k elements of nums should hold the
final result. It does not matter what you leave beyond the first k elements.

Return k after placing the final result in the first k slots of nums.

Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1)
extra memory.
*/

package main

import "fmt"

func remove_duplicates(arr []int) []int {
	n := len(arr)
	if n == 0 || n == 1 {
		return arr
	}
	j := 0
	for i := 0; i < n-1; i++ {
		if arr[i] != arr[i+1] {
			arr[j] = arr[i]
			j++
		}
	}
	arr[j] = arr[n-1]
	j++
	return arr[:j]
}

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(remove_duplicates(arr))
}
