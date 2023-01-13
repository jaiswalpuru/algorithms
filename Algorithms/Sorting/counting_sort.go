package main

import (
	"fmt"
	"math"
)

/*
	Counting sort : maintains a map of frequency
	of each element and sorts based on that
*/

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

func sort(arr []int) []int {
	min_val, max_val := math.MaxInt64, math.MinInt64
	for i := 0; i < len(arr); i++ {
		min_val = min(min_val, arr[i])
		max_val = max(max_val, arr[i])
	}
	size := max_val - min_val + 1
	freq := make([]int, size)
	//count the freq here by arr[i]-min_val
	for i := 0; i < len(arr); i++ {
		freq[arr[i]-min_val]++
	}
	k := 0
	for i := 0; i < len(freq); i++ {
		for freq[i] > 0 {
			arr[k] = i + min_val
			freq[i]--
			k++
		}
	}
	return arr
}

func main() {
	fmt.Println(sort([]int{9, 9, 9, 8, 8, 8, 0, 0, 0, -5, -5, -5, 5, 5, 5}))
}
