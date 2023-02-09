package main

import (
	"fmt"
	"math"
)

// This algorithm is for range min queries
// built on hybrid approaches. In this the
// concept of cartesian tree is used which
// is which is basically to represent a block
// of array using a cartesian number. This is
// solved using stack. In which every to push
// operation is encoded as 1 and every pop operation
// is encoded as 0;
// Complexity of Fischer Heun <O(n), O(1)>

func fischerHeun(arr []int, queries [][]int) {
	fmt.Println("Nums : ", arr)
	size := len(arr)
	blockSize := 3         //int((0.5) * math.Log(size) / math.Log(4))
	blockMinima := []int{} // to store the block minima

	// this is not required here but in java I used it create array for RMQFullTable objects
	// so the maximum size of the array will be as shown below
	//totalCartesianNumber := int(math.Pow(4, float64(blockSize)))

	// create the block minima
	for i := 0; i < size; i += blockSize {
		minVal := arr[i]
		for j := i; j < min(i+blockSize, size); j++ {
			minVal = min(minVal, arr[j])
		}
		blockMinima = append(blockMinima, minVal)
	}

	fmt.Println("Block Minima : ", blockMinima)

	blockNumberToArr := make(map[int][]int)     // mapping of block number to array
	blockToCartesianNumber := make(map[int]int) // mapping of block number to cartesian number

	rmqFullTable := make(map[int][][]int)

	for i := 0; i < len(blockMinima); i++ {
		blockNumberToArr[i] = arr[i*blockSize : min(i*blockSize+blockSize, size)]
		blockToCartesianNumber[i] = getCartesianNumber(blockNumberToArr[i])
		rmqFullTable[blockToCartesianNumber[i]] = getFullTablePreprocess(blockNumberToArr[i])
	}

	fmt.Println("Block number to array mapping : ", blockNumberToArr)
	fmt.Println("Block number to cartesian number mapping : ", blockToCartesianNumber)

	// prepare the sparse table for block minima
	rmqSparseTableBlockMinima := getSparseTable(blockMinima)
	fmt.Println("RMQ Sparse Table for block minima : ", blockMinima)
	for i := 0; i < len(rmqSparseTableBlockMinima); i++ {
		fmt.Println(rmqSparseTableBlockMinima[i])
	}

	//Block Number to Array
	fmt.Println("Block Number to array")
	for i := 0; i < len(blockMinima); i++ {
		fmt.Println("Block Number : ", i)
		fmt.Println(blockNumberToArr[i])
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Mapping of cartesian number to full table")
	for k, v := range rmqFullTable {
		fmt.Println("Cartesian number : ", k)
		for i := 0; i < len(v); i++ {
			fmt.Println(v[i])
		}
		fmt.Println()
	}

	//query the results
	for i := 0; i < len(queries); i++ {
		l, r := queries[i][0], queries[i][1]

		if !isSafe(l, r, size) {
			fmt.Println("Invalid query : ", queries[i])
			continue
		}

		if l == r {
			fmt.Println("Minimum value in range : ", queries[i], " = ", arr[l])
		}

		minValue := math.MaxInt64

		//get the left and right block
		leftBlock := l / blockSize
		rightBlock := r / blockSize

		fmt.Println("Left Block : ", leftBlock)
		fmt.Println("Right Block : ", rightBlock)

		if leftBlock+1 <= rightBlock-1 {
			minValue = min(minValue, querySparseTable(rmqSparseTableBlockMinima, leftBlock+1, rightBlock-1))
			fmt.Println("Minimum from ", leftBlock+1, " and ", rightBlock-1, "  = ", minValue)
		}

		//get the cartesian number for left block
		leftBlockCartesianNumber := blockToCartesianNumber[leftBlock]
		fmt.Println("Left Block Cartesian number : ", leftBlockCartesianNumber)
		minValue = min(minValue, queryFullTable(arr, rmqFullTable[leftBlockCartesianNumber], l%blockSize, min(leftBlock*blockSize+blockSize-1, r)%blockSize))

		// get the cartesian number for right block
		rightBlockCartesianNumber := blockToCartesianNumber[rightBlock]
		fmt.Println("Right Block Cartesian number : ", rightBlockCartesianNumber)
		minValue = min(minValue, queryFullTable(arr, rmqFullTable[rightBlockCartesianNumber], max(l, rightBlock*blockSize)%blockSize, r%blockSize))

		fmt.Println("Minimum value in range ,", queries[i], " = ", minValue)
	}

}

func queryFullTable(arr []int, table [][]int, l, r int) int {
	return arr[table[l][r]]
}

func getSparseTable(nums []int) [][]int {
	m := len(nums)
	n := int(math.Log2(float64(m))) + 1
	sparseTable := make([][]int, n)
	for i := 0; i < n; i++ {
		sparseTable[i] = make([]int, m)
	}

	// fill the first row as it is
	for i := 0; i < m; i++ {
		sparseTable[0][i] = nums[i]
	}

	// now start comparing the elements in
	// powers of two
	for i := 1; i < n; i++ {
		for j := 0; j+(1<<(i-1)) < m; j++ {
			sparseTable[i][j] = min(sparseTable[i-1][j], sparseTable[i-1][j+(1<<(i-1))])
		}
	}
	return sparseTable
}

func querySparseTable(sparseTable [][]int, l, r int) int {
	k := int(math.Log2(float64(r - l + 1)))
	res := min(sparseTable[k][l], sparseTable[k][r-(1<<k)+1])
	return res
}

func getFullTablePreprocess(nums []int) [][]int {
	n := len(nums)
	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
	}
	// 1. Fill the diagonal elements
	for i := 0; i < n; i++ {
		table[i][i] = i
	}
	//now start filling the adjacent elements
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[table[i][j-1]] > nums[j] {
				table[i][j] = j
			} else {
				table[i][j] = table[i][j-1]
			}
		}
	}
	return table
}

func getCartesianNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	stack := []int{nums[0]}
	bits := []int{1}

	for i := 1; i < len(nums); i++ {
		for len(stack) > 0 && stack[len(stack)-1] > nums[i] {
			stack = stack[:len(stack)-1]
			bits = append(bits, 0)
		}
		stack = append(stack, nums[i])
		bits = append(bits, 1)
	}

	for len(stack) > 0 {
		stack = stack[:len(stack)-1]
		bits = append(bits, 0)
	}

	j := 0
	cartesianNumber := 0
	for i := len(bits) - 1; i >= 0; i-- {
		cartesianNumber += (1 << j) * bits[i]
		j++
	}

	return cartesianNumber

}

func isSafe(i, j, size int) bool {
	return i >= 0 && j >= 0 && i < size && j < size && i <= j
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{66, 13, 59, 75, 55, 2, 6, 17, 25, 84, 44, 0}
	queries := [][]int{{0, 1}}
	fischerHeun(nums, queries)
}
