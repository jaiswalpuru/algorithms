package main

import (
	"fmt"
	"sort"
)

func append_k_integers_with_minimal_sum(arr []int, k int) int64 {
	sort.Ints(arr)
	n := len(arr)
	val := int64(1)
	time := 0
	sum := int64(0)
	for i := 0; i < n; i++ {
		var start int64 = val
		var end int64 = int64(arr[i])

		for start < end && time < k {
			sum += start
			time++
			start++
		}

		val = int64(arr[i]) + 1
	}
	if time < k {
		end := val + int64(k-time-1)
		start := val - 1

		sum += (end * (end + 1) / 2) - (start * (start + 1) / 2)
	}
	return sum
}

func main() {
	fmt.Println(append_k_integers_with_minimal_sum([]int{1, 4, 25, 10, 25}, 2))
}
