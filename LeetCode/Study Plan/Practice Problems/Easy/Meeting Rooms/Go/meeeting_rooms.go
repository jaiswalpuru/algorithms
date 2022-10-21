package main

import (
	"fmt"
	"sort"
)

func meeting_rooms(arr [][]int) bool {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] < arr[j][0]
	})
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i][0] < arr[i-1][1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(meeting_rooms([][]int{
		{0, 30}, {5, 10}, {15, 20},
	}))
}
