package main

import "fmt"
import "sort"

func partition_array_such_that_max_diff_is_k(arr []int, k int) int {
	n := len(arr)
	sort.Ints(arr)
	val := arr[0]
	par := 1 
	for i := 1; i < n; i++ {
		if arr[i]-val > k {
			val = arr[i]
			par++
		}
	}
	return par
}

func main() {
	fmt.Println(partition_array_such_that_max_diff_is_k([]int{3, 6, 1, 2, 5}, 2))
}
