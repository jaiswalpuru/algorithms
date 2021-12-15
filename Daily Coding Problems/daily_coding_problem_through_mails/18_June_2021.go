/*
	Given a N by M matrix of numbers, print out the matrix in a clockwise spiral.

	For example, given the following matrix:

	[[1,  2,  3,  4,  5],
	 [6,  7,  8,  9,  10],
	 [11, 12, 13, 14, 15],
	 [16, 17, 18, 19, 20]]
	You should print out the following:

	1
	2
	3
	4
	5
	10
	15
	20
	19
	18
	17
	16
	11
	6
	7
	8
	9
	14
	13
	12
*/

package main

import "fmt"

func spiral(m, n int, arr [][]int) {
	k := 0 //starting row index
	l := 0 //starting col index
	//m -> ending row index
	//n -> ending col index

	for k < m && l < n {

		//first row and from the remaining rows
		for i := l; i < n; i++ {
			fmt.Println(arr[k][i])
		}
		k++

		//last column from the remaining cols
		for i := k; i < m; i++ {
			fmt.Println(arr[i][n-1])
		}
		n--

		//print last row from the remaining rows
		if k < m {
			for i := n - 1; i >= l; i-- {
				fmt.Println(arr[m-1][i])
			}
			m--
		}

		//print the first col from the remaining cols
		if l < n {
			for i := m - 1; i >= k; i-- {
				fmt.Println(arr[i][l])
			}
			l++
		}
	}

}

func main() {
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}
	spiral(len(arr), len(arr[0]), arr)
}
