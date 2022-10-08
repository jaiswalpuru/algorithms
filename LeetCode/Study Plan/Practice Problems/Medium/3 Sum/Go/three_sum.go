package main

import (
	"fmt"
	"sort"
)

//---------------------two pointers approach--------------------------
func three_sum(arr []int) [][]int {
	n := len(arr)
	sort.Ints(arr)
	res := [][]int{}
	for i := 0; i < n && arr[i] <= 0; i++ {
		if i == 0 || arr[i] != arr[i-1] {
			val := two_sum(arr, i)
			res = append(res, val...)
		}
	}
	return res
}

func two_sum(arr []int, cur_ind int) [][]int {
	i, j := cur_ind+1, len(arr)-1
	res := [][]int{}
	for i < j {
		val := arr[i] + arr[j] + arr[cur_ind]
		if val == 0 {
			if cur_ind != i && cur_ind != j {
				res = append(res, []int{arr[cur_ind], arr[i], arr[j]})
			}
			i++
			j--
			for i < j && arr[i] == arr[i-1] {
				i++
			}
		} else if val < 0 {
			i++
		} else {
			j--
		}
	}
	return res
}

//---------------------two pointers approach--------------------------

//---------------------hash set approach--------------------------
func three_sum_hash(arr []int) [][]int {
	n := len(arr)
	sort.Ints(arr)
	res := [][]int{}
	for i := 0; i < n && arr[i] <= 0; i++ {
		if i == 0 || arr[i] != arr[i-1] {
			val := two_sum_hash(arr, i)
			res = append(res, val...)
		}
	}
	return res
}

func two_sum_hash(arr []int, cur_ind int) [][]int {
	res := [][]int{}
	seen := make(map[int]bool)
	for j := cur_ind + 1; j < len(arr); j++ {
		complement := -arr[cur_ind] - arr[j]
		if seen[complement] {
			res = append(res, []int{arr[cur_ind], arr[j], complement})
			for j+1 < len(arr) && arr[j] == arr[j+1] {
				j++
			}
		}
		seen[arr[j]] = true
	}
	return res
}

//---------------------hash setapproach--------------------------

func main() {
	fmt.Println(three_sum_hash([]int{-1, 0, 1, 2, -1, -4}))
}
