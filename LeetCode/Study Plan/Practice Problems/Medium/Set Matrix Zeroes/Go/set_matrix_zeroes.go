package main

func set_matrix_zeroes(arr [][]int) {
	mp := make(map[int][]int)
	n, m := len(arr), len(arr[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if arr[i][j] == 0 {
				mp[i] = append(mp[i], j)
			}
		}
	}
	for k, v := range mp {
		i, j := k, v
		for k := 0; k < m; k++ {
			arr[i][k] = 0
		}
		for k := 0; k < len(j); k++ {
			col := j[k]
			for l := 0; l < n; l++ {
				arr[l][col] = 0
			}
		}
	}
}

func main() {
	arr := [][]int{
		{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5},
	}
	set_matrix_zeroes(arr)
}
