package main

import "fmt"

func min_max(arr []int) int {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}
	res := make([]int, n/2)
	for i := 0; i < n/2; i++ {
		if i%2 == 0 {
			res[i] = min(arr[2*i], arr[2*i+1])
		} else {
			res[i] = max(arr[2*i], arr[2*i+1])
		}
	}
	val := min_max(res)
	return val
}

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

func main() {
	fmt.Println(min_max([]int{1, 3, 5, 2, 4, 8, 2, 2}))
}
