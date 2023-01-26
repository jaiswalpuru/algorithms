package main

import "fmt"

// In brute force 2 we build the table in O(N^3)
// This table can be built in O(N^2)
// start by filling the diagonal elements and then its aadjacent elements

func rangeMinQuery(nums []int, queries [][]int) {
	n := len(nums)
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
	}
	// 1. Fill the diagonal elements
	for i := 0; i < n; i++ {
		table[i][i] = nums[i]
	}
	//now start filling the adjacent elements
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j > i {
				table[i][j] = min(table[i][j-1], nums[j])
			}
		}
	}

	//now answer the queries
	for i := 0; i < len(queries); i++ {
		q := queries[i]
		fmt.Println("Query : ", q, " min element : ", table[q[0]][q[1]])
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
