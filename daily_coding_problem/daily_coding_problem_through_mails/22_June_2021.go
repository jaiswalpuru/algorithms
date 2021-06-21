/*
	Given a list of integers, return the largest product that can be made by multiplying any three integers.

	For example, if the list is [-10, -10, 5, 2], we should return 500, since that's -10 * -10 * 5.

	You can assume the list has at least three integers.
*/

package main

import (
	"fmt"
	"math"
	"sort"
)

const (
	INT_MIN = int(math.MinInt32)
	INT_MAX = int(math.MaxInt32)
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//get max prod without sorting
func largest_prod_without_sort(arr []int) int {

	max_one, max_two, max_three := INT_MIN, INT_MIN, INT_MIN
	min_one, min_two := INT_MAX, INT_MAX

	for i := 0; i < len(arr); i++ {
		if arr[i] > max_one {
			max_three = max_two
			max_two = max_one
			max_one = arr[i]
		} else if arr[i] > max_two {
			max_three = max_two
			max_two = arr[i]
		} else if arr[i] > max_three {
			max_three = arr[i]
		}

		//update first min and second min
		if arr[i] < min_one {
			min_two = min_one
			min_one = arr[i]
		} else if arr[i] < min_two {
			min_two = arr[i]
		}
	}
	return max(min_one*min_two*max_one, max_one*max_two*max_three)
}

//get the max prod using sorting O(nlogn)
func largest_prod_sort(arr []int) int {
	sort.Ints(arr)
	n := len(arr) - 1
	t1, t2 := arr[0]*arr[1]*arr[n], arr[n]*arr[n-1]*arr[n-2]
	return max(t1, t2)
}

func main() {
	fmt.Println(largest_prod_sort([]int{-10, -10, 5, 2}))
	fmt.Println(largest_prod_without_sort([]int{-10, -10, 5, 2}))
}
