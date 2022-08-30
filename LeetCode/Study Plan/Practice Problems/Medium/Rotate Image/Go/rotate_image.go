package main

import "fmt"

func rotate_image(arr [][]int) {
	//transpose + reflex = will rotate the matrix by 90
	//transpose
	n, m := len(arr), len(arr[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i < j {
				arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m/2; j++ {
			arr[i][j], arr[i][n-j-1] = arr[i][n-j-1], arr[i][j]
		}
	}
	fmt.Println(arr)
}

func main() {
	rotate_image([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
}
