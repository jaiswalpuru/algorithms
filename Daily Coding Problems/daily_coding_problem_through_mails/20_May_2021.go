/*
	The power set of a set is the set of all its subsets. Write a function that, given a set,
	generates its power set.

	For example, given the set {1, 2, 3}, it should return {{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}.

	You may also use a list or array to represent a set.
*/

package main

import (
	"fmt"
	"math"
)

var r = [][]int{}

func all_subset(arr *[]int, subset []int, cur_ind int) {
	r = append(r, subset)
	for i := cur_ind; i < len(*arr); i++ {
		subset = append(subset, (*arr)[i])
		all_subset(arr, subset, i+1)
		subset = subset[:len(subset)-1]
	}
}

func all_subset_iterative(arr []int) [][]int {
	power_set_size := int(math.Pow(2, float64(len(arr))))
	res := make([][]int, power_set_size)
	for i := 0; i < power_set_size; i++ {
		temp := []int{}
		for j := 0; j < len(arr); j++ {
			if (i & (1 << j)) != 0 {
				temp = append(temp, arr[j])
			}
		}
		res[i] = append(res[i], temp...)

	}
	return res
}

func main() {
	arr := []int{1, 2, 3}
	subset := make([]int, 0)
	all_subset(&arr, subset, 0)
	fmt.Println(r)
	res := all_subset_iterative(arr)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
