package main

import "fmt"

//-----------------------------------recursion---------------------------------------------
func maximum_score_from_performing_multiplication_operations(nums []int, mul []int) int {
	res := 0
	recur(nums, mul, 0, 0, &res)
	return res
}

func recur(nums []int, mul []int, i, curr_sum int, res *int) {
	if i == len(mul) {
		*res = max(*res, curr_sum)
		return
	}
	recur(nums[1:], mul, i+1, curr_sum+(nums[0]*mul[i]), res)
	recur(nums[:len(nums)-1], mul, i+1, curr_sum+(nums[len(nums)-1]*mul[i]), res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//-----------------------------------recursion---------------------------------------------

//-----------------------------------memoization---------------------------------------------(Runtime error)
func maximum_score_from_performing_multiplication_operations_memo(nums []int, mul []int) int {
	memo := make([][][]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([][]int, len(nums))
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = make([]int, len(mul))
			for k := 0; k < len(memo[i][j]); k++ {
				memo[i][j][k] = -1
			}
		}
	}
	return _memo(nums, mul, 0, len(nums)-1, 0, &memo)
}

func _memo(nums []int, mul []int, left, right, ind int, memo *[][][]int) int {
	if ind == len(mul) {
		return 0
	}
	if (*memo)[left][right][ind] != -1 {
		return (*memo)[left][right][ind]
	}
	l := nums[left]*mul[ind] + _memo(nums, mul, left+1, right, ind+1, memo)
	r := nums[right]*mul[ind] + _memo(nums, mul, left, right-1, ind+1, memo)
	(*memo)[left][right][ind] = max(l, r)
	return (*memo)[left][right][ind]
}

//-----------------------------------memoization---------------------------------------------

//-----------------------------------memoization--------------------------------------------- (runtime error)
func maximum_score_from_performing_multiplication_operations_memo_better(nums []int, mul []int) int {
	memo := make([][]int, len(nums)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(mul)+1)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}
	return _memo_better(nums, mul, 0, 0, &memo)
}

func _memo_better(nums, mul []int, op, left int, memo *[][]int) int {
	n := len(nums)
	if op == len(mul) {
		return 0
	}
	if (*memo)[op][left] != -1 {
		return (*memo)[op][left]
	}
	l := nums[left]*mul[op] + _memo_better(nums, mul, op+1, left+1, memo)
	r := nums[n-1-(op-left)]*mul[op] + _memo_better(nums, mul, op+1, left, memo)
	(*memo)[op][left] = max(l, r)
	return (*memo)[op][left]
}

//-----------------------------------memoization---------------------------------------------

//-----------------------------------memoization---------------------------------------------
func maximum_score_from_performing_multiplication_operations_memo_space_optimizied(nums []int, mul []int) int {
	n := len(nums)
	m := len(mul)
	dp := make([]int, m+1)
	for op := m - 1; op >= 0; op-- {
		next_row := make([]int, m+1)
		copy(next_row, dp)
		for left := op; left >= 0; left-- {
			dp[left] = max(mul[op]*nums[left]+next_row[left+1], mul[op]*nums[n-1-(op-left)]+next_row[left])
		}
	}
	return dp[0]
}

//-----------------------------------memoization---------------------------------------------

func main() {
	fmt.Println(maximum_score_from_performing_multiplication_operations_memo_space_optimizied([]int{1, 2, 3}, []int{3, 2, 1}))
}
