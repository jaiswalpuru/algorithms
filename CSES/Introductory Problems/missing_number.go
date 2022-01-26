package main

import "fmt"

func missing_number(n int, arr []int) int {
	total_sum := n * (n + 1) / 2
	arr_sum := 0
	m := len(arr)
	for i := 0; i < m; i++ {
		arr_sum += arr[i]
	}
	return total_sum - arr_sum
}

func main() {
	fmt.Println(missing_number(5, []int{2, 3, 1, 5}))
}
