package main

import "fmt"

func running_sum(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i] + arr[i-1]
	}
	return arr
}

func main() {
	fmt.Println(running_sum([]int{1, 2, 3, 4}))
}
