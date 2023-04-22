package main

import "fmt"

func remove_elements(arr []int, val int) int {
	n := len(arr)

	k := 0
	for i := 0; i < n; i++ {
		if arr[i] != val {
			arr[k] = arr[i]
			k++
		}
	}
	return k
}

func main() {
	fmt.Println(remove_elements([]int{3, 2, 2, 3}, 3))
}
