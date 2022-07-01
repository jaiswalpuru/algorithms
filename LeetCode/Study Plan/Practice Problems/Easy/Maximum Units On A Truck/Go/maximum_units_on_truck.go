package main

import (
	"fmt"
	"sort"
)

func maximum_units_on_truck(arr [][]int, size int) int {
	res := 0
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] < arr[j][0]
		}
		return arr[i][1] > arr[j][1]
	})
	for i := 0; i < len(arr); i++ {
		if size == 0 {
			break
		}
		if size >= arr[i][0] {
			size -= arr[i][0]
			res += arr[i][0] * arr[i][1]
		} else {
			res += size * arr[i][1]
			size = 0
		}
	}
	return res
}

func main() {
	fmt.Println(maximum_units_on_truck([][]int{
		{1, 3}, {2, 2}, {3, 1},
	}, 4))
}
