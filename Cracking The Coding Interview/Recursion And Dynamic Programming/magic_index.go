package main

import "fmt"

/*
	here we already know that the array is sorted
	below solutions are for the arrays which contain distinct elements
*/

//--------------straight forward approach----------------
func magic_index(arr []int) int {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == i {
			return i
		}
	}
	return -1
}

//--------------straight forward approach----------------

//--------------can binary search help----------------
func magic_index_b_search(arr []int) int {
	return b_search(arr, 0, len(arr))
}

func b_search(arr []int, low, high int) int {
	res := -1
	for low < high {
		mid := (low + high) / 2
		if mid == arr[mid] {
			res = mid
		}
		if mid > arr[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return res
}

//--------------can binary search help----------------

//--------------arrays which contains duplicate elements----------------
func magic_index_dup(arr []int) int {
	return b_search_dup(arr, 0, len(arr)-1)
}

func b_search_dup(arr []int, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] == mid {
		return mid
	}
	left := min(mid-1, arr[mid])
	l := b_search_dup(arr, low, left)
	if l >= 0 {
		return l
	}
	right := max(mid+1, arr[mid])
	r := b_search_dup(arr, right, high)
	return r
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//--------------arrays which contains duplicate elements----------------

func main() {
	fmt.Println(magic_index_dup([]int{-40, -20, -1, 1, 2, 3, 5, 7, 8, 9, 12, 13}))
}
