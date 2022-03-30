package main

import "fmt"

func binary_search(arr []int, l, h, target int) bool {
	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == target {
			return true
		}
		if arr[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func search_in_2D_matrix(arr [][]int, target int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		if binary_search(arr[i], 0, len(arr[i])-1, target) {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(search_in_2D_matrix([][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	}, 3))
}
