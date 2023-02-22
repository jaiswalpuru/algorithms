package main

import "fmt"

func shipWithinDays(weights []int, days int) int {
	maxWt := weights[0]
	totalWt := weights[0]
	size := len(weights)

	for i := 0; i < size; i++ {
		maxWt = max(maxWt, weights[i])
		totalWt += weights[i]
	}

	low := maxWt
	high := totalWt

	for low < high {
		mid := (low + high) / 2
		if isPossible(days, mid, weights) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func isPossible(days int, capacity int, weights []int) bool {
	totalDays := 1
	totalWt := 0
	size := len(weights)

	for i := 0; i < size; i++ {
		totalWt += weights[i]
		if totalWt > capacity {
			totalDays++
			totalWt = weights[i]
		}
	}

	return totalDays <= days
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
