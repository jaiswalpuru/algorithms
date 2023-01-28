package main

import (
	"fmt"
	"math"
)

// using block partition
// in this approach we divide the array into
// blocks and calculate the minimum for each
// block. and based on the query we see to which
// block it belongs to and return res
// O(sqrt(N)) -> query time and pre-processing time : O(N)

func rangeMinQuery(nums []int, queries [][]int) {
	block := []int{}
	b := int(math.Sqrt(float64(len(nums)-1))) + 1
	for i := 0; i < len(nums); i += b {
		minVal := int(1e7)
		for j := i; j < min(i+b, len(nums)); j++ {
			minVal = min(minVal, nums[j])
		}
		block = append(block, minVal)
	}
	for i := 0; i < len(queries); i++ {
		l, r := queries[i][0], queries[i][1]
		lBlock, rBlock := l/b, r/b
		minVal := int(1e7)
		//fetch the minimum values for in between blocks
		for k := lBlock + 1; k <= rBlock-1; k++ {
			minVal = min(minVal, block[k])
		}
		//now get the values for block lBlock and r Block
		if l == lBlock*b && r > (lBlock*b-1) {
			minVal = min(minVal, block[lBlock])
		} else {
			for j := l; j < min((lBlock*b+b), r+1); j++ {
				minVal = min(minVal, nums[j])
			}
		}
		if r == rBlock*b+b-1 && l < rBlock*b {
			minVal = min(minVal, block[rBlock])
		} else {
			for j := max(l, rBlock*b); j <= r; j++ {
				minVal = min(minVal, nums[j])
			}
		}
		fmt.Println("for query ", queries[i], " minimum value = ", minVal)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{6, 32, 58, 24, 94, 86, 16, 2, 99, 28, 39, 40, 64, 47, 70, 13, 12}
	queries := [][]int{{1, 14}, {0, 7}, {0, 1}, {1, 3}, {15, 16}}
	rangeMinQuery(nums, queries)
}
