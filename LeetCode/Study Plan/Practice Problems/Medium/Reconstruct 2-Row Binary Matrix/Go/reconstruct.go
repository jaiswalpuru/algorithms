package main

import "fmt"

func reconstruct(upper, lower int, arr []int) [][]int {
	res := make([][]int, 2)
	res[0] = make([]int, len(arr))
	res[1] = make([]int, len(arr))

	m := len(arr)
	for i := 0; i < m; i++ {
		if arr[i] == 1 {
			if upper >= lower {
				res[0][i] = 1
				upper--
			} else {
				res[1][i] = 1
				lower--
			}
		} else if arr[i] == 2 {
			res[0][i] = 1
			res[1][i] = 1
			upper--
			lower--
		}
	}
	if lower != 0 || upper != 0 {
		return nil
	}
	return res
}

func main() {
	fmt.Println(reconstruct(5, 5, []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}))
}
