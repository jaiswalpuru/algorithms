package main

import "fmt"

func minimum_average_difference(arr []int) int {
	n := len(arr)
	left := make([]int, n)
	right := make([]int, n)
	left[0] = arr[0]
	right[n-1] = arr[n-1]
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + arr[i]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] + arr[i]
	}
	avg := right[0] / n
	ind := n - 1
	l, r := 1, n-1
	for i := 1; i < n; i++ {
		if avg == abs(right[i]/r-left[i-1]/l) {
			ind = min(ind, i-1)
		}
		if avg > abs(right[i]/r-left[i-1]/l) {
			ind = i - 1
			avg = abs(right[i]/r - left[i-1]/l)
		}
		l++
		r--
	}
	return ind
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minimum_average_difference([]int{4, 2, 0}))
}
