package main

import "fmt"

func single_number(arr []int) int {
	n := len(arr)
	x := arr[0]
	for i := 1; i < n; i++ {
		x ^= arr[i]
	}
	return x
}

func main() {
	fmt.Println(single_number([]int{2, 2, 1}))
}
