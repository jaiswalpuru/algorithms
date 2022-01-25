package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func count_submatrices(arr [][]int) int {
	n, m := len(arr), len(arr[0])
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			if arr[i][j] == 1 {
				arr[i][j] = arr[i][j] + arr[i][j-1]
			}
		}
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res += arr[i][j]

			local_min := arr[i][j]
			for k := i + 1; k < n; k++ {
				if arr[k][j] == 0 {
					break
				}

				local_min = min(local_min, arr[k][j])
				res += local_min
			}
		}
	}
	return res
}

func main() {
	fmt.Println(count_submatrices([][]int{
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 0},
	}))
}
