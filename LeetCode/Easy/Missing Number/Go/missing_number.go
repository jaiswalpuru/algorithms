package main

import "fmt"

func missing_number(a []int) int {
	n := len(a)
	sum := n * (n + 1) / 2
	total_sum := 0
	for i := 0; i < n; i++ {
		total_sum += a[i]
	}

	return sum - total_sum
}

func main() {
	fmt.Println(missing_number([]int{3, 0, 1}))
}
