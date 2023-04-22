package main

import (
	"fmt"
	"sort"
)

func partition_k_equal_sum_subset(arr []int, k int) bool {
	n := len(arr)
	visited := make([]bool, n)

	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	if sum%k != 0 {
		return false
	}
	//optimized
	sort.Ints(arr)
	return _partition_k_equal_sum_subset_backtrack(arr, 0, k, 0, sum/k, 0, &visited)
}

func _partition_k_equal_sum_subset_backtrack(arr []int, ind, k int, count int, target_sum int, curr_sum int, visited *[]bool) bool {
	if count == k-1 {
		return true
	}

	if curr_sum > target_sum {
		return false
	}

	if curr_sum == target_sum {
		return _partition_k_equal_sum_subset_backtrack(arr, 0, k, count+1, target_sum, 0, visited)
	}

	for i := ind; i < len(arr); i++ {
		if !(*visited)[i] {
			(*visited)[i] = true
			if _partition_k_equal_sum_subset_backtrack(arr, i+1, k, count, target_sum, curr_sum+arr[i], visited) {
				return true
			}
			(*visited)[i] = false
		}
	}
	return false
}

func main() {
	fmt.Println(partition_k_equal_sum_subset([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}
