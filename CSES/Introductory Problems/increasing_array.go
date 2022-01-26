package main

import "fmt"

func increasing_array(n int, arr []int) int {
	min_moves := 0
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			min_moves += arr[i-1] - arr[i]
		}
	}
	return min_moves
}

func main() {
	fmt.Println(increasing_array(5, []int{3, 2, 5, 1, 7}))
}
