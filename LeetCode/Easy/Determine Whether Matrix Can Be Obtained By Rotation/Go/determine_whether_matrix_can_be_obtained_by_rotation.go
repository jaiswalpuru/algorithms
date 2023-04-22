package main

import "fmt"

func determine_whether_matrix_can_be_obtained_by_rotation(mat, target [][]int) bool {
	n, m := len(mat), len(mat[0])
	flag := true
	for k := 0; k < 4; k++ {

		//transpose
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if i < j {
					mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
				}
			}
		}
		//reflex
		for i := 0; i < n; i++ {
			for j := 0; j < m/2; j++ {
				mat[i][j], mat[i][m-j-1] = mat[i][m-1-j], mat[i][j]
			}
		}

		//compare
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if mat[i][j] != target[i][j] {
					return false
				}
			}
		}
		if flag {
			return true
		} else {
			flag = true
		}
	}
	return false
}

func main() {
	fmt.Println(determine_whether_matrix_can_be_obtained_by_rotation([][]int{
		{0, 1}, {1, 0},
	}, [][]int{
		{1, 0}, {0, 1},
	}))
}
