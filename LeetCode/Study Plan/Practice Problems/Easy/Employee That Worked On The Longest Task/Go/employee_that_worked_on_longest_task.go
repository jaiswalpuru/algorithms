package main

import (
	"fmt"
)

func employee_that_worked_on_longest_task(n int, arr [][]int) int {
	l := len(arr)
	res := arr[0][1] - 0
	ind := arr[0][0]
	for i := 1; i < l; i++ {
		if arr[i][1]-arr[i-1][1] > res {
			res = arr[i][1] - arr[i-1][1]
			ind = arr[i][0]
		}
		if arr[i][1]-arr[i-1][1] == res {
			ind = min(ind, arr[i][0])
		}
	}
	return ind
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(employee_that_worked_on_longest_task(26, [][]int{
		{1, 1}, {3, 7}, {2, 12}, {7, 17},
	}))
}
