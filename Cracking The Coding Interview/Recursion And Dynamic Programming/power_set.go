package main

import (
	"fmt"
	"math"
)

//---------------------recursion--------------------
func power_set(arr []int) [][]int {
	res := [][]int{}
	_recur(arr, 0, &res, []int{})
	return res
}

func _recur(arr []int, ind int, res *[][]int, set []int) {
	if ind == len(arr) {
		*res = append(*res, set)
		return
	}
	t := append(set, arr[ind])
	_recur(arr, ind+1, res, t)
	t = t[:len(t)-1]
	_recur(arr, ind+1, res, t)
}

//---------------------recursion--------------------

//---------------------using combinatorics--------------------
func power_set_bit(arr []int) [][]int {
	n := int(math.Pow(2, float64(len(arr))))
	res := [][]int{}
	for i := 0; i < n; i++ {
		set := get_set(arr, i)
		res = append(res, set)
	}
	return res
}

func get_set(arr []int, i int) []int {
	r := []int{}
	ind := 0
	for k := i; k > 0; k >>= 1 {
		if (k & 1) == 1 {
			r = append(r, arr[ind])
		}
		ind++
	}
	return r
}

//---------------------using combinatorics--------------------

func main() {
	fmt.Println(power_set_bit([]int{1, 2, 3}))
}
