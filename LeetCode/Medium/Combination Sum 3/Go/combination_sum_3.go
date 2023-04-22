package main

import "fmt"

func combination_sum(k, n int) [][]int {
	res := [][]int{}
	_combination_sum(k, n, &res, []int{}, 1, 0)
	return res
}

func _combination_sum(k, n int, res *[][]int, temp []int, start int, sum int) {
	if len(temp) == k {
		if sum == n {
			*res = append(*res, append([]int{}, temp...))
		}
		return
	}

	for i := start; i <= 9; i++ {
		if sum+i <= n {
			set := append(temp, i)
			_combination_sum(k, n, res, set, i+1, sum+i)
			set = set[:len(set)-1]
		}
	}
}

func main() {
	fmt.Println(combination_sum(3, 7))
}
