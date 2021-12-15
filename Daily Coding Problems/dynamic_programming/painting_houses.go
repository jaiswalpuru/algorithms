/*
A builder is looking to build n houses that can be of k different colors. She has a goal of minimizing cost while
ensuring that no two neighboring houses are of the same color.

Given a n by k matrix where the entry at the ith row and jth column represents the cost to build the ith house
with the jth color, return the minimum cost required to achieve this goal.
*/

package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMinValInList(arr []int) int {
	minVal := int(math.MaxInt32)
	for i := 0; i < len(arr); i++ {
		minVal = min(minVal, arr[i])
	}
	return minVal
}

func min_cost_paint_houses(cost [][]int) int {
	k := len(cost[0])
	soln := make([][]int, 0)
	temp_cost := []int{}
	for i := 0; i < k; i++ {
		temp_cost = append(temp_cost, 0)
	}
	soln = append(soln, temp_cost)

	for r, row := range cost {
		row_cost, temp_cost := []int{}, []int{}
		for c, val := range row {
			for i := 0; i < k; i++ {
				if i != c {
					temp_cost = append(temp_cost, soln[r][i]+val)
				}
			}
			fmt.Println(temp_cost)
			row_cost = append(row_cost, getMinValInList(temp_cost))
			fmt.Println(row_cost)
		}
		soln = append(soln, row_cost)
	}
	return getMinValInList(soln[len(soln)-1])
}

//point of optimization as we are only looking at the last row we don't need to store all the values

func main() {
	cost := [][]int{
		{14, 2, 11}, {11, 14, 5}, {14, 3, 10},
	}

	fmt.Println(min_cost_paint_houses(cost))
}
