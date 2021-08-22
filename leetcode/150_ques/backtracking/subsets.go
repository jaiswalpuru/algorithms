/*
Generate all possible subsets of a given set.
2^n
*/

package main

import (
	"fmt"
	"math"
)

func power_set(arr []int) [][]int {
	res := [][]int{}
	subset := []int{}

	powerset(arr, subset, 0, &res)
	return res
}

func powerset(arr []int, subset []int, start int, res *[][]int) {
	if idx == len(arr) {
		*res = append(*res, append([]int{}, subset...))
		return
	}
	power_set(arr, cur, start+1, res)
	power_set(arr, append(subset, arr[start]), start+1, res)
}

func power_set_non_recursive(arr []int) {
	n := int(math.Pow(2, float64(len(arr))))
	m := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i&(1<<j) == 0 {
				fmt.Print(arr[j], " ")
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println(power_set([]int{1, 2, 3}))
	power_set_non_recursive([]int{1, 2, 3})
}
