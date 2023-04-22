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

//----------------------brute force---------------------------
func trapping_rain_water_brute_force(height []int) int {
	ans := 0
	n := len(height)
	for i := 0; i < n; i++ {
		left_max, right_max := 0, 0
		for j := i; j >= 0; j-- {
			left_max = max(left_max, height[j])
		}
		for j := i; j < n; j++ {
			right_max = max(right_max, height[j])
		}
		ans += (min(left_max, right_max) - height[i])
	}
	return ans
}

//----------------------brute force---------------------------

//----------------------O(n) solution---------------------------
func trapping_rain_water(height []int) int {
	n := len(height)
	res := 0
	left := make([]int, n)
	right := make([]int, n)
	left[0], right[n-1] = height[0], height[n-1]
	for i := 1; i < n; i++ {
		left[i] = max(height[i], left[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = max(height[i], right[i+1])
	}
	for i := 0; i < n; i++ {
		res += min(left[i], right[i]) - height[i]
	}
	return res
}

//----------------------O(n) solution---------------------------

func main() {
	fmt.Println(trapping_rain_water([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
