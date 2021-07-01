/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]


Example 2:

Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]


Example 3:

Input: matrix = [[1]]
Output: [[1]]
Example 4:

Input: matrix = [[1,2],[3,4]]
Output: [[3,1],[4,2]]
*/

package main

import "fmt"

func rotate_image(arr [][]int) [][]int {
	n := len(arr)

	//image of matrix
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := arr[i][j]
			arr[i][j] = arr[n-1-j][i]
			arr[n-1-j][i] = arr[n-1-i][n-1-j]
			arr[n-1-i][n-1-j] = arr[j][n-1-i]
			arr[j][n-1-i] = temp
		}
	}

	return arr
}

func main() {

	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(rotate_image(arr))

	arr = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	fmt.Println(rotate_image(arr))

	arr = [][]int{
		{1},
	}
	fmt.Println(rotate_image(arr))

	arr = [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(rotate_image(arr))
}
