package main

import "fmt"
import "sort"

func sort_array_by_parity(arr []int) []int {
	sort.Slice(arr, func(i,j int)bool {
		return arr[i]%2 < arr[j]%2
	})
	return arr
}

func main() {
	fmt.Println(sort_array_by_parity([]int{3,1,2,4}))
}