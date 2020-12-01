package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }

//I=1 all coins are showing head initially and I=2 all coins are tails initially
//Q=1 elephants need to guess total number of coins showing heads which Q=2 means
//guess is to be made for tails

func main() {

	var t int
	scanf("%d\n", &t)

	for i := 0; i < t; i++ {

		var g int
		scanf("%d\n", &g)

		for j := 0; j < g; j++ {
			var i, n, q int
			scanf("%d %d %d\n", &i, &n, &q)
			if n%2 == 0 || i == q {
				printf("%d\n", n/2)
			} else {
				printf("%d\n", n/2+1)
			}
		}

	}
	defer writer.Flush()
}
