package main

import "fmt"

func combination(n int, k int) [][]int {
	res := [][]int{}
	_combination(n, k, &res, 1, []int{})
	return res
}

func _combination(n int, k int, res *[][]int, ind int, set []int) {
	if ind == n+1 {
		if k == len(set) {
			*res = append(*res, append([]int{}, set...))
		}
		return
	}

	temp := append(set, ind)
	_combination(n, k, res, ind+1, temp)
	temp = temp[:len(temp)-1]
	_combination(n, k, res, ind+1, temp)
}

func main() {
	fmt.Println(combination(4, 2))
}
