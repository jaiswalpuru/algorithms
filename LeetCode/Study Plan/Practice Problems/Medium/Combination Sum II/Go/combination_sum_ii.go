package main

import "fmt"
import "sort"

func combination_sum_II(arr []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(arr)
	_combination_sum_II(arr, &res, []int{}, target, 0)
	return res
}

func _combination_sum_II(arr []int, res *[][]int, set []int, target int, ind int) {
	if target == 0 {
		*res = append(*res, append([]int{}, set...))
		return
	}

	for i:=ind;i<len(arr);i++{
		if i>ind && arr[i] == arr[i-1] {
			continue
		}

		if arr[i] <= target {
			temp := append(set, arr[i])
			_combination_sum_II(arr, res, temp, target-arr[i], i+1)
			temp = temp[:len(temp)-1]
		}
	}
}

func main() {
	fmt.Println(combination_sum_II([]int{10,1,2,7,6,1,5}, 8))
}