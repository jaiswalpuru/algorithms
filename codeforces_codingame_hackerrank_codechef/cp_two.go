package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func getInititalPos(arr [][]rune) (int, int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 'O' {
				return i, j
			}
		}
	}
	return 0, 0
}

func getRobotDir(arr [][]rune, x, y int, dir rune) rune {
	if dir == 'U' {
		if arr[x-1][y] == '#' {
			dir = 'R'
		}
	}
	if dir == 'D' {
		if arr[x+1][y] == '#' {
			dir = 'L'
		}
	}
	if dir == 'L' {
		if arr[x][y-1] == '#' {
			dir = 'U'
		}
	}
	if dir == 'R' {
		if arr[x][y+1] == '#' {
			dir = 'D'
		}
	}
	return dir
}

func moveRobot(arr [][]rune, x, y, n *int, dir *rune) {
	d := getRobotDir(arr, *x, *y, *dir)
	*dir = d
	if d == 'U' {
		if arr[*x-1][*y] != '#' {
			*x--
			*n--
		} else {
			d = getRobotDir(arr, *x, *y, 'R')
			moveRobot(arr, x, y, n, &d)
		}
	}
	if d == 'D' {
		if arr[*x+1][*y] != '#' {
			*x++
			*n--
		} else {
			d = getRobotDir(arr, *x, *y, 'L')
			moveRobot(arr, x, y, n, &d)
		}
	}
	if d == 'L' {
		if arr[*x][*y-1] != '#' {
			*y--
			*n--
		} else {
			d = getRobotDir(arr, *x, *y, 'U')
			moveRobot(arr, x, y, n, &d)
		}
	}
	if d == 'R' {
		if arr[*x][*y+1] != '#' {
			*y++
			*n--
		} else {
			d = getRobotDir(arr, *x, *y, 'D')
			moveRobot(arr, x, y, n, &d)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var w, h int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &w, &h)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	inp := [][]rune{}

	for i := 0; i < h; i++ {
		scanner.Scan()
		line := scanner.Text()
		inp = append(inp, []rune(line))
	}

	x, y := getInititalPos(inp)
	dir := 'U'
	f := false
	if n >= 100 {
		for {
			if n%2 == 0 {
				fmt.Println(n / 2)
				n = n / 2
				f = true
			} else {
				break
			}
		}
		if f {
			n = n * 2
		}
	}
	fmt.Println(n)
	for {
		moveRobot(inp, &x, &y, &n, &dir)
		if n <= 0 {
			break
		}
	}

	fmt.Println(y, x)
}
