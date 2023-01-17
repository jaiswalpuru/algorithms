package main

import (
	"fmt"
	"math"
)

/**
 * // This is the ArrayReader's API interface.
 * // You should not implement it, or speculate about its implementation
 * type ArrayReader struct {
 * }
 *
 * func (this *ArrayReader) get(index int) int {}
 */

func search(reader ArrayReader, target int) int {
	l, r := 0, int(1e4)
	for l <= r {
		mid := l + (r-l)/2
		val := reader.get(mid)
		if val == math.MaxInt32 {
			r = mid - 1
		} else {
			fmt.Println(l, r)
			if val == target {
				return mid
			}
			if val > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
