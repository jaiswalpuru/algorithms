package main

import (
	"fmt"
	"math"
)

//-----------------------------------bad brute force-----------------------------------
func is_132_pattern_bf(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i] < arr[k] && arr[j] > arr[k] {
					return true
				}
			}
		}
	}

	return false
}

//---------------------------------------------------------------------------------------

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//-----------------------------------better brute force------------------------
func is_132_pattern_bf_eff(arr []int) bool {
	n := len(arr)
	min_val := math.MaxInt64
	for j := 0; j < n; j++ {
		min_val = min(min_val, arr[j])
		for k := j + 1; k < n; k++ {
			if arr[j] > arr[k] && min_val < arr[k] {
				return true
			}
		}
	}
	return false
}

//---------------------------------------------------------------------------

//-------------------------------using monostack-----------------------------
func is_132_pattern(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	stack := []int{}
	min_arr := make([]int, n)
	min_arr[0] = arr[0]
	for i := 1; i < n; i++ {
		min_arr[i] = min(min_arr[i-1], arr[i])
	}

	for j := n - 1; j >= 0; j-- {
		if min_arr[j] < arr[j] {
			for len(stack) > 0 && stack[len(stack)-1] <= min_arr[j] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] < arr[j] {
				return true
			}
			stack = append(stack, arr[j])
		}
	}
	return false
}

//-------------------------------using monostack-----------------------------

func main() {
	fmt.Println(is_132_pattern([]int{1, 2, 3, 4}))
}
