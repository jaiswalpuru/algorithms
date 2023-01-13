package main

import "fmt"

/*
	Selection sort algorithm O(n^2)
	In this we find the index of
	smallest value and swap it the
	current index
*/

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

func main() {
	fmt.Println(sort([]int{5, 4, 3, 2, 1, -110110101}))
}
