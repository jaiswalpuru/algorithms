/*
Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
*/

package main

import "fmt"

func binary_search(arr []int, target int) bool {
	l, h := 0, len(arr)-1
	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func search(arr [][]int, target int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		if binary_search(arr[i], target) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(search([][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	}, 3))
}
