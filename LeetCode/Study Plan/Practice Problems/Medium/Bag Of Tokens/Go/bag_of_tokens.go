package main

import (
	"fmt"
	"sort"
)

func bag_of_tokens(arr []int, power int) int {
	sort.Ints(arr)
	i, j := 0, len(arr)-1
	max_score := 0
	curr := 0
	for i <= j {
		if power >= arr[i] {
			curr++
			power -= arr[i]
			i++
		} else {
			if curr >= 1 {
				curr--
				power += arr[j]
				j--
			} else {
				break
			}
		}
		max_score = max(max_score, curr)
	}
	return max_score
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(bag_of_tokens([]int{100}, 50))
}
