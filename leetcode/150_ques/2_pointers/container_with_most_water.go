/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0).
Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

Notice that you may not slant the container.

Example 1:
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:
Input: height = [1,1]
Output: 1

Example 3:
Input: height = [4,3,2,1,4]
Output: 16

Example 4:
Input: height = [1,2,1]
Output: 2
*/

package main

import "fmt"

func container_with_most_water(height []int) int {
	low, high := 0, len(height)-1
	area := 0

	for low <= high {
		if height[low] < height[high] {
			t := (high - low) * height[low]
			area = max(area, t)
			low++
		} else {
			t := (high - low) * height[high]
			area = max(area, t)
			high--
		}
	}

	return area
}

func main() {
	fmt.Println(container_with_most_water([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(container_with_most_water([]int{1, 1}))
	fmt.Println(container_with_most_water([]int{4, 3, 2, 1, 4}))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
