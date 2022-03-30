package main

import "fmt"

func spiral_matrix(arr [][]int) []int {
	n, m := len(arr), len(arr[0])
	up, down, left, right := 0, n-1, 0, m-1

	res := []int{}
	for len(res) < n*m {
		for j := left; j <= right; j++ {
			res = append(res, arr[up][j])
		}

		for i := up + 1; i <= down; i++ {
			res = append(res, arr[i][right])
		}

		if up != down {
			for j := right - 1; j >= left; j-- {
				res = append(res, arr[down][j])
			}
		}

		if left != right {
			for i := down - 1; i > up; i-- {
				res = append(res, arr[i][left])
			}
		}

		left++
		down--
		up++
		right--
	}
	return res
}

func main() {
	fmt.Println(spiral_matrix([][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}))
}
