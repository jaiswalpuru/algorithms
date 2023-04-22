package main

import (
	"fmt"
	"sort"
)

//---------------------two pointers approach--------------------------
func three_sum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			res = append(res, two_sum(nums, i)...)
		}
	}
	return res
}

func two_sum(nums []int, ind int) [][]int {
	i, j := ind+1, len(nums)-1
	res := [][]int{}
	for i < j {
		sum := nums[i] + nums[j] + nums[ind]
		if sum == 0 {
			res = append(res, []int{nums[ind], nums[i], nums[j]})
			i++
			j--
			for i < j && nums[i] == nums[i-1] {
				i++
			}
		} else if sum < 0 {
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
