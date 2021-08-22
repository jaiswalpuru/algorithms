/*
Given an array of integers temperatures represents the daily temperatures, return an array
answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature.
If there is no future day for which this is possible, keep answer[i] == 0 instead.

Example 1:
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Example 2:
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Example 3:
Input: temperatures = [30,60,90]
Output: [1,1,0]
*/

package main

import "fmt"

func daily_temp(temperatures []int) []int {
	n := len(temperatures)
	if n == 0 {
		return nil
	}
	days_wait := make([]int, n)
	type pair struct{ val, ind int }
	stack := []pair{}

	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > stack[len(stack)-1].val {
			days_wait[stack[len(stack)-1].ind] = i - stack[len(stack)-1].ind
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, pair{val: temperatures[i], ind: i})
	}

	return days_wait
}

func main() {
	fmt.Println(daily_temp([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(daily_temp([]int{30, 40, 50, 60}))
	fmt.Println(daily_temp([]int{30, 60, 90}))
}
