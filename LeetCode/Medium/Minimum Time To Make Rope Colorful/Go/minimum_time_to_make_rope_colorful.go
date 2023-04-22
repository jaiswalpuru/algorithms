package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimum_time_to_make_rope_colorful(colors string, needed_time []int) int {
	n := len(colors)
	cnt := 1
	max_time := needed_time[0]
	sum := needed_time[0]
	min_time := 0
	for i := 1; i < n; i++ {
		if colors[i] == colors[i-1] {
			max_time = max(max_time, needed_time[i])
			cnt++
			sum += needed_time[i]
		} else {
			if cnt > 1 {
				min_time += sum - max_time
			}
			cnt = 1
			max_time = needed_time[i]
			sum = needed_time[i]
		}
	}
	if cnt > 1 {
		min_time += sum - max_time
	}
	return min_time
}

func main() {
	fmt.Println(minimum_time_to_make_rope_colorful("abaac", []int{1, 2, 3, 4, 5}))
}
