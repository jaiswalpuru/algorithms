/*
Generate power set of a set with the set having duplicates but power set should not have duplicates
*/

package main

import (
	"fmt"
	"sort"
)

func power_set(arr []int) [][]int {
	res := [][]int{}
	subset := []int{}
	sort.Ints(arr)
	backtrack(arr, subset, 0, &res)
	return res
}

func backtrack(arr []int, subset []int, start int, res *[][]int) {
	*res = append(*res, append([]int{}, subset...))
	for i := start; i < len(arr); i++ {
		if i > start && arr[i] == arr[i-1] {
			continue
		}
		temp := append(subset, arr[i])
		backtrack(arr, temp, i+1, res)
		temp = temp[0 : len(temp)-1]
	}
}

func main() {
	fmt.Println(power_set([]int{1, 2, 2}))
}
