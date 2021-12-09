/*
Given an m x n matrix, return all elements of the matrix in spiral order.

Example 1:
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2:
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

package main

import "fmt"

func print_spiral_matrix(arr [][]int) {
	s, t, u, v := 0, 0, len(arr)-1, len(arr[0])-1
	for s <= u && t <= v {
		//first row
		for i := s; i <= v; i++ {
			fmt.Print(arr[s][i])
		}
		s++
		//down
		for i := s; i <= u; i++ {
			fmt.Print(arr[i][v])
		}
		v--

		//print the last rows from the remaining rows
		if s <= u {
			for i := v; i >= 0; i-- {
				fmt.Print(arr[u][i])
			}
			u--
		}

		//print the first column from remaining column
		if t <= v {
			for i := u; i >= s; i-- {
				fmt.Print(arr[i][t])
			}
			t++
		}
	}
	fmt.Println()
}

func main() {
	print_spiral_matrix([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
}
