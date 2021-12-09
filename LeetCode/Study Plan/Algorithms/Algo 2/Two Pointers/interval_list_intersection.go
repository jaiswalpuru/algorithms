/*
You are given two lists of closed intervals, firstList and secondList, where firstList[i] = [starti, endi] and
secondList[j] = [startj, endj]. Each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

A closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.

The intersection of two closed intervals is a set of real numbers that are either empty or represented as a closed interval.
For example, the intersection of [1, 3] and [2, 4] is [2, 3].
*/

package main

import (
	"fmt"
)

type Pair struct{ start, end int }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func interval_intersection(a []Pair, b []Pair) []Pair {
	i, j := 0, 0
	n, m := len(a), len(b)
	res := make([]Pair, 0)
	for i < n && j < m {
		low := max(a[i].start, b[j].start)
		high := min(a[i].end, b[j].end)
		if low <= high {
			res = append(res, Pair{low, high})
		}

		if a[i].end < b[j].end {
			i++
		} else {
			j++
		}
	}
	return res
}

func main() {
	fmt.Println(interval_intersection([]Pair{{0, 2}, {5, 10}, {13, 23}, {24, 25}}, []Pair{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))
	fmt.Println(interval_intersection([]Pair{{4, 11}}, []Pair{{1, 2}, {8, 11}, {12, 13}, {14, 15}, {17, 19}}))
	fmt.Println(interval_intersection([]Pair{{5, 10}}, []Pair{{3, 10}}))
	fmt.Println(interval_intersection([]Pair{{3, 5}, {9, 20}}, []Pair{{4, 5}, {7, 10}, {11, 12}, {14, 15}, {16, 20}}))
}
