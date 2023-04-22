/*
You are given an integer array nums of length n.

Assume arrk to be an array obtained by rotating nums by k positions clock-wise. We define the rotation function
F on nums as follow:

F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1].
Return the maximum value of F(0), F(1), ..., F(n-1).

The test cases are generated so that the answer fits in a 32-bit integer.
*/

package main

import (
	"fmt"
)

func rotate_function(arr []int) int {
	n := len(arr)

	sum := 0
	f_0 := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		f_0 += i * arr[i]
	}

	max := f_0
	for i := 1; i < n; i++ {
		f_0 += sum - (n * arr[n-i])
		if max < f_0 {
			max = f_0
		}
	}

	return max
}

func main() {
	fmt.Println(rotate_function([]int{4, 3, 2, 6}))
	fmt.Println(rotate_function([]int{100}))
}
