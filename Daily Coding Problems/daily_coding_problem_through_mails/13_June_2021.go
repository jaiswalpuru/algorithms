/*
	Given a multiset of integers, return whether it can be partitioned into two subsets whose sums are the same.

	For example, given the multiset {15, 5, 20, 10, 35, 15, 10}, it would return true,
	since we can split it up into {15, 5, 10, 15, 10} and {20, 35}, which both add up to 55.

	Given the multiset {15, 5, 20, 10, 35}, it would return false,
	since we can't split it up into two subsets that add up to the same sum.
*/

package main

import (
	"fmt"
	"sort"
)

func split(arr []int) ([]int, []int) {
	sort.Ints(arr)

	left, right := 0, len(arr)-1
	sp_one, sp_two := []int{}, []int{}

	sum_one, sum_two := arr[left], arr[right]
	sp_one, sp_two = append(sp_one, arr[left]), append(sp_two, arr[right])

	if sum_one == sum_two {
		return sp_one, sp_two
	} else {
		left++
		right--
		for left <= right {
			if sum_one < sum_two {
				sum_one += arr[left]
				sp_one = append(sp_one, arr[left])
				left++
			} else if sum_one > sum_two {
				sum_one += arr[left]
				sum_two += arr[right]
				sp_one = append(sp_one, arr[left])
				sp_two = append(sp_two, arr[right])
				left++
				right--
			}
		}
		if sum_one == sum_two {
			return sp_one, sp_two
		} else {
			return nil, nil
		}
	}
}

func main() {
	arr := []int{15, 5, 20, 10, 35, 15, 10}
	//arr := []int{15, 5, 20, 10, 35}
	sp_one, sp_two := split(arr)
	if sp_one == nil || sp_two == nil {
		fmt.Println(false)
	} else {
		fmt.Println(sp_one)
		fmt.Println(sp_two)
	}
}
