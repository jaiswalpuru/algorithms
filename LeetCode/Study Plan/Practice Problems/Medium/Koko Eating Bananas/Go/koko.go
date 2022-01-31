package main

import (
	"fmt"
	"math"
)

//---------------------------brute force---------------------------
func eat_all(piles []int, h int) int {
	n := len(piles)
	t := 1

	for {
		temp := h
		for i := 0; i < n; i++ {
			temp -= int(math.Ceil(float64(piles[i]) / float64(t)))
			if temp < 0 {
				break
			}
		}
		if temp < 0 {
			t++
		}
		if temp >= 0 {
			break
		}
	}
	return t
}

//-----------------------------------------------------------------

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//------------------------------bsearch-------------------------------
func eat_all_bsearch(piles []int, h int) int {
	n := len(piles)
	max_val := math.MinInt64
	for i := 0; i < n; i++ {
		max_val = max(max_val, piles[i])
	}
	l, r := 1, max_val

	for l < r {
		mid := (l + r) / 2

		hrs := 0
		for i := 0; i < n; i++ {
			hrs += int(math.Ceil(float64(piles[i]) / float64(mid)))
		}

		if hrs <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

//------------------------------bsearch--------------------------------

func main() {
	fmt.Println(eat_all_bsearch([]int{3, 6, 7, 11}, 8))
}
