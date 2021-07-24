/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's,
and return the matrix.

You must do it in place.

Example 1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Constraints:

m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1


Follow up:

A straightforward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/
package main

import (
	"fmt"
)

func set_matrix_zeroes(arr [][]int) [][]int {
	n := len(arr)
	m := len(arr[0])
	is_col := false

	for i := 0; i < n; i++ {
		if arr[i][0] == 0 {
			is_col = true
		}

		for j := 1; j < m; j++ {
			if arr[i][j] == 0 {
				arr[0][j] = 0
				arr[i][0] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if arr[i][0] == 0 || arr[0][j] == 0 {
				arr[i][j] = 0
			}
		}
	}

	//see if the first row needs to set to zero as well
	if arr[0][0] == 0 {
		for j := 0; j < m; j++ {
			arr[0][j] = 0
		}
	}

	if is_col {
		for i := 0; i < n; i++ {
			arr[i][0] = 0
		}
	}
	return arr
}

func main() {
	fmt.Println(set_matrix_zeroes([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}))

	fmt.Println(set_matrix_zeroes([][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}))
}
