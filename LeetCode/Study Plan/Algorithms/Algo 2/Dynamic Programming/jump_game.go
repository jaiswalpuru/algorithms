/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array
represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/

package main

import "fmt"

func jump_game(arr []int) bool {
	n := len(arr)

	for i := 1; i < n; i++ {
		if arr[i-1] > 0 {
			if arr[i-1] > arr[i] {
				arr[i] = arr[i-1] - 1
			}
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(jump_game([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump_game([]int{3, 2, 1, 0, 4}))
}
