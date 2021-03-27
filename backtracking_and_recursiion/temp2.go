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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var w, h, countX, countY int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &w, &h, &countX, &countY)

	inputX := make([]int, countX)
	for i := 0; i < countX; i++ {
		fmt.Scan(&inputX[i])
	}

	inputY := make([]int, countY)
	for i := 0; i < countY; i++ {
		fmt.Scan(&inputY[i])
	}
}
