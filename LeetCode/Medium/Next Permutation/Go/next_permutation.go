package main

import (
	"fmt"
)

func next_permutation(arr []int) []int {
	n := len(arr)
	i, j := n-2, 0

	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}

	if i >= 0 {
		j = n - 1
		for arr[j] <= arr[i] {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	reverse(&arr, i+1)
	return arr
}

func reverse(arr *[]int, i int) {
	j, k := i, len(*arr)-1
	for j < k {
		(*arr)[j], (*arr)[k] = (*arr)[k], (*arr)[j]
		j++
		k--
	}
}

func main() {
	fmt.Println(next_permutation([]int{1, 2, 6, 7, 5}))
}
