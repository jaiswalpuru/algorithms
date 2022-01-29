package main

import "fmt"

func non_decreasing_array(arr []int) bool {
	cnt := 0
	n := len(arr)

	for i := 1; i < n; i++ {
		if arr[i-1] > arr[i] {
			if cnt == 1 {
				return false
			}

			cnt++
			if i < 2 || arr[i-2] <= arr[i] {
				arr[i-1] = arr[i]
			} else {
				arr[i] = arr[i-1]
			}
		}
	}
	return true
}

func main() {
	fmt.Println(non_decreasing_array([]int{4, 2, 3}))
}
