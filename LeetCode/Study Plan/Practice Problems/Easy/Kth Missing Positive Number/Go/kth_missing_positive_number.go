package main

import "fmt"

func kth_missing_positive_number(arr []int, k int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid]-mid-1 < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l + k
}

func main() {
	fmt.Println(kth_missing_positive_number([]int{2, 3, 4, 7, 11}, 5))
}
