package main

import "fmt"

//-----------------------------brute force-----------------------------
func subarray_sum(arr []int, k int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		sum := arr[i]
		if sum == k {
			res++
		}
		for j := i + 1; j < n; j++ {
			sum += arr[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}

//------------------------------------------------------------------------

//-------------------------------opitimizing using map---------------------
func subarray_sum_map(arr []int, k int) int {
	n := len(arr)
	res := 0
	sum := 0
	visited := make(map[int]int)
	for i := 0; i < n; i++ {
		sum += arr[i]
		if sum == k {
			res++
		}
		if _, ok := visited[sum-k]; ok {
			res += visited[sum-k]
		}
		visited[sum]++
	}
	return res
}

//-------------------------------------------------------------------------

func main() {
	fmt.Println(subarray_sum_map([]int{1, 2, 3}, 3))
}
