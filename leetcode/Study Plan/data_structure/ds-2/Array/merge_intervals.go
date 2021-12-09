/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.
*/

package main

import (
	"fmt"
	"sort"
)

type _int [][]int

func (a _int) Len() int           { return len(a) }
func (a _int) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a _int) Less(i, j int) bool { return a[i][0] < a[j][0] }

func merge_intervals(arr _int) [][]int {
	res := make([][]int, 0)
	n := len(arr)
	sort.Sort(arr)
	temp := []int{arr[0][0], arr[0][1]}
	for i := 1; i < n; i++ {
		if temp[1] >= arr[i][0] && temp[1] <= arr[i][1] {
			temp[1] = arr[i][1]
		} else if temp[1] >= arr[i][0] && temp[1] > arr[i][1] {
			continue
		} else {
			res = append(res, temp)
			temp = []int{arr[i][0], arr[i][1]}
		}
	}
	res = append(res, temp)
	return res
}

func main() {
	fmt.Println(merge_intervals(_int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge_intervals(_int{{1, 4}, {4, 5}}))
	fmt.Println(merge_intervals(_int{{1, 4}, {2, 3}}))
}
