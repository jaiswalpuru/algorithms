package main

import (
	"fmt"
	"math"
)

//----------------------using monotonic stack----------------------
func verify(arr []int) bool {
	n := len(arr)

	stack := []int{}
	root := math.MinInt64
	for i := 0; i < n; i++ {
		if arr[i] < root {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] < arr[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	return true
}

//------------------------------------------------------------------

func main() {
	fmt.Println(verify([]int{5, 2, 1, 3, 6}))
}
