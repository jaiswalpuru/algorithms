package main

import "fmt"

func maximum_candies_allocated_to_k_children(arr []int, k int64) int {
	l, r := 0, int(1e7)

	for l < r {
		sum := int64(0)
		mid := (l + r + 1) / 2
		for i := 0; i < len(arr); i++ {
			sum += int64(arr[i] / mid)
		}
		if k > sum {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

func main() {
	fmt.Println(maximum_candies_allocated_to_k_children([]int{5, 8, 6}, 3))
}
