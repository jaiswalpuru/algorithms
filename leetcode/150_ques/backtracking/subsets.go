/*
Generate all possible subsets of a given set.
2^n
*/

package main

import "fmt"

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

func main() {
	fmt.Println(power_set([]int{1, 2, 3}))
}
