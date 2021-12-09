/*
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Example 2:
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

Example 3:
Input: intervals = [], newInterval = [5,7]
Output: [[5,7]]

Example 4:
Input: intervals = [[1,5]], newInterval = [2,3]
Output: [[1,5]]

Example 5:
Input: intervals = [[1,5]], newInterval = [2,7]
Output: [[1,7]]
*/

package main

import "fmt"

func intervals(interval [][]int, new_interval []int) [][]int {
	//interval is already sorted
	temp := [][]int{}
	res := [][]int{}
	n := len(interval)
	f := false
	if n == 0 {
		temp = append(temp, new_interval)
	}

	for i := 0; i < n; i++ {
		if new_interval[0] < interval[i][0] {
			f = true
			temp = append(temp, new_interval)
			temp = append(temp, interval[i])
		} else {
			temp = append(temp, interval[i])
		}
	}

	if !f {
		temp = append(temp, new_interval)
	}

	res = append(res, temp[0])
	n = len(temp)
	for i := 1; i < n; i++ {
		if res[len(res)-1][1] >= temp[i][0] {
			if res[len(res)-1][1] <= temp[i][1] {
				res[len(res)-1][1] = temp[i][1]
			}
		} else {
			res = append(res, temp[i])
		}
	}
	return res

}

func main() {
	fmt.Println(intervals([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	fmt.Println(intervals([][]int{}, []int{5, 7}))
	fmt.Println(intervals([][]int{{1, 5}}, []int{2, 7}))
}
