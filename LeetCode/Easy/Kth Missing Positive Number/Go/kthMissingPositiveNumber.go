package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid]-mid-1 < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l + k
}

func main() {
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5))
}
