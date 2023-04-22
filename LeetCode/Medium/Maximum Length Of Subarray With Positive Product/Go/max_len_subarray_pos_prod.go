package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------------bad brute force------------------------------------ (will give wrong answer due to overflow)
func max_len(arr []int) int {
	n := len(arr)
	ans := 0
	mul_val := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res := 1
			for k := i; k <= j; k++ {
				res = (res * arr[k])
			}
			if res >= mul_val {
				mul_val = res
				ans = max(ans, j-i+1)
			}
		}
	}
	return ans
}

//----------------------------------------------------------------------------------------

func max_len_eff(arr []int) int {
	n := len(arr)
	pos := make([]int, n)
	neg := make([]int, n)
	if arr[0] > 0 {
		pos[0] = 1
	} else if arr[0] < 0 {
		neg[0] = 1
	}
	ans := pos[0]
	for i := 1; i < n; i++ {
		if arr[i] > 0 {
			pos[i] = 1 + pos[i-1]
			if neg[i-1] > 0 {
				neg[i] = 1 + neg[i-1]
			} else {
				neg[i] = 0
			}
		} else if arr[i] < 0 {
			if neg[i-1] > 0 {
				pos[i] = 1 + neg[i-1]
			} else {
				pos[i] = 0
			}
			neg[i] = 1 + pos[i-1]
		}
		ans = max(ans, pos[i])
	}
	return ans
}

func main() {
	fmt.Println(max_len_eff([]int{70, -18, 75, -72, -69, -84, 64, -65, 0, -82, 62, 54, -63, -85, 53, -60, -59, 29, 32, 59, -54,
		-29, -45, 0, -10, 22, 42, -37, -16, 0, -7, -76, -34, 37, -10, 2, -59, -24, 85, 45, -81, 56, 86}))
}
