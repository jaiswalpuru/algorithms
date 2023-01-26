package main

import "fmt"

// Here the first approach would be loop through each
// query and for each query (u, v) run a loop from
// u to v and return the minimum value
// O(n^2)

func rangeMinQuery(nums []int, queries [][]int) {
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		u, v := query[0], query[1]
		minVal := int(1e7)
		for j := u; j <= v; j++ {
			minVal = min(minVal, nums[j])
		}
		fmt.Println("For query : ", query, " minimum value is : ", minVal)
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
