package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		n, m, t int
	)
	scanf("%d\n", &t)
	for ; t > 0; t-- {
		scanf("%d %d\n", &n, &m)
		a := make([]int, n)
		b := make([]int, m)
		sumA, sumB, swap := 0, 0, 0
		for i := 0; i < n; i++ {
			if i == n-1 {
				scanf("%d\n", &a[i])
				sumA += a[i]
				break
			}
			scanf("%d ", &a[i])
			sumA += a[i]
		}
		for i := 0; i < m; i++ {
			if i == m-1 {
				scanf("%d\n", &b[i])
				sumB += b[i]
				break
			}
			scanf("%d ", &b[i])
			sumB += b[i]
		}

		if sumA <= sumB {
			sort.Ints(a)
			sort.Ints(b)
			i, j := 0, m-1
			for (i < n) && (j > -1) {
				swap++
				sumA = (sumA - a[i]) + b[j]
				sumB = (sumB - b[j]) + a[i]
				a[i], b[j] = b[j], a[i]
				if sumA > sumB {
					break
				}
				i++
				j--
			}
		}
		if sumA > sumB {
			printf("%d\n", swap)
		} else {
			printf("%d\n", -1)
		}
	}
}
