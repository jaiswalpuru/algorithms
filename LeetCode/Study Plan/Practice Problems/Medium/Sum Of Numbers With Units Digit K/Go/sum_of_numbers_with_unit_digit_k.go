package main

import "fmt"
import "math"

//---------------------------brute force---------------------------------
func sum_of_numbers_with_unit_digit_k(num, k int) int {
	res := math.MaxInt64
	_recur(num, &res, k, 1, []int{}, 0)
	if res == math.MaxInt64 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func _recur(num int, res *int, k int, start int, arr []int, sum int) {
	if sum >= num {
		if sum == num {
			*res = min(*res, len(arr))
		}
		return
	}

	if start > num {
		return
	}

	if start%10 == k {
		if k == 0 {
			temp := append(arr, start)
			_recur(num, res, k, start, temp, start+sum)
			temp = temp[:len(temp)-1]
		} else {
			temp := append(arr, start)
			_recur(num, res, k, start, temp, start+sum)
			temp = temp[:len(temp)-1]
		}
	}
	_recur(num, res, k, start+1, arr, sum)
}

//-----------------------------------------------------------------------

//------------------------memoization---------------------------------
func sum_of_number_with_unit_digit_k_memo(num, k int) int {
	if num == 0 {
		return 0
	}

	arr := []int{}
	for i := 1; i <= num; i++ {
		if i%10 == k {
			arr = append(arr, i)
		}
	}

	n := len(arr)
	memo := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, num+1)
		for j := 0; j < num+1; j++ {
			memo[i][j] = -1
		}
	}
	_memo(num, n, &memo, arr)
	if memo[n][num] == math.MaxInt64-1 {
		return -1
	}
	return memo[n][num]

}

func _memo(num int, n int, memo *[][]int, arr []int) int {
	if num == 0 {
		return 0
	}

	if n == 0 {
		return math.MaxInt64 - 1
	}

	if (*memo)[n][num] != -1 {
		return (*memo)[n][num]
	}

	if arr[n-1] <= num {
		(*memo)[n][num] = min(1+_memo(num-arr[n-1], n, memo, arr), _memo(num, n-1, memo, arr))
	} else {
		(*memo)[n][num] = _memo(num, n-1, memo, arr)
	}

	return (*memo)[n][num]
}

//--------------------------------------------------------------------

func main() {
	fmt.Println(sum_of_number_with_unit_digit_k_memo(58, 9))
}
