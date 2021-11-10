/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0).
Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

Notice that you may not slant the container.
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func container_with_most_water(arr []int) int {
	n := len(arr)
	i, j := 0, n-1
	max_water := 0
	for i < j {
		max_water = max(max_water, min(arr[i], arr[j])*(j-i))
		if arr[i] < arr[j] {
			i++
		} else {
			j--
		}
	}
	return max_water
}

func main() {
	fmt.Println(container_with_most_water([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
