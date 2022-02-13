package main

import (
	"fmt"
	"math"
	"sort"
)

func remove_minimum_beans(arr []int) int {
	n := len(arr)
	sort.Ints(arr)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	res := math.MaxInt64
	for i := 0; i < n; i++ {
		res = min(res, sum-((n-i)*arr[i]))
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(remove_minimum_beans([]int{2, 10, 2, 3}))
}
