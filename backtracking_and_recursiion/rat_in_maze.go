package main

//some global var
var (
	ROW int
	COL int
)

func isValid(mat [][]int, x, y int) bool {
	if (x <= ROW && y <= COL) && mat[x][y] != 1 {
		return true
	}
	return false
}

func findPathInMaze(mat [][]int, x, y int, path [][]int) bool {

	if x == ROW-1 && y == COL-1 {
		path[x][y] = 1
		return true
	}

	if isValid(mat, x, y) {
		path[x][y] = 1

		if findPathInMaze(mat, x+1, y, path) {
			return true
		}

		if findPathInMaze(mat, x, y+1, path) {
			return true
		}

		path[x][y] = 0
		return false
	}

	return false
}

func main() {

}
