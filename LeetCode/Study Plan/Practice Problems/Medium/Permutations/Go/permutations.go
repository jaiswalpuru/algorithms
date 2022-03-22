package main

import "fmt"

func permutations(arr []int) [][]int {
	res := [][]int{}
	visited := make(map[int]bool)
	_permutations(arr, &res, []int{}, &visited)
	return res
}

func _permutations(arr []int, res *[][]int, set []int, visited *map[int]bool) {
	if len(set) == len(arr) {
		*res = append(*res, append([]int{}, set...))
		return
	}

	for i:=0; i<len(arr); i++ {
		if !(*visited)[i] {
			temp := append(set, arr[i])
			(*visited)[i] = true
			_permutations(arr, res, temp, visited)
			temp = temp[:len(temp)-1]
			(*visited)[i] = false
		}
	}
}

func permutation_two(arr []int) [][]int {
	res := [][]int{}
	_permutations_two(&arr, &res, 0)
	return res
}

func _permutations_two(arr *[]int, res *[][]int, start int) {
	if len(*arr) == start {
		*res = append(*res, append([]int{}, *arr...))
		return
	}

	for i:= start; i<len(*arr); i++ {
		(*arr)[i], (*arr)[start] = (*arr)[start], (*arr)[i]
		_permutations_two(arr, res, start+1)
		(*arr)[i], (*arr)[start] = (*arr)[start], (*arr)[i]
	}
}

func main() {
	fmt.Println(permutation_two([]int{1, 2, 3}))
}