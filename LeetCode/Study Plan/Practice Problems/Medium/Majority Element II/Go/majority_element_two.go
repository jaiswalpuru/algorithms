package main

import "fmt"

func majority_element_two(arr []int) []int {
	n := len(arr)
	res := []int{}
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		cnt[arr[i]]++
	}
	for k, v := range cnt {
		if v > n/3 {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	fmt.Println(majority_element_two([]int{3, 2, 3}))
}
