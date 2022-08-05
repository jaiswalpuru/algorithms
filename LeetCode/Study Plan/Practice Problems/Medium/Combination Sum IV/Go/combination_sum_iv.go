package main

import "fmt"

//------------------------brute force ---------------------------------------
func combination_sum_iv(arr []int, target int) int {
	res := 0
	_recur(arr, target, 0, []int{}, &res)
	return res
}

func _recur(arr []int, target int, sum int, temp []int, res *int) {
	if sum > target {
		return
	}
	if sum == target {
		*res++
		return
	}

	for i := 0; i < len(arr); i++ {
		if sum+arr[i] > target {
			continue
		}
		t := sum + arr[i]
		set := append(temp, arr[i])
		_recur(arr, target, t, set, res)
	}
}

//-----------------------brute force------------------------------------------

//-----------------------memoization------------------------------------------
func combination_sum_iv_memo(arr []int, target int) int {
	memo := make(map[int]int)
	return _memo(arr, target, &memo)
}

func _memo(arr []int, target int, memo *map[int]int) int {
	if target == 0 {
		return 1
	}

	if _, ok := (*memo)[target]; ok {
		return (*memo)[target]
	}

	res := 0
	for i := 0; i < len(arr); i++ {
		if target-arr[i] < 0 {
			continue
		}
		res += _memo(arr, target-arr[i], memo)
	}
	(*memo)[target] = res
	return (*memo)[target]
}

//-----------------------memoization------------------------------------------

func main() {
	fmt.Println(combination_sum_iv_memo([]int{1, 2, 3}, 4))
}
