package main

import "fmt"
import "strconv"
import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//---------------------------------brute force----------------------------------
func longest_binary_subsequence_less_than_or_equal_to_k(str string, k int) int {
	res := math.MinInt64
	_recur(str, k, &res, "", 0)
	return res
}

func _recur(str string, k int, res *int, t string, ind int) {
	if ind == len(str) {
		val, _ := strconv.ParseInt(t, 2, 64)
		if int(val) <= k {
			*res = max(*res, len(t))
		}
		return
	}

	temp := t + string(str[ind])
	_recur(str, k, res, temp, ind+1)
	temp = temp[:len(temp)-1]
	_recur(str, k, res, temp, ind+1)
}

//-----------------------------------------------------------------------------

//--------------------------------efficient approach---------------------------
func longest_binary_subsequence_less_than_or_equal_to_k_eff(str string, k int) int {
	n := len(str)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	return _recur_memo(str, n-1, 0, k, &memo)
}

func _recur_memo(str string, ind, size, sum int, memo *[][]int) int {
	if ind < 0 {
		return 0
	}

	if (*memo)[ind][size] != -1 {
		return (*memo)[ind][size]
	}

	not_take := _recur_memo(str, ind-1, size, sum, memo)
	take := 0
	if sum-int(str[ind]-'0')*int(math.Pow(2, float64(size))) >= 0 {
		take = 1 + _recur_memo(str, ind-1, size+1, sum-int(str[ind]-'0')*int(math.Pow(2, float64(size))), memo)
	}
	(*memo)[ind][size] = max(not_take, take)
	return (*memo)[ind][size]

}

//--------------------------------efficient approach---------------------------

func main() {
	fmt.Println(longest_binary_subsequence_less_than_or_equal_to_k_eff("111100010000011101001110001111000000001011101111111110111000011111011000010101110100110110001111001001011001010011010000011111101001101000000101101001110110000111101011000101", 11713332))
}
