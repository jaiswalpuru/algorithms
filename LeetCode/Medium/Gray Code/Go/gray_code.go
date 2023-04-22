package main

import "fmt"

//---------------------backtrack-------------------------
func gray_code(n int) []int {
	res := []int{0}
	set := make(map[int]bool)
	set[0] = true
	_backtrack(&res, &set, n)
	return res
}

func _backtrack(res *[]int, set *map[int]bool, n int) bool {
	if len(*res) == 1<<n {
		return true
	}

	curr := (*res)[len(*res)-1]
	for i := 0; i < n; i++ {
		next := curr ^ (1 << i)
		if !(*set)[next] {
			(*set)[next] = true
			*res = append(*res, next)
			if _backtrack(res, set, n) {
				return true
			}
			(*set)[next] = false
			(*res) = (*res)[:len(*res)-1]
		}
	}
	return false
}

//-------------------backtrack----------------------------

//---------------------recursion---------------------------
func gray_code_recursion(n int) []int {
	res := []int{}
	_recur(&res, n)
	return res
}

func _recur(res *[]int, n int) {
	if n == 0 {
		*res = append(*res, 0)
		return
	}

	_recur(res, n-1)
	curr_len := len(*res)
	mask := 1 << (n - 1)
	for i := curr_len - 1; i >= 0; i-- {
		*res = append(*res, (*res)[i]|mask)
	}
}

func main() {
	fmt.Println(gray_code_recursion(16))
}
