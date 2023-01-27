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
	n := len(nums)
	m := int(math.Log2(float64(n))) + 1
	sparseTable := make([][]int, n)
	for i := 0; i < n; i++ {
		sparseTable[i] = make([]int, m)
	}
	//create the sparse table
	for i := 0; i < n; i++ {
		sparseTable[i][0] = nums[i]
	}
	for j := 1; j < m; j++ {
		for i := 0; i+(1<<(j-1)) < n; i++ {
			sparseTable[i][j] = min(sparseTable[i][j-1], sparseTable[i+(1<<(j-1))][j-1])
		}
	}
	for i := 0; i < len(queries); i++ {
		l, r := queries[i][0], queries[i][1]
		k := int(math.Log2(float64(r - l + 1)))
		pow2 := int(math.Pow(2, float64(k)))
		fmt.Println("minimum value in range : ", queries[i], " = ", min(sparseTable[l][k], sparseTable[r-pow2+1][k]))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{24, 32, 58, 6, 94, 86, 16, 20}
	queries := [][]int{{2, 7}}
	rangeMinQuery(nums, queries)
}
