package main

import "fmt"

func replace_elements_with_greater_elements_on_right_side(arr []int) []int {
	val := arr[len(arr)-1]
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > val {
			val, arr[i] = arr[i], val
		} else {
			arr[i] = val
		}
	}
	return arr
}

func main() {
	fmt.Println(replace_elements_with_greater_elements_on_right_side([]int{17, 18, 5, 4, 6, 1}))
}
