package main

import "fmt"

func apply_operations_to_array(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		if arr[i] == arr[i+1] {
			arr[i] *= 2
			arr[i+1] = 0
		}
	}
	i, j := 0, 0
	for i < n && j < n {
		if arr[j] != 0 {
			arr[i] = arr[j]
			i++
			j++
		} else {
			j++
		}
	}
	for k := i; k < n; k++ {
		arr[k] = 0
	}
	return arr
}

func main() {
	fmt.Println(apply_operations_to_array([]int{1, 2, 2, 1, 1, 0}))
}
