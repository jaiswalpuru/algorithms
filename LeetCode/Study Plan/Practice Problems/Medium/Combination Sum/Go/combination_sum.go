package main

import "fmt"

func combination_sum(arr []int, target int) [][]int {
	res := [][]int{}
	_combiantion_sum(arr, target, []int{}, 0, &res)
	return res
}

func _combiantion_sum(arr []int, target int, set []int, ind int, res *[][]int) {
	if ind >= len(arr) {
		if target == 0 {
			*res = append(*res, append([]int{}, set...))
		}
		return
	}

	if arr[ind] <= target {
		temp := append(set, arr[ind])
		_combiantion_sum(arr, target-arr[ind], temp, ind, res)
		temp = temp[:len(temp)-1]
	}
	_combiantion_sum(arr, target, set, ind+1, res)

}

func main() {
	fmt.Println(combination_sum([]int{2, 3, 6, 7}, 7))
}
