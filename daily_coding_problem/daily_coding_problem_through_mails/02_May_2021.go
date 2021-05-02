/*
A builder is looking to build a row of N houses that can be of K different colors.
He has a goal of minimizing cost while ensuring that no two neighboring houses are of the same color.

Given an N by K matrix where the nth row and kth column represents the cost to build the nth house with kth color,
return the minimum cost which achieves this goal.
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

func minCost(arr []int, j int) int {
	temp := make([]int, len(arr)-1)
	//cloning so that it doesn't deep copy
	for i, k := 0, 0; i < len(arr); i++ {
		if i != j {
			temp[k] = arr[i]
			k++
		}
	}
	return getMinValInList(temp)
}

func main() {
	//N=4 houses
	cost := [][]int{{2, 5}, {1, 6}, {2, 7}, {4, 3}}

	totalCost := [][]int{{0, 0}, {0, 0}, {0, 0}, {0, 0}}
	numColors := len(cost[0])

	for i := 0; i < 4; i++ {
		if i == 0 {
			for j := 0; j < numColors; j++ {
				totalCost[i][j] = cost[i][j]
			}
		} else {
			for j := 0; j < numColors; j++ {
				costTemp := cost[i][j]
				totalCost[i][j] = costTemp + minCost(totalCost[i-1], j)
			}
		}
	}
	fmt.Println("Total minimum cost of the houses : ", getMinValInList(totalCost[len(totalCost)-1]))
}
