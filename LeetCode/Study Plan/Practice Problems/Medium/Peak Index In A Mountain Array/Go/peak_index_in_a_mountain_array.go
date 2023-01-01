package main

import "fmt"

func peak_index_in_a_mountain_array(arr []int) int {
	return b_search(arr, 0, len(arr)-1)
}

func b_search(arr []int, l, r int) int {
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	fmt.Println(peak_index_in_a_mountain_array([]int{0, 1, 0}))
}
