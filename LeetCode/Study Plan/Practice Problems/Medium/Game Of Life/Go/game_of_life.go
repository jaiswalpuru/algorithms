package main

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

func is_safe(i, j, n, m int) bool {
	if i < 0 || j < 0 || i >= n || j >= m {
		return false
	}
	return true
}

func game_of_life(arr [][]int) {
	n, m := len(arr), len(arr[0])

	temp := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			live := 0
			for k := 0; k < 8; k++ {
				x, y := i+directions[k][0], j+directions[k][1]
				if is_safe(x, y, n, m) {
					if arr[x][y] == 1 {
						live++
					}
				}
			}

			if arr[i][j] == 1 && live < 2 {
				temp = append(temp, []int{i, j, 0})
			} else if arr[i][j] == 1 && live > 3 {
				temp = append(temp, []int{i, j, 0})
			} else if arr[i][j] == 0 && live == 3 {
				temp = append(temp, []int{i, j, 1})
			} else if arr[i][j] == 1 && (live == 2 || live == 3) {
				temp = append(temp, []int{i, j, 1})
			}
		}
	}

	for i := 0; i < len(temp); i++ {
		arr[temp[i][0]][temp[i][1]] = temp[i][2]
	}
}

func main() {
	game_of_life([][]int{
		{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
	})
}
