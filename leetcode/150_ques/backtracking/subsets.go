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
	*res = append(*res, subset)
	for i := start; i < len(arr); i++ {
		temp := append(subset, arr[i])
		powerset(arr, temp, i+1, res)
		temp = temp[0 : len(temp)-1]
	}
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
