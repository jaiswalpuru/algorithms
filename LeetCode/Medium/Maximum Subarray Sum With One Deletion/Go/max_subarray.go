package main

import (
	"fmt"
	"math"
)

//-----------------------bad solution brute force dont try to use this----------------------------------------------
func max_subarray(arr []int) int {
	n := len(arr)
	max_val := int(math.MinInt64)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0
			//calculating the sum for the subarray
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
			visited := make(map[int]int)
			max_val = max(max_val, sum)
			if j-i <= 1 {
				continue
			}
			for k := i; k <= j; k++ {
				visited[sum-arr[k]]++
			}

			for k := range visited {
				max_val = max(max_val, k)
			}
		}
	}
	return max_val
}

//----------------------efficient soln uses kadane's algo--------------------------------------------
func max_subarray_bf(arr []int) int {
	n := len(arr)
	max_end, max_start, ans := make([]int, n), make([]int, n), arr[0]
	max_end[0] = arr[0]
	for i := 1; i < n; i++ {
		max_end[i] = max(arr[i], max_end[i-1]+arr[i])
		ans = max(ans, max_end[i])
	}

	max_start[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		max_start[i] = max(arr[i], max_start[i+1]+arr[i])
	}

	for i := 1; i < n-1; i++ {
		ans = max(ans, max_start[i+1]+max_end[i-1])
	}
	return ans
}

//--------------------------------------------------------------------------------------------------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max_subarray_bf([]int{1, -2, 0, 3}))
}
