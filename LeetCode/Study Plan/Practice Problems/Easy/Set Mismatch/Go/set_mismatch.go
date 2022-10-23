package main

import "fmt"

func set_mismatch(arr []int) []int {
	n := len(arr)
	res := make([]int, 2)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[arr[i]]++
	}
	for i := 1; i <= n; i++ {
		if mp[i] > 1 {
			res[0] = i
		}
		if mp[i] == 0 {
			res[1] = i
		}
	}
	return res
}

func main() {
	fmt.Println(set_mismatch([]int{1, 2, 2, 4}))
}
