package main

import (
	"fmt"
	"sort"
)

func matchsticks_to_square(arr []int) bool {
	n := len(arr)
	perimeter := 0
	for i := 0; i < n; i++ {
		perimeter += arr[i]
	}
	possible_side := perimeter / 4
	if possible_side*4 != perimeter {
		return false
	}

	sort.Ints(arr)
	//reverse
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return _recur(arr, [4]int{}, 0, possible_side)
}

func _recur(arr []int, temp [4]int, ind int, possible_side int) bool {
	if ind == len(arr) {
		return temp[0] == temp[1] && temp[1] == temp[2] && temp[2] == temp[3]
	}

	curr := arr[ind]
	for i := 0; i < 4; i++ {
		if curr+temp[i] <= possible_side {
			temp[i] += curr
			if _recur(arr, temp, ind+1, possible_side) {
				return true
			}
			temp[i] -= curr
		}
	}
	return false
}

func main() {
	fmt.Println(matchsticks_to_square([]int{1, 1, 2, 2, 2}))
}
