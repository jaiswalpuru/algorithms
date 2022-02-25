package main

import "fmt"

func permute(arr []int) [][]int {
	res := [][]int{}
	_permute(arr, &res, 0)
	return res
}

func _permute(arr []int, res *[][]int, start int) {
	if start == len(arr) {
		*res = append(*res, append([]int{}, arr...))
		return
	}

	visited := make(map[int]struct{})
	for i := start; i < len(arr); i++ {
		if _, ok := visited[arr[i]]; ok {
			continue
		}
		arr[i], arr[start] = arr[start], arr[i]
		_permute(arr, res, start+1)
		arr[i], arr[start] = arr[start], arr[i]
		visited[arr[i]] = struct{}{}
	}
}

func main() {
	fmt.Println(permute([]int{1, 1, 2}))
}
