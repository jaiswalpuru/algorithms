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
	total := n + m
	half := (total + 1) / 2
	l, r := 0, n-1
	for {
		la := (l + r) >> 1  //left partition of a
		lb := half - la - 2 //left partition of b (-2 because of 0 zero indexing)
		a_left, a_right := math.MinInt64, math.MaxInt64
		b_left, b_right := math.MinInt64, math.MaxInt64
		if la >= 0 {
			a_left = a[la]
		}
		if la+1 < n {
			a_right = a[la+1]
		}
		if lb >= 0 {
			b_left = b[lb]
		}
		if lb+1 < m {
			b_right = b[lb+1]
		}
		if a_left <= b_right && b_left <= a_right {
			if total%2 == 0 {
				return float64((max(a_left, b_left) + min(a_right, b_right))) / 2
			}
			return float64(max(a_left, b_left))

		} else if a_left > b_right {
			r = la - 1
		} else {
			l = la + 1
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
