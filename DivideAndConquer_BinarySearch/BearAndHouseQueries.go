package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	res    string
	a      int
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

//perform binary search to get all the sides which are to the right as the figure is symmetrical
//on origin

func isInside(x, y int) bool {
	printf("? %d %d\n", x, y)
	writer.Flush()
	fmt.Println(x, y)
	scanf("%s\n", &res)
	writer.Flush()
	if res == "YES" {
		return true
	}
	return false
}

func main() {
	defer writer.Flush()

	//side of the square
	squareX := -1 + sort.Search(1001, func(x int) bool { return !isInside(x, 0) })
	squareX = squareX * 2
	//fmt.Println(x)

	//get the y axis coordinates
	heightY := -1 + sort.Search(1001, func(y int) bool { return !isInside(0, y) })
	//fmt.Println(height)

	//get coordinates of triangle
	side := -1 + sort.Search(1001, func(x int) bool { return !isInside(x, squareX) })
	side = side * 2
	//fmt.Println(side)

	printf("! %d\n", (squareX*squareX)+(heightY-squareX)*side/2)
	writer.Flush()
}
