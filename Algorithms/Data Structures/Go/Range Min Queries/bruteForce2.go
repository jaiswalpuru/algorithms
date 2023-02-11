package main

import "fmt"

// One other approach is to preprocess all the queries.
// make a matrix of n*n where n is the size of array.

func rangeMinQuery(nums []int, queries [][]int) {
	n := len(nums)
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			minVal := int(1e7)
			for k := i; k <= j; k++ {
				minVal = min(minVal, nums[k])
			}
			table[i][j] = minVal
		}
	}
	fmt.Println("Preprocessed matrix")
	for i := 0; i < n; i++ {
		fmt.Println(table[i])
	}
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		if query[1] < query[0] {
			panic("check your query ")
		}
		fmt.Println(table[query[0]][query[1]])
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	queries := [][]int{{2, 3}, {4, 8}}
	rangeMinQuery(nums, queries)
}
