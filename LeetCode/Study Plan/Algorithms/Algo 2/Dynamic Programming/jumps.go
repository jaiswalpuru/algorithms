/*
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

You can assume that you can always reach the last index.
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jumps(arr []int) int {
	curr := 0
	n := len(arr)
	farthest, jumps := 0, 0
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+arr[i])
		if i == curr {
			jumps += 1
			curr = farthest
		}
	}
	return jumps
}

func main() {
	fmt.Println(jumps([]int{2, 3, 0, 1, 4}))
	fmt.Println(jumps([]int{1, 2}))
}
