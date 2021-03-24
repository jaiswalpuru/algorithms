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

func main() {
	defer writer.Flush()
	var (
		n, d, k, t int
	)

	scanf("%d\n", &t)

	for ; t > 0; t-- {
		scanf("%d %d %d\n", &n, &k, &d)
		arr := make([]int, n)
		sum := 0

		for i := 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &arr[i])
				break
			}
			scanf("%d ", &arr[i])
		}

		for i := 0; i < n; i++ {
			sum += arr[i]
		}

		totalProb := sum / k

		if totalProb >= d {
			printf("%d\n", d)
		} else {
			printf("%d\n", totalProb)
		}

	}
}
