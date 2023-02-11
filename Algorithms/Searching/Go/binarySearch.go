package main

import "fmt"

//binary search O(lgn)
// there are different templates for binary search
//strongly recommended to practice those on LC

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	return b_search(arr, l, r, target)
}

func b_search(arr []int, l, r, target int) int {
	if l > r {
		return -1
	}
	mid := (l + r) / 2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		return b_search(arr, l, mid-1, target)
	}
	return b_search(arr, mid+1, r, target)
}

//array should be sorted in order to apply binary search
//there are different variations see LC binary search tag.
func main() {
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
}
