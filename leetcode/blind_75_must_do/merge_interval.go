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

type _int [][]int

func (a _int) Len() int           { return len(a) }
func (a _int) Less(i, j int) bool { return a[i][0] < a[j][0] }
func (a _int) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func merge_intervals(arr _int) _int {
	sort.Sort(arr)
	res := [][]int{}
	res = append(res, arr[0])
	n := arr.Len()
	for i := 1; i < n; i++ {
		if res[len(res)-1][1] >= arr[i][0] {
			if res[len(res)-1][1] <= arr[i][1] {
				res[len(res)-1][1] = arr[i][1]
			}
		} else {
			res = append(res, arr[i])
		}
	}
	return res
}

func main() {
	fmt.Println(merge_intervals(_int{
		{1, 3},
		{2, 6},
		{15, 18},
		{8, 10},
	}))
	fmt.Println(merge_intervals(_int{
		{1, 4}, {2, 3},
	}))
}
