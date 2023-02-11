package main

import (
	"fmt"
	"math"
)

// Here the implementation assumes that queries are valid.
// Few simple checks can be done to check if queries are
// valid.

// This is hybrid two approach
// which develops upon strategies which
// includes -> sparse, block etc
// to improve the preprocess time along
// with query time.
// Complexity : <O(N), O(logn)>

func hybridTwo(nums []int, queries [][]int) {
	size := len(nums)
	blockSize := int(math.Log2(float64(size)))
	blockMinima := []int{}

	fmt.Println("Nums : ", nums)

	//creating the block minima array
	for i := 0; i < size; i += blockSize {
		minVal := nums[i]
		for j := i; j < min(i+blockSize, size); j++ {
			minVal = min(minVal, nums[j])
		}
		blockMinima = append(blockMinima, minVal)
	}
	fmt.Println("Block Minima : ", blockMinima)

	// map to keep track of sparse table for each block
	blockLevelSparseMap := make(map[int][][]int)
	for i := 0; i < len(blockMinima); i++ {
		blockLevelSparseMap[i] = preProcessTable(nums[i*blockSize : i*blockSize+blockSize])
	}

	// create sparse table for the block minima
	blockMinimaSparseTable := preProcessTable(blockMinima)

	//block minima sparse table (summary array)
	fmt.Println("Block minima sparse table")
	printTable(blockMinimaSparseTable)

	// block level sparse map which has mapping of block number and sparse table
	for k, v := range blockLevelSparseMap {
		fmt.Println("Block Number : ", k)
		printTable(v)
	}

	//query
	for i := 0; i < len(queries); i++ {

		l, r := queries[i][0], queries[i][1]
		minValue := math.MaxInt64

		if !isSafe(l, r, size) {
			fmt.Println("Invalid query : ", queries[i])
			continue
		}

		if l == r {
			fmt.Println("Minimum Value in range : ", queries[i], " = ", nums[l])
			continue
		}

		leftBlockNumber := l / blockSize
		rightBlockNumber := r / blockSize

		fmt.Println("Left Block : ", leftBlockNumber)
		fmt.Println("Right Block : ", rightBlockNumber)
		fmt.Println()

		// get the minimum from between blocks in the summary.
		if leftBlockNumber+1 <= rightBlockNumber-1 {
			minValue = min(minValue,
				query(blockMinimaSparseTable, leftBlockNumber+1, rightBlockNumber-1))
		}

		// get the minimum from the left block
		leftBlockMinima := query(blockLevelSparseMap[leftBlockNumber],
			l%blockSize,
			min(leftBlockNumber*blockSize+blockSize-1, r)%blockSize)
		fmt.Println("Minimum value from block : ", leftBlockNumber, " = ", leftBlockMinima)

		// get the minimum from the right block itself
		rightBlockMinima := query(blockLevelSparseMap[rightBlockNumber],
			max(l, rightBlockNumber*blockSize)%blockSize, r%blockSize)
		fmt.Println("Minimum value from block : ", rightBlockNumber, " = ", rightBlockMinima)

		//update the minimum value
		minValue = min(minValue, min(leftBlockMinima, rightBlockMinima))

		fmt.Println("Minimum value in range : ", queries[i], " = ", minValue)
		fmt.Println()
	}

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

func query(sparseTable [][]int, l, r int) int {
	k := int(math.Log2(float64(r - l + 1)))
	return min(sparseTable[k][l], sparseTable[k][r-(1<<k)+1])
}

func isSafe(i, j, size int) bool {
	return i >= 0 && j >= 0 && i < size && j < size && i <= j
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
	nums := []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84, 62, 64, 33, 83, 27}
	queries := [][]int{{0, 14}}
	hybridTwo(nums, queries)
}
