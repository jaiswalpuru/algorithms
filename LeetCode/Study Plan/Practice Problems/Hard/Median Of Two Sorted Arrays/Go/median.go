package main

import (
	"fmt"
	"math"
)

//-----------------------------------brute force merge the two arrays-------------
func median_bf(a, b []int) float64 {
	m, n := len(a), len(b)
	c := make([]int, m+n)

	i, j, k := 0, 0, 0
	for i < m && j < n {
		if a[i] > b[j] {
			c[k] = b[j]
			j++
		} else {
			c[k] = a[i]
			i++
		}
		k++
	}

	if i < m {
		for p := i; p < m; p++ {
			c[k] = a[p]
			k++
		}
	}

	if j < n {
		for p := j; p < n; p++ {
			c[k] = b[p]
			k++
		}
	}

	if len(c)%2 == 0 {
		return float64((c[len(c)/2] + c[len(c)/2-1])) / 2
	}

	return float64(c[len(c)/2])
}

//--------------------------------------------------------------------------------

//-----------------------------------efficient approach using binary search-----------------------------------
func median(a, b []int) float64 {
	if len(a) > len(b) {
		return median(b, a)
	}

	n, m := len(a), len(b)
	l, h := 0, n

	for l <= h {
		cut1 := (l + h) / 2
		cut2 := (n+m+1)/2 - cut1
		var left1, left2, right1, right2 int

		if cut1 == 0 {
			left1 = math.MinInt64
		} else {
			left1 = a[cut1-1]
		}
		if cut2 == 0 {
			left2 = math.MinInt64
		} else {
			left2 = b[cut2-1]
		}

		if cut1 == n {
			right1 = math.MaxInt64
		} else {
			right1 = a[cut1]
		}

		if cut2 == m {
			right2 = math.MaxInt64
		} else {
			right2 = b[cut2]
		}

		if left1 <= right2 && left2 <= right1 {
			if (n+m)%2 == 0 {
				return float64(max(left1, left2)+min(right1, right2)) / 2
			} else {
				return float64(max(left1, left2))
			}
		} else if left1 > right2 {
			h = cut1 - 1
		} else {
			l = cut1 + 1
		}
	}
	return 0
}

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

//-----------------------------------------------------------------------------------------

func main() {
	fmt.Println(median([]int{1, 3}, []int{2}))
}
