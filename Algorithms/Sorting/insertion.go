package main

import "fmt"

/*
	Insertion sort : worst case O(n^2)
	this algorithm can be thought of as
	moving a wall from left to right and
	keeping the left of the wall sorted.
*/

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		temp := arr[i]
		j := i
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
	return arr
}

func main() {
	fmt.Println(sort([]int{5, 4, 3, 2, 1, -111101010}))
}
