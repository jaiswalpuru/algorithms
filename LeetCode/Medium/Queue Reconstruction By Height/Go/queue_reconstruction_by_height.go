package main

import "fmt"
import "sort"

func queue_reconstruction_by_height(arr [][]int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[i][1] < arr[j][1]
		}
		return arr[i][0] > arr[j][0]
	})

	res := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i]
		for j := arr[i][1]; j < i; j++ {
			res[j], res[i] = res[i], res[j]
		}
	}
	return res
}

func main() {
	fmt.Println(queue_reconstruction_by_height([][]int{
		{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1},
	}))
}
