package main

var (
	row  = make([][]int, 10)
	col  = make([][]int, 10)
	cell = make([][]int, 10)
)

func init() {
	for i := 0; i < 10; i++ {
		row[i] = make([]int, 10)
		col[i] = make([]int, 10)
		cell[i] = make([]int, 10)
	}
}

func checkForEmptyCell(arr [9][9]int) (int, int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if arr[i][j] == 0 {
				return i, j
			}
		}
	}
	return 0, 0
}

func isSafe(arr [9][9]int, i, j, val int) bool {

	//row check
	for k := 0; k <= 9; k++ {
		if arr[i][k] == val {
			return false
		}
	}

	//col check
	for k := 0; k <= 9; k++ {
		if arr[k][j] == val {
			return false
		}
	}

	//3 x 3 cell check
	startRow, startCol := i-i%3, j-j%3
	for k := 0; k < 3; k++ {
		for q := 0; q < 3; q++ {
			if arr[k+startRow][q+startCol] == val {
				return false
			}
		}
	}
	return true
}

func sudoku(arr [9][9]int) bool {
	var i, j int
	i, j = checkForEmptyCell(arr)
	for k := 1; k <= 9; k++ {
		if isSafe(arr, i, j, k) {
			arr[i][j] = k
			if sudoku(arr) {
				return true
			}
			arr[i][j] = 0
		}
	}
	return false
}

func main() {

}
