package main

import "fmt"

func spiral_matrix_two(n int) [][]int {
	mat := make([][]int, n)
	for i:=0;i<n;i++{
		mat[i] = make([]int, n)
	}

	start := 1

	row,col,row_size,col_size := 0,0,n,n

	for start <= n*n {
		//fill the first row
		for i:=col; i<col_size; i++{
			mat[row][i] = start
			start++
		}

		//fill the last column
		row++
		for i:=row;i<row_size; i++{
			mat[i][col_size-1] = start
			start++
		}

		col_size--
		for i:=col_size-1;i>=col;i--{
			mat[row_size-1][i] = start
			start++
		}

		row_size--
		for i:=row_size-1; i>=row;i--{
			mat[i][col] = start
			start++
		}
		col++

	}

	return mat
}

func main () {
	fmt.Println(spiral_matrix_two(20))
}