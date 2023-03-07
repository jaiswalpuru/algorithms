package main

import (
	"fmt"
	"sort"
)

func minimumTime(time []int, totalTrips int) int64 {
	sort.Ints(time)

	res := int64(-1)
	size := len(time)
	l := int64(1)
	r := int64(time[size-1] * totalTrips)

	for l <= r {
		mid := l + (r-l)/2
		if isPossible(mid, totalTrips, time) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return res
}

func isPossible(t int64, totalTrips int, time []int) bool {
	trips := int64(0)
	size := len(time)

	for i := 0; i < size; i++ {
		trips += t / int64(time[i])
	}

	return trips >= int64(totalTrips)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTime([]int{1, 2, 3}, 5))
}
