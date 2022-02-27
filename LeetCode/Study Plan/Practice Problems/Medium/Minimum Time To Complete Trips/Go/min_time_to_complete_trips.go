package main

import (
	"fmt"
	"sort"
)

func min_time(time []int, total_trips int) int64 {
	sort.Ints(time)
	n := len(time)
	max_time := time[n-1] * total_trips
	low, high := 0, max_time
	for low < high {
		mid_time := (low + high) / 2

		sum := 0
		for i := 0; i < n; i++ {
			sum += (mid_time / time[i])
		}
		if sum >= total_trips {
			high = mid_time
		} else {
			low = mid_time + 1
		}
	}
	return int64(high)
}

func main() {
	fmt.Println(min_time([]int{1, 2, 3}, 5))
}
