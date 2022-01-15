package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func count_submatrice(arr [][]int) int {
	n, m := len(arr), len(arr[0])

	res := 0
	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if arr[row][col] == 1 {
				if row == 0 || col == 0 {
					res += 1
				} else {
					val := min(arr[row-1][col], min(arr[row][col-1], arr[row-1][col-1])) + arr[row][col]
					arr[row][col] = val
					res += val
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(count_submatrice([][]int{
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 1, 1, 1},
	}))
}
