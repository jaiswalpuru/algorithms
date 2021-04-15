package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	matrix             [][]rune
	bearMode           bool     = false
	inverted           bool     = false
	curDir             rune     = 'S'
	startX, startY     int      = -1, -1
	t1x, t1y, t2x, t2y int      = -1, -1, -1, -1
	path               []string = []string{}
	canReach           int      = 100 * 100
)

func isSafe(x, y int) bool {
	if matrix[x][y] == '#' {
		return false
	} else if matrix[x][y] == 'X' && bearMode {
		matrix[x][y] = ' '
		return true
	} else if matrix[x][y] == 'X' && !bearMode {
		return false
	} else {
		return true
	}
}

func updateNextState() {
	if !inverted {
		if isSafe(startX+1, startY) {
			startX++
			curDir = 'S'
			path = append(path, "SOUTH")

		} else if isSafe(startX, startY+1) {
			startY++
			curDir = 'E'
			path = append(path, "EAST")

		} else if isSafe(startX-1, startY) {
			startX--
			curDir = 'N'
			path = append(path, "NORTH")

		} else if isSafe(startX, startY-1) {
			startY--
			curDir = 'W'
			path = append(path, "WEST")

		}
	} else {
		if isSafe(startX, startY-1) {
			startY--
			curDir = 'W'
			path = append(path, "WEST")

		} else if isSafe(startX-1, startY) {
			startX--
			curDir = 'N'
			path = append(path, "NORTH")

		} else if isSafe(startX, startY+1) {
			startY++
			curDir = 'E'
			path = append(path, "EAST")

		} else if isSafe(startX+1, startY) {
			startX++
			curDir = 'S'
			path = append(path, "SOUTH")

		}
	}
}

func stateMachine() {
	for {
		if canReach < 0 {
			fmt.Println("LOOP")
			break
		}
		canReach--
		if matrix[startX][startY] == '$' {
			for i := 0; i < len(path); i++ {
				fmt.Println(path[i])
			}
			break
		}

		if matrix[startX][startY] == 'S' {
			curDir = 'S'
		} else if matrix[startX][startY] == 'E' {
			curDir = 'E'
		} else if matrix[startX][startY] == 'W' {
			curDir = 'W'
		} else if matrix[startX][startY] == 'N' {
			curDir = 'N'
		} else if matrix[startX][startY] == 'B' {
			bearMode = !bearMode
		} else if matrix[startX][startY] == 'I' {
			inverted = !inverted
		} else if matrix[startX][startY] == 'T' {
			if t1x == startX && t1y == startY {
				startX, startY = t2x, t2y
			} else {
				startX, startY = t1x, t1y
			}
		}

		if curDir == 'S' {
			if isSafe(startX+1, startY) {
				path = append(path, "SOUTH")
				startX++
			} else {
				updateNextState()
			}
		} else if curDir == 'E' {
			if isSafe(startX, startY+1) {
				path = append(path, "EAST")
				startY++
			} else {
				updateNextState()
			}
		} else if curDir == 'W' {
			if isSafe(startX, startY-1) {
				path = append(path, "WEST")
				startY--
			} else {
				updateNextState()
			}
		} else if curDir == 'N' {
			if isSafe(startX-1, startY) {
				path = append(path, "NORTH")
				startX--
			} else {
				updateNextState()
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L, C int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L, &C)
	matrix = make([][]rune, L)
	for i := 0; i < L; i++ {
		matrix[i] = make([]rune, C)
		scanner.Scan()
		row := scanner.Text()
		rowRune := []rune(row)
		for j := 0; j < len(rowRune); j++ {
			matrix[i][j] = rowRune[j]
			if matrix[i][j] == '@' {
				startX, startY = i, j
			}
			if matrix[i][j] == 'T' && t1x == -1 {
				t1x, t1y = i, j
			} else if matrix[i][j] == 'T' && t2x == -1 {
				t2x, t2y = i, j
			}
		}
	}
	stateMachine()
}
