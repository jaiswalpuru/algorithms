package main

import "fmt"

func sum_of_absolute_diff_in_sorted_array(arr []int) []int {
	n := len(arr)
	sum := 0
	res := make([]int, n)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	curr_sum := 0
	for i := 0; i < n; i++ {
		curr_sum += arr[i]
		sum -= arr[i]
		l := n - 1 - i
		res[i] = sum - arr[i]*l + arr[i]*(n-l) - curr_sum
	}
	return res
}

func main() {
	fmt.Println(sum_of_absolute_diff_in_sorted_array([]int{2, 3, 5}))
}
