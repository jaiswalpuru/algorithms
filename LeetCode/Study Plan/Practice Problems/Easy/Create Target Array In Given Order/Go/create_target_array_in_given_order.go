package main

import "fmt"

func create_target_array_in_given_order(arr []int, index []int) []int {
	n := len(arr)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		val := res[index[i]]
		res[index[i]] = arr[i]
		for j := index[i] + 1; j < n; j++ {
			t := res[j]
			res[j] = val
			val = t
		}
	}
	return res
}

func main() {
	fmt.Println(create_target_array_in_given_order([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}
