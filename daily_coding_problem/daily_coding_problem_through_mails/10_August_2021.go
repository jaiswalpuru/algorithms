/*
Given a sorted list of integers, square the elements and give the output in sorted order.

For example, given [-9, -2, 0, 2, 3], return [0, 4, 4, 9, 81].
*/

package main

import "fmt"

func square_sort(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return nil
	}

	stack := []int{}
	res := []int{}
	for i := 0; i < n; i++ {
		if arr[i] < 0 {
			stack = append(stack, arr[i]*arr[i])
		} else {
			res = append(res, arr[i]*arr[i])
		}
	}
	ans := []int{}
	p, q := len(res), len(stack)
	i := 0
	for i < p && q > 0 {
		if res[i] > stack[q-1] {
			ans = append(ans, stack[q-1])
			q--
		} else if res[i] < stack[q-1] {
			ans = append(ans, res[i])
			i++
		} else {
			ans = append(ans, res[i], stack[q-1])
			i++
			q--
		}
	}

	for i := q - 1; i >= 0; i-- {
		ans = append(ans, stack[i])
	}
	return ans
}

func main() {
	fmt.Println(square_sort([]int{-9, -2, 0, 2, 3}))
}
