package main

import "fmt"

func check_if_number_is_majority_element_in_sorted_array(arr []int, val int) bool {
	cnt := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == val {
			cnt++
		}
	}
	return cnt > n/2
}

func main() {
	fmt.Println(check_if_number_is_majority_element_in_sorted_array([]int{2, 4, 5, 5, 5, 5, 5, 6, 6}, 5))
}
