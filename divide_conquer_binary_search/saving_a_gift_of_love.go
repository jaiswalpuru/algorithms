package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
)

func main() {

	var (
		t int //test cases
		x int // the distance of byteland from chef's town

	)
	scanf("%d\n", &t)

	for ; t > 0; t-- {

	}
}
