package main

import (
	"fmt"
	"sort"
)

func largest_perimeter_triangle(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	for i := n - 3; i >= 0; i-- {
		if arr[i]+arr[i+1] > arr[i+2] {
			return arr[i] + arr[i+1] + arr[i+2]
		}
	}
	return 0
}

func main() {
	fmt.Println(largest_perimeter_triangle([]int{2, 1, 2}))
}
