/*
Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.

A falling path starts at any element in the first row and chooses the element in the next row that
is either directly below or diagonally left/right. Specifically, the next element from position (row, col)
will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).

Example 1:
Input: matrix = [[2,1,3],[6,5,4],[7,8,9]]
Output: 13
Explanation: There are two falling paths with a minimum sum underlined below:
[[2,1,3],      [[2,1,3],
 [6,5,4],       [6,5,4],
 [7,8,9]]       [7,8,9]]

 Example 2:
Input: matrix = [[-19,57],[-40,-5]]
Output: -59
Explanation: The falling path with a minimum sum is underlined below:
[[-19,57],
 [-40,-5]]

 Example 3:
Input: matrix = [[-48]]
Output: -48
*/

package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func min_falling_path_sum(arr [][]int) int {
	n := len(arr)
	for i := n - 2; i >= 0; i-- {
		arr[i][0] = arr[i][0] + min(arr[i+1][0], arr[i+1][1])
		for j := 1; j < n-1; j++ {
			arr[i][j] = arr[i][j] + min(arr[i+1][j-1], min(arr[i+1][j], arr[i+1][j+1]))
		}
		arr[i][n-1] = arr[i][n-1] + min(arr[i+1][n-1-1], arr[i+1][n-1])
	}
	for i := 1; i < n; i++ {
		arr[0][0] = min(arr[0][0], arr[0][i])
	}
	return arr[0][0]
}

func main() {
	fmt.Println(min_falling_path_sum([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))

	fmt.Println(min_falling_path_sum([][]int{
		{-19, 57},
		{-40, -5},
	}))

	fmt.Println(min_falling_path_sum([][]int{
		{-48},
	}))
}
