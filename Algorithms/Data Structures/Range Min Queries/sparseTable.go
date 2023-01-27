package main

import (
	"fmt"
	"math"
)

// so it if we see while creating the dp table earlier
// in brute force approach there were many duplicates
// the idea is to create a 2d table of n * log(n) size.
// preprocessing complexity is O(n*log(n))
// query time is O(1)

func rangeMinQuery(nums []int, queries [][]int) {
	m := len(nums)
	n := int(math.Log2(float64(m))) + 1
	sparseTable := make([][]int, n)
	for i := 0; i < n; i++ {
		sparseTable[i] = make([]int, m)
	}
	//create the sparse table
	//Fill the first row
	for i := 0; i < m; i++ {
		sparseTable[0][i] = nums[i]
	}
	// (j+(1<<(i-1))) for checking which neighbors to compare
	// so lets say if j=0 and i=1 we just need to compare the adjacent elements. when i=1
	// i.e we are calculating the row with 2^1
	// eg1 : j = 0 && i = 1 value of (j + (1<<(i-1))) = 1
	// eg2 : j = 1 && i = 2 value = 1 + (1<<(2-1)) = 3
	fmt.Println(sparseTable[0])
	for i := 1; i < n; i++ {
		for j := 0; j+(1<<(i-1)) < m; j++ {
			sparseTable[i][j] = min(sparseTable[i-1][j], sparseTable[i-1][j+(1<<(i-1))])
		}
		fmt.Println(sparseTable[i])
	}

	// now answer the queries which will take O(1)
	for i := 0; i < len(queries); i++ {
		l, r := queries[i][0], queries[i][1]
		k := int(math.Log2(float64(r - l + 1)))
		res := min(sparseTable[k][l], sparseTable[k][r-(1<<k)+1])
		fmt.Println("For query : ", queries[i], " minimum value = ", res)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{24, 32, 58, 600, 94, 86, 16, 20}
	queries := [][]int{{0, 1}, {2, 7}, {1, 6}, {5, 7}}
	rangeMinQuery(nums, queries)
}
