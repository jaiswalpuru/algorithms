package main

import (
	"fmt"
	"math"
)

// This is hybrid one approach
// which develops upon strategies which
// includes -> sparse, block etc
// to improve the preprocess time along
// with query time.
// Complexity : <O(N), O(logn)>

func hybridOne(nums []int, queries [][]int) {
	fmt.Println(nums)
	size := len(nums)
	blockSize := int(math.Log2(float64(size)))
	blockMinima := []int{}

	// map to keep track of array each block
	blockLevelArr := make(map[int][]int)
	k := 0
	//creating the block minima array
	for i := 0; i < size; i += blockSize {
		minVal := nums[i]
		tempArr := []int{}
		for j := i; j < min(i+blockSize, size); j++ {
			minVal = min(minVal, nums[j])
			tempArr = append(tempArr, nums[j])
		}
		blockMinima = append(blockMinima, minVal)
		blockLevelArr[k] = tempArr
		k++
	}
	fmt.Println("Block Minima : ", blockMinima)

	// create sparse table for the block minima
	blockMinimaSparseTable := preProcessTable(blockMinima)

	//block minima sparse table (summary array)
	fmt.Println("Block minima sparse table")
	printTable(blockMinimaSparseTable)

	// block level sparse map which has mapping of block number and sparse table
	for k, v := range blockLevelArr {
		fmt.Println("Block Number : ", k)
		fmt.Println(v)
		fmt.Println()
	}

	// miniumum value for all the queries
	for i := 0; i < len(queries); i++ {
		fmt.Println("Processing query : ", queries[i])
		minValue := math.MaxInt64

		low, high := queries[i][0], queries[i][1]

		if !isSafe(low, high, size) {
			fmt.Println("Invlaid case ", " low : ", low, " , high : ", high)
			continue
		}

		if low == high {
			fmt.Println("Mininum value for query : ", queries[i], " = ", nums[low])
			continue
		}

		leftBlock := low / blockSize
		rightBlock := high / blockSize

		fmt.Println("Left Block : ", leftBlock)
		fmt.Println("Right Block : ", rightBlock)
		fmt.Println()

		if leftBlock+1 <= rightBlock-1 {
			minValue = min(minValue, querySparseTable(blockMinimaSparseTable, leftBlock+1, rightBlock-1))
		}

		//get minimum from left block
		leftBlockMinima := getMinimumFromArr(blockLevelArr[leftBlock], low%blockSize,
			min(leftBlock*blockSize+blockSize-1, high)%blockSize)

		//get minimum from right block
		rightBlockMinima := getMinimumFromArr(blockLevelArr[rightBlock],
			max(low, rightBlock*blockSize)%blockSize, high%blockSize)

		// update the minimum value
		minValue = min(minValue, min(leftBlockMinima, rightBlockMinima))

		fmt.Println("Minimum value in range : ", queries[i], " = ", minValue)
		fmt.Println()
	}
}

func querySparseTable(sparseTable [][]int, l, r int) int {
	k := int(math.Log2(float64(r - l + 1)))
	res := min(sparseTable[k][l], sparseTable[k][r-(1<<k)+1])
	return res
}

func preProcessTable(nums []int) [][]int {
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
	for i := 1; i < n; i++ {
		for j := 0; j+(1<<(i-1)) < m; j++ {
			sparseTable[i][j] = min(sparseTable[i-1][j], sparseTable[i-1][j+(1<<(i-1))])
		}
	}
	return sparseTable
}

func isSafe(i, j, size int) bool { return i >= 0 && j >= 0 && i <= j && i < size && j < size }

func getMinimumFromArr(arr []int, low, high int) int {
	minVal := math.MaxInt64
	for i := low; i <= high; i++ {
		minVal = min(minVal, arr[i])
	}
	return minVal
}

func printTable(nums [][]int) {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	fmt.Println()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{40, 32, 58, 24, 12, 86, 16, 28, 99, 2, 39, 40, 64, 47, 70, 13}
	queries := [][]int{{0, 15}}
	hybridOne(nums, queries)
}
