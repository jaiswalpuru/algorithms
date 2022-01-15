package main

import "fmt"

func print_sub_count(arr []int, k int) int {
	sum := 0
	return _print_sub_count(arr, k, &sum, 0)
}

func _print_sub_count(arr []int, k int, sum *int, start int) int {
	if start >= len(arr) {
		if k == *sum {
			return 1
		}
		return 0
	}

	*sum += arr[start]
	l := _print_sub_count(arr, k, sum, start+1)
	*sum -= arr[start]
	r := _print_sub_count(arr, k, sum, start+1)
	return l + r
}

func main() {
	fmt.Println(print_sub_count([]int{4, 3, 2, 1, 5, 6}, 6))
}
