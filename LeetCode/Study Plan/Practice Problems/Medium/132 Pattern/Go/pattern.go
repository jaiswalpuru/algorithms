package main

import "fmt"
import "math"

//------------------------------brute force-------------------------------(do not try this)
func pattern(arr []int) bool {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i] < arr[k] && arr[k] < arr[j] {
					return true
				}
			}
		}
	}
	return false
}

//--------------------------------------------------------------------------------

//-----------------------------a better brute force------------------------------
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func pattern_better(arr []int) bool {
	n := len(arr)
	min_val := math.MaxInt64
	for i := 0; i < n; i++ {
		min_val = min(min_val, arr[i])
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] && min_val < arr[j] {
				return true
			}
		}
	}
	return false
}

//-------------------------------------------------------------------------------

//-------------------------------using mono stack---------------------------------
func pattern_eff(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	st := []int{}
	min_val := make([]int, n)
	min_val[0] = arr[0]
	for i := 1; i < n; i++ {
		min_val[i] = min(min_val[i-1], arr[i])
	}

	for j := n - 1; j >= 0; j-- {
		if min_val[j] < arr[j] {
			for len(st) > 0 && st[len(st)-1] <= min_val[j] {
				st = st[:len(st)-1]
			}
			if len(st) > 0 && st[len(st)-1] < arr[j] {
				return true
			}
			st = append(st, arr[j])
		}
	}
	return false
}

//-------------------------------using mono stack---------------------------------

func main() {
	fmt.Println(pattern_eff([]int{3,1,4,2}))
}
