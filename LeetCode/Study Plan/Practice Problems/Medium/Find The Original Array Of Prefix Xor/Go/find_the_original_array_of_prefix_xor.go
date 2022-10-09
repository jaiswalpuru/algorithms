package main

import "fmt"

func find_the_original_array_of_prefix_xor(arr []int) []int {
	n := len(arr)
	for i := n - 1; i >= 1; i-- {
		arr[i] = arr[i] ^ arr[i-1]
	}
	return arr
}

func main() {
	fmt.Println(find_the_original_array_of_prefix_xor([]int{5, 2, 0, 3, 1}))
}
