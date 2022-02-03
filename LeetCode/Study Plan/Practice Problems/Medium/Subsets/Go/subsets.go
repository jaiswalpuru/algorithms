package main

import "fmt"

func subsets(arr []int) [][]int {
	res := make([][]int, 0)
	_subsets(arr, 0, &res, []int{})
	return res
}

func _subsets(arr []int, ind int, res *[][]int, set []int) {
	if ind >= len(arr) {
		*res = append(*res, append([]int{}, set...))
		return
	}

	temp := append(set, arr[ind])
	_subsets(arr, ind+1, res, temp)
	temp = temp[:len(temp)-1]
	_subsets(arr, ind+1, res, temp)
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
