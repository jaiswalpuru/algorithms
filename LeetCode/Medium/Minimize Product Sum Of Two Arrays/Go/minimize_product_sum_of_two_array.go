package main

import (
	"fmt"
	"sort"
)

func minimize_product_sum_of_two_arrays(a []int, b []int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	sum := 0
	n := len(a)
	for i := 0; i < n; i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func main() {
	fmt.Println(minimize_product_sum_of_two_arrays([]int{5, 3, 4, 2}, []int{4, 2, 2, 5}))
}
