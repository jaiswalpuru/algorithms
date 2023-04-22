package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	maxPile := math.MinInt64
	size := len(piles)

	for i := 0; i < size; i++ {
		maxPile = max(maxPile, piles[i])
	}

	l := 1
	r := maxPile

	for l < r {
		mid := l + (r-l)/2
		if isPossible(h, mid, piles) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}

func isPossible(h, banana int, piles []int) bool {
	hrs := 0
	size := len(piles)

	for i := 0; i < size; i++ {
		hrs += int(math.Ceil(float64(piles[i]) / float64(banana)))
	}

	return hrs <= h
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
