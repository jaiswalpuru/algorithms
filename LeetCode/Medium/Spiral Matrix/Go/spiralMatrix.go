package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	rowSize := len(matrix)
	colSize := len(matrix[0])
	up := 0
	down := rowSize - 1
	left := 0
	right := colSize - 1

	spiralArr := []int{}
	for len(spiralArr) < rowSize*colSize {
		for j := left; j <= right; j++ {
			spiralArr = append(spiralArr, matrix[up][j])
		}

		for i := up + 1; i <= down; i++ {
			spiralArr = append(spiralArr, matrix[i][right])
		}

		if up != down {
			for j := right - 1; j >= left; j-- {
				spiralArr = append(spiralArr, matrix[down][j])
			}
		}

		if left != right {
			for i := down - 1; i > up; i-- {
				spiralArr = append(spiralArr, matrix[i][left])
			}
		}

		up++
		down--
		left++
		right--
	}
	return spiralArr
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}))
}
