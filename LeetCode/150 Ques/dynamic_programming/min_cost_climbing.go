/*
You are given an integer array cost where cost[i] is the cost of ith step on a staircase.
Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index 0, or the step with index 1.

Return the minimum cost to reach the top of the floor.

Example 1:
Input: cost = [10,15,20]
Output: 15
Explanation: Cheapest is: start on cost[1], pay that cost, and go to the top.

Example 2:
Input: cost = [1,100,1,1,1,100,1,1,100,1]
Output: 6
Explanation: Cheapest is: start on cost[0], and only step on 1s, skipping cost[3].
*/

package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func min_cost_climbing(arr []int) int {
	n := len(arr)
	for i := 2; i < n; i++ {
		arr[i] += min(arr[i-1], arr[i-2])
	}
	return min(arr[n-1], arr[n-2])
}

func main() {
	fmt.Println(min_cost_climbing([]int{10, 15, 20}))
	fmt.Println(min_cost_climbing([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
