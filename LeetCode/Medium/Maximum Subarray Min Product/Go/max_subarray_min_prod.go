package main

import (
	"fmt"
	"math"
)

func max_subarray_min_prod(arr []int) int {
	n := len(arr)
	stack := [][]int{}
	prefix_sum := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefix_sum[i+1] = prefix_sum[i] + arr[i]
	}

	res := int(math.MinInt64)
	for i, val := range arr {
		curr := i
		for len(stack) > 0 && stack[len(stack)-1][1] > val {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			last_index, last_val := pop[0], pop[1]
			total := prefix_sum[i] - prefix_sum[last_index]
			res = max(res, total*last_val)
			curr = last_index
		}
		stack = append(stack, []int{curr, val})
	}

	//pop everything out from the stack using 0 as the sentinel value
	curr := len(arr)
	val := 0
	for len(stack) > 0 && stack[len(stack)-1][1] > val {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		last_index, last_val := pop[0], pop[1]
		total := prefix_sum[len(arr)] - prefix_sum[last_index]
		res = max(res, total*last_val)
		curr = last_index
	}
	stack = append(stack, []int{curr, val}) //note required
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(max_subarray_min_prod([]int{1, 2, 3, 2}))
}
