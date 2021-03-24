package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	scanf  = func(f string, args ...interface{}) { fmt.Fscanf(reader, f, args...) }
	printf = func(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
)

type pair struct {
	x, y int
}

func main() {
	defer writer.Flush()

	var t int
	scanf("%d\n", &t)

	for ; t > 0; t-- {
		var (
			n, k, x, y int
		)
		scanf("%d %d %d %d\n", &n, &k, &x, &y)
		a1 := []pair{{n, y + n - x}, {y + n - x, n}, {0, x - y}, {x - y, 0}} //x>y
		a2 := []pair{{x + n - y, n}, {n, n - y + x}, {y - x, 0}, {0, y - x}} //y>x
		k = (k - 1) % 4
		if x > y {
			//fmt.Println("x>y")
			//(n,y+n-x),(y+n-x,n),(0,x-y),(x-y,0)
			printf("%d %d\n", a1[k].x, a1[k].y)
		} else if x < y {
			//(x+n-y,n),(n,n-y+x),(y-x,0),(0,y-x)
			printf("%d %d\n", a2[k].x, a2[k].y)
		} else if x == y {
			printf("%d %d\n", n, n)
		}
	}
}
