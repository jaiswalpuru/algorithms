package main

import "fmt"

func maximum_erasure_value(arr []int) int {
	ind_map := make(map[int]bool)
	l := 0
	sum := 0
	res := 0
	for i := 0; i < len(arr); i++ {
		for ind_map[arr[i]] {
			ind_map[arr[l]] = false
			sum -= arr[l]
			l++
		}
		sum += arr[i]
		ind_map[arr[i]] = true
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximum_erasure_value([]int{4, 2, 4, 5, 6}))
}
