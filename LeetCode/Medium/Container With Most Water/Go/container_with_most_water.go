package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func container_with_most_water(arr []int) int {
	i, j := 0, len(arr)-1

	max_area := math.MinInt64

	for i < j {
		if arr[i] > arr[j] {
			max_area = max(max_area, arr[j]*abs(j-i))
			j--
		} else {
			max_area = max(max_area, arr[i]*abs(j-i))
			i++
		}
	}
	return max_area
}

func main() {
	fmt.Println(container_with_most_water([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
