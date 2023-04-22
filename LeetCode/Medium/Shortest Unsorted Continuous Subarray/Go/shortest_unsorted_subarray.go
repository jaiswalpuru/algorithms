package main

import (
	"fmt"
	"math"
	"sort"
)

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

//-------------------------------------------bad brute force-------------------------------------------
func solve(arr []int) int {
	n := len(arr)
	res := n
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			max_val, min_val, prev := int(math.MinInt64), int(math.MaxInt64), int(math.MinInt64)

			for k := i; k < j; k++ {
				min_val = min(min_val, arr[k])
				max_val = max(max_val, arr[k])
			}

			//this is just checking if the size has already been calculated
			if (i > 0 && arr[i-1] > min_val) || (j < n && arr[j] < max_val) {
				continue
			}

			k := 0
			for k < i && prev <= arr[k] {
				prev = arr[k]
				k++
			}
			if i == 1 && j == 6 {
				fmt.Println(k)
			}

			if k != i {
				continue
			}
			k = j
			for k < n && prev <= arr[k] {
				prev = arr[k]
				k++
			}

			if k == n {
				res = min(res, j-i)
			}
		}
	}
	return res
}

//--------------------------------------------------------------------------------------------------------------------------

//------------------------------------------Brute force optimized-----------------------------------------------------------

func solve_optimized(arr []int) int {
	n := len(arr)
	l, r := n, 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[i] {
				l = min(l, i)
				r = max(r, j)
			}
		}
	}
	if r-l < 0 {
		return 0
	} else {
		return r - l + 1
	}
}

//----------------------------------------------------------------------------------------------------------------------------

//-------------------------------------------------------Using Sorting---------------------------------------------------------
func solve_using_sorting(arr []int) int {
	n := len(arr)
	arr_temp := make([]int, n)
	copy(arr_temp, arr)
	sort.Ints(arr_temp)
	l, r := -1, -1
	for i := 0; i < n; i++ {
		if arr[i] != arr_temp[i] {
			if l == -1 {
				l = i
			} else {
				r = i
			}
		}
	}
	if r-l < 0 {
		return 0
	}
	return r - l + 1
}

//---------------------------------------------------------------------------------------------------------------------------

//-------------------------------------------Using Stack-----------------------------------------------------------------

func solve_using_stack(arr []int) int {
	n := len(arr)
	stack := []int{}
	l, r := n, 0
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
			l = min(l, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] < arr[i] {
			r = max(r, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	if r-l < 0 {
		return 0
	}
	return r - l + 1
}

//---------------------------------------------------------------------------------------------------------------------------
func main() {
	fmt.Println(solve_using_stack([]int{2, 6, 4, 8, 10, 9, 15}))
}
