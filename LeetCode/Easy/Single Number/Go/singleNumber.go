package main

import "fmt"

func singleNumber(arr []int) int {
	n := len(arr)
	x := arr[0]
	for i := 1; i < n; i++ {
		x ^= arr[i]
	}
	return x
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
