/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].


Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
*/

package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	start, end int
}

type P []Pair

func (p P) Len() int           { return len(p) }
func (p P) Less(i, j int) bool { return p[i].end <= p[j].start }
func (p P) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func merge_interval(a [][]int) P {
	res := P{}
	n := len(a)
	data := P{}
	for i := 0; i < n; i++ {
		data = append(data, Pair{start: a[i][0], end: a[i][1]})
	}

	sort.Sort(data)
	//fmt.Println(data)
	temp := data[0]
	for i := 1; i < n; i++ {
		if temp.end >= data[i].start {
			temp.end = data[i].end
		} else {
			res = append(res, temp)
			temp = data[i]
		}
		if i == n-1 {
			res = append(res, temp)
		}
	}

	return res
}

func main() {
	fmt.Println(merge_interval([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
	fmt.Println(merge_interval([][]int{
		{1, 4}, {4, 5},
	}))
}
