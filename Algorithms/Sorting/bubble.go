package main

import "fmt"

/*
	Bubble sort worst case : O(N^2)
*/

func sort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ { //n-1-i ignoring the comparisons which have already been placed correctly
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(sort([]int{5, 4, 3, 2, 1}))
}
